package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceQOS() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceQOSRead,
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
			"fip_committed_burst_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fip_committed_information_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fip_peak_burst_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fip_peak_information_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fip_rate_limiting_active": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"bum_committed_burst_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bum_committed_information_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bum_peak_burst_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bum_peak_information_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"bum_rate_limiting_active": &schema.Schema{
				Type:     schema.TypeBool,
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
			"rate_limiting_active": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"active": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"peak": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_class": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"rewrite_forwarding_class": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"egress_fip_committed_burst_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"egress_fip_committed_information_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"egress_fip_peak_burst_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"egress_fip_peak_information_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"committed_burst_size": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"committed_information_rate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"trusted_forwarding_class": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"assoc_qos_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_dscp_forwarding_class_table_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_dscp_forwarding_class_table_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"burst": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_zone": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_subnet_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_vport": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_subnet": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_zone_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_l2_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_bridge_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_l2_domain", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_host_interface", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_host_interface": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_domain_template", "parent_policy_decision"},
			},
			"parent_domain_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_policy_decision"},
			},
			"parent_policy_decision": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
		},
	}
}

func dataSourceQOSRead(d *schema.ResourceData, m interface{}) error {
	filteredQOSs := vspk.QOSsList{}
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
	if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet_template"); ok {
		parent := &vspk.SubnetTemplate{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_zone_template"); ok {
		parent := &vspk.ZoneTemplate{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_bridge_interface"); ok {
		parent := &vspk.BridgeInterface{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_host_interface"); ok {
		parent := &vspk.HostInterface{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_policy_decision"); ok {
		parent := &vspk.PolicyDecision{ID: attr.(string)}
		filteredQOSs, err = parent.QOSs(fetchFilter)
		if err != nil {
			return err
		}
	}

	QOS := &vspk.QOS{}

	if len(filteredQOSs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredQOSs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		QOS = filteredQOSs[0]
	}

	d.Set("fip_committed_burst_size", QOS.FIPCommittedBurstSize)
	d.Set("fip_committed_information_rate", QOS.FIPCommittedInformationRate)
	d.Set("fip_peak_burst_size", QOS.FIPPeakBurstSize)
	d.Set("fip_peak_information_rate", QOS.FIPPeakInformationRate)
	d.Set("fip_rate_limiting_active", QOS.FIPRateLimitingActive)
	d.Set("bum_committed_burst_size", QOS.BUMCommittedBurstSize)
	d.Set("bum_committed_information_rate", QOS.BUMCommittedInformationRate)
	d.Set("bum_peak_burst_size", QOS.BUMPeakBurstSize)
	d.Set("bum_peak_information_rate", QOS.BUMPeakInformationRate)
	d.Set("bum_rate_limiting_active", QOS.BUMRateLimitingActive)
	d.Set("name", QOS.Name)
	d.Set("last_updated_by", QOS.LastUpdatedBy)
	d.Set("rate_limiting_active", QOS.RateLimitingActive)
	d.Set("active", QOS.Active)
	d.Set("peak", QOS.Peak)
	d.Set("service_class", QOS.ServiceClass)
	d.Set("description", QOS.Description)
	d.Set("rewrite_forwarding_class", QOS.RewriteForwardingClass)
	d.Set("egress_fip_committed_burst_size", QOS.EgressFIPCommittedBurstSize)
	d.Set("egress_fip_committed_information_rate", QOS.EgressFIPCommittedInformationRate)
	d.Set("egress_fip_peak_burst_size", QOS.EgressFIPPeakBurstSize)
	d.Set("egress_fip_peak_information_rate", QOS.EgressFIPPeakInformationRate)
	d.Set("entity_scope", QOS.EntityScope)
	d.Set("committed_burst_size", QOS.CommittedBurstSize)
	d.Set("committed_information_rate", QOS.CommittedInformationRate)
	d.Set("trusted_forwarding_class", QOS.TrustedForwardingClass)
	d.Set("assoc_qos_id", QOS.AssocQosId)
	d.Set("associated_dscp_forwarding_class_table_id", QOS.AssociatedDSCPForwardingClassTableID)
	d.Set("associated_dscp_forwarding_class_table_name", QOS.AssociatedDSCPForwardingClassTableName)
	d.Set("burst", QOS.Burst)
	d.Set("external_id", QOS.ExternalID)

	d.Set("id", QOS.Identifier())
	d.Set("parent_id", QOS.ParentID)
	d.Set("parent_type", QOS.ParentType)
	d.Set("owner", QOS.Owner)

	d.SetId(QOS.Identifier())

	return nil
}
