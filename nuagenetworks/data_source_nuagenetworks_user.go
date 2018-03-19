package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ldapuser_dn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mobile_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avatar_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avatar_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_group"},
			},
		},
	}
}

func dataSourceUserRead(d *schema.ResourceData, m interface{}) error {
	filteredUsers := vspk.UsersList{}
	err := &bambou.Error{}
	fetchFilter := &bambou.FetchingInfo{}

	filters, filtersOk := d.GetOk("filter")
	if filtersOk {
		fetchFilter = bambou.NewFetchingInfo()
		for _, v := range filters.(*schema.Set).List() {
			m := v.(map[string]interface{})
			if fetchFilter.Filter != "" {
				fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string), m["operator"].(string), m["value"].(string))
			} else {
				fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
			}

		}
	}
	if attr, ok := d.GetOk("parent_group"); ok {
		parent := &vspk.Group{ID: attr.(string)}
		filteredUsers, err = parent.Users(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredUsers, err = parent.Users(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredUsers, err = parent.Users(fetchFilter)
		if err != nil {
			return err
		}
	}

	User := &vspk.User{}

	if len(filteredUsers) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredUsers) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	User = filteredUsers[0]

	d.Set("ldapuser_dn", User.LDAPUserDN)
	d.Set("management_mode", User.ManagementMode)
	d.Set("password", User.Password)
	d.Set("last_name", User.LastName)
	d.Set("last_updated_by", User.LastUpdatedBy)
	d.Set("first_name", User.FirstName)
	d.Set("disabled", User.Disabled)
	d.Set("email", User.Email)
	d.Set("entity_scope", User.EntityScope)
	d.Set("mobile_number", User.MobileNumber)
	d.Set("user_name", User.UserName)
	d.Set("avatar_data", User.AvatarData)
	d.Set("avatar_type", User.AvatarType)
	d.Set("external_id", User.ExternalID)

	d.Set("id", User.Identifier())
	d.Set("parent_id", User.ParentID)
	d.Set("parent_type", User.ParentType)
	d.Set("owner", User.Owner)

	d.SetId(User.Identifier())

	return nil
}
