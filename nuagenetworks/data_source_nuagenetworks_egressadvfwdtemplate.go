package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceEgressAdvFwdTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEgressAdvFwdTemplateRead,
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
			"description": {
				Type:     schema.TypeString,
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

func dataSourceEgressAdvFwdTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredEgressAdvFwdTemplates := vspk.EgressAdvFwdTemplatesList{}
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
		filteredEgressAdvFwdTemplates, err = parent.EgressAdvFwdTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		filteredEgressAdvFwdTemplates, err = parent.EgressAdvFwdTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredEgressAdvFwdTemplates, err = parent.EgressAdvFwdTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredEgressAdvFwdTemplates, err = parent.EgressAdvFwdTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	EgressAdvFwdTemplate := &vspk.EgressAdvFwdTemplate{}

	if len(filteredEgressAdvFwdTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredEgressAdvFwdTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	EgressAdvFwdTemplate = filteredEgressAdvFwdTemplates[0]

	d.Set("name", EgressAdvFwdTemplate.Name)
	d.Set("last_updated_by", EgressAdvFwdTemplate.LastUpdatedBy)
	d.Set("active", EgressAdvFwdTemplate.Active)
	d.Set("description", EgressAdvFwdTemplate.Description)
	d.Set("entity_scope", EgressAdvFwdTemplate.EntityScope)
	d.Set("policy_state", EgressAdvFwdTemplate.PolicyState)
	d.Set("priority", EgressAdvFwdTemplate.Priority)
	d.Set("priority_type", EgressAdvFwdTemplate.PriorityType)
	d.Set("associated_live_entity_id", EgressAdvFwdTemplate.AssociatedLiveEntityID)
	d.Set("auto_generate_priority", EgressAdvFwdTemplate.AutoGeneratePriority)
	d.Set("external_id", EgressAdvFwdTemplate.ExternalID)

	d.Set("id", EgressAdvFwdTemplate.Identifier())
	d.Set("parent_id", EgressAdvFwdTemplate.ParentID)
	d.Set("parent_type", EgressAdvFwdTemplate.ParentType)
	d.Set("owner", EgressAdvFwdTemplate.Owner)

	d.SetId(EgressAdvFwdTemplate.Identifier())

	return nil
}
