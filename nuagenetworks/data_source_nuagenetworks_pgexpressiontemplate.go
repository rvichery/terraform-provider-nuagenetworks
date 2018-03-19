package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourcePGExpressionTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePGExpressionTemplateRead,
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
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expression": {
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

func dataSourcePGExpressionTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredPGExpressionTemplates := vspk.PGExpressionTemplatesList{}
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
		filteredPGExpressionTemplates, err = parent.PGExpressionTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredPGExpressionTemplates, err = parent.PGExpressionTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	PGExpressionTemplate := &vspk.PGExpressionTemplate{}

	if len(filteredPGExpressionTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPGExpressionTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	PGExpressionTemplate = filteredPGExpressionTemplates[0]

	d.Set("name", PGExpressionTemplate.Name)
	d.Set("last_updated_by", PGExpressionTemplate.LastUpdatedBy)
	d.Set("description", PGExpressionTemplate.Description)
	d.Set("entity_scope", PGExpressionTemplate.EntityScope)
	d.Set("expression", PGExpressionTemplate.Expression)
	d.Set("external_id", PGExpressionTemplate.ExternalID)

	d.Set("id", PGExpressionTemplate.Identifier())
	d.Set("parent_id", PGExpressionTemplate.ParentID)
	d.Set("parent_type", PGExpressionTemplate.ParentType)
	d.Set("owner", PGExpressionTemplate.Owner)

	d.SetId(PGExpressionTemplate.Identifier())

	return nil
}
