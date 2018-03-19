package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVCenterDataCenter() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVCenterDataCenterRead,
        Schema: map[string]*schema.Schema{
            "filter": dataSourceFiltersSchema(),
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vrs_configuration_time_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "v_require_nuage_metadata": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "managed_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data_dns1": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data_dns2": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "data_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "datapath_sync_timeout": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "secondary_data_uplink_dhcp_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "secondary_data_uplink_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "secondary_data_uplink_interface": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "secondary_data_uplink_mtu": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "secondary_data_uplink_primary_controller": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "secondary_data_uplink_secondary_controller": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "secondary_data_uplink_underlay_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "secondary_nuage_controller": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "deleted_from_vcenter": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "memory_size_in_gb": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "remote_syslog_server_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "remote_syslog_server_port": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "remote_syslog_server_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "generic_split_activation": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "separate_data_network": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_mirror_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "metadata_server_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "metadata_server_listen_port": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "metadata_server_port": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "metadata_service_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "network_uplink_interface": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_uplink_interface_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_uplink_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "network_uplink_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "revertive_controller_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "revertive_timer": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "nfs_log_server": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nfs_mount_path": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mgmt_dns1": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mgmt_dns2": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mgmt_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mgmt_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dhcp_relay_server": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mirror_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "disable_gro_on_datapath": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "disable_lro_on_datapath": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "site_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "allow_data_dhcp": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "allow_mgmt_dhcp": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "flow_eviction_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "vm_network_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enable_vrs_resource_reservation": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "configured_metrics_push_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "portgroup_metadata": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "nova_client_version": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "nova_identity_url_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_auth_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_endpoint": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_tenant": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_service_username": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_metadata_shared_secret": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_os_keystone_username": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_project_domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_project_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_region_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nova_user_domain_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "upgrade_package_password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "upgrade_package_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "upgrade_package_username": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "upgrade_script_time_limit": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "cpu_count": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "primary_data_uplink_underlay_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "primary_nuage_controller": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vrs_password": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vrs_user_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_vcenter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "static_route": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "static_route_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "static_route_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ntp_server1": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ntp_server2": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mtu": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "multi_vmssupport": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "multicast_receive_interface": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_receive_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_receive_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_receive_range": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_send_interface": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_send_interface_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_send_interface_netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "multicast_source_portgroup": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "customized_script_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ovf_url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "avrs_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "avrs_profile": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vcenter": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceVCenterDataCenterRead(d *schema.ResourceData, m interface{}) error {
    filteredVCenterDataCenters := vspk.VCenterDataCentersList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    parent := &vspk.VCenter{ID: d.Get("parent_vcenter").(string)}
    filteredVCenterDataCenters, err = parent.VCenterDataCenters(fetchFilter)
    if err != nil {
        return err
    }

    VCenterDataCenter := &vspk.VCenterDataCenter{}

    if len(filteredVCenterDataCenters) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVCenterDataCenters) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VCenterDataCenter = filteredVCenterDataCenters[0]

    d.Set("vrs_configuration_time_limit", VCenterDataCenter.VRSConfigurationTimeLimit)
    d.Set("v_require_nuage_metadata", VCenterDataCenter.VRequireNuageMetadata)
    d.Set("name", VCenterDataCenter.Name)
    d.Set("managed_object_id", VCenterDataCenter.ManagedObjectID)
    d.Set("last_updated_by", VCenterDataCenter.LastUpdatedBy)
    d.Set("data_dns1", VCenterDataCenter.DataDNS1)
    d.Set("data_dns2", VCenterDataCenter.DataDNS2)
    d.Set("data_gateway", VCenterDataCenter.DataGateway)
    d.Set("data_network_portgroup", VCenterDataCenter.DataNetworkPortgroup)
    d.Set("datapath_sync_timeout", VCenterDataCenter.DatapathSyncTimeout)
    d.Set("secondary_data_uplink_dhcp_enabled", VCenterDataCenter.SecondaryDataUplinkDHCPEnabled)
    d.Set("secondary_data_uplink_enabled", VCenterDataCenter.SecondaryDataUplinkEnabled)
    d.Set("secondary_data_uplink_interface", VCenterDataCenter.SecondaryDataUplinkInterface)
    d.Set("secondary_data_uplink_mtu", VCenterDataCenter.SecondaryDataUplinkMTU)
    d.Set("secondary_data_uplink_primary_controller", VCenterDataCenter.SecondaryDataUplinkPrimaryController)
    d.Set("secondary_data_uplink_secondary_controller", VCenterDataCenter.SecondaryDataUplinkSecondaryController)
    d.Set("secondary_data_uplink_underlay_id", VCenterDataCenter.SecondaryDataUplinkUnderlayID)
    d.Set("secondary_nuage_controller", VCenterDataCenter.SecondaryNuageController)
    d.Set("deleted_from_vcenter", VCenterDataCenter.DeletedFromVCenter)
    d.Set("memory_size_in_gb", VCenterDataCenter.MemorySizeInGB)
    d.Set("remote_syslog_server_ip", VCenterDataCenter.RemoteSyslogServerIP)
    d.Set("remote_syslog_server_port", VCenterDataCenter.RemoteSyslogServerPort)
    d.Set("remote_syslog_server_type", VCenterDataCenter.RemoteSyslogServerType)
    d.Set("generic_split_activation", VCenterDataCenter.GenericSplitActivation)
    d.Set("separate_data_network", VCenterDataCenter.SeparateDataNetwork)
    d.Set("personality", VCenterDataCenter.Personality)
    d.Set("description", VCenterDataCenter.Description)
    d.Set("destination_mirror_port", VCenterDataCenter.DestinationMirrorPort)
    d.Set("metadata_server_ip", VCenterDataCenter.MetadataServerIP)
    d.Set("metadata_server_listen_port", VCenterDataCenter.MetadataServerListenPort)
    d.Set("metadata_server_port", VCenterDataCenter.MetadataServerPort)
    d.Set("metadata_service_enabled", VCenterDataCenter.MetadataServiceEnabled)
    d.Set("network_uplink_interface", VCenterDataCenter.NetworkUplinkInterface)
    d.Set("network_uplink_interface_gateway", VCenterDataCenter.NetworkUplinkInterfaceGateway)
    d.Set("network_uplink_interface_ip", VCenterDataCenter.NetworkUplinkInterfaceIp)
    d.Set("network_uplink_interface_netmask", VCenterDataCenter.NetworkUplinkInterfaceNetmask)
    d.Set("revertive_controller_enabled", VCenterDataCenter.RevertiveControllerEnabled)
    d.Set("revertive_timer", VCenterDataCenter.RevertiveTimer)
    d.Set("nfs_log_server", VCenterDataCenter.NfsLogServer)
    d.Set("nfs_mount_path", VCenterDataCenter.NfsMountPath)
    d.Set("mgmt_dns1", VCenterDataCenter.MgmtDNS1)
    d.Set("mgmt_dns2", VCenterDataCenter.MgmtDNS2)
    d.Set("mgmt_gateway", VCenterDataCenter.MgmtGateway)
    d.Set("mgmt_network_portgroup", VCenterDataCenter.MgmtNetworkPortgroup)
    d.Set("dhcp_relay_server", VCenterDataCenter.DhcpRelayServer)
    d.Set("mirror_network_portgroup", VCenterDataCenter.MirrorNetworkPortgroup)
    d.Set("disable_gro_on_datapath", VCenterDataCenter.DisableGROOnDatapath)
    d.Set("disable_lro_on_datapath", VCenterDataCenter.DisableLROOnDatapath)
    d.Set("site_id", VCenterDataCenter.SiteId)
    d.Set("allow_data_dhcp", VCenterDataCenter.AllowDataDHCP)
    d.Set("allow_mgmt_dhcp", VCenterDataCenter.AllowMgmtDHCP)
    d.Set("flow_eviction_threshold", VCenterDataCenter.FlowEvictionThreshold)
    d.Set("vm_network_portgroup", VCenterDataCenter.VmNetworkPortgroup)
    d.Set("enable_vrs_resource_reservation", VCenterDataCenter.EnableVRSResourceReservation)
    d.Set("entity_scope", VCenterDataCenter.EntityScope)
    d.Set("configured_metrics_push_interval", VCenterDataCenter.ConfiguredMetricsPushInterval)
    d.Set("portgroup_metadata", VCenterDataCenter.PortgroupMetadata)
    d.Set("nova_client_version", VCenterDataCenter.NovaClientVersion)
    d.Set("nova_identity_url_version", VCenterDataCenter.NovaIdentityURLVersion)
    d.Set("nova_metadata_service_auth_url", VCenterDataCenter.NovaMetadataServiceAuthUrl)
    d.Set("nova_metadata_service_endpoint", VCenterDataCenter.NovaMetadataServiceEndpoint)
    d.Set("nova_metadata_service_password", VCenterDataCenter.NovaMetadataServicePassword)
    d.Set("nova_metadata_service_tenant", VCenterDataCenter.NovaMetadataServiceTenant)
    d.Set("nova_metadata_service_username", VCenterDataCenter.NovaMetadataServiceUsername)
    d.Set("nova_metadata_shared_secret", VCenterDataCenter.NovaMetadataSharedSecret)
    d.Set("nova_os_keystone_username", VCenterDataCenter.NovaOSKeystoneUsername)
    d.Set("nova_project_domain_name", VCenterDataCenter.NovaProjectDomainName)
    d.Set("nova_project_name", VCenterDataCenter.NovaProjectName)
    d.Set("nova_region_name", VCenterDataCenter.NovaRegionName)
    d.Set("nova_user_domain_name", VCenterDataCenter.NovaUserDomainName)
    d.Set("upgrade_package_password", VCenterDataCenter.UpgradePackagePassword)
    d.Set("upgrade_package_url", VCenterDataCenter.UpgradePackageURL)
    d.Set("upgrade_package_username", VCenterDataCenter.UpgradePackageUsername)
    d.Set("upgrade_script_time_limit", VCenterDataCenter.UpgradeScriptTimeLimit)
    d.Set("cpu_count", VCenterDataCenter.CpuCount)
    d.Set("primary_data_uplink_underlay_id", VCenterDataCenter.PrimaryDataUplinkUnderlayID)
    d.Set("primary_nuage_controller", VCenterDataCenter.PrimaryNuageController)
    d.Set("vrs_password", VCenterDataCenter.VrsPassword)
    d.Set("vrs_user_name", VCenterDataCenter.VrsUserName)
    d.Set("associated_vcenter_id", VCenterDataCenter.AssociatedVCenterID)
    d.Set("static_route", VCenterDataCenter.StaticRoute)
    d.Set("static_route_gateway", VCenterDataCenter.StaticRouteGateway)
    d.Set("static_route_netmask", VCenterDataCenter.StaticRouteNetmask)
    d.Set("ntp_server1", VCenterDataCenter.NtpServer1)
    d.Set("ntp_server2", VCenterDataCenter.NtpServer2)
    d.Set("mtu", VCenterDataCenter.Mtu)
    d.Set("multi_vmssupport", VCenterDataCenter.MultiVMSsupport)
    d.Set("multicast_receive_interface", VCenterDataCenter.MulticastReceiveInterface)
    d.Set("multicast_receive_interface_ip", VCenterDataCenter.MulticastReceiveInterfaceIP)
    d.Set("multicast_receive_interface_netmask", VCenterDataCenter.MulticastReceiveInterfaceNetmask)
    d.Set("multicast_receive_range", VCenterDataCenter.MulticastReceiveRange)
    d.Set("multicast_send_interface", VCenterDataCenter.MulticastSendInterface)
    d.Set("multicast_send_interface_ip", VCenterDataCenter.MulticastSendInterfaceIP)
    d.Set("multicast_send_interface_netmask", VCenterDataCenter.MulticastSendInterfaceNetmask)
    d.Set("multicast_source_portgroup", VCenterDataCenter.MulticastSourcePortgroup)
    d.Set("customized_script_url", VCenterDataCenter.CustomizedScriptURL)
    d.Set("ovf_url", VCenterDataCenter.OvfURL)
    d.Set("avrs_enabled", VCenterDataCenter.AvrsEnabled)
    d.Set("avrs_profile", VCenterDataCenter.AvrsProfile)
    d.Set("external_id", VCenterDataCenter.ExternalID)
    
    d.Set("id", VCenterDataCenter.Identifier())
    d.Set("parent_id", VCenterDataCenter.ParentID)
    d.Set("parent_type", VCenterDataCenter.ParentType)
    d.Set("owner", VCenterDataCenter.Owner)

    d.SetId(VCenterDataCenter.Identifier())
    
    return nil
}