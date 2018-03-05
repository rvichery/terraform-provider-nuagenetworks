package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceConnectionendpoint() *schema.Resource {
	return &schema.Resource{
		Create: resourceConnectionendpointCreate,
		Read:   resourceConnectionendpointRead,
		Update: resourceConnectionendpointUpdate,
		Delete: resourceConnectionendpointDelete,

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
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IPV4",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_point_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SOURCE",
			},
			"parent_infrastructure_access_profile": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceConnectionendpointCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Connectionendpoint object
	o := &vspk.Connectionendpoint{
		IPAddress: d.Get("ip_address").(string),
		Name:      d.Get("name").(string),
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("end_point_type"); ok {
		o.EndPointType = attr.(string)
	}
	parent := &vspk.InfrastructureAccessProfile{ID: d.Get("parent_infrastructure_access_profile").(string)}
	err := parent.CreateConnectionendpoint(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceConnectionendpointRead(d, m)
}

func resourceConnectionendpointRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Connectionendpoint{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ip_address", o.IPAddress)
	d.Set("ip_type", o.IPType)
	d.Set("name", o.Name)
	d.Set("description", o.Description)
	d.Set("end_point_type", o.EndPointType)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceConnectionendpointUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Connectionendpoint{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.IPAddress = d.Get("ip_address").(string)
	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("end_point_type"); ok {
		o.EndPointType = attr.(string)
	}

	o.Save()

	return nil
}

func resourceConnectionendpointDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Connectionendpoint{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
