package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVNFThresholdPolicy() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVNFThresholdPolicyRead,
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
            "cpu_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "memory_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "min_occurrence": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "monit_interval": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "storage_threshold": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceVNFThresholdPolicyRead(d *schema.ResourceData, m interface{}) error {
    filteredVNFThresholdPolicies := vspk.VNFThresholdPoliciesList{}
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
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredVNFThresholdPolicies, err = parent.VNFThresholdPolicies(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredVNFThresholdPolicies, err = parent.VNFThresholdPolicies(fetchFilter)
        if err != nil {
            return err
        }
    }

    VNFThresholdPolicy := &vspk.VNFThresholdPolicy{}

    if len(filteredVNFThresholdPolicies) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVNFThresholdPolicies) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VNFThresholdPolicy = filteredVNFThresholdPolicies[0]

    d.Set("cpu_threshold", VNFThresholdPolicy.CPUThreshold)
    d.Set("name", VNFThresholdPolicy.Name)
    d.Set("action", VNFThresholdPolicy.Action)
    d.Set("memory_threshold", VNFThresholdPolicy.MemoryThreshold)
    d.Set("description", VNFThresholdPolicy.Description)
    d.Set("min_occurrence", VNFThresholdPolicy.MinOccurrence)
    d.Set("monit_interval", VNFThresholdPolicy.MonitInterval)
    d.Set("storage_threshold", VNFThresholdPolicy.StorageThreshold)
    
    d.Set("id", VNFThresholdPolicy.Identifier())
    d.Set("parent_id", VNFThresholdPolicy.ParentID)
    d.Set("parent_type", VNFThresholdPolicy.ParentType)
    d.Set("owner", VNFThresholdPolicy.Owner)

    d.SetId(VNFThresholdPolicy.Identifier())
    
    return nil
}