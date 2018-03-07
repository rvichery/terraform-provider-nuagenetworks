package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceDomainTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainTemplateCreate,
		Read:   resourceDomainTemplateRead,
		Update: resourceDomainTemplateUpdate,
		Delete: resourceDomainTemplateDelete,
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
			"dpi": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DISABLED",
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
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_change_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_bgp_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_multicast_channel_map_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_pat_mapper_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"multicast": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDomainTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize DomainTemplate object
	o := &vspk.DomainTemplate{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("dpi"); ok {
		o.DPI = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("encryption"); ok {
		o.Encryption = attr.(string)
	}
	if attr, ok := d.GetOk("policy_change_status"); ok {
		o.PolicyChangeStatus = attr.(string)
	}
	if attr, ok := d.GetOk("associated_bgp_profile_id"); ok {
		o.AssociatedBGPProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
		o.AssociatedMulticastChannelMapID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_pat_mapper_id"); ok {
		o.AssociatedPATMapperID = attr.(string)
	}
	if attr, ok := d.GetOk("multicast"); ok {
		o.Multicast = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateDomainTemplate(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("domains"); ok {
		o.AssignDomains(attr.(vspk.DomainsList))
	}
	return resourceDomainTemplateRead(d, m)
}

func resourceDomainTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DomainTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("dpi", o.DPI)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("encryption", o.Encryption)
	d.Set("entity_scope", o.EntityScope)
	d.Set("policy_change_status", o.PolicyChangeStatus)
	d.Set("associated_bgp_profile_id", o.AssociatedBGPProfileID)
	d.Set("associated_multicast_channel_map_id", o.AssociatedMulticastChannelMapID)
	d.Set("associated_pat_mapper_id", o.AssociatedPATMapperID)
	d.Set("multicast", o.Multicast)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceDomainTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DomainTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("dpi"); ok {
		o.DPI = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("encryption"); ok {
		o.Encryption = attr.(string)
	}
	if attr, ok := d.GetOk("policy_change_status"); ok {
		o.PolicyChangeStatus = attr.(string)
	}
	if attr, ok := d.GetOk("associated_bgp_profile_id"); ok {
		o.AssociatedBGPProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_multicast_channel_map_id"); ok {
		o.AssociatedMulticastChannelMapID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_pat_mapper_id"); ok {
		o.AssociatedPATMapperID = attr.(string)
	}
	if attr, ok := d.GetOk("multicast"); ok {
		o.Multicast = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceDomainTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DomainTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
