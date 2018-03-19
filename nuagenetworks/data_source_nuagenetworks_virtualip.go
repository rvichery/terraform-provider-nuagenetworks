package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVirtualIP() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVirtualIPRead,
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
            "mac": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ip_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "virtual_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_floating_ip_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "subnet_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_redirection_target": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vport", "parent_subnet"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_redirection_target", "parent_subnet"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_redirection_target", "parent_vport"},
            },
        },
    }
}


func dataSourceVirtualIPRead(d *schema.ResourceData, m interface{}) error {
    filteredVirtualIPs := vspk.VirtualIPsList{}
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
    if attr, ok := d.GetOk("parent_redirection_target"); ok {
        parent := &vspk.RedirectionTarget{ID: attr.(string)}
        filteredVirtualIPs, err = parent.VirtualIPs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredVirtualIPs, err = parent.VirtualIPs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredVirtualIPs, err = parent.VirtualIPs(fetchFilter)
        if err != nil {
            return err
        }
    }

    VirtualIP := &vspk.VirtualIP{}

    if len(filteredVirtualIPs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVirtualIPs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VirtualIP = filteredVirtualIPs[0]

    d.Set("mac", VirtualIP.MAC)
    d.Set("ip_type", VirtualIP.IPType)
    d.Set("last_updated_by", VirtualIP.LastUpdatedBy)
    d.Set("virtual_ip", VirtualIP.VirtualIP)
    d.Set("entity_scope", VirtualIP.EntityScope)
    d.Set("associated_floating_ip_id", VirtualIP.AssociatedFloatingIPID)
    d.Set("subnet_id", VirtualIP.SubnetID)
    d.Set("external_id", VirtualIP.ExternalID)
    
    d.Set("id", VirtualIP.Identifier())
    d.Set("parent_id", VirtualIP.ParentID)
    d.Set("parent_type", VirtualIP.ParentType)
    d.Set("owner", VirtualIP.Owner)

    d.SetId(VirtualIP.Identifier())
    
    return nil
}