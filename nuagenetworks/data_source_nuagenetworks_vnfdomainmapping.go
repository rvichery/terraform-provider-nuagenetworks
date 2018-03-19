package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceVNFDomainMapping() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceVNFDomainMappingRead,
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
            "segmentation_id": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "segmentation_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ns_gateway_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ns_gateway_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func dataSourceVNFDomainMappingRead(d *schema.ResourceData, m interface{}) error {
    filteredVNFDomainMappings := vspk.VNFDomainMappingsList{}
    err := &bambou.Error{}
    fetchFilter := &bambou.FetchingInfo{}
    
    filters, filtersOk := d.GetOk("filter")
    if filtersOk {
        fetchFilter = bambou.NewFetchingInfo()
        for _, v := range filters.(*schema.Set).List() {
            m := v.(map[string]interface{})
            if fetchFilter.Filter != "" {
                fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string),  m["operator"].(string),  m["value"].(string))
            } else {
                fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
            }
           
        }
    }
    parent := &vspk.Domain{ID: d.Get("parent_domain").(string)}
    filteredVNFDomainMappings, err = parent.VNFDomainMappings(fetchFilter)
    if err != nil {
        return err
    }

    VNFDomainMapping := &vspk.VNFDomainMapping{}

    if len(filteredVNFDomainMappings) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredVNFDomainMappings) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    VNFDomainMapping = filteredVNFDomainMappings[0]

    d.Set("segmentation_id", VNFDomainMapping.SegmentationID)
    d.Set("segmentation_type", VNFDomainMapping.SegmentationType)
    d.Set("associated_ns_gateway_id", VNFDomainMapping.AssociatedNSGatewayID)
    d.Set("associated_ns_gateway_name", VNFDomainMapping.AssociatedNSGatewayName)
    
    d.Set("id", VNFDomainMapping.Identifier())
    d.Set("parent_id", VNFDomainMapping.ParentID)
    d.Set("parent_type", VNFDomainMapping.ParentType)
    d.Set("owner", VNFDomainMapping.Owner)

    d.SetId(VNFDomainMapping.Identifier())
    
    return nil
}