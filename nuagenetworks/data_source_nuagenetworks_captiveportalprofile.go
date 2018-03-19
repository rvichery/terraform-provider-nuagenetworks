package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceCaptivePortalProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCaptivePortalProfileRead,
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
			"captive_page": {
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
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portal_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ssid_connection": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ssid_connection"},
			},
		},
	}
}

func dataSourceCaptivePortalProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredCaptivePortalProfiles := vspk.CaptivePortalProfilesList{}
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
	if attr, ok := d.GetOk("parent_ssid_connection"); ok {
		parent := &vspk.SSIDConnection{ID: attr.(string)}
		filteredCaptivePortalProfiles, err = parent.CaptivePortalProfiles(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredCaptivePortalProfiles, err = parent.CaptivePortalProfiles(fetchFilter)
		if err != nil {
			return err
		}
	}

	CaptivePortalProfile := &vspk.CaptivePortalProfile{}

	if len(filteredCaptivePortalProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredCaptivePortalProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	CaptivePortalProfile = filteredCaptivePortalProfiles[0]

	d.Set("name", CaptivePortalProfile.Name)
	d.Set("captive_page", CaptivePortalProfile.CaptivePage)
	d.Set("last_updated_by", CaptivePortalProfile.LastUpdatedBy)
	d.Set("description", CaptivePortalProfile.Description)
	d.Set("entity_scope", CaptivePortalProfile.EntityScope)
	d.Set("portal_type", CaptivePortalProfile.PortalType)
	d.Set("external_id", CaptivePortalProfile.ExternalID)

	d.Set("id", CaptivePortalProfile.Identifier())
	d.Set("parent_id", CaptivePortalProfile.ParentID)
	d.Set("parent_type", CaptivePortalProfile.ParentType)
	d.Set("owner", CaptivePortalProfile.Owner)

	d.SetId(CaptivePortalProfile.Identifier())

	return nil
}
