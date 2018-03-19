package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceL2DomainTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceL2DomainTemplateRead,
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
			"dhcp_managed": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dpi": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_gateway": {
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
			"address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"encryption": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_change_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_global_mac": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_multicast_channel_map_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_ipv6_address": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceL2DomainTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredL2DomainTemplates := vspk.L2DomainTemplatesList{}
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
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	filteredL2DomainTemplates, err = parent.L2DomainTemplates(fetchFilter)
	if err != nil {
		return err
	}

	L2DomainTemplate := &vspk.L2DomainTemplate{}

	if len(filteredL2DomainTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredL2DomainTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	L2DomainTemplate = filteredL2DomainTemplates[0]

	d.Set("dhcp_managed", L2DomainTemplate.DHCPManaged)
	d.Set("dpi", L2DomainTemplate.DPI)
	d.Set("ip_type", L2DomainTemplate.IPType)
	d.Set("ipv6_address", L2DomainTemplate.IPv6Address)
	d.Set("ipv6_gateway", L2DomainTemplate.IPv6Gateway)
	d.Set("name", L2DomainTemplate.Name)
	d.Set("last_updated_by", L2DomainTemplate.LastUpdatedBy)
	d.Set("gateway", L2DomainTemplate.Gateway)
	d.Set("address", L2DomainTemplate.Address)
	d.Set("description", L2DomainTemplate.Description)
	d.Set("netmask", L2DomainTemplate.Netmask)
	d.Set("encryption", L2DomainTemplate.Encryption)
	d.Set("entity_scope", L2DomainTemplate.EntityScope)
	d.Set("entity_state", L2DomainTemplate.EntityState)
	d.Set("policy_change_status", L2DomainTemplate.PolicyChangeStatus)
	d.Set("use_global_mac", L2DomainTemplate.UseGlobalMAC)
	d.Set("associated_multicast_channel_map_id", L2DomainTemplate.AssociatedMulticastChannelMapID)
	d.Set("multicast", L2DomainTemplate.Multicast)
	d.Set("external_id", L2DomainTemplate.ExternalID)
	d.Set("dynamic_ipv6_address", L2DomainTemplate.DynamicIpv6Address)

	d.Set("id", L2DomainTemplate.Identifier())
	d.Set("parent_id", L2DomainTemplate.ParentID)
	d.Set("parent_type", L2DomainTemplate.ParentType)
	d.Set("owner", L2DomainTemplate.Owner)

	d.SetId(L2DomainTemplate.Identifier())

	return nil
}
