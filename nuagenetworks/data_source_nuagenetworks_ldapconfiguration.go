package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceLDAPConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLDAPConfigurationRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"accept_all_certificates": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_dn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_name_prefix": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_name_suffix": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_dn_template": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_name_attribute": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authorization_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"authorizing_user_dn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceLDAPConfigurationRead(d *schema.ResourceData, m interface{}) error {
	filteredLDAPConfigurations := vspk.LDAPConfigurationsList{}
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
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	filteredLDAPConfigurations, err = parent.LDAPConfigurations(fetchFilter)
	if err != nil {
		return err
	}

	LDAPConfiguration := &vspk.LDAPConfiguration{}

	if len(filteredLDAPConfigurations) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredLDAPConfigurations) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		LDAPConfiguration = filteredLDAPConfigurations[0]
	}

	d.Set("ssl_enabled", LDAPConfiguration.SSLEnabled)
	d.Set("password", LDAPConfiguration.Password)
	d.Set("last_updated_by", LDAPConfiguration.LastUpdatedBy)
	d.Set("accept_all_certificates", LDAPConfiguration.AcceptAllCertificates)
	d.Set("certificate", LDAPConfiguration.Certificate)
	d.Set("server", LDAPConfiguration.Server)
	d.Set("enabled", LDAPConfiguration.Enabled)
	d.Set("entity_scope", LDAPConfiguration.EntityScope)
	d.Set("port", LDAPConfiguration.Port)
	d.Set("group_dn", LDAPConfiguration.GroupDN)
	d.Set("group_name_prefix", LDAPConfiguration.GroupNamePrefix)
	d.Set("group_name_suffix", LDAPConfiguration.GroupNameSuffix)
	d.Set("user_dn_template", LDAPConfiguration.UserDNTemplate)
	d.Set("user_name_attribute", LDAPConfiguration.UserNameAttribute)
	d.Set("authorization_enabled", LDAPConfiguration.AuthorizationEnabled)
	d.Set("authorizing_user_dn", LDAPConfiguration.AuthorizingUserDN)
	d.Set("external_id", LDAPConfiguration.ExternalID)

	d.Set("id", LDAPConfiguration.Identifier())
	d.Set("parent_id", LDAPConfiguration.ParentID)
	d.Set("parent_type", LDAPConfiguration.ParentType)
	d.Set("owner", LDAPConfiguration.Owner)

	d.SetId(LDAPConfiguration.Identifier())

	return nil
}
