package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceContainerResync() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceContainerResyncRead,
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
				ConflictsWith: []string{"parent_container"},
			},
			"parent_container": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet"},
			},
		},
	}
}

func dataSourceContainerResyncRead(d *schema.ResourceData, m interface{}) error {
	filteredContainerResyncs := vspk.ContainerResyncsList{}
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
		filteredContainerResyncs, err = parent.ContainerResyncs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_container"); ok {
		parent := &vspk.Container{ID: attr.(string)}
		filteredContainerResyncs, err = parent.ContainerResyncs(fetchFilter)
		if err != nil {
			return err
		}
	}

	ContainerResync := &vspk.ContainerResync{}

	if len(filteredContainerResyncs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredContainerResyncs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	ContainerResync = filteredContainerResyncs[0]

	d.Set("last_request_timestamp", ContainerResync.LastRequestTimestamp)
	d.Set("last_time_resync_initiated", ContainerResync.LastTimeResyncInitiated)
	d.Set("last_updated_by", ContainerResync.LastUpdatedBy)
	d.Set("entity_scope", ContainerResync.EntityScope)
	d.Set("status", ContainerResync.Status)
	d.Set("external_id", ContainerResync.ExternalID)

	d.Set("id", ContainerResync.Identifier())
	d.Set("parent_id", ContainerResync.ParentID)
	d.Set("parent_type", ContainerResync.ParentType)
	d.Set("owner", ContainerResync.Owner)

	d.SetId(ContainerResync.Identifier())

	return nil
}
