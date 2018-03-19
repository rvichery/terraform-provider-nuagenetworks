package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVirtualFirewallPolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVirtualFirewallPolicyRead,
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
			"default_install_acl_implicit_rules": {
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
			"associated_egress_template_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ingress_template_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_live_entity_id": {
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

func dataSourceVirtualFirewallPolicyRead(d *schema.ResourceData, m interface{}) error {
	filteredVirtualFirewallPolicies := vspk.VirtualFirewallPoliciesList{}
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
		filteredVirtualFirewallPolicies, err = parent.VirtualFirewallPolicies(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		filteredVirtualFirewallPolicies, err = parent.VirtualFirewallPolicies(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredVirtualFirewallPolicies, err = parent.VirtualFirewallPolicies(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredVirtualFirewallPolicies, err = parent.VirtualFirewallPolicies(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredVirtualFirewallPolicies, err = parent.VirtualFirewallPolicies(fetchFilter)
		if err != nil {
			return err
		}
	}

	VirtualFirewallPolicy := &vspk.VirtualFirewallPolicy{}

	if len(filteredVirtualFirewallPolicies) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVirtualFirewallPolicies) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VirtualFirewallPolicy = filteredVirtualFirewallPolicies[0]

	d.Set("name", VirtualFirewallPolicy.Name)
	d.Set("last_updated_by", VirtualFirewallPolicy.LastUpdatedBy)
	d.Set("active", VirtualFirewallPolicy.Active)
	d.Set("default_allow_ip", VirtualFirewallPolicy.DefaultAllowIP)
	d.Set("default_allow_non_ip", VirtualFirewallPolicy.DefaultAllowNonIP)
	d.Set("default_install_acl_implicit_rules", VirtualFirewallPolicy.DefaultInstallACLImplicitRules)
	d.Set("description", VirtualFirewallPolicy.Description)
	d.Set("allow_address_spoof", VirtualFirewallPolicy.AllowAddressSpoof)
	d.Set("entity_scope", VirtualFirewallPolicy.EntityScope)
	d.Set("policy_state", VirtualFirewallPolicy.PolicyState)
	d.Set("priority", VirtualFirewallPolicy.Priority)
	d.Set("priority_type", VirtualFirewallPolicy.PriorityType)
	d.Set("associated_egress_template_id", VirtualFirewallPolicy.AssociatedEgressTemplateID)
	d.Set("associated_ingress_template_id", VirtualFirewallPolicy.AssociatedIngressTemplateID)
	d.Set("associated_live_entity_id", VirtualFirewallPolicy.AssociatedLiveEntityID)
	d.Set("auto_generate_priority", VirtualFirewallPolicy.AutoGeneratePriority)
	d.Set("external_id", VirtualFirewallPolicy.ExternalID)

	d.Set("id", VirtualFirewallPolicy.Identifier())
	d.Set("parent_id", VirtualFirewallPolicy.ParentID)
	d.Set("parent_type", VirtualFirewallPolicy.ParentType)
	d.Set("owner", VirtualFirewallPolicy.Owner)

	d.SetId(VirtualFirewallPolicy.Identifier())

	return nil
}
