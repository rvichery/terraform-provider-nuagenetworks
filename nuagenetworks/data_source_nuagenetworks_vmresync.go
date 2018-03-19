package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceVMResync() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVMResyncRead,
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
			"last_request_timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_time_resync_initiated": {
				Type:     schema.TypeInt,
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
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vm"},
			},
			"parent_vm": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet"},
			},
		},
	}
}

func dataSourceVMResyncRead(d *schema.ResourceData, m interface{}) error {
	filteredVMResyncs := vspk.VMResyncsList{}
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
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredVMResyncs, err = parent.VMResyncs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm"); ok {
		parent := &vspk.VM{ID: attr.(string)}
		filteredVMResyncs, err = parent.VMResyncs(fetchFilter)
		if err != nil {
			return err
		}
	}

	VMResync := &vspk.VMResync{}

	if len(filteredVMResyncs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVMResyncs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VMResync = filteredVMResyncs[0]

	d.Set("last_request_timestamp", VMResync.LastRequestTimestamp)
	d.Set("last_time_resync_initiated", VMResync.LastTimeResyncInitiated)
	d.Set("last_updated_by", VMResync.LastUpdatedBy)
	d.Set("entity_scope", VMResync.EntityScope)
	d.Set("status", VMResync.Status)
	d.Set("external_id", VMResync.ExternalID)

	d.Set("id", VMResync.Identifier())
	d.Set("parent_id", VMResync.ParentID)
	d.Set("parent_type", VMResync.ParentType)
	d.Set("owner", VMResync.Owner)

	d.SetId(VMResync.Identifier())

	return nil
}
