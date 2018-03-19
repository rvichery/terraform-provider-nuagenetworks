package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceCustomProperty() *schema.Resource {
	return &schema.Resource{
		Create: resourceCustomPropertyCreate,
		Read:   resourceCustomPropertyRead,
		Update: resourceCustomPropertyUpdate,
		Delete: resourceCustomPropertyDelete,
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
			"attribute_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attribute_value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_uplink_connection": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCustomPropertyCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize CustomProperty object
	o := &vspk.CustomProperty{}
	if attr, ok := d.GetOk("attribute_name"); ok {
		o.AttributeName = attr.(string)
	}
	if attr, ok := d.GetOk("attribute_value"); ok {
		o.AttributeValue = attr.(string)
	}
	parent := &vspk.UplinkConnection{ID: d.Get("parent_uplink_connection").(string)}
	err := parent.CreateCustomProperty(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceCustomPropertyRead(d, m)
}

func resourceCustomPropertyRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CustomProperty{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("attribute_name", o.AttributeName)
	d.Set("attribute_value", o.AttributeValue)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceCustomPropertyUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CustomProperty{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("attribute_name"); ok {
		o.AttributeName = attr.(string)
	}
	if attr, ok := d.GetOk("attribute_value"); ok {
		o.AttributeValue = attr.(string)
	}

	o.Save()

	return nil
}

func resourceCustomPropertyDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CustomProperty{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
