package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourcePolicyDecision() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePolicyDecisionRead,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"egress_acls": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"egress_qos": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fip_acls": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ingress_acls": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ingress_adv_fwd": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ingress_external_service_acls": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"qos": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"stats": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_container_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vm_interface", "parent_bridge_interface", "parent_host_interface"},
			},
			"parent_vm_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_bridge_interface", "parent_host_interface"},
			},
			"parent_bridge_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_vm_interface", "parent_host_interface"},
			},
			"parent_host_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container_interface", "parent_vm_interface", "parent_bridge_interface"},
			},
		},
	}
}

func dataSourcePolicyDecisionRead(d *schema.ResourceData, m interface{}) error {
	filteredPolicyDecisions := vspk.PolicyDecisionsList{}
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
		filteredPolicyDecisions, err = parent.PolicyDecisions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vm_interface"); ok {
		parent := &vspk.VMInterface{ID: attr.(string)}
		filteredPolicyDecisions, err = parent.PolicyDecisions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
		parent := &vspk.BridgeInterface{ID: attr.(string)}
		filteredPolicyDecisions, err = parent.PolicyDecisions(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_host_interface"); ok {
		parent := &vspk.HostInterface{ID: attr.(string)}
		filteredPolicyDecisions, err = parent.PolicyDecisions(fetchFilter)
		if err != nil {
			return err
		}
	}

	PolicyDecision := &vspk.PolicyDecision{}

	if len(filteredPolicyDecisions) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredPolicyDecisions) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		PolicyDecision = filteredPolicyDecisions[0]
	}

	d.Set("last_updated_by", PolicyDecision.LastUpdatedBy)
	d.Set("egress_acls", PolicyDecision.EgressACLs)
	d.Set("egress_qos", PolicyDecision.EgressQos)
	d.Set("fip_acls", PolicyDecision.FipACLs)
	d.Set("ingress_acls", PolicyDecision.IngressACLs)
	d.Set("ingress_adv_fwd", PolicyDecision.IngressAdvFwd)
	d.Set("ingress_external_service_acls", PolicyDecision.IngressExternalServiceACLs)
	d.Set("entity_scope", PolicyDecision.EntityScope)
	d.Set("qos", PolicyDecision.Qos)
	d.Set("stats", PolicyDecision.Stats)
	d.Set("external_id", PolicyDecision.ExternalID)

	d.Set("id", PolicyDecision.Identifier())
	d.Set("parent_id", PolicyDecision.ParentID)
	d.Set("parent_type", PolicyDecision.ParentType)
	d.Set("owner", PolicyDecision.Owner)

	d.SetId(PolicyDecision.Identifier())

	return nil
}
