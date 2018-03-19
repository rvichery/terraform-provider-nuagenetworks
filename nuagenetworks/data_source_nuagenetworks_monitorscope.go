package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceMonitorscope() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMonitorscopeRead,
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
			"read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"destination_nsgs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"allow_all_destination_nsgs": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_all_source_nsgs": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"source_nsgs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"parent_ns_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_application", "parent_network_performance_measurement"},
			},
			"parent_application": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ns_gateway", "parent_network_performance_measurement"},
			},
			"parent_network_performance_measurement": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ns_gateway", "parent_application"},
			},
		},
	}
}

func dataSourceMonitorscopeRead(d *schema.ResourceData, m interface{}) error {
	filteredMonitorscopes := vspk.MonitorscopesList{}
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
	if attr, ok := d.GetOk("parent_ns_gateway"); ok {
		parent := &vspk.NSGateway{ID: attr.(string)}
		filteredMonitorscopes, err = parent.Monitorscopes(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_application"); ok {
		parent := &vspk.Application{ID: attr.(string)}
		filteredMonitorscopes, err = parent.Monitorscopes(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_network_performance_measurement"); ok {
		parent := &vspk.NetworkPerformanceMeasurement{ID: attr.(string)}
		filteredMonitorscopes, err = parent.Monitorscopes(fetchFilter)
		if err != nil {
			return err
		}
	}

	Monitorscope := &vspk.Monitorscope{}

	if len(filteredMonitorscopes) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredMonitorscopes) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Monitorscope = filteredMonitorscopes[0]

	d.Set("name", Monitorscope.Name)
	d.Set("read_only", Monitorscope.ReadOnly)
	d.Set("destination_nsgs", Monitorscope.DestinationNSGs)
	d.Set("allow_all_destination_nsgs", Monitorscope.AllowAllDestinationNSGs)
	d.Set("allow_all_source_nsgs", Monitorscope.AllowAllSourceNSGs)
	d.Set("source_nsgs", Monitorscope.SourceNSGs)

	d.Set("id", Monitorscope.Identifier())
	d.Set("parent_id", Monitorscope.ParentID)
	d.Set("parent_type", Monitorscope.ParentType)
	d.Set("owner", Monitorscope.Owner)

	d.SetId(Monitorscope.Identifier())

	return nil
}
