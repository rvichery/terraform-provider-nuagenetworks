package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceKeyServerMonitorSeed() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceKeyServerMonitorSeedRead,
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
            "seed_traffic_authentication_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "seed_traffic_encryption_algorithm": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "seed_traffic_encryption_key_lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "lifetime": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "creation_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "start_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_key_server_monitor": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceKeyServerMonitorSeedRead(d *schema.ResourceData, m interface{}) error {
    filteredKeyServerMonitorSeeds := vspk.KeyServerMonitorSeedsList{}
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
    parent := &vspk.KeyServerMonitor{ID: d.Get("parent_key_server_monitor").(string)}
    filteredKeyServerMonitorSeeds, err = parent.KeyServerMonitorSeeds(fetchFilter)
    if err != nil {
        return err
    }

    KeyServerMonitorSeed := &vspk.KeyServerMonitorSeed{}

    if len(filteredKeyServerMonitorSeeds) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredKeyServerMonitorSeeds) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    KeyServerMonitorSeed = filteredKeyServerMonitorSeeds[0]

    d.Set("last_updated_by", KeyServerMonitorSeed.LastUpdatedBy)
    d.Set("seed_traffic_authentication_algorithm", KeyServerMonitorSeed.SeedTrafficAuthenticationAlgorithm)
    d.Set("seed_traffic_encryption_algorithm", KeyServerMonitorSeed.SeedTrafficEncryptionAlgorithm)
    d.Set("seed_traffic_encryption_key_lifetime", KeyServerMonitorSeed.SeedTrafficEncryptionKeyLifetime)
    d.Set("lifetime", KeyServerMonitorSeed.Lifetime)
    d.Set("entity_scope", KeyServerMonitorSeed.EntityScope)
    d.Set("creation_time", KeyServerMonitorSeed.CreationTime)
    d.Set("start_time", KeyServerMonitorSeed.StartTime)
    d.Set("external_id", KeyServerMonitorSeed.ExternalID)
    
    d.Set("id", KeyServerMonitorSeed.Identifier())
    d.Set("parent_id", KeyServerMonitorSeed.ParentID)
    d.Set("parent_type", KeyServerMonitorSeed.ParentType)
    d.Set("owner", KeyServerMonitorSeed.Owner)

    d.SetId(KeyServerMonitorSeed.Identifier())
    
    return nil
}