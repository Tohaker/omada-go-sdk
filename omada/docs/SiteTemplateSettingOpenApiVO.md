# SiteTemplateSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteMulticastRateLimitSetting** | Pointer to [**NewMcastRateLimitSettingOpenApiVO**](NewMcastRateLimitSettingOpenApiVO.md) |  | [optional] 
**AdvancedFeature** | Pointer to [**AdvancedFeatureOpenApiVO**](AdvancedFeatureOpenApiVO.md) |  | [optional] 
**AirtimeFairness** | Pointer to [**AirtimeFairnessSettingOpenApiVO**](AirtimeFairnessSettingOpenApiVO.md) |  | [optional] 
**BandSteering** | Pointer to [**BandSteeringOpenApiVO**](BandSteeringOpenApiVO.md) |  | [optional] 
**BandSteeringForMultiBand** | Pointer to [**BandSteeringMultiBandOpenApiVO**](BandSteeringMultiBandOpenApiVO.md) |  | [optional] 
**BeaconControl** | Pointer to [**BeaconControlOpenApiVO**](BeaconControlOpenApiVO.md) |  | [optional] 
**ChannelLimit** | Pointer to [**NewChannelLimitSettingOpenApiVO**](NewChannelLimitSettingOpenApiVO.md) |  | [optional] 
**Led** | Pointer to [**SiteLedSetting**](SiteLedSetting.md) |  | [optional] 
**Lldp** | Pointer to [**SiteApLldpSettingOpenApiVO**](SiteApLldpSettingOpenApiVO.md) |  | [optional] 
**Mesh** | Pointer to [**NewMeshSettingOpenApiVO**](NewMeshSettingOpenApiVO.md) |  | [optional] 
**RemoteLog** | Pointer to [**RemoteLogSettingOpenApiVO**](RemoteLogSettingOpenApiVO.md) |  | [optional] 
**Roaming** | Pointer to [**NewRoamingSettingOpenApiVO**](NewRoamingSettingOpenApiVO.md) |  | [optional] 
**SiteTemplate** | Pointer to [**SiteTemplateOpenApiVO**](SiteTemplateOpenApiVO.md) |  | [optional] 

## Methods

### NewSiteTemplateSettingOpenApiVO

`func NewSiteTemplateSettingOpenApiVO() *SiteTemplateSettingOpenApiVO`

NewSiteTemplateSettingOpenApiVO instantiates a new SiteTemplateSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTemplateSettingOpenApiVOWithDefaults

`func NewSiteTemplateSettingOpenApiVOWithDefaults() *SiteTemplateSettingOpenApiVO`

NewSiteTemplateSettingOpenApiVOWithDefaults instantiates a new SiteTemplateSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteMulticastRateLimitSetting

`func (o *SiteTemplateSettingOpenApiVO) GetSiteMulticastRateLimitSetting() NewMcastRateLimitSettingOpenApiVO`

GetSiteMulticastRateLimitSetting returns the SiteMulticastRateLimitSetting field if non-nil, zero value otherwise.

### GetSiteMulticastRateLimitSettingOk

`func (o *SiteTemplateSettingOpenApiVO) GetSiteMulticastRateLimitSettingOk() (*NewMcastRateLimitSettingOpenApiVO, bool)`

GetSiteMulticastRateLimitSettingOk returns a tuple with the SiteMulticastRateLimitSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteMulticastRateLimitSetting

`func (o *SiteTemplateSettingOpenApiVO) SetSiteMulticastRateLimitSetting(v NewMcastRateLimitSettingOpenApiVO)`

SetSiteMulticastRateLimitSetting sets SiteMulticastRateLimitSetting field to given value.

### HasSiteMulticastRateLimitSetting

`func (o *SiteTemplateSettingOpenApiVO) HasSiteMulticastRateLimitSetting() bool`

HasSiteMulticastRateLimitSetting returns a boolean if a field has been set.

### GetAdvancedFeature

`func (o *SiteTemplateSettingOpenApiVO) GetAdvancedFeature() AdvancedFeatureOpenApiVO`

GetAdvancedFeature returns the AdvancedFeature field if non-nil, zero value otherwise.

### GetAdvancedFeatureOk

