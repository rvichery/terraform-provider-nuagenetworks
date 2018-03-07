package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceApplicationBinding() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplicationBindingRead,
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
			"read_only": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"associated_application_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_applicationperformancemanagement": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_application"},
			},
			"parent_application": &schema.Schema{
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_applicationperformancemanagement"},
			},
		},
	}
}

func dataSourceApplicationBindingRead(d *schema.ResourceData, m interface{}) error {
	filteredApplicationBindings := vspk.ApplicationBindingsList{}
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
	if attr, ok := d.GetOk("parent_applicationperformancemanagement"); ok {
		parent := &vspk.Applicationperformancemanagement{ID: attr.(string)}
		filteredApplicationBindings, err = parent.ApplicationBindings(fetchFilter)
		if err != nil {
			return err
		}
	} else if attr, ok := d.GetOk("parent_application"); ok {
		parent := &vspk.Application{ID: attr.(string)}
		filteredApplicationBindings, err = parent.ApplicationBindings(fetchFilter)
		if err != nil {
			return err
		}
	}

	ApplicationBinding := &vspk.ApplicationBinding{}

	if len(filteredApplicationBindings) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredApplicationBindings) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		ApplicationBinding = filteredApplicationBindings[0]
	}

	d.Set("read_only", ApplicationBinding.ReadOnly)
	d.Set("priority", ApplicationBinding.Priority)
	d.Set("associated_application_id", ApplicationBinding.AssociatedApplicationID)

	d.Set("id", ApplicationBinding.Identifier())
	d.Set("parent_id", ApplicationBinding.ParentID)
	d.Set("parent_type", ApplicationBinding.ParentType)
	d.Set("owner", ApplicationBinding.Owner)

	d.SetId(ApplicationBinding.Identifier())

	return nil
}
