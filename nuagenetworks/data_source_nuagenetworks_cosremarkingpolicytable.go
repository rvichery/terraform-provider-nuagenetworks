package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceCOSRemarkingPolicyTable() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCOSRemarkingPolicyTableRead,
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
				Optional: true,
			},
		},
	}
}

func dataSourceCOSRemarkingPolicyTableRead(d *schema.ResourceData, m interface{}) error {
	filteredCOSRemarkingPolicyTables := vspk.COSRemarkingPolicyTablesList{}
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
		filteredCOSRemarkingPolicyTables, err = parent.COSRemarkingPolicyTables(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredCOSRemarkingPolicyTables, err = parent.COSRemarkingPolicyTables(fetchFilter)
		if err != nil {
			return err
		}
	}

	COSRemarkingPolicyTable := &vspk.COSRemarkingPolicyTable{}

	if len(filteredCOSRemarkingPolicyTables) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredCOSRemarkingPolicyTables) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	COSRemarkingPolicyTable = filteredCOSRemarkingPolicyTables[0]

	d.Set("name", COSRemarkingPolicyTable.Name)
	d.Set("last_updated_by", COSRemarkingPolicyTable.LastUpdatedBy)
	d.Set("description", COSRemarkingPolicyTable.Description)
	d.Set("entity_scope", COSRemarkingPolicyTable.EntityScope)
	d.Set("external_id", COSRemarkingPolicyTable.ExternalID)

	d.Set("id", COSRemarkingPolicyTable.Identifier())
	d.Set("parent_id", COSRemarkingPolicyTable.ParentID)
	d.Set("parent_type", COSRemarkingPolicyTable.ParentType)
	d.Set("owner", COSRemarkingPolicyTable.Owner)

	d.SetId(COSRemarkingPolicyTable.Identifier())

	return nil
}
