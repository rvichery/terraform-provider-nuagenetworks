package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceRedirectionTargetTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceRedirectionTargetTemplateRead,
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
			"redundancy_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_point_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"trigger_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_l2_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain_template"},
			},
			"parent_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l2_domain_template"},
			},
		},
	}
}

func dataSourceRedirectionTargetTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredRedirectionTargetTemplates := vspk.RedirectionTargetTemplatesList{}
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
		filteredRedirectionTargetTemplates, err = parent.RedirectionTargetTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredRedirectionTargetTemplates, err = parent.RedirectionTargetTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	RedirectionTargetTemplate := &vspk.RedirectionTargetTemplate{}

	if len(filteredRedirectionTargetTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredRedirectionTargetTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	RedirectionTargetTemplate = filteredRedirectionTargetTemplates[0]

	d.Set("name", RedirectionTargetTemplate.Name)
	d.Set("last_updated_by", RedirectionTargetTemplate.LastUpdatedBy)
	d.Set("redundancy_enabled", RedirectionTargetTemplate.RedundancyEnabled)
	d.Set("description", RedirectionTargetTemplate.Description)
	d.Set("end_point_type", RedirectionTargetTemplate.EndPointType)
	d.Set("entity_scope", RedirectionTargetTemplate.EntityScope)
	d.Set("trigger_type", RedirectionTargetTemplate.TriggerType)
	d.Set("external_id", RedirectionTargetTemplate.ExternalID)

	d.Set("id", RedirectionTargetTemplate.Identifier())
	d.Set("parent_id", RedirectionTargetTemplate.ParentID)
	d.Set("parent_type", RedirectionTargetTemplate.ParentType)
	d.Set("owner", RedirectionTargetTemplate.Owner)

	d.SetId(RedirectionTargetTemplate.Identifier())

	return nil
}
