package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceIKEGatewayProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceIKEGatewayProfileCreate,
		Read:   resourceIKEGatewayProfileRead,
		Update: resourceIKEGatewayProfileUpdate,
		Delete: resourceIKEGatewayProfileDelete,
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
			"ike_gateway_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ike_gateway_identifier_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ID_IPV4_ADDR",
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"service_class": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"anti_replay_check": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ike_authentication_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ike_authentication_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ike_encryption_profile_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_ike_gateway_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceIKEGatewayProfileCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize IKEGatewayProfile object
	o := &vspk.IKEGatewayProfile{}
	if attr, ok := d.GetOk("ike_gateway_identifier"); ok {
		o.IKEGatewayIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("ike_gateway_identifier_type"); ok {
		o.IKEGatewayIdentifierType = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("service_class"); ok {
		o.ServiceClass = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("anti_replay_check"); ok {
		o.AntiReplayCheck = attr.(bool)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_authentication_id"); ok {
		o.AssociatedIKEAuthenticationID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_authentication_type"); ok {
		o.AssociatedIKEAuthenticationType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_encryption_profile_id"); ok {
		o.AssociatedIKEEncryptionProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_gateway_id"); ok {
		o.AssociatedIKEGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateIKEGatewayProfile(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceIKEGatewayProfileRead(d, m)
}

func resourceIKEGatewayProfileRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGatewayProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("ike_gateway_identifier", o.IKEGatewayIdentifier)
	d.Set("ike_gateway_identifier_type", o.IKEGatewayIdentifierType)
	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("service_class", o.ServiceClass)
	d.Set("description", o.Description)
	d.Set("anti_replay_check", o.AntiReplayCheck)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_enterprise_id", o.AssociatedEnterpriseID)
	d.Set("associated_ike_authentication_id", o.AssociatedIKEAuthenticationID)
	d.Set("associated_ike_authentication_type", o.AssociatedIKEAuthenticationType)
	d.Set("associated_ike_encryption_profile_id", o.AssociatedIKEEncryptionProfileID)
	d.Set("associated_ike_gateway_id", o.AssociatedIKEGatewayID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceIKEGatewayProfileUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGatewayProfile{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("ike_gateway_identifier"); ok {
		o.IKEGatewayIdentifier = attr.(string)
	}
	if attr, ok := d.GetOk("ike_gateway_identifier_type"); ok {
		o.IKEGatewayIdentifierType = attr.(string)
	}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("service_class"); ok {
		o.ServiceClass = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("anti_replay_check"); ok {
		o.AntiReplayCheck = attr.(bool)
	}
	if attr, ok := d.GetOk("associated_enterprise_id"); ok {
		o.AssociatedEnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_authentication_id"); ok {
		o.AssociatedIKEAuthenticationID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_authentication_type"); ok {
		o.AssociatedIKEAuthenticationType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_encryption_profile_id"); ok {
		o.AssociatedIKEEncryptionProfileID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_ike_gateway_id"); ok {
		o.AssociatedIKEGatewayID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceIKEGatewayProfileDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.IKEGatewayProfile{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
