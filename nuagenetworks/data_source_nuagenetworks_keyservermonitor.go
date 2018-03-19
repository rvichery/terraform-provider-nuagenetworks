package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceKeyServerMonitor() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceKeyServerMonitorRead,
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
            "last_update_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "gateway_secured_data_record_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "keyserver_monitor_encrypted_sek_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "keyserver_monitor_encrypted_seed_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "keyserver_monitor_sek_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "keyserver_monitor_seed_count": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "enterprise_secured_data_record_count": &schema.Schema{
                Type:     schema.TypeInt,
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


func dataSourceKeyServerMonitorRead(d *schema.ResourceData, m interface{}) error {
    filteredKeyServerMonitors := vspk.KeyServerMonitorsList{}
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
    filteredKeyServerMonitors, err = parent.KeyServerMonitors(fetchFilter)
    if err != nil {
        return err
    }

    KeyServerMonitor := &vspk.KeyServerMonitor{}

    if len(filteredKeyServerMonitors) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredKeyServerMonitors) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    KeyServerMonitor = filteredKeyServerMonitors[0]

    d.Set("last_update_time", KeyServerMonitor.LastUpdateTime)
    d.Set("last_updated_by", KeyServerMonitor.LastUpdatedBy)
    d.Set("gateway_secured_data_record_count", KeyServerMonitor.GatewaySecuredDataRecordCount)
    d.Set("keyserver_monitor_encrypted_sek_count", KeyServerMonitor.KeyserverMonitorEncryptedSEKCount)
    d.Set("keyserver_monitor_encrypted_seed_count", KeyServerMonitor.KeyserverMonitorEncryptedSeedCount)
    d.Set("keyserver_monitor_sek_count", KeyServerMonitor.KeyserverMonitorSEKCount)
    d.Set("keyserver_monitor_seed_count", KeyServerMonitor.KeyserverMonitorSeedCount)
    d.Set("enterprise_secured_data_record_count", KeyServerMonitor.EnterpriseSecuredDataRecordCount)
    d.Set("entity_scope", KeyServerMonitor.EntityScope)
    d.Set("external_id", KeyServerMonitor.ExternalID)
    
    d.Set("id", KeyServerMonitor.Identifier())
    d.Set("parent_id", KeyServerMonitor.ParentID)
    d.Set("parent_type", KeyServerMonitor.ParentType)
    d.Set("owner", KeyServerMonitor.Owner)

    d.SetId(KeyServerMonitor.Identifier())
    
    return nil
}