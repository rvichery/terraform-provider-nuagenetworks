package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
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
			"ldapuser_dn": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mobile_number": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"avatar_data": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"avatar_type": {
				Type:     schema.TypeString,
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

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize User object
	o := &vspk.User{
		Password:  d.Get("password").(string),
		LastName:  d.Get("last_name").(string),
		FirstName: d.Get("first_name").(string),
		Email:     d.Get("email").(string),
		UserName:  d.Get("user_name").(string),
	}
	if attr, ok := d.GetOk("ldapuser_dn"); ok {
		o.LDAPUserDN = attr.(string)
	}
	if attr, ok := d.GetOk("management_mode"); ok {
		o.ManagementMode = attr.(string)
	}
	if attr, ok := d.GetOk("disabled"); ok {
		Disabled := attr.(bool)
		o.Disabled = &Disabled
	}
	if attr, ok := d.GetOk("mobile_number"); ok {
		o.MobileNumber = attr.(string)
	}
	if attr, ok := d.GetOk("avatar_data"); ok {
		o.AvatarData = attr.(string)
	}
	if attr, ok := d.GetOk("avatar_type"); ok {
		o.AvatarType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateUser(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceUserRead(d, m)
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.User{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ldapuser_dn", o.LDAPUserDN)
	d.Set("management_mode", o.ManagementMode)
	d.Set("password", o.Password)
	d.Set("last_name", o.LastName)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("first_name", o.FirstName)
	d.Set("disabled", o.Disabled)
	d.Set("email", o.Email)
	d.Set("entity_scope", o.EntityScope)
	d.Set("mobile_number", o.MobileNumber)
	d.Set("user_name", o.UserName)
	d.Set("avatar_data", o.AvatarData)
	d.Set("avatar_type", o.AvatarType)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.User{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Password = d.Get("password").(string)
	o.LastName = d.Get("last_name").(string)
	o.FirstName = d.Get("first_name").(string)
	o.Email = d.Get("email").(string)
	o.UserName = d.Get("user_name").(string)

	if attr, ok := d.GetOk("ldapuser_dn"); ok {
		o.LDAPUserDN = attr.(string)
	}
	if attr, ok := d.GetOk("management_mode"); ok {
		o.ManagementMode = attr.(string)
	}
	if attr, ok := d.GetOk("disabled"); ok {
		Disabled := attr.(bool)
		o.Disabled = &Disabled
	}
	if attr, ok := d.GetOk("mobile_number"); ok {
		o.MobileNumber = attr.(string)
	}
	if attr, ok := d.GetOk("avatar_data"); ok {
		o.AvatarData = attr.(string)
	}
	if attr, ok := d.GetOk("avatar_type"); ok {
		o.AvatarType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.User{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
