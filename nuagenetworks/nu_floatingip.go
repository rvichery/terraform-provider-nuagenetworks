package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceFloatingIp() *schema.Resource {
	return &schema.Resource{
		Create: resourceFloatingIpCreate,
		Read:   resourceFloatingIpRead,
		Update: resourceFloatingIpUpdate,
		Delete: resourceFloatingIpDelete,

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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"access_control": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assigned": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"assigned_to_object_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_shared_network_resource_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFloatingIpCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize FloatingIp object
	o := &vspk.FloatingIp{
		AssociatedSharedNetworkResourceID: d.Get("associated_shared_network_resource_id").(string),
	}
	if attr, ok := d.GetOk("access_control"); ok {
		o.AccessControl = attr.(bool)
	}
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("assigned"); ok {
		o.Assigned = attr.(bool)
	}
	if attr, ok := d.GetOk("assigned_to_object_type"); ok {
		o.AssignedToObjectType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	err := parent.CreateFloatingIp(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceFloatingIpRead(d, m)
}

func resourceFloatingIpRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIp{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("access_control", o.AccessControl)
	d.Set("address", o.Address)
	d.Set("entity_scope", o.EntityScope)
	d.Set("assigned", o.Assigned)
	d.Set("assigned_to_object_type", o.AssignedToObjectType)
	d.Set("associated_shared_network_resource_id", o.AssociatedSharedNetworkResourceID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceFloatingIpUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIp{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.AssociatedSharedNetworkResourceID = d.Get("associated_shared_network_resource_id").(string)

	if attr, ok := d.GetOk("access_control"); ok {
		o.AccessControl = attr.(bool)
	}
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("assigned"); ok {
		o.Assigned = attr.(bool)
	}
	if attr, ok := d.GetOk("assigned_to_object_type"); ok {
		o.AssignedToObjectType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceFloatingIpDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIp{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
