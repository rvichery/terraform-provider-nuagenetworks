package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVCenterHypervisor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVCenterHypervisorRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_agent_moid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_agent_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_configuration_time_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vrs_metrics_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_mgmt_hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"v_require_nuage_metadata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_vrs_deployed_date": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"data_dns1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_dns2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datapath_sync_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scope": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"secondary_data_uplink_dhcp_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"secondary_data_uplink_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"secondary_data_uplink_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_data_uplink_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_data_uplink_mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondary_data_uplink_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_data_uplink_primary_controller": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_data_uplink_secondary_controller": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_data_uplink_underlay_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondary_nuage_controller": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_size_in_gb": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_syslog_server_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_syslog_server_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"remote_syslog_server_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"removed_from_vcenter_inventory": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"generic_split_activation": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"separate_data_network": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"deployment_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"personality": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_mirror_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_server_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_server_listen_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"metadata_server_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"metadata_service_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"network_uplink_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_uplink_interface_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_uplink_interface_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_uplink_interface_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revertive_controller_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"revertive_timer": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"nfs_log_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nfs_mount_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_dns1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_dns2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_gro_on_datapath": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disable_lro_on_datapath": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"site_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_data_dhcp": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_mgmt_dhcp": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"flow_eviction_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vm_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_vrs_resource_reservation": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configured_metrics_push_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"toolbox_deployment_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"toolbox_group": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"toolbox_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"toolbox_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"toolbox_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portgroup_metadata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"nova_client_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"nova_identity_url_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_auth_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_tenant": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_shared_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_os_keystone_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_project_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_project_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_region_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_user_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_package_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_package_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_package_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_script_time_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upgrade_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_timedout": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_data_uplink_underlay_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"primary_nuage_controller": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_route": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_route_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_route_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"successfully_applied_upgrade_package_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"successfully_applied_upgrade_package_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"successfully_applied_upgrade_package_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"successfully_applied_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multi_vmssupport": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"multicast_receive_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_receive_interface_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_receive_interface_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_receive_range": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_send_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_send_interface_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_send_interface_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_source_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customized_script_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"available_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ovf_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avrs_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"avrs_profile": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vcenter_cluster": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vcenter_data_center"},
			},
			"parent_vcenter_data_center": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vcenter_cluster"},
			},
		},
	}
}

func dataSourceVCenterHypervisorRead(d *schema.ResourceData, m interface{}) error {
	filteredVCenterHypervisors := vspk.VCenterHypervisorsList{}
	err := &bambou.Error{}
	fetchFilter := &bambou.FetchingInfo{}

	filters, filtersOk := d.GetOk("filter")
	if filtersOk {
		fetchFilter = bambou.NewFetchingInfo()
		for _, v := range filters.(*schema.Set).List() {
			m := v.(map[string]interface{})
			if fetchFilter.Filter != "" {
				fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string), m["operator"].(string), m["value"].(string))
			} else {
				fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
			}

		}
	}
	if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
		parent := &vspk.VCenterCluster{ID: attr.(string)}
		filteredVCenterHypervisors, err = parent.VCenterHypervisors(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
		parent := &vspk.VCenterDataCenter{ID: attr.(string)}
		filteredVCenterHypervisors, err = parent.VCenterHypervisors(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredVCenterHypervisors, err = parent.VCenterHypervisors(fetchFilter)
		if err != nil {
			return err
		}
	}

	VCenterHypervisor := &vspk.VCenterHypervisor{}

	if len(filteredVCenterHypervisors) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVCenterHypervisors) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VCenterHypervisor = filteredVCenterHypervisors[0]

	d.Set("vcenter_ip", VCenterHypervisor.VCenterIP)
	d.Set("vcenter_password", VCenterHypervisor.VCenterPassword)
	d.Set("vcenter_user", VCenterHypervisor.VCenterUser)
	d.Set("vrs_agent_moid", VCenterHypervisor.VRSAgentMOID)
	d.Set("vrs_agent_name", VCenterHypervisor.VRSAgentName)
	d.Set("vrs_configuration_time_limit", VCenterHypervisor.VRSConfigurationTimeLimit)
	d.Set("vrs_metrics_id", VCenterHypervisor.VRSMetricsID)
	d.Set("vrs_mgmt_hostname", VCenterHypervisor.VRSMgmtHostname)
	d.Set("vrs_state", VCenterHypervisor.VRSState)
	d.Set("v_require_nuage_metadata", VCenterHypervisor.VRequireNuageMetadata)
	d.Set("name", VCenterHypervisor.Name)
	d.Set("managed_object_id", VCenterHypervisor.ManagedObjectID)
	d.Set("last_updated_by", VCenterHypervisor.LastUpdatedBy)
	d.Set("last_vrs_deployed_date", VCenterHypervisor.LastVRSDeployedDate)
	d.Set("data_dns1", VCenterHypervisor.DataDNS1)
	d.Set("data_dns2", VCenterHypervisor.DataDNS2)
	d.Set("data_gateway", VCenterHypervisor.DataGateway)
	d.Set("data_ip_address", VCenterHypervisor.DataIPAddress)
	d.Set("data_netmask", VCenterHypervisor.DataNetmask)
	d.Set("data_network_portgroup", VCenterHypervisor.DataNetworkPortgroup)
	d.Set("datapath_sync_timeout", VCenterHypervisor.DatapathSyncTimeout)
	d.Set("scope", VCenterHypervisor.Scope)
	d.Set("secondary_data_uplink_dhcp_enabled", VCenterHypervisor.SecondaryDataUplinkDHCPEnabled)
	d.Set("secondary_data_uplink_enabled", VCenterHypervisor.SecondaryDataUplinkEnabled)
	d.Set("secondary_data_uplink_ip", VCenterHypervisor.SecondaryDataUplinkIP)
	d.Set("secondary_data_uplink_interface", VCenterHypervisor.SecondaryDataUplinkInterface)
	d.Set("secondary_data_uplink_mtu", VCenterHypervisor.SecondaryDataUplinkMTU)
	d.Set("secondary_data_uplink_netmask", VCenterHypervisor.SecondaryDataUplinkNetmask)
	d.Set("secondary_data_uplink_primary_controller", VCenterHypervisor.SecondaryDataUplinkPrimaryController)
	d.Set("secondary_data_uplink_secondary_controller", VCenterHypervisor.SecondaryDataUplinkSecondaryController)
	d.Set("secondary_data_uplink_underlay_id", VCenterHypervisor.SecondaryDataUplinkUnderlayID)
	d.Set("secondary_nuage_controller", VCenterHypervisor.SecondaryNuageController)
	d.Set("memory_size_in_gb", VCenterHypervisor.MemorySizeInGB)
	d.Set("remote_syslog_server_ip", VCenterHypervisor.RemoteSyslogServerIP)
	d.Set("remote_syslog_server_port", VCenterHypervisor.RemoteSyslogServerPort)
	d.Set("remote_syslog_server_type", VCenterHypervisor.RemoteSyslogServerType)
	d.Set("removed_from_vcenter_inventory", VCenterHypervisor.RemovedFromVCenterInventory)
	d.Set("generic_split_activation", VCenterHypervisor.GenericSplitActivation)
	d.Set("separate_data_network", VCenterHypervisor.SeparateDataNetwork)
	d.Set("deployment_count", VCenterHypervisor.DeploymentCount)
	d.Set("personality", VCenterHypervisor.Personality)
	d.Set("description", VCenterHypervisor.Description)
	d.Set("destination_mirror_port", VCenterHypervisor.DestinationMirrorPort)
	d.Set("metadata_server_ip", VCenterHypervisor.MetadataServerIP)
	d.Set("metadata_server_listen_port", VCenterHypervisor.MetadataServerListenPort)
	d.Set("metadata_server_port", VCenterHypervisor.MetadataServerPort)
	d.Set("metadata_service_enabled", VCenterHypervisor.MetadataServiceEnabled)
	d.Set("network_uplink_interface", VCenterHypervisor.NetworkUplinkInterface)
	d.Set("network_uplink_interface_gateway", VCenterHypervisor.NetworkUplinkInterfaceGateway)
	d.Set("network_uplink_interface_ip", VCenterHypervisor.NetworkUplinkInterfaceIp)
	d.Set("network_uplink_interface_netmask", VCenterHypervisor.NetworkUplinkInterfaceNetmask)
	d.Set("revertive_controller_enabled", VCenterHypervisor.RevertiveControllerEnabled)
	d.Set("revertive_timer", VCenterHypervisor.RevertiveTimer)
	d.Set("nfs_log_server", VCenterHypervisor.NfsLogServer)
	d.Set("nfs_mount_path", VCenterHypervisor.NfsMountPath)
	d.Set("mgmt_dns1", VCenterHypervisor.MgmtDNS1)
	d.Set("mgmt_dns2", VCenterHypervisor.MgmtDNS2)
	d.Set("mgmt_gateway", VCenterHypervisor.MgmtGateway)
	d.Set("mgmt_ip_address", VCenterHypervisor.MgmtIPAddress)
	d.Set("mgmt_netmask", VCenterHypervisor.MgmtNetmask)
	d.Set("mgmt_network_portgroup", VCenterHypervisor.MgmtNetworkPortgroup)
	d.Set("dhcp_relay_server", VCenterHypervisor.DhcpRelayServer)
	d.Set("mirror_network_portgroup", VCenterHypervisor.MirrorNetworkPortgroup)
	d.Set("disable_gro_on_datapath", VCenterHypervisor.DisableGROOnDatapath)
	d.Set("disable_lro_on_datapath", VCenterHypervisor.DisableLROOnDatapath)
	d.Set("site_id", VCenterHypervisor.SiteId)
	d.Set("allow_data_dhcp", VCenterHypervisor.AllowDataDHCP)
	d.Set("allow_mgmt_dhcp", VCenterHypervisor.AllowMgmtDHCP)
	d.Set("flow_eviction_threshold", VCenterHypervisor.FlowEvictionThreshold)
	d.Set("vm_network_portgroup", VCenterHypervisor.VmNetworkPortgroup)
	d.Set("enable_vrs_resource_reservation", VCenterHypervisor.EnableVRSResourceReservation)
	d.Set("entity_scope", VCenterHypervisor.EntityScope)
	d.Set("configured_metrics_push_interval", VCenterHypervisor.ConfiguredMetricsPushInterval)
	d.Set("toolbox_deployment_mode", VCenterHypervisor.ToolboxDeploymentMode)
	d.Set("toolbox_group", VCenterHypervisor.ToolboxGroup)
	d.Set("toolbox_ip", VCenterHypervisor.ToolboxIP)
	d.Set("toolbox_password", VCenterHypervisor.ToolboxPassword)
	d.Set("toolbox_user_name", VCenterHypervisor.ToolboxUserName)
	d.Set("portgroup_metadata", VCenterHypervisor.PortgroupMetadata)
	d.Set("nova_client_version", VCenterHypervisor.NovaClientVersion)
	d.Set("nova_identity_url_version", VCenterHypervisor.NovaIdentityURLVersion)
	d.Set("nova_metadata_service_auth_url", VCenterHypervisor.NovaMetadataServiceAuthUrl)
	d.Set("nova_metadata_service_endpoint", VCenterHypervisor.NovaMetadataServiceEndpoint)
	d.Set("nova_metadata_service_password", VCenterHypervisor.NovaMetadataServicePassword)
	d.Set("nova_metadata_service_tenant", VCenterHypervisor.NovaMetadataServiceTenant)
	d.Set("nova_metadata_service_username", VCenterHypervisor.NovaMetadataServiceUsername)
	d.Set("nova_metadata_shared_secret", VCenterHypervisor.NovaMetadataSharedSecret)
	d.Set("nova_os_keystone_username", VCenterHypervisor.NovaOSKeystoneUsername)
	d.Set("nova_project_domain_name", VCenterHypervisor.NovaProjectDomainName)
	d.Set("nova_project_name", VCenterHypervisor.NovaProjectName)
	d.Set("nova_region_name", VCenterHypervisor.NovaRegionName)
	d.Set("nova_user_domain_name", VCenterHypervisor.NovaUserDomainName)
	d.Set("upgrade_package_password", VCenterHypervisor.UpgradePackagePassword)
	d.Set("upgrade_package_url", VCenterHypervisor.UpgradePackageURL)
	d.Set("upgrade_package_username", VCenterHypervisor.UpgradePackageUsername)
	d.Set("upgrade_script_time_limit", VCenterHypervisor.UpgradeScriptTimeLimit)
	d.Set("upgrade_status", VCenterHypervisor.UpgradeStatus)
	d.Set("upgrade_timedout", VCenterHypervisor.UpgradeTimedout)
	d.Set("cpu_count", VCenterHypervisor.CpuCount)
	d.Set("primary_data_uplink_underlay_id", VCenterHypervisor.PrimaryDataUplinkUnderlayID)
	d.Set("primary_nuage_controller", VCenterHypervisor.PrimaryNuageController)
	d.Set("vrs_id", VCenterHypervisor.VrsId)
	d.Set("vrs_password", VCenterHypervisor.VrsPassword)
	d.Set("vrs_user_name", VCenterHypervisor.VrsUserName)
	d.Set("static_route", VCenterHypervisor.StaticRoute)
	d.Set("static_route_gateway", VCenterHypervisor.StaticRouteGateway)
	d.Set("static_route_netmask", VCenterHypervisor.StaticRouteNetmask)
	d.Set("ntp_server1", VCenterHypervisor.NtpServer1)
	d.Set("ntp_server2", VCenterHypervisor.NtpServer2)
	d.Set("mtu", VCenterHypervisor.Mtu)
	d.Set("successfully_applied_upgrade_package_password", VCenterHypervisor.SuccessfullyAppliedUpgradePackagePassword)
	d.Set("successfully_applied_upgrade_package_url", VCenterHypervisor.SuccessfullyAppliedUpgradePackageURL)
	d.Set("successfully_applied_upgrade_package_username", VCenterHypervisor.SuccessfullyAppliedUpgradePackageUsername)
	d.Set("successfully_applied_version", VCenterHypervisor.SuccessfullyAppliedVersion)
	d.Set("multi_vmssupport", VCenterHypervisor.MultiVMSsupport)
	d.Set("multicast_receive_interface", VCenterHypervisor.MulticastReceiveInterface)
	d.Set("multicast_receive_interface_ip", VCenterHypervisor.MulticastReceiveInterfaceIP)
	d.Set("multicast_receive_interface_netmask", VCenterHypervisor.MulticastReceiveInterfaceNetmask)
	d.Set("multicast_receive_range", VCenterHypervisor.MulticastReceiveRange)
	d.Set("multicast_send_interface", VCenterHypervisor.MulticastSendInterface)
	d.Set("multicast_send_interface_ip", VCenterHypervisor.MulticastSendInterfaceIP)
	d.Set("multicast_send_interface_netmask", VCenterHypervisor.MulticastSendInterfaceNetmask)
	d.Set("multicast_source_portgroup", VCenterHypervisor.MulticastSourcePortgroup)
	d.Set("customized_script_url", VCenterHypervisor.CustomizedScriptURL)
	d.Set("available_networks", VCenterHypervisor.AvailableNetworks)
	d.Set("ovf_url", VCenterHypervisor.OvfURL)
	d.Set("avrs_enabled", VCenterHypervisor.AvrsEnabled)
	d.Set("avrs_profile", VCenterHypervisor.AvrsProfile)
	d.Set("external_id", VCenterHypervisor.ExternalID)
	d.Set("hypervisor_ip", VCenterHypervisor.HypervisorIP)
	d.Set("hypervisor_password", VCenterHypervisor.HypervisorPassword)
	d.Set("hypervisor_user", VCenterHypervisor.HypervisorUser)

	d.Set("id", VCenterHypervisor.Identifier())
	d.Set("parent_id", VCenterHypervisor.ParentID)
	d.Set("parent_type", VCenterHypervisor.ParentType)
	d.Set("owner", VCenterHypervisor.Owner)

	d.SetId(VCenterHypervisor.Identifier())

	return nil
}
