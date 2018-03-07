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
			"parent_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"passive_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"admin_state": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"dead_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"hello_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"message_digest_keys": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"metric": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"interface_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_subnet_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"authentication_key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ospf_area": &schema.Schema{
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
	} else {
		OSPFInterface = filteredOSPFInterfaces[0]
	}

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
