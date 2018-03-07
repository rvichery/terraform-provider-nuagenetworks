package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceAddressRange() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAddressRangeRead,
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
			"dhcp_pool_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_address": &schema.Schema{
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
			"parent_shared_network_resource": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet_template", "parent_subnet", "parent_l2_domain_template", "parent_l2_domain"},
			},
			"parent_subnet_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_subnet", "parent_l2_domain_template", "parent_l2_domain"},
			},
			"parent_subnet": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_subnet_template", "parent_l2_domain_template", "parent_l2_domain"},
			},
			"parent_l2_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_subnet_template", "parent_subnet", "parent_l2_domain"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_shared_network_resource", "parent_subnet_template", "parent_subnet", "parent_l2_domain_template"},
			},
		},
	}
}

func dataSourceAddressRangeRead(d *schema.ResourceData, m interface{}) error {
	filteredAddressRanges := vspk.AddressRangesList{}
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
	if attr, ok := d.GetOk("parent_shared_network_resource"); ok {
		parent := &vspk.SharedNetworkResource{ID: attr.(string)}
		filteredAddressRanges, err = parent.AddressRanges(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet_template"); ok {
		parent := &vspk.SubnetTemplate{ID: attr.(string)}
		filteredAddressRanges, err = parent.AddressRanges(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredAddressRanges, err = parent.AddressRanges(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		filteredAddressRanges, err = parent.AddressRanges(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredAddressRanges, err = parent.AddressRanges(fetchFilter)
		if err != nil {
			return err
		}
	}

	AddressRange := &vspk.AddressRange{}

	if len(filteredAddressRanges) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredAddressRanges) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		AddressRange = filteredAddressRanges[0]
	}

	d.Set("dhcp_pool_type", AddressRange.DHCPPoolType)
	d.Set("ip_type", AddressRange.IPType)
	d.Set("last_updated_by", AddressRange.LastUpdatedBy)
	d.Set("max_address", AddressRange.MaxAddress)
	d.Set("min_address", AddressRange.MinAddress)
	d.Set("entity_scope", AddressRange.EntityScope)
	d.Set("external_id", AddressRange.ExternalID)

	d.Set("id", AddressRange.Identifier())
	d.Set("parent_id", AddressRange.ParentID)
	d.Set("parent_type", AddressRange.ParentType)
	d.Set("owner", AddressRange.Owner)

	d.SetId(AddressRange.Identifier())

	return nil
}
