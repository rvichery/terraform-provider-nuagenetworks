package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDSCPRemarkingPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDSCPRemarkingPolicyRead,
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
			"dscp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"forwarding_class": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_dscp_remarking_policy_table": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceDSCPRemarkingPolicyRead(d *schema.ResourceData, m interface{}) error {
	filteredDSCPRemarkingPolicies := vspk.DSCPRemarkingPoliciesList{}
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
	parent := &vspk.DSCPRemarkingPolicyTable{ID: d.Get("parent_dscp_remarking_policy_table").(string)}
	filteredDSCPRemarkingPolicies, err = parent.DSCPRemarkingPolicies(fetchFilter)
	if err != nil {
		return err
	}

	DSCPRemarkingPolicy := &vspk.DSCPRemarkingPolicy{}

	if len(filteredDSCPRemarkingPolicies) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDSCPRemarkingPolicies) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	DSCPRemarkingPolicy = filteredDSCPRemarkingPolicies[0]

	d.Set("dscp", DSCPRemarkingPolicy.DSCP)
	d.Set("last_updated_by", DSCPRemarkingPolicy.LastUpdatedBy)
	d.Set("entity_scope", DSCPRemarkingPolicy.EntityScope)
	d.Set("forwarding_class", DSCPRemarkingPolicy.ForwardingClass)
	d.Set("external_id", DSCPRemarkingPolicy.ExternalID)

	d.Set("id", DSCPRemarkingPolicy.Identifier())
	d.Set("parent_id", DSCPRemarkingPolicy.ParentID)
	d.Set("parent_type", DSCPRemarkingPolicy.ParentType)
	d.Set("owner", DSCPRemarkingPolicy.Owner)

	d.SetId(DSCPRemarkingPolicy.Identifier())

	return nil
}
