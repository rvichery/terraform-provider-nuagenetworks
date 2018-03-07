package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceL2DomainTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceL2DomainTemplateCreate,
		Read:   resourceL2DomainTemplateRead,
		Update: resourceL2DomainTemplateUpdate,
		Delete: resourceL2DomainTemplateDelete,
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
			"dhcp_managed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"dpi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DISABLED",
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
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"entity_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"policy_change_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_global_mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DISABLED",
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
				Default:  false,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceL2DomainTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize L2DomainTemplate object
	o := &vspk.L2DomainTemplate{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("dhcp_managed"); ok {
		o.DHCPManaged = attr.(bool)
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
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("encryption"); ok {
		o.Encryption = attr.(string)
	}
	if attr, ok := d.GetOk("entity_state"); ok {
		o.EntityState = attr.(string)
	}
	if attr, ok := d.GetOk("policy_change_status"); ok {
		o.PolicyChangeStatus = attr.(string)
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
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateL2DomainTemplate(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceL2DomainTemplateRead(d, m)
}

func resourceL2DomainTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L2DomainTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("dhcp_managed", o.DHCPManaged)
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
	d.Set("entity_state", o.EntityState)
	d.Set("policy_change_status", o.PolicyChangeStatus)
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

func resourceL2DomainTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L2DomainTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("dhcp_managed"); ok {
		o.DHCPManaged = attr.(bool)
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
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("encryption"); ok {
		o.Encryption = attr.(string)
	}
	if attr, ok := d.GetOk("entity_state"); ok {
		o.EntityState = attr.(string)
	}
	if attr, ok := d.GetOk("policy_change_status"); ok {
		o.PolicyChangeStatus = attr.(string)
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

func resourceL2DomainTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.L2DomainTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
