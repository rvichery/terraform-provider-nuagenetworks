package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVM() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVMRead,
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
            "l2_domain_ids": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "vrsid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "uuid": &schema.Schema{
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
            "reason_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "delete_expiry": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "delete_mode": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "resync_info": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "site_identifier": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "interfaces": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "enterprise_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "domain_ids": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "zone_ids": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "orchestration_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "user_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "user_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "subnet_ids": &schema.Schema{
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "hypervisor_ip": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_qos": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_zone": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_vrs": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_egress_acl_template", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_egress_acl_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_user", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_user": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_enterprise", "parent_ingress_acl_template"},
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_ingress_acl_template"},
            },
            "parent_ingress_acl_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_qos", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vrs", "parent_l2_domain", "parent_egress_acl_template", "parent_user", "parent_enterprise"},
            },
        },
    }
}


func dataSourceVMRead(d *schema.ResourceData, m interface{}) error {
    filteredVMs := vspk.VMsList{}
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
    if attr, ok := d.GetOk("parent_qos"); ok {
        parent := &vspk.QOS{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_zone"); ok {
        parent := &vspk.Zone{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vrs"); ok {
        parent := &vspk.VRS{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_egress_acl_template"); ok {
        parent := &vspk.EgressACLTemplate{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_user"); ok {
        parent := &vspk.User{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_ingress_acl_template"); ok {
        parent := &vspk.IngressACLTemplate{ID: attr.(string)}
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredVMs, err = parent.VMs(fetchFilter)
        if err != nil {
            return err
        }
    }

    VM := &vspk.VM{}

    if len(filteredVMs) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVMs) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VM = filteredVMs[0]

    d.Set("l2_domain_ids", VM.L2DomainIDs)
    d.Set("vrsid", VM.VRSID)
    d.Set("uuid", VM.UUID)
    d.Set("name", VM.Name)
    d.Set("last_updated_by", VM.LastUpdatedBy)
    d.Set("reason_type", VM.ReasonType)
    d.Set("delete_expiry", VM.DeleteExpiry)
    d.Set("delete_mode", VM.DeleteMode)
    d.Set("resync_info", VM.ResyncInfo)
    d.Set("site_identifier", VM.SiteIdentifier)
    d.Set("interfaces", VM.Interfaces)
    d.Set("enterprise_id", VM.EnterpriseID)
    d.Set("enterprise_name", VM.EnterpriseName)
    d.Set("entity_scope", VM.EntityScope)
    d.Set("domain_ids", VM.DomainIDs)
    d.Set("zone_ids", VM.ZoneIDs)
    d.Set("orchestration_id", VM.OrchestrationID)
    d.Set("user_id", VM.UserID)
    d.Set("user_name", VM.UserName)
    d.Set("status", VM.Status)
    d.Set("subnet_ids", VM.SubnetIDs)
    d.Set("external_id", VM.ExternalID)
    d.Set("hypervisor_ip", VM.HypervisorIP)
    
    d.Set("id", VM.Identifier())
    d.Set("parent_id", VM.ParentID)
    d.Set("parent_type", VM.ParentType)
    d.Set("owner", VM.Owner)

    d.SetId(VM.Identifier())
    
    return nil
}