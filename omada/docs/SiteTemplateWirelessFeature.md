# SiteTemplateWirelessFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedFeature** | Pointer to [**AdvancedFeatureVO**](AdvancedFeatureVO.md) |  | [optional] 
**AirtimeFairness** | Pointer to [**AirtimeFairnessSettingVO**](AirtimeFairnessSettingVO.md) |  | [optional] 
**BandSteering** | Pointer to [**BandSteeringVO**](BandSteeringVO.md) |  | [optional] 
**BandSteeringForMultiBand** | Pointer to [**BandSteeringMultiBandVO**](BandSteeringMultiBandVO.md) |  | [optional] 
**BeaconControl** | Pointer to [**BeaconControlVO**](BeaconControlVO.md) |  | [optional] 
**Lldp** | Pointer to [**SiteApLldpSettingVO**](SiteApLldpSettingVO.md) |  | [optional] 
**McastRateLimit** | Pointer to [**McastRateLimitSettingVO**](McastRateLimitSettingVO.md) |  | [optional] 
**Mesh** | Pointer to [**MeshSettingVO**](MeshSettingVO.md) |  | [optional] 
**Roaming** | Pointer to [**RoamingSettingVO**](RoamingSettingVO.md) |  | [optional] 

## Methods

### NewSiteTemplateWirelessFeature

`func NewSiteTemplateWirelessFeature() *SiteTemplateWirelessFeature`

NewSiteTemplateWirelessFeature instantiates a new SiteTemplateWirelessFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTemplateWirelessFeatureWithDefaults

`func NewSiteTemplateWirelessFeatureWithDefaults() *SiteTemplateWirelessFeature`

NewSiteTemplateWirelessFeatureWithDefaults instantiates a new SiteTemplateWirelessFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedFeature

`func (o *SiteTemplateWirelessFeature) GetAdvancedFeature() AdvancedFeatureVO`

GetAdvancedFeature returns the AdvancedFeature field if non-nil, zero value otherwise.

### GetAdvancedFeatureOk

`func (o *SiteTemplateWirelessFeature) GetAdvancedFeatureOk() (*AdvancedFeatureVO, bool)`

GetAdvancedFeatureOk returns a tuple with the AdvancedFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedFeature

`func (o *SiteTemplateWirelessFeature) SetAdvancedFeature(v AdvancedFeatureVO)`

SetAdvancedFeature sets AdvancedFeature field to given value.

### HasAdvancedFeature

`func (o *SiteTemplateWirelessFeature) HasAdvancedFeature() bool`

HasAdvancedFeature returns a boolean if a field has been set.

### GetAirtimeFairness

`func (o *SiteTemplateWirelessFeature) GetAirtimeFairness() AirtimeFairnessSettingVO`

GetAirtimeFairness returns the AirtimeFairness field if non-nil, zero value otherwise.

### GetAirtimeFairnessOk

`func (o *SiteTemplateWirelessFeature) GetAirtimeFairnessOk() (*AirtimeFairnessSettingVO, bool)`

GetAirtimeFairnessOk returns a tuple with the AirtimeFairness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirtimeFairness

`func (o *SiteTemplateWirelessFeature) SetAirtimeFairness(v AirtimeFairnessSettingVO)`

SetAirtimeFairness sets AirtimeFairness field to given value.

### HasAirtimeFairness

`func (o *SiteTemplateWirelessFeature) HasAirtimeFairness() bool`

HasAirtimeFairness returns a boolean if a field has been set.

### GetBandSteering

`func (o *SiteTemplateWirelessFeature) GetBandSteering() BandSteeringVO`

GetBandSteering returns the BandSteering field if non-nil, zero value otherwise.

### GetBandSteeringOk

`func (o *SiteTemplateWirelessFeature) GetBandSteeringOk() (*BandSteeringVO, bool)`

GetBandSteeringOk returns a tuple with the BandSteering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteering

