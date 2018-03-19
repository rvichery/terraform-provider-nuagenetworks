package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVSP() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVSPRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "location": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "product_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}


func dataSourceVSPRead(d *schema.ResourceData, m interface{}) error {
    filteredVSPs := vspk.VSPsList{}
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
    filteredVSPs, err = parent.VSPs(fetchFilter)
    if err != nil {
        return err
    }

    VSP := &vspk.VSP{}

    if len(filteredVSPs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVSPs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VSP = filteredVSPs[0]

    d.Set("name", VSP.Name)
    d.Set("last_updated_by", VSP.LastUpdatedBy)
    d.Set("description", VSP.Description)
    d.Set("entity_scope", VSP.EntityScope)
    d.Set("location", VSP.Location)
    d.Set("product_version", VSP.ProductVersion)
    d.Set("external_id", VSP.ExternalID)
    
    d.Set("id", VSP.Identifier())
    d.Set("parent_id", VSP.ParentID)
    d.Set("parent_type", VSP.ParentType)
    d.Set("owner", VSP.Owner)

    d.SetId(VSP.Identifier())
    
    return nil
}