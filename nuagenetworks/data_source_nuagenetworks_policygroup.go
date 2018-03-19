package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourcePolicyGroup() *schema.Resource {
    return &schema.Resource{
        Read: dataSourcePolicyGroupRead,
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
            "evpn_community_tag": &schema.Schema{
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
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "description": &schema.Schema{
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
            "external": &schema.Schema{
                Type:     schema.TypeBool,
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
                ConflictsWith: []string{"parent_domain", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
            },
            "parent_vport": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
            },
            "parent_vm_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
            },
            "parent_bridge_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_vm_interface", "parent_l2_domain", "parent_host_interface"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_host_interface"},
            },
            "parent_host_interface": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain"},
            },
        },
    }
}


func dataSourcePolicyGroupRead(d *schema.ResourceData, m interface{}) error {
    filteredPolicyGroups := vspk.PolicyGroupsList{}
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
        filteredPolicyGroups, err = parent.PolicyGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        filteredPolicyGroups, err = parent.PolicyGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vport"); ok {
        parent := &vspk.VPort{ID: attr.(string)}
        filteredPolicyGroups, err = parent.PolicyGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vm_interface"); ok {
        parent := &vspk.VMInterface{ID: attr.(string)}
        filteredPolicyGroups, err = parent.PolicyGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
        parent := &vspk.BridgeInterface{ID: attr.(string)}
        filteredPolicyGroups, err = parent.PolicyGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredPolicyGroups, err = parent.PolicyGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_host_interface"); ok {
        parent := &vspk.HostInterface{ID: attr.(string)}
        filteredPolicyGroups, err = parent.PolicyGroups(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredPolicyGroups, err = parent.PolicyGroups(fetchFilter)
        if err != nil {
            return err
        }
    }

    PolicyGroup := &vspk.PolicyGroup{}

    if len(filteredPolicyGroups) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredPolicyGroups) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    PolicyGroup = filteredPolicyGroups[0]

    d.Set("evpn_community_tag", PolicyGroup.EVPNCommunityTag)
    d.Set("name", PolicyGroup.Name)
    d.Set("last_updated_by", PolicyGroup.LastUpdatedBy)
    d.Set("template_id", PolicyGroup.TemplateID)
    d.Set("description", PolicyGroup.Description)
    d.Set("entity_scope", PolicyGroup.EntityScope)
    d.Set("policy_group_id", PolicyGroup.PolicyGroupID)
    d.Set("external", PolicyGroup.External)
    d.Set("external_id", PolicyGroup.ExternalID)
    d.Set("type", PolicyGroup.Type)
    
    d.Set("id", PolicyGroup.Identifier())
    d.Set("parent_id", PolicyGroup.ParentID)
    d.Set("parent_type", PolicyGroup.ParentType)
    d.Set("owner", PolicyGroup.Owner)

    d.SetId(PolicyGroup.Identifier())
    
    return nil
}