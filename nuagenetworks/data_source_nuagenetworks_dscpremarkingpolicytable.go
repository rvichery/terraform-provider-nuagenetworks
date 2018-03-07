package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDSCPRemarkingPolicyTable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDSCPRemarkingPolicyTableRead,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
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
				Optional: true,
			},
		},
	}
}

func dataSourceDSCPRemarkingPolicyTableRead(d *schema.ResourceData, m interface{}) error {
	filteredDSCPRemarkingPolicyTables := vspk.DSCPRemarkingPolicyTablesList{}
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
	if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredDSCPRemarkingPolicyTables, err = parent.DSCPRemarkingPolicyTables(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredDSCPRemarkingPolicyTables, err = parent.DSCPRemarkingPolicyTables(fetchFilter)
		if err != nil {
			return err
		}
	}

	DSCPRemarkingPolicyTable := &vspk.DSCPRemarkingPolicyTable{}

	if len(filteredDSCPRemarkingPolicyTables) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDSCPRemarkingPolicyTables) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		DSCPRemarkingPolicyTable = filteredDSCPRemarkingPolicyTables[0]
	}

	d.Set("name", DSCPRemarkingPolicyTable.Name)
	d.Set("last_updated_by", DSCPRemarkingPolicyTable.LastUpdatedBy)
	d.Set("description", DSCPRemarkingPolicyTable.Description)
	d.Set("entity_scope", DSCPRemarkingPolicyTable.EntityScope)
	d.Set("external_id", DSCPRemarkingPolicyTable.ExternalID)

	d.Set("id", DSCPRemarkingPolicyTable.Identifier())
	d.Set("parent_id", DSCPRemarkingPolicyTable.ParentID)
	d.Set("parent_type", DSCPRemarkingPolicyTable.ParentType)
	d.Set("owner", DSCPRemarkingPolicyTable.Owner)

	d.SetId(DSCPRemarkingPolicyTable.Identifier())

	return nil
}
