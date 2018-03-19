package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceEgressACLTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceEgressACLTemplateRead,
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
            "active": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "default_allow_ip": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "default_allow_non_ip": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "default_install_acl_implicit_rules": &schema.Schema{
                Type:     schema.TypeBool,
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
            "policy_state": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "priority_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_live_entity_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_virtual_firewall_policy_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "auto_generate_priority": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain_template", "parent_l2_domain", "parent_domain_template"},
            },
            "parent_l2_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_domain_template"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain_template", "parent_domain_template"},
            },
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain", "parent_l2_domain_template", "parent_l2_domain"},
            },
        },
    }
}


func dataSourceEgressACLTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredEgressACLTemplates := vspk.EgressACLTemplatesList{}
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
        filteredEgressACLTemplates, err = parent.EgressACLTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
        parent := &vspk.L2DomainTemplate{ID: attr.(string)}
        filteredEgressACLTemplates, err = parent.EgressACLTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        filteredEgressACLTemplates, err = parent.EgressACLTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        filteredEgressACLTemplates, err = parent.EgressACLTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredEgressACLTemplates, err = parent.EgressACLTemplates(fetchFilter)
        if err != nil {
            return err
        }
    }

    EgressACLTemplate := &vspk.EgressACLTemplate{}

    if len(filteredEgressACLTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredEgressACLTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    EgressACLTemplate = filteredEgressACLTemplates[0]

    d.Set("name", EgressACLTemplate.Name)
    d.Set("last_updated_by", EgressACLTemplate.LastUpdatedBy)
    d.Set("active", EgressACLTemplate.Active)
    d.Set("default_allow_ip", EgressACLTemplate.DefaultAllowIP)
    d.Set("default_allow_non_ip", EgressACLTemplate.DefaultAllowNonIP)
    d.Set("default_install_acl_implicit_rules", EgressACLTemplate.DefaultInstallACLImplicitRules)
    d.Set("description", EgressACLTemplate.Description)
    d.Set("entity_scope", EgressACLTemplate.EntityScope)
    d.Set("policy_state", EgressACLTemplate.PolicyState)
    d.Set("priority", EgressACLTemplate.Priority)
    d.Set("priority_type", EgressACLTemplate.PriorityType)
    d.Set("associated_live_entity_id", EgressACLTemplate.AssociatedLiveEntityID)
    d.Set("associated_virtual_firewall_policy_id", EgressACLTemplate.AssociatedVirtualFirewallPolicyID)
    d.Set("auto_generate_priority", EgressACLTemplate.AutoGeneratePriority)
    d.Set("external_id", EgressACLTemplate.ExternalID)
    
    d.Set("id", EgressACLTemplate.Identifier())
    d.Set("parent_id", EgressACLTemplate.ParentID)
    d.Set("parent_type", EgressACLTemplate.ParentType)
    d.Set("owner", EgressACLTemplate.Owner)

    d.SetId(EgressACLTemplate.Identifier())
    
    return nil
}