package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubnetCreate,
		Read:   resourceSubnetRead,
		Update: resourceSubnetUpdate,
		Delete: resourceSubnetDelete,
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
			"pat_enabled": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_relay_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dpi": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INHERITED",
			},
			"ip_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_mac_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"access_restriction_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"advertise": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"default_action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"service_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resource_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "STANDARD",
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"flow_collection_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INHERITED",
			},
			"vn_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"encryption": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"underlay": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"underlay_enabled": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"route_distinguisher": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"route_target": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"split_subnet": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_global_mac": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_multicast_channel_map_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_shared_network_resource_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"public": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"multi_home_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"multicast": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_ipv6_address": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"parent_zone": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSubnetCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Subnet object
	o := &vspk.Subnet{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("pat_enabled"); ok {
		o.PATEnabled = attr.(string)
	}
	if attr, ok := d.GetOk("dhcp_relay_status"); ok {
		o.DHCPRelayStatus = attr.(string)
	}
	if attr, ok := d.GetOk("dpi"); ok {
		o.DPI = attr.(string)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address"); ok {
		o.IPv6Address = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_gateway"); ok {
		o.IPv6Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("maintenance_mode"); ok {
		o.MaintenanceMode = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("gateway_mac_address"); ok {
		o.GatewayMACAddress = attr.(string)
	}
	if attr, ok := d.GetOk("access_restriction_enabled"); ok {
		o.AccessRestrictionEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("advertise"); ok {
		o.Advertise = attr.(bool)
	}
	if attr, ok := d.GetOk("default_action"); ok {
		o.DefaultAction = attr.(string)
	}
	if attr, ok := d.GetOk("template_id"); ok {
		o.TemplateID = attr.(string)
	}
	if attr, ok := d.GetOk("service_id"); ok {
		o.ServiceID = attr.(int)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("resource_type"); ok {
		o.ResourceType = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("flow_collection_enabled"); ok {
		o.FlowCollectionEnabled = attr.(string)
	}
	if attr, ok := d.GetOk("vn_id"); ok {
		o.VnId = attr.(int)
	}
	if attr, ok := d.GetOk("encryption"); ok {
		o.Encryption = attr.(string)
	}
	if attr, ok := d.GetOk("underlay"); ok {
		o.Underlay = attr.(bool)
	}
	if attr, ok := d.GetOk("underlay_enabled"); ok {
		o.UnderlayEnabled = attr.(string)
	}
	if attr, ok := d.GetOk("entity_state"); ok {
		o.EntityState = attr.(string)
	}
	if attr, ok := d.GetOk("policy_group_id"); ok {
		o.PolicyGroupID = attr.(int)
	}
	if attr, ok := d.GetOk("route_distinguisher"); ok {
		o.RouteDistinguisher = attr.(string)
	}
	if attr, ok := d.GetOk("route_target"); ok {
		o.RouteTarget = attr.(string)
	}
	if attr, ok := d.GetOk("split_subnet"); ok {
		o.SplitSubnet = attr.(bool)
	}
	if attr, ok := d.GetOk("proxy_arp"); ok {
		o.ProxyARP = attr.(bool)
	}
	if attr, ok := d.GetOk("use_global_mac"); ok {
		o.UseGlobalMAC = attr.(string)
	}
	if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
		o.AssociatedMulticastChannelMapID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_shared_network_resource_id"); ok {
		o.AssociatedSharedNetworkResourceID = attr.(string)
	}
	if attr, ok := d.GetOk("public"); ok {
		o.Public = attr.(bool)
	}
	if attr, ok := d.GetOk("multi_home_enabled"); ok {
		o.MultiHomeEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("multicast"); ok {
		o.Multicast = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_ipv6_address"); ok {
		o.DynamicIpv6Address = attr.(bool)
	}
	parent := &vspk.Zone{ID: d.Get("parent_zone").(string)}
	err := parent.CreateSubnet(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("ikegatewayconnections"); ok {
		o.AssignIKEGatewayConnections(attr.(vspk.IKEGatewayConnectionsList))
	}
	return resourceSubnetRead(d, m)
}

func resourceSubnetRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Subnet{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("pat_enabled", o.PATEnabled)
	d.Set("dhcp_relay_status", o.DHCPRelayStatus)
	d.Set("dpi", o.DPI)
	d.Set("ip_type", o.IPType)
	d.Set("ipv6_address", o.IPv6Address)
	d.Set("ipv6_gateway", o.IPv6Gateway)
	d.Set("maintenance_mode", o.MaintenanceMode)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("gateway", o.Gateway)
	d.Set("gateway_mac_address", o.GatewayMACAddress)
	d.Set("access_restriction_enabled", o.AccessRestrictionEnabled)
	d.Set("address", o.Address)
	d.Set("advertise", o.Advertise)
	d.Set("default_action", o.DefaultAction)
	d.Set("template_id", o.TemplateID)
	d.Set("service_id", o.ServiceID)
	d.Set("description", o.Description)
	d.Set("resource_type", o.ResourceType)
	d.Set("netmask", o.Netmask)
	d.Set("flow_collection_enabled", o.FlowCollectionEnabled)
	d.Set("vn_id", o.VnId)
	d.Set("encryption", o.Encryption)
	d.Set("underlay", o.Underlay)
	d.Set("underlay_enabled", o.UnderlayEnabled)
	d.Set("entity_scope", o.EntityScope)
	d.Set("entity_state", o.EntityState)
	d.Set("policy_group_id", o.PolicyGroupID)
	d.Set("route_distinguisher", o.RouteDistinguisher)
	d.Set("route_target", o.RouteTarget)
	d.Set("split_subnet", o.SplitSubnet)
	d.Set("proxy_arp", o.ProxyARP)
	d.Set("use_global_mac", o.UseGlobalMAC)
	d.Set("associated_multicast_channel_map_id", o.AssociatedMulticastChannelMapID)
	d.Set("associated_shared_network_resource_id", o.AssociatedSharedNetworkResourceID)
	d.Set("public", o.Public)
	d.Set("multi_home_enabled", o.MultiHomeEnabled)
	d.Set("multicast", o.Multicast)
	d.Set("external_id", o.ExternalID)
	d.Set("dynamic_ipv6_address", o.DynamicIpv6Address)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceSubnetUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Subnet{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("pat_enabled"); ok {
		o.PATEnabled = attr.(string)
	}
	if attr, ok := d.GetOk("dhcp_relay_status"); ok {
		o.DHCPRelayStatus = attr.(string)
	}
	if attr, ok := d.GetOk("dpi"); ok {
		o.DPI = attr.(string)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address"); ok {
		o.IPv6Address = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_gateway"); ok {
		o.IPv6Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("maintenance_mode"); ok {
		o.MaintenanceMode = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("gateway_mac_address"); ok {
		o.GatewayMACAddress = attr.(string)
	}
	if attr, ok := d.GetOk("access_restriction_enabled"); ok {
		o.AccessRestrictionEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("advertise"); ok {
		o.Advertise = attr.(bool)
	}
	if attr, ok := d.GetOk("default_action"); ok {
		o.DefaultAction = attr.(string)
	}
	if attr, ok := d.GetOk("template_id"); ok {
		o.TemplateID = attr.(string)
	}
	if attr, ok := d.GetOk("service_id"); ok {
		o.ServiceID = attr.(int)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("resource_type"); ok {
		o.ResourceType = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("flow_collection_enabled"); ok {
		o.FlowCollectionEnabled = attr.(string)
	}
	if attr, ok := d.GetOk("vn_id"); ok {
		o.VnId = attr.(int)
	}
	if attr, ok := d.GetOk("encryption"); ok {
		o.Encryption = attr.(string)
	}
	if attr, ok := d.GetOk("underlay"); ok {
		o.Underlay = attr.(bool)
	}
	if attr, ok := d.GetOk("underlay_enabled"); ok {
		o.UnderlayEnabled = attr.(string)
	}
	if attr, ok := d.GetOk("entity_state"); ok {
		o.EntityState = attr.(string)
	}
	if attr, ok := d.GetOk("policy_group_id"); ok {
		o.PolicyGroupID = attr.(int)
	}
	if attr, ok := d.GetOk("route_distinguisher"); ok {
		o.RouteDistinguisher = attr.(string)
	}
	if attr, ok := d.GetOk("route_target"); ok {
		o.RouteTarget = attr.(string)
	}
	if attr, ok := d.GetOk("split_subnet"); ok {
		o.SplitSubnet = attr.(bool)
	}
	if attr, ok := d.GetOk("proxy_arp"); ok {
		o.ProxyARP = attr.(bool)
	}
	if attr, ok := d.GetOk("use_global_mac"); ok {
		o.UseGlobalMAC = attr.(string)
	}
	if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
		o.AssociatedMulticastChannelMapID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_shared_network_resource_id"); ok {
		o.AssociatedSharedNetworkResourceID = attr.(string)
	}
	if attr, ok := d.GetOk("public"); ok {
		o.Public = attr.(bool)
	}
	if attr, ok := d.GetOk("multi_home_enabled"); ok {
		o.MultiHomeEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("multicast"); ok {
		o.Multicast = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_ipv6_address"); ok {
		o.DynamicIpv6Address = attr.(bool)
	}

	o.Save()

	return nil
}

func resourceSubnetDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Subnet{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
