package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceEgressAdvFwdEntryTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceEgressAdvFwdEntryTemplateCreate,
		Read:   resourceEgressAdvFwdEntryTemplateRead,
		Update: resourceEgressAdvFwdEntryTemplateUpdate,
		Delete: resourceEgressAdvFwdEntryTemplateDelete,
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
				Required: true,
			},
			"icmp_code": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"icmp_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fc_override": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ipv6_address_override": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dscp": {
				Type:     schema.TypeString,
				Required: true,
			},
			"failsafe_datapath": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "FAIL_TO_BLOCK",
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address_override": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"redirect_vport_tag_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"destination_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"mirror_destination_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"flow_logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enterprise_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"location_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"location_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"policy_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"source_port": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"uplink_preference": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_live_entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"stats_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"stats_logging_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"ether_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_egress_adv_fwd_template": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEgressAdvFwdEntryTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize EgressAdvFwdEntryTemplate object
	o := &vspk.EgressAdvFwdEntryTemplate{
		ACLTemplateName: d.Get("acl_template_name").(string),
		DSCP:            d.Get("dscp").(string),
		Name:            d.Get("name").(string),
		Action:          d.Get("action").(string),
		LocationType:    d.Get("location_type").(string),
		EtherType:       d.Get("ether_type").(string),
	}
	if attr, ok := d.GetOk("icmp_code"); ok {
		o.ICMPCode = attr.(string)
	}
	if attr, ok := d.GetOk("icmp_type"); ok {
		o.ICMPType = attr.(string)
	}
	if attr, ok := d.GetOk("fc_override"); ok {
		o.FCOverride = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address_override"); ok {
		o.IPv6AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("failsafe_datapath"); ok {
		o.FailsafeDatapath = attr.(string)
	}
	if attr, ok := d.GetOk("address_override"); ok {
		o.AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("redirect_vport_tag_id"); ok {
		o.RedirectVPortTagID = attr.(string)
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
	if attr, ok := d.GetOk("source_port"); ok {
		o.SourcePort = attr.(string)
	}
	if attr, ok := d.GetOk("uplink_preference"); ok {
		o.UplinkPreference = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("protocol"); ok {
		o.Protocol = attr.(string)
	}
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("stats_logging_enabled"); ok {
		o.StatsLoggingEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.EgressAdvFwdTemplate{ID: d.Get("parent_egress_adv_fwd_template").(string)}
	err := parent.CreateEgressAdvFwdEntryTemplate(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceEgressAdvFwdEntryTemplateRead(d, m)
}

func resourceEgressAdvFwdEntryTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressAdvFwdEntryTemplate{
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
	d.Set("fc_override", o.FCOverride)
	d.Set("ipv6_address_override", o.IPv6AddressOverride)
	d.Set("dscp", o.DSCP)
	d.Set("failsafe_datapath", o.FailsafeDatapath)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("action", o.Action)
	d.Set("address_override", o.AddressOverride)
	d.Set("redirect_vport_tag_id", o.RedirectVPortTagID)
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
	d.Set("uplink_preference", o.UplinkPreference)
	d.Set("priority", o.Priority)
	d.Set("protocol", o.Protocol)
	d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
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

func resourceEgressAdvFwdEntryTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressAdvFwdEntryTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.ACLTemplateName = d.Get("acl_template_name").(string)
	o.DSCP = d.Get("dscp").(string)
	o.Name = d.Get("name").(string)
	o.Action = d.Get("action").(string)
	o.LocationType = d.Get("location_type").(string)
	o.EtherType = d.Get("ether_type").(string)

	if attr, ok := d.GetOk("icmp_code"); ok {
		o.ICMPCode = attr.(string)
	}
	if attr, ok := d.GetOk("icmp_type"); ok {
		o.ICMPType = attr.(string)
	}
	if attr, ok := d.GetOk("fc_override"); ok {
		o.FCOverride = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address_override"); ok {
		o.IPv6AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("failsafe_datapath"); ok {
		o.FailsafeDatapath = attr.(string)
	}
	if attr, ok := d.GetOk("address_override"); ok {
		o.AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("redirect_vport_tag_id"); ok {
		o.RedirectVPortTagID = attr.(string)
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
	if attr, ok := d.GetOk("source_port"); ok {
		o.SourcePort = attr.(string)
	}
	if attr, ok := d.GetOk("uplink_preference"); ok {
		o.UplinkPreference = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("protocol"); ok {
		o.Protocol = attr.(string)
	}
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("stats_logging_enabled"); ok {
		o.StatsLoggingEnabled = attr.(bool)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceEgressAdvFwdEntryTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.EgressAdvFwdEntryTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
