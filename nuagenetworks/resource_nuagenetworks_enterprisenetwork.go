package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceEnterpriseNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceEnterpriseNetworkCreate,
		Read:   resourceEnterpriseNetworkRead,
		Update: resourceEnterpriseNetworkUpdate,
		Delete: resourceEnterpriseNetworkDelete,
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
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address": &schema.Schema{
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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEnterpriseNetworkCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize EnterpriseNetwork object
	o := &vspk.EnterpriseNetwork{
		Name:    d.Get("name").(string),
		Address: d.Get("address").(string),
		Netmask: d.Get("netmask").(string),
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address"); ok {
		o.IPv6Address = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateEnterpriseNetwork(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("networkmacrogroups"); ok {
		o.AssignNetworkMacroGroups(attr.(vspk.NetworkMacroGroupsList))
	}
	return resourceEnterpriseNetworkRead(d, m)
}

func resourceEnterpriseNetworkRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterpriseNetwork{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ip_type", o.IPType)
	d.Set("ipv6_address", o.IPv6Address)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("address", o.Address)
	d.Set("netmask", o.Netmask)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceEnterpriseNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterpriseNetwork{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Address = d.Get("address").(string)
	o.Netmask = d.Get("netmask").(string)

	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address"); ok {
		o.IPv6Address = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceEnterpriseNetworkDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EnterpriseNetwork{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
