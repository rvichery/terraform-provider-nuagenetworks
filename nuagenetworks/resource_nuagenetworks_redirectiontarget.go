package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceRedirectionTarget() *schema.Resource {
	return &schema.Resource{
		Create: resourceRedirectionTargetCreate,
		Read:   resourceRedirectionTargetRead,
		Update: resourceRedirectionTargetUpdate,
		Delete: resourceRedirectionTargetDelete,
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
			"esi": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"redundancy_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"template_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"virtual_network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_point_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceRedirectionTargetCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize RedirectionTarget object
	o := &vspk.RedirectionTarget{
		Name:         d.Get("name").(string),
		EndPointType: d.Get("end_point_type").(string),
	}
	if attr, ok := d.GetOk("esi"); ok {
		o.ESI = attr.(string)
	}
	if attr, ok := d.GetOk("redundancy_enabled"); ok {
		RedundancyEnabled := attr.(bool)
		o.RedundancyEnabled = &RedundancyEnabled
	}
	if attr, ok := d.GetOk("template_id"); ok {
		o.TemplateID = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("virtual_network_id"); ok {
		o.VirtualNetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("trigger_type"); ok {
		o.TriggerType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateRedirectionTarget(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		err := parent.CreateRedirectionTarget(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("vports"); ok {
		o.AssignVPorts(attr.(vspk.VPortsList))
	}
	return resourceRedirectionTargetRead(d, m)
}

func resourceRedirectionTargetRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.RedirectionTarget{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("esi", o.ESI)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("redundancy_enabled", o.RedundancyEnabled)
	d.Set("template_id", o.TemplateID)
	d.Set("description", o.Description)
	d.Set("virtual_network_id", o.VirtualNetworkID)
	d.Set("end_point_type", o.EndPointType)
	d.Set("entity_scope", o.EntityScope)
	d.Set("trigger_type", o.TriggerType)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceRedirectionTargetUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.RedirectionTarget{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.EndPointType = d.Get("end_point_type").(string)

	if attr, ok := d.GetOk("esi"); ok {
		o.ESI = attr.(string)
	}
	if attr, ok := d.GetOk("redundancy_enabled"); ok {
		RedundancyEnabled := attr.(bool)
		o.RedundancyEnabled = &RedundancyEnabled
	}
	if attr, ok := d.GetOk("template_id"); ok {
		o.TemplateID = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("virtual_network_id"); ok {
		o.VirtualNetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("trigger_type"); ok {
		o.TriggerType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceRedirectionTargetDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.RedirectionTarget{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
