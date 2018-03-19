package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceAlarm() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAlarmRead,
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
			"target_object": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"acknowledged": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"severity": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_condition": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"number_of_occurances": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_redundancy_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_tca": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_vsg_redundant_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_ike_gateway_connection": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_wan_service": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_vsd": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_container": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_vrs": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_hsc": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_vlan": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_ssid_connection": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_vm": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_ns_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_vsc": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_ns_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_redundant_gateway_group", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_ns_redundant_gateway_group": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_wireless_port", "parent_enterprise"},
			},
			"parent_wireless_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_enterprise"},
			},
			"parent_enterprise": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_tca", "parent_vsg_redundant_port", "parent_ike_gateway_connection", "parent_wan_service", "parent_vsd", "parent_vport", "parent_port", "parent_container", "parent_vrs", "parent_hsc", "parent_vlan", "parent_ssid_connection", "parent_gateway", "parent_vm", "parent_ns_gateway", "parent_vsc", "parent_ns_port", "parent_ns_redundant_gateway_group", "parent_wireless_port"},
			},
		},
	}
}

func dataSourceAlarmRead(d *schema.ResourceData, m interface{}) error {
	filteredAlarms := vspk.AlarmsList{}
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
	if attr, ok := d.GetOk("parent_redundancy_group"); ok {
		parent := &vspk.RedundancyGroup{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_tca"); ok {
		parent := &vspk.TCA{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vsg_redundant_port"); ok {
		parent := &vspk.VsgRedundantPort{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ike_gateway_connection"); ok {
		parent := &vspk.IKEGatewayConnection{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_wan_service"); ok {
		parent := &vspk.WANService{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vsd"); ok {
		parent := &vspk.VSD{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_port"); ok {
		parent := &vspk.Port{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_container"); ok {
		parent := &vspk.Container{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vrs"); ok {
		parent := &vspk.VRS{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_hsc"); ok {
		parent := &vspk.HSC{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ssid_connection"); ok {
		parent := &vspk.SSIDConnection{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm"); ok {
		parent := &vspk.VM{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
		parent := &vspk.NSGateway{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vsc"); ok {
		parent := &vspk.VSC{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_port"); ok {
		parent := &vspk.NSPort{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_redundant_gateway_group"); ok {
		parent := &vspk.NSRedundantGatewayGroup{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_wireless_port"); ok {
		parent := &vspk.WirelessPort{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredAlarms, err = parent.Alarms(fetchFilter)
		if err != nil {
			return err
		}
	}

	Alarm := &vspk.Alarm{}

	if len(filteredAlarms) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredAlarms) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Alarm = filteredAlarms[0]

	d.Set("name", Alarm.Name)
	d.Set("target_object", Alarm.TargetObject)
	d.Set("last_updated_by", Alarm.LastUpdatedBy)
	d.Set("acknowledged", Alarm.Acknowledged)
	d.Set("reason", Alarm.Reason)
	d.Set("description", Alarm.Description)
	d.Set("severity", Alarm.Severity)
	d.Set("timestamp", Alarm.Timestamp)
	d.Set("enterprise_id", Alarm.EnterpriseID)
	d.Set("entity_scope", Alarm.EntityScope)
	d.Set("error_condition", Alarm.ErrorCondition)
	d.Set("number_of_occurances", Alarm.NumberOfOccurances)
	d.Set("external_id", Alarm.ExternalID)

	d.Set("id", Alarm.Identifier())
	d.Set("parent_id", Alarm.ParentID)
	d.Set("parent_type", Alarm.ParentType)
	d.Set("owner", Alarm.Owner)

	d.SetId(Alarm.Identifier())

	return nil
}
