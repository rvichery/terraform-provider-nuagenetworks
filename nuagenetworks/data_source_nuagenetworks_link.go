package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceLink() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLinkRead,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"acceptance_criteria": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_destination_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_destination_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_destination_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_source_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_source_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_source_type": &schema.Schema{
				Type:     schema.TypeString,
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
			"parent_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceLinkRead(d *schema.ResourceData, m interface{}) error {
	filteredLinks := vspk.LinksList{}
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
	parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
	filteredLinks, err = parent.Links(fetchFilter)
	if err != nil {
		return err
	}

	Link := &vspk.Link{}

	if len(filteredLinks) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredLinks) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		Link = filteredLinks[0]
	}

	d.Set("last_updated_by", Link.LastUpdatedBy)
	d.Set("acceptance_criteria", Link.AcceptanceCriteria)
	d.Set("read_only", Link.ReadOnly)
	d.Set("entity_scope", Link.EntityScope)
	d.Set("associated_destination_id", Link.AssociatedDestinationID)
	d.Set("associated_destination_name", Link.AssociatedDestinationName)
	d.Set("associated_destination_type", Link.AssociatedDestinationType)
	d.Set("associated_source_id", Link.AssociatedSourceID)
	d.Set("associated_source_name", Link.AssociatedSourceName)
	d.Set("associated_source_type", Link.AssociatedSourceType)
	d.Set("external_id", Link.ExternalID)
	d.Set("type", Link.Type)

	d.Set("id", Link.Identifier())
	d.Set("parent_id", Link.ParentID)
	d.Set("parent_type", Link.ParentType)
	d.Set("owner", Link.Owner)

	d.SetId(Link.Identifier())

	return nil
}
