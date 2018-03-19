/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// SystemConfigIdentity represents the Identity of the object
var SystemConfigIdentity = bambou.Identity{
	Name:     "systemconfig",
	Category: "systemconfigs",
}

// SystemConfigsList represents a list of SystemConfigs
type SystemConfigsList []*SystemConfig

// SystemConfigsAncestor is the interface that an ancestor of a SystemConfig must implement.
// An Ancestor is defined as an entity that has SystemConfig as a descendant.
// An Ancestor can get a list of its child SystemConfigs, but not necessarily create one.
type SystemConfigsAncestor interface {
	SystemConfigs(*bambou.FetchingInfo) (SystemConfigsList, *bambou.Error)
}

// SystemConfigsParent is the interface that a parent of a SystemConfig must implement.
// A Parent is defined as an entity that has SystemConfig as a child.
// A Parent is an Ancestor which can create a SystemConfig.
type SystemConfigsParent interface {
	SystemConfigsAncestor
	CreateSystemConfig(*SystemConfig) *bambou.Error
}

// SystemConfig represents the model of a systemconfig
type SystemConfig struct {
	ID                                                string      `json:"ID,omitempty"`
	ParentID                                          string      `json:"parentID,omitempty"`
	ParentType                                        string      `json:"parentType,omitempty"`
	Owner                                             string      `json:"owner,omitempty"`
	AARFlowStatsInterval                              int         `json:"AARFlowStatsInterval"`
	AARProbeStatsInterval                             int         `json:"AARProbeStatsInterval"`
	ACLAllowOrigin                                    string      `json:"ACLAllowOrigin,omitempty"`
	ECMPCount                                         int         `json:"ECMPCount"`
	LDAPSyncInterval                                  int         `json:"LDAPSyncInterval"`
	LDAPTrustStoreCertifcate                          string      `json:"LDAPTrustStoreCertifcate,omitempty"`
	LDAPTrustStorePassword                            string      `json:"LDAPTrustStorePassword,omitempty"`
	ADGatewayPurgeTime                                int         `json:"ADGatewayPurgeTime"`
	RDLowerLimit                                      int         `json:"RDLowerLimit"`
	RDPublicNetworkLowerLimit                         int         `json:"RDPublicNetworkLowerLimit"`
	RDPublicNetworkUpperLimit                         int         `json:"RDPublicNetworkUpperLimit"`
	RDUpperLimit                                      int         `json:"RDUpperLimit"`
	ZFBBootstrapEnabled                               bool        `json:"ZFBBootstrapEnabled"`
	ZFBRequestRetryTimer                              int         `json:"ZFBRequestRetryTimer"`
	ZFBSchedulerStaleRequestTimeout                   int         `json:"ZFBSchedulerStaleRequestTimeout"`
	PGIDLowerLimit                                    interface{} `json:"PGIDLowerLimit,omitempty"`
	PGIDUpperLimit                                    interface{} `json:"PGIDUpperLimit,omitempty"`
	DHCPOptionSize                                    int         `json:"DHCPOptionSize"`
	VLANIDLowerLimit                                  int         `json:"VLANIDLowerLimit"`
	VLANIDUpperLimit                                  int         `json:"VLANIDUpperLimit"`
	VMCacheSize                                       int         `json:"VMCacheSize"`
	VMPurgeTime                                       int         `json:"VMPurgeTime"`
	VMResyncDeletionWaitTime                          int         `json:"VMResyncDeletionWaitTime"`
	VMResyncOutstandingInterval                       int         `json:"VMResyncOutstandingInterval"`
	VMUnreachableCleanupTime                          int         `json:"VMUnreachableCleanupTime"`
	VMUnreachableTime                                 int         `json:"VMUnreachableTime"`
	VNFTaskTimeout                                    int         `json:"VNFTaskTimeout"`
	VNIDLowerLimit                                    int         `json:"VNIDLowerLimit"`
	VNIDPublicNetworkLowerLimit                       int         `json:"VNIDPublicNetworkLowerLimit"`
	VNIDPublicNetworkUpperLimit                       int         `json:"VNIDPublicNetworkUpperLimit"`
	VNIDUpperLimit                                    int         `json:"VNIDUpperLimit"`
	APIKeyRenewalInterval                             int         `json:"APIKeyRenewalInterval"`
	APIKeyValidity                                    int         `json:"APIKeyValidity"`
	VPortInitStatefulTimer                            int         `json:"VPortInitStatefulTimer"`
	LRUCacheSizePerSubnet                             int         `json:"LRUCacheSizePerSubnet"`
	VSCOnSameVersionAsVSD                             bool        `json:"VSCOnSameVersionAsVSD"`
	VSDReadOnlyMode                                   bool        `json:"VSDReadOnlyMode"`
	VSDUpgradeIsComplete                              bool        `json:"VSDUpgradeIsComplete"`
	ASNumber                                          int         `json:"ASNumber"`
	VSSStatsInterval                                  int         `json:"VSSStatsInterval"`
	RTLowerLimit                                      int         `json:"RTLowerLimit"`
	RTPublicNetworkLowerLimit                         int         `json:"RTPublicNetworkLowerLimit"`
	RTPublicNetworkUpperLimit                         int         `json:"RTPublicNetworkUpperLimit"`
	RTUpperLimit                                      int         `json:"RTUpperLimit"`
	EVPNBGPCommunityTagASNumber                       int         `json:"EVPNBGPCommunityTagASNumber"`
	EVPNBGPCommunityTagLowerLimit                     int         `json:"EVPNBGPCommunityTagLowerLimit"`
	EVPNBGPCommunityTagUpperLimit                     int         `json:"EVPNBGPCommunityTagUpperLimit"`
	PageMaxSize                                       int         `json:"pageMaxSize"`
	PageSize                                          int         `json:"pageSize"`
	LastUpdatedBy                                     string      `json:"lastUpdatedBy,omitempty"`
	MaxFailedLogins                                   int         `json:"maxFailedLogins"`
	MaxResponse                                       int         `json:"maxResponse"`
	AccumulateLicensesEnabled                         bool        `json:"accumulateLicensesEnabled"`
	VcinLoadBalancerIP                                string      `json:"vcinLoadBalancerIP,omitempty"`
	PerDomainVlanIdEnabled                            bool        `json:"perDomainVlanIdEnabled"`
	PerformancePathSelectionVNID                      int         `json:"performancePathSelectionVNID"`
	ServiceIDUpperLimit                               int         `json:"serviceIDUpperLimit"`
	KeyServerMonitorEnabled                           bool        `json:"keyServerMonitorEnabled"`
	KeyServerVSDDataSynchronizationInterval           int         `json:"keyServerVSDDataSynchronizationInterval"`
	OffsetCustomerID                                  int         `json:"offsetCustomerID"`
	OffsetServiceID                                   int         `json:"offsetServiceID"`
	VirtualFirewallRulesEnabled                       bool        `json:"virtualFirewallRulesEnabled"`
	EjbcaNSGCertificateProfile                        string      `json:"ejbcaNSGCertificateProfile,omitempty"`
	EjbcaNSGEndEntityProfile                          string      `json:"ejbcaNSGEndEntityProfile,omitempty"`
	EjbcaOCSPResponderCN                              string      `json:"ejbcaOCSPResponderCN,omitempty"`
	EjbcaOCSPResponderURI                             string      `json:"ejbcaOCSPResponderURI,omitempty"`
	EjbcaVspRootCa                                    string      `json:"ejbcaVspRootCa,omitempty"`
	AlarmsMaxPerObject                                int         `json:"alarmsMaxPerObject"`
	ElasticClusterName                                string      `json:"elasticClusterName,omitempty"`
	AllowEnterpriseAvatarOnNSG                        bool        `json:"allowEnterpriseAvatarOnNSG"`
	GlobalMACAddress                                  string      `json:"globalMACAddress,omitempty"`
	FlowCollectionEnabled                             bool        `json:"flowCollectionEnabled"`
	InactiveTimeout                                   int         `json:"inactiveTimeout"`
	InfrastructureBGPASNumber                         int         `json:"infrastructureBGPASNumber"`
	EntityScope                                       string      `json:"entityScope,omitempty"`
	DomainTunnelType                                  string      `json:"domainTunnelType,omitempty"`
	PostProcessorThreadsCount                         int         `json:"postProcessorThreadsCount"`
	GroupKeyDefaultSEKGenerationInterval              int         `json:"groupKeyDefaultSEKGenerationInterval"`
	GroupKeyDefaultSEKLifetime                        int         `json:"groupKeyDefaultSEKLifetime"`
	GroupKeyDefaultSEKPayloadEncryptionAlgorithm      string      `json:"groupKeyDefaultSEKPayloadEncryptionAlgorithm,omitempty"`
	GroupKeyDefaultSEKPayloadSigningAlgorithm         string      `json:"groupKeyDefaultSEKPayloadSigningAlgorithm,omitempty"`
	GroupKeyDefaultSeedGenerationInterval             int         `json:"groupKeyDefaultSeedGenerationInterval"`
	GroupKeyDefaultSeedLifetime                       int         `json:"groupKeyDefaultSeedLifetime"`
	GroupKeyDefaultSeedPayloadAuthenticationAlgorithm string      `json:"groupKeyDefaultSeedPayloadAuthenticationAlgorithm,omitempty"`
	GroupKeyDefaultSeedPayloadEncryptionAlgorithm     string      `json:"groupKeyDefaultSeedPayloadEncryptionAlgorithm,omitempty"`
	GroupKeyDefaultSeedPayloadSigningAlgorithm        string      `json:"groupKeyDefaultSeedPayloadSigningAlgorithm,omitempty"`
	GroupKeyDefaultTrafficAuthenticationAlgorithm     string      `json:"groupKeyDefaultTrafficAuthenticationAlgorithm,omitempty"`
	GroupKeyDefaultTrafficEncryptionAlgorithm         string      `json:"groupKeyDefaultTrafficEncryptionAlgorithm,omitempty"`
	GroupKeyDefaultTrafficEncryptionKeyLifetime       int         `json:"groupKeyDefaultTrafficEncryptionKeyLifetime"`
	GroupKeyGenerationIntervalOnForcedReKey           int         `json:"groupKeyGenerationIntervalOnForcedReKey"`
	GroupKeyGenerationIntervalOnRevoke                int         `json:"groupKeyGenerationIntervalOnRevoke"`
	GroupKeyMinimumSEKGenerationInterval              int         `json:"groupKeyMinimumSEKGenerationInterval"`
	GroupKeyMinimumSEKLifetime                        int         `json:"groupKeyMinimumSEKLifetime"`
	GroupKeyMinimumSeedGenerationInterval             int         `json:"groupKeyMinimumSeedGenerationInterval"`
	GroupKeyMinimumSeedLifetime                       int         `json:"groupKeyMinimumSeedLifetime"`
	GroupKeyMinimumTrafficEncryptionKeyLifetime       int         `json:"groupKeyMinimumTrafficEncryptionKeyLifetime"`
	NsgBootstrapEndpoint                              string      `json:"nsgBootstrapEndpoint,omitempty"`
	NsgConfigEndpoint                                 string      `json:"nsgConfigEndpoint,omitempty"`
	NsgLocalUiUrl                                     string      `json:"nsgLocalUiUrl,omitempty"`
	EsiID                                             int         `json:"esiID"`
	CsprootAuthenticationMethod                       string      `json:"csprootAuthenticationMethod,omitempty"`
	StackTraceEnabled                                 bool        `json:"stackTraceEnabled"`
	StatefulACLNonTCPTimeout                          int         `json:"statefulACLNonTCPTimeout"`
	StatefulACLTCPTimeout                             int         `json:"statefulACLTCPTimeout"`
	StaticWANServicePurgeTime                         int         `json:"staticWANServicePurgeTime"`
	StatisticsEnabled                                 bool        `json:"statisticsEnabled"`
	StatsCollectorAddress                             string      `json:"statsCollectorAddress,omitempty"`
	StatsCollectorPort                                string      `json:"statsCollectorPort,omitempty"`
	StatsCollectorProtoBufPort                        string      `json:"statsCollectorProtoBufPort,omitempty"`
	StatsMaxDataPoints                                int         `json:"statsMaxDataPoints"`
	StatsMinDuration                                  int         `json:"statsMinDuration"`
	StatsNumberOfDataPoints                           int         `json:"statsNumberOfDataPoints"`
	StatsTSDBServerAddress                            string      `json:"statsTSDBServerAddress,omitempty"`
	StickyECMPIdleTimeout                             int         `json:"stickyECMPIdleTimeout"`
	AttachProbeToIPsecNPM                             bool        `json:"attachProbeToIPsecNPM"`
	AttachProbeToVXLANNPM                             bool        `json:"attachProbeToVXLANNPM"`
	SubnetResyncInterval                              int         `json:"subnetResyncInterval"`
	SubnetResyncOutstandingInterval                   int         `json:"subnetResyncOutstandingInterval"`
	CustomerIDUpperLimit                              int         `json:"customerIDUpperLimit"`
	CustomerKey                                       string      `json:"customerKey,omitempty"`
	AvatarBasePath                                    string      `json:"avatarBasePath,omitempty"`
	AvatarBaseURL                                     string      `json:"avatarBaseURL,omitempty"`
	EventLogCleanupInterval                           int         `json:"eventLogCleanupInterval"`
	EventLogEntryMaxAge                               int         `json:"eventLogEntryMaxAge"`
	EventProcessorInterval                            int         `json:"eventProcessorInterval"`
	EventProcessorMaxEventsCount                      int         `json:"eventProcessorMaxEventsCount"`
	EventProcessorTimeout                             int         `json:"eventProcessorTimeout"`
	TwoFactorCodeExpiry                               int         `json:"twoFactorCodeExpiry"`
	TwoFactorCodeLength                               int         `json:"twoFactorCodeLength"`
	TwoFactorCodeSeedLength                           int         `json:"twoFactorCodeSeedLength"`
	ExternalID                                        string      `json:"externalID,omitempty"`
	DynamicWANServiceDiffTime                         int         `json:"dynamicWANServiceDiffTime"`
	SyslogDestinationHost                             string      `json:"syslogDestinationHost,omitempty"`
	SyslogDestinationPort                             int         `json:"syslogDestinationPort"`
	SysmonCleanupTaskInterval                         int         `json:"sysmonCleanupTaskInterval"`
	SysmonNodePresenceTimeout                         int         `json:"sysmonNodePresenceTimeout"`
	SysmonProbeResponseTimeout                        int         `json:"sysmonProbeResponseTimeout"`
	SystemAvatarData                                  string      `json:"systemAvatarData,omitempty"`
	SystemAvatarType                                  string      `json:"systemAvatarType,omitempty"`
}

// NewSystemConfig returns a new *SystemConfig
func NewSystemConfig() *SystemConfig {

	return &SystemConfig{
		AARFlowStatsInterval:        30,
		AARProbeStatsInterval:       30,
		ZFBRequestRetryTimer:        30,
		PGIDLowerLimit:              65536,
		PGIDUpperLimit:              2147483647,
		VMCacheSize:                 5000,
		VMPurgeTime:                 60,
		VMResyncDeletionWaitTime:    2,
		VMResyncOutstandingInterval: 1000,
		VMUnreachableCleanupTime:    7200,
		VMUnreachableTime:           3600,
		VNFTaskTimeout:              3600,
		VPortInitStatefulTimer:      300,
		VSSStatsInterval:            30,
		PageMaxSize:                 500,
		PageSize:                    50,
		AccumulateLicensesEnabled:   false,
		PerDomainVlanIdEnabled:      false,
		VirtualFirewallRulesEnabled: false,
		ElasticClusterName:          "nuage_elasticsearch",
		AllowEnterpriseAvatarOnNSG:  true,
		InfrastructureBGPASNumber:   65500,
		CsprootAuthenticationMethod: "LOCAL",
		StatsMinDuration:            2592000,
		StickyECMPIdleTimeout:       0,
		AttachProbeToIPsecNPM:       false,
		AttachProbeToVXLANNPM:       false,
		SubnetResyncInterval:        10,
		DynamicWANServiceDiffTime:   1,
	}
}

// Identity returns the Identity of the object.
func (o *SystemConfig) Identity() bambou.Identity {

	return SystemConfigIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *SystemConfig) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *SystemConfig) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the SystemConfig from the server
func (o *SystemConfig) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the SystemConfig into the server
func (o *SystemConfig) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the SystemConfig from the server
func (o *SystemConfig) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// Metadatas retrieves the list of child Metadatas of the SystemConfig
func (o *SystemConfig) Metadatas(info *bambou.FetchingInfo) (MetadatasList, *bambou.Error) {

	var list MetadatasList
	err := bambou.CurrentSession().FetchChildren(o, MetadataIdentity, &list, info)
	return list, err
}

// CreateMetadata creates a new child Metadata under the SystemConfig
func (o *SystemConfig) CreateMetadata(child *Metadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}

// GlobalMetadatas retrieves the list of child GlobalMetadatas of the SystemConfig
func (o *SystemConfig) GlobalMetadatas(info *bambou.FetchingInfo) (GlobalMetadatasList, *bambou.Error) {

	var list GlobalMetadatasList
	err := bambou.CurrentSession().FetchChildren(o, GlobalMetadataIdentity, &list, info)
	return list, err
}

// CreateGlobalMetadata creates a new child GlobalMetadata under the SystemConfig
func (o *SystemConfig) CreateGlobalMetadata(child *GlobalMetadata) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
