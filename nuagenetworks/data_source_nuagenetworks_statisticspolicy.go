package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceStatisticsPolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceStatisticsPolicyRead,
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
            "data_collection_frequency": &schema.Schema{
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
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_vport", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_zone", "parent_vport", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_zone", "parent_domain", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_address_map", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
            },
            "parent_address_map": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_address_map", "parent_ns_port", "parent_patnat_pool"},
            },
            "parent_ns_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_patnat_pool"},
            },
            "parent_patnat_pool": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_ns_port"},
            },
        },
    }
}


func dataSourceStatisticsPolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredStatisticsPolicies := vspk.StatisticsPoliciesList{}
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
    if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        filteredStatisticsPolicies, err = parent.StatisticsPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredStatisticsPolicies, err = parent.StatisticsPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredStatisticsPolicies, err = parent.StatisticsPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredStatisticsPolicies, err = parent.StatisticsPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_address_map"); ok {
        parent := &vspk.AddressMap{ID: attr.(string)}
        filteredStatisticsPolicies, err = parent.StatisticsPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredStatisticsPolicies, err = parent.StatisticsPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_port"); ok {
        parent := &vspk.NSPort{ID: attr.(string)}
        filteredStatisticsPolicies, err = parent.StatisticsPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_patnat_pool"); ok {
        parent := &vspk.PATNATPool{ID: attr.(string)}
        filteredStatisticsPolicies, err = parent.StatisticsPolicies(fetchFilter)
        if err != nil {
            return err
        }
    }

    StatisticsPolicy := &vspk.StatisticsPolicy{}

    if len(filteredStatisticsPolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredStatisticsPolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    StatisticsPolicy = filteredStatisticsPolicies[0]

    d.Set("name", StatisticsPolicy.Name)
    d.Set("last_updated_by", StatisticsPolicy.LastUpdatedBy)
    d.Set("data_collection_frequency", StatisticsPolicy.DataCollectionFrequency)
    d.Set("description", StatisticsPolicy.Description)
    d.Set("entity_scope", StatisticsPolicy.EntityScope)
    d.Set("external_id", StatisticsPolicy.ExternalID)
    
    d.Set("id", StatisticsPolicy.Identifier())
    d.Set("parent_id", StatisticsPolicy.ParentID)
    d.Set("parent_type", StatisticsPolicy.ParentType)
    d.Set("owner", StatisticsPolicy.Owner)

    d.SetId(StatisticsPolicy.Identifier())
    
    return nil
}