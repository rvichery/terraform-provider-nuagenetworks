package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceIngressQOSPolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceIngressQOSPolicyRead,
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
            "parent_queue_associated_rate_limiter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
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
            "assoc_egress_qos_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "queue1_associated_rate_limiter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "queue1_forwarding_classes": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "queue2_associated_rate_limiter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "queue2_forwarding_classes": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "queue3_associated_rate_limiter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "queue3_forwarding_classes": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "queue4_associated_rate_limiter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "queue4_forwarding_classes": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
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


func dataSourceIngressQOSPolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredIngressQOSPolicies := vspk.IngressQOSPoliciesList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredIngressQOSPolicies, err = parent.IngressQOSPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredIngressQOSPolicies, err = parent.IngressQOSPolicies(fetchFilter)
        if err != nil {
            return err
        }
    }

    IngressQOSPolicy := &vspk.IngressQOSPolicy{}

    if len(filteredIngressQOSPolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredIngressQOSPolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    IngressQOSPolicy = filteredIngressQOSPolicies[0]

    d.Set("name", IngressQOSPolicy.Name)
    d.Set("parent_queue_associated_rate_limiter_id", IngressQOSPolicy.ParentQueueAssociatedRateLimiterID)
    d.Set("last_updated_by", IngressQOSPolicy.LastUpdatedBy)
    d.Set("description", IngressQOSPolicy.Description)
    d.Set("entity_scope", IngressQOSPolicy.EntityScope)
    d.Set("assoc_egress_qos_id", IngressQOSPolicy.AssocEgressQosId)
    d.Set("queue1_associated_rate_limiter_id", IngressQOSPolicy.Queue1AssociatedRateLimiterID)
    d.Set("queue1_forwarding_classes", IngressQOSPolicy.Queue1ForwardingClasses)
    d.Set("queue2_associated_rate_limiter_id", IngressQOSPolicy.Queue2AssociatedRateLimiterID)
    d.Set("queue2_forwarding_classes", IngressQOSPolicy.Queue2ForwardingClasses)
    d.Set("queue3_associated_rate_limiter_id", IngressQOSPolicy.Queue3AssociatedRateLimiterID)
    d.Set("queue3_forwarding_classes", IngressQOSPolicy.Queue3ForwardingClasses)
    d.Set("queue4_associated_rate_limiter_id", IngressQOSPolicy.Queue4AssociatedRateLimiterID)
    d.Set("queue4_forwarding_classes", IngressQOSPolicy.Queue4ForwardingClasses)
    d.Set("external_id", IngressQOSPolicy.ExternalID)
    
    d.Set("id", IngressQOSPolicy.Identifier())
    d.Set("parent_id", IngressQOSPolicy.ParentID)
    d.Set("parent_type", IngressQOSPolicy.ParentType)
    d.Set("owner", IngressQOSPolicy.Owner)

    d.SetId(IngressQOSPolicy.Identifier())
    
    return nil
}