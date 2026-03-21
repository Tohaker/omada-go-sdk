# SiteServiceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirtimeFairness** | Pointer to [**AirtimeFairnessSettingOpenApiVO**](AirtimeFairnessSettingOpenApiVO.md) |  | [optional] 
**Alert** | Pointer to [**AlertSettingOpenApiVO**](AlertSettingOpenApiVO.md) |  | [optional] 
**BandSteering** | Pointer to [**BandSteeringOpenApiVO**](BandSteeringOpenApiVO.md) |  | [optional] 
**BandSteeringForMultiBand** | Pointer to [**BandSteeringMultiBandOpenApiVO**](BandSteeringMultiBandOpenApiVO.md) |  | [optional] 
**BeaconControl** | Pointer to [**BeaconControlOpenApiVO**](BeaconControlOpenApiVO.md) |  | [optional] 
**Led** | Pointer to [**SiteLedSetting**](SiteLedSetting.md) |  | [optional] 
**Lldp** | Pointer to [**SiteApLldpSettingOpenApiVO**](SiteApLldpSettingOpenApiVO.md) |  | [optional] 
**Mesh** | Pointer to [**NewMeshSettingOpenApiVO**](NewMeshSettingOpenApiVO.md) |  | [optional] 
**RemoteLog** | Pointer to [**RemoteLogSettingOpenApiVO**](RemoteLogSettingOpenApiVO.md) |  | [optional] 
**Roaming** | Pointer to [**NewRoamingSettingOpenApiVO**](NewRoamingSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewSiteServiceOpenApiVO

`func NewSiteServiceOpenApiVO() *SiteServiceOpenApiVO`

NewSiteServiceOpenApiVO instantiates a new SiteServiceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteServiceOpenApiVOWithDefaults

`func NewSiteServiceOpenApiVOWithDefaults() *SiteServiceOpenApiVO`

NewSiteServiceOpenApiVOWithDefaults instantiates a new SiteServiceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirtimeFairness

`func (o *SiteServiceOpenApiVO) GetAirtimeFairness() AirtimeFairnessSettingOpenApiVO`

GetAirtimeFairness returns the AirtimeFairness field if non-nil, zero value otherwise.

### GetAirtimeFairnessOk

`func (o *SiteServiceOpenApiVO) GetAirtimeFairnessOk() (*AirtimeFairnessSettingOpenApiVO, bool)`

GetAirtimeFairnessOk returns a tuple with the AirtimeFairness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirtimeFairness

`func (o *SiteServiceOpenApiVO) SetAirtimeFairness(v AirtimeFairnessSettingOpenApiVO)`

SetAirtimeFairness sets AirtimeFairness field to given value.

### HasAirtimeFairness

`func (o *SiteServiceOpenApiVO) HasAirtimeFairness() bool`

HasAirtimeFairness returns a boolean if a field has been set.

### GetAlert

`func (o *SiteServiceOpenApiVO) GetAlert() AlertSettingOpenApiVO`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *SiteServiceOpenApiVO) GetAlertOk() (*AlertSettingOpenApiVO, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *SiteServiceOpenApiVO) SetAlert(v AlertSettingOpenApiVO)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *SiteServiceOpenApiVO) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetBandSteering

`func (o *SiteServiceOpenApiVO) GetBandSteering() BandSteeringOpenApiVO`

GetBandSteering returns the BandSteering field if non-nil, zero value otherwise.

### GetBandSteeringOk

`func (o *SiteServiceOpenApiVO) GetBandSteeringOk() (*BandSteeringOpenApiVO, bool)`

GetBandSteeringOk returns a tuple with the BandSteering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteering

`func (o *SiteServiceOpenApiVO) SetBandSteering(v BandSteeringOpenApiVO)`

SetBandSteering sets BandSteering field to given value.

### HasBandSteering

`func (o *SiteServiceOpenApiVO) HasBandSteering() bool`

HasBandSteering returns a boolean if a field has been set.

### GetBandSteeringForMultiBand

`func (o *SiteServiceOpenApiVO) GetBandSteeringForMultiBand() BandSteeringMultiBandOpenApiVO`

GetBandSteeringForMultiBand returns the BandSteeringForMultiBand field if non-nil, zero value otherwise.

### GetBandSteeringForMultiBandOk

`func (o *SiteServiceOpenApiVO) GetBandSteeringForMultiBandOk() (*BandSteeringMultiBandOpenApiVO, bool)`

GetBandSteeringForMultiBandOk returns a tuple with the BandSteeringForMultiBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringForMultiBand

`func (o *SiteServiceOpenApiVO) SetBandSteeringForMultiBand(v BandSteeringMultiBandOpenApiVO)`

SetBandSteeringForMultiBand sets BandSteeringForMultiBand field to given value.

### HasBandSteeringForMultiBand

`func (o *SiteServiceOpenApiVO) HasBandSteeringForMultiBand() bool`

HasBandSteeringForMultiBand returns a boolean if a field has been set.

### GetBeaconControl

`func (o *SiteServiceOpenApiVO) GetBeaconControl() BeaconControlOpenApiVO`

GetBeaconControl returns the BeaconControl field if non-nil, zero value otherwise.

### GetBeaconControlOk

`func (o *SiteServiceOpenApiVO) GetBeaconControlOk() (*BeaconControlOpenApiVO, bool)`

GetBeaconControlOk returns a tuple with the BeaconControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconControl

`func (o *SiteServiceOpenApiVO) SetBeaconControl(v BeaconControlOpenApiVO)`

SetBeaconControl sets BeaconControl field to given value.

### HasBeaconControl

`func (o *SiteServiceOpenApiVO) HasBeaconControl() bool`

HasBeaconControl returns a boolean if a field has been set.

### GetLed

`func (o *SiteServiceOpenApiVO) GetLed() SiteLedSetting`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *SiteServiceOpenApiVO) GetLedOk() (*SiteLedSetting, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *SiteServiceOpenApiVO) SetLed(v SiteLedSetting)`

SetLed sets Led field to given value.

### HasLed

`func (o *SiteServiceOpenApiVO) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLldp

`func (o *SiteServiceOpenApiVO) GetLldp() SiteApLldpSettingOpenApiVO`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *SiteServiceOpenApiVO) GetLldpOk() (*SiteApLldpSettingOpenApiVO, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *SiteServiceOpenApiVO) SetLldp(v SiteApLldpSettingOpenApiVO)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *SiteServiceOpenApiVO) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetMesh

`func (o *SiteServiceOpenApiVO) GetMesh() NewMeshSettingOpenApiVO`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *SiteServiceOpenApiVO) GetMeshOk() (*NewMeshSettingOpenApiVO, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *SiteServiceOpenApiVO) SetMesh(v NewMeshSettingOpenApiVO)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *SiteServiceOpenApiVO) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetRemoteLog

`func (o *SiteServiceOpenApiVO) GetRemoteLog() RemoteLogSettingOpenApiVO`

GetRemoteLog returns the RemoteLog field if non-nil, zero value otherwise.

### GetRemoteLogOk

`func (o *SiteServiceOpenApiVO) GetRemoteLogOk() (*RemoteLogSettingOpenApiVO, bool)`

GetRemoteLogOk returns a tuple with the RemoteLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLog

`func (o *SiteServiceOpenApiVO) SetRemoteLog(v RemoteLogSettingOpenApiVO)`

SetRemoteLog sets RemoteLog field to given value.

### HasRemoteLog

`func (o *SiteServiceOpenApiVO) HasRemoteLog() bool`

HasRemoteLog returns a boolean if a field has been set.

### GetRoaming

`func (o *SiteServiceOpenApiVO) GetRoaming() NewRoamingSettingOpenApiVO`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *SiteServiceOpenApiVO) GetRoamingOk() (*NewRoamingSettingOpenApiVO, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *SiteServiceOpenApiVO) SetRoaming(v NewRoamingSettingOpenApiVO)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *SiteServiceOpenApiVO) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


