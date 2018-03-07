package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceMirrorDestination() *schema.Resource {
	return &schema.Resource{
		Create: resourceMirrorDestinationCreate,
		Read:   resourceMirrorDestinationRead,
		Update: resourceMirrorDestinationUpdate,
		Delete: resourceMirrorDestinationDelete,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"destination_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
		},
	}
}

func resourceMirrorDestinationCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize MirrorDestination object
	o := &vspk.MirrorDestination{}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("service_id"); ok {
		o.ServiceId = attr.(int)
	}
	if attr, ok := d.GetOk("destination_ip"); ok {
		o.DestinationIp = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateMirrorDestination(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceMirrorDestinationRead(d, m)
}

func resourceMirrorDestinationRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.MirrorDestination{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("service_id", o.ServiceId)
	d.Set("destination_ip", o.DestinationIp)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceMirrorDestinationUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.MirrorDestination{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("service_id"); ok {
		o.ServiceId = attr.(int)
	}
	if attr, ok := d.GetOk("destination_ip"); ok {
		o.DestinationIp = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceMirrorDestinationDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.MirrorDestination{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
