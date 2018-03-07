package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceSubnetTemplate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSubnetTemplateRead,
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
			"dpi": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_gateway": &schema.Schema{
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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"encryption": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"split_subnet": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"proxy_arp": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"use_global_mac": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_multicast_channel_map_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_ipv6_address": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"parent_zone_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain_template"},
			},
			"parent_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone_template"},
			},
		},
	}
}

func dataSourceSubnetTemplateRead(d *schema.ResourceData, m interface{}) error {
	filteredSubnetTemplates := vspk.SubnetTemplatesList{}
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
	if attr, ok := d.GetOk("parent_zone_template"); ok {
		parent := &vspk.ZoneTemplate{ID: attr.(string)}
		filteredSubnetTemplates, err = parent.SubnetTemplates(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredSubnetTemplates, err = parent.SubnetTemplates(fetchFilter)
		if err != nil {
			return err
		}
	}

	SubnetTemplate := &vspk.SubnetTemplate{}

	if len(filteredSubnetTemplates) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredSubnetTemplates) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		SubnetTemplate = filteredSubnetTemplates[0]
	}

	d.Set("dpi", SubnetTemplate.DPI)
	d.Set("ip_type", SubnetTemplate.IPType)
	d.Set("ipv6_address", SubnetTemplate.IPv6Address)
	d.Set("ipv6_gateway", SubnetTemplate.IPv6Gateway)
	d.Set("name", SubnetTemplate.Name)
	d.Set("last_updated_by", SubnetTemplate.LastUpdatedBy)
	d.Set("gateway", SubnetTemplate.Gateway)
	d.Set("address", SubnetTemplate.Address)
	d.Set("description", SubnetTemplate.Description)
	d.Set("netmask", SubnetTemplate.Netmask)
	d.Set("encryption", SubnetTemplate.Encryption)
	d.Set("entity_scope", SubnetTemplate.EntityScope)
	d.Set("split_subnet", SubnetTemplate.SplitSubnet)
	d.Set("proxy_arp", SubnetTemplate.ProxyARP)
	d.Set("use_global_mac", SubnetTemplate.UseGlobalMAC)
	d.Set("associated_multicast_channel_map_id", SubnetTemplate.AssociatedMulticastChannelMapID)
	d.Set("multicast", SubnetTemplate.Multicast)
	d.Set("external_id", SubnetTemplate.ExternalID)
	d.Set("dynamic_ipv6_address", SubnetTemplate.DynamicIpv6Address)

	d.Set("id", SubnetTemplate.Identifier())
	d.Set("parent_id", SubnetTemplate.ParentID)
	d.Set("parent_type", SubnetTemplate.ParentType)
	d.Set("owner", SubnetTemplate.Owner)

	d.SetId(SubnetTemplate.Identifier())

	return nil
}