`func (o *SiteTemplateSettingOpenApiVO) GetAdvancedFeatureOk() (*AdvancedFeatureOpenApiVO, bool)`

GetAdvancedFeatureOk returns a tuple with the AdvancedFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedFeature

`func (o *SiteTemplateSettingOpenApiVO) SetAdvancedFeature(v AdvancedFeatureOpenApiVO)`

SetAdvancedFeature sets AdvancedFeature field to given value.

### HasAdvancedFeature

`func (o *SiteTemplateSettingOpenApiVO) HasAdvancedFeature() bool`

HasAdvancedFeature returns a boolean if a field has been set.

### GetAirtimeFairness

`func (o *SiteTemplateSettingOpenApiVO) GetAirtimeFairness() AirtimeFairnessSettingOpenApiVO`

GetAirtimeFairness returns the AirtimeFairness field if non-nil, zero value otherwise.

### GetAirtimeFairnessOk

`func (o *SiteTemplateSettingOpenApiVO) GetAirtimeFairnessOk() (*AirtimeFairnessSettingOpenApiVO, bool)`

GetAirtimeFairnessOk returns a tuple with the AirtimeFairness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirtimeFairness

`func (o *SiteTemplateSettingOpenApiVO) SetAirtimeFairness(v AirtimeFairnessSettingOpenApiVO)`

SetAirtimeFairness sets AirtimeFairness field to given value.

### HasAirtimeFairness

`func (o *SiteTemplateSettingOpenApiVO) HasAirtimeFairness() bool`

HasAirtimeFairness returns a boolean if a field has been set.

### GetBandSteering

`func (o *SiteTemplateSettingOpenApiVO) GetBandSteering() BandSteeringOpenApiVO`

GetBandSteering returns the BandSteering field if non-nil, zero value otherwise.

### GetBandSteeringOk

`func (o *SiteTemplateSettingOpenApiVO) GetBandSteeringOk() (*BandSteeringOpenApiVO, bool)`

GetBandSteeringOk returns a tuple with the BandSteering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteering

`func (o *SiteTemplateSettingOpenApiVO) SetBandSteering(v BandSteeringOpenApiVO)`

SetBandSteering sets BandSteering field to given value.

### HasBandSteering

`func (o *SiteTemplateSettingOpenApiVO) HasBandSteering() bool`

HasBandSteering returns a boolean if a field has been set.

### GetBandSteeringForMultiBand

`func (o *SiteTemplateSettingOpenApiVO) GetBandSteeringForMultiBand() BandSteeringMultiBandOpenApiVO`

GetBandSteeringForMultiBand returns the BandSteeringForMultiBand field if non-nil, zero value otherwise.

### GetBandSteeringForMultiBandOk

`func (o *SiteTemplateSettingOpenApiVO) GetBandSteeringForMultiBandOk() (*BandSteeringMultiBandOpenApiVO, bool)`

GetBandSteeringForMultiBandOk returns a tuple with the BandSteeringForMultiBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringForMultiBand

`func (o *SiteTemplateSettingOpenApiVO) SetBandSteeringForMultiBand(v BandSteeringMultiBandOpenApiVO)`

SetBandSteeringForMultiBand sets BandSteeringForMultiBand field to given value.

### HasBandSteeringForMultiBand

`func (o *SiteTemplateSettingOpenApiVO) HasBandSteeringForMultiBand() bool`

HasBandSteeringForMultiBand returns a boolean if a field has been set.

### GetBeaconControl

`func (o *SiteTemplateSettingOpenApiVO) GetBeaconControl() BeaconControlOpenApiVO`

GetBeaconControl returns the BeaconControl field if non-nil, zero value otherwise.

### GetBeaconControlOk

`func (o *SiteTemplateSettingOpenApiVO) GetBeaconControlOk() (*BeaconControlOpenApiVO, bool)`

GetBeaconControlOk returns a tuple with the BeaconControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconControl

`func (o *SiteTemplateSettingOpenApiVO) SetBeaconControl(v BeaconControlOpenApiVO)`

SetBeaconControl sets BeaconControl field to given value.

### HasBeaconControl

`func (o *SiteTemplateSettingOpenApiVO) HasBeaconControl() bool`

