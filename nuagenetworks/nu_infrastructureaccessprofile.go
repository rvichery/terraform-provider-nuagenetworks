package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceInfrastructureAccessProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceInfrastructureAccessProfileCreate,
		Read:   resourceInfrastructureAccessProfileRead,
		Update: resourceInfrastructureAccessProfileUpdate,
		Delete: resourceInfrastructureAccessProfileDelete,

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
			"ssh_auth_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PASSWORD_AND_KEY_BASED",
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
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
			"enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_ip_filter": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DISABLED",
			},
			"user_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "nuage",
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceInfrastructureAccessProfileCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize InfrastructureAccessProfile object
	o := &vspk.InfrastructureAccessProfile{
		Name:     d.Get("name").(string),
		Password: d.Get("password").(string),
	}
	if attr, ok := d.GetOk("ssh_auth_mode"); ok {
		o.SSHAuthMode = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("source_ip_filter"); ok {
		o.SourceIPFilter = attr.(string)
	}
	if attr, ok := d.GetOk("user_name"); ok {
		o.UserName = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := m.(*vspk.Me)
	err := parent.CreateInfrastructureAccessProfile(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceInfrastructureAccessProfileRead(d, m)
}

func resourceInfrastructureAccessProfileRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureAccessProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ssh_auth_mode", o.SSHAuthMode)
	d.Set("name", o.Name)
	d.Set("password", o.Password)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("description", o.Description)
	d.Set("enterprise_id", o.EnterpriseID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("source_ip_filter", o.SourceIPFilter)
	d.Set("user_name", o.UserName)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceInfrastructureAccessProfileUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureAccessProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.Password = d.Get("password").(string)

	if attr, ok := d.GetOk("ssh_auth_mode"); ok {
		o.SSHAuthMode = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("source_ip_filter"); ok {
		o.SourceIPFilter = attr.(string)
	}
	if attr, ok := d.GetOk("user_name"); ok {
		o.UserName = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceInfrastructureAccessProfileDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.InfrastructureAccessProfile{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
