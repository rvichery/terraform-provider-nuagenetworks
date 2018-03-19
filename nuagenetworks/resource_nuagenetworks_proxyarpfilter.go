package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceProxyARPFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceProxyARPFilterCreate,
		Read:   resourceProxyARPFilterRead,
		Update: resourceProxyARPFilterUpdate,
		Delete: resourceProxyARPFilterDelete,
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
			"ip_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IPV4",
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"min_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceProxyARPFilterCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize ProxyARPFilter object
	o := &vspk.ProxyARPFilter{
		MaxAddress: d.Get("max_address").(string),
		MinAddress: d.Get("min_address").(string),
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Subnet{ID: d.Get("parent_subnet").(string)}
	err := parent.CreateProxyARPFilter(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceProxyARPFilterRead(d, m)
}

func resourceProxyARPFilterRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ProxyARPFilter{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ip_type", o.IPType)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("max_address", o.MaxAddress)
	d.Set("min_address", o.MinAddress)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceProxyARPFilterUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ProxyARPFilter{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.MaxAddress = d.Get("max_address").(string)
	o.MinAddress = d.Get("min_address").(string)

	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceProxyARPFilterDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ProxyARPFilter{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
