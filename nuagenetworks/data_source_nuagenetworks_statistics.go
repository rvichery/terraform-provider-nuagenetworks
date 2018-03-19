package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceStatistics() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceStatisticsRead,
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
            "version": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "end_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "start_time": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "stats_data": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "number_of_data_points": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "parent_container_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_ingress_adv_fwd_entry_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_egress_adv_fwd_entry_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_egress_acl_entry_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_ingress_external_service_template_entry": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_vm_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_bridge_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_address_map": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_virtual_firewall_rule": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_ingress_acl_entry_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_host_interface", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_host_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_ns_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_patnat_pool", "parent_wireless_port"},
            },
            "parent_patnat_pool": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_wireless_port"},
            },
            "parent_wireless_port": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_ingress_adv_fwd_entry_template", "parent_egress_adv_fwd_entry_template", "parent_zone", "parent_egress_acl_entry_template", "parent_ingress_external_service_template_entry", "parent_domain", "parent_vport", "parent_subnet", "parent_vlan", "parent_vm_interface", "parent_bridge_interface", "parent_address_map", "parent_virtual_firewall_rule", "parent_ingress_acl_entry_template", "parent_l2_domain", "parent_host_interface", "parent_ns_port", "parent_patnat_pool"},
            },
        },
    }
}


func dataSourceStatisticsRead(d *schema.ResourceData, m interface{}) error {
    filteredStatistics := vspk.StatisticsList{}
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
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ingress_adv_fwd_entry_template"); ok {
        parent := &vspk.IngressAdvFwdEntryTemplate{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_egress_adv_fwd_entry_template"); ok {
        parent := &vspk.EgressAdvFwdEntryTemplate{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_egress_acl_entry_template"); ok {
        parent := &vspk.EgressACLEntryTemplate{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ingress_external_service_template_entry"); ok {
        parent := &vspk.IngressExternalServiceTemplateEntry{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vm_interface"); ok {
        parent := &vspk.VMInterface{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
        parent := &vspk.BridgeInterface{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_address_map"); ok {
        parent := &vspk.AddressMap{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_virtual_firewall_rule"); ok {
        parent := &vspk.VirtualFirewallRule{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ingress_acl_entry_template"); ok {
        parent := &vspk.IngressACLEntryTemplate{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_host_interface"); ok {
        parent := &vspk.HostInterface{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ns_port"); ok {
        parent := &vspk.NSPort{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_patnat_pool"); ok {
        parent := &vspk.PATNATPool{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_wireless_port"); ok {
        parent := &vspk.WirelessPort{ID: attr.(string)}
        filteredStatistics, err = parent.Statistics(fetchFilter)
        if err != nil {
            return err
        }
    }

    Statistics := &vspk.Statistics{}

    if len(filteredStatistics) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredStatistics) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    Statistics = filteredStatistics[0]

    d.Set("version", Statistics.Version)
    d.Set("end_time", Statistics.EndTime)
    d.Set("start_time", Statistics.StartTime)
    d.Set("stats_data", Statistics.StatsData)
    d.Set("number_of_data_points", Statistics.NumberOfDataPoints)
    
    d.Set("id", Statistics.Identifier())
    d.Set("parent_id", Statistics.ParentID)
    d.Set("parent_type", Statistics.ParentType)
    d.Set("owner", Statistics.Owner)

    d.SetId(Statistics.Identifier())
    
    return nil
}