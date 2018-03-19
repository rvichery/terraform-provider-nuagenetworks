package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePATNATPool() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePATNATPoolRead,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address_range": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "default_patip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "permitted_action": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "end_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "end_source_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_gateway_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_subnet_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_vlan_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "start_address_range": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "start_source_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "dynamic_source_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_gateway", "parent_ns_gateway", "parent_enterprise"},
            },
            "parent_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan", "parent_ns_gateway", "parent_enterprise"},
            },
            "parent_ns_gateway": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan", "parent_gateway", "parent_enterprise"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan", "parent_gateway", "parent_ns_gateway"},
            },
        },
    }
}


func dataSourcePATNATPoolRead(d *schema.ResourceData, m interface{}) error {
    filteredPATNATPools := vspk.PATNATPoolsList{}
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
    if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        filteredPATNATPools, err = parent.PATNATPools(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_gateway"); ok {
        parent := &vspk.Gateway{ID: attr.(string)}
        filteredPATNATPools, err = parent.PATNATPools(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
        parent := &vspk.NSGateway{ID: attr.(string)}
        filteredPATNATPools, err = parent.PATNATPools(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredPATNATPools, err = parent.PATNATPools(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredPATNATPools, err = parent.PATNATPools(fetchFilter)
        if err != nil {
            return err
        }
    }

    PATNATPool := &vspk.PATNATPool{}

    if len(filteredPATNATPools) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPATNATPools) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    PATNATPool = filteredPATNATPools[0]

    d.Set("name", PATNATPool.Name)
    d.Set("last_updated_by", PATNATPool.LastUpdatedBy)
    d.Set("address_range", PATNATPool.AddressRange)
    d.Set("default_patip", PATNATPool.DefaultPATIP)
    d.Set("permitted_action", PATNATPool.PermittedAction)
    d.Set("description", PATNATPool.Description)
    d.Set("end_address_range", PATNATPool.EndAddressRange)
    d.Set("end_source_address", PATNATPool.EndSourceAddress)
    d.Set("entity_scope", PATNATPool.EntityScope)
    d.Set("associated_gateway_id", PATNATPool.AssociatedGatewayId)
    d.Set("associated_gateway_type", PATNATPool.AssociatedGatewayType)
    d.Set("associated_subnet_id", PATNATPool.AssociatedSubnetId)
    d.Set("associated_vlan_id", PATNATPool.AssociatedVlanId)
    d.Set("start_address_range", PATNATPool.StartAddressRange)
    d.Set("start_source_address", PATNATPool.StartSourceAddress)
    d.Set("external_id", PATNATPool.ExternalID)
    d.Set("dynamic_source_enabled", PATNATPool.DynamicSourceEnabled)
    
    d.Set("id", PATNATPool.Identifier())
    d.Set("parent_id", PATNATPool.ParentID)
    d.Set("parent_type", PATNATPool.ParentType)
    d.Set("owner", PATNATPool.Owner)

    d.SetId(PATNATPool.Identifier())
    
    return nil
}