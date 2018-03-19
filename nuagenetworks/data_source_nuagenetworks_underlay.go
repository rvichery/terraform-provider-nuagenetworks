package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceUnderlay() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUnderlayRead,
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
			"underlay_id": {
				Type:     schema.TypeFloat,
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
			"parent_uplink_connection": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceUnderlayRead(d *schema.ResourceData, m interface{}) error {
	filteredUnderlays := vspk.UnderlaysList{}
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
	if attr, ok := d.GetOk("parent_uplink_connection"); ok {
		parent := &vspk.UplinkConnection{ID: attr.(string)}
		filteredUnderlays, err = parent.Underlays(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredUnderlays, err = parent.Underlays(fetchFilter)
		if err != nil {
			return err
		}
	}

	Underlay := &vspk.Underlay{}

	if len(filteredUnderlays) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredUnderlays) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Underlay = filteredUnderlays[0]

	d.Set("name", Underlay.Name)
	d.Set("last_updated_by", Underlay.LastUpdatedBy)
	d.Set("description", Underlay.Description)
	d.Set("underlay_id", Underlay.UnderlayId)
	d.Set("entity_scope", Underlay.EntityScope)
	d.Set("external_id", Underlay.ExternalID)

	d.Set("id", Underlay.Identifier())
	d.Set("parent_id", Underlay.ParentID)
	d.Set("parent_type", Underlay.ParentType)
	d.Set("owner", Underlay.Owner)

	d.SetId(Underlay.Identifier())

	return nil
}
