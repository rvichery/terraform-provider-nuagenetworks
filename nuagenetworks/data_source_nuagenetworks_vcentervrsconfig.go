package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVCenterVRSConfig() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVCenterVRSConfigRead,
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
			"v_require_nuage_metadata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
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
			"data_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datapath_sync_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondary_nuage_controller": {
				Type:     schema.TypeString,
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
			"personality": {
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
			"mgmt_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_server": {
				Type:     schema.TypeString,
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
			"entity_scope": {
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
			"nova_region_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_nuage_controller": {
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
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVCenterVRSConfigRead(d *schema.ResourceData, m interface{}) error {
	filteredVCenterVRSConfigs := vspk.VCenterVRSConfigsList{}
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
	parent := m.(*vspk.Me)
	filteredVCenterVRSConfigs, err = parent.VCenterVRSConfigs(fetchFilter)
	if err != nil {
		return err
	}

	VCenterVRSConfig := &vspk.VCenterVRSConfig{}

	if len(filteredVCenterVRSConfigs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVCenterVRSConfigs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VCenterVRSConfig = filteredVCenterVRSConfigs[0]

	d.Set("v_require_nuage_metadata", VCenterVRSConfig.VRequireNuageMetadata)
	d.Set("last_updated_by", VCenterVRSConfig.LastUpdatedBy)
	d.Set("data_dns1", VCenterVRSConfig.DataDNS1)
	d.Set("data_dns2", VCenterVRSConfig.DataDNS2)
	d.Set("data_gateway", VCenterVRSConfig.DataGateway)
	d.Set("data_network_portgroup", VCenterVRSConfig.DataNetworkPortgroup)
	d.Set("datapath_sync_timeout", VCenterVRSConfig.DatapathSyncTimeout)
	d.Set("secondary_nuage_controller", VCenterVRSConfig.SecondaryNuageController)
	d.Set("generic_split_activation", VCenterVRSConfig.GenericSplitActivation)
	d.Set("separate_data_network", VCenterVRSConfig.SeparateDataNetwork)
	d.Set("personality", VCenterVRSConfig.Personality)
	d.Set("metadata_server_ip", VCenterVRSConfig.MetadataServerIP)
	d.Set("metadata_server_listen_port", VCenterVRSConfig.MetadataServerListenPort)
	d.Set("metadata_server_port", VCenterVRSConfig.MetadataServerPort)
	d.Set("metadata_service_enabled", VCenterVRSConfig.MetadataServiceEnabled)
	d.Set("network_uplink_interface", VCenterVRSConfig.NetworkUplinkInterface)
	d.Set("network_uplink_interface_gateway", VCenterVRSConfig.NetworkUplinkInterfaceGateway)
	d.Set("network_uplink_interface_ip", VCenterVRSConfig.NetworkUplinkInterfaceIp)
	d.Set("network_uplink_interface_netmask", VCenterVRSConfig.NetworkUplinkInterfaceNetmask)
	d.Set("nfs_log_server", VCenterVRSConfig.NfsLogServer)
	d.Set("nfs_mount_path", VCenterVRSConfig.NfsMountPath)
	d.Set("mgmt_dns1", VCenterVRSConfig.MgmtDNS1)
	d.Set("mgmt_dns2", VCenterVRSConfig.MgmtDNS2)
	d.Set("mgmt_gateway", VCenterVRSConfig.MgmtGateway)
	d.Set("mgmt_network_portgroup", VCenterVRSConfig.MgmtNetworkPortgroup)
	d.Set("dhcp_relay_server", VCenterVRSConfig.DhcpRelayServer)
	d.Set("site_id", VCenterVRSConfig.SiteId)
	d.Set("allow_data_dhcp", VCenterVRSConfig.AllowDataDHCP)
	d.Set("allow_mgmt_dhcp", VCenterVRSConfig.AllowMgmtDHCP)
	d.Set("flow_eviction_threshold", VCenterVRSConfig.FlowEvictionThreshold)
	d.Set("vm_network_portgroup", VCenterVRSConfig.VmNetworkPortgroup)
	d.Set("entity_scope", VCenterVRSConfig.EntityScope)
	d.Set("portgroup_metadata", VCenterVRSConfig.PortgroupMetadata)
	d.Set("nova_client_version", VCenterVRSConfig.NovaClientVersion)
	d.Set("nova_metadata_service_auth_url", VCenterVRSConfig.NovaMetadataServiceAuthUrl)
	d.Set("nova_metadata_service_endpoint", VCenterVRSConfig.NovaMetadataServiceEndpoint)
	d.Set("nova_metadata_service_password", VCenterVRSConfig.NovaMetadataServicePassword)
	d.Set("nova_metadata_service_tenant", VCenterVRSConfig.NovaMetadataServiceTenant)
	d.Set("nova_metadata_service_username", VCenterVRSConfig.NovaMetadataServiceUsername)
	d.Set("nova_metadata_shared_secret", VCenterVRSConfig.NovaMetadataSharedSecret)
	d.Set("nova_region_name", VCenterVRSConfig.NovaRegionName)
	d.Set("primary_nuage_controller", VCenterVRSConfig.PrimaryNuageController)
	d.Set("vrs_password", VCenterVRSConfig.VrsPassword)
	d.Set("vrs_user_name", VCenterVRSConfig.VrsUserName)
	d.Set("static_route", VCenterVRSConfig.StaticRoute)
	d.Set("static_route_gateway", VCenterVRSConfig.StaticRouteGateway)
	d.Set("static_route_netmask", VCenterVRSConfig.StaticRouteNetmask)
	d.Set("ntp_server1", VCenterVRSConfig.NtpServer1)
	d.Set("ntp_server2", VCenterVRSConfig.NtpServer2)
	d.Set("mtu", VCenterVRSConfig.Mtu)
	d.Set("multi_vmssupport", VCenterVRSConfig.MultiVMSsupport)
	d.Set("multicast_receive_interface", VCenterVRSConfig.MulticastReceiveInterface)
	d.Set("multicast_receive_interface_ip", VCenterVRSConfig.MulticastReceiveInterfaceIP)
	d.Set("multicast_receive_interface_netmask", VCenterVRSConfig.MulticastReceiveInterfaceNetmask)
	d.Set("multicast_receive_range", VCenterVRSConfig.MulticastReceiveRange)
	d.Set("multicast_send_interface", VCenterVRSConfig.MulticastSendInterface)
	d.Set("multicast_send_interface_ip", VCenterVRSConfig.MulticastSendInterfaceIP)
	d.Set("multicast_send_interface_netmask", VCenterVRSConfig.MulticastSendInterfaceNetmask)
	d.Set("multicast_source_portgroup", VCenterVRSConfig.MulticastSourcePortgroup)
	d.Set("customized_script_url", VCenterVRSConfig.CustomizedScriptURL)
	d.Set("external_id", VCenterVRSConfig.ExternalID)

	d.Set("id", VCenterVRSConfig.Identifier())
	d.Set("parent_id", VCenterVRSConfig.ParentID)
	d.Set("parent_type", VCenterVRSConfig.ParentType)
	d.Set("owner", VCenterVRSConfig.Owner)

	d.SetId(VCenterVRSConfig.Identifier())

	return nil
}
