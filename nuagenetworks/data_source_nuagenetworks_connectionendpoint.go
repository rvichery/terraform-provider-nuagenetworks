package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceConnectionendpoint() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceConnectionendpointRead,
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
            "ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "end_point_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_infrastructure_access_profile": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceConnectionendpointRead(d *schema.ResourceData, m interface{}) error {
    filteredConnectionendpoints := vspk.ConnectionendpointsList{}
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
    parent := &vspk.InfrastructureAccessProfile{ID: d.Get("parent_infrastructure_access_profile").(string)}
    filteredConnectionendpoints, err = parent.Connectionendpoints(fetchFilter)
    if err != nil {
        return err
    }

    Connectionendpoint := &vspk.Connectionendpoint{}

    if len(filteredConnectionendpoints) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredConnectionendpoints) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Connectionendpoint = filteredConnectionendpoints[0]

    d.Set("ip_address", Connectionendpoint.IPAddress)
    d.Set("ip_type", Connectionendpoint.IPType)
    d.Set("name", Connectionendpoint.Name)
    d.Set("description", Connectionendpoint.Description)
    d.Set("end_point_type", Connectionendpoint.EndPointType)
    
    d.Set("id", Connectionendpoint.Identifier())
    d.Set("parent_id", Connectionendpoint.ParentID)
    d.Set("parent_type", Connectionendpoint.ParentType)
    d.Set("owner", Connectionendpoint.Owner)

    d.SetId(Connectionendpoint.Identifier())
    
    return nil
}