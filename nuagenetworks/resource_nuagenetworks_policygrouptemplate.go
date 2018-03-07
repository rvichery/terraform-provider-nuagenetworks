package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourcePolicyGroupTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourcePolicyGroupTemplateCreate,
		Read:   resourcePolicyGroupTemplateRead,
		Update: resourcePolicyGroupTemplateUpdate,
		Delete: resourcePolicyGroupTemplateDelete,
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
			"evpn_community_tag": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"external": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_l2_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain_template"},
			},
			"parent_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l2_domain_template"},
			},
		},
	}
}

func resourcePolicyGroupTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize PolicyGroupTemplate object
	o := &vspk.PolicyGroupTemplate{
		Name: d.Get("name").(string),
		Type: d.Get("type").(string),
	}
	if attr, ok := d.GetOk("evpn_community_tag"); ok {
		o.EVPNCommunityTag = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external"); ok {
		o.External = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		err := parent.CreatePolicyGroupTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		err := parent.CreatePolicyGroupTemplate(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourcePolicyGroupTemplateRead(d, m)
}

func resourcePolicyGroupTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyGroupTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("evpn_community_tag", o.EVPNCommunityTag)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external", o.External)
	d.Set("external_id", o.ExternalID)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourcePolicyGroupTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyGroupTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Type = d.Get("type").(string)

	if attr, ok := d.GetOk("evpn_community_tag"); ok {
		o.EVPNCommunityTag = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external"); ok {
		o.External = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourcePolicyGroupTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PolicyGroupTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
