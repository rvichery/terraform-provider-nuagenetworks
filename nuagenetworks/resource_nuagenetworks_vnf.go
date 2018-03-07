package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
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
			"vnf_descriptor_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vnf_descriptor_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"nsg_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nsg_system_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ns_gateway_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"task_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_known_error": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"memory_mb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vendor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"metadata_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"allowed_actions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"enterprise_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_attached_to_descriptor": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"associated_vnf_metadata_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_vnf_threshold_policy_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"storage_gb": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_enterprise": &schema.Schema{
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
		o.CPUCount = attr.(int)
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
		o.MemoryMB = attr.(int)
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
		o.IsAttachedToDescriptor = attr.(bool)
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
		o.StorageGB = attr.(int)
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
		o.CPUCount = attr.(int)
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
		o.MemoryMB = attr.(int)
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
		o.IsAttachedToDescriptor = attr.(bool)
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
		o.StorageGB = attr.(int)
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
