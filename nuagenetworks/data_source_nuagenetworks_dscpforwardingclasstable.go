package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceDSCPForwardingClassTable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDSCPForwardingClassTableRead,
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
			"description": {
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
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceDSCPForwardingClassTableRead(d *schema.ResourceData, m interface{}) error {
	filteredDSCPForwardingClassTables := vspk.DSCPForwardingClassTablesList{}
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
	filteredDSCPForwardingClassTables, err = parent.DSCPForwardingClassTables(fetchFilter)
	if err != nil {
		return err
	}

	DSCPForwardingClassTable := &vspk.DSCPForwardingClassTable{}

	if len(filteredDSCPForwardingClassTables) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDSCPForwardingClassTables) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	DSCPForwardingClassTable = filteredDSCPForwardingClassTables[0]

	d.Set("name", DSCPForwardingClassTable.Name)
	d.Set("last_updated_by", DSCPForwardingClassTable.LastUpdatedBy)
	d.Set("description", DSCPForwardingClassTable.Description)
	d.Set("entity_scope", DSCPForwardingClassTable.EntityScope)
	d.Set("external_id", DSCPForwardingClassTable.ExternalID)

	d.Set("id", DSCPForwardingClassTable.Identifier())
	d.Set("parent_id", DSCPForwardingClassTable.ParentID)
	d.Set("parent_type", DSCPForwardingClassTable.ParentType)
	d.Set("owner", DSCPForwardingClassTable.Owner)

	d.SetId(DSCPForwardingClassTable.Identifier())

	return nil
}
