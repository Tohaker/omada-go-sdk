# OswMlagMemberPortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | Pointer to **int32** |  | [optional] 
**MacLagNames** | Pointer to **map[string]string** |  | [optional] 
**MacPorts** | Pointer to **map[string][]int32** |  | [optional] 
**MlagEnable** | Pointer to **bool** |  | [optional] 

## Methods

### NewOswMlagMemberPortVO

`func NewOswMlagMemberPortVO() *OswMlagMemberPortVO`

NewOswMlagMemberPortVO instantiates a new OswMlagMemberPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswMlagMemberPortVOWithDefaults

`func NewOswMlagMemberPortVOWithDefaults() *OswMlagMemberPortVO`

NewOswMlagMemberPortVOWithDefaults instantiates a new OswMlagMemberPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *OswMlagMemberPortVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswMlagMemberPortVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswMlagMemberPortVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *OswMlagMemberPortVO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetMacLagNames

`func (o *OswMlagMemberPortVO) GetMacLagNames() map[string]string`

GetMacLagNames returns the MacLagNames field if non-nil, zero value otherwise.

### GetMacLagNamesOk

`func (o *OswMlagMemberPortVO) GetMacLagNamesOk() (*map[string]string, bool)`

GetMacLagNamesOk returns a tuple with the MacLagNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLagNames

`func (o *OswMlagMemberPortVO) SetMacLagNames(v map[string]string)`

SetMacLagNames sets MacLagNames field to given value.

### HasMacLagNames

`func (o *OswMlagMemberPortVO) HasMacLagNames() bool`

HasMacLagNames returns a boolean if a field has been set.

### GetMacPorts

`func (o *OswMlagMemberPortVO) GetMacPorts() map[string][]int32`

GetMacPorts returns the MacPorts field if non-nil, zero value otherwise.

### GetMacPortsOk

`func (o *OswMlagMemberPortVO) GetMacPortsOk() (*map[string][]int32, bool)`

GetMacPortsOk returns a tuple with the MacPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPorts

`func (o *OswMlagMemberPortVO) SetMacPorts(v map[string][]int32)`

SetMacPorts sets MacPorts field to given value.

### HasMacPorts

`func (o *OswMlagMemberPortVO) HasMacPorts() bool`

HasMacPorts returns a boolean if a field has been set.

### GetMlagEnable

`func (o *OswMlagMemberPortVO) GetMlagEnable() bool`

GetMlagEnable returns the MlagEnable field if non-nil, zero value otherwise.

### GetMlagEnableOk

`func (o *OswMlagMemberPortVO) GetMlagEnableOk() (*bool, bool)`

GetMlagEnableOk returns a tuple with the MlagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagEnable

`func (o *OswMlagMemberPortVO) SetMlagEnable(v bool)`

SetMlagEnable sets MlagEnable field to given value.

### HasMlagEnable

`func (o *OswMlagMemberPortVO) HasMlagEnable() bool`

HasMlagEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


