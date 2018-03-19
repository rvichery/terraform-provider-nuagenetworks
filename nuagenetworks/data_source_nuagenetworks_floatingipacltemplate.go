package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceFloatingIPACLTemplate() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceFloatingIPACLTemplateRead,
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
            "auto_generate_priority": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain_template"},
            },
            "parent_domain_template": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain"},
            },
        },
    }
}


func dataSourceFloatingIPACLTemplateRead(d *schema.ResourceData, m interface{}) error {
    filteredFloatingIPACLTemplates := vspk.FloatingIPACLTemplatesList{}
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
        filteredFloatingIPACLTemplates, err = parent.FloatingIPACLTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_domain_template"); ok {
        parent := &vspk.DomainTemplate{ID: attr.(string)}
        filteredFloatingIPACLTemplates, err = parent.FloatingIPACLTemplates(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredFloatingIPACLTemplates, err = parent.FloatingIPACLTemplates(fetchFilter)
        if err != nil {
            return err
        }
    }

    FloatingIPACLTemplate := &vspk.FloatingIPACLTemplate{}

    if len(filteredFloatingIPACLTemplates) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredFloatingIPACLTemplates) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    FloatingIPACLTemplate = filteredFloatingIPACLTemplates[0]

    d.Set("name", FloatingIPACLTemplate.Name)
    d.Set("last_updated_by", FloatingIPACLTemplate.LastUpdatedBy)
    d.Set("active", FloatingIPACLTemplate.Active)
    d.Set("default_allow_ip", FloatingIPACLTemplate.DefaultAllowIP)
    d.Set("default_allow_non_ip", FloatingIPACLTemplate.DefaultAllowNonIP)
    d.Set("description", FloatingIPACLTemplate.Description)
    d.Set("entity_scope", FloatingIPACLTemplate.EntityScope)
    d.Set("policy_state", FloatingIPACLTemplate.PolicyState)
    d.Set("priority", FloatingIPACLTemplate.Priority)
    d.Set("priority_type", FloatingIPACLTemplate.PriorityType)
    d.Set("associated_live_entity_id", FloatingIPACLTemplate.AssociatedLiveEntityID)
    d.Set("auto_generate_priority", FloatingIPACLTemplate.AutoGeneratePriority)
    d.Set("external_id", FloatingIPACLTemplate.ExternalID)
    
    d.Set("id", FloatingIPACLTemplate.Identifier())
    d.Set("parent_id", FloatingIPACLTemplate.ParentID)
    d.Set("parent_type", FloatingIPACLTemplate.ParentType)
    d.Set("owner", FloatingIPACLTemplate.Owner)

    d.SetId(FloatingIPACLTemplate.Identifier())
    
    return nil
}