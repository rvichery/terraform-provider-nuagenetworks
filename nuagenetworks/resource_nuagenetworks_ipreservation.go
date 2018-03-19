package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceIPReservation() *schema.Resource {
	return &schema.Resource{
		Create: resourceIPReservationCreate,
		Read:   resourceIPReservationRead,
		Update: resourceIPReservationUpdate,
		Delete: resourceIPReservationDelete,
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
			"mac": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"dynamic_allocation_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"parent_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIPReservationCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize IPReservation object
	o := &vspk.IPReservation{
		MAC:       d.Get("mac").(string),
		IPAddress: d.Get("ip_address").(string),
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_allocation_enabled"); ok {
		o.DynamicAllocationEnabled = attr.(bool)
	}
	parent := &vspk.Subnet{ID: d.Get("parent_subnet").(string)}
	err := parent.CreateIPReservation(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceIPReservationRead(d, m)
}

func resourceIPReservationRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IPReservation{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("mac", o.MAC)
	d.Set("ip_address", o.IPAddress)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)
	d.Set("dynamic_allocation_enabled", o.DynamicAllocationEnabled)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceIPReservationUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IPReservation{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.MAC = d.Get("mac").(string)
	o.IPAddress = d.Get("ip_address").(string)

	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("dynamic_allocation_enabled"); ok {
		o.DynamicAllocationEnabled = attr.(bool)
	}

	o.Save()

	return nil
}

func resourceIPReservationDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IPReservation{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
