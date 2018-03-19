package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceRoutingPolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceRoutingPolicyRead,
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
            "default_action": &schema.Schema{
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
            "policy_definition": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain"},
            },
        },
    }
}


func dataSourceRoutingPolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredRoutingPolicies := vspk.RoutingPoliciesList{}
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
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredRoutingPolicies, err = parent.RoutingPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredRoutingPolicies, err = parent.RoutingPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredRoutingPolicies, err = parent.RoutingPolicies(fetchFilter)
        if err != nil {
            return err
        }
    }

    RoutingPolicy := &vspk.RoutingPolicy{}

    if len(filteredRoutingPolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredRoutingPolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    RoutingPolicy = filteredRoutingPolicies[0]

    d.Set("name", RoutingPolicy.Name)
    d.Set("default_action", RoutingPolicy.DefaultAction)
    d.Set("description", RoutingPolicy.Description)
    d.Set("entity_scope", RoutingPolicy.EntityScope)
    d.Set("policy_definition", RoutingPolicy.PolicyDefinition)
    d.Set("external_id", RoutingPolicy.ExternalID)
    
    d.Set("id", RoutingPolicy.Identifier())
    d.Set("parent_id", RoutingPolicy.ParentID)
    d.Set("parent_type", RoutingPolicy.ParentType)
    d.Set("owner", RoutingPolicy.Owner)

    d.SetId(RoutingPolicy.Identifier())
    
    return nil
}