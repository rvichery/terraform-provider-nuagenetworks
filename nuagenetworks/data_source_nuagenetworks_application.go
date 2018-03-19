package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceApplication() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceApplicationRead,
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
            "dscp": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "bandwidth": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "read_only": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "performance_monitor_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "destination_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enable_pps": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "one_way_delay": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "one_way_jitter": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "one_way_loss": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "post_classification_path": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "source_port": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "app_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "optimize_path_selection": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "pre_classification_path": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "protocol": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_l7_application_signature_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ether_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "symmetry": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "parent_l7applicationsignature": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l7applicationsignature"},
            },
        },
    }
}


func dataSourceApplicationRead(d *schema.ResourceData, m interface{}) error {
    filteredApplications := vspk.ApplicationsList{}
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
    if attr, ok := d.GetOk("parent_l7applicationsignature"); ok {
        parent := &vspk.L7applicationsignature{ID: attr.(string)}
        filteredApplications, err = parent.Applications(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredApplications, err = parent.Applications(fetchFilter)
        if err != nil {
            return err
        }
    }

    Application := &vspk.Application{}

    if len(filteredApplications) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredApplications) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Application = filteredApplications[0]

    d.Set("dscp", Application.DSCP)
    d.Set("name", Application.Name)
    d.Set("bandwidth", Application.Bandwidth)
    d.Set("last_updated_by", Application.LastUpdatedBy)
    d.Set("read_only", Application.ReadOnly)
    d.Set("performance_monitor_type", Application.PerformanceMonitorType)
    d.Set("description", Application.Description)
    d.Set("destination_ip", Application.DestinationIP)
    d.Set("destination_port", Application.DestinationPort)
    d.Set("enable_pps", Application.EnablePPS)
    d.Set("one_way_delay", Application.OneWayDelay)
    d.Set("one_way_jitter", Application.OneWayJitter)
    d.Set("one_way_loss", Application.OneWayLoss)
    d.Set("entity_scope", Application.EntityScope)
    d.Set("post_classification_path", Application.PostClassificationPath)
    d.Set("source_ip", Application.SourceIP)
    d.Set("source_port", Application.SourcePort)
    d.Set("app_id", Application.AppId)
    d.Set("optimize_path_selection", Application.OptimizePathSelection)
    d.Set("pre_classification_path", Application.PreClassificationPath)
    d.Set("protocol", Application.Protocol)
    d.Set("associated_l7_application_signature_id", Application.AssociatedL7ApplicationSignatureID)
    d.Set("ether_type", Application.EtherType)
    d.Set("external_id", Application.ExternalID)
    d.Set("symmetry", Application.Symmetry)
    
    d.Set("id", Application.Identifier())
    d.Set("parent_id", Application.ParentID)
    d.Set("parent_type", Application.ParentType)
    d.Set("owner", Application.Owner)

    d.SetId(Application.Identifier())
    
    return nil
}