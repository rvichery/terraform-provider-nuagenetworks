package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVSC() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVSCRead,
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
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_state_change": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"peak_cpuusage": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"peak_memory_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"messages": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"disks": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"already_marked_for_unavailable": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"unavailable_timestamp": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsds": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_cpuusage": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"current_memory_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"average_cpuusage": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"average_memory_usage": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vsp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVSCRead(d *schema.ResourceData, m interface{}) error {
	filteredVSCs := vspk.VSCsList{}
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
	filteredVSCs, err = parent.VSCs(fetchFilter)
	if err != nil {
		return err
	}

	VSC := &vspk.VSC{}

	if len(filteredVSCs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVSCs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		VSC = filteredVSCs[0]
	}

	d.Set("name", VSC.Name)
	d.Set("management_ip", VSC.ManagementIP)
	d.Set("last_state_change", VSC.LastStateChange)
	d.Set("last_updated_by", VSC.LastUpdatedBy)
	d.Set("address", VSC.Address)
	d.Set("peak_cpuusage", VSC.PeakCPUUsage)
	d.Set("peak_memory_usage", VSC.PeakMemoryUsage)
	d.Set("description", VSC.Description)
	d.Set("messages", VSC.Messages)
	d.Set("disks", VSC.Disks)
	d.Set("already_marked_for_unavailable", VSC.AlreadyMarkedForUnavailable)
	d.Set("unavailable_timestamp", VSC.UnavailableTimestamp)
	d.Set("entity_scope", VSC.EntityScope)
	d.Set("location", VSC.Location)
	d.Set("product_version", VSC.ProductVersion)
	d.Set("vsds", VSC.Vsds)
	d.Set("status", VSC.Status)
	d.Set("current_cpuusage", VSC.CurrentCPUUsage)
	d.Set("current_memory_usage", VSC.CurrentMemoryUsage)
	d.Set("average_cpuusage", VSC.AverageCPUUsage)
	d.Set("average_memory_usage", VSC.AverageMemoryUsage)
	d.Set("external_id", VSC.ExternalID)

	d.Set("id", VSC.Identifier())
	d.Set("parent_id", VSC.ParentID)
	d.Set("parent_type", VSC.ParentType)
	d.Set("owner", VSC.Owner)

	d.SetId(VSC.Identifier())

	return nil
}
