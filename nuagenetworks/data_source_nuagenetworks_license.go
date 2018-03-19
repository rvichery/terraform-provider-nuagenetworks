package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceLicense() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceLicenseRead,
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
			"major_release": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_supported_versions": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"phone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_encryption": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_entities": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"license_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"license_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"minor_release": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"zip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"city": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowed_avrsgs_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allowed_avrss_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allowed_cpes_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allowed_nics_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allowed_vms_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allowed_vrsgs_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"allowed_vrss_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"encryption_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"unique_license_identifier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"company": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"country": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_cluster_license": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"street": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customer_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"expiration_date": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"expiry_timestamp": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceLicenseRead(d *schema.ResourceData, m interface{}) error {
	filteredLicenses := vspk.LicensesList{}
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
	parent := m.(*vspk.Me)
	filteredLicenses, err = parent.Licenses(fetchFilter)
	if err != nil {
		return err
	}

	License := &vspk.License{}

	if len(filteredLicenses) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredLicenses) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	License = filteredLicenses[0]

	d.Set("major_release", License.MajorRelease)
	d.Set("last_updated_by", License.LastUpdatedBy)
	d.Set("additional_supported_versions", License.AdditionalSupportedVersions)
	d.Set("phone", License.Phone)
	d.Set("license", License.License)
	d.Set("license_encryption", License.LicenseEncryption)
	d.Set("license_entities", License.LicenseEntities)
	d.Set("license_id", License.LicenseID)
	d.Set("license_type", License.LicenseType)
	d.Set("minor_release", License.MinorRelease)
	d.Set("zip", License.Zip)
	d.Set("city", License.City)
	d.Set("allowed_avrsgs_count", License.AllowedAVRSGsCount)
	d.Set("allowed_avrss_count", License.AllowedAVRSsCount)
	d.Set("allowed_cpes_count", License.AllowedCPEsCount)
	d.Set("allowed_nics_count", License.AllowedNICsCount)
	d.Set("allowed_vms_count", License.AllowedVMsCount)
	d.Set("allowed_vrsgs_count", License.AllowedVRSGsCount)
	d.Set("allowed_vrss_count", License.AllowedVRSsCount)
	d.Set("email", License.Email)
	d.Set("encryption_mode", License.EncryptionMode)
	d.Set("unique_license_identifier", License.UniqueLicenseIdentifier)
	d.Set("entity_scope", License.EntityScope)
	d.Set("company", License.Company)
	d.Set("country", License.Country)
	d.Set("product_version", License.ProductVersion)

	d.Set("is_cluster_license", License.IsClusterLicense)
	d.Set("user_name", License.UserName)
	d.Set("state", License.State)
	d.Set("street", License.Street)
	d.Set("customer_key", License.CustomerKey)
	d.Set("expiration_date", License.ExpirationDate)
	d.Set("expiry_timestamp", License.ExpiryTimestamp)
	d.Set("external_id", License.ExternalID)

	d.Set("id", License.Identifier())
	d.Set("parent_id", License.ParentID)
	d.Set("parent_type", License.ParentType)
	d.Set("owner", License.Owner)

	d.SetId(License.Identifier())

	return nil
}
