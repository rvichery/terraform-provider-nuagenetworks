package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceSubnetTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceSubnetTemplateCreate,
		Read:   resourceSubnetTemplateRead,
		Update: resourceSubnetTemplateUpdate,
		Delete: resourceSubnetTemplateDelete,
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
			"dpi": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INHERITED",
			},
			"ip_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Required: true,
			},
			"encryption": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"split_subnet": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"use_global_mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_multicast_channel_map_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multicast": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_ipv6_address": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"parent_zone_template": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceSubnetTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize SubnetTemplate object
	o := &vspk.SubnetTemplate{
		Name:    d.Get("name").(string),
		Address: d.Get("address").(string),
		Netmask: d.Get("netmask").(string),
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
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("encryption"); ok {
		o.Encryption = attr.(string)
	}
	if attr, ok := d.GetOk("split_subnet"); ok {
		SplitSubnet := attr.(bool)
		o.SplitSubnet = &SplitSubnet
	}
	if attr, ok := d.GetOk("proxy_arp"); ok {
		ProxyARP := attr.(bool)
		o.ProxyARP = &ProxyARP
	}
	if attr, ok := d.GetOk("use_global_mac"); ok {
		o.UseGlobalMAC = attr.(string)
	}
	if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
		o.AssociatedMulticastChannelMapID = attr.(string)
	}
	if attr, ok := d.GetOk("multicast"); ok {
		o.Multicast = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_ipv6_address"); ok {
		DynamicIpv6Address := attr.(bool)
		o.DynamicIpv6Address = &DynamicIpv6Address
	}
	parent := &vspk.ZoneTemplate{ID: d.Get("parent_zone_template").(string)}
	err := parent.CreateSubnetTemplate(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceSubnetTemplateRead(d, m)
}

func resourceSubnetTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SubnetTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("dpi", o.DPI)
	d.Set("ip_type", o.IPType)
	d.Set("ipv6_address", o.IPv6Address)
	d.Set("ipv6_gateway", o.IPv6Gateway)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("gateway", o.Gateway)
	d.Set("address", o.Address)
	d.Set("description", o.Description)
	d.Set("netmask", o.Netmask)
	d.Set("encryption", o.Encryption)
	d.Set("entity_scope", o.EntityScope)
	d.Set("split_subnet", o.SplitSubnet)
	d.Set("proxy_arp", o.ProxyARP)
	d.Set("use_global_mac", o.UseGlobalMAC)
	d.Set("associated_multicast_channel_map_id", o.AssociatedMulticastChannelMapID)
	d.Set("multicast", o.Multicast)
	d.Set("external_id", o.ExternalID)
	d.Set("dynamic_ipv6_address", o.DynamicIpv6Address)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceSubnetTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SubnetTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Address = d.Get("address").(string)
	o.Netmask = d.Get("netmask").(string)

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
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("encryption"); ok {
		o.Encryption = attr.(string)
	}
	if attr, ok := d.GetOk("split_subnet"); ok {
		SplitSubnet := attr.(bool)
		o.SplitSubnet = &SplitSubnet
	}
	if attr, ok := d.GetOk("proxy_arp"); ok {
		ProxyARP := attr.(bool)
		o.ProxyARP = &ProxyARP
	}
	if attr, ok := d.GetOk("use_global_mac"); ok {
		o.UseGlobalMAC = attr.(string)
	}
	if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
		o.AssociatedMulticastChannelMapID = attr.(string)
	}
	if attr, ok := d.GetOk("multicast"); ok {
		o.Multicast = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_ipv6_address"); ok {
		DynamicIpv6Address := attr.(bool)
		o.DynamicIpv6Address = &DynamicIpv6Address
	}

	o.Save()

	return nil
}

func resourceSubnetTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SubnetTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
