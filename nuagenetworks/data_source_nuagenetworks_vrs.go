package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVRS() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVRSRead,
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
			"jsonrpc_connection_state": {
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
			"parent_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"last_event_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_event_object": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_event_timestamp": {
				Type:     schema.TypeInt,
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
			"db_synced": {
				Type:     schema.TypeBool,
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
			"peer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"personality": {
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
			"revert_behavior_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"revert_completed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"revert_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"revert_failed_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"licensed_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cluster_node_role": {
				Type:     schema.TypeString,
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
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uptime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"primary_vsc_connection_lost": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"product_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_resilient": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vsc_config_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsc_current_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multi_nic_vport_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"number_of_bridge_interfaces": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_containers": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_host_interfaces": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_virtual_machines": {
				Type:     schema.TypeInt,
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
			"dynamic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hypervisor_connection_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container", "parent_hsc", "parent_vm", "parent_vsc"},
			},
			"parent_container": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_hsc", "parent_vm", "parent_vsc"},
			},
			"parent_hsc": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_container", "parent_vm", "parent_vsc"},
			},
			"parent_vm": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_container", "parent_hsc", "parent_vsc"},
			},
			"parent_vsc": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_container", "parent_hsc", "parent_vm"},
			},
		},
	}
}

func dataSourceVRSRead(d *schema.ResourceData, m interface{}) error {
	filteredVRSs := vspk.VRSsList{}
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
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredVRSs, err = parent.VRSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_container"); ok {
		parent := &vspk.Container{ID: attr.(string)}
		filteredVRSs, err = parent.VRSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_hsc"); ok {
		parent := &vspk.HSC{ID: attr.(string)}
		filteredVRSs, err = parent.VRSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm"); ok {
		parent := &vspk.VM{ID: attr.(string)}
		filteredVRSs, err = parent.VRSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vsc"); ok {
		parent := &vspk.VSC{ID: attr.(string)}
		filteredVRSs, err = parent.VRSs(fetchFilter)
		if err != nil {
			return err
		}
	}

	VRS := &vspk.VRS{}

	if len(filteredVRSs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVRSs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VRS = filteredVRSs[0]

	d.Set("jsonrpc_connection_state", VRS.JSONRPCConnectionState)
	d.Set("name", VRS.Name)
	d.Set("management_ip", VRS.ManagementIP)
	d.Set("parent_ids", VRS.ParentIDs)
	d.Set("last_event_name", VRS.LastEventName)
	d.Set("last_event_object", VRS.LastEventObject)
	d.Set("last_event_timestamp", VRS.LastEventTimestamp)
	d.Set("last_state_change", VRS.LastStateChange)
	d.Set("last_updated_by", VRS.LastUpdatedBy)
	d.Set("db_synced", VRS.DbSynced)
	d.Set("address", VRS.Address)
	d.Set("peak_cpuusage", VRS.PeakCPUUsage)
	d.Set("peak_memory_usage", VRS.PeakMemoryUsage)
	d.Set("peer", VRS.Peer)
	d.Set("personality", VRS.Personality)
	d.Set("description", VRS.Description)
	d.Set("messages", VRS.Messages)
	d.Set("revert_behavior_enabled", VRS.RevertBehaviorEnabled)
	d.Set("revert_completed", VRS.RevertCompleted)
	d.Set("revert_count", VRS.RevertCount)
	d.Set("revert_failed_count", VRS.RevertFailedCount)
	d.Set("licensed_state", VRS.LicensedState)
	d.Set("disks", VRS.Disks)
	d.Set("cluster_node_role", VRS.ClusterNodeRole)
	d.Set("entity_scope", VRS.EntityScope)
	d.Set("location", VRS.Location)
	d.Set("role", VRS.Role)
	d.Set("uptime", VRS.Uptime)
	d.Set("primary_vsc_connection_lost", VRS.PrimaryVSCConnectionLost)
	d.Set("product_version", VRS.ProductVersion)
	d.Set("is_resilient", VRS.IsResilient)
	d.Set("vsc_config_state", VRS.VscConfigState)
	d.Set("vsc_current_state", VRS.VscCurrentState)
	d.Set("status", VRS.Status)
	d.Set("multi_nic_vport_enabled", VRS.MultiNICVPortEnabled)
	d.Set("number_of_bridge_interfaces", VRS.NumberOfBridgeInterfaces)
	d.Set("number_of_containers", VRS.NumberOfContainers)
	d.Set("number_of_host_interfaces", VRS.NumberOfHostInterfaces)
	d.Set("number_of_virtual_machines", VRS.NumberOfVirtualMachines)
	d.Set("current_cpuusage", VRS.CurrentCPUUsage)
	d.Set("current_memory_usage", VRS.CurrentMemoryUsage)
	d.Set("average_cpuusage", VRS.AverageCPUUsage)
	d.Set("average_memory_usage", VRS.AverageMemoryUsage)
	d.Set("external_id", VRS.ExternalID)
	d.Set("dynamic", VRS.Dynamic)
	d.Set("hypervisor_connection_state", VRS.HypervisorConnectionState)
	d.Set("hypervisor_identifier", VRS.HypervisorIdentifier)
	d.Set("hypervisor_name", VRS.HypervisorName)
	d.Set("hypervisor_type", VRS.HypervisorType)

	d.Set("id", VRS.Identifier())
	d.Set("parent_id", VRS.ParentID)
	d.Set("parent_type", VRS.ParentType)
	d.Set("owner", VRS.Owner)

	d.SetId(VRS.Identifier())

	return nil
}
