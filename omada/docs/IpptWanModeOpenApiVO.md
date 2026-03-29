# IpptWanModeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | whether to enable wan Settings Override | 
**Interval** | Pointer to **int32** | Time taken to check the connection of wan port. | [optional] 
**Unit** | Pointer to **int32** | Unit. 0: minute, 1: second. | [optional] 

## Methods

### NewIpptWanModeOpenApiVO

`func NewIpptWanModeOpenApiVO(enable bool, ) *IpptWanModeOpenApiVO`

NewIpptWanModeOpenApiVO instantiates a new IpptWanModeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpptWanModeOpenApiVOWithDefaults

`func NewIpptWanModeOpenApiVOWithDefaults() *IpptWanModeOpenApiVO`

NewIpptWanModeOpenApiVOWithDefaults instantiates a new IpptWanModeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IpptWanModeOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IpptWanModeOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IpptWanModeOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetInterval

`func (o *IpptWanModeOpenApiVO) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *IpptWanModeOpenApiVO) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *IpptWanModeOpenApiVO) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *IpptWanModeOpenApiVO) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetUnit

`func (o *IpptWanModeOpenApiVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *IpptWanModeOpenApiVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *IpptWanModeOpenApiVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *IpptWanModeOpenApiVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


