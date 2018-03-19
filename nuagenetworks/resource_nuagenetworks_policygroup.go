package nuagenetworks

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/rvichery/vspk-go/vspk"
)

func resourcePolicyGroup() *schema.Resource {
    return &schema.Resource{
        Create: resourcePolicyGroupCreate,
        Read:   resourcePolicyGroupRead,
        Update: resourcePolicyGroupUpdate,
        Delete: resourcePolicyGroupDelete,
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
            "evpn_community_tag": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
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
            "template_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "description": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "entity_scope": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
            },
            "policy_group_id": &schema.Schema{
                Type:     schema.TypeInt,
                Optional: true,
            },
            "external": &schema.Schema{
                Type:     schema.TypeBool,
                Optional: true,
            },
            "external_id": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
            },
            "type": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "parent_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_l2_domain"},
            },
            "parent_l2_domain": &schema.Schema{
                Type:     schema.TypeString,
                Optional: true,
                ConflictsWith: []string{"parent_domain"},
            },
        },
    }
}

func resourcePolicyGroupCreate(d *schema.ResourceData, m interface{}) error {

    // Initialize PolicyGroup object
    o := &vspk.PolicyGroup{
        Name: d.Get("name").(string),
        Type: d.Get("type").(string),
    }
    if attr, ok := d.GetOk("evpn_community_tag"); ok {
        o.EVPNCommunityTag = attr.(string)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("policy_group_id"); ok {
        o.PolicyGroupID = attr.(int)
    }
    if attr, ok := d.GetOk("external"); ok {
        o.External = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }
    if attr, ok := d.GetOk("parent_domain"); ok {
        parent := &vspk.Domain{ID: attr.(string)}
        err := parent.CreatePolicyGroup(o)
        if err != nil {
            return err
        }
    }
    if attr, ok := d.GetOk("parent_l2_domain"); ok {
        parent := &vspk.L2Domain{ID: attr.(string)}
        err := parent.CreatePolicyGroup(o)
        if err != nil {
            return err
        }
    }
    
    

    d.SetId(o.Identifier())
    if attr, ok := d.GetOk("vports"); ok {
        o.AssignVPorts(attr.(vspk.VPortsList))
    }
    return resourcePolicyGroupRead(d, m)
}

func resourcePolicyGroupRead(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PolicyGroup{
        ID: d.Id(),
    }

    err := o.Fetch()
    if err != nil {
        d.SetId("")
        return nil
    }

    d.Set("evpn_community_tag", o.EVPNCommunityTag)
    d.Set("name", o.Name)
    d.Set("last_updated_by", o.LastUpdatedBy)
    d.Set("template_id", o.TemplateID)
    d.Set("description", o.Description)
    d.Set("entity_scope", o.EntityScope)
    d.Set("policy_group_id", o.PolicyGroupID)
    d.Set("external", o.External)
    d.Set("external_id", o.ExternalID)
    d.Set("type", o.Type)
    
    d.Set("id", o.Identifier())
    d.Set("parent_id", o.ParentID)
    d.Set("parent_type", o.ParentType)
    d.Set("owner", o.Owner)

    return nil
}

func resourcePolicyGroupUpdate(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PolicyGroup{
        ID: d.Id(),
    }
    
    err := o.Fetch()
    if err != nil {
        return err
    }
    
    o.Name = d.Get("name").(string)
    o.Type = d.Get("type").(string)
    
    if attr, ok := d.GetOk("evpn_community_tag"); ok {
        o.EVPNCommunityTag = attr.(string)
    }
    if attr, ok := d.GetOk("template_id"); ok {
        o.TemplateID = attr.(string)
    }
    if attr, ok := d.GetOk("description"); ok {
        o.Description = attr.(string)
    }
    if attr, ok := d.GetOk("policy_group_id"); ok {
        o.PolicyGroupID = attr.(int)
    }
    if attr, ok := d.GetOk("external"); ok {
        o.External = attr.(bool)
    }
    if attr, ok := d.GetOk("external_id"); ok {
        o.ExternalID = attr.(string)
    }

    o.Save()

    return nil
}

func resourcePolicyGroupDelete(d *schema.ResourceData, m interface{}) error {
    o := &vspk.PolicyGroup{
        ID: d.Id(),
    }

    err := o.Delete()
    if err != nil {
        return err
    }

    return nil
}