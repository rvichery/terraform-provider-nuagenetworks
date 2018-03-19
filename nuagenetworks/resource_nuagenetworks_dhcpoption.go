package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceDHCPOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPOptionCreate,
		Read:   resourceDHCPOptionRead,
		Update: resourceDHCPOptionUpdate,
		Delete: resourceDHCPOptionDelete,
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
			"value": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"actual_type": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"actual_values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"length": {
				Type:     schema.TypeString,
				Optional: true,
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
			"type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_shared_network_resource": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain"},
			},
			"parent_zone": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_zone", "parent_vport", "parent_subnet", "parent_l2_domain"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_zone", "parent_domain", "parent_subnet", "parent_l2_domain"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_zone", "parent_domain", "parent_vport", "parent_l2_domain"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_zone", "parent_domain", "parent_vport", "parent_subnet"},
			},
		},
	}
}

func resourceDHCPOptionCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize DHCPOption object
	o := &vspk.DHCPOption{}
	if attr, ok := d.GetOk("value"); ok {
		o.Value = attr.(string)
	}
	if attr, ok := d.GetOk("actual_type"); ok {
		o.ActualType = attr.(int)
	}
	if attr, ok := d.GetOk("actual_values"); ok {
		o.ActualValues = attr.([]interface{})
	}
	if attr, ok := d.GetOk("length"); ok {
		o.Length = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		o.Type = attr.(string)
	}
	if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		err := parent.CreateDHCPOption(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		err := parent.CreateDHCPOption(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateDHCPOption(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		err := parent.CreateDHCPOption(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		err := parent.CreateDHCPOption(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		err := parent.CreateDHCPOption(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceDHCPOptionRead(d, m)
}

func resourceDHCPOptionRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DHCPOption{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("value", o.Value)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("actual_type", o.ActualType)
	d.Set("actual_values", o.ActualValues)
	d.Set("length", o.Length)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceDHCPOptionUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DHCPOption{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("value"); ok {
		o.Value = attr.(string)
	}
	if attr, ok := d.GetOk("actual_type"); ok {
		o.ActualType = attr.(int)
	}
	if attr, ok := d.GetOk("actual_values"); ok {
		o.ActualValues = attr.([]interface{})
	}
	if attr, ok := d.GetOk("length"); ok {
		o.Length = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("type"); ok {
		o.Type = attr.(string)
	}

	o.Save()

	return nil
}

func resourceDHCPOptionDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DHCPOption{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
