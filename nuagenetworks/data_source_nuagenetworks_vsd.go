package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVSD() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVSDRead,
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
            "url": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "management_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_state_change": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "peak_cpuusage": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "peak_memory_usage": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "peer_addresses": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "messages": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "disks": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "already_marked_for_unavailable": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "unavailable_timestamp": &schema.Schema{
                Type:     schema.TypeInt,
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
            "mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "product_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "current_cpuusage": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "current_memory_usage": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "average_cpuusage": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "average_memory_usage": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_vsp": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceVSDRead(d *schema.ResourceData, m interface{}) error {
    filteredVSDs := vspk.VSDsList{}
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
    parent := &vspk.VSP{ID: d.Get("parent_vsp").(string)}
    filteredVSDs, err = parent.VSDs(fetchFilter)
    if err != nil {
        return err
    }

    VSD := &vspk.VSD{}

    if len(filteredVSDs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVSDs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VSD = filteredVSDs[0]

    d.Set("url", VSD.URL)
    d.Set("name", VSD.Name)
    d.Set("management_ip", VSD.ManagementIP)
    d.Set("last_state_change", VSD.LastStateChange)
    d.Set("last_updated_by", VSD.LastUpdatedBy)
    d.Set("address", VSD.Address)
    d.Set("peak_cpuusage", VSD.PeakCPUUsage)
    d.Set("peak_memory_usage", VSD.PeakMemoryUsage)
    d.Set("peer_addresses", VSD.PeerAddresses)
    d.Set("description", VSD.Description)
    d.Set("messages", VSD.Messages)
    d.Set("disks", VSD.Disks)
    d.Set("already_marked_for_unavailable", VSD.AlreadyMarkedForUnavailable)
    d.Set("unavailable_timestamp", VSD.UnavailableTimestamp)
    d.Set("entity_scope", VSD.EntityScope)
    d.Set("location", VSD.Location)
    d.Set("mode", VSD.Mode)
    d.Set("product_version", VSD.ProductVersion)
    d.Set("status", VSD.Status)
    d.Set("current_cpuusage", VSD.CurrentCPUUsage)
    d.Set("current_memory_usage", VSD.CurrentMemoryUsage)
    d.Set("average_cpuusage", VSD.AverageCPUUsage)
    d.Set("average_memory_usage", VSD.AverageMemoryUsage)
    d.Set("external_id", VSD.ExternalID)
    
    d.Set("id", VSD.Identifier())
    d.Set("parent_id", VSD.ParentID)
    d.Set("parent_type", VSD.ParentType)
    d.Set("owner", VSD.Owner)

    d.SetId(VSD.Identifier())
    
    return nil
}