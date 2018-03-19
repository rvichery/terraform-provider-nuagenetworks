package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceDUCGroupBinding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDUCGroupBindingRead,
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
			"one_way_delay": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_duc_group_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_nsg_group": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceDUCGroupBindingRead(d *schema.ResourceData, m interface{}) error {
	filteredDUCGroupBindings := vspk.DUCGroupBindingsList{}
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
	parent := &vspk.NSGGroup{ID: d.Get("parent_nsg_group").(string)}
	filteredDUCGroupBindings, err = parent.DUCGroupBindings(fetchFilter)
	if err != nil {
		return err
	}

	DUCGroupBinding := &vspk.DUCGroupBinding{}

	if len(filteredDUCGroupBindings) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDUCGroupBindings) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	DUCGroupBinding = filteredDUCGroupBindings[0]

	d.Set("one_way_delay", DUCGroupBinding.OneWayDelay)
	d.Set("priority", DUCGroupBinding.Priority)
	d.Set("associated_duc_group_id", DUCGroupBinding.AssociatedDUCGroupID)

	d.Set("id", DUCGroupBinding.Identifier())
	d.Set("parent_id", DUCGroupBinding.ParentID)
	d.Set("parent_type", DUCGroupBinding.ParentType)
	d.Set("owner", DUCGroupBinding.Owner)

	d.SetId(DUCGroupBinding.Identifier())

	return nil
}
