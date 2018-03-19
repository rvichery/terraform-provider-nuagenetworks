package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceVSD() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVSDRead,
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
			"url": {
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
			"peer_addresses": {
				Type:     schema.TypeString,
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
			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_version": {
				Type:     schema.TypeString,
				Computed: true,
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
			"parent_vsp": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVSDRead(d *schema.ResourceData, m interface{}) error {
	filteredVSDs := vspk.VSDsList{}
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
	filteredVSDs, err = parent.VSDs(fetchFilter)
	if err != nil {
		return err
	}

	VSD := &vspk.VSD{}

	if len(filteredVSDs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVSDs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VSD = filteredVSDs[0]

	d.Set("url", VSD.URL)
	d.Set("name", VSD.Name)
	d.Set("management_ip", VSD.ManagementIP)
	d.Set("last_state_change", VSD.LastStateChange)
	d.Set("last_updated_by", VSD.LastUpdatedBy)
	d.Set("address", VSD.Address)
	d.Set("peak_cpuusage", VSD.PeakCPUUsage)
	d.Set("peak_memory_usage", VSD.PeakMemoryUsage)
	d.Set("peer_addresses", VSD.PeerAddresses)
	d.Set("description", VSD.Description)
	d.Set("messages", VSD.Messages)
	d.Set("disks", VSD.Disks)
	d.Set("already_marked_for_unavailable", VSD.AlreadyMarkedForUnavailable)
	d.Set("unavailable_timestamp", VSD.UnavailableTimestamp)
	d.Set("entity_scope", VSD.EntityScope)
	d.Set("location", VSD.Location)
	d.Set("mode", VSD.Mode)
	d.Set("product_version", VSD.ProductVersion)
	d.Set("status", VSD.Status)
	d.Set("current_cpuusage", VSD.CurrentCPUUsage)
	d.Set("current_memory_usage", VSD.CurrentMemoryUsage)
	d.Set("average_cpuusage", VSD.AverageCPUUsage)
	d.Set("average_memory_usage", VSD.AverageMemoryUsage)
	d.Set("external_id", VSD.ExternalID)

	d.Set("id", VSD.Identifier())
	d.Set("parent_id", VSD.ParentID)
	d.Set("parent_type", VSD.ParentType)
	d.Set("owner", VSD.Owner)

	d.SetId(VSD.Identifier())

	return nil
}
