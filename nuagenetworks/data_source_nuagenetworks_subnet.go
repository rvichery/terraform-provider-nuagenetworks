package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceSubnet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSubnetRead,
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
			"pat_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dpi": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_mode": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway_mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"access_restriction_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"advertise": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default_action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_collection_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vn_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"underlay": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"underlay_enabled": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"route_distinguisher": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_target": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"split_subnet": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"use_global_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_multicast_channel_map_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_shared_network_resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"multi_home_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"multicast": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_ipv6_address": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parent_zone": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet_template", "parent_ike_gateway_connection", "parent_domain", "parent_ns_gateway", "parent_pat_mapper"},
			},
			"parent_subnet_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_ike_gateway_connection", "parent_domain", "parent_ns_gateway", "parent_pat_mapper"},
			},
			"parent_ike_gateway_connection": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_ns_gateway", "parent_pat_mapper"},
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_ike_gateway_connection", "parent_ns_gateway", "parent_pat_mapper"},
			},
			"parent_ns_gateway": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_ike_gateway_connection", "parent_domain", "parent_pat_mapper"},
			},
			"parent_pat_mapper": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_ike_gateway_connection", "parent_domain", "parent_ns_gateway"},
			},
		},
	}
}

func dataSourceSubnetRead(d *schema.ResourceData, m interface{}) error {
	filteredSubnets := vspk.SubnetsList{}
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
	if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		filteredSubnets, err = parent.Subnets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet_template"); ok {
		parent := &vspk.SubnetTemplate{ID: attr.(string)}
		filteredSubnets, err = parent.Subnets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ike_gateway_connection"); ok {
		parent := &vspk.IKEGatewayConnection{ID: attr.(string)}
		filteredSubnets, err = parent.Subnets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredSubnets, err = parent.Subnets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
		parent := &vspk.NSGateway{ID: attr.(string)}
		filteredSubnets, err = parent.Subnets(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_pat_mapper"); ok {
		parent := &vspk.PATMapper{ID: attr.(string)}
		filteredSubnets, err = parent.Subnets(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredSubnets, err = parent.Subnets(fetchFilter)
		if err != nil {
			return err
		}
	}

	Subnet := &vspk.Subnet{}

	if len(filteredSubnets) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredSubnets) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		Subnet = filteredSubnets[0]
	}

	d.Set("pat_enabled", Subnet.PATEnabled)
	d.Set("dhcp_relay_status", Subnet.DHCPRelayStatus)
	d.Set("dpi", Subnet.DPI)
	d.Set("ip_type", Subnet.IPType)
	d.Set("ipv6_address", Subnet.IPv6Address)
	d.Set("ipv6_gateway", Subnet.IPv6Gateway)
	d.Set("maintenance_mode", Subnet.MaintenanceMode)
	d.Set("name", Subnet.Name)
	d.Set("last_updated_by", Subnet.LastUpdatedBy)
	d.Set("gateway", Subnet.Gateway)
	d.Set("gateway_mac_address", Subnet.GatewayMACAddress)
	d.Set("access_restriction_enabled", Subnet.AccessRestrictionEnabled)
	d.Set("address", Subnet.Address)
	d.Set("advertise", Subnet.Advertise)
	d.Set("default_action", Subnet.DefaultAction)
	d.Set("template_id", Subnet.TemplateID)
	d.Set("service_id", Subnet.ServiceID)
	d.Set("description", Subnet.Description)
	d.Set("resource_type", Subnet.ResourceType)
	d.Set("netmask", Subnet.Netmask)
	d.Set("flow_collection_enabled", Subnet.FlowCollectionEnabled)
	d.Set("vn_id", Subnet.VnId)
	d.Set("encryption", Subnet.Encryption)
	d.Set("underlay", Subnet.Underlay)
	d.Set("underlay_enabled", Subnet.UnderlayEnabled)
	d.Set("entity_scope", Subnet.EntityScope)
	d.Set("entity_state", Subnet.EntityState)
	d.Set("policy_group_id", Subnet.PolicyGroupID)
	d.Set("route_distinguisher", Subnet.RouteDistinguisher)
	d.Set("route_target", Subnet.RouteTarget)
	d.Set("split_subnet", Subnet.SplitSubnet)
	d.Set("proxy_arp", Subnet.ProxyARP)
	d.Set("use_global_mac", Subnet.UseGlobalMAC)
	d.Set("associated_multicast_channel_map_id", Subnet.AssociatedMulticastChannelMapID)
	d.Set("associated_shared_network_resource_id", Subnet.AssociatedSharedNetworkResourceID)
	d.Set("public", Subnet.Public)
	d.Set("multi_home_enabled", Subnet.MultiHomeEnabled)
	d.Set("multicast", Subnet.Multicast)
	d.Set("external_id", Subnet.ExternalID)
	d.Set("dynamic_ipv6_address", Subnet.DynamicIpv6Address)

	d.Set("id", Subnet.Identifier())
	d.Set("parent_id", Subnet.ParentID)
	d.Set("parent_type", Subnet.ParentType)
	d.Set("owner", Subnet.Owner)

	d.SetId(Subnet.Identifier())

	return nil
}
