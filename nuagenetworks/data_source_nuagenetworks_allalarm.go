package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceAllAlarm() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceAllAlarmRead,
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
            "target_object": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "acknowledged": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "reason": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "severity": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "timestamp": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "error_condition": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "number_of_occurances": &schema.Schema{
                Type:     schema.TypeInt,
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


func dataSourceAllAlarmRead(d *schema.ResourceData, m interface{}) error {
    filteredAllAlarms := vspk.AllAlarmsList{}
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
    filteredAllAlarms, err = parent.AllAlarms(fetchFilter)
    if err != nil {
        return err
    }

    AllAlarm := &vspk.AllAlarm{}

    if len(filteredAllAlarms) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredAllAlarms) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    AllAlarm = filteredAllAlarms[0]

    d.Set("name", AllAlarm.Name)
    d.Set("target_object", AllAlarm.TargetObject)
    d.Set("last_updated_by", AllAlarm.LastUpdatedBy)
    d.Set("acknowledged", AllAlarm.Acknowledged)
    d.Set("reason", AllAlarm.Reason)
    d.Set("description", AllAlarm.Description)
    d.Set("severity", AllAlarm.Severity)
    d.Set("timestamp", AllAlarm.Timestamp)
    d.Set("enterprise_id", AllAlarm.EnterpriseID)
    d.Set("entity_scope", AllAlarm.EntityScope)
    d.Set("error_condition", AllAlarm.ErrorCondition)
    d.Set("number_of_occurances", AllAlarm.NumberOfOccurances)
    d.Set("external_id", AllAlarm.ExternalID)
    
    d.Set("id", AllAlarm.Identifier())
    d.Set("parent_id", AllAlarm.ParentID)
    d.Set("parent_type", AllAlarm.ParentType)
    d.Set("owner", AllAlarm.Owner)

    d.SetId(AllAlarm.Identifier())
    
    return nil
}