package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourcePerformanceMonitor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePerformanceMonitorRead,
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
			"payload_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"read_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"service_class": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_target_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hold_down_timer": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"probe_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"number_of_packets": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ike_gateway_connection": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_ike_gateway_connection"},
			},
		},
	}
}

func dataSourcePerformanceMonitorRead(d *schema.ResourceData, m interface{}) error {
	filteredPerformanceMonitors := vspk.PerformanceMonitorsList{}
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
	if attr, ok := d.GetOk("parent_ike_gateway_connection"); ok {
		parent := &vspk.IKEGatewayConnection{ID: attr.(string)}
		filteredPerformanceMonitors, err = parent.PerformanceMonitors(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredPerformanceMonitors, err = parent.PerformanceMonitors(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredPerformanceMonitors, err = parent.PerformanceMonitors(fetchFilter)
		if err != nil {
			return err
		}
	}

	PerformanceMonitor := &vspk.PerformanceMonitor{}

	if len(filteredPerformanceMonitors) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPerformanceMonitors) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	PerformanceMonitor = filteredPerformanceMonitors[0]

	d.Set("name", PerformanceMonitor.Name)
	d.Set("last_updated_by", PerformanceMonitor.LastUpdatedBy)
	d.Set("payload_size", PerformanceMonitor.PayloadSize)
	d.Set("read_only", PerformanceMonitor.ReadOnly)
	d.Set("service_class", PerformanceMonitor.ServiceClass)
	d.Set("description", PerformanceMonitor.Description)
	d.Set("destination_target_list", PerformanceMonitor.DestinationTargetList)
	d.Set("timeout", PerformanceMonitor.Timeout)
	d.Set("interval", PerformanceMonitor.Interval)
	d.Set("entity_scope", PerformanceMonitor.EntityScope)
	d.Set("hold_down_timer", PerformanceMonitor.HoldDownTimer)
	d.Set("probe_type", PerformanceMonitor.ProbeType)
	d.Set("number_of_packets", PerformanceMonitor.NumberOfPackets)
	d.Set("external_id", PerformanceMonitor.ExternalID)

	d.Set("id", PerformanceMonitor.Identifier())
	d.Set("parent_id", PerformanceMonitor.ParentID)
	d.Set("parent_type", PerformanceMonitor.ParentType)
	d.Set("owner", PerformanceMonitor.Owner)

	d.SetId(PerformanceMonitor.Identifier())

	return nil
}
