package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceIKEGatewayProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIKEGatewayProfileRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ike_gateway_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ike_gateway_identifier_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_class": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"anti_replay_check": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ike_authentication_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ike_authentication_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ike_encryption_profile_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_ike_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_enterprise": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceIKEGatewayProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredIKEGatewayProfiles := vspk.IKEGatewayProfilesList{}
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
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	filteredIKEGatewayProfiles, err = parent.IKEGatewayProfiles(fetchFilter)
	if err != nil {
		return err
	}

	IKEGatewayProfile := &vspk.IKEGatewayProfile{}

	if len(filteredIKEGatewayProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIKEGatewayProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	IKEGatewayProfile = filteredIKEGatewayProfiles[0]

	d.Set("ike_gateway_identifier", IKEGatewayProfile.IKEGatewayIdentifier)
	d.Set("ike_gateway_identifier_type", IKEGatewayProfile.IKEGatewayIdentifierType)
	d.Set("name", IKEGatewayProfile.Name)
	d.Set("last_updated_by", IKEGatewayProfile.LastUpdatedBy)
	d.Set("service_class", IKEGatewayProfile.ServiceClass)
	d.Set("description", IKEGatewayProfile.Description)
	d.Set("anti_replay_check", IKEGatewayProfile.AntiReplayCheck)
	d.Set("entity_scope", IKEGatewayProfile.EntityScope)
	d.Set("associated_enterprise_id", IKEGatewayProfile.AssociatedEnterpriseID)
	d.Set("associated_ike_authentication_id", IKEGatewayProfile.AssociatedIKEAuthenticationID)
	d.Set("associated_ike_authentication_type", IKEGatewayProfile.AssociatedIKEAuthenticationType)
	d.Set("associated_ike_encryption_profile_id", IKEGatewayProfile.AssociatedIKEEncryptionProfileID)
	d.Set("associated_ike_gateway_id", IKEGatewayProfile.AssociatedIKEGatewayID)
	d.Set("external_id", IKEGatewayProfile.ExternalID)

	d.Set("id", IKEGatewayProfile.Identifier())
	d.Set("parent_id", IKEGatewayProfile.ParentID)
	d.Set("parent_type", IKEGatewayProfile.ParentType)
	d.Set("owner", IKEGatewayProfile.Owner)

	d.SetId(IKEGatewayProfile.Identifier())

	return nil
}
