package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
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
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dpi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "INHERITED",
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"split_subnet": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"use_global_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_multicast_channel_map_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_ipv6_address": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"parent_zone_template": &schema.Schema{
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
	if attr, ok := d.GetOk("multicast"); ok {
		o.Multicast = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_ipv6_address"); ok {
		o.DynamicIpv6Address = attr.(bool)
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
