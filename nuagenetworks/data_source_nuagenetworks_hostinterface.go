package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceHostInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHostInterfaceRead,
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
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vport_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vport_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tier_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_decision_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"zone_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_floating_ip_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"attached_network_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"attached_network_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_l2_domain"},
			},
			"parent_vport": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_l2_domain"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport"},
			},
		},
	}
}

func dataSourceHostInterfaceRead(d *schema.ResourceData, m interface{}) error {
	filteredHostInterfaces := vspk.HostInterfacesList{}
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
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredHostInterfaces, err = parent.HostInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredHostInterfaces, err = parent.HostInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredHostInterfaces, err = parent.HostInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredHostInterfaces, err = parent.HostInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	}

	HostInterface := &vspk.HostInterface{}

	if len(filteredHostInterfaces) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredHostInterfaces) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		HostInterface = filteredHostInterfaces[0]
	}

	d.Set("mac", HostInterface.MAC)
	d.Set("ip_address", HostInterface.IPAddress)
	d.Set("vport_id", HostInterface.VPortID)
	d.Set("vport_name", HostInterface.VPortName)
	d.Set("name", HostInterface.Name)
	d.Set("last_updated_by", HostInterface.LastUpdatedBy)
	d.Set("gateway", HostInterface.Gateway)
	d.Set("netmask", HostInterface.Netmask)
	d.Set("network_name", HostInterface.NetworkName)
	d.Set("tier_id", HostInterface.TierID)
	d.Set("entity_scope", HostInterface.EntityScope)
	d.Set("policy_decision_id", HostInterface.PolicyDecisionID)
	d.Set("domain_id", HostInterface.DomainID)
	d.Set("domain_name", HostInterface.DomainName)
	d.Set("zone_id", HostInterface.ZoneID)
	d.Set("zone_name", HostInterface.ZoneName)
	d.Set("associated_floating_ip_address", HostInterface.AssociatedFloatingIPAddress)
	d.Set("attached_network_id", HostInterface.AttachedNetworkID)
	d.Set("attached_network_type", HostInterface.AttachedNetworkType)
	d.Set("external_id", HostInterface.ExternalID)

	d.Set("id", HostInterface.Identifier())
	d.Set("parent_id", HostInterface.ParentID)
	d.Set("parent_type", HostInterface.ParentType)
	d.Set("owner", HostInterface.Owner)

	d.SetId(HostInterface.Identifier())

	return nil
}
