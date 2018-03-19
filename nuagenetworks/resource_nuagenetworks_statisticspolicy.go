package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func resourceStatisticsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceStatisticsPolicyCreate,
		Read:   resourceStatisticsPolicyRead,
		Update: resourceStatisticsPolicyUpdate,
		Delete: resourceStatisticsPolicyDelete,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"data_collection_frequency": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_zone": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain", "parent_vport", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_vport", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_vport": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_subnet": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_address_map", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_address_map": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_l2_domain", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_l2_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_address_map", "parent_ns_port", "parent_patnat_pool"},
			},
			"parent_ns_port": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_patnat_pool"},
			},
			"parent_patnat_pool": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_zone", "parent_domain", "parent_vport", "parent_subnet", "parent_address_map", "parent_l2_domain", "parent_ns_port"},
			},
		},
	}
}

func resourceStatisticsPolicyCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize StatisticsPolicy object
	o := &vspk.StatisticsPolicy{
		Name: d.Get("name").(string),
		DataCollectionFrequency: d.Get("data_collection_frequency").(int),
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_zone"); ok {
		parent := &vspk.Zone{ID: attr.(string)}
		err := parent.CreateStatisticsPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateStatisticsPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vport"); ok {
		parent := &vspk.VPort{ID: attr.(string)}
		err := parent.CreateStatisticsPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_subnet"); ok {
		parent := &vspk.Subnet{ID: attr.(string)}
		err := parent.CreateStatisticsPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_address_map"); ok {
		parent := &vspk.AddressMap{ID: attr.(string)}
		err := parent.CreateStatisticsPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_l2_domain"); ok {
		parent := &vspk.L2Domain{ID: attr.(string)}
		err := parent.CreateStatisticsPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_ns_port"); ok {
		parent := &vspk.NSPort{ID: attr.(string)}
		err := parent.CreateStatisticsPolicy(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_patnat_pool"); ok {
		parent := &vspk.PATNATPool{ID: attr.(string)}
		err := parent.CreateStatisticsPolicy(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceStatisticsPolicyRead(d, m)
}

func resourceStatisticsPolicyRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.StatisticsPolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("data_collection_frequency", o.DataCollectionFrequency)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceStatisticsPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.StatisticsPolicy{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	o.Name = d.Get("name").(string)
	o.DataCollectionFrequency = d.Get("data_collection_frequency").(int)

	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceStatisticsPolicyDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.StatisticsPolicy{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
