package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourcePublicNetworkMacro() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePublicNetworkMacroRead,
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
			"ip_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipv6_address": &schema.Schema{
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
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"netmask": &schema.Schema{
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
			"parent_enterprise": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourcePublicNetworkMacroRead(d *schema.ResourceData, m interface{}) error {
	filteredPublicNetworkMacros := vspk.PublicNetworkMacrosList{}
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
	filteredPublicNetworkMacros, err = parent.PublicNetworkMacros(fetchFilter)
	if err != nil {
		return err
	}

	PublicNetworkMacro := &vspk.PublicNetworkMacro{}

	if len(filteredPublicNetworkMacros) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPublicNetworkMacros) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		PublicNetworkMacro = filteredPublicNetworkMacros[0]
	}

	d.Set("ip_type", PublicNetworkMacro.IPType)
	d.Set("ipv6_address", PublicNetworkMacro.IPv6Address)
	d.Set("name", PublicNetworkMacro.Name)
	d.Set("last_updated_by", PublicNetworkMacro.LastUpdatedBy)
	d.Set("address", PublicNetworkMacro.Address)
	d.Set("netmask", PublicNetworkMacro.Netmask)
	d.Set("entity_scope", PublicNetworkMacro.EntityScope)
	d.Set("external_id", PublicNetworkMacro.ExternalID)

	d.Set("id", PublicNetworkMacro.Identifier())
	d.Set("parent_id", PublicNetworkMacro.ParentID)
	d.Set("parent_type", PublicNetworkMacro.ParentType)
	d.Set("owner", PublicNetworkMacro.Owner)

	d.SetId(PublicNetworkMacro.Identifier())

	return nil
}
