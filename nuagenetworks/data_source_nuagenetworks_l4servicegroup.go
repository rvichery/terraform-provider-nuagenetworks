package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceL4ServiceGroup() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceL4ServiceGroupRead,
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
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceL4ServiceGroupRead(d *schema.ResourceData, m interface{}) error {
    filteredL4ServiceGroups := vspk.L4ServiceGroupsList{}
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
    parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
    filteredL4ServiceGroups, err = parent.L4ServiceGroups(fetchFilter)
    if err != nil {
        return err
    }

    L4ServiceGroup := &vspk.L4ServiceGroup{}

    if len(filteredL4ServiceGroups) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredL4ServiceGroups) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    L4ServiceGroup = filteredL4ServiceGroups[0]

    d.Set("name", L4ServiceGroup.Name)
    d.Set("last_updated_by", L4ServiceGroup.LastUpdatedBy)
    d.Set("description", L4ServiceGroup.Description)
    d.Set("entity_scope", L4ServiceGroup.EntityScope)
    d.Set("external_id", L4ServiceGroup.ExternalID)
    
    d.Set("id", L4ServiceGroup.Identifier())
    d.Set("parent_id", L4ServiceGroup.ParentID)
    d.Set("parent_type", L4ServiceGroup.ParentType)
    d.Set("owner", L4ServiceGroup.Owner)

    d.SetId(L4ServiceGroup.Identifier())
    
    return nil
}