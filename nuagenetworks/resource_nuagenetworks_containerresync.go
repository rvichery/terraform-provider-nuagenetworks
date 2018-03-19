package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceContainerResync() *schema.Resource {
	return &schema.Resource{
		Create: resourceContainerResyncCreate,
		Read:   resourceContainerResyncRead,
		Update: resourceContainerResyncUpdate,
		Delete: resourceContainerResyncDelete,
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
			"last_request_timestamp": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_time_resync_initiated": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_container"},
			},
			"parent_container": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_subnet"},
			},
		},
	}
}

func resourceContainerResyncCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize ContainerResync object
	o := &vspk.ContainerResync{}
	if attr, ok := d.GetOk("last_request_timestamp"); ok {
		o.LastRequestTimestamp = attr.(int)
	}
	if attr, ok := d.GetOk("last_time_resync_initiated"); ok {
		o.LastTimeResyncInitiated = attr.(int)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		err := parent.CreateContainerResync(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_container"); ok {
		parent := &vspk.Container{ID: attr.(string)}
		err := parent.CreateContainerResync(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceContainerResyncRead(d, m)
}

func resourceContainerResyncRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ContainerResync{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_request_timestamp", o.LastRequestTimestamp)
	d.Set("last_time_resync_initiated", o.LastTimeResyncInitiated)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("entity_scope", o.EntityScope)
	d.Set("status", o.Status)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceContainerResyncUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ContainerResync{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("last_request_timestamp"); ok {
		o.LastRequestTimestamp = attr.(int)
	}
	if attr, ok := d.GetOk("last_time_resync_initiated"); ok {
		o.LastTimeResyncInitiated = attr.(int)
	}
	if attr, ok := d.GetOk("status"); ok {
		o.Status = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceContainerResyncDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.ContainerResync{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
