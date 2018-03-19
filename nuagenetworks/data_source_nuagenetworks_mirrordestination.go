package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceMirrorDestination() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMirrorDestinationRead,
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
			"service_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"destination_ip": {
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

func dataSourceMirrorDestinationRead(d *schema.ResourceData, m interface{}) error {
	filteredMirrorDestinations := vspk.MirrorDestinationsList{}
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
	filteredMirrorDestinations, err = parent.MirrorDestinations(fetchFilter)
	if err != nil {
		return err
	}

	MirrorDestination := &vspk.MirrorDestination{}

	if len(filteredMirrorDestinations) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredMirrorDestinations) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	MirrorDestination = filteredMirrorDestinations[0]

	d.Set("name", MirrorDestination.Name)
	d.Set("last_updated_by", MirrorDestination.LastUpdatedBy)
	d.Set("service_id", MirrorDestination.ServiceId)
	d.Set("destination_ip", MirrorDestination.DestinationIp)
	d.Set("entity_scope", MirrorDestination.EntityScope)
	d.Set("external_id", MirrorDestination.ExternalID)

	d.Set("id", MirrorDestination.Identifier())
	d.Set("parent_id", MirrorDestination.ParentID)
	d.Set("parent_type", MirrorDestination.ParentType)
	d.Set("owner", MirrorDestination.Owner)

	d.SetId(MirrorDestination.Identifier())

	return nil
}
