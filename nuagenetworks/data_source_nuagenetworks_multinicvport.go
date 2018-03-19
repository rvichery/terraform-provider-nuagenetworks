package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceMultiNICVPort() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMultiNICVPortRead,
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
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vrs": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceMultiNICVPortRead(d *schema.ResourceData, m interface{}) error {
	filteredMultiNICVPorts := vspk.MultiNICVPortsList{}
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
	parent := &vspk.VRS{ID: d.Get("parent_vrs").(string)}
	filteredMultiNICVPorts, err = parent.MultiNICVPorts(fetchFilter)
	if err != nil {
		return err
	}

	MultiNICVPort := &vspk.MultiNICVPort{}

	if len(filteredMultiNICVPorts) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredMultiNICVPorts) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	MultiNICVPort = filteredMultiNICVPorts[0]

	d.Set("name", MultiNICVPort.Name)
	d.Set("last_updated_by", MultiNICVPort.LastUpdatedBy)
	d.Set("entity_scope", MultiNICVPort.EntityScope)
	d.Set("external_id", MultiNICVPort.ExternalID)

	d.Set("id", MultiNICVPort.Identifier())
	d.Set("parent_id", MultiNICVPort.ParentID)
	d.Set("parent_type", MultiNICVPort.ParentType)
	d.Set("owner", MultiNICVPort.Owner)

	d.SetId(MultiNICVPort.Identifier())

	return nil
}
