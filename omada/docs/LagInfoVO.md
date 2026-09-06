# LagInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | Pointer to **int32** |  | [optional] 
**MlagEnable** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewLagInfoVO

`func NewLagInfoVO() *LagInfoVO`

NewLagInfoVO instantiates a new LagInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLagInfoVOWithDefaults

`func NewLagInfoVOWithDefaults() *LagInfoVO`

NewLagInfoVOWithDefaults instantiates a new LagInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *LagInfoVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *LagInfoVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *LagInfoVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *LagInfoVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetMlagEnable

`func (o *LagInfoVO) GetMlagEnable() bool`

GetMlagEnable returns the MlagEnable field if non-nil, zero value otherwise.

### GetMlagEnableOk

`func (o *LagInfoVO) GetMlagEnableOk() (*bool, bool)`

GetMlagEnableOk returns a tuple with the MlagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagEnable

`func (o *LagInfoVO) SetMlagEnable(v bool)`

SetMlagEnable sets MlagEnable field to given value.

### HasMlagEnable

`func (o *LagInfoVO) HasMlagEnable() bool`

HasMlagEnable returns a boolean if a field has been set.

### GetName

`func (o *LagInfoVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LagInfoVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LagInfoVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LagInfoVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *LagInfoVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *LagInfoVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *LagInfoVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *LagInfoVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


