package nuagenetworks

import (
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
    "github.com/nuagenetworks/go-bambou/bambou"
)

func dataSourceIKEGatewayConnection() *schema.Resource {
    return &schema.Resource{
        Read: dataSourceIKEGatewayConnectionRead,
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
            "nsg_identifier": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsg_identifier_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "nsg_role": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "mark": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "sequence": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "allow_any_subnet": &schema.Schema{
                Type:     schema.TypeBool,
                Computed: true,
            },
            "unencrypted_psk": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "port_vlan_name": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "priority": &schema.Schema{
                Type:     schema.TypeInt,
                Computed: true,
            },
            "associated_ike_authentication_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ike_authentication_type": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ike_encryption_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_ike_gateway_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "associated_vlanid": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Computed: true,
            },
            "parent_subnet": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_vlan"},
            },
            "parent_vlan": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_subnet"},
            },
        },
    }
}


func dataSourceIKEGatewayConnectionRead(d *schema.ResourceData, m interface{}) error {
    filteredIKEGatewayConnections := vspk.IKEGatewayConnectionsList{}
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
    if attr, ok := d.GetOk("parent_subnet"); ok {
        parent := &vspk.Subnet{ID: attr.(string)}
        filteredIKEGatewayConnections, err = parent.IKEGatewayConnections(fetchFilter)
        if err != nil {
            return err
        }
    } else if attr, ok := d.GetOk("parent_vlan"); ok {
        parent := &vspk.VLAN{ID: attr.(string)}
        filteredIKEGatewayConnections, err = parent.IKEGatewayConnections(fetchFilter)
        if err != nil {
            return err
        }
    }

    IKEGatewayConnection := &vspk.IKEGatewayConnection{}

    if len(filteredIKEGatewayConnections) < 1 {
        return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
    }

    if len(filteredIKEGatewayConnections) > 1 {
        return fmt.Errorf("Your query returned more than one result. Please try a more " +
            "specific search criteria.")
    }
    
    IKEGatewayConnection = filteredIKEGatewayConnections[0]

    d.Set("nsg_identifier", IKEGatewayConnection.NSGIdentifier)
    d.Set("nsg_identifier_type", IKEGatewayConnection.NSGIdentifierType)
    d.Set("nsg_role", IKEGatewayConnection.NSGRole)
    d.Set("name", IKEGatewayConnection.Name)
    d.Set("mark", IKEGatewayConnection.Mark)
    d.Set("last_updated_by", IKEGatewayConnection.LastUpdatedBy)
    d.Set("sequence", IKEGatewayConnection.Sequence)
    d.Set("allow_any_subnet", IKEGatewayConnection.AllowAnySubnet)
    d.Set("unencrypted_psk", IKEGatewayConnection.UnencryptedPSK)
    d.Set("entity_scope", IKEGatewayConnection.EntityScope)
    d.Set("port_vlan_name", IKEGatewayConnection.PortVLANName)
    d.Set("priority", IKEGatewayConnection.Priority)
    d.Set("associated_ike_authentication_id", IKEGatewayConnection.AssociatedIKEAuthenticationID)
    d.Set("associated_ike_authentication_type", IKEGatewayConnection.AssociatedIKEAuthenticationType)
    d.Set("associated_ike_encryption_profile_id", IKEGatewayConnection.AssociatedIKEEncryptionProfileID)
    d.Set("associated_ike_gateway_profile_id", IKEGatewayConnection.AssociatedIKEGatewayProfileID)
    d.Set("associated_vlanid", IKEGatewayConnection.AssociatedVLANID)
    d.Set("external_id", IKEGatewayConnection.ExternalID)
    
    d.Set("id", IKEGatewayConnection.Identifier())
    d.Set("parent_id", IKEGatewayConnection.ParentID)
    d.Set("parent_type", IKEGatewayConnection.ParentType)
    d.Set("owner", IKEGatewayConnection.Owner)

    d.SetId(IKEGatewayConnection.Identifier())
    
    return nil
}