package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceQosPolicer() *schema.Resource {
	return &schema.Resource{
		Create: resourceQosPolicerCreate,
		Read:   resourceQosPolicerRead,
		Update: resourceQosPolicerUpdate,
		Delete: resourceQosPolicerDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rate": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"burst": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceQosPolicerCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize QosPolicer object
	o := &vspk.QosPolicer{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("rate"); ok {
		Rate := attr.(int)
		o.Rate = &Rate
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("burst"); ok {
		Burst := attr.(int)
		o.Burst = &Burst
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateQosPolicer(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceQosPolicerRead(d, m)
}

func resourceQosPolicerRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.QosPolicer{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("rate", o.Rate)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("burst", o.Burst)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceQosPolicerUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.QosPolicer{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("rate"); ok {
		Rate := attr.(int)
		o.Rate = &Rate
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("burst"); ok {
		Burst := attr.(int)
		o.Burst = &Burst
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceQosPolicerDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.QosPolicer{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
