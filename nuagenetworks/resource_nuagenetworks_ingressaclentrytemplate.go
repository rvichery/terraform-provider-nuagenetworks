package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceIngressACLEntryTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceIngressACLEntryTemplateCreate,
		Read:   resourceIngressACLEntryTemplateRead,
		Update: resourceIngressACLEntryTemplateUpdate,
		Delete: resourceIngressACLEntryTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"acl_template_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icmp_code": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"icmp_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address_override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mirror_destination_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"flow_logging_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enterprise_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"policy_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_l7_application_signature_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_live_entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_traffic_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_traffic_type_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"stateful": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"stats_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stats_logging_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ether_type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"overlay_mirror_destination_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_l2_domain", "parent_ingress_acl_template"},
			},
			"parent_l2_domain": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_ingress_acl_template"},
			},
			"parent_ingress_acl_template": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_l2_domain"},
			},
		},
	}
}

func resourceIngressACLEntryTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize IngressACLEntryTemplate object
	o := &vspk.IngressACLEntryTemplate{
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
		o.FlowLoggingEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("location_id"); ok {
		o.LocationID = attr.(string)
	}
	if attr, ok := d.GetOk("policy_state"); ok {
		o.PolicyState = attr.(string)
	}
	if attr, ok := d.GetOk("domain_name"); ok {
		o.DomainName = attr.(string)
	}
	if attr, ok := d.GetOk("source_port"); ok {
		o.SourcePort = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
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
		o.Stateful = attr.(bool)
	}
	if attr, ok := d.GetOk("stats_logging_enabled"); ok {
		o.StatsLoggingEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("overlay_mirror_destination_id"); ok {
		o.OverlayMirrorDestinationID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateIngressACLEntryTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		err := parent.CreateIngressACLEntryTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateIngressACLEntryTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_ingress_acl_template"); ok {
		parent := &vspk.IngressACLTemplate{ID: attr.(string)}
		err := parent.CreateIngressACLEntryTemplate(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceIngressACLEntryTemplateRead(d, m)
}

func resourceIngressACLEntryTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IngressACLEntryTemplate{
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
	d.Set("overlay_mirror_destination_id", o.OverlayMirrorDestinationID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceIngressACLEntryTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IngressACLEntryTemplate{
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
		o.FlowLoggingEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("location_id"); ok {
		o.LocationID = attr.(string)
	}
	if attr, ok := d.GetOk("policy_state"); ok {
		o.PolicyState = attr.(string)
	}
	if attr, ok := d.GetOk("domain_name"); ok {
		o.DomainName = attr.(string)
	}
	if attr, ok := d.GetOk("source_port"); ok {
		o.SourcePort = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
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
		o.Stateful = attr.(bool)
	}
	if attr, ok := d.GetOk("stats_logging_enabled"); ok {
		o.StatsLoggingEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("overlay_mirror_destination_id"); ok {
		o.OverlayMirrorDestinationID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceIngressACLEntryTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IngressACLEntryTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
