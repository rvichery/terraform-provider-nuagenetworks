package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceSSIDConnection() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceSSIDConnectionRead,
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
            "passphrase": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "redirect_option": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "redirect_url": &schema.Schema{
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
            "white_list": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "black_list": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "interface_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "vport_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "broadcast_ssid": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "associated_captive_portal_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_egress_qos_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "authentication_mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_wireless_port": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceSSIDConnectionRead(d *schema.ResourceData, m interface{}) error {
    filteredSSIDConnections := vspk.SSIDConnectionsList{}
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
    parent := &vspk.WirelessPort{ID: d.Get("parent_wireless_port").(string)}
    filteredSSIDConnections, err = parent.SSIDConnections(fetchFilter)
    if err != nil {
        return err
    }

    SSIDConnection := &vspk.SSIDConnection{}

    if len(filteredSSIDConnections) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredSSIDConnections) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    SSIDConnection = filteredSSIDConnections[0]

    d.Set("name", SSIDConnection.Name)
    d.Set("passphrase", SSIDConnection.Passphrase)
    d.Set("redirect_option", SSIDConnection.RedirectOption)
    d.Set("redirect_url", SSIDConnection.RedirectURL)
    d.Set("generic_config", SSIDConnection.GenericConfig)
    d.Set("description", SSIDConnection.Description)
    d.Set("white_list", SSIDConnection.WhiteList)
    d.Set("black_list", SSIDConnection.BlackList)
    d.Set("interface_name", SSIDConnection.InterfaceName)
    d.Set("vport_id", SSIDConnection.VportID)
    d.Set("broadcast_ssid", SSIDConnection.BroadcastSSID)
    d.Set("associated_captive_portal_profile_id", SSIDConnection.AssociatedCaptivePortalProfileID)
    d.Set("associated_egress_qos_policy_id", SSIDConnection.AssociatedEgressQOSPolicyID)
    d.Set("authentication_mode", SSIDConnection.AuthenticationMode)
    
    d.Set("id", SSIDConnection.Identifier())
    d.Set("parent_id", SSIDConnection.ParentID)
    d.Set("parent_type", SSIDConnection.ParentType)
    d.Set("owner", SSIDConnection.Owner)

    d.SetId(SSIDConnection.Identifier())
    
    return nil
}