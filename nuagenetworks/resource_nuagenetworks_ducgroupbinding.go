package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceDUCGroupBinding() *schema.Resource {
	return &schema.Resource{
		Create: resourceDUCGroupBindingCreate,
		Read:   resourceDUCGroupBindingRead,
		Update: resourceDUCGroupBindingUpdate,
		Delete: resourceDUCGroupBindingDelete,
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
			"one_way_delay": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  50,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"associated_duc_group_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_nsg_group": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDUCGroupBindingCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize DUCGroupBinding object
	o := &vspk.DUCGroupBinding{}
	if attr, ok := d.GetOk("one_way_delay"); ok {
		o.OneWayDelay = attr.(int)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("associated_duc_group_id"); ok {
		o.AssociatedDUCGroupID = attr.(string)
	}
	parent := &vspk.NSGGroup{ID: d.Get("parent_nsg_group").(string)}
	err := parent.CreateDUCGroupBinding(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceDUCGroupBindingRead(d, m)
}

func resourceDUCGroupBindingRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DUCGroupBinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("one_way_delay", o.OneWayDelay)
	d.Set("priority", o.Priority)
	d.Set("associated_duc_group_id", o.AssociatedDUCGroupID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceDUCGroupBindingUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DUCGroupBinding{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("one_way_delay"); ok {
		o.OneWayDelay = attr.(int)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("associated_duc_group_id"); ok {
		o.AssociatedDUCGroupID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceDUCGroupBindingDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DUCGroupBinding{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