`func (o *SiteTemplateWirelessFeature) SetBandSteering(v BandSteeringVO)`

SetBandSteering sets BandSteering field to given value.

### HasBandSteering

`func (o *SiteTemplateWirelessFeature) HasBandSteering() bool`

HasBandSteering returns a boolean if a field has been set.

### GetBandSteeringForMultiBand

`func (o *SiteTemplateWirelessFeature) GetBandSteeringForMultiBand() BandSteeringMultiBandVO`

GetBandSteeringForMultiBand returns the BandSteeringForMultiBand field if non-nil, zero value otherwise.

### GetBandSteeringForMultiBandOk

`func (o *SiteTemplateWirelessFeature) GetBandSteeringForMultiBandOk() (*BandSteeringMultiBandVO, bool)`

GetBandSteeringForMultiBandOk returns a tuple with the BandSteeringForMultiBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringForMultiBand

`func (o *SiteTemplateWirelessFeature) SetBandSteeringForMultiBand(v BandSteeringMultiBandVO)`

SetBandSteeringForMultiBand sets BandSteeringForMultiBand field to given value.

### HasBandSteeringForMultiBand

`func (o *SiteTemplateWirelessFeature) HasBandSteeringForMultiBand() bool`

HasBandSteeringForMultiBand returns a boolean if a field has been set.

### GetBeaconControl

`func (o *SiteTemplateWirelessFeature) GetBeaconControl() BeaconControlVO`

GetBeaconControl returns the BeaconControl field if non-nil, zero value otherwise.

### GetBeaconControlOk

`func (o *SiteTemplateWirelessFeature) GetBeaconControlOk() (*BeaconControlVO, bool)`

GetBeaconControlOk returns a tuple with the BeaconControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconControl

`func (o *SiteTemplateWirelessFeature) SetBeaconControl(v BeaconControlVO)`

SetBeaconControl sets BeaconControl field to given value.

### HasBeaconControl

`func (o *SiteTemplateWirelessFeature) HasBeaconControl() bool`

HasBeaconControl returns a boolean if a field has been set.

### GetLldp

`func (o *SiteTemplateWirelessFeature) GetLldp() SiteApLldpSettingVO`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *SiteTemplateWirelessFeature) GetLldpOk() (*SiteApLldpSettingVO, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *SiteTemplateWirelessFeature) SetLldp(v SiteApLldpSettingVO)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *SiteTemplateWirelessFeature) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetMcastRateLimit

`func (o *SiteTemplateWirelessFeature) GetMcastRateLimit() McastRateLimitSettingVO`

GetMcastRateLimit returns the McastRateLimit field if non-nil, zero value otherwise.

### GetMcastRateLimitOk

`func (o *SiteTemplateWirelessFeature) GetMcastRateLimitOk() (*McastRateLimitSettingVO, bool)`

GetMcastRateLimitOk returns a tuple with the McastRateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcastRateLimit

`func (o *SiteTemplateWirelessFeature) SetMcastRateLimit(v McastRateLimitSettingVO)`

SetMcastRateLimit sets McastRateLimit field to given value.

### HasMcastRateLimit

`func (o *SiteTemplateWirelessFeature) HasMcastRateLimit() bool`

HasMcastRateLimit returns a boolean if a field has been set.

### GetMesh

`func (o *SiteTemplateWirelessFeature) GetMesh() MeshSettingVO`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *SiteTemplateWirelessFeature) GetMeshOk() (*MeshSettingVO, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *SiteTemplateWirelessFeature) SetMesh(v MeshSettingVO)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *SiteTemplateWirelessFeature) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetRoaming

`func (o *SiteTemplateWirelessFeature) GetRoaming() RoamingSettingVO`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *SiteTemplateWirelessFeature) GetRoamingOk() (*RoamingSettingVO, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *SiteTemplateWirelessFeature) SetRoaming(v RoamingSettingVO)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *SiteTemplateWirelessFeature) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


