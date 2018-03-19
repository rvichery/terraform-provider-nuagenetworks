package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceIKEPSK() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIKEPSKRead,
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
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signing_certificate_serial_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"encrypted_psk": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"encrypting_certificate_serial_number": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"unencrypted_psk": {
				Type:     schema.TypeString,
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
			"auto_created": {
				Type:     schema.TypeBool,
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

func dataSourceIKEPSKRead(d *schema.ResourceData, m interface{}) error {
	filteredIKEPSKs := vspk.IKEPSKsList{}
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
	filteredIKEPSKs, err = parent.IKEPSKs(fetchFilter)
	if err != nil {
		return err
	}

	IKEPSK := &vspk.IKEPSK{}

	if len(filteredIKEPSKs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredIKEPSKs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	IKEPSK = filteredIKEPSKs[0]

	d.Set("name", IKEPSK.Name)
	d.Set("last_updated_by", IKEPSK.LastUpdatedBy)
	d.Set("description", IKEPSK.Description)
	d.Set("signature", IKEPSK.Signature)
	d.Set("signing_certificate_serial_number", IKEPSK.SigningCertificateSerialNumber)
	d.Set("encrypted_psk", IKEPSK.EncryptedPSK)
	d.Set("encrypting_certificate_serial_number", IKEPSK.EncryptingCertificateSerialNumber)
	d.Set("unencrypted_psk", IKEPSK.UnencryptedPSK)
	d.Set("entity_scope", IKEPSK.EntityScope)
	d.Set("associated_enterprise_id", IKEPSK.AssociatedEnterpriseID)
	d.Set("auto_created", IKEPSK.AutoCreated)
	d.Set("external_id", IKEPSK.ExternalID)

	d.Set("id", IKEPSK.Identifier())
	d.Set("parent_id", IKEPSK.ParentID)
	d.Set("parent_type", IKEPSK.ParentType)
	d.Set("owner", IKEPSK.Owner)

	d.SetId(IKEPSK.Identifier())

	return nil
}
