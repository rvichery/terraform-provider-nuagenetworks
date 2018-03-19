package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceUplinkRD() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceUplinkRDRead,
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
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"route_distinguisher": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uplink_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l2_domain"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain"},
			},
		},
	}
}

func dataSourceUplinkRDRead(d *schema.ResourceData, m interface{}) error {
	filteredUplinkRDs := vspk.UplinkRDsList{}
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
		filteredUplinkRDs, err = parent.UplinkRDs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredUplinkRDs, err = parent.UplinkRDs(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredUplinkRDs, err = parent.UplinkRDs(fetchFilter)
		if err != nil {
			return err
		}
	}

	UplinkRD := &vspk.UplinkRD{}

	if len(filteredUplinkRDs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredUplinkRDs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	UplinkRD = filteredUplinkRDs[0]

	d.Set("last_updated_by", UplinkRD.LastUpdatedBy)
	d.Set("entity_scope", UplinkRD.EntityScope)
	d.Set("route_distinguisher", UplinkRD.RouteDistinguisher)
	d.Set("uplink_type", UplinkRD.UplinkType)
	d.Set("external_id", UplinkRD.ExternalID)

	d.Set("id", UplinkRD.Identifier())
	d.Set("parent_id", UplinkRD.ParentID)
	d.Set("parent_type", UplinkRD.ParentType)
	d.Set("owner", UplinkRD.Owner)

	d.SetId(UplinkRD.Identifier())

	return nil
}
