package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceCSNATPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceCSNATPoolCreate,
		Read:   resourceCSNATPoolRead,
		Update: resourceCSNATPoolUpdate,
		Delete: resourceCSNATPoolDelete,
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
			"end_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"start_address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_link": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCSNATPoolCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize CSNATPool object
	o := &vspk.CSNATPool{
		EndAddress:   d.Get("end_address").(string),
		StartAddress: d.Get("start_address").(string),
	}
	parent := &vspk.Link{ID: d.Get("parent_link").(string)}
	err := parent.CreateCSNATPool(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceCSNATPoolRead(d, m)
}

func resourceCSNATPoolRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CSNATPool{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("end_address", o.EndAddress)
	d.Set("start_address", o.StartAddress)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceCSNATPoolUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CSNATPool{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.EndAddress = d.Get("end_address").(string)
	o.StartAddress = d.Get("start_address").(string)

	o.Save()

	return nil
}

func resourceCSNATPoolDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CSNATPool{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
