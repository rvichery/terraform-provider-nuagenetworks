package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourcePort() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePortRead,
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
			"template_id": {
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
			"port_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_resilient": {
				Type:     schema.TypeBool,
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
			"associated_redundant_port_id": {
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
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_auto_discovered_gateway", "parent_gateway"},
			},
			"parent_auto_discovered_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_gateway"},
			},
			"parent_gateway": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_redundancy_group", "parent_auto_discovered_gateway"},
			},
		},
	}
}

func dataSourcePortRead(d *schema.ResourceData, m interface{}) error {
	filteredPorts := vspk.PortsList{}
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
		filteredPorts, err = parent.Ports(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_auto_discovered_gateway"); ok {
		parent := &vspk.AutoDiscoveredGateway{ID: attr.(string)}
		filteredPorts, err = parent.Ports(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		filteredPorts, err = parent.Ports(fetchFilter)
		if err != nil {
			return err
		}
	}

	Port := &vspk.Port{}

	if len(filteredPorts) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPorts) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Port = filteredPorts[0]

	d.Set("vlan_range", Port.VLANRange)
	d.Set("name", Port.Name)
	d.Set("last_updated_by", Port.LastUpdatedBy)
	d.Set("template_id", Port.TemplateID)
	d.Set("permitted_action", Port.PermittedAction)
	d.Set("description", Port.Description)
	d.Set("physical_name", Port.PhysicalName)
	d.Set("entity_scope", Port.EntityScope)
	d.Set("port_type", Port.PortType)
	d.Set("is_resilient", Port.IsResilient)
	d.Set("use_user_mnemonic", Port.UseUserMnemonic)
	d.Set("user_mnemonic", Port.UserMnemonic)
	d.Set("associated_egress_qos_policy_id", Port.AssociatedEgressQOSPolicyID)
	d.Set("associated_redundant_port_id", Port.AssociatedRedundantPortID)
	d.Set("status", Port.Status)
	d.Set("external_id", Port.ExternalID)

	d.Set("id", Port.Identifier())
	d.Set("parent_id", Port.ParentID)
	d.Set("parent_type", Port.ParentType)
	d.Set("owner", Port.Owner)

	d.SetId(Port.Identifier())

	return nil
}
