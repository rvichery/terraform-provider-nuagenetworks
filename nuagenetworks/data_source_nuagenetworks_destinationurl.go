package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceDestinationurl() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceDestinationurlRead,
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
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_method": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"packet_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"percentage_weight": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"down_threshold_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"probe_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_tier": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceDestinationurlRead(d *schema.ResourceData, m interface{}) error {
	filteredDestinationurls := vspk.DestinationurlsList{}
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
	parent := &vspk.Tier{ID: d.Get("parent_tier").(string)}
	filteredDestinationurls, err = parent.Destinationurls(fetchFilter)
	if err != nil {
		return err
	}

	Destinationurl := &vspk.Destinationurl{}

	if len(filteredDestinationurls) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredDestinationurls) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	Destinationurl = filteredDestinationurls[0]

	d.Set("url", Destinationurl.URL)
	d.Set("http_method", Destinationurl.HTTPMethod)
	d.Set("packet_count", Destinationurl.PacketCount)
	d.Set("last_updated_by", Destinationurl.LastUpdatedBy)
	d.Set("percentage_weight", Destinationurl.PercentageWeight)
	d.Set("timeout", Destinationurl.Timeout)
	d.Set("entity_scope", Destinationurl.EntityScope)
	d.Set("down_threshold_count", Destinationurl.DownThresholdCount)
	d.Set("probe_interval", Destinationurl.ProbeInterval)
	d.Set("external_id", Destinationurl.ExternalID)

	d.Set("id", Destinationurl.Identifier())
	d.Set("parent_id", Destinationurl.ParentID)
	d.Set("parent_type", Destinationurl.ParentType)
	d.Set("owner", Destinationurl.Owner)

	d.SetId(Destinationurl.Identifier())

	return nil
}
