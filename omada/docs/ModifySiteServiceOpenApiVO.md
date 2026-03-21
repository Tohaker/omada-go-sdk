# ModifySiteServiceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AirtimeFairness** | Pointer to [**AirtimeFairnessSettingOpenApiVO**](AirtimeFairnessSettingOpenApiVO.md) |  | [optional] 
**BandSteering** | Pointer to [**BandSteeringOpenApiVO**](BandSteeringOpenApiVO.md) |  | [optional] 
**BandSteeringForMultiBand** | Pointer to [**BandSteeringMultiBandOpenApiVO**](BandSteeringMultiBandOpenApiVO.md) |  | [optional] 
**BeaconControl** | Pointer to [**BeaconControlOpenApiVO**](BeaconControlOpenApiVO.md) |  | [optional] 
**Led** | Pointer to [**SiteLedSetting**](SiteLedSetting.md) |  | [optional] 
**Lldp** | Pointer to [**SiteApLldpSettingOpenApiVO**](SiteApLldpSettingOpenApiVO.md) |  | [optional] 
**Mesh** | Pointer to [**NewMeshSettingOpenApiVO**](NewMeshSettingOpenApiVO.md) |  | [optional] 
**RemoteLog** | Pointer to [**RemoteLogSettingOpenApiVO**](RemoteLogSettingOpenApiVO.md) |  | [optional] 
**Roaming** | Pointer to [**NewRoamingSettingOpenApiVO**](NewRoamingSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewModifySiteServiceOpenApiVO

`func NewModifySiteServiceOpenApiVO() *ModifySiteServiceOpenApiVO`

NewModifySiteServiceOpenApiVO instantiates a new ModifySiteServiceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModifySiteServiceOpenApiVOWithDefaults

`func NewModifySiteServiceOpenApiVOWithDefaults() *ModifySiteServiceOpenApiVO`

NewModifySiteServiceOpenApiVOWithDefaults instantiates a new ModifySiteServiceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAirtimeFairness

`func (o *ModifySiteServiceOpenApiVO) GetAirtimeFairness() AirtimeFairnessSettingOpenApiVO`

GetAirtimeFairness returns the AirtimeFairness field if non-nil, zero value otherwise.

### GetAirtimeFairnessOk

`func (o *ModifySiteServiceOpenApiVO) GetAirtimeFairnessOk() (*AirtimeFairnessSettingOpenApiVO, bool)`

GetAirtimeFairnessOk returns a tuple with the AirtimeFairness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAirtimeFairness

`func (o *ModifySiteServiceOpenApiVO) SetAirtimeFairness(v AirtimeFairnessSettingOpenApiVO)`

SetAirtimeFairness sets AirtimeFairness field to given value.

### HasAirtimeFairness

`func (o *ModifySiteServiceOpenApiVO) HasAirtimeFairness() bool`

HasAirtimeFairness returns a boolean if a field has been set.

### GetBandSteering

`func (o *ModifySiteServiceOpenApiVO) GetBandSteering() BandSteeringOpenApiVO`

GetBandSteering returns the BandSteering field if non-nil, zero value otherwise.

### GetBandSteeringOk

`func (o *ModifySiteServiceOpenApiVO) GetBandSteeringOk() (*BandSteeringOpenApiVO, bool)`

GetBandSteeringOk returns a tuple with the BandSteering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteering

`func (o *ModifySiteServiceOpenApiVO) SetBandSteering(v BandSteeringOpenApiVO)`

SetBandSteering sets BandSteering field to given value.

### HasBandSteering

`func (o *ModifySiteServiceOpenApiVO) HasBandSteering() bool`

HasBandSteering returns a boolean if a field has been set.

### GetBandSteeringForMultiBand

`func (o *ModifySiteServiceOpenApiVO) GetBandSteeringForMultiBand() BandSteeringMultiBandOpenApiVO`

GetBandSteeringForMultiBand returns the BandSteeringForMultiBand field if non-nil, zero value otherwise.

### GetBandSteeringForMultiBandOk

`func (o *ModifySiteServiceOpenApiVO) GetBandSteeringForMultiBandOk() (*BandSteeringMultiBandOpenApiVO, bool)`

GetBandSteeringForMultiBandOk returns a tuple with the BandSteeringForMultiBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandSteeringForMultiBand

`func (o *ModifySiteServiceOpenApiVO) SetBandSteeringForMultiBand(v BandSteeringMultiBandOpenApiVO)`

SetBandSteeringForMultiBand sets BandSteeringForMultiBand field to given value.

### HasBandSteeringForMultiBand

`func (o *ModifySiteServiceOpenApiVO) HasBandSteeringForMultiBand() bool`

HasBandSteeringForMultiBand returns a boolean if a field has been set.

### GetBeaconControl

`func (o *ModifySiteServiceOpenApiVO) GetBeaconControl() BeaconControlOpenApiVO`

GetBeaconControl returns the BeaconControl field if non-nil, zero value otherwise.

### GetBeaconControlOk

`func (o *ModifySiteServiceOpenApiVO) GetBeaconControlOk() (*BeaconControlOpenApiVO, bool)`

GetBeaconControlOk returns a tuple with the BeaconControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconControl

`func (o *ModifySiteServiceOpenApiVO) SetBeaconControl(v BeaconControlOpenApiVO)`

SetBeaconControl sets BeaconControl field to given value.

### HasBeaconControl

`func (o *ModifySiteServiceOpenApiVO) HasBeaconControl() bool`

HasBeaconControl returns a boolean if a field has been set.

### GetLed

`func (o *ModifySiteServiceOpenApiVO) GetLed() SiteLedSetting`

GetLed returns the Led field if non-nil, zero value otherwise.

### GetLedOk

`func (o *ModifySiteServiceOpenApiVO) GetLedOk() (*SiteLedSetting, bool)`

GetLedOk returns a tuple with the Led field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLed

`func (o *ModifySiteServiceOpenApiVO) SetLed(v SiteLedSetting)`

SetLed sets Led field to given value.

### HasLed

`func (o *ModifySiteServiceOpenApiVO) HasLed() bool`

HasLed returns a boolean if a field has been set.

### GetLldp

`func (o *ModifySiteServiceOpenApiVO) GetLldp() SiteApLldpSettingOpenApiVO`

GetLldp returns the Lldp field if non-nil, zero value otherwise.

### GetLldpOk

`func (o *ModifySiteServiceOpenApiVO) GetLldpOk() (*SiteApLldpSettingOpenApiVO, bool)`

GetLldpOk returns a tuple with the Lldp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLldp

`func (o *ModifySiteServiceOpenApiVO) SetLldp(v SiteApLldpSettingOpenApiVO)`

SetLldp sets Lldp field to given value.

### HasLldp

`func (o *ModifySiteServiceOpenApiVO) HasLldp() bool`

HasLldp returns a boolean if a field has been set.

### GetMesh

`func (o *ModifySiteServiceOpenApiVO) GetMesh() NewMeshSettingOpenApiVO`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *ModifySiteServiceOpenApiVO) GetMeshOk() (*NewMeshSettingOpenApiVO, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *ModifySiteServiceOpenApiVO) SetMesh(v NewMeshSettingOpenApiVO)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *ModifySiteServiceOpenApiVO) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetRemoteLog

