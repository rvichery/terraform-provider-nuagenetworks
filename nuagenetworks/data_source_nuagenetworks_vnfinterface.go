package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceVNFInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNFInterfaceRead,
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
			"vnfuuid": {
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
			"attached_network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attached_network_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vnf"},
			},
			"parent_vnf": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport"},
			},
		},
	}
}

func dataSourceVNFInterfaceRead(d *schema.ResourceData, m interface{}) error {
	filteredVNFInterfaces := vspk.VNFInterfacesList{}
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
		filteredVNFInterfaces, err = parent.VNFInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vnf"); ok {
		parent := &vspk.VNF{ID: attr.(string)}
		filteredVNFInterfaces, err = parent.VNFInterfaces(fetchFilter)
		if err != nil {
			return err
		}
	}

	VNFInterface := &vspk.VNFInterface{}

	if len(filteredVNFInterfaces) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVNFInterfaces) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VNFInterface = filteredVNFInterfaces[0]

	d.Set("mac", VNFInterface.MAC)
	d.Set("vnfuuid", VNFInterface.VNFUUID)
	d.Set("ip_address", VNFInterface.IPAddress)
	d.Set("vport_id", VNFInterface.VPortID)
	d.Set("vport_name", VNFInterface.VPortName)
	d.Set("name", VNFInterface.Name)
	d.Set("gateway", VNFInterface.Gateway)
	d.Set("netmask", VNFInterface.Netmask)
	d.Set("network_name", VNFInterface.NetworkName)
	d.Set("policy_decision_id", VNFInterface.PolicyDecisionID)
	d.Set("domain_id", VNFInterface.DomainID)
	d.Set("domain_name", VNFInterface.DomainName)
	d.Set("zone_id", VNFInterface.ZoneID)
	d.Set("zone_name", VNFInterface.ZoneName)
	d.Set("attached_network_id", VNFInterface.AttachedNetworkID)
	d.Set("attached_network_type", VNFInterface.AttachedNetworkType)
	d.Set("type", VNFInterface.Type)

	d.Set("id", VNFInterface.Identifier())
	d.Set("parent_id", VNFInterface.ParentID)
	d.Set("parent_type", VNFInterface.ParentType)
	d.Set("owner", VNFInterface.Owner)

	d.SetId(VNFInterface.Identifier())

	return nil
}
