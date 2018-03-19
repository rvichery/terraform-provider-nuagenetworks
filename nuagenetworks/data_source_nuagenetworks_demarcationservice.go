package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceDemarcationService() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceDemarcationServiceRead,
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
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_vlanid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_link": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceDemarcationServiceRead(d *schema.ResourceData, m interface{}) error {
    filteredDemarcationServices := vspk.DemarcationServicesList{}
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
    parent := &vspk.Link{ID: d.Get("parent_link").(string)}
    filteredDemarcationServices, err = parent.DemarcationServices(fetchFilter)
    if err != nil {
        return err
    }

    DemarcationService := &vspk.DemarcationService{}

    if len(filteredDemarcationServices) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredDemarcationServices) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    DemarcationService = filteredDemarcationServices[0]

    d.Set("route_distinguisher", DemarcationService.RouteDistinguisher)
    d.Set("priority", DemarcationService.Priority)
    d.Set("associated_gateway_id", DemarcationService.AssociatedGatewayID)
    d.Set("associated_vlanid", DemarcationService.AssociatedVLANID)
    d.Set("type", DemarcationService.Type)
    
    d.Set("id", DemarcationService.Identifier())
    d.Set("parent_id", DemarcationService.ParentID)
    d.Set("parent_type", DemarcationService.ParentType)
    d.Set("owner", DemarcationService.Owner)

    d.SetId(DemarcationService.Identifier())
    
    return nil
}