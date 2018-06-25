package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceBGPProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceBGPProfileCreate,
		Read:   resourceBGPProfileRead,
		Update: resourceBGPProfileUpdate,
		Delete: resourceBGPProfileDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"dampening_half_life": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
			"dampening_max_suppress": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  60,
			},
			"dampening_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dampening_reuse": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  750,
			},
			"dampening_suppress": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  3000,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_export_routing_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_import_routing_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceBGPProfileCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize BGPProfile object
	o := &vspk.BGPProfile{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("dampening_half_life"); ok {
		DampeningHalfLife := attr.(int)
		o.DampeningHalfLife = &DampeningHalfLife
	}
	if attr, ok := d.GetOk("dampening_max_suppress"); ok {
		DampeningMaxSuppress := attr.(int)
		o.DampeningMaxSuppress = &DampeningMaxSuppress
	}
	if attr, ok := d.GetOk("dampening_name"); ok {
		o.DampeningName = attr.(string)
	}
	if attr, ok := d.GetOk("dampening_reuse"); ok {
		DampeningReuse := attr.(int)
		o.DampeningReuse = &DampeningReuse
	}
	if attr, ok := d.GetOk("dampening_suppress"); ok {
		DampeningSuppress := attr.(int)
		o.DampeningSuppress = &DampeningSuppress
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
		o.AssociatedExportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
		o.AssociatedImportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateBGPProfile(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceBGPProfileRead(d, m)
}

func resourceBGPProfileRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BGPProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("dampening_half_life", o.DampeningHalfLife)
	d.Set("dampening_max_suppress", o.DampeningMaxSuppress)
	d.Set("dampening_name", o.DampeningName)
	d.Set("dampening_reuse", o.DampeningReuse)
	d.Set("dampening_suppress", o.DampeningSuppress)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_export_routing_policy_id", o.AssociatedExportRoutingPolicyID)
	d.Set("associated_import_routing_policy_id", o.AssociatedImportRoutingPolicyID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceBGPProfileUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BGPProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("dampening_half_life"); ok {
		DampeningHalfLife := attr.(int)
		o.DampeningHalfLife = &DampeningHalfLife
	}
	if attr, ok := d.GetOk("dampening_max_suppress"); ok {
		DampeningMaxSuppress := attr.(int)
		o.DampeningMaxSuppress = &DampeningMaxSuppress
	}
	if attr, ok := d.GetOk("dampening_name"); ok {
		o.DampeningName = attr.(string)
	}
	if attr, ok := d.GetOk("dampening_reuse"); ok {
		DampeningReuse := attr.(int)
		o.DampeningReuse = &DampeningReuse
	}
	if attr, ok := d.GetOk("dampening_suppress"); ok {
		DampeningSuppress := attr.(int)
		o.DampeningSuppress = &DampeningSuppress
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("associated_export_routing_policy_id"); ok {
		o.AssociatedExportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_import_routing_policy_id"); ok {
		o.AssociatedImportRoutingPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceBGPProfileDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.BGPProfile{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
