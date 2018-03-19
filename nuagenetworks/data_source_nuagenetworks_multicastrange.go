package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceMultiCastRange() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceMultiCastRangeRead,
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
            "max_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "min_address": &schema.Schema{
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
            "parent_multi_cast_channel_map": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceMultiCastRangeRead(d *schema.ResourceData, m interface{}) error {
    filteredMultiCastRanges := vspk.MultiCastRangesList{}
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
    parent := &vspk.MultiCastChannelMap{ID: d.Get("parent_multi_cast_channel_map").(string)}
    filteredMultiCastRanges, err = parent.MultiCastRanges(fetchFilter)
    if err != nil {
        return err
    }

    MultiCastRange := &vspk.MultiCastRange{}

    if len(filteredMultiCastRanges) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredMultiCastRanges) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    MultiCastRange = filteredMultiCastRanges[0]

    d.Set("last_updated_by", MultiCastRange.LastUpdatedBy)
    d.Set("max_address", MultiCastRange.MaxAddress)
    d.Set("min_address", MultiCastRange.MinAddress)
    d.Set("entity_scope", MultiCastRange.EntityScope)
    d.Set("external_id", MultiCastRange.ExternalID)
    
    d.Set("id", MultiCastRange.Identifier())
    d.Set("parent_id", MultiCastRange.ParentID)
    d.Set("parent_type", MultiCastRange.ParentType)
    d.Set("owner", MultiCastRange.Owner)

    d.SetId(MultiCastRange.Identifier())
    
    return nil
}