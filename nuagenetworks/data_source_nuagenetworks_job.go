package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceJob() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceJobRead,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"result": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"progress": &schema.Schema{
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"assoc_entity_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vport": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_gateway", "parent_ns_gateway", "parent_enterprise"},
			},
			"parent_gateway": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_ns_gateway", "parent_enterprise"},
			},
			"parent_ns_gateway": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_gateway", "parent_enterprise"},
			},
			"parent_enterprise": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vport", "parent_gateway", "parent_ns_gateway"},
			},
		},
	}
}

func dataSourceJobRead(d *schema.ResourceData, m interface{}) error {
	filteredJobs := vspk.JobsList{}
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
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredJobs, err = parent.Jobs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_gateway"); ok {
		parent := &vspk.Gateway{ID: attr.(string)}
		filteredJobs, err = parent.Jobs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_ns_gateway"); ok {
		parent := &vspk.NSGateway{ID: attr.(string)}
		filteredJobs, err = parent.Jobs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_enterprise"); ok {
		parent := &vspk.Enterprise{ID: attr.(string)}
		filteredJobs, err = parent.Jobs(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredJobs, err = parent.Jobs(fetchFilter)
		if err != nil {
			return err
		}
	}

	Job := &vspk.Job{}

	if len(filteredJobs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredJobs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		Job = filteredJobs[0]
	}

	d.Set("parameters", Job.Parameters)
	d.Set("last_updated_by", Job.LastUpdatedBy)
	d.Set("result", Job.Result)
	d.Set("entity_scope", Job.EntityScope)
	d.Set("command", Job.Command)
	d.Set("progress", Job.Progress)
	d.Set("assoc_entity_type", Job.AssocEntityType)
	d.Set("status", Job.Status)
	d.Set("external_id", Job.ExternalID)

	d.Set("id", Job.Identifier())
	d.Set("parent_id", Job.ParentID)
	d.Set("parent_type", Job.ParentType)
	d.Set("owner", Job.Owner)

	d.SetId(Job.Identifier())

	return nil
}
