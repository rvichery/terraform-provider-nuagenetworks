package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceOSPFInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceOSPFInterfaceRead,
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
			"passive_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dead_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hello_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"message_digest_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"metric": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_subnet_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"authentication_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ospf_area": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceOSPFInterfaceRead(d *schema.ResourceData, m interface{}) error {
	filteredOSPFInterfaces := vspk.OSPFInterfacesList{}
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
	parent := &vspk.OSPFArea{ID: d.Get("parent_ospf_area").(string)}
	filteredOSPFInterfaces, err = parent.OSPFInterfaces(fetchFilter)
	if err != nil {
		return err
	}

	OSPFInterface := &vspk.OSPFInterface{}

	if len(filteredOSPFInterfaces) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredOSPFInterfaces) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	OSPFInterface = filteredOSPFInterfaces[0]

	d.Set("name", OSPFInterface.Name)
	d.Set("passive_enabled", OSPFInterface.PassiveEnabled)
	d.Set("last_updated_by", OSPFInterface.LastUpdatedBy)
	d.Set("admin_state", OSPFInterface.AdminState)
	d.Set("dead_interval", OSPFInterface.DeadInterval)
	d.Set("hello_interval", OSPFInterface.HelloInterval)
	d.Set("description", OSPFInterface.Description)
	d.Set("message_digest_keys", OSPFInterface.MessageDigestKeys)
	d.Set("metric", OSPFInterface.Metric)
	d.Set("interface_type", OSPFInterface.InterfaceType)
	d.Set("entity_scope", OSPFInterface.EntityScope)
	d.Set("priority", OSPFInterface.Priority)
	d.Set("associated_subnet_id", OSPFInterface.AssociatedSubnetID)
	d.Set("mtu", OSPFInterface.Mtu)
	d.Set("authentication_key", OSPFInterface.AuthenticationKey)
	d.Set("authentication_type", OSPFInterface.AuthenticationType)
	d.Set("external_id", OSPFInterface.ExternalID)

	d.Set("id", OSPFInterface.Identifier())
	d.Set("parent_id", OSPFInterface.ParentID)
	d.Set("parent_type", OSPFInterface.ParentType)
	d.Set("owner", OSPFInterface.Owner)

	d.SetId(OSPFInterface.Identifier())

	return nil
}
