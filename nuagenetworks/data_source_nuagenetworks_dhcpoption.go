package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDHCPOption() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDHCPOptionRead,
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
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"actual_type": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"actual_values": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"length": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_container_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_shared_network_resource": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_zone": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_zone", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_vport": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_zone", "parent_domain", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_subnet": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_zone", "parent_domain", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_vm_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_bridge_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_host_interface"},
			},
			"parent_host_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_shared_network_resource", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain"},
			},
		},
	}
}

func dataSourceDHCPOptionRead(d *schema.ResourceData, m interface{}) error {
	filteredDHCPOptions := vspk.DHCPOptionsList{}
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
	if attr, ok := d.GetOk("parent_container_interface"); ok {
		parent := &vspk.ContainerInterface{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm_interface"); ok {
		parent := &vspk.VMInterface{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
		parent := &vspk.BridgeInterface{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_host_interface"); ok {
		parent := &vspk.HostInterface{ID: attr.(string)}
		filteredDHCPOptions, err = parent.DHCPOptions(fetchFilter)
		if err != nil {
			return err
		}
	}

	DHCPOption := &vspk.DHCPOption{}

	if len(filteredDHCPOptions) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDHCPOptions) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		DHCPOption = filteredDHCPOptions[0]
	}

	d.Set("value", DHCPOption.Value)
	d.Set("last_updated_by", DHCPOption.LastUpdatedBy)
	d.Set("actual_type", DHCPOption.ActualType)
	d.Set("actual_values", DHCPOption.ActualValues)
	d.Set("length", DHCPOption.Length)
	d.Set("entity_scope", DHCPOption.EntityScope)
	d.Set("external_id", DHCPOption.ExternalID)
	d.Set("type", DHCPOption.Type)

	d.Set("id", DHCPOption.Identifier())
	d.Set("parent_id", DHCPOption.ParentID)
	d.Set("parent_type", DHCPOption.ParentType)
	d.Set("owner", DHCPOption.Owner)

	d.SetId(DHCPOption.Identifier())

	return nil
}
