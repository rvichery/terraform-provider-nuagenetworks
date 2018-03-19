package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePSNATPool() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePSNATPoolRead,
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
            "end_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "start_address": &schema.Schema{
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


func dataSourcePSNATPoolRead(d *schema.ResourceData, m interface{}) error {
    filteredPSNATPools := vspk.PSNATPoolsList{}
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
    filteredPSNATPools, err = parent.PSNATPools(fetchFilter)
    if err != nil {
        return err
    }

    PSNATPool := &vspk.PSNATPool{}

    if len(filteredPSNATPools) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPSNATPools) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    PSNATPool = filteredPSNATPools[0]

    d.Set("end_address", PSNATPool.EndAddress)
    d.Set("start_address", PSNATPool.StartAddress)
    
    d.Set("id", PSNATPool.Identifier())
    d.Set("parent_id", PSNATPool.ParentID)
    d.Set("parent_type", PSNATPool.ParentType)
    d.Set("owner", PSNATPool.Owner)

    d.SetId(PSNATPool.Identifier())
    
    return nil
}