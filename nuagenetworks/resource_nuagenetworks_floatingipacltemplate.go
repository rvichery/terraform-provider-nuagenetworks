package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceFloatingIPACLTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceFloatingIPACLTemplateCreate,
		Read:   resourceFloatingIPACLTemplateRead,
		Update: resourceFloatingIPACLTemplateUpdate,
		Delete: resourceFloatingIPACLTemplateDelete,
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
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_allow_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_allow_non_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_live_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"auto_generate_priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain_template"},
			},
			"parent_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain"},
			},
		},
	}
}

func resourceFloatingIPACLTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize FloatingIPACLTemplate object
	o := &vspk.FloatingIPACLTemplate{}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("active"); ok {
		o.Active = attr.(bool)
	}
	if attr, ok := d.GetOk("default_allow_ip"); ok {
		o.DefaultAllowIP = attr.(bool)
	}
	if attr, ok := d.GetOk("default_allow_non_ip"); ok {
		o.DefaultAllowNonIP = attr.(bool)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("policy_state"); ok {
		o.PolicyState = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("priority_type"); ok {
		o.PriorityType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("auto_generate_priority"); ok {
		o.AutoGeneratePriority = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateFloatingIPACLTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateFloatingIPACLTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		err := parent.CreateFloatingIPACLTemplate(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceFloatingIPACLTemplateRead(d, m)
}

func resourceFloatingIPACLTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIPACLTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("active", o.Active)
	d.Set("default_allow_ip", o.DefaultAllowIP)
	d.Set("default_allow_non_ip", o.DefaultAllowNonIP)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("policy_state", o.PolicyState)
	d.Set("priority", o.Priority)
	d.Set("priority_type", o.PriorityType)
	d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
	d.Set("auto_generate_priority", o.AutoGeneratePriority)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceFloatingIPACLTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIPACLTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("active"); ok {
		o.Active = attr.(bool)
	}
	if attr, ok := d.GetOk("default_allow_ip"); ok {
		o.DefaultAllowIP = attr.(bool)
	}
	if attr, ok := d.GetOk("default_allow_non_ip"); ok {
		o.DefaultAllowNonIP = attr.(bool)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("policy_state"); ok {
		o.PolicyState = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("priority_type"); ok {
		o.PriorityType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("auto_generate_priority"); ok {
		o.AutoGeneratePriority = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceFloatingIPACLTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIPACLTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
