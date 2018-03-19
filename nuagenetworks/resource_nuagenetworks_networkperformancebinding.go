package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceNetworkPerformanceBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetworkPerformanceBindingCreate,
		Read:   resourceNetworkPerformanceBindingRead,
		Update: resourceNetworkPerformanceBindingUpdate,
		Delete: resourceNetworkPerformanceBindingDelete,
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
				Optional: true,
				Default:  false,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"associated_network_measurement_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l2_domain"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain"},
			},
		},
	}
}

func resourceNetworkPerformanceBindingCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize NetworkPerformanceBinding object
	o := &vspk.NetworkPerformanceBinding{}
	if attr, ok := d.GetOk("read_only"); ok {
		o.ReadOnly = attr.(bool)
	}
	if attr, ok := d.GetOk("associated_network_measurement_id"); ok {
		o.AssociatedNetworkMeasurementID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateNetworkPerformanceBinding(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		err := parent.CreateNetworkPerformanceBinding(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceNetworkPerformanceBindingRead(d, m)
}

func resourceNetworkPerformanceBindingRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NetworkPerformanceBinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("read_only", o.ReadOnly)
	d.Set("priority", o.Priority)
	d.Set("associated_network_measurement_id", o.AssociatedNetworkMeasurementID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceNetworkPerformanceBindingUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NetworkPerformanceBinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("read_only"); ok {
		o.ReadOnly = attr.(bool)
	}
	if attr, ok := d.GetOk("associated_network_measurement_id"); ok {
		o.AssociatedNetworkMeasurementID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceNetworkPerformanceBindingDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NetworkPerformanceBinding{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
