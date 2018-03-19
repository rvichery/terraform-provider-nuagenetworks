package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourcePATMapper() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePATMapperRead,
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
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePATMapperRead(d *schema.ResourceData, m interface{}) error {
	filteredPATMappers := vspk.PATMappersList{}
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
	parent := m.(*vspk.Me)
	filteredPATMappers, err = parent.PATMappers(fetchFilter)
	if err != nil {
		return err
	}

	PATMapper := &vspk.PATMapper{}

	if len(filteredPATMappers) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPATMappers) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	PATMapper = filteredPATMappers[0]

	d.Set("name", PATMapper.Name)
	d.Set("last_updated_by", PATMapper.LastUpdatedBy)
	d.Set("description", PATMapper.Description)
	d.Set("entity_scope", PATMapper.EntityScope)
	d.Set("external_id", PATMapper.ExternalID)

	d.Set("id", PATMapper.Identifier())
	d.Set("parent_id", PATMapper.ParentID)
	d.Set("parent_type", PATMapper.ParentType)
	d.Set("owner", PATMapper.Owner)

	d.SetId(PATMapper.Identifier())

	return nil
}
