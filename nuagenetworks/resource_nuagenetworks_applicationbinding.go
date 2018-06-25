package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceApplicationBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationBindingCreate,
		Read:   resourceApplicationBindingRead,
		Update: resourceApplicationBindingUpdate,
		Delete: resourceApplicationBindingDelete,
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
			"read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_application_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_applicationperformancemanagement": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApplicationBindingCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize ApplicationBinding object
	o := &vspk.ApplicationBinding{
		AssociatedApplicationID: d.Get("associated_application_id").(string),
	}
	if attr, ok := d.GetOk("read_only"); ok {
		ReadOnly := attr.(bool)
		o.ReadOnly = &ReadOnly
	}
	parent := &vspk.Applicationperformancemanagement{ID: d.Get("parent_applicationperformancemanagement").(string)}
	err := parent.CreateApplicationBinding(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceApplicationBindingRead(d, m)
}

func resourceApplicationBindingRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ApplicationBinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("read_only", o.ReadOnly)
	d.Set("priority", o.Priority)
	d.Set("associated_application_id", o.AssociatedApplicationID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceApplicationBindingUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ApplicationBinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.AssociatedApplicationID = d.Get("associated_application_id").(string)

	if attr, ok := d.GetOk("read_only"); ok {
		ReadOnly := attr.(bool)
		o.ReadOnly = &ReadOnly
	}

	o.Save()

	return nil
}

func resourceApplicationBindingDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ApplicationBinding{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
