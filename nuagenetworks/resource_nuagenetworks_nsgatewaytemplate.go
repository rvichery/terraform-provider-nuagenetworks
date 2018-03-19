package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourceNSGatewayTemplate() *schema.Resource {
    return &schema.Resource{
        Create: resourceNSGatewayTemplateCreate,
        Read:   resourceNSGatewayTemplateRead,
        Update: resourceNSGatewayTemplateUpdate,
        Delete: resourceNSGatewayTemplateDelete,
        Importer: &schema.ResourceImporter{
            State: schema.ImportStatePassthrough,
        },
        Schema: map[string]*schema.Schema{
            "id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "parent_type": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "owner": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "ssh_service": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "ENABLED",
            },
            "name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "last_updated_by": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "personality": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "infrastructure_access_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "infrastructure_profile_id": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "instance_ssh_override": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Default: "DISALLOWED",
            },
            "enterprise_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
        },
    }
}

func resourceNSGatewayTemplateCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize NSGatewayTemplate object
    o := &vspk.NSGatewayTemplate{
        Name: d.Get("name").(string),
        InfrastructureProfileID: d.Get("infrastructure_profile_id").(string),
    }
    if attr, ok := d.GetOk("ssh_service"); ok {
        o.SSHService = attr.(string)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("infrastructure_access_profile_id"); ok {
        o.InfrastructureAccessProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("instance_ssh_override"); ok {
        o.InstanceSSHOverride = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    parent := m.(*vspk.Me)
    err := parent.CreateNSGatewayTemplate(o)
    if err != nil {
        return err
    }
    
    

    d.SetId(o.Identifier())
    return resourceNSGatewayTemplateRead(d, m)
}

func resourceNSGatewayTemplateRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSGatewayTemplate{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("ssh_service", o.SSHService)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("personality", o.Personality)
    d.Set("description", o.Description)
    d.Set("infrastructure_access_profile_id", o.InfrastructureAccessProfileID)
    d.Set("infrastructure_profile_id", o.InfrastructureProfileID)
    d.Set("instance_ssh_override", o.InstanceSSHOverride)
    d.Set("enterprise_id", o.EnterpriseID)
    d.Set("entity_scope", o.EntityScope)
    d.Set("external_id", o.ExternalID)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourceNSGatewayTemplateUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSGatewayTemplate{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.InfrastructureProfileID = d.Get("infrastructure_profile_id").(string)
    
    if attr, ok := d.GetOk("ssh_service"); ok {
        o.SSHService = attr.(string)
    }
    if attr, ok := d.GetOk("personality"); ok {
        o.Personality = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("infrastructure_access_profile_id"); ok {
        o.InfrastructureAccessProfileID = attr.(string)
    }
    if attr, ok := d.GetOk("instance_ssh_override"); ok {
        o.InstanceSSHOverride = attr.(string)
    }
    if attr, ok := d.GetOk("enterprise_id"); ok {
        o.EnterpriseID = attr.(string)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourceNSGatewayTemplateDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.NSGatewayTemplate{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}