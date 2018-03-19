package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVCenterHypervisor() *schema.Resource {
	return &schema.Resource{
		Create: resourceVCenterHypervisorCreate,
		Read:   resourceVCenterHypervisorRead,
		Update: resourceVCenterHypervisorUpdate,
		Delete: resourceVCenterHypervisorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcenter_user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrs_agent_moid": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrs_agent_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrs_configuration_time_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vrs_metrics_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrs_mgmt_hostname": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrs_state": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NOT_DEPLOYED",
			},
			"v_require_nuage_metadata": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"managed_object_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_vrs_deployed_date": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"data_dns1": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_dns2": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_network_portgroup": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datapath_sync_timeout": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"scope": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"secondary_data_uplink_dhcp_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"secondary_data_uplink_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"secondary_data_uplink_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_data_uplink_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_data_uplink_mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"secondary_data_uplink_netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_data_uplink_primary_controller": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_data_uplink_secondary_controller": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"secondary_data_uplink_underlay_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"secondary_nuage_controller": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory_size_in_gb": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DEFAULT_4",
			},
			"remote_syslog_server_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_syslog_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  514,
			},
			"remote_syslog_server_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "NONE",
			},
			"removed_from_vcenter_inventory": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"generic_split_activation": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"separate_data_network": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"deployment_count": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"personality": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"destination_mirror_port": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "no_mirror",
			},
			"metadata_server_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata_server_listen_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"metadata_server_port": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"metadata_service_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"network_uplink_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uplink_interface_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uplink_interface_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_uplink_interface_netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"revertive_controller_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"revertive_timer": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"nfs_log_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nfs_mount_path": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_dns1": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_dns2": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_network_portgroup": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dhcp_relay_server": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mirror_network_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"disable_gro_on_datapath": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"disable_lro_on_datapath": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"site_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"allow_data_dhcp": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"allow_mgmt_dhcp": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"flow_eviction_threshold": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vm_network_portgroup": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enable_vrs_resource_reservation": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configured_metrics_push_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"toolbox_deployment_mode": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"toolbox_group": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolbox_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolbox_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"toolbox_user_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"portgroup_metadata": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"nova_client_version": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nova_identity_url_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_metadata_service_auth_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_metadata_service_endpoint": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_metadata_service_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_metadata_service_tenant": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_metadata_service_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_metadata_shared_secret": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_os_keystone_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_project_domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_project_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_region_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"nova_user_domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upgrade_package_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upgrade_package_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upgrade_package_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upgrade_script_time_limit": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"upgrade_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"upgrade_timedout": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"cpu_count": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DEFAULT_2",
			},
			"primary_data_uplink_underlay_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"primary_nuage_controller": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrs_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrs_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrs_user_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_route": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_route_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_route_netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_server1": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_server2": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"successfully_applied_upgrade_package_password": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"successfully_applied_upgrade_package_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"successfully_applied_upgrade_package_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"successfully_applied_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multi_vmssupport": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"multicast_receive_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast_receive_interface_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast_receive_interface_netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast_receive_range": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast_send_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast_send_interface_ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast_send_interface_netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast_source_portgroup": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"customized_script_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"available_networks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ovf_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"avrs_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"avrs_profile": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AVRS_25G",
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"hypervisor_ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hypervisor_password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"hypervisor_user": {
				Type:     schema.TypeString,
				Required: true,
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

func resourceVCenterHypervisorCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VCenterHypervisor object
	o := &vspk.VCenterHypervisor{
		Name:                       d.Get("name").(string),
		DataNetworkPortgroup:       d.Get("data_network_portgroup").(string),
		SecondaryDataUplinkEnabled: d.Get("secondary_data_uplink_enabled").(bool),
		Description:                d.Get("description").(string),
		RevertiveControllerEnabled: d.Get("revertive_controller_enabled").(bool),
		RevertiveTimer:             d.Get("revertive_timer").(int),
		MgmtNetworkPortgroup:       d.Get("mgmt_network_portgroup").(string),
		VmNetworkPortgroup:         d.Get("vm_network_portgroup").(string),
		HypervisorIP:               d.Get("hypervisor_ip").(string),
		HypervisorPassword:         d.Get("hypervisor_password").(string),
		HypervisorUser:             d.Get("hypervisor_user").(string),
	}
	if attr, ok := d.GetOk("vcenter_ip"); ok {
		o.VCenterIP = attr.(string)
	}
	if attr, ok := d.GetOk("vcenter_password"); ok {
		o.VCenterPassword = attr.(string)
	}
	if attr, ok := d.GetOk("vcenter_user"); ok {
		o.VCenterUser = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_agent_moid"); ok {
		o.VRSAgentMOID = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_agent_name"); ok {
		o.VRSAgentName = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_configuration_time_limit"); ok {
		o.VRSConfigurationTimeLimit = attr.(int)
	}
	if attr, ok := d.GetOk("vrs_metrics_id"); ok {
		o.VRSMetricsID = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_mgmt_hostname"); ok {
		o.VRSMgmtHostname = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_state"); ok {
		o.VRSState = attr.(string)
	}
	if attr, ok := d.GetOk("v_require_nuage_metadata"); ok {
		o.VRequireNuageMetadata = attr.(bool)
	}
	if attr, ok := d.GetOk("managed_object_id"); ok {
		o.ManagedObjectID = attr.(string)
	}
	if attr, ok := d.GetOk("last_vrs_deployed_date"); ok {
		o.LastVRSDeployedDate = attr.(float64)
	}
	if attr, ok := d.GetOk("data_dns1"); ok {
		o.DataDNS1 = attr.(string)
	}
	if attr, ok := d.GetOk("data_dns2"); ok {
		o.DataDNS2 = attr.(string)
	}
	if attr, ok := d.GetOk("data_gateway"); ok {
		o.DataGateway = attr.(string)
	}
	if attr, ok := d.GetOk("data_ip_address"); ok {
		o.DataIPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("data_netmask"); ok {
		o.DataNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
		o.DatapathSyncTimeout = attr.(int)
	}
	if attr, ok := d.GetOk("scope"); ok {
		o.Scope = attr.(bool)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_dhcp_enabled"); ok {
		o.SecondaryDataUplinkDHCPEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_ip"); ok {
		o.SecondaryDataUplinkIP = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_interface"); ok {
		o.SecondaryDataUplinkInterface = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_mtu"); ok {
		o.SecondaryDataUplinkMTU = attr.(int)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_netmask"); ok {
		o.SecondaryDataUplinkNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_primary_controller"); ok {
		o.SecondaryDataUplinkPrimaryController = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_secondary_controller"); ok {
		o.SecondaryDataUplinkSecondaryController = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_underlay_id"); ok {
		o.SecondaryDataUplinkUnderlayID = attr.(int)
	}
	if attr, ok := d.GetOk("secondary_nuage_controller"); ok {
		o.SecondaryNuageController = attr.(string)
	}
	if attr, ok := d.GetOk("memory_size_in_gb"); ok {
		o.MemorySizeInGB = attr.(string)
	}
	if attr, ok := d.GetOk("remote_syslog_server_ip"); ok {
		o.RemoteSyslogServerIP = attr.(string)
	}
	if attr, ok := d.GetOk("remote_syslog_server_port"); ok {
		o.RemoteSyslogServerPort = attr.(int)
	}
	if attr, ok := d.GetOk("remote_syslog_server_type"); ok {
		o.RemoteSyslogServerType = attr.(string)
	}
	if attr, ok := d.GetOk("removed_from_vcenter_inventory"); ok {
		o.RemovedFromVCenterInventory = attr.(bool)
	}
	if attr, ok := d.GetOk("generic_split_activation"); ok {
		o.GenericSplitActivation = attr.(bool)
	}
	if attr, ok := d.GetOk("separate_data_network"); ok {
		o.SeparateDataNetwork = attr.(bool)
	}
	if attr, ok := d.GetOk("deployment_count"); ok {
		o.DeploymentCount = attr.(int)
	}
	if attr, ok := d.GetOk("personality"); ok {
		o.Personality = attr.(string)
	}
	if attr, ok := d.GetOk("destination_mirror_port"); ok {
		o.DestinationMirrorPort = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_server_ip"); ok {
		o.MetadataServerIP = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_server_listen_port"); ok {
		o.MetadataServerListenPort = attr.(int)
	}
	if attr, ok := d.GetOk("metadata_server_port"); ok {
		o.MetadataServerPort = attr.(int)
	}
	if attr, ok := d.GetOk("metadata_service_enabled"); ok {
		o.MetadataServiceEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("network_uplink_interface"); ok {
		o.NetworkUplinkInterface = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_gateway"); ok {
		o.NetworkUplinkInterfaceGateway = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_ip"); ok {
		o.NetworkUplinkInterfaceIp = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_netmask"); ok {
		o.NetworkUplinkInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("nfs_log_server"); ok {
		o.NfsLogServer = attr.(string)
	}
	if attr, ok := d.GetOk("nfs_mount_path"); ok {
		o.NfsMountPath = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_dns1"); ok {
		o.MgmtDNS1 = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_dns2"); ok {
		o.MgmtDNS2 = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_gateway"); ok {
		o.MgmtGateway = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_ip_address"); ok {
		o.MgmtIPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_netmask"); ok {
		o.MgmtNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("dhcp_relay_server"); ok {
		o.DhcpRelayServer = attr.(string)
	}
	if attr, ok := d.GetOk("mirror_network_portgroup"); ok {
		o.MirrorNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("disable_gro_on_datapath"); ok {
		o.DisableGROOnDatapath = attr.(bool)
	}
	if attr, ok := d.GetOk("disable_lro_on_datapath"); ok {
		o.DisableLROOnDatapath = attr.(bool)
	}
	if attr, ok := d.GetOk("site_id"); ok {
		o.SiteId = attr.(string)
	}
	if attr, ok := d.GetOk("allow_data_dhcp"); ok {
		o.AllowDataDHCP = attr.(bool)
	}
	if attr, ok := d.GetOk("allow_mgmt_dhcp"); ok {
		o.AllowMgmtDHCP = attr.(bool)
	}
	if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
		o.FlowEvictionThreshold = attr.(int)
	}
	if attr, ok := d.GetOk("enable_vrs_resource_reservation"); ok {
		o.EnableVRSResourceReservation = attr.(bool)
	}
	if attr, ok := d.GetOk("configured_metrics_push_interval"); ok {
		o.ConfiguredMetricsPushInterval = attr.(int)
	}
	if attr, ok := d.GetOk("toolbox_deployment_mode"); ok {
		o.ToolboxDeploymentMode = attr.(bool)
	}
	if attr, ok := d.GetOk("toolbox_group"); ok {
		o.ToolboxGroup = attr.(string)
	}
	if attr, ok := d.GetOk("toolbox_ip"); ok {
		o.ToolboxIP = attr.(string)
	}
	if attr, ok := d.GetOk("toolbox_password"); ok {
		o.ToolboxPassword = attr.(string)
	}
	if attr, ok := d.GetOk("toolbox_user_name"); ok {
		o.ToolboxUserName = attr.(string)
	}
	if attr, ok := d.GetOk("portgroup_metadata"); ok {
		o.PortgroupMetadata = attr.(bool)
	}
	if attr, ok := d.GetOk("nova_client_version"); ok {
		o.NovaClientVersion = attr.(int)
	}
	if attr, ok := d.GetOk("nova_identity_url_version"); ok {
		o.NovaIdentityURLVersion = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_auth_url"); ok {
		o.NovaMetadataServiceAuthUrl = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_endpoint"); ok {
		o.NovaMetadataServiceEndpoint = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_password"); ok {
		o.NovaMetadataServicePassword = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_tenant"); ok {
		o.NovaMetadataServiceTenant = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_username"); ok {
		o.NovaMetadataServiceUsername = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_shared_secret"); ok {
		o.NovaMetadataSharedSecret = attr.(string)
	}
	if attr, ok := d.GetOk("nova_os_keystone_username"); ok {
		o.NovaOSKeystoneUsername = attr.(string)
	}
	if attr, ok := d.GetOk("nova_project_domain_name"); ok {
		o.NovaProjectDomainName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_project_name"); ok {
		o.NovaProjectName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_region_name"); ok {
		o.NovaRegionName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_user_domain_name"); ok {
		o.NovaUserDomainName = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_password"); ok {
		o.UpgradePackagePassword = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_url"); ok {
		o.UpgradePackageURL = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_username"); ok {
		o.UpgradePackageUsername = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_script_time_limit"); ok {
		o.UpgradeScriptTimeLimit = attr.(int)
	}
	if attr, ok := d.GetOk("upgrade_status"); ok {
		o.UpgradeStatus = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_timedout"); ok {
		o.UpgradeTimedout = attr.(bool)
	}
	if attr, ok := d.GetOk("cpu_count"); ok {
		o.CpuCount = attr.(string)
	}
	if attr, ok := d.GetOk("primary_data_uplink_underlay_id"); ok {
		o.PrimaryDataUplinkUnderlayID = attr.(int)
	}
	if attr, ok := d.GetOk("primary_nuage_controller"); ok {
		o.PrimaryNuageController = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_id"); ok {
		o.VrsId = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_password"); ok {
		o.VrsPassword = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_user_name"); ok {
		o.VrsUserName = attr.(string)
	}
	if attr, ok := d.GetOk("static_route"); ok {
		o.StaticRoute = attr.(string)
	}
	if attr, ok := d.GetOk("static_route_gateway"); ok {
		o.StaticRouteGateway = attr.(string)
	}
	if attr, ok := d.GetOk("static_route_netmask"); ok {
		o.StaticRouteNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server1"); ok {
		o.NtpServer1 = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server2"); ok {
		o.NtpServer2 = attr.(string)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		o.Mtu = attr.(int)
	}
	if attr, ok := d.GetOk("successfully_applied_upgrade_package_password"); ok {
		o.SuccessfullyAppliedUpgradePackagePassword = attr.(string)
	}
	if attr, ok := d.GetOk("successfully_applied_upgrade_package_url"); ok {
		o.SuccessfullyAppliedUpgradePackageURL = attr.(string)
	}
	if attr, ok := d.GetOk("successfully_applied_upgrade_package_username"); ok {
		o.SuccessfullyAppliedUpgradePackageUsername = attr.(string)
	}
	if attr, ok := d.GetOk("successfully_applied_version"); ok {
		o.SuccessfullyAppliedVersion = attr.(string)
	}
	if attr, ok := d.GetOk("multi_vmssupport"); ok {
		o.MultiVMSsupport = attr.(bool)
	}
	if attr, ok := d.GetOk("multicast_receive_interface"); ok {
		o.MulticastReceiveInterface = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_interface_ip"); ok {
		o.MulticastReceiveInterfaceIP = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_interface_netmask"); ok {
		o.MulticastReceiveInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_range"); ok {
		o.MulticastReceiveRange = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface"); ok {
		o.MulticastSendInterface = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface_ip"); ok {
		o.MulticastSendInterfaceIP = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface_netmask"); ok {
		o.MulticastSendInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_source_portgroup"); ok {
		o.MulticastSourcePortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("customized_script_url"); ok {
		o.CustomizedScriptURL = attr.(string)
	}
	if attr, ok := d.GetOk("available_networks"); ok {
		o.AvailableNetworks = attr.([]interface{})
	}
	if attr, ok := d.GetOk("ovf_url"); ok {
		o.OvfURL = attr.(string)
	}
	if attr, ok := d.GetOk("avrs_enabled"); ok {
		o.AvrsEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("avrs_profile"); ok {
		o.AvrsProfile = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_vcenter_cluster"); ok {
		parent := &vspk.VCenterCluster{ID: attr.(string)}
		err := parent.CreateVCenterHypervisor(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vcenter_data_center"); ok {
		parent := &vspk.VCenterDataCenter{ID: attr.(string)}
		err := parent.CreateVCenterHypervisor(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceVCenterHypervisorRead(d, m)
}

func resourceVCenterHypervisorRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VCenterHypervisor{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("vcenter_ip", o.VCenterIP)
	d.Set("vcenter_password", o.VCenterPassword)
	d.Set("vcenter_user", o.VCenterUser)
	d.Set("vrs_agent_moid", o.VRSAgentMOID)
	d.Set("vrs_agent_name", o.VRSAgentName)
	d.Set("vrs_configuration_time_limit", o.VRSConfigurationTimeLimit)
	d.Set("vrs_metrics_id", o.VRSMetricsID)
	d.Set("vrs_mgmt_hostname", o.VRSMgmtHostname)
	d.Set("vrs_state", o.VRSState)
	d.Set("v_require_nuage_metadata", o.VRequireNuageMetadata)
	d.Set("name", o.Name)
	d.Set("managed_object_id", o.ManagedObjectID)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("last_vrs_deployed_date", o.LastVRSDeployedDate)
	d.Set("data_dns1", o.DataDNS1)
	d.Set("data_dns2", o.DataDNS2)
	d.Set("data_gateway", o.DataGateway)
	d.Set("data_ip_address", o.DataIPAddress)
	d.Set("data_netmask", o.DataNetmask)
	d.Set("data_network_portgroup", o.DataNetworkPortgroup)
	d.Set("datapath_sync_timeout", o.DatapathSyncTimeout)
	d.Set("scope", o.Scope)
	d.Set("secondary_data_uplink_dhcp_enabled", o.SecondaryDataUplinkDHCPEnabled)
	d.Set("secondary_data_uplink_enabled", o.SecondaryDataUplinkEnabled)
	d.Set("secondary_data_uplink_ip", o.SecondaryDataUplinkIP)
	d.Set("secondary_data_uplink_interface", o.SecondaryDataUplinkInterface)
	d.Set("secondary_data_uplink_mtu", o.SecondaryDataUplinkMTU)
	d.Set("secondary_data_uplink_netmask", o.SecondaryDataUplinkNetmask)
	d.Set("secondary_data_uplink_primary_controller", o.SecondaryDataUplinkPrimaryController)
	d.Set("secondary_data_uplink_secondary_controller", o.SecondaryDataUplinkSecondaryController)
	d.Set("secondary_data_uplink_underlay_id", o.SecondaryDataUplinkUnderlayID)
	d.Set("secondary_nuage_controller", o.SecondaryNuageController)
	d.Set("memory_size_in_gb", o.MemorySizeInGB)
	d.Set("remote_syslog_server_ip", o.RemoteSyslogServerIP)
	d.Set("remote_syslog_server_port", o.RemoteSyslogServerPort)
	d.Set("remote_syslog_server_type", o.RemoteSyslogServerType)
	d.Set("removed_from_vcenter_inventory", o.RemovedFromVCenterInventory)
	d.Set("generic_split_activation", o.GenericSplitActivation)
	d.Set("separate_data_network", o.SeparateDataNetwork)
	d.Set("deployment_count", o.DeploymentCount)
	d.Set("personality", o.Personality)
	d.Set("description", o.Description)
	d.Set("destination_mirror_port", o.DestinationMirrorPort)
	d.Set("metadata_server_ip", o.MetadataServerIP)
	d.Set("metadata_server_listen_port", o.MetadataServerListenPort)
	d.Set("metadata_server_port", o.MetadataServerPort)
	d.Set("metadata_service_enabled", o.MetadataServiceEnabled)
	d.Set("network_uplink_interface", o.NetworkUplinkInterface)
	d.Set("network_uplink_interface_gateway", o.NetworkUplinkInterfaceGateway)
	d.Set("network_uplink_interface_ip", o.NetworkUplinkInterfaceIp)
	d.Set("network_uplink_interface_netmask", o.NetworkUplinkInterfaceNetmask)
	d.Set("revertive_controller_enabled", o.RevertiveControllerEnabled)
	d.Set("revertive_timer", o.RevertiveTimer)
	d.Set("nfs_log_server", o.NfsLogServer)
	d.Set("nfs_mount_path", o.NfsMountPath)
	d.Set("mgmt_dns1", o.MgmtDNS1)
	d.Set("mgmt_dns2", o.MgmtDNS2)
	d.Set("mgmt_gateway", o.MgmtGateway)
	d.Set("mgmt_ip_address", o.MgmtIPAddress)
	d.Set("mgmt_netmask", o.MgmtNetmask)
	d.Set("mgmt_network_portgroup", o.MgmtNetworkPortgroup)
	d.Set("dhcp_relay_server", o.DhcpRelayServer)
	d.Set("mirror_network_portgroup", o.MirrorNetworkPortgroup)
	d.Set("disable_gro_on_datapath", o.DisableGROOnDatapath)
	d.Set("disable_lro_on_datapath", o.DisableLROOnDatapath)
	d.Set("site_id", o.SiteId)
	d.Set("allow_data_dhcp", o.AllowDataDHCP)
	d.Set("allow_mgmt_dhcp", o.AllowMgmtDHCP)
	d.Set("flow_eviction_threshold", o.FlowEvictionThreshold)
	d.Set("vm_network_portgroup", o.VmNetworkPortgroup)
	d.Set("enable_vrs_resource_reservation", o.EnableVRSResourceReservation)
	d.Set("entity_scope", o.EntityScope)
	d.Set("configured_metrics_push_interval", o.ConfiguredMetricsPushInterval)
	d.Set("toolbox_deployment_mode", o.ToolboxDeploymentMode)
	d.Set("toolbox_group", o.ToolboxGroup)
	d.Set("toolbox_ip", o.ToolboxIP)
	d.Set("toolbox_password", o.ToolboxPassword)
	d.Set("toolbox_user_name", o.ToolboxUserName)
	d.Set("portgroup_metadata", o.PortgroupMetadata)
	d.Set("nova_client_version", o.NovaClientVersion)
	d.Set("nova_identity_url_version", o.NovaIdentityURLVersion)
	d.Set("nova_metadata_service_auth_url", o.NovaMetadataServiceAuthUrl)
	d.Set("nova_metadata_service_endpoint", o.NovaMetadataServiceEndpoint)
	d.Set("nova_metadata_service_password", o.NovaMetadataServicePassword)
	d.Set("nova_metadata_service_tenant", o.NovaMetadataServiceTenant)
	d.Set("nova_metadata_service_username", o.NovaMetadataServiceUsername)
	d.Set("nova_metadata_shared_secret", o.NovaMetadataSharedSecret)
	d.Set("nova_os_keystone_username", o.NovaOSKeystoneUsername)
	d.Set("nova_project_domain_name", o.NovaProjectDomainName)
	d.Set("nova_project_name", o.NovaProjectName)
	d.Set("nova_region_name", o.NovaRegionName)
	d.Set("nova_user_domain_name", o.NovaUserDomainName)
	d.Set("upgrade_package_password", o.UpgradePackagePassword)
	d.Set("upgrade_package_url", o.UpgradePackageURL)
	d.Set("upgrade_package_username", o.UpgradePackageUsername)
	d.Set("upgrade_script_time_limit", o.UpgradeScriptTimeLimit)
	d.Set("upgrade_status", o.UpgradeStatus)
	d.Set("upgrade_timedout", o.UpgradeTimedout)
	d.Set("cpu_count", o.CpuCount)
	d.Set("primary_data_uplink_underlay_id", o.PrimaryDataUplinkUnderlayID)
	d.Set("primary_nuage_controller", o.PrimaryNuageController)
	d.Set("vrs_id", o.VrsId)
	d.Set("vrs_password", o.VrsPassword)
	d.Set("vrs_user_name", o.VrsUserName)
	d.Set("static_route", o.StaticRoute)
	d.Set("static_route_gateway", o.StaticRouteGateway)
	d.Set("static_route_netmask", o.StaticRouteNetmask)
	d.Set("ntp_server1", o.NtpServer1)
	d.Set("ntp_server2", o.NtpServer2)
	d.Set("mtu", o.Mtu)
	d.Set("successfully_applied_upgrade_package_password", o.SuccessfullyAppliedUpgradePackagePassword)
	d.Set("successfully_applied_upgrade_package_url", o.SuccessfullyAppliedUpgradePackageURL)
	d.Set("successfully_applied_upgrade_package_username", o.SuccessfullyAppliedUpgradePackageUsername)
	d.Set("successfully_applied_version", o.SuccessfullyAppliedVersion)
	d.Set("multi_vmssupport", o.MultiVMSsupport)
	d.Set("multicast_receive_interface", o.MulticastReceiveInterface)
	d.Set("multicast_receive_interface_ip", o.MulticastReceiveInterfaceIP)
	d.Set("multicast_receive_interface_netmask", o.MulticastReceiveInterfaceNetmask)
	d.Set("multicast_receive_range", o.MulticastReceiveRange)
	d.Set("multicast_send_interface", o.MulticastSendInterface)
	d.Set("multicast_send_interface_ip", o.MulticastSendInterfaceIP)
	d.Set("multicast_send_interface_netmask", o.MulticastSendInterfaceNetmask)
	d.Set("multicast_source_portgroup", o.MulticastSourcePortgroup)
	d.Set("customized_script_url", o.CustomizedScriptURL)
	d.Set("available_networks", o.AvailableNetworks)
	d.Set("ovf_url", o.OvfURL)
	d.Set("avrs_enabled", o.AvrsEnabled)
	d.Set("avrs_profile", o.AvrsProfile)
	d.Set("external_id", o.ExternalID)
	d.Set("hypervisor_ip", o.HypervisorIP)
	d.Set("hypervisor_password", o.HypervisorPassword)
	d.Set("hypervisor_user", o.HypervisorUser)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVCenterHypervisorUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VCenterHypervisor{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.DataNetworkPortgroup = d.Get("data_network_portgroup").(string)
	o.SecondaryDataUplinkEnabled = d.Get("secondary_data_uplink_enabled").(bool)
	o.Description = d.Get("description").(string)
	o.RevertiveControllerEnabled = d.Get("revertive_controller_enabled").(bool)
	o.RevertiveTimer = d.Get("revertive_timer").(int)
	o.MgmtNetworkPortgroup = d.Get("mgmt_network_portgroup").(string)
	o.VmNetworkPortgroup = d.Get("vm_network_portgroup").(string)
	o.HypervisorIP = d.Get("hypervisor_ip").(string)
	o.HypervisorPassword = d.Get("hypervisor_password").(string)
	o.HypervisorUser = d.Get("hypervisor_user").(string)

	if attr, ok := d.GetOk("vcenter_ip"); ok {
		o.VCenterIP = attr.(string)
	}
	if attr, ok := d.GetOk("vcenter_password"); ok {
		o.VCenterPassword = attr.(string)
	}
	if attr, ok := d.GetOk("vcenter_user"); ok {
		o.VCenterUser = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_agent_moid"); ok {
		o.VRSAgentMOID = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_agent_name"); ok {
		o.VRSAgentName = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_configuration_time_limit"); ok {
		o.VRSConfigurationTimeLimit = attr.(int)
	}
	if attr, ok := d.GetOk("vrs_metrics_id"); ok {
		o.VRSMetricsID = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_mgmt_hostname"); ok {
		o.VRSMgmtHostname = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_state"); ok {
		o.VRSState = attr.(string)
	}
	if attr, ok := d.GetOk("v_require_nuage_metadata"); ok {
		o.VRequireNuageMetadata = attr.(bool)
	}
	if attr, ok := d.GetOk("managed_object_id"); ok {
		o.ManagedObjectID = attr.(string)
	}
	if attr, ok := d.GetOk("last_vrs_deployed_date"); ok {
		o.LastVRSDeployedDate = attr.(float64)
	}
	if attr, ok := d.GetOk("data_dns1"); ok {
		o.DataDNS1 = attr.(string)
	}
	if attr, ok := d.GetOk("data_dns2"); ok {
		o.DataDNS2 = attr.(string)
	}
	if attr, ok := d.GetOk("data_gateway"); ok {
		o.DataGateway = attr.(string)
	}
	if attr, ok := d.GetOk("data_ip_address"); ok {
		o.DataIPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("data_netmask"); ok {
		o.DataNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("datapath_sync_timeout"); ok {
		o.DatapathSyncTimeout = attr.(int)
	}
	if attr, ok := d.GetOk("scope"); ok {
		o.Scope = attr.(bool)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_dhcp_enabled"); ok {
		o.SecondaryDataUplinkDHCPEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_ip"); ok {
		o.SecondaryDataUplinkIP = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_interface"); ok {
		o.SecondaryDataUplinkInterface = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_mtu"); ok {
		o.SecondaryDataUplinkMTU = attr.(int)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_netmask"); ok {
		o.SecondaryDataUplinkNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_primary_controller"); ok {
		o.SecondaryDataUplinkPrimaryController = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_secondary_controller"); ok {
		o.SecondaryDataUplinkSecondaryController = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_data_uplink_underlay_id"); ok {
		o.SecondaryDataUplinkUnderlayID = attr.(int)
	}
	if attr, ok := d.GetOk("secondary_nuage_controller"); ok {
		o.SecondaryNuageController = attr.(string)
	}
	if attr, ok := d.GetOk("memory_size_in_gb"); ok {
		o.MemorySizeInGB = attr.(string)
	}
	if attr, ok := d.GetOk("remote_syslog_server_ip"); ok {
		o.RemoteSyslogServerIP = attr.(string)
	}
	if attr, ok := d.GetOk("remote_syslog_server_port"); ok {
		o.RemoteSyslogServerPort = attr.(int)
	}
	if attr, ok := d.GetOk("remote_syslog_server_type"); ok {
		o.RemoteSyslogServerType = attr.(string)
	}
	if attr, ok := d.GetOk("removed_from_vcenter_inventory"); ok {
		o.RemovedFromVCenterInventory = attr.(bool)
	}
	if attr, ok := d.GetOk("generic_split_activation"); ok {
		o.GenericSplitActivation = attr.(bool)
	}
	if attr, ok := d.GetOk("separate_data_network"); ok {
		o.SeparateDataNetwork = attr.(bool)
	}
	if attr, ok := d.GetOk("deployment_count"); ok {
		o.DeploymentCount = attr.(int)
	}
	if attr, ok := d.GetOk("personality"); ok {
		o.Personality = attr.(string)
	}
	if attr, ok := d.GetOk("destination_mirror_port"); ok {
		o.DestinationMirrorPort = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_server_ip"); ok {
		o.MetadataServerIP = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_server_listen_port"); ok {
		o.MetadataServerListenPort = attr.(int)
	}
	if attr, ok := d.GetOk("metadata_server_port"); ok {
		o.MetadataServerPort = attr.(int)
	}
	if attr, ok := d.GetOk("metadata_service_enabled"); ok {
		o.MetadataServiceEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("network_uplink_interface"); ok {
		o.NetworkUplinkInterface = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_gateway"); ok {
		o.NetworkUplinkInterfaceGateway = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_ip"); ok {
		o.NetworkUplinkInterfaceIp = attr.(string)
	}
	if attr, ok := d.GetOk("network_uplink_interface_netmask"); ok {
		o.NetworkUplinkInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("nfs_log_server"); ok {
		o.NfsLogServer = attr.(string)
	}
	if attr, ok := d.GetOk("nfs_mount_path"); ok {
		o.NfsMountPath = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_dns1"); ok {
		o.MgmtDNS1 = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_dns2"); ok {
		o.MgmtDNS2 = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_gateway"); ok {
		o.MgmtGateway = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_ip_address"); ok {
		o.MgmtIPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("mgmt_netmask"); ok {
		o.MgmtNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("dhcp_relay_server"); ok {
		o.DhcpRelayServer = attr.(string)
	}
	if attr, ok := d.GetOk("mirror_network_portgroup"); ok {
		o.MirrorNetworkPortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("disable_gro_on_datapath"); ok {
		o.DisableGROOnDatapath = attr.(bool)
	}
	if attr, ok := d.GetOk("disable_lro_on_datapath"); ok {
		o.DisableLROOnDatapath = attr.(bool)
	}
	if attr, ok := d.GetOk("site_id"); ok {
		o.SiteId = attr.(string)
	}
	if attr, ok := d.GetOk("allow_data_dhcp"); ok {
		o.AllowDataDHCP = attr.(bool)
	}
	if attr, ok := d.GetOk("allow_mgmt_dhcp"); ok {
		o.AllowMgmtDHCP = attr.(bool)
	}
	if attr, ok := d.GetOk("flow_eviction_threshold"); ok {
		o.FlowEvictionThreshold = attr.(int)
	}
	if attr, ok := d.GetOk("enable_vrs_resource_reservation"); ok {
		o.EnableVRSResourceReservation = attr.(bool)
	}
	if attr, ok := d.GetOk("configured_metrics_push_interval"); ok {
		o.ConfiguredMetricsPushInterval = attr.(int)
	}
	if attr, ok := d.GetOk("toolbox_deployment_mode"); ok {
		o.ToolboxDeploymentMode = attr.(bool)
	}
	if attr, ok := d.GetOk("toolbox_group"); ok {
		o.ToolboxGroup = attr.(string)
	}
	if attr, ok := d.GetOk("toolbox_ip"); ok {
		o.ToolboxIP = attr.(string)
	}
	if attr, ok := d.GetOk("toolbox_password"); ok {
		o.ToolboxPassword = attr.(string)
	}
	if attr, ok := d.GetOk("toolbox_user_name"); ok {
		o.ToolboxUserName = attr.(string)
	}
	if attr, ok := d.GetOk("portgroup_metadata"); ok {
		o.PortgroupMetadata = attr.(bool)
	}
	if attr, ok := d.GetOk("nova_client_version"); ok {
		o.NovaClientVersion = attr.(int)
	}
	if attr, ok := d.GetOk("nova_identity_url_version"); ok {
		o.NovaIdentityURLVersion = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_auth_url"); ok {
		o.NovaMetadataServiceAuthUrl = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_endpoint"); ok {
		o.NovaMetadataServiceEndpoint = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_password"); ok {
		o.NovaMetadataServicePassword = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_tenant"); ok {
		o.NovaMetadataServiceTenant = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_service_username"); ok {
		o.NovaMetadataServiceUsername = attr.(string)
	}
	if attr, ok := d.GetOk("nova_metadata_shared_secret"); ok {
		o.NovaMetadataSharedSecret = attr.(string)
	}
	if attr, ok := d.GetOk("nova_os_keystone_username"); ok {
		o.NovaOSKeystoneUsername = attr.(string)
	}
	if attr, ok := d.GetOk("nova_project_domain_name"); ok {
		o.NovaProjectDomainName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_project_name"); ok {
		o.NovaProjectName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_region_name"); ok {
		o.NovaRegionName = attr.(string)
	}
	if attr, ok := d.GetOk("nova_user_domain_name"); ok {
		o.NovaUserDomainName = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_password"); ok {
		o.UpgradePackagePassword = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_url"); ok {
		o.UpgradePackageURL = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_package_username"); ok {
		o.UpgradePackageUsername = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_script_time_limit"); ok {
		o.UpgradeScriptTimeLimit = attr.(int)
	}
	if attr, ok := d.GetOk("upgrade_status"); ok {
		o.UpgradeStatus = attr.(string)
	}
	if attr, ok := d.GetOk("upgrade_timedout"); ok {
		o.UpgradeTimedout = attr.(bool)
	}
	if attr, ok := d.GetOk("cpu_count"); ok {
		o.CpuCount = attr.(string)
	}
	if attr, ok := d.GetOk("primary_data_uplink_underlay_id"); ok {
		o.PrimaryDataUplinkUnderlayID = attr.(int)
	}
	if attr, ok := d.GetOk("primary_nuage_controller"); ok {
		o.PrimaryNuageController = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_id"); ok {
		o.VrsId = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_password"); ok {
		o.VrsPassword = attr.(string)
	}
	if attr, ok := d.GetOk("vrs_user_name"); ok {
		o.VrsUserName = attr.(string)
	}
	if attr, ok := d.GetOk("static_route"); ok {
		o.StaticRoute = attr.(string)
	}
	if attr, ok := d.GetOk("static_route_gateway"); ok {
		o.StaticRouteGateway = attr.(string)
	}
	if attr, ok := d.GetOk("static_route_netmask"); ok {
		o.StaticRouteNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server1"); ok {
		o.NtpServer1 = attr.(string)
	}
	if attr, ok := d.GetOk("ntp_server2"); ok {
		o.NtpServer2 = attr.(string)
	}
	if attr, ok := d.GetOk("mtu"); ok {
		o.Mtu = attr.(int)
	}
	if attr, ok := d.GetOk("successfully_applied_upgrade_package_password"); ok {
		o.SuccessfullyAppliedUpgradePackagePassword = attr.(string)
	}
	if attr, ok := d.GetOk("successfully_applied_upgrade_package_url"); ok {
		o.SuccessfullyAppliedUpgradePackageURL = attr.(string)
	}
	if attr, ok := d.GetOk("successfully_applied_upgrade_package_username"); ok {
		o.SuccessfullyAppliedUpgradePackageUsername = attr.(string)
	}
	if attr, ok := d.GetOk("successfully_applied_version"); ok {
		o.SuccessfullyAppliedVersion = attr.(string)
	}
	if attr, ok := d.GetOk("multi_vmssupport"); ok {
		o.MultiVMSsupport = attr.(bool)
	}
	if attr, ok := d.GetOk("multicast_receive_interface"); ok {
		o.MulticastReceiveInterface = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_interface_ip"); ok {
		o.MulticastReceiveInterfaceIP = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_interface_netmask"); ok {
		o.MulticastReceiveInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_receive_range"); ok {
		o.MulticastReceiveRange = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface"); ok {
		o.MulticastSendInterface = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface_ip"); ok {
		o.MulticastSendInterfaceIP = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_send_interface_netmask"); ok {
		o.MulticastSendInterfaceNetmask = attr.(string)
	}
	if attr, ok := d.GetOk("multicast_source_portgroup"); ok {
		o.MulticastSourcePortgroup = attr.(string)
	}
	if attr, ok := d.GetOk("customized_script_url"); ok {
		o.CustomizedScriptURL = attr.(string)
	}
	if attr, ok := d.GetOk("available_networks"); ok {
		o.AvailableNetworks = attr.([]interface{})
	}
	if attr, ok := d.GetOk("ovf_url"); ok {
		o.OvfURL = attr.(string)
	}
	if attr, ok := d.GetOk("avrs_enabled"); ok {
		o.AvrsEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("avrs_profile"); ok {
		o.AvrsProfile = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceVCenterHypervisorDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VCenterHypervisor{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
