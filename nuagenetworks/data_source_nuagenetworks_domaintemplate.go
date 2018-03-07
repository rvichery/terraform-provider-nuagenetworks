package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDomainTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDomainTemplateRead,
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
			"dpi": &schema.Schema{
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
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_change_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_bgp_profile_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_multicast_channel_map_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_pat_mapper_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain"},
			},
		},
	}
}

func dataSourceDomainTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredDomainTemplates := vspk.DomainTemplatesList{}
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
		filteredDomainTemplates, err = parent.DomainTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredDomainTemplates, err = parent.DomainTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	DomainTemplate := &vspk.DomainTemplate{}

	if len(filteredDomainTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDomainTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		DomainTemplate = filteredDomainTemplates[0]
	}

	d.Set("dpi", DomainTemplate.DPI)
	d.Set("name", DomainTemplate.Name)
	d.Set("last_updated_by", DomainTemplate.LastUpdatedBy)
	d.Set("description", DomainTemplate.Description)
	d.Set("encryption", DomainTemplate.Encryption)
	d.Set("entity_scope", DomainTemplate.EntityScope)
	d.Set("policy_change_status", DomainTemplate.PolicyChangeStatus)
	d.Set("associated_bgp_profile_id", DomainTemplate.AssociatedBGPProfileID)
	d.Set("associated_multicast_channel_map_id", DomainTemplate.AssociatedMulticastChannelMapID)
	d.Set("associated_pat_mapper_id", DomainTemplate.AssociatedPATMapperID)
	d.Set("multicast", DomainTemplate.Multicast)
	d.Set("external_id", DomainTemplate.ExternalID)

	d.Set("id", DomainTemplate.Identifier())
	d.Set("parent_id", DomainTemplate.ParentID)
	d.Set("parent_type", DomainTemplate.ParentType)
	d.Set("owner", DomainTemplate.Owner)

	d.SetId(DomainTemplate.Identifier())

	return nil
}
