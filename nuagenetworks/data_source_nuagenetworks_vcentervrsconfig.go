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
			"v_require_nuage_metadata": &schema.Schema{
				Type:     schema.TypeBool,
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
			"secondary_nuage_controller": &schema.Schema{
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
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
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
			"nova_region_name": &schema.Schema{
				Type:     schema.TypeString,
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
			"external_id": &schema.Schema{
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
	} else {
		VCenterVRSConfig = filteredVCenterVRSConfigs[0]
	}

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
