package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceCaptivePortalProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceCaptivePortalProfileCreate,
		Read:   resourceCaptivePortalProfileRead,
		Update: resourceCaptivePortalProfileUpdate,
		Delete: resourceCaptivePortalProfileDelete,

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
				Required: true,
			},
			"captive_page": &schema.Schema{
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
			"portal_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

func resourceCaptivePortalProfileCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize CaptivePortalProfile object
	o := &vspk.CaptivePortalProfile{
		Name:        d.Get("name").(string),
		CaptivePage: d.Get("captive_page").(string),
		PortalType:  d.Get("portal_type").(string),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateCaptivePortalProfile(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceCaptivePortalProfileRead(d, m)
}

func resourceCaptivePortalProfileRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CaptivePortalProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("captive_page", o.CaptivePage)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("portal_type", o.PortalType)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceCaptivePortalProfileUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CaptivePortalProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.CaptivePage = d.Get("captive_page").(string)
	o.PortalType = d.Get("portal_type").(string)

	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceCaptivePortalProfileDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.CaptivePortalProfile{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
