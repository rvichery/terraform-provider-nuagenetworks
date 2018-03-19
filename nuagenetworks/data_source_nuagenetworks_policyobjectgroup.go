package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePolicyObjectGroup() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePolicyObjectGroupRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
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


func dataSourcePolicyObjectGroupRead(d *schema.ResourceData, m interface{}) error {
    filteredPolicyObjectGroups := vspk.PolicyObjectGroupsList{}
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
    filteredPolicyObjectGroups, err = parent.PolicyObjectGroups(fetchFilter)
    if err != nil {
        return err
    }

    PolicyObjectGroup := &vspk.PolicyObjectGroup{}

    if len(filteredPolicyObjectGroups) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPolicyObjectGroups) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    PolicyObjectGroup = filteredPolicyObjectGroups[0]

    d.Set("name", PolicyObjectGroup.Name)
    d.Set("description", PolicyObjectGroup.Description)
    d.Set("type", PolicyObjectGroup.Type)
    
    d.Set("id", PolicyObjectGroup.Identifier())
    d.Set("parent_id", PolicyObjectGroup.ParentID)
    d.Set("parent_type", PolicyObjectGroup.ParentType)
    d.Set("owner", PolicyObjectGroup.Owner)

    d.SetId(PolicyObjectGroup.Identifier())
    
    return nil
}