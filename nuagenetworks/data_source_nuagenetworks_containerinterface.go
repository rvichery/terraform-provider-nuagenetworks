package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceContainerInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContainerInterfaceRead,
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
			"network_id": {
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
			"endpoint_id": {
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
			"container_uuid": {
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
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_subnet", "parent_container", "parent_l2_domain"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_vport", "parent_subnet", "parent_container", "parent_l2_domain"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_subnet", "parent_container", "parent_l2_domain"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_container", "parent_l2_domain"},
			},
			"parent_container": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_container"},
			},
		},
	}
}

func dataSourceContainerInterfaceRead(d *schema.ResourceData, m interface{}) error {
	filteredContainerInterfaces := vspk.ContainerInterfacesList{}
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
		filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_container"); ok {
		parent := &vspk.Container{ID: attr.(string)}
		filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredContainerInterfaces, err = parent.ContainerInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	}

	ContainerInterface := &vspk.ContainerInterface{}

	if len(filteredContainerInterfaces) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredContainerInterfaces) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	ContainerInterface = filteredContainerInterfaces[0]

	d.Set("mac", ContainerInterface.MAC)
	d.Set("ip_address", ContainerInterface.IPAddress)
	d.Set("vport_id", ContainerInterface.VPortID)
	d.Set("vport_name", ContainerInterface.VPortName)
	d.Set("name", ContainerInterface.Name)
	d.Set("last_updated_by", ContainerInterface.LastUpdatedBy)
	d.Set("gateway", ContainerInterface.Gateway)
	d.Set("netmask", ContainerInterface.Netmask)
	d.Set("network_id", ContainerInterface.NetworkID)
	d.Set("network_name", ContainerInterface.NetworkName)
	d.Set("tier_id", ContainerInterface.TierID)
	d.Set("endpoint_id", ContainerInterface.EndpointID)
	d.Set("entity_scope", ContainerInterface.EntityScope)
	d.Set("policy_decision_id", ContainerInterface.PolicyDecisionID)
	d.Set("domain_id", ContainerInterface.DomainID)
	d.Set("domain_name", ContainerInterface.DomainName)
	d.Set("zone_id", ContainerInterface.ZoneID)
	d.Set("zone_name", ContainerInterface.ZoneName)
	d.Set("container_uuid", ContainerInterface.ContainerUUID)
	d.Set("associated_floating_ip_address", ContainerInterface.AssociatedFloatingIPAddress)
	d.Set("attached_network_id", ContainerInterface.AttachedNetworkID)
	d.Set("attached_network_type", ContainerInterface.AttachedNetworkType)
	d.Set("multi_nic_vport_name", ContainerInterface.MultiNICVPortName)
	d.Set("external_id", ContainerInterface.ExternalID)

	d.Set("id", ContainerInterface.Identifier())
	d.Set("parent_id", ContainerInterface.ParentID)
	d.Set("parent_type", ContainerInterface.ParentType)
	d.Set("owner", ContainerInterface.Owner)

	d.SetId(ContainerInterface.Identifier())

	return nil
}
