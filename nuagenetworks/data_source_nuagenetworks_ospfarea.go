package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceOSPFArea() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOSPFAreaRead,
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
			"redistribute_external_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"default_metric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"default_originate_option": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aggregate_area_range": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"aggregate_area_range_nssa": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"area_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"area_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"summaries_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"suppress_area_range": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"suppress_area_range_nssa": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ospf_instance": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceOSPFAreaRead(d *schema.ResourceData, m interface{}) error {
	filteredOSPFAreas := vspk.OSPFAreasList{}
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
	parent := &vspk.OSPFInstance{ID: d.Get("parent_ospf_instance").(string)}
	filteredOSPFAreas, err = parent.OSPFAreas(fetchFilter)
	if err != nil {
		return err
	}

	OSPFArea := &vspk.OSPFArea{}

	if len(filteredOSPFAreas) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredOSPFAreas) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		OSPFArea = filteredOSPFAreas[0]
	}

	d.Set("last_updated_by", OSPFArea.LastUpdatedBy)
	d.Set("redistribute_external_enabled", OSPFArea.RedistributeExternalEnabled)
	d.Set("default_metric", OSPFArea.DefaultMetric)
	d.Set("default_originate_option", OSPFArea.DefaultOriginateOption)
	d.Set("description", OSPFArea.Description)
	d.Set("aggregate_area_range", OSPFArea.AggregateAreaRange)
	d.Set("aggregate_area_range_nssa", OSPFArea.AggregateAreaRangeNSSA)
	d.Set("entity_scope", OSPFArea.EntityScope)
	d.Set("area_id", OSPFArea.AreaID)
	d.Set("area_type", OSPFArea.AreaType)
	d.Set("summaries_enabled", OSPFArea.SummariesEnabled)
	d.Set("suppress_area_range", OSPFArea.SuppressAreaRange)
	d.Set("suppress_area_range_nssa", OSPFArea.SuppressAreaRangeNSSA)
	d.Set("external_id", OSPFArea.ExternalID)

	d.Set("id", OSPFArea.Identifier())
	d.Set("parent_id", OSPFArea.ParentID)
	d.Set("parent_type", OSPFArea.ParentType)
	d.Set("owner", OSPFArea.Owner)

	d.SetId(OSPFArea.Identifier())

	return nil
}
