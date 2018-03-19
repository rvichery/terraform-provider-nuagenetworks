package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceWirelessPort() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceWirelessPortRead,
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
            "generic_config": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "physical_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "wifi_frequency_band": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "wifi_mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "country_code": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "frequency_channel": &schema.Schema{
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


func dataSourceWirelessPortRead(d *schema.ResourceData, m interface{}) error {
    filteredWirelessPorts := vspk.WirelessPortsList{}
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
    filteredWirelessPorts, err = parent.WirelessPorts(fetchFilter)
    if err != nil {
        return err
    }

    WirelessPort := &vspk.WirelessPort{}

    if len(filteredWirelessPorts) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredWirelessPorts) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    WirelessPort = filteredWirelessPorts[0]

    d.Set("name", WirelessPort.Name)
    d.Set("generic_config", WirelessPort.GenericConfig)
    d.Set("description", WirelessPort.Description)
    d.Set("physical_name", WirelessPort.PhysicalName)
    d.Set("wifi_frequency_band", WirelessPort.WifiFrequencyBand)
    d.Set("wifi_mode", WirelessPort.WifiMode)
    d.Set("port_type", WirelessPort.PortType)
    d.Set("country_code", WirelessPort.CountryCode)
    d.Set("frequency_channel", WirelessPort.FrequencyChannel)
    
    d.Set("id", WirelessPort.Identifier())
    d.Set("parent_id", WirelessPort.ParentID)
    d.Set("parent_type", WirelessPort.ParentType)
    d.Set("owner", WirelessPort.Owner)

    d.SetId(WirelessPort.Identifier())
    
    return nil
}