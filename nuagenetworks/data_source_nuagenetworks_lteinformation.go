package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceLTEInformation() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLTEInformationRead,
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
			"lte_connection_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ns_port": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceLTEInformationRead(d *schema.ResourceData, m interface{}) error {
	filteredLTEInformations := vspk.LTEInformationsList{}
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
	parent := &vspk.NSPort{ID: d.Get("parent_ns_port").(string)}
	filteredLTEInformations, err = parent.LTEInformations(fetchFilter)
	if err != nil {
		return err
	}

	LTEInformation := &vspk.LTEInformation{}

	if len(filteredLTEInformations) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredLTEInformations) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	LTEInformation = filteredLTEInformations[0]

	d.Set("lte_connection_info", LTEInformation.LTEConnectionInfo)

	d.Set("id", LTEInformation.Identifier())
	d.Set("parent_id", LTEInformation.ParentID)
	d.Set("parent_type", LTEInformation.ParentType)
	d.Set("owner", LTEInformation.Owner)

	d.SetId(LTEInformation.Identifier())

	return nil
}
