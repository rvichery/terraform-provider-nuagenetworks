package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceApplicationperformancemanagementbinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationperformancemanagementbindingCreate,
		Read:   resourceApplicationperformancemanagementbindingRead,
		Update: resourceApplicationperformancemanagementbindingUpdate,
		Delete: resourceApplicationperformancemanagementbindingDelete,
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
			"associated_application_performance_management_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceApplicationperformancemanagementbindingCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Applicationperformancemanagementbinding object
	o := &vspk.Applicationperformancemanagementbinding{
		AssociatedApplicationPerformanceManagementID: d.Get("associated_application_performance_management_id").(string),
	}
	if attr, ok := d.GetOk("read_only"); ok {
		ReadOnly := attr.(bool)
		o.ReadOnly = &ReadOnly
	}
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	err := parent.CreateApplicationperformancemanagementbinding(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceApplicationperformancemanagementbindingRead(d, m)
}

func resourceApplicationperformancemanagementbindingRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Applicationperformancemanagementbinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("read_only", o.ReadOnly)
	d.Set("priority", o.Priority)
	d.Set("associated_application_performance_management_id", o.AssociatedApplicationPerformanceManagementID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceApplicationperformancemanagementbindingUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Applicationperformancemanagementbinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.AssociatedApplicationPerformanceManagementID = d.Get("associated_application_performance_management_id").(string)

	if attr, ok := d.GetOk("read_only"); ok {
		ReadOnly := attr.(bool)
		o.ReadOnly = &ReadOnly
	}

	o.Save()

	return nil
}

func resourceApplicationperformancemanagementbindingDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Applicationperformancemanagementbinding{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
