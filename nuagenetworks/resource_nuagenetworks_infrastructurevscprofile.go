package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceInfrastructureVscProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceInfrastructureVscProfileCreate,
		Read:   resourceInfrastructureVscProfileRead,
		Update: resourceInfrastructureVscProfileUpdate,
		Delete: resourceInfrastructureVscProfileDelete,
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
			"second_controller": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"first_controller": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"probe_interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  5000,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceInfrastructureVscProfileCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize InfrastructureVscProfile object
	o := &vspk.InfrastructureVscProfile{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("second_controller"); ok {
		o.SecondController = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("first_controller"); ok {
		o.FirstController = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("probe_interval"); ok {
		ProbeInterval := attr.(int)
		o.ProbeInterval = &ProbeInterval
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateInfrastructureVscProfile(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceInfrastructureVscProfileRead(d, m)
}

func resourceInfrastructureVscProfileRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureVscProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("second_controller", o.SecondController)
	d.Set("description", o.Description)
	d.Set("first_controller", o.FirstController)
	d.Set("enterprise_id", o.EnterpriseID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("probe_interval", o.ProbeInterval)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceInfrastructureVscProfileUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureVscProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("second_controller"); ok {
		o.SecondController = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("first_controller"); ok {
		o.FirstController = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("probe_interval"); ok {
		ProbeInterval := attr.(int)
		o.ProbeInterval = &ProbeInterval
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceInfrastructureVscProfileDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureVscProfile{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
