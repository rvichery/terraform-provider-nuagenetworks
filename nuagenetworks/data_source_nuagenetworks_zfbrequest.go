package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceZFBRequest() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceZFBRequestRead,
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
            "mac_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zfb_approval_status": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zfb_bootstrap_enabled": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "zfb_info": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "zfb_request_retry_timer": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "sku": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "ip_address": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "cpu_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsg_version": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "uuid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "family": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "last_connected_time": &schema.Schema{
                Type:     schema.TypeFloat,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "serial_number": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "hostname": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_enterprise_name": &schema.Schema{
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
            "status_string": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_enterprise": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}


func dataSourceZFBRequestRead(d *schema.ResourceData, m interface{}) error {
    filteredZFBRequests := vspk.ZFBRequestsList{}
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
    if attr, ok := d.GetOk("parent_enterprise"); ok {
        parent := &vspk.Enterprise{ID: attr.(string)}
        filteredZFBRequests, err = parent.ZFBRequests(fetchFilter)
        if err != nil {
            return err
        }
    } else {
        parent := m.(*vspk.Me)
        filteredZFBRequests, err = parent.ZFBRequests(fetchFilter)
        if err != nil {
            return err
        }
    }

    ZFBRequest := &vspk.ZFBRequest{}

    if len(filteredZFBRequests) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredZFBRequests) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    ZFBRequest = filteredZFBRequests[0]

    d.Set("mac_address", ZFBRequest.MACAddress)
    d.Set("zfb_approval_status", ZFBRequest.ZFBApprovalStatus)
    d.Set("zfb_bootstrap_enabled", ZFBRequest.ZFBBootstrapEnabled)
    d.Set("zfb_info", ZFBRequest.ZFBInfo)
    d.Set("zfb_request_retry_timer", ZFBRequest.ZFBRequestRetryTimer)
    d.Set("sku", ZFBRequest.SKU)
    d.Set("ip_address", ZFBRequest.IPAddress)
    d.Set("cpu_type", ZFBRequest.CPUType)
    d.Set("nsg_version", ZFBRequest.NSGVersion)
    d.Set("uuid", ZFBRequest.UUID)
    d.Set("family", ZFBRequest.Family)
    d.Set("last_connected_time", ZFBRequest.LastConnectedTime)
    d.Set("last_updated_by", ZFBRequest.LastUpdatedBy)
    d.Set("serial_number", ZFBRequest.SerialNumber)
    d.Set("entity_scope", ZFBRequest.EntityScope)
    d.Set("hostname", ZFBRequest.Hostname)
    d.Set("associated_enterprise_id", ZFBRequest.AssociatedEnterpriseID)
    d.Set("associated_enterprise_name", ZFBRequest.AssociatedEnterpriseName)
    d.Set("associated_ns_gateway_id", ZFBRequest.AssociatedNSGatewayID)
    d.Set("associated_ns_gateway_name", ZFBRequest.AssociatedNSGatewayName)
    d.Set("status_string", ZFBRequest.StatusString)
    d.Set("external_id", ZFBRequest.ExternalID)
    
    d.Set("id", ZFBRequest.Identifier())
    d.Set("parent_id", ZFBRequest.ParentID)
    d.Set("parent_type", ZFBRequest.ParentType)
    d.Set("owner", ZFBRequest.Owner)

    d.SetId(ZFBRequest.Identifier())
    
    return nil
}