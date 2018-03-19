package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceStaticRoute() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceStaticRouteRead,
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
            "bfd_enabled": &schema.Schema{
                Type:     schema.TypeBool,
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
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "netmask": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "next_hop_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "route_distinguisher": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_subnet_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_container_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_shared_network_resource", "parent_domain", "parent_vm_interface", "parent_host_interface"},
            },
            "parent_shared_network_resource": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vm_interface", "parent_host_interface"},
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_vm_interface", "parent_host_interface"},
            },
            "parent_vm_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_domain", "parent_host_interface"},
            },
            "parent_host_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_domain", "parent_vm_interface"},
            },
        },
    }
}


func dataSourceStaticRouteRead(d *schema.ResourceData, m interface{}) error {
    filteredStaticRoutes := vspk.StaticRoutesList{}
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
    if attr, ok := d.GetOk("parent_container_interface"); ok {
        parent := &vspk.ContainerInterface{ID: attr.(string)}
        filteredStaticRoutes, err = parent.StaticRoutes(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
        parent := &vspk.SharedNetworkResource{ID: attr.(string)}
        filteredStaticRoutes, err = parent.StaticRoutes(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredStaticRoutes, err = parent.StaticRoutes(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vm_interface"); ok {
        parent := &vspk.VMInterface{ID: attr.(string)}
        filteredStaticRoutes, err = parent.StaticRoutes(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_host_interface"); ok {
        parent := &vspk.HostInterface{ID: attr.(string)}
        filteredStaticRoutes, err = parent.StaticRoutes(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredStaticRoutes, err = parent.StaticRoutes(fetchFilter)
        if err != nil {
            return err
        }
    }

    StaticRoute := &vspk.StaticRoute{}

    if len(filteredStaticRoutes) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredStaticRoutes) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    StaticRoute = filteredStaticRoutes[0]

    d.Set("bfd_enabled", StaticRoute.BFDEnabled)
    d.Set("ip_type", StaticRoute.IPType)
    d.Set("ipv6_address", StaticRoute.IPv6Address)
    d.Set("last_updated_by", StaticRoute.LastUpdatedBy)
    d.Set("address", StaticRoute.Address)
    d.Set("netmask", StaticRoute.Netmask)
    d.Set("next_hop_ip", StaticRoute.NextHopIp)
    d.Set("entity_scope", StaticRoute.EntityScope)
    d.Set("route_distinguisher", StaticRoute.RouteDistinguisher)
    d.Set("associated_subnet_id", StaticRoute.AssociatedSubnetID)
    d.Set("external_id", StaticRoute.ExternalID)
    d.Set("type", StaticRoute.Type)
    
    d.Set("id", StaticRoute.Identifier())
    d.Set("parent_id", StaticRoute.ParentID)
    d.Set("parent_type", StaticRoute.ParentType)
    d.Set("owner", StaticRoute.Owner)

    d.SetId(StaticRoute.Identifier())
    
    return nil
}