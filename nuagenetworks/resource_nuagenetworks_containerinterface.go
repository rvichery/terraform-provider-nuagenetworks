package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceContainerInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceContainerInterfaceCreate,
		Read:   resourceContainerInterfaceRead,
		Update: resourceContainerInterfaceUpdate,
		Delete: resourceContainerInterfaceDelete,
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
			"mac": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vport_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vport_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"endpoint_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"policy_decision_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"zone_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"container_uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_floating_ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attached_network_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"attached_network_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"multi_nic_vport_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_container": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceContainerInterfaceCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize ContainerInterface object
	o := &vspk.ContainerInterface{}
	if attr, ok := d.GetOk("mac"); ok {
		o.MAC = attr.(string)
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("vport_id"); ok {
		o.VPortID = attr.(string)
	}
	if attr, ok := d.GetOk("vport_name"); ok {
		o.VPortName = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("network_id"); ok {
		o.NetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("network_name"); ok {
		o.NetworkName = attr.(string)
	}
	if attr, ok := d.GetOk("tier_id"); ok {
		o.TierID = attr.(string)
	}
	if attr, ok := d.GetOk("endpoint_id"); ok {
		o.EndpointID = attr.(string)
	}
	if attr, ok := d.GetOk("policy_decision_id"); ok {
		o.PolicyDecisionID = attr.(string)
	}
	if attr, ok := d.GetOk("domain_id"); ok {
		o.DomainID = attr.(string)
	}
	if attr, ok := d.GetOk("domain_name"); ok {
		o.DomainName = attr.(string)
	}
	if attr, ok := d.GetOk("zone_id"); ok {
		o.ZoneID = attr.(string)
	}
	if attr, ok := d.GetOk("zone_name"); ok {
		o.ZoneName = attr.(string)
	}
	if attr, ok := d.GetOk("container_uuid"); ok {
		o.ContainerUUID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_floating_ip_address"); ok {
		o.AssociatedFloatingIPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("attached_network_id"); ok {
		o.AttachedNetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("attached_network_type"); ok {
		o.AttachedNetworkType = attr.(string)
	}
	if attr, ok := d.GetOk("multi_nic_vport_name"); ok {
		o.MultiNICVPortName = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Container{ID: d.Get("parent_container").(string)}
	err := parent.CreateContainerInterface(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceContainerInterfaceRead(d, m)
}

func resourceContainerInterfaceRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ContainerInterface{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("mac", o.MAC)
	d.Set("ip_address", o.IPAddress)
	d.Set("vport_id", o.VPortID)
	d.Set("vport_name", o.VPortName)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("gateway", o.Gateway)
	d.Set("netmask", o.Netmask)
	d.Set("network_id", o.NetworkID)
	d.Set("network_name", o.NetworkName)
	d.Set("tier_id", o.TierID)
	d.Set("endpoint_id", o.EndpointID)
	d.Set("entity_scope", o.EntityScope)
	d.Set("policy_decision_id", o.PolicyDecisionID)
	d.Set("domain_id", o.DomainID)
	d.Set("domain_name", o.DomainName)
	d.Set("zone_id", o.ZoneID)
	d.Set("zone_name", o.ZoneName)
	d.Set("container_uuid", o.ContainerUUID)
	d.Set("associated_floating_ip_address", o.AssociatedFloatingIPAddress)
	d.Set("attached_network_id", o.AttachedNetworkID)
	d.Set("attached_network_type", o.AttachedNetworkType)
	d.Set("multi_nic_vport_name", o.MultiNICVPortName)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceContainerInterfaceUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ContainerInterface{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("mac"); ok {
		o.MAC = attr.(string)
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("vport_id"); ok {
		o.VPortID = attr.(string)
	}
	if attr, ok := d.GetOk("vport_name"); ok {
		o.VPortName = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("network_id"); ok {
		o.NetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("network_name"); ok {
		o.NetworkName = attr.(string)
	}
	if attr, ok := d.GetOk("tier_id"); ok {
		o.TierID = attr.(string)
	}
	if attr, ok := d.GetOk("endpoint_id"); ok {
		o.EndpointID = attr.(string)
	}
	if attr, ok := d.GetOk("policy_decision_id"); ok {
		o.PolicyDecisionID = attr.(string)
	}
	if attr, ok := d.GetOk("domain_id"); ok {
		o.DomainID = attr.(string)
	}
	if attr, ok := d.GetOk("domain_name"); ok {
		o.DomainName = attr.(string)
	}
	if attr, ok := d.GetOk("zone_id"); ok {
		o.ZoneID = attr.(string)
	}
	if attr, ok := d.GetOk("zone_name"); ok {
		o.ZoneName = attr.(string)
	}
	if attr, ok := d.GetOk("container_uuid"); ok {
		o.ContainerUUID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_floating_ip_address"); ok {
		o.AssociatedFloatingIPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("attached_network_id"); ok {
		o.AttachedNetworkID = attr.(string)
	}
	if attr, ok := d.GetOk("attached_network_type"); ok {
		o.AttachedNetworkType = attr.(string)
	}
	if attr, ok := d.GetOk("multi_nic_vport_name"); ok {
		o.MultiNICVPortName = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceContainerInterfaceDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ContainerInterface{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
