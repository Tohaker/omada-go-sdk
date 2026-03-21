# CliConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliConfig** | **string** | CLI configuration content | 
**CliType** | Pointer to **int32** | CLI configuration type, it should be a value as follows: 0：site CLI； 1：device CLI. | [optional] 
**Description** | Pointer to **string** | CLI configuration description, it should be within the range of 0 - 256 characters. | [optional] 
**DeviceType** | Pointer to **string** | Device type for CLI configuration application, it should be a value as follows: switch. | [optional] 
**Devices** | Pointer to [**[]DeviceCliVO**](DeviceCliVO.md) | List of devices bound to the CLI configuration. Only device CLI has this field. | [optional] 
**Id** | Pointer to **string** | CLI configuration ID | [optional] 
**Model** | Pointer to **string** | Device model of model cli | [optional] 
**ModelVersion** | Pointer to **string** | Device model version of model cli | [optional] 
**Name** | Pointer to **string** | CLI configuration name | [optional] 
**OmadacId** | Pointer to **string** | Omada ID | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Stacks** | Pointer to [**[]StackCliVO**](StackCliVO.md) | List of stacks bound to the CLI configuration, only device CLI has this field. | [optional] 
**Status** | Pointer to **int32** | CLI configuration status, it should be a value as follows: 0: active, 1: inactive | [optional] 

## Methods

### NewCliConfigOpenApiVO

`func NewCliConfigOpenApiVO(cliConfig string, ) *CliConfigOpenApiVO`

NewCliConfigOpenApiVO instantiates a new CliConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCliConfigOpenApiVOWithDefaults

`func NewCliConfigOpenApiVOWithDefaults() *CliConfigOpenApiVO`

NewCliConfigOpenApiVOWithDefaults instantiates a new CliConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliConfig

`func (o *CliConfigOpenApiVO) GetCliConfig() string`

GetCliConfig returns the CliConfig field if non-nil, zero value otherwise.

### GetCliConfigOk

`func (o *CliConfigOpenApiVO) GetCliConfigOk() (*string, bool)`

GetCliConfigOk returns a tuple with the CliConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliConfig

`func (o *CliConfigOpenApiVO) SetCliConfig(v string)`

SetCliConfig sets CliConfig field to given value.


### GetCliType

`func (o *CliConfigOpenApiVO) GetCliType() int32`

GetCliType returns the CliType field if non-nil, zero value otherwise.

### GetCliTypeOk

`func (o *CliConfigOpenApiVO) GetCliTypeOk() (*int32, bool)`

GetCliTypeOk returns a tuple with the CliType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliType

`func (o *CliConfigOpenApiVO) SetCliType(v int32)`

SetCliType sets CliType field to given value.

### HasCliType

`func (o *CliConfigOpenApiVO) HasCliType() bool`

HasCliType returns a boolean if a field has been set.

### GetDescription

`func (o *CliConfigOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CliConfigOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CliConfigOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CliConfigOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceType

`func (o *CliConfigOpenApiVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CliConfigOpenApiVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CliConfigOpenApiVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *CliConfigOpenApiVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDevices

`func (o *CliConfigOpenApiVO) GetDevices() []DeviceCliVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *CliConfigOpenApiVO) GetDevicesOk() (*[]DeviceCliVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *CliConfigOpenApiVO) SetDevices(v []DeviceCliVO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *CliConfigOpenApiVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetId

`func (o *CliConfigOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CliConfigOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CliConfigOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CliConfigOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModel

`func (o *CliConfigOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CliConfigOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CliConfigOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CliConfigOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *CliConfigOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *CliConfigOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *CliConfigOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *CliConfigOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *CliConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CliConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CliConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CliConfigOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *CliConfigOpenApiVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *CliConfigOpenApiVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *CliConfigOpenApiVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *CliConfigOpenApiVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetSiteId

`func (o *CliConfigOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CliConfigOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CliConfigOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CliConfigOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStacks

`func (o *CliConfigOpenApiVO) GetStacks() []StackCliVO`

GetStacks returns the Stacks field if non-nil, zero value otherwise.

### GetStacksOk

`func (o *CliConfigOpenApiVO) GetStacksOk() (*[]StackCliVO, bool)`

GetStacksOk returns a tuple with the Stacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStacks

`func (o *CliConfigOpenApiVO) SetStacks(v []StackCliVO)`

SetStacks sets Stacks field to given value.

### HasStacks

`func (o *CliConfigOpenApiVO) HasStacks() bool`

HasStacks returns a boolean if a field has been set.

### GetStatus

`func (o *CliConfigOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CliConfigOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CliConfigOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CliConfigOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