HasBeaconControl returns a boolean if a field has been set.

### GetChannelLimit

`func (o *SiteTemplateSettingOpenApiVO) GetChannelLimit() NewChannelLimitSettingOpenApiVO`

GetChannelLimit returns the ChannelLimit field if non-nil, zero value otherwise.

### GetChannelLimitOk

`func (o *SiteTemplateSettingOpenApiVO) GetChannelLimitOk() (*NewChannelLimitSettingOpenApiVO, bool)`

GetChannelLimitOk returns a tuple with the ChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimit

`func (o *SiteTemplateSettingOpenApiVO) SetChannelLimit(v NewChannelLimitSettingOpenApiVO)`

SetChannelLimit sets ChannelLimit field to given value.

### HasChannelLimit

`func (o *SiteTemplateSettingOpenApiVO) HasChannelLimit() bool`

HasChannelLimit returns a boolean if a field has been set.

### GetLed

`func (o *SiteTemplateSettingOpenApiVO) GetLed() SiteLedSetting`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *SiteTemplateSettingOpenApiVO) GetLedOk() (*SiteLedSetting, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *SiteTemplateSettingOpenApiVO) SetLed(v SiteLedSetting)`

SetLed sets Led field to given value.

### HasLed

`func (o *SiteTemplateSettingOpenApiVO) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLldp

`func (o *SiteTemplateSettingOpenApiVO) GetLldp() SiteApLldpSettingOpenApiVO`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *SiteTemplateSettingOpenApiVO) GetLldpOk() (*SiteApLldpSettingOpenApiVO, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *SiteTemplateSettingOpenApiVO) SetLldp(v SiteApLldpSettingOpenApiVO)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *SiteTemplateSettingOpenApiVO) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetMesh

`func (o *SiteTemplateSettingOpenApiVO) GetMesh() NewMeshSettingOpenApiVO`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *SiteTemplateSettingOpenApiVO) GetMeshOk() (*NewMeshSettingOpenApiVO, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *SiteTemplateSettingOpenApiVO) SetMesh(v NewMeshSettingOpenApiVO)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *SiteTemplateSettingOpenApiVO) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetRemoteLog

`func (o *SiteTemplateSettingOpenApiVO) GetRemoteLog() RemoteLogSettingOpenApiVO`

GetRemoteLog returns the RemoteLog field if non-nil, zero value otherwise.

### GetRemoteLogOk

`func (o *SiteTemplateSettingOpenApiVO) GetRemoteLogOk() (*RemoteLogSettingOpenApiVO, bool)`

GetRemoteLogOk returns a tuple with the RemoteLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLog

`func (o *SiteTemplateSettingOpenApiVO) SetRemoteLog(v RemoteLogSettingOpenApiVO)`

SetRemoteLog sets RemoteLog field to given value.

### HasRemoteLog

`func (o *SiteTemplateSettingOpenApiVO) HasRemoteLog() bool`

HasRemoteLog returns a boolean if a field has been set.

### GetRoaming

`func (o *SiteTemplateSettingOpenApiVO) GetRoaming() NewRoamingSettingOpenApiVO`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *SiteTemplateSettingOpenApiVO) GetRoamingOk() (*NewRoamingSettingOpenApiVO, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *SiteTemplateSettingOpenApiVO) SetRoaming(v NewRoamingSettingOpenApiVO)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *SiteTemplateSettingOpenApiVO) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.

### GetSiteTemplate

`func (o *SiteTemplateSettingOpenApiVO) GetSiteTemplate() SiteTemplateOpenApiVO`

GetSiteTemplate returns the SiteTemplate field if non-nil, zero value otherwise.

### GetSiteTemplateOk

`func (o *SiteTemplateSettingOpenApiVO) GetSiteTemplateOk() (*SiteTemplateOpenApiVO, bool)`

GetSiteTemplateOk returns a tuple with the SiteTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTemplate

`func (o *SiteTemplateSettingOpenApiVO) SetSiteTemplate(v SiteTemplateOpenApiVO)`

SetSiteTemplate sets SiteTemplate field to given value.

### HasSiteTemplate

`func (o *SiteTemplateSettingOpenApiVO) HasSiteTemplate() bool`

HasSiteTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


