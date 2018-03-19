package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceLtestatistics() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceLtestatisticsRead,
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
            "version": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "end_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "start_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "stats_data": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceLtestatisticsRead(d *schema.ResourceData, m interface{}) error {
    filteredLtestatistics := vspk.LtestatisticsList{}
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
    parent := &vspk.VLAN{ID: d.Get("parent_vlan").(string)}
    filteredLtestatistics, err = parent.Ltestatistics(fetchFilter)
    if err != nil {
        return err
    }

    Ltestatistics := &vspk.Ltestatistics{}

    if len(filteredLtestatistics) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredLtestatistics) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Ltestatistics = filteredLtestatistics[0]

    d.Set("version", Ltestatistics.Version)
    d.Set("end_time", Ltestatistics.EndTime)
    d.Set("start_time", Ltestatistics.StartTime)
    d.Set("stats_data", Ltestatistics.StatsData)
    
    d.Set("id", Ltestatistics.Identifier())
    d.Set("parent_id", Ltestatistics.ParentID)
    d.Set("parent_type", Ltestatistics.ParentType)
    d.Set("owner", Ltestatistics.Owner)

    d.SetId(Ltestatistics.Identifier())
    
    return nil
}