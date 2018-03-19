package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceLocation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLocationRead,
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
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latitude": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_geocode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"time_zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locality": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"longitude": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"country": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ns_gateway": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceLocationRead(d *schema.ResourceData, m interface{}) error {
	filteredLocations := vspk.LocationsList{}
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
	parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
	filteredLocations, err = parent.Locations(fetchFilter)
	if err != nil {
		return err
	}

	Location := &vspk.Location{}

	if len(filteredLocations) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredLocations) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Location = filteredLocations[0]

	d.Set("last_updated_by", Location.LastUpdatedBy)
	d.Set("latitude", Location.Latitude)
	d.Set("address", Location.Address)
	d.Set("ignore_geocode", Location.IgnoreGeocode)
	d.Set("time_zone_id", Location.TimeZoneID)
	d.Set("entity_scope", Location.EntityScope)
	d.Set("locality", Location.Locality)
	d.Set("longitude", Location.Longitude)
	d.Set("country", Location.Country)
	d.Set("state", Location.State)
	d.Set("external_id", Location.ExternalID)

	d.Set("id", Location.Identifier())
	d.Set("parent_id", Location.ParentID)
	d.Set("parent_type", Location.ParentType)
	d.Set("owner", Location.Owner)

	d.SetId(Location.Identifier())

	return nil
}
