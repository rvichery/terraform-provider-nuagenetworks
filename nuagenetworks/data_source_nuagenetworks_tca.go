package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceTCA() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTCARead,
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
			"url_end_point": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_policy_group_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"period": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"metric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"throttle_time": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"disable": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"display_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_container_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_zone": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_zone", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_vport": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_zone", "parent_domain", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_subnet": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_zone", "parent_domain", "parent_vport", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_vm_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_bridge_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_l2_domain", "parent_host_interface"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_host_interface"},
			},
			"parent_host_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_vm_interface", "parent_bridge_interface", "parent_l2_domain"},
			},
		},
	}
}

func dataSourceTCARead(d *schema.ResourceData, m interface{}) error {
	filteredTCAs := vspk.TCAsList{}
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
	if attr, ok := d.GetOk("parent_container_interface"); ok {
		parent := &vspk.ContainerInterface{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm_interface"); ok {
		parent := &vspk.VMInterface{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
		parent := &vspk.BridgeInterface{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_host_interface"); ok {
		parent := &vspk.HostInterface{ID: attr.(string)}
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	} else {
		parent := m.(*vspk.Me)
		filteredTCAs, err = parent.TCAs(fetchFilter)
		if err != nil {
			return err
		}
	}

	TCA := &vspk.TCA{}

	if len(filteredTCAs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredTCAs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		TCA = filteredTCAs[0]
	}

	d.Set("url_end_point", TCA.URLEndPoint)
	d.Set("name", TCA.Name)
	d.Set("target_policy_group_id", TCA.TargetPolicyGroupID)
	d.Set("last_updated_by", TCA.LastUpdatedBy)
	d.Set("action", TCA.Action)
	d.Set("period", TCA.Period)
	d.Set("description", TCA.Description)
	d.Set("metric", TCA.Metric)
	d.Set("threshold", TCA.Threshold)
	d.Set("throttle_time", TCA.ThrottleTime)
	d.Set("disable", TCA.Disable)
	d.Set("display_status", TCA.DisplayStatus)
	d.Set("entity_scope", TCA.EntityScope)

	d.Set("status", TCA.Status)
	d.Set("external_id", TCA.ExternalID)
	d.Set("type", TCA.Type)

	d.Set("id", TCA.Identifier())
	d.Set("parent_id", TCA.ParentID)
	d.Set("parent_type", TCA.ParentType)
	d.Set("owner", TCA.Owner)

	d.SetId(TCA.Identifier())

	return nil
}
