package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceNetworkLayout() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceNetworkLayoutRead,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "service_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "route_reflector_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "autonomous_system_num": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceNetworkLayoutRead(d *schema.ResourceData, m interface{}) error {
    filteredNetworkLayouts := vspk.NetworkLayoutsList{}
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
    parent := m.(*vspk.Me)
    filteredNetworkLayouts, err = parent.NetworkLayouts(fetchFilter)
    if err != nil {
        return err
    }

    NetworkLayout := &vspk.NetworkLayout{}

    if len(filteredNetworkLayouts) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredNetworkLayouts) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    NetworkLayout = filteredNetworkLayouts[0]

    d.Set("last_updated_by", NetworkLayout.LastUpdatedBy)
    d.Set("service_type", NetworkLayout.ServiceType)
    d.Set("entity_scope", NetworkLayout.EntityScope)
    d.Set("route_reflector_ip", NetworkLayout.RouteReflectorIP)
    d.Set("autonomous_system_num", NetworkLayout.AutonomousSystemNum)
    d.Set("external_id", NetworkLayout.ExternalID)
    
    d.Set("id", NetworkLayout.Identifier())
    d.Set("parent_id", NetworkLayout.ParentID)
    d.Set("parent_type", NetworkLayout.ParentType)
    d.Set("owner", NetworkLayout.Owner)

    d.SetId(NetworkLayout.Identifier())
    
    return nil
}