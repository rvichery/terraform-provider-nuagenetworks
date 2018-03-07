package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceIPReservation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPReservationRead,
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
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_allocation_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parent_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceIPReservationRead(d *schema.ResourceData, m interface{}) error {
	filteredIPReservations := vspk.IPReservationsList{}
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
	parent := &vspk.Subnet{ID: d.Get("parent_subnet").(string)}
	filteredIPReservations, err = parent.IPReservations(fetchFilter)
	if err != nil {
		return err
	}

	IPReservation := &vspk.IPReservation{}

	if len(filteredIPReservations) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIPReservations) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		IPReservation = filteredIPReservations[0]
	}

	d.Set("mac", IPReservation.MAC)
	d.Set("ip_address", IPReservation.IPAddress)
	d.Set("last_updated_by", IPReservation.LastUpdatedBy)
	d.Set("entity_scope", IPReservation.EntityScope)
	d.Set("external_id", IPReservation.ExternalID)
	d.Set("dynamic_allocation_enabled", IPReservation.DynamicAllocationEnabled)

	d.Set("id", IPReservation.Identifier())
	d.Set("parent_id", IPReservation.ParentID)
	d.Set("parent_type", IPReservation.ParentType)
	d.Set("owner", IPReservation.Owner)

	d.SetId(IPReservation.Identifier())

	return nil
}
