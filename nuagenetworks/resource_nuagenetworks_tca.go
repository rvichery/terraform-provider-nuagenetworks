package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceTCA() *schema.Resource {
	return &schema.Resource{
		Create: resourceTCACreate,
		Read:   resourceTCARead,
		Update: resourceTCAUpdate,
		Delete: resourceTCADelete,
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
			"url_end_point": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"target_policy_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"throttle_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"disable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"display_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_zone": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain"},
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_vport", "parent_subnet", "parent_l2_domain"},
			},
			"parent_vport": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_subnet", "parent_l2_domain"},
			},
			"parent_subnet": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_l2_domain"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet"},
			},
		},
	}
}

func resourceTCACreate(d *schema.ResourceData, m interface{}) error {

	// Initialize TCA object
	o := &vspk.TCA{
		Name:      d.Get("name").(string),
		Action:    d.Get("action").(string),
		Period:    d.Get("period").(int),
		Metric:    d.Get("metric").(string),
		Threshold: d.Get("threshold").(int),
		Type:      d.Get("type").(string),
	}
	if attr, ok := d.GetOk("url_end_point"); ok {
		o.URLEndPoint = attr.(string)
	}
	if attr, ok := d.GetOk("target_policy_group_id"); ok {
		o.TargetPolicyGroupID = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("throttle_time"); ok {
		o.ThrottleTime = attr.(int)
	}
	if attr, ok := d.GetOk("disable"); ok {
		o.Disable = attr.(bool)
	}
	if attr, ok := d.GetOk("display_status"); ok {
		o.DisplayStatus = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		err := parent.CreateTCA(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateTCA(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		err := parent.CreateTCA(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		err := parent.CreateTCA(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		err := parent.CreateTCA(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceTCARead(d, m)
}

func resourceTCARead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.TCA{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("url_end_point", o.URLEndPoint)
	d.Set("name", o.Name)
	d.Set("target_policy_group_id", o.TargetPolicyGroupID)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("action", o.Action)
	d.Set("period", o.Period)
	d.Set("description", o.Description)
	d.Set("metric", o.Metric)
	d.Set("threshold", o.Threshold)
	d.Set("throttle_time", o.ThrottleTime)
	d.Set("disable", o.Disable)
	d.Set("display_status", o.DisplayStatus)
	d.Set("entity_scope", o.EntityScope)

	d.Set("status", o.Status)
	d.Set("external_id", o.ExternalID)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceTCAUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.TCA{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Action = d.Get("action").(string)
	o.Period = d.Get("period").(int)
	o.Metric = d.Get("metric").(string)
	o.Threshold = d.Get("threshold").(int)
	o.Type = d.Get("type").(string)

	if attr, ok := d.GetOk("url_end_point"); ok {
		o.URLEndPoint = attr.(string)
	}
	if attr, ok := d.GetOk("target_policy_group_id"); ok {
		o.TargetPolicyGroupID = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("throttle_time"); ok {
		o.ThrottleTime = attr.(int)
	}
	if attr, ok := d.GetOk("disable"); ok {
		o.Disable = attr.(bool)
	}
	if attr, ok := d.GetOk("display_status"); ok {
		o.DisplayStatus = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceTCADelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.TCA{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
