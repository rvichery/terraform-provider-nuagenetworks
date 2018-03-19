package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceInfrastructureAccessProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInfrastructureAccessProfileRead,
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
			"ssh_auth_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_ip_filter": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceInfrastructureAccessProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredInfrastructureAccessProfiles := vspk.InfrastructureAccessProfilesList{}
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
	parent := m.(*vspk.Me)
	filteredInfrastructureAccessProfiles, err = parent.InfrastructureAccessProfiles(fetchFilter)
	if err != nil {
		return err
	}

	InfrastructureAccessProfile := &vspk.InfrastructureAccessProfile{}

	if len(filteredInfrastructureAccessProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredInfrastructureAccessProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	InfrastructureAccessProfile = filteredInfrastructureAccessProfiles[0]

	d.Set("ssh_auth_mode", InfrastructureAccessProfile.SSHAuthMode)
	d.Set("name", InfrastructureAccessProfile.Name)
	d.Set("password", InfrastructureAccessProfile.Password)
	d.Set("last_updated_by", InfrastructureAccessProfile.LastUpdatedBy)
	d.Set("description", InfrastructureAccessProfile.Description)
	d.Set("enterprise_id", InfrastructureAccessProfile.EnterpriseID)
	d.Set("entity_scope", InfrastructureAccessProfile.EntityScope)
	d.Set("source_ip_filter", InfrastructureAccessProfile.SourceIPFilter)
	d.Set("user_name", InfrastructureAccessProfile.UserName)
	d.Set("external_id", InfrastructureAccessProfile.ExternalID)

	d.Set("id", InfrastructureAccessProfile.Identifier())
	d.Set("parent_id", InfrastructureAccessProfile.ParentID)
	d.Set("parent_type", InfrastructureAccessProfile.ParentType)
	d.Set("owner", InfrastructureAccessProfile.Owner)

	d.SetId(InfrastructureAccessProfile.Identifier())

	return nil
}
