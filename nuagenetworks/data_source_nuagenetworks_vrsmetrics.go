package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVRSMetrics() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVRSMetricsRead,
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
			"al_ubr0_status": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"cpu_utilization": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"vrs_process": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vrsvsc_status": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"re_deploy": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"receiving_metrics": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"memory_utilization": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"jesxmon_process": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_disk_partition_utilization": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"root_disk_partition_utilization": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"applied_metrics_push_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_vcenter_hypervisor_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"current_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vcenter_hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVRSMetricsRead(d *schema.ResourceData, m interface{}) error {
	filteredVRSMetrics := vspk.VRSMetricsList{}
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
	parent := &vspk.VCenterHypervisor{ID: d.Get("parent_vcenter_hypervisor").(string)}
	filteredVRSMetrics, err = parent.VRSMetrics(fetchFilter)
	if err != nil {
		return err
	}

	VRSMetrics := &vspk.VRSMetrics{}

	if len(filteredVRSMetrics) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVRSMetrics) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		VRSMetrics = filteredVRSMetrics[0]
	}

	d.Set("al_ubr0_status", VRSMetrics.ALUbr0Status)
	d.Set("cpu_utilization", VRSMetrics.CPUUtilization)
	d.Set("vrs_process", VRSMetrics.VRSProcess)
	d.Set("vrsvsc_status", VRSMetrics.VRSVSCStatus)
	d.Set("last_updated_by", VRSMetrics.LastUpdatedBy)
	d.Set("re_deploy", VRSMetrics.ReDeploy)
	d.Set("receiving_metrics", VRSMetrics.ReceivingMetrics)
	d.Set("memory_utilization", VRSMetrics.MemoryUtilization)
	d.Set("jesxmon_process", VRSMetrics.JesxmonProcess)
	d.Set("entity_scope", VRSMetrics.EntityScope)
	d.Set("log_disk_partition_utilization", VRSMetrics.LogDiskPartitionUtilization)
	d.Set("root_disk_partition_utilization", VRSMetrics.RootDiskPartitionUtilization)
	d.Set("applied_metrics_push_interval", VRSMetrics.AppliedMetricsPushInterval)
	d.Set("associated_vcenter_hypervisor_id", VRSMetrics.AssociatedVCenterHypervisorID)
	d.Set("current_version", VRSMetrics.CurrentVersion)
	d.Set("external_id", VRSMetrics.ExternalID)

	d.Set("id", VRSMetrics.Identifier())
	d.Set("parent_id", VRSMetrics.ParentID)
	d.Set("parent_type", VRSMetrics.ParentType)
	d.Set("owner", VRSMetrics.Owner)

	d.SetId(VRSMetrics.Identifier())

	return nil
}
