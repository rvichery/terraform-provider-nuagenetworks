package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourcePolicyEntry() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePolicyEntryRead,
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
			"match_criteria": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"actions": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_policy_statement": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePolicyEntryRead(d *schema.ResourceData, m interface{}) error {
	filteredPolicyEntries := vspk.PolicyEntriesList{}
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
	parent := &vspk.PolicyStatement{ID: d.Get("parent_policy_statement").(string)}
	filteredPolicyEntries, err = parent.PolicyEntries(fetchFilter)
	if err != nil {
		return err
	}

	PolicyEntry := &vspk.PolicyEntry{}

	if len(filteredPolicyEntries) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPolicyEntries) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	PolicyEntry = filteredPolicyEntries[0]

	d.Set("name", PolicyEntry.Name)
	d.Set("match_criteria", PolicyEntry.MatchCriteria)
	d.Set("actions", PolicyEntry.Actions)
	d.Set("description", PolicyEntry.Description)

	d.Set("id", PolicyEntry.Identifier())
	d.Set("parent_id", PolicyEntry.ParentID)
	d.Set("parent_type", PolicyEntry.ParentType)
	d.Set("owner", PolicyEntry.Owner)

	d.SetId(PolicyEntry.Identifier())

	return nil
}
