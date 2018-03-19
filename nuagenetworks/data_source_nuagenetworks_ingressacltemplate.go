package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceIngressACLTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIngressACLTemplateRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default_allow_ip": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default_allow_non_ip": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_address_spoof": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"priority_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assoc_acl_template_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_live_entity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_virtual_firewall_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_generate_priority": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l2_domain_template", "parent_l2_domain", "parent_domain_template"},
			},
			"parent_l2_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_l2_domain", "parent_domain_template"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_l2_domain_template", "parent_domain_template"},
			},
			"parent_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_l2_domain_template", "parent_l2_domain"},
			},
		},
	}
}

func dataSourceIngressACLTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredIngressACLTemplates := vspk.IngressACLTemplatesList{}
	err := &bambou.Error{}
	fetchFilter := &bambou.FetchingInfo{}

	filters, filtersOk := d.GetOk("filter")
	if filtersOk {
		fetchFilter = bambou.NewFetchingInfo()
		for _, v := range filters.(*schema.Set).List() {
			m := v.(map[string]interface{})
			if fetchFilter.Filter != "" {
				fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string), m["operator"].(string), m["value"].(string))
			} else {
				fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
			}

		}
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredIngressACLTemplates, err = parent.IngressACLTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		filteredIngressACLTemplates, err = parent.IngressACLTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredIngressACLTemplates, err = parent.IngressACLTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredIngressACLTemplates, err = parent.IngressACLTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredIngressACLTemplates, err = parent.IngressACLTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	IngressACLTemplate := &vspk.IngressACLTemplate{}

	if len(filteredIngressACLTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIngressACLTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	IngressACLTemplate = filteredIngressACLTemplates[0]

	d.Set("name", IngressACLTemplate.Name)
	d.Set("last_updated_by", IngressACLTemplate.LastUpdatedBy)
	d.Set("active", IngressACLTemplate.Active)
	d.Set("default_allow_ip", IngressACLTemplate.DefaultAllowIP)
	d.Set("default_allow_non_ip", IngressACLTemplate.DefaultAllowNonIP)
	d.Set("description", IngressACLTemplate.Description)
	d.Set("allow_address_spoof", IngressACLTemplate.AllowAddressSpoof)
	d.Set("entity_scope", IngressACLTemplate.EntityScope)
	d.Set("policy_state", IngressACLTemplate.PolicyState)
	d.Set("priority", IngressACLTemplate.Priority)
	d.Set("priority_type", IngressACLTemplate.PriorityType)
	d.Set("assoc_acl_template_id", IngressACLTemplate.AssocAclTemplateId)
	d.Set("associated_live_entity_id", IngressACLTemplate.AssociatedLiveEntityID)
	d.Set("associated_virtual_firewall_policy_id", IngressACLTemplate.AssociatedVirtualFirewallPolicyID)
	d.Set("auto_generate_priority", IngressACLTemplate.AutoGeneratePriority)
	d.Set("external_id", IngressACLTemplate.ExternalID)

	d.Set("id", IngressACLTemplate.Identifier())
	d.Set("parent_id", IngressACLTemplate.ParentID)
	d.Set("parent_type", IngressACLTemplate.ParentType)
	d.Set("owner", IngressACLTemplate.Owner)

	d.SetId(IngressACLTemplate.Identifier())

	return nil
}
