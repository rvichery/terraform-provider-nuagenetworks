package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,

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
			"ldap_group_dn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"management_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"account_restrictions": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"restriction_date": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private": &schema.Schema{
				Type:     schema.TypeBool,
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

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize Group object
	o := &vspk.Group{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("ldap_group_dn"); ok {
		o.LDAPGroupDN = attr.(string)
	}
	if attr, ok := d.GetOk("management_mode"); ok {
		o.ManagementMode = attr.(string)
	}
	if attr, ok := d.GetOk("account_restrictions"); ok {
		o.AccountRestrictions = attr.(bool)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("restriction_date"); ok {
		o.RestrictionDate = attr.(float64)
	}
	if attr, ok := d.GetOk("role"); ok {
		o.Role = attr.(string)
	}
	if attr, ok := d.GetOk("private"); ok {
		o.Private = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateGroup(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	if attr, ok := d.GetOk("users"); ok {
		o.AssignUsers(attr.(vspk.UsersList))
	}
	return resourceGroupRead(d, m)
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Group{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ldap_group_dn", o.LDAPGroupDN)
	d.Set("name", o.Name)
	d.Set("management_mode", o.ManagementMode)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("account_restrictions", o.AccountRestrictions)
	d.Set("description", o.Description)
	d.Set("restriction_date", o.RestrictionDate)
	d.Set("entity_scope", o.EntityScope)
	d.Set("role", o.Role)
	d.Set("private", o.Private)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Group{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("ldap_group_dn"); ok {
		o.LDAPGroupDN = attr.(string)
	}
	if attr, ok := d.GetOk("management_mode"); ok {
		o.ManagementMode = attr.(string)
	}
	if attr, ok := d.GetOk("account_restrictions"); ok {
		o.AccountRestrictions = attr.(bool)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("restriction_date"); ok {
		o.RestrictionDate = attr.(float64)
	}
	if attr, ok := d.GetOk("role"); ok {
		o.Role = attr.(string)
	}
	if attr, ok := d.GetOk("private"); ok {
		o.Private = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.Group{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
