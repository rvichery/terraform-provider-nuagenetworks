package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVNFInterfaceDescriptor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNFInterfaceDescriptorRead,
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vnf_descriptor": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVNFInterfaceDescriptorRead(d *schema.ResourceData, m interface{}) error {
	filteredVNFInterfaceDescriptors := vspk.VNFInterfaceDescriptorsList{}
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
	parent := &vspk.VNFDescriptor{ID: d.Get("parent_vnf_descriptor").(string)}
	filteredVNFInterfaceDescriptors, err = parent.VNFInterfaceDescriptors(fetchFilter)
	if err != nil {
		return err
	}

	VNFInterfaceDescriptor := &vspk.VNFInterfaceDescriptor{}

	if len(filteredVNFInterfaceDescriptors) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVNFInterfaceDescriptors) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		VNFInterfaceDescriptor = filteredVNFInterfaceDescriptors[0]
	}

	d.Set("name", VNFInterfaceDescriptor.Name)
	d.Set("type", VNFInterfaceDescriptor.Type)

	d.Set("id", VNFInterfaceDescriptor.Identifier())
	d.Set("parent_id", VNFInterfaceDescriptor.ParentID)
	d.Set("parent_type", VNFInterfaceDescriptor.ParentType)
	d.Set("owner", VNFInterfaceDescriptor.Owner)

	d.SetId(VNFInterfaceDescriptor.Identifier())

	return nil
}
