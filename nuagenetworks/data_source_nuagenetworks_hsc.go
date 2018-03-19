package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceHSC() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHSCRead,
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
			"management_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_state_change": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"peak_cpuusage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"peak_memory_usage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"messages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"disks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"already_marked_for_unavailable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"unavailable_timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"model": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsds": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_cpuusage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"current_memory_usage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"average_cpuusage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"average_memory_usage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vsp": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceHSCRead(d *schema.ResourceData, m interface{}) error {
	filteredHSCs := vspk.HSCsList{}
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
	parent := &vspk.VSP{ID: d.Get("parent_vsp").(string)}
	filteredHSCs, err = parent.HSCs(fetchFilter)
	if err != nil {
		return err
	}

	HSC := &vspk.HSC{}

	if len(filteredHSCs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredHSCs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	HSC = filteredHSCs[0]

	d.Set("name", HSC.Name)
	d.Set("management_ip", HSC.ManagementIP)
	d.Set("last_state_change", HSC.LastStateChange)
	d.Set("last_updated_by", HSC.LastUpdatedBy)
	d.Set("address", HSC.Address)
	d.Set("peak_cpuusage", HSC.PeakCPUUsage)
	d.Set("peak_memory_usage", HSC.PeakMemoryUsage)
	d.Set("description", HSC.Description)
	d.Set("messages", HSC.Messages)
	d.Set("disks", HSC.Disks)
	d.Set("already_marked_for_unavailable", HSC.AlreadyMarkedForUnavailable)
	d.Set("unavailable_timestamp", HSC.UnavailableTimestamp)
	d.Set("entity_scope", HSC.EntityScope)
	d.Set("location", HSC.Location)
	d.Set("model", HSC.Model)
	d.Set("product_version", HSC.ProductVersion)
	d.Set("vsds", HSC.Vsds)
	d.Set("status", HSC.Status)
	d.Set("current_cpuusage", HSC.CurrentCPUUsage)
	d.Set("current_memory_usage", HSC.CurrentMemoryUsage)
	d.Set("average_cpuusage", HSC.AverageCPUUsage)
	d.Set("average_memory_usage", HSC.AverageMemoryUsage)
	d.Set("external_id", HSC.ExternalID)
	d.Set("type", HSC.Type)

	d.Set("id", HSC.Identifier())
	d.Set("parent_id", HSC.ParentID)
	d.Set("parent_type", HSC.ParentType)
	d.Set("owner", HSC.Owner)

	d.SetId(HSC.Identifier())

	return nil
}
