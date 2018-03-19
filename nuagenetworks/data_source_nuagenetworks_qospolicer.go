package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceQosPolicer() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceQosPolicerRead,
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
            "rate": &schema.Schema{
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
            "burst": &schema.Schema{
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


func dataSourceQosPolicerRead(d *schema.ResourceData, m interface{}) error {
    filteredQosPolicers := vspk.QosPolicersList{}
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
    filteredQosPolicers, err = parent.QosPolicers(fetchFilter)
    if err != nil {
        return err
    }

    QosPolicer := &vspk.QosPolicer{}

    if len(filteredQosPolicers) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredQosPolicers) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    QosPolicer = filteredQosPolicers[0]

    d.Set("name", QosPolicer.Name)
    d.Set("last_updated_by", QosPolicer.LastUpdatedBy)
    d.Set("rate", QosPolicer.Rate)
    d.Set("description", QosPolicer.Description)
    d.Set("entity_scope", QosPolicer.EntityScope)
    d.Set("burst", QosPolicer.Burst)
    d.Set("external_id", QosPolicer.ExternalID)
    
    d.Set("id", QosPolicer.Identifier())
    d.Set("parent_id", QosPolicer.ParentID)
    d.Set("parent_type", QosPolicer.ParentType)
    d.Set("owner", QosPolicer.Owner)

    d.SetId(QosPolicer.Identifier())
    
    return nil
}