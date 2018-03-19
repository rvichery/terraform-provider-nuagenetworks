package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceVsgRedundantPort() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVsgRedundantPortRead,
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
			"vlan_range": {
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
			"permitted_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"physical_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_peer1_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_peer2_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_user_mnemonic": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"user_mnemonic": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_egress_qos_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_redundancy_group": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVsgRedundantPortRead(d *schema.ResourceData, m interface{}) error {
	filteredVsgRedundantPorts := vspk.VsgRedundantPortsList{}
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
	parent := &vspk.RedundancyGroup{ID: d.Get("parent_redundancy_group").(string)}
	filteredVsgRedundantPorts, err = parent.VsgRedundantPorts(fetchFilter)
	if err != nil {
		return err
	}

	VsgRedundantPort := &vspk.VsgRedundantPort{}

	if len(filteredVsgRedundantPorts) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVsgRedundantPorts) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VsgRedundantPort = filteredVsgRedundantPorts[0]

	d.Set("vlan_range", VsgRedundantPort.VLANRange)
	d.Set("name", VsgRedundantPort.Name)
	d.Set("last_updated_by", VsgRedundantPort.LastUpdatedBy)
	d.Set("permitted_action", VsgRedundantPort.PermittedAction)
	d.Set("description", VsgRedundantPort.Description)
	d.Set("physical_name", VsgRedundantPort.PhysicalName)
	d.Set("entity_scope", VsgRedundantPort.EntityScope)
	d.Set("port_peer1_id", VsgRedundantPort.PortPeer1ID)
	d.Set("port_peer2_id", VsgRedundantPort.PortPeer2ID)
	d.Set("port_type", VsgRedundantPort.PortType)
	d.Set("use_user_mnemonic", VsgRedundantPort.UseUserMnemonic)
	d.Set("user_mnemonic", VsgRedundantPort.UserMnemonic)
	d.Set("associated_egress_qos_policy_id", VsgRedundantPort.AssociatedEgressQOSPolicyID)
	d.Set("status", VsgRedundantPort.Status)
	d.Set("external_id", VsgRedundantPort.ExternalID)

	d.Set("id", VsgRedundantPort.Identifier())
	d.Set("parent_id", VsgRedundantPort.ParentID)
	d.Set("parent_type", VsgRedundantPort.ParentType)
	d.Set("owner", VsgRedundantPort.Owner)

	d.SetId(VsgRedundantPort.Identifier())

	return nil
}
