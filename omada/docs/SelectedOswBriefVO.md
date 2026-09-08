# SelectedOswBriefVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevCap** | Pointer to [**OswDevCapVO**](OswDevCapVO.md) |  | [optional] 
**Lags** | Pointer to [**[]OswLagVO**](OswLagVO.md) | Switch lag List. | [optional] 
**Mac** | Pointer to **string** | Switch mac address. | [optional] 
**Model** | Pointer to **string** | Switch model. | [optional] 
**ModelVersion** | Pointer to **string** | Switch model version. | [optional] 
**Name** | Pointer to **string** | Switch name. | [optional] 
**OswStackData** | Pointer to [**OswStackDataVO**](OswStackDataVO.md) |  | [optional] 
**Ports** | Pointer to [**[]OswPortVO**](OswPortVO.md) | Switch port List. | [optional] 
**ShowModel** | Pointer to **string** | Switch show model. | [optional] 
**StackDevice** | Pointer to **bool** | Indicate whether this device is a stacked device. | [optional] 
**Status** | Pointer to **int32** | Switch status. | [optional] 
**StatusCategory** | Pointer to **int32** | Switch status category. | [optional] 

## Methods

### NewSelectedOswBriefVO

`func NewSelectedOswBriefVO() *SelectedOswBriefVO`

NewSelectedOswBriefVO instantiates a new SelectedOswBriefVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectedOswBriefVOWithDefaults

`func NewSelectedOswBriefVOWithDefaults() *SelectedOswBriefVO`

NewSelectedOswBriefVOWithDefaults instantiates a new SelectedOswBriefVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevCap

`func (o *SelectedOswBriefVO) GetDevCap() OswDevCapVO`

GetDevCap returns the DevCap field if non-nil, zero value otherwise.

### GetDevCapOk

`func (o *SelectedOswBriefVO) GetDevCapOk() (*OswDevCapVO, bool)`

GetDevCapOk returns a tuple with the DevCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevCap

`func (o *SelectedOswBriefVO) SetDevCap(v OswDevCapVO)`

SetDevCap sets DevCap field to given value.

### HasDevCap

`func (o *SelectedOswBriefVO) HasDevCap() bool`

HasDevCap returns a boolean if a field has been set.

### GetLags

`func (o *SelectedOswBriefVO) GetLags() []OswLagVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *SelectedOswBriefVO) GetLagsOk() (*[]OswLagVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *SelectedOswBriefVO) SetLags(v []OswLagVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *SelectedOswBriefVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMac

`func (o *SelectedOswBriefVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SelectedOswBriefVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SelectedOswBriefVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SelectedOswBriefVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *SelectedOswBriefVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SelectedOswBriefVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SelectedOswBriefVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SelectedOswBriefVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *SelectedOswBriefVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SelectedOswBriefVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SelectedOswBriefVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SelectedOswBriefVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *SelectedOswBriefVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelectedOswBriefVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelectedOswBriefVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelectedOswBriefVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOswStackData

`func (o *SelectedOswBriefVO) GetOswStackData() OswStackDataVO`

GetOswStackData returns the OswStackData field if non-nil, zero value otherwise.

### GetOswStackDataOk

`func (o *SelectedOswBriefVO) GetOswStackDataOk() (*OswStackDataVO, bool)`

GetOswStackDataOk returns a tuple with the OswStackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswStackData

`func (o *SelectedOswBriefVO) SetOswStackData(v OswStackDataVO)`

SetOswStackData sets OswStackData field to given value.

### HasOswStackData

`func (o *SelectedOswBriefVO) HasOswStackData() bool`

HasOswStackData returns a boolean if a field has been set.

### GetPorts

`func (o *SelectedOswBriefVO) GetPorts() []OswPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *SelectedOswBriefVO) GetPortsOk() (*[]OswPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *SelectedOswBriefVO) SetPorts(v []OswPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *SelectedOswBriefVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetShowModel

`func (o *SelectedOswBriefVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *SelectedOswBriefVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *SelectedOswBriefVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *SelectedOswBriefVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackDevice

`func (o *SelectedOswBriefVO) GetStackDevice() bool`

GetStackDevice returns the StackDevice field if non-nil, zero value otherwise.

### GetStackDeviceOk

`func (o *SelectedOswBriefVO) GetStackDeviceOk() (*bool, bool)`

GetStackDeviceOk returns a tuple with the StackDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackDevice

`func (o *SelectedOswBriefVO) SetStackDevice(v bool)`

SetStackDevice sets StackDevice field to given value.

### HasStackDevice

`func (o *SelectedOswBriefVO) HasStackDevice() bool`

HasStackDevice returns a boolean if a field has been set.

### GetStatus

`func (o *SelectedOswBriefVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SelectedOswBriefVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SelectedOswBriefVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SelectedOswBriefVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *SelectedOswBriefVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *SelectedOswBriefVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *SelectedOswBriefVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *SelectedOswBriefVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


