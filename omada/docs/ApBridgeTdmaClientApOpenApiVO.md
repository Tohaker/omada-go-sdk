# ApBridgeTdmaClientApOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Bridge TDMA Client AP mac. | [optional] 
**Model** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**Priority** | Pointer to **int32** | Bridge TDMA Client AP priority. 0: high, 1:base, 2:low. | [optional] 

## Methods

### NewApBridgeTdmaClientApOpenApiVO

`func NewApBridgeTdmaClientApOpenApiVO() *ApBridgeTdmaClientApOpenApiVO`

NewApBridgeTdmaClientApOpenApiVO instantiates a new ApBridgeTdmaClientApOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApBridgeTdmaClientApOpenApiVOWithDefaults

`func NewApBridgeTdmaClientApOpenApiVOWithDefaults() *ApBridgeTdmaClientApOpenApiVO`

NewApBridgeTdmaClientApOpenApiVOWithDefaults instantiates a new ApBridgeTdmaClientApOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *ApBridgeTdmaClientApOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApBridgeTdmaClientApOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApBridgeTdmaClientApOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ApBridgeTdmaClientApOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ApBridgeTdmaClientApOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApBridgeTdmaClientApOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApBridgeTdmaClientApOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApBridgeTdmaClientApOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApBridgeTdmaClientApOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApBridgeTdmaClientApOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApBridgeTdmaClientApOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApBridgeTdmaClientApOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ApBridgeTdmaClientApOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApBridgeTdmaClientApOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApBridgeTdmaClientApOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApBridgeTdmaClientApOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriority

`func (o *ApBridgeTdmaClientApOpenApiVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ApBridgeTdmaClientApOpenApiVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ApBridgeTdmaClientApOpenApiVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ApBridgeTdmaClientApOpenApiVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


