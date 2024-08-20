package constants

// PreNetworkConfigScriptWithNmstatectl script runs on hosts before network manager service starts in order to apply
// user's provided network configuration on the host.
// If the user provides static network configuration, the network config files will be stored in directory
// /etc/assisted/network in the following structure:
// /etc/assisted/network/
//
//		+-- host1
//		|      +--- *.yml
//		|      +--- mac_interface.ini
//		+-- host2
//		      +--- *.yml
//		      +--- mac_interface.ini
//	 1. *.yml - files provided by the user with additional embedded captures to make it nmpolicy
//	 2. mac_interface.ini - the file contains mapping of mac-address to logical interface name.
//	    There are two usages for the file:
//	      1. Map logical interface name to MAC Address of the host. The logical interface name is a
//	         name provided by the user for the interface. It will be replaced by execute nmstate apply with the
//	         actual network interface name.
//	      2. Identify if a host directory belongs to the current host by matching a MAC Address
//	         from the mapping file with host network interfaces.
//
// Applying the network configuration of each host will be done by:
//  1. Associate the current host with its matching hostX directories. The association will be done by
//     matching host's mac addresses with those in mac_interface.ini.
//  2. Replace logical interface name with the interface name as set on the host by nmstate
//  3. Copy the nmconnection files to the target directory. That will be
//     '/etc/NetworkManager/system-connections' when the script is part of the full ISO and
//     '/etc/coreos-firstboot-network' when the script is part of the minimal ISO.
const PreNetworkConfigScriptWithNmstatectl = `#!/bin/bash

PATH_PREFIX=${PATH_PREFIX:=''}

# The directory that contains nm config files of all nodes
NM_CONFIG_DIR=${PATH_PREFIX}/etc/assisted/network
MAC_NIC_MAPPING_FILE=mac_interface.ini
NMSTATECTL=${PATH_PREFIX}/usr/bin/nmstatectl
NMSTATE_DIR=${PATH_PREFIX}/etc/nmstate

mkdir -p ${NMSTATE_DIR}

if [[ ! -d "$NM_CONFIG_DIR" ]]
then
  echo "Error (exiting): Expected to find the directory $NM_CONFIG_DIR on the host but this was not present."
  exit 0
fi

# A map of host mac addresses to interface names
declare -A host_macs_to_hw_iface

# Find destination directory based on ISO mode
if [[ -f ${PATH_PREFIX}/etc/initrd-release ]]; then
  ETC_NETWORK_MANAGER="${PATH_PREFIX}/etc/coreos-firstboot-network"
else
  ETC_NETWORK_MANAGER="${PATH_PREFIX}/etc/NetworkManager/system-connections"
fi
echo "Info: ETC_NETWORK_MANAGER was set to $ETC_NETWORK_MANAGER"

# Create a map of host mac addresses to their network interfaces
function map_host_macs_to_interfaces() {
  SYS_CLASS_NET_DIR="${PATH_PREFIX}/sys/class/net"
  for nic in $( ls $SYS_CLASS_NET_DIR )
  do
    mac=$(cat $SYS_CLASS_NET_DIR/$nic/address | tr '[:lower:]' '[:upper:]')
    [[ -n "$mac" ]] && host_macs_to_hw_iface[$mac]=$nic
  done
}

function process_host_directories_by_mac_address {
  pattern="$(echo -n ${!host_macs_to_hw_iface[@]} | sed 's/  */|/g')"
  if [[ -z "$pattern" ]] ; then
    return
  fi
  
  for host_src_dir in $(ls -1 -d ${NM_CONFIG_DIR}/host* || echo)
  do
    mapping_file="${host_src_dir}/${MAC_NIC_MAPPING_FILE}"
    if [[ ! -f "$mapping_file" ]]
    then
      echo "Warning: Mapping file $mapping_file is missing. Skipping on directory $host_src_dir"
      continue
    fi

    if [[ -z "$(ls -A $host_src_dir/*.yml)" ]]
    then
      echo "Warning: Host directory does not contain any yml files, skipping"
      continue
    fi

    if grep -q -i -E "$pattern" $mapping_file; then
      echo "Info: Found host directory: $(basename $host_src_dir) , copying configuration"
      
      cp "${host_src_dir}"/*.yml ${NMSTATE_DIR}
      ${NMSTATECTL} service

      # in full iso the nmstate service already put the keyfiles under /etc/NetworkManager/system-connections/
      # in minimal ISO only, we need to create /etc/coreos-firstboot-network and copy the files from /etc/NetworkManager/system-connections/ to /etc/coreos-firstboot-network
      if [[ -f ${PATH_PREFIX}/etc/initrd-release ]]; then
        mkdir -p ${ETC_NETWORK_MANAGER}
        cp /etc/NetworkManager/system-connections/* ${ETC_NETWORK_MANAGER}
        /usr/sbin/coreos-copy-firstboot-network
      fi

    fi
  done
}


echo "PreNetworkConfig Start"

# Get the mac to host nic mapping from local machine
map_host_macs_to_interfaces

process_host_directories_by_mac_address

echo "PreNetworkConfig End"
`

const MinimalISONetworkConfigServiceNmstatectl = `
[Unit]
Description=Assisted static network config
DefaultDependencies=no
After=nm-initrd.service systemd-udev-settle.service

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/usr/local/bin/pre-network-manager-config.sh
`
