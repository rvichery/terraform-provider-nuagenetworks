package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceEgressACLEntryTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceEgressACLEntryTemplateCreate,
		Read:   resourceEgressACLEntryTemplateRead,
		Update: resourceEgressACLEntryTemplateUpdate,
		Delete: resourceEgressACLEntryTemplateDelete,
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
			"acl_template_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_code": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icmp_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ipv6_address_override": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dscp": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address_override": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"destination_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mirror_destination_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"flow_logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"enterprise_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policy_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"source_port": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_l7_application_signature_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_live_entity_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_traffic_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_traffic_type_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stateful": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"stats_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"stats_logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ether_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_egress_acl_template": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEgressACLEntryTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize EgressACLEntryTemplate object
	o := &vspk.EgressACLEntryTemplate{
		DSCP:         d.Get("dscp").(string),
		Action:       d.Get("action").(string),
		LocationType: d.Get("location_type").(string),
		EtherType:    d.Get("ether_type").(string),
	}
	if attr, ok := d.GetOk("acl_template_name"); ok {
		o.ACLTemplateName = attr.(string)
	}
	if attr, ok := d.GetOk("icmp_code"); ok {
		o.ICMPCode = attr.(string)
	}
	if attr, ok := d.GetOk("icmp_type"); ok {
		o.ICMPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address_override"); ok {
		o.IPv6AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("address_override"); ok {
		o.AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("destination_port"); ok {
		o.DestinationPort = attr.(string)
	}
	if attr, ok := d.GetOk("network_id"); ok {
		o.NetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("network_type"); ok {
		o.NetworkType = attr.(string)
	}
	if attr, ok := d.GetOk("mirror_destination_id"); ok {
		o.MirrorDestinationID = attr.(string)
	}
	if attr, ok := d.GetOk("flow_logging_enabled"); ok {
		FlowLoggingEnabled := attr.(bool)
		o.FlowLoggingEnabled = &FlowLoggingEnabled
	}
	if attr, ok := d.GetOk("location_id"); ok {
		o.LocationID = attr.(string)
	}
	if attr, ok := d.GetOk("policy_state"); ok {
		o.PolicyState = attr.(string)
	}
	if attr, ok := d.GetOk("source_port"); ok {
		o.SourcePort = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		Priority := attr.(int)
		o.Priority = &Priority
	}
	if attr, ok := d.GetOk("protocol"); ok {
		o.Protocol = attr.(string)
	}
	if attr, ok := d.GetOk("associated_l7_application_signature_id"); ok {
		o.AssociatedL7ApplicationSignatureID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_traffic_type"); ok {
		o.AssociatedTrafficType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_traffic_type_id"); ok {
		o.AssociatedTrafficTypeID = attr.(string)
	}
	if attr, ok := d.GetOk("stateful"); ok {
		Stateful := attr.(bool)
		o.Stateful = &Stateful
	}
	if attr, ok := d.GetOk("stats_logging_enabled"); ok {
		StatsLoggingEnabled := attr.(bool)
		o.StatsLoggingEnabled = &StatsLoggingEnabled
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.EgressACLTemplate{ID: d.Get("parent_egress_acl_template").(string)}
	err := parent.CreateEgressACLEntryTemplate(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceEgressACLEntryTemplateRead(d, m)
}

func resourceEgressACLEntryTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressACLEntryTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("acl_template_name", o.ACLTemplateName)
	d.Set("icmp_code", o.ICMPCode)
	d.Set("icmp_type", o.ICMPType)
	d.Set("ipv6_address_override", o.IPv6AddressOverride)
	d.Set("dscp", o.DSCP)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("action", o.Action)
	d.Set("address_override", o.AddressOverride)
	d.Set("description", o.Description)
	d.Set("destination_port", o.DestinationPort)
	d.Set("network_id", o.NetworkID)
	d.Set("network_type", o.NetworkType)
	d.Set("mirror_destination_id", o.MirrorDestinationID)
	d.Set("flow_logging_enabled", o.FlowLoggingEnabled)
	d.Set("enterprise_name", o.EnterpriseName)
	d.Set("entity_scope", o.EntityScope)
	d.Set("location_id", o.LocationID)
	d.Set("location_type", o.LocationType)
	d.Set("policy_state", o.PolicyState)
	d.Set("domain_name", o.DomainName)
	d.Set("source_port", o.SourcePort)
	d.Set("priority", o.Priority)
	d.Set("protocol", o.Protocol)
	d.Set("associated_l7_application_signature_id", o.AssociatedL7ApplicationSignatureID)
	d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
	d.Set("associated_traffic_type", o.AssociatedTrafficType)
	d.Set("associated_traffic_type_id", o.AssociatedTrafficTypeID)
	d.Set("stateful", o.Stateful)
	d.Set("stats_id", o.StatsID)
	d.Set("stats_logging_enabled", o.StatsLoggingEnabled)
	d.Set("ether_type", o.EtherType)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceEgressACLEntryTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressACLEntryTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.DSCP = d.Get("dscp").(string)
	o.Action = d.Get("action").(string)
	o.LocationType = d.Get("location_type").(string)
	o.EtherType = d.Get("ether_type").(string)

	if attr, ok := d.GetOk("acl_template_name"); ok {
		o.ACLTemplateName = attr.(string)
	}
	if attr, ok := d.GetOk("icmp_code"); ok {
		o.ICMPCode = attr.(string)
	}
	if attr, ok := d.GetOk("icmp_type"); ok {
		o.ICMPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address_override"); ok {
		o.IPv6AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("address_override"); ok {
		o.AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("destination_port"); ok {
		o.DestinationPort = attr.(string)
	}
	if attr, ok := d.GetOk("network_id"); ok {
		o.NetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("network_type"); ok {
		o.NetworkType = attr.(string)
	}
	if attr, ok := d.GetOk("mirror_destination_id"); ok {
		o.MirrorDestinationID = attr.(string)
	}
	if attr, ok := d.GetOk("flow_logging_enabled"); ok {
		FlowLoggingEnabled := attr.(bool)
		o.FlowLoggingEnabled = &FlowLoggingEnabled
	}
	if attr, ok := d.GetOk("location_id"); ok {
		o.LocationID = attr.(string)
	}
	if attr, ok := d.GetOk("policy_state"); ok {
		o.PolicyState = attr.(string)
	}
	if attr, ok := d.GetOk("source_port"); ok {
		o.SourcePort = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		Priority := attr.(int)
		o.Priority = &Priority
	}
	if attr, ok := d.GetOk("protocol"); ok {
		o.Protocol = attr.(string)
	}
	if attr, ok := d.GetOk("associated_l7_application_signature_id"); ok {
		o.AssociatedL7ApplicationSignatureID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_traffic_type"); ok {
		o.AssociatedTrafficType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_traffic_type_id"); ok {
		o.AssociatedTrafficTypeID = attr.(string)
	}
	if attr, ok := d.GetOk("stateful"); ok {
		Stateful := attr.(bool)
		o.Stateful = &Stateful
	}
	if attr, ok := d.GetOk("stats_logging_enabled"); ok {
		StatsLoggingEnabled := attr.(bool)
		o.StatsLoggingEnabled = &StatsLoggingEnabled
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceEgressACLEntryTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressACLEntryTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
