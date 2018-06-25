package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceQOS() *schema.Resource {
	return &schema.Resource{
		Create: resourceQOSCreate,
		Read:   resourceQOSRead,
		Update: resourceQOSUpdate,
		Delete: resourceQOSDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip_committed_burst_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip_committed_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip_peak_burst_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip_peak_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fip_rate_limiting_active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"bum_committed_burst_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bum_committed_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bum_peak_burst_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bum_peak_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bum_rate_limiting_active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rate_limiting_active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"peak": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_class": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rewrite_forwarding_class": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"egress_fip_committed_burst_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egress_fip_committed_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egress_fip_peak_burst_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"egress_fip_peak_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"committed_burst_size": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"committed_information_rate": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"trusted_forwarding_class": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"assoc_qos_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_dscp_forwarding_class_table_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_dscp_forwarding_class_table_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"burst": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_zone": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
			"parent_subnet_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
			"parent_zone_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
			"parent_l2_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
			"parent_bridge_interface": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_l2_domain", "parent_host_interface", "parent_domain_template"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_host_interface", "parent_domain_template"},
			},
			"parent_host_interface": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_domain_template"},
			},
			"parent_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_subnet_template", "parent_domain", "parent_vport", "parent_subnet", "parent_zone_template", "parent_l2_domain_template", "parent_bridge_interface", "parent_l2_domain", "parent_host_interface"},
			},
		},
	}
}

func resourceQOSCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize QOS object
	o := &vspk.QOS{
		Name: d.Get("name").(string),
	}
	if attr, ok := d.GetOk("fip_committed_burst_size"); ok {
		o.FIPCommittedBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("fip_committed_information_rate"); ok {
		o.FIPCommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("fip_peak_burst_size"); ok {
		o.FIPPeakBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("fip_peak_information_rate"); ok {
		o.FIPPeakInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("fip_rate_limiting_active"); ok {
		FIPRateLimitingActive := attr.(bool)
		o.FIPRateLimitingActive = &FIPRateLimitingActive
	}
	if attr, ok := d.GetOk("bum_committed_burst_size"); ok {
		o.BUMCommittedBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("bum_committed_information_rate"); ok {
		o.BUMCommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("bum_peak_burst_size"); ok {
		o.BUMPeakBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("bum_peak_information_rate"); ok {
		o.BUMPeakInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("bum_rate_limiting_active"); ok {
		BUMRateLimitingActive := attr.(bool)
		o.BUMRateLimitingActive = &BUMRateLimitingActive
	}
	if attr, ok := d.GetOk("rate_limiting_active"); ok {
		RateLimitingActive := attr.(bool)
		o.RateLimitingActive = &RateLimitingActive
	}
	if attr, ok := d.GetOk("active"); ok {
		Active := attr.(bool)
		o.Active = &Active
	}
	if attr, ok := d.GetOk("peak"); ok {
		o.Peak = attr.(string)
	}
	if attr, ok := d.GetOk("service_class"); ok {
		o.ServiceClass = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("rewrite_forwarding_class"); ok {
		RewriteForwardingClass := attr.(bool)
		o.RewriteForwardingClass = &RewriteForwardingClass
	}
	if attr, ok := d.GetOk("egress_fip_committed_burst_size"); ok {
		o.EgressFIPCommittedBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("egress_fip_committed_information_rate"); ok {
		o.EgressFIPCommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("egress_fip_peak_burst_size"); ok {
		o.EgressFIPPeakBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("egress_fip_peak_information_rate"); ok {
		o.EgressFIPPeakInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("committed_burst_size"); ok {
		o.CommittedBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("committed_information_rate"); ok {
		o.CommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("trusted_forwarding_class"); ok {
		TrustedForwardingClass := attr.(bool)
		o.TrustedForwardingClass = &TrustedForwardingClass
	}
	if attr, ok := d.GetOk("assoc_qos_id"); ok {
		o.AssocQosId = attr.(string)
	}
	if attr, ok := d.GetOk("associated_dscp_forwarding_class_table_id"); ok {
		o.AssociatedDSCPForwardingClassTableID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_dscp_forwarding_class_table_name"); ok {
		o.AssociatedDSCPForwardingClassTableName = attr.(string)
	}
	if attr, ok := d.GetOk("burst"); ok {
		o.Burst = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet_template"); ok {
		parent := &vspk.SubnetTemplate{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_zone_template"); ok {
		parent := &vspk.ZoneTemplate{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain_template"); ok {
		parent := &vspk.L2DomainTemplate{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_bridge_interface"); ok {
		parent := &vspk.BridgeInterface{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_host_interface"); ok {
		parent := &vspk.HostInterface{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		err := parent.CreateQOS(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceQOSRead(d, m)
}

func resourceQOSRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.QOS{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("fip_committed_burst_size", o.FIPCommittedBurstSize)
	d.Set("fip_committed_information_rate", o.FIPCommittedInformationRate)
	d.Set("fip_peak_burst_size", o.FIPPeakBurstSize)
	d.Set("fip_peak_information_rate", o.FIPPeakInformationRate)
	d.Set("fip_rate_limiting_active", o.FIPRateLimitingActive)
	d.Set("bum_committed_burst_size", o.BUMCommittedBurstSize)
	d.Set("bum_committed_information_rate", o.BUMCommittedInformationRate)
	d.Set("bum_peak_burst_size", o.BUMPeakBurstSize)
	d.Set("bum_peak_information_rate", o.BUMPeakInformationRate)
	d.Set("bum_rate_limiting_active", o.BUMRateLimitingActive)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("rate_limiting_active", o.RateLimitingActive)
	d.Set("active", o.Active)
	d.Set("peak", o.Peak)
	d.Set("service_class", o.ServiceClass)
	d.Set("description", o.Description)
	d.Set("rewrite_forwarding_class", o.RewriteForwardingClass)
	d.Set("egress_fip_committed_burst_size", o.EgressFIPCommittedBurstSize)
	d.Set("egress_fip_committed_information_rate", o.EgressFIPCommittedInformationRate)
	d.Set("egress_fip_peak_burst_size", o.EgressFIPPeakBurstSize)
	d.Set("egress_fip_peak_information_rate", o.EgressFIPPeakInformationRate)
	d.Set("entity_scope", o.EntityScope)
	d.Set("committed_burst_size", o.CommittedBurstSize)
	d.Set("committed_information_rate", o.CommittedInformationRate)
	d.Set("trusted_forwarding_class", o.TrustedForwardingClass)
	d.Set("assoc_qos_id", o.AssocQosId)
	d.Set("associated_dscp_forwarding_class_table_id", o.AssociatedDSCPForwardingClassTableID)
	d.Set("associated_dscp_forwarding_class_table_name", o.AssociatedDSCPForwardingClassTableName)
	d.Set("burst", o.Burst)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceQOSUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.QOS{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("fip_committed_burst_size"); ok {
		o.FIPCommittedBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("fip_committed_information_rate"); ok {
		o.FIPCommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("fip_peak_burst_size"); ok {
		o.FIPPeakBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("fip_peak_information_rate"); ok {
		o.FIPPeakInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("fip_rate_limiting_active"); ok {
		FIPRateLimitingActive := attr.(bool)
		o.FIPRateLimitingActive = &FIPRateLimitingActive
	}
	if attr, ok := d.GetOk("bum_committed_burst_size"); ok {
		o.BUMCommittedBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("bum_committed_information_rate"); ok {
		o.BUMCommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("bum_peak_burst_size"); ok {
		o.BUMPeakBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("bum_peak_information_rate"); ok {
		o.BUMPeakInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("bum_rate_limiting_active"); ok {
		BUMRateLimitingActive := attr.(bool)
		o.BUMRateLimitingActive = &BUMRateLimitingActive
	}
	if attr, ok := d.GetOk("rate_limiting_active"); ok {
		RateLimitingActive := attr.(bool)
		o.RateLimitingActive = &RateLimitingActive
	}
	if attr, ok := d.GetOk("active"); ok {
		Active := attr.(bool)
		o.Active = &Active
	}
	if attr, ok := d.GetOk("peak"); ok {
		o.Peak = attr.(string)
	}
	if attr, ok := d.GetOk("service_class"); ok {
		o.ServiceClass = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("rewrite_forwarding_class"); ok {
		RewriteForwardingClass := attr.(bool)
		o.RewriteForwardingClass = &RewriteForwardingClass
	}
	if attr, ok := d.GetOk("egress_fip_committed_burst_size"); ok {
		o.EgressFIPCommittedBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("egress_fip_committed_information_rate"); ok {
		o.EgressFIPCommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("egress_fip_peak_burst_size"); ok {
		o.EgressFIPPeakBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("egress_fip_peak_information_rate"); ok {
		o.EgressFIPPeakInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("committed_burst_size"); ok {
		o.CommittedBurstSize = attr.(string)
	}
	if attr, ok := d.GetOk("committed_information_rate"); ok {
		o.CommittedInformationRate = attr.(string)
	}
	if attr, ok := d.GetOk("trusted_forwarding_class"); ok {
		TrustedForwardingClass := attr.(bool)
		o.TrustedForwardingClass = &TrustedForwardingClass
	}
	if attr, ok := d.GetOk("assoc_qos_id"); ok {
		o.AssocQosId = attr.(string)
	}
	if attr, ok := d.GetOk("associated_dscp_forwarding_class_table_id"); ok {
		o.AssociatedDSCPForwardingClassTableID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_dscp_forwarding_class_table_name"); ok {
		o.AssociatedDSCPForwardingClassTableName = attr.(string)
	}
	if attr, ok := d.GetOk("burst"); ok {
		o.Burst = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceQOSDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.QOS{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
