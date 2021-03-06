package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceVNF() *schema.Resource {
	return &schema.Resource{
		Create: resourceVNFCreate,
		Read:   resourceVNFRead,
		Update: resourceVNFUpdate,
		Delete: resourceVNFDelete,
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
			"vnf_descriptor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vnf_descriptor_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"nsg_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsg_system_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ns_gateway_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"task_state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"last_known_error": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"memory_mb": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"metadata_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowed_actions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enterprise_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"is_attached_to_descriptor": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"associated_vnf_metadata_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_vnf_threshold_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"storage_gb": {
				Type:     schema.TypeInt,
				Optional: true,
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

func resourceVNFCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize VNF object
	o := &vspk.VNF{
		VNFDescriptorID: d.Get("vnf_descriptor_id").(string),
		NSGatewayID:     d.Get("ns_gateway_id").(string),
		Name:            d.Get("name").(string),
	}
	if attr, ok := d.GetOk("vnf_descriptor_name"); ok {
		o.VNFDescriptorName = attr.(string)
	}
	if attr, ok := d.GetOk("cpu_count"); ok {
		CPUCount := attr.(int)
		o.CPUCount = &CPUCount
	}
	if attr, ok := d.GetOk("nsg_name"); ok {
		o.NSGName = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_system_id"); ok {
		o.NSGSystemID = attr.(string)
	}
	if attr, ok := d.GetOk("task_state"); ok {
		o.TaskState = attr.(string)
	}
	if attr, ok := d.GetOk("last_known_error"); ok {
		o.LastKnownError = attr.(string)
	}
	if attr, ok := d.GetOk("memory_mb"); ok {
		MemoryMB := attr.(int)
		o.MemoryMB = &MemoryMB
	}
	if attr, ok := d.GetOk("vendor"); ok {
		o.Vendor = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_id"); ok {
		o.MetadataID = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("is_attached_to_descriptor"); ok {
		IsAttachedToDescriptor := attr.(bool)
		o.IsAttachedToDescriptor = &IsAttachedToDescriptor
	}
	if attr, ok := d.GetOk("associated_vnf_metadata_id"); ok {
		o.AssociatedVNFMetadataID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vnf_threshold_policy_id"); ok {
		o.AssociatedVNFThresholdPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("storage_gb"); ok {
		StorageGB := attr.(int)
		o.StorageGB = &StorageGB
	}
	parent := &vspk.Enterprise{ID: d.Get("parent_enterprise").(string)}
	err := parent.CreateVNF(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceVNFRead(d, m)
}

func resourceVNFRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNF{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("vnf_descriptor_id", o.VNFDescriptorID)
	d.Set("vnf_descriptor_name", o.VNFDescriptorName)
	d.Set("cpu_count", o.CPUCount)
	d.Set("nsg_name", o.NSGName)
	d.Set("nsg_system_id", o.NSGSystemID)
	d.Set("ns_gateway_id", o.NSGatewayID)
	d.Set("name", o.Name)
	d.Set("task_state", o.TaskState)
	d.Set("last_known_error", o.LastKnownError)
	d.Set("memory_mb", o.MemoryMB)
	d.Set("vendor", o.Vendor)
	d.Set("description", o.Description)
	d.Set("metadata_id", o.MetadataID)
	d.Set("allowed_actions", o.AllowedActions)
	d.Set("enterprise_id", o.EnterpriseID)
	d.Set("is_attached_to_descriptor", o.IsAttachedToDescriptor)
	d.Set("associated_vnf_metadata_id", o.AssociatedVNFMetadataID)
	d.Set("associated_vnf_threshold_policy_id", o.AssociatedVNFThresholdPolicyID)
	d.Set("status", o.Status)
	d.Set("storage_gb", o.StorageGB)
	d.Set("type", o.Type)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceVNFUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNF{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.VNFDescriptorID = d.Get("vnf_descriptor_id").(string)
	o.NSGatewayID = d.Get("ns_gateway_id").(string)
	o.Name = d.Get("name").(string)

	if attr, ok := d.GetOk("vnf_descriptor_name"); ok {
		o.VNFDescriptorName = attr.(string)
	}
	if attr, ok := d.GetOk("cpu_count"); ok {
		CPUCount := attr.(int)
		o.CPUCount = &CPUCount
	}
	if attr, ok := d.GetOk("nsg_name"); ok {
		o.NSGName = attr.(string)
	}
	if attr, ok := d.GetOk("nsg_system_id"); ok {
		o.NSGSystemID = attr.(string)
	}
	if attr, ok := d.GetOk("task_state"); ok {
		o.TaskState = attr.(string)
	}
	if attr, ok := d.GetOk("last_known_error"); ok {
		o.LastKnownError = attr.(string)
	}
	if attr, ok := d.GetOk("memory_mb"); ok {
		MemoryMB := attr.(int)
		o.MemoryMB = &MemoryMB
	}
	if attr, ok := d.GetOk("vendor"); ok {
		o.Vendor = attr.(string)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("metadata_id"); ok {
		o.MetadataID = attr.(string)
	}
	if attr, ok := d.GetOk("enterprise_id"); ok {
		o.EnterpriseID = attr.(string)
	}
	if attr, ok := d.GetOk("is_attached_to_descriptor"); ok {
		IsAttachedToDescriptor := attr.(bool)
		o.IsAttachedToDescriptor = &IsAttachedToDescriptor
	}
	if attr, ok := d.GetOk("associated_vnf_metadata_id"); ok {
		o.AssociatedVNFMetadataID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_vnf_threshold_policy_id"); ok {
		o.AssociatedVNFThresholdPolicyID = attr.(string)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("storage_gb"); ok {
		StorageGB := attr.(int)
		o.StorageGB = &StorageGB
	}

	o.Save()

	return nil
}

func resourceVNFDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.VNF{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
