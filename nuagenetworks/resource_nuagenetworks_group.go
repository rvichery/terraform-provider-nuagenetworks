package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,
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
			"ldap_group_dn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"account_restrictions": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"restriction_date": {
				Type:     schema.TypeFloat,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"private": {
				Type:     schema.TypeBool,
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
		AccountRestrictions := attr.(bool)
		o.AccountRestrictions = &AccountRestrictions
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
		Private := attr.(bool)
		o.Private = &Private
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
		AccountRestrictions := attr.(bool)
		o.AccountRestrictions = &AccountRestrictions
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
		Private := attr.(bool)
		o.Private = &Private
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
