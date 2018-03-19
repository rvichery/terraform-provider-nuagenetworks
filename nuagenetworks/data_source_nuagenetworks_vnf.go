package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVNF() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNFRead,
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
			"vnf_descriptor_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vnf_descriptor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"nsg_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsg_system_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ns_gateway_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_known_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_mb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowed_actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_attached_to_descriptor": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"associated_vnf_metadata_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_vnf_threshold_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": {
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

func dataSourceVNFRead(d *schema.ResourceData, m interface{}) error {
	filteredVNFs := vspk.VNFsList{}
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
	filteredVNFs, err = parent.VNFs(fetchFilter)
	if err != nil {
		return err
	}

	VNF := &vspk.VNF{}

	if len(filteredVNFs) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVNFs) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VNF = filteredVNFs[0]

	d.Set("vnf_descriptor_id", VNF.VNFDescriptorID)
	d.Set("vnf_descriptor_name", VNF.VNFDescriptorName)
	d.Set("cpu_count", VNF.CPUCount)
	d.Set("nsg_name", VNF.NSGName)
	d.Set("nsg_system_id", VNF.NSGSystemID)
	d.Set("ns_gateway_id", VNF.NSGatewayID)
	d.Set("name", VNF.Name)
	d.Set("task_state", VNF.TaskState)
	d.Set("last_known_error", VNF.LastKnownError)
	d.Set("memory_mb", VNF.MemoryMB)
	d.Set("vendor", VNF.Vendor)
	d.Set("description", VNF.Description)
	d.Set("metadata_id", VNF.MetadataID)
	d.Set("allowed_actions", VNF.AllowedActions)
	d.Set("enterprise_id", VNF.EnterpriseID)
	d.Set("is_attached_to_descriptor", VNF.IsAttachedToDescriptor)
	d.Set("associated_vnf_metadata_id", VNF.AssociatedVNFMetadataID)
	d.Set("associated_vnf_threshold_policy_id", VNF.AssociatedVNFThresholdPolicyID)
	d.Set("status", VNF.Status)
	d.Set("storage_gb", VNF.StorageGB)
	d.Set("type", VNF.Type)

	d.Set("id", VNF.Identifier())
	d.Set("parent_id", VNF.ParentID)
	d.Set("parent_type", VNF.ParentType)
	d.Set("owner", VNF.Owner)

	d.SetId(VNF.Identifier())

	return nil
}
