package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceCustomProperty() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceCustomPropertyRead,
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
            "attribute_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "attribute_value": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_uplink_connection": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceCustomPropertyRead(d *schema.ResourceData, m interface{}) error {
    filteredCustomProperties := vspk.CustomPropertiesList{}
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
    parent := &vspk.UplinkConnection{ID: d.Get("parent_uplink_connection").(string)}
    filteredCustomProperties, err = parent.CustomProperties(fetchFilter)
    if err != nil {
        return err
    }

    CustomProperty := &vspk.CustomProperty{}

    if len(filteredCustomProperties) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredCustomProperties) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    CustomProperty = filteredCustomProperties[0]

    d.Set("attribute_name", CustomProperty.AttributeName)
    d.Set("attribute_value", CustomProperty.AttributeValue)
    
    d.Set("id", CustomProperty.Identifier())
    d.Set("parent_id", CustomProperty.ParentID)
    d.Set("parent_type", CustomProperty.ParentType)
    d.Set("owner", CustomProperty.Owner)

    d.SetId(CustomProperty.Identifier())
    
    return nil
}