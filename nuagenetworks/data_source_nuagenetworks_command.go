package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceCommand() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceCommandRead,
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
			"last_updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"detailed_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"detailed_status_code": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entity_scope": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"command_information": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_param": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_param_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"full_command": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_ns_gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceCommandRead(d *schema.ResourceData, m interface{}) error {
	filteredCommands := vspk.CommandsList{}
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
	parent := &vspk.NSGateway{ID: d.Get("parent_ns_gateway").(string)}
	filteredCommands, err = parent.Commands(fetchFilter)
	if err != nil {
		return err
	}

	Command := &vspk.Command{}

	if len(filteredCommands) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredCommands) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	} else {
		Command = filteredCommands[0]
	}

	d.Set("last_updated_by", Command.LastUpdatedBy)
	d.Set("detailed_status", Command.DetailedStatus)
	d.Set("detailed_status_code", Command.DetailedStatusCode)
	d.Set("entity_scope", Command.EntityScope)
	d.Set("command", Command.Command)
	d.Set("command_information", Command.CommandInformation)
	d.Set("associated_param", Command.AssociatedParam)
	d.Set("associated_param_type", Command.AssociatedParamType)
	d.Set("status", Command.Status)
	d.Set("full_command", Command.FullCommand)
	d.Set("summary", Command.Summary)
	d.Set("override", Command.Override)
	d.Set("external_id", Command.ExternalID)

	d.Set("id", Command.Identifier())
	d.Set("parent_id", Command.ParentID)
	d.Set("parent_type", Command.ParentType)
	d.Set("owner", Command.Owner)

	d.SetId(Command.Identifier())

	return nil
}
