package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceAddressMap() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAddressMapRead,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_patnat_pool_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_port": &schema.Schema{
				Type:     schema.TypeInt,
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
			"parent_patnat_pool": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceAddressMapRead(d *schema.ResourceData, m interface{}) error {
	filteredAddressMaps := vspk.AddressMapsList{}
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
	parent := &vspk.PATNATPool{ID: d.Get("parent_patnat_pool").(string)}
	filteredAddressMaps, err = parent.AddressMaps(fetchFilter)
	if err != nil {
		return err
	}

	AddressMap := &vspk.AddressMap{}

	if len(filteredAddressMaps) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredAddressMaps) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		AddressMap = filteredAddressMaps[0]
	}

	d.Set("last_updated_by", AddressMap.LastUpdatedBy)
	d.Set("entity_scope", AddressMap.EntityScope)
	d.Set("private_ip", AddressMap.PrivateIP)
	d.Set("private_port", AddressMap.PrivatePort)
	d.Set("associated_patnat_pool_id", AddressMap.AssociatedPATNATPoolID)
	d.Set("public_ip", AddressMap.PublicIP)
	d.Set("public_port", AddressMap.PublicPort)
	d.Set("external_id", AddressMap.ExternalID)
	d.Set("type", AddressMap.Type)

	d.Set("id", AddressMap.Identifier())
	d.Set("parent_id", AddressMap.ParentID)
	d.Set("parent_type", AddressMap.ParentType)
	d.Set("owner", AddressMap.Owner)

	d.SetId(AddressMap.Identifier())

	return nil
}
