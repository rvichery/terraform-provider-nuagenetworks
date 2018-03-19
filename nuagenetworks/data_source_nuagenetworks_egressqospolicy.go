package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceEgressQOSPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEgressQOSPolicyRead,
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
			"parent_queue_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_service_class": {
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
			"assoc_egress_qos_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_cos_remarking_policy_table_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_dscp_remarking_policy_table_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"queue1_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"queue1_forwarding_classes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"queue2_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"queue2_forwarding_classes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"queue3_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"queue3_forwarding_classes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"queue4_associated_rate_limiter_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"queue4_forwarding_classes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceEgressQOSPolicyRead(d *schema.ResourceData, m interface{}) error {
	filteredEgressQOSPolicies := vspk.EgressQOSPoliciesList{}
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
		filteredEgressQOSPolicies, err = parent.EgressQOSPolicies(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredEgressQOSPolicies, err = parent.EgressQOSPolicies(fetchFilter)
		if err != nil {
			return err
		}
	}

	EgressQOSPolicy := &vspk.EgressQOSPolicy{}

	if len(filteredEgressQOSPolicies) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredEgressQOSPolicies) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	EgressQOSPolicy = filteredEgressQOSPolicies[0]

	d.Set("name", EgressQOSPolicy.Name)
	d.Set("parent_queue_associated_rate_limiter_id", EgressQOSPolicy.ParentQueueAssociatedRateLimiterID)
	d.Set("last_updated_by", EgressQOSPolicy.LastUpdatedBy)
	d.Set("default_service_class", EgressQOSPolicy.DefaultServiceClass)
	d.Set("description", EgressQOSPolicy.Description)
	d.Set("entity_scope", EgressQOSPolicy.EntityScope)
	d.Set("assoc_egress_qos_id", EgressQOSPolicy.AssocEgressQosId)
	d.Set("associated_cos_remarking_policy_table_id", EgressQOSPolicy.AssociatedCOSRemarkingPolicyTableID)
	d.Set("associated_dscp_remarking_policy_table_id", EgressQOSPolicy.AssociatedDSCPRemarkingPolicyTableID)
	d.Set("queue1_associated_rate_limiter_id", EgressQOSPolicy.Queue1AssociatedRateLimiterID)
	d.Set("queue1_forwarding_classes", EgressQOSPolicy.Queue1ForwardingClasses)
	d.Set("queue2_associated_rate_limiter_id", EgressQOSPolicy.Queue2AssociatedRateLimiterID)
	d.Set("queue2_forwarding_classes", EgressQOSPolicy.Queue2ForwardingClasses)
	d.Set("queue3_associated_rate_limiter_id", EgressQOSPolicy.Queue3AssociatedRateLimiterID)
	d.Set("queue3_forwarding_classes", EgressQOSPolicy.Queue3ForwardingClasses)
	d.Set("queue4_associated_rate_limiter_id", EgressQOSPolicy.Queue4AssociatedRateLimiterID)
	d.Set("queue4_forwarding_classes", EgressQOSPolicy.Queue4ForwardingClasses)
	d.Set("external_id", EgressQOSPolicy.ExternalID)

	d.Set("id", EgressQOSPolicy.Identifier())
	d.Set("parent_id", EgressQOSPolicy.ParentID)
	d.Set("parent_type", EgressQOSPolicy.ParentType)
	d.Set("owner", EgressQOSPolicy.Owner)

	d.SetId(EgressQOSPolicy.Identifier())

	return nil
}
