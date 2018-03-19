package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAutodiscovereddatacenter() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAutodiscovereddatacenterRead,
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
            "managed_object_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_vcenter_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vcenter": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceAutodiscovereddatacenterRead(d *schema.ResourceData, m interface{}) error {
    filteredAutodiscovereddatacenters := vspk.AutodiscovereddatacentersList{}
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
    parent := &vspk.VCenter{ID: d.Get("parent_vcenter").(string)}
    filteredAutodiscovereddatacenters, err = parent.Autodiscovereddatacenters(fetchFilter)
    if err != nil {
        return err
    }

    Autodiscovereddatacenter := &vspk.Autodiscovereddatacenter{}

    if len(filteredAutodiscovereddatacenters) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAutodiscovereddatacenters) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Autodiscovereddatacenter = filteredAutodiscovereddatacenters[0]

    d.Set("name", Autodiscovereddatacenter.Name)
    d.Set("managed_object_id", Autodiscovereddatacenter.ManagedObjectID)
    d.Set("last_updated_by", Autodiscovereddatacenter.LastUpdatedBy)
    d.Set("entity_scope", Autodiscovereddatacenter.EntityScope)
    d.Set("associated_vcenter_id", Autodiscovereddatacenter.AssociatedVCenterID)
    d.Set("external_id", Autodiscovereddatacenter.ExternalID)
    
    d.Set("id", Autodiscovereddatacenter.Identifier())
    d.Set("parent_id", Autodiscovereddatacenter.ParentID)
    d.Set("parent_type", Autodiscovereddatacenter.ParentType)
    d.Set("owner", Autodiscovereddatacenter.Owner)

    d.SetId(Autodiscovereddatacenter.Identifier())
    
    return nil
}