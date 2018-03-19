package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceZone() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceZoneRead,
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
            "dpi": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ipv6_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "maintenance_mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
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
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "flow_collection_enabled": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "encryption": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "policy_group_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "associated_multicast_channel_map_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "public_zone": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "multicast": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "number_of_hosts_in_subnets": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dynamic_ipv6_address": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceZoneRead(d *schema.ResourceData, m interface{}) error {
    filteredZones := vspk.ZonesList{}
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
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredZones, err = parent.Zones(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredZones, err = parent.Zones(fetchFilter)
        if err != nil {
            return err
        }
    }

    Zone := &vspk.Zone{}

    if len(filteredZones) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredZones) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Zone = filteredZones[0]

    d.Set("dpi", Zone.DPI)
    d.Set("ip_type", Zone.IPType)
    d.Set("ipv6_address", Zone.IPv6Address)
    d.Set("maintenance_mode", Zone.MaintenanceMode)
    d.Set("name", Zone.Name)
    d.Set("last_updated_by", Zone.LastUpdatedBy)
    d.Set("address", Zone.Address)
    d.Set("template_id", Zone.TemplateID)
    d.Set("description", Zone.Description)
    d.Set("netmask", Zone.Netmask)
    d.Set("flow_collection_enabled", Zone.FlowCollectionEnabled)
    d.Set("encryption", Zone.Encryption)
    d.Set("entity_scope", Zone.EntityScope)
    d.Set("policy_group_id", Zone.PolicyGroupID)
    d.Set("associated_multicast_channel_map_id", Zone.AssociatedMulticastChannelMapID)
    d.Set("public_zone", Zone.PublicZone)
    d.Set("multicast", Zone.Multicast)
    d.Set("number_of_hosts_in_subnets", Zone.NumberOfHostsInSubnets)
    d.Set("external_id", Zone.ExternalID)
    d.Set("dynamic_ipv6_address", Zone.DynamicIpv6Address)
    
    d.Set("id", Zone.Identifier())
    d.Set("parent_id", Zone.ParentID)
    d.Set("parent_type", Zone.ParentType)
    d.Set("owner", Zone.Owner)

    d.SetId(Zone.Identifier())
    
    return nil
}