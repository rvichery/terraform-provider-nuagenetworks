package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceBGPProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceBGPProfileRead,
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
			"dampening_half_life": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening_max_suppress": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dampening_reuse": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"dampening_suppress": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_export_routing_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_import_routing_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceBGPProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredBGPProfiles := vspk.BGPProfilesList{}
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
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredBGPProfiles, err = parent.BGPProfiles(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredBGPProfiles, err = parent.BGPProfiles(fetchFilter)
		if err != nil {
			return err
		}
	}

	BGPProfile := &vspk.BGPProfile{}

	if len(filteredBGPProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredBGPProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		BGPProfile = filteredBGPProfiles[0]
	}

	d.Set("name", BGPProfile.Name)
	d.Set("dampening_half_life", BGPProfile.DampeningHalfLife)
	d.Set("dampening_max_suppress", BGPProfile.DampeningMaxSuppress)
	d.Set("dampening_name", BGPProfile.DampeningName)
	d.Set("dampening_reuse", BGPProfile.DampeningReuse)
	d.Set("dampening_suppress", BGPProfile.DampeningSuppress)
	d.Set("description", BGPProfile.Description)
	d.Set("entity_scope", BGPProfile.EntityScope)
	d.Set("associated_export_routing_policy_id", BGPProfile.AssociatedExportRoutingPolicyID)
	d.Set("associated_import_routing_policy_id", BGPProfile.AssociatedImportRoutingPolicyID)
	d.Set("external_id", BGPProfile.ExternalID)

	d.Set("id", BGPProfile.Identifier())
	d.Set("parent_id", BGPProfile.ParentID)
	d.Set("parent_type", BGPProfile.ParentType)
	d.Set("owner", BGPProfile.Owner)

	d.SetId(BGPProfile.Identifier())

	return nil
}
