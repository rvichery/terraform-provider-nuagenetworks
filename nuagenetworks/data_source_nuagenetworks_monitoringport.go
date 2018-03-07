package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceMonitoringPort() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMonitoringPortRead,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_state_change": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"access": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"resiliency_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"resilient": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uplink": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vrs": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_hsc", "parent_vsc"},
			},
			"parent_hsc": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vrs", "parent_vsc"},
			},
			"parent_vsc": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vrs", "parent_hsc"},
			},
		},
	}
}

func dataSourceMonitoringPortRead(d *schema.ResourceData, m interface{}) error {
	filteredMonitoringPorts := vspk.MonitoringPortsList{}
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
	if attr, ok := d.GetOk("parent_vrs"); ok {
		parent := &vspk.VRS{ID: attr.(string)}
		filteredMonitoringPorts, err = parent.MonitoringPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_hsc"); ok {
		parent := &vspk.HSC{ID: attr.(string)}
		filteredMonitoringPorts, err = parent.MonitoringPorts(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vsc"); ok {
		parent := &vspk.VSC{ID: attr.(string)}
		filteredMonitoringPorts, err = parent.MonitoringPorts(fetchFilter)
		if err != nil {
			return err
		}
	}

	MonitoringPort := &vspk.MonitoringPort{}

	if len(filteredMonitoringPorts) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredMonitoringPorts) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		MonitoringPort = filteredMonitoringPorts[0]
	}

	d.Set("name", MonitoringPort.Name)
	d.Set("last_state_change", MonitoringPort.LastStateChange)
	d.Set("access", MonitoringPort.Access)
	d.Set("description", MonitoringPort.Description)
	d.Set("resiliency_state", MonitoringPort.ResiliencyState)
	d.Set("resilient", MonitoringPort.Resilient)
	d.Set("entity_scope", MonitoringPort.EntityScope)
	d.Set("uplink", MonitoringPort.Uplink)
	d.Set("state", MonitoringPort.State)
	d.Set("external_id", MonitoringPort.ExternalID)

	d.Set("id", MonitoringPort.Identifier())
	d.Set("parent_id", MonitoringPort.ParentID)
	d.Set("parent_type", MonitoringPort.ParentType)
	d.Set("owner", MonitoringPort.Owner)

	d.SetId(MonitoringPort.Identifier())

	return nil
}
