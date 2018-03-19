package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceZFBAutoAssignment() *schema.Resource {
	return &schema.Resource{
		Create: resourceZFBAutoAssignmentCreate,
		Read:   resourceZFBAutoAssignmentRead,
		Update: resourceZFBAutoAssignmentUpdate,
		Delete: resourceZFBAutoAssignmentDelete,
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
			"zfb_match_attribute": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"zfb_match_attribute_values": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"associated_enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_enterprise_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceZFBAutoAssignmentCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize ZFBAutoAssignment object
	o := &vspk.ZFBAutoAssignment{
		Name:     d.Get("name").(string),
		Priority: d.Get("priority").(int),
	}
	if attr, ok := d.GetOk("zfb_match_attribute"); ok {
		o.ZFBMatchAttribute = attr.(string)
	}
	if attr, ok := d.GetOk("zfb_match_attribute_values"); ok {
		o.ZFBMatchAttributeValues = attr.([]interface{})
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_name"); ok {
		o.AssociatedEnterpriseName = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateZFBAutoAssignment(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceZFBAutoAssignmentRead(d, m)
}

func resourceZFBAutoAssignmentRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ZFBAutoAssignment{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("zfb_match_attribute", o.ZFBMatchAttribute)
	d.Set("zfb_match_attribute_values", o.ZFBMatchAttributeValues)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("priority", o.Priority)
	d.Set("associated_enterprise_id", o.AssociatedEnterpriseID)
	d.Set("associated_enterprise_name", o.AssociatedEnterpriseName)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceZFBAutoAssignmentUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ZFBAutoAssignment{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Priority = d.Get("priority").(int)

	if attr, ok := d.GetOk("zfb_match_attribute"); ok {
		o.ZFBMatchAttribute = attr.(string)
	}
	if attr, ok := d.GetOk("zfb_match_attribute_values"); ok {
		o.ZFBMatchAttributeValues = attr.([]interface{})
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_enterprise_name"); ok {
		o.AssociatedEnterpriseName = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceZFBAutoAssignmentDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ZFBAutoAssignment{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
