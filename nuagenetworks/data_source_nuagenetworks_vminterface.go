package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVMInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVMInterfaceRead,
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
			"mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vmuuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vport_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vport_name": {
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
			"gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tier_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_decision_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_floating_ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attached_network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attached_network_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multi_nic_vport_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_zone": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_subnet", "parent_vm", "parent_l2_domain"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_vport", "parent_subnet", "parent_vm", "parent_l2_domain"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_subnet", "parent_vm", "parent_l2_domain"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_vm", "parent_l2_domain"},
			},
			"parent_vm": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm"},
			},
		},
	}
}

func dataSourceVMInterfaceRead(d *schema.ResourceData, m interface{}) error {
	filteredVMInterfaces := vspk.VMInterfacesList{}
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
	if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		filteredVMInterfaces, err = parent.VMInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredVMInterfaces, err = parent.VMInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredVMInterfaces, err = parent.VMInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredVMInterfaces, err = parent.VMInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm"); ok {
		parent := &vspk.VM{ID: attr.(string)}
		filteredVMInterfaces, err = parent.VMInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredVMInterfaces, err = parent.VMInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredVMInterfaces, err = parent.VMInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	}

	VMInterface := &vspk.VMInterface{}

	if len(filteredVMInterfaces) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVMInterfaces) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VMInterface = filteredVMInterfaces[0]

	d.Set("mac", VMInterface.MAC)
	d.Set("vmuuid", VMInterface.VMUUID)
	d.Set("ip_address", VMInterface.IPAddress)
	d.Set("vport_id", VMInterface.VPortID)
	d.Set("vport_name", VMInterface.VPortName)
	d.Set("name", VMInterface.Name)
	d.Set("last_updated_by", VMInterface.LastUpdatedBy)
	d.Set("gateway", VMInterface.Gateway)
	d.Set("netmask", VMInterface.Netmask)
	d.Set("network_name", VMInterface.NetworkName)
	d.Set("tier_id", VMInterface.TierID)
	d.Set("entity_scope", VMInterface.EntityScope)
	d.Set("policy_decision_id", VMInterface.PolicyDecisionID)
	d.Set("domain_id", VMInterface.DomainID)
	d.Set("domain_name", VMInterface.DomainName)
	d.Set("zone_id", VMInterface.ZoneID)
	d.Set("zone_name", VMInterface.ZoneName)
	d.Set("associated_floating_ip_address", VMInterface.AssociatedFloatingIPAddress)
	d.Set("attached_network_id", VMInterface.AttachedNetworkID)
	d.Set("attached_network_type", VMInterface.AttachedNetworkType)
	d.Set("multi_nic_vport_name", VMInterface.MultiNICVPortName)
	d.Set("external_id", VMInterface.ExternalID)

	d.Set("id", VMInterface.Identifier())
	d.Set("parent_id", VMInterface.ParentID)
	d.Set("parent_type", VMInterface.ParentType)
	d.Set("owner", VMInterface.Owner)

	d.SetId(VMInterface.Identifier())

	return nil
}
