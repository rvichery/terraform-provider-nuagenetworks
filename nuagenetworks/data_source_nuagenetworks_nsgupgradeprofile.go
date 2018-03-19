package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceNSGUpgradeProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceNSGUpgradeProfileRead,
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
			"name": {
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
			"metadata_upgrade_path": {
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
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceNSGUpgradeProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredNSGUpgradeProfiles := vspk.NSGUpgradeProfilesList{}
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
	filteredNSGUpgradeProfiles, err = parent.NSGUpgradeProfiles(fetchFilter)
	if err != nil {
		return err
	}

	NSGUpgradeProfile := &vspk.NSGUpgradeProfile{}

	if len(filteredNSGUpgradeProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredNSGUpgradeProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	NSGUpgradeProfile = filteredNSGUpgradeProfiles[0]

	d.Set("name", NSGUpgradeProfile.Name)
	d.Set("last_updated_by", NSGUpgradeProfile.LastUpdatedBy)
	d.Set("description", NSGUpgradeProfile.Description)
	d.Set("metadata_upgrade_path", NSGUpgradeProfile.MetadataUpgradePath)
	d.Set("enterprise_id", NSGUpgradeProfile.EnterpriseID)
	d.Set("entity_scope", NSGUpgradeProfile.EntityScope)
	d.Set("external_id", NSGUpgradeProfile.ExternalID)

	d.Set("id", NSGUpgradeProfile.Identifier())
	d.Set("parent_id", NSGUpgradeProfile.ParentID)
	d.Set("parent_type", NSGUpgradeProfile.ParentType)
	d.Set("owner", NSGUpgradeProfile.Owner)

	d.SetId(NSGUpgradeProfile.Identifier())

	return nil
}
