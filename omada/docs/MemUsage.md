# MemUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device MAC | [optional] 
**MemUsage** | Pointer to **int32** | Device memory usage such as 60 : 60% | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 

## Methods

### NewMemUsage

`func NewMemUsage() *MemUsage`

NewMemUsage instantiates a new MemUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemUsageWithDefaults

`func NewMemUsageWithDefaults() *MemUsage`

NewMemUsageWithDefaults instantiates a new MemUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *MemUsage) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MemUsage) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MemUsage) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MemUsage) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemUsage

`func (o *MemUsage) GetMemUsage() int32`

GetMemUsage returns the MemUsage field if non-nil, zero value otherwise.

### GetMemUsageOk

`func (o *MemUsage) GetMemUsageOk() (*int32, bool)`

GetMemUsageOk returns a tuple with the MemUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUsage

`func (o *MemUsage) SetMemUsage(v int32)`

SetMemUsage sets MemUsage field to given value.

### HasMemUsage

`func (o *MemUsage) HasMemUsage() bool`

HasMemUsage returns a boolean if a field has been set.

### GetModel

`func (o *MemUsage) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MemUsage) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MemUsage) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MemUsage) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *MemUsage) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *MemUsage) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *MemUsage) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *MemUsage) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *MemUsage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemUsage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemUsage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemUsage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *MemUsage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemUsage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemUsage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MemUsage) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


