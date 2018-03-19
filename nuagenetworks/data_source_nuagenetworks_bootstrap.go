package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceBootstrap() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceBootstrapRead,
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
            "zfb_info": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zfb_match_attribute": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zfb_match_value": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "installer_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceBootstrapRead(d *schema.ResourceData, m interface{}) error {
    filteredBootstraps := vspk.BootstrapsList{}
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
    parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
    filteredBootstraps, err = parent.Bootstraps(fetchFilter)
    if err != nil {
        return err
    }

    Bootstrap := &vspk.Bootstrap{}

    if len(filteredBootstraps) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredBootstraps) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Bootstrap = filteredBootstraps[0]

    d.Set("zfb_info", Bootstrap.ZFBInfo)
    d.Set("zfb_match_attribute", Bootstrap.ZFBMatchAttribute)
    d.Set("zfb_match_value", Bootstrap.ZFBMatchValue)
    d.Set("last_updated_by", Bootstrap.LastUpdatedBy)
    d.Set("installer_id", Bootstrap.InstallerID)
    d.Set("entity_scope", Bootstrap.EntityScope)
    d.Set("status", Bootstrap.Status)
    d.Set("external_id", Bootstrap.ExternalID)
    
    d.Set("id", Bootstrap.Identifier())
    d.Set("parent_id", Bootstrap.ParentID)
    d.Set("parent_type", Bootstrap.ParentType)
    d.Set("owner", Bootstrap.Owner)

    d.SetId(Bootstrap.Identifier())
    
    return nil
}