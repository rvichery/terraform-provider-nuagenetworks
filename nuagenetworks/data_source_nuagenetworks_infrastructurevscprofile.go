package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceInfrastructureVscProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInfrastructureVscProfileRead,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"second_controller": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_controller": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"probe_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceInfrastructureVscProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredInfrastructureVscProfiles := vspk.InfrastructureVscProfilesList{}
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
	filteredInfrastructureVscProfiles, err = parent.InfrastructureVscProfiles(fetchFilter)
	if err != nil {
		return err
	}

	InfrastructureVscProfile := &vspk.InfrastructureVscProfile{}

	if len(filteredInfrastructureVscProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredInfrastructureVscProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		InfrastructureVscProfile = filteredInfrastructureVscProfiles[0]
	}

	d.Set("name", InfrastructureVscProfile.Name)
	d.Set("last_updated_by", InfrastructureVscProfile.LastUpdatedBy)
	d.Set("second_controller", InfrastructureVscProfile.SecondController)
	d.Set("description", InfrastructureVscProfile.Description)
	d.Set("first_controller", InfrastructureVscProfile.FirstController)
	d.Set("enterprise_id", InfrastructureVscProfile.EnterpriseID)
	d.Set("entity_scope", InfrastructureVscProfile.EntityScope)
	d.Set("probe_interval", InfrastructureVscProfile.ProbeInterval)
	d.Set("external_id", InfrastructureVscProfile.ExternalID)

	d.Set("id", InfrastructureVscProfile.Identifier())
	d.Set("parent_id", InfrastructureVscProfile.ParentID)
	d.Set("parent_type", InfrastructureVscProfile.ParentType)
	d.Set("owner", InfrastructureVscProfile.Owner)

	d.SetId(InfrastructureVscProfile.Identifier())

	return nil
}
