package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceDomainFIPAclTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDomainFIPAclTemplateRead,
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
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entries": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain_template"},
			},
			"parent_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain"},
			},
		},
	}
}

func dataSourceDomainFIPAclTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredDomainFIPAclTemplates := vspk.DomainFIPAclTemplatesList{}
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
		filteredDomainFIPAclTemplates, err = parent.DomainFIPAclTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredDomainFIPAclTemplates, err = parent.DomainFIPAclTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredDomainFIPAclTemplates, err = parent.DomainFIPAclTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	DomainFIPAclTemplate := &vspk.DomainFIPAclTemplate{}

	if len(filteredDomainFIPAclTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDomainFIPAclTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	DomainFIPAclTemplate = filteredDomainFIPAclTemplates[0]

	d.Set("name", DomainFIPAclTemplate.Name)
	d.Set("last_updated_by", DomainFIPAclTemplate.LastUpdatedBy)
	d.Set("active", DomainFIPAclTemplate.Active)
	d.Set("default_allow_ip", DomainFIPAclTemplate.DefaultAllowIP)
	d.Set("default_allow_non_ip", DomainFIPAclTemplate.DefaultAllowNonIP)
	d.Set("description", DomainFIPAclTemplate.Description)
	d.Set("entity_scope", DomainFIPAclTemplate.EntityScope)
	d.Set("entries", DomainFIPAclTemplate.Entries)
	d.Set("policy_state", DomainFIPAclTemplate.PolicyState)
	d.Set("priority", DomainFIPAclTemplate.Priority)
	d.Set("priority_type", DomainFIPAclTemplate.PriorityType)
	d.Set("associated_live_entity_id", DomainFIPAclTemplate.AssociatedLiveEntityID)
	d.Set("external_id", DomainFIPAclTemplate.ExternalID)

	d.Set("id", DomainFIPAclTemplate.Identifier())
	d.Set("parent_id", DomainFIPAclTemplate.ParentID)
	d.Set("parent_type", DomainFIPAclTemplate.ParentType)
	d.Set("owner", DomainFIPAclTemplate.Owner)

	d.SetId(DomainFIPAclTemplate.Identifier())

	return nil
}