`func (o *ModifySiteServiceOpenApiVO) GetRemoteLog() RemoteLogSettingOpenApiVO`

GetRemoteLog returns the RemoteLog field if non-nil, zero value otherwise.

### GetRemoteLogOk

`func (o *ModifySiteServiceOpenApiVO) GetRemoteLogOk() (*RemoteLogSettingOpenApiVO, bool)`

GetRemoteLogOk returns a tuple with the RemoteLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLog

`func (o *ModifySiteServiceOpenApiVO) SetRemoteLog(v RemoteLogSettingOpenApiVO)`

SetRemoteLog sets RemoteLog field to given value.

### HasRemoteLog

`func (o *ModifySiteServiceOpenApiVO) HasRemoteLog() bool`

HasRemoteLog returns a boolean if a field has been set.

### GetRoaming

`func (o *ModifySiteServiceOpenApiVO) GetRoaming() NewRoamingSettingOpenApiVO`

GetRoaming returns the Roaming field if non-nil, zero value otherwise.

### GetRoamingOk

`func (o *ModifySiteServiceOpenApiVO) GetRoamingOk() (*NewRoamingSettingOpenApiVO, bool)`

GetRoamingOk returns a tuple with the Roaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoaming

`func (o *ModifySiteServiceOpenApiVO) SetRoaming(v NewRoamingSettingOpenApiVO)`

SetRoaming sets Roaming field to given value.

### HasRoaming

`func (o *ModifySiteServiceOpenApiVO) HasRoaming() bool`

HasRoaming returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


