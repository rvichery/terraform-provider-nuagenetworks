package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceZoneTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceZoneTemplateRead,
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
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
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
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceZoneTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredZoneTemplates := vspk.ZoneTemplatesList{}
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
    parent := &vspk.DomainTemplate{ID: d.Get("parent_domain_template").(string)}
    filteredZoneTemplates, err = parent.ZoneTemplates(fetchFilter)
    if err != nil {
        return err
    }

    ZoneTemplate := &vspk.ZoneTemplate{}

    if len(filteredZoneTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredZoneTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ZoneTemplate = filteredZoneTemplates[0]

    d.Set("dpi", ZoneTemplate.DPI)
    d.Set("ip_type", ZoneTemplate.IPType)
    d.Set("ipv6_address", ZoneTemplate.IPv6Address)
    d.Set("name", ZoneTemplate.Name)
    d.Set("last_updated_by", ZoneTemplate.LastUpdatedBy)
    d.Set("address", ZoneTemplate.Address)
    d.Set("description", ZoneTemplate.Description)
    d.Set("netmask", ZoneTemplate.Netmask)
    d.Set("encryption", ZoneTemplate.Encryption)
    d.Set("entity_scope", ZoneTemplate.EntityScope)
    d.Set("associated_multicast_channel_map_id", ZoneTemplate.AssociatedMulticastChannelMapID)
    d.Set("public_zone", ZoneTemplate.PublicZone)
    d.Set("multicast", ZoneTemplate.Multicast)
    d.Set("number_of_hosts_in_subnets", ZoneTemplate.NumberOfHostsInSubnets)
    d.Set("external_id", ZoneTemplate.ExternalID)
    d.Set("dynamic_ipv6_address", ZoneTemplate.DynamicIpv6Address)
    
    d.Set("id", ZoneTemplate.Identifier())
    d.Set("parent_id", ZoneTemplate.ParentID)
    d.Set("parent_type", ZoneTemplate.ParentType)
    d.Set("owner", ZoneTemplate.Owner)

    d.SetId(ZoneTemplate.Identifier())
    
    return nil
}