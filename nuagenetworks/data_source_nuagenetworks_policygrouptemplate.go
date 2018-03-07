package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourcePolicyGroupTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePolicyGroupTemplateRead,
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
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
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
			"parent_l2_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain_template"},
			},
			"parent_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l2_domain_template"},
			},
		},
	}
}

func dataSourcePolicyGroupTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredPolicyGroupTemplates := vspk.PolicyGroupTemplatesList{}
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
	if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		filteredPolicyGroupTemplates, err = parent.PolicyGroupTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredPolicyGroupTemplates, err = parent.PolicyGroupTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	PolicyGroupTemplate := &vspk.PolicyGroupTemplate{}

	if len(filteredPolicyGroupTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPolicyGroupTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		PolicyGroupTemplate = filteredPolicyGroupTemplates[0]
	}

	d.Set("evpn_community_tag", PolicyGroupTemplate.EVPNCommunityTag)
	d.Set("name", PolicyGroupTemplate.Name)
	d.Set("last_updated_by", PolicyGroupTemplate.LastUpdatedBy)
	d.Set("description", PolicyGroupTemplate.Description)
	d.Set("entity_scope", PolicyGroupTemplate.EntityScope)
	d.Set("external", PolicyGroupTemplate.External)
	d.Set("external_id", PolicyGroupTemplate.ExternalID)
	d.Set("type", PolicyGroupTemplate.Type)

	d.Set("id", PolicyGroupTemplate.Identifier())
	d.Set("parent_id", PolicyGroupTemplate.ParentID)
	d.Set("parent_type", PolicyGroupTemplate.ParentType)
	d.Set("owner", PolicyGroupTemplate.Owner)

	d.SetId(PolicyGroupTemplate.Identifier())

	return nil
}
