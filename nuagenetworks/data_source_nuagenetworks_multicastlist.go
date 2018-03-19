package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceMultiCastList() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceMultiCastListRead,
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
            "mcast_type": &schema.Schema{
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
            "parent_enterprise_profile": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise_profile"},
            },
        },
    }
}


func dataSourceMultiCastListRead(d *schema.ResourceData, m interface{}) error {
    filteredMultiCastLists := vspk.MultiCastListsList{}
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
    if attr, ok := d.GetOk("parent_enterprise_profile"); ok {
        parent := &vspk.EnterpriseProfile{ID: attr.(string)}
        filteredMultiCastLists, err = parent.MultiCastLists(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredMultiCastLists, err = parent.MultiCastLists(fetchFilter)
        if err != nil {
            return err
        }
    }

    MultiCastList := &vspk.MultiCastList{}

    if len(filteredMultiCastLists) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredMultiCastLists) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    MultiCastList = filteredMultiCastLists[0]

    d.Set("last_updated_by", MultiCastList.LastUpdatedBy)
    d.Set("mcast_type", MultiCastList.McastType)
    d.Set("entity_scope", MultiCastList.EntityScope)
    d.Set("external_id", MultiCastList.ExternalID)
    
    d.Set("id", MultiCastList.Identifier())
    d.Set("parent_id", MultiCastList.ParentID)
    d.Set("parent_type", MultiCastList.ParentType)
    d.Set("owner", MultiCastList.Owner)

    d.SetId(MultiCastList.Identifier())
    
    return nil
}