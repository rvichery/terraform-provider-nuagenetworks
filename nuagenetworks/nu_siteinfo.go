package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceSiteInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceSiteInfoCreate,
		Read:   resourceSiteInfoRead,
		Update: resourceSiteInfoUpdate,
		Delete: resourceSiteInfoDelete,

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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_identifier": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"xmpp_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceSiteInfoCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize SiteInfo object
	o := &vspk.SiteInfo{
		Name:       d.Get("name").(string),
		Address:    d.Get("address").(string),
		XmppDomain: d.Get("xmpp_domain").(string),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("site_identifier"); ok {
		o.SiteIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateSiteInfo(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceSiteInfoRead(d, m)
}

func resourceSiteInfoRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SiteInfo{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("address", o.Address)
	d.Set("description", o.Description)
	d.Set("site_identifier", o.SiteIdentifier)
	d.Set("xmpp_domain", o.XmppDomain)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceSiteInfoUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SiteInfo{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Address = d.Get("address").(string)
	o.XmppDomain = d.Get("xmpp_domain").(string)

	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("site_identifier"); ok {
		o.SiteIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceSiteInfoDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.SiteInfo{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
