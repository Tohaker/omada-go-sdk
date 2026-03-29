# RrmIncompatibleApOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device MAC. | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**ModelVersion** | Pointer to **string** | Model version. | [optional] 
**Name** | Pointer to **string** | Device name. | [optional] 
**ShowModel** | Pointer to **string** | The model name of device. | [optional] 
**Status** | Pointer to **int32** | Detailed status. | [optional] 
**StatusCategory** | Pointer to **int32** | Status category. 0 : Disconnected. 1: Connected. 2: Pending. 3: Heartbeat Missed. 4: Isolated. | [optional] 
**Type** | Pointer to **string** | Device type. | [optional] 
**Version** | Pointer to **string** | Simplified version of firmware,for example:2.5.0 | [optional] 

## Methods

### NewRrmIncompatibleApOpenApiVO

`func NewRrmIncompatibleApOpenApiVO() *RrmIncompatibleApOpenApiVO`

NewRrmIncompatibleApOpenApiVO instantiates a new RrmIncompatibleApOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRrmIncompatibleApOpenApiVOWithDefaults

`func NewRrmIncompatibleApOpenApiVOWithDefaults() *RrmIncompatibleApOpenApiVO`

NewRrmIncompatibleApOpenApiVOWithDefaults instantiates a new RrmIncompatibleApOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *RrmIncompatibleApOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *RrmIncompatibleApOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *RrmIncompatibleApOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *RrmIncompatibleApOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *RrmIncompatibleApOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RrmIncompatibleApOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RrmIncompatibleApOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *RrmIncompatibleApOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *RrmIncompatibleApOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *RrmIncompatibleApOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *RrmIncompatibleApOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *RrmIncompatibleApOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *RrmIncompatibleApOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RrmIncompatibleApOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RrmIncompatibleApOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RrmIncompatibleApOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShowModel

`func (o *RrmIncompatibleApOpenApiVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *RrmIncompatibleApOpenApiVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *RrmIncompatibleApOpenApiVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *RrmIncompatibleApOpenApiVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStatus

`func (o *RrmIncompatibleApOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RrmIncompatibleApOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RrmIncompatibleApOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RrmIncompatibleApOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *RrmIncompatibleApOpenApiVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *RrmIncompatibleApOpenApiVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *RrmIncompatibleApOpenApiVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *RrmIncompatibleApOpenApiVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *RrmIncompatibleApOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RrmIncompatibleApOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RrmIncompatibleApOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RrmIncompatibleApOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *RrmIncompatibleApOpenApiVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *RrmIncompatibleApOpenApiVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *RrmIncompatibleApOpenApiVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *RrmIncompatibleApOpenApiVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


