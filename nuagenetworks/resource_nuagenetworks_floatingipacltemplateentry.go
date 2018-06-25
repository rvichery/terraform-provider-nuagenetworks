package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceFloatingIPACLTemplateEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourceFloatingIPACLTemplateEntryCreate,
		Read:   resourceFloatingIPACLTemplateEntryRead,
		Update: resourceFloatingIPACLTemplateEntryUpdate,
		Delete: resourceFloatingIPACLTemplateEntryDelete,
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
				Optional: true,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
				Required: true,
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
				Optional: true,
				Computed: true,
			},
			"policy_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Required: true,
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
			"associated_live_entity_id": {
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
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_floating_ipacl_template": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFloatingIPACLTemplateEntryCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize FloatingIPACLTemplateEntry object
	o := &vspk.FloatingIPACLTemplateEntry{
		ACLTemplateName: d.Get("acl_template_name").(string),
		EnterpriseName:  d.Get("enterprise_name").(string),
		DomainName:      d.Get("domain_name").(string),
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
	if attr, ok := d.GetOk("dscp"); ok {
		o.DSCP = attr.(string)
	}
	if attr, ok := d.GetOk("action"); ok {
		o.Action = attr.(string)
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
	if attr, ok := d.GetOk("location_type"); ok {
		o.LocationType = attr.(string)
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
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("stateful"); ok {
		Stateful := attr.(bool)
		o.Stateful = &Stateful
	}
	if attr, ok := d.GetOk("stats_logging_enabled"); ok {
		StatsLoggingEnabled := attr.(bool)
		o.StatsLoggingEnabled = &StatsLoggingEnabled
	}
	if attr, ok := d.GetOk("ether_type"); ok {
		o.EtherType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.FloatingIPACLTemplate{ID: d.Get("parent_floating_ipacl_template").(string)}
	err := parent.CreateFloatingIPACLTemplateEntry(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceFloatingIPACLTemplateEntryRead(d, m)
}

func resourceFloatingIPACLTemplateEntryRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIPACLTemplateEntry{
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
	d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
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

func resourceFloatingIPACLTemplateEntryUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIPACLTemplateEntry{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.ACLTemplateName = d.Get("acl_template_name").(string)
	o.EnterpriseName = d.Get("enterprise_name").(string)
	o.DomainName = d.Get("domain_name").(string)

	if attr, ok := d.GetOk("icmp_code"); ok {
		o.ICMPCode = attr.(string)
	}
	if attr, ok := d.GetOk("icmp_type"); ok {
		o.ICMPType = attr.(string)
	}
	if attr, ok := d.GetOk("ipv6_address_override"); ok {
		o.IPv6AddressOverride = attr.(string)
	}
	if attr, ok := d.GetOk("dscp"); ok {
		o.DSCP = attr.(string)
	}
	if attr, ok := d.GetOk("action"); ok {
		o.Action = attr.(string)
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
	if attr, ok := d.GetOk("location_type"); ok {
		o.LocationType = attr.(string)
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
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("stateful"); ok {
		Stateful := attr.(bool)
		o.Stateful = &Stateful
	}
	if attr, ok := d.GetOk("stats_logging_enabled"); ok {
		StatsLoggingEnabled := attr.(bool)
		o.StatsLoggingEnabled = &StatsLoggingEnabled
	}
	if attr, ok := d.GetOk("ether_type"); ok {
		o.EtherType = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceFloatingIPACLTemplateEntryDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.FloatingIPACLTemplateEntry{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
