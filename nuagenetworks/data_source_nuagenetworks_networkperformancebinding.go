package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNetworkPerformanceBinding() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNetworkPerformanceBindingRead,
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
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "associated_network_measurement_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain", "parent_network_performance_measurement"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_network_performance_measurement"},
            },
            "parent_network_performance_measurement": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain"},
            },
        },
    }
}


func dataSourceNetworkPerformanceBindingRead(d *schema.ResourceData, m interface{}) error {
    filteredNetworkPerformanceBindings := vspk.NetworkPerformanceBindingsList{}
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
        filteredNetworkPerformanceBindings, err = parent.NetworkPerformanceBindings(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredNetworkPerformanceBindings, err = parent.NetworkPerformanceBindings(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_network_performance_measurement"); ok {
        parent := &vspk.NetworkPerformanceMeasurement{ID: attr.(string)}
        filteredNetworkPerformanceBindings, err = parent.NetworkPerformanceBindings(fetchFilter)
        if err != nil {
            return err
        }
    }

    NetworkPerformanceBinding := &vspk.NetworkPerformanceBinding{}

    if len(filteredNetworkPerformanceBindings) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNetworkPerformanceBindings) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NetworkPerformanceBinding = filteredNetworkPerformanceBindings[0]

    d.Set("read_only", NetworkPerformanceBinding.ReadOnly)
    d.Set("priority", NetworkPerformanceBinding.Priority)
    d.Set("associated_network_measurement_id", NetworkPerformanceBinding.AssociatedNetworkMeasurementID)
    
    d.Set("id", NetworkPerformanceBinding.Identifier())
    d.Set("parent_id", NetworkPerformanceBinding.ParentID)
    d.Set("parent_type", NetworkPerformanceBinding.ParentType)
    d.Set("owner", NetworkPerformanceBinding.Owner)

    d.SetId(NetworkPerformanceBinding.Identifier())
    
    return nil
}