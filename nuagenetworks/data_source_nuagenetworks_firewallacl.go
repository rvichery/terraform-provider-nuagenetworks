package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceFirewallAcl() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFirewallAclRead,
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
			"rule_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain"},
			},
		},
	}
}

func dataSourceFirewallAclRead(d *schema.ResourceData, m interface{}) error {
	filteredFirewallAcls := vspk.FirewallAclsList{}
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
		filteredFirewallAcls, err = parent.FirewallAcls(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredFirewallAcls, err = parent.FirewallAcls(fetchFilter)
		if err != nil {
			return err
		}
	}

	FirewallAcl := &vspk.FirewallAcl{}

	if len(filteredFirewallAcls) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredFirewallAcls) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	FirewallAcl = filteredFirewallAcls[0]

	d.Set("name", FirewallAcl.Name)
	d.Set("active", FirewallAcl.Active)
	d.Set("default_allow_ip", FirewallAcl.DefaultAllowIP)
	d.Set("default_allow_non_ip", FirewallAcl.DefaultAllowNonIP)
	d.Set("description", FirewallAcl.Description)
	d.Set("rule_ids", FirewallAcl.RuleIds)

	d.Set("id", FirewallAcl.Identifier())
	d.Set("parent_id", FirewallAcl.ParentID)
	d.Set("parent_type", FirewallAcl.ParentType)
	d.Set("owner", FirewallAcl.Owner)

	d.SetId(FirewallAcl.Identifier())

	return nil
}
