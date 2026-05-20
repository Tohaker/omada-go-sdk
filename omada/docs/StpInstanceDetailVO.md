# StpInstanceDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | instance ID | [optional] 
**Stp** | **int32** | stp, 1: STP / 2: RSTP / 3: MSTP / 4: RPVST / 0: OFF | 
**Vlan** | Pointer to **int32** | instance vlanId | [optional] 

## Methods

### NewStpInstanceDetailVO

`func NewStpInstanceDetailVO(stp int32, ) *StpInstanceDetailVO`

NewStpInstanceDetailVO instantiates a new StpInstanceDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStpInstanceDetailVOWithDefaults

`func NewStpInstanceDetailVOWithDefaults() *StpInstanceDetailVO`

NewStpInstanceDetailVOWithDefaults instantiates a new StpInstanceDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StpInstanceDetailVO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StpInstanceDetailVO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StpInstanceDetailVO) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *StpInstanceDetailVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStp

`func (o *StpInstanceDetailVO) GetStp() int32`

GetStp returns the Stp field if non-nil, zero value otherwise.

### GetStpOk

`func (o *StpInstanceDetailVO) GetStpOk() (*int32, bool)`

GetStpOk returns a tuple with the Stp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStp

`func (o *StpInstanceDetailVO) SetStp(v int32)`

SetStp sets Stp field to given value.


### GetVlan

`func (o *StpInstanceDetailVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *StpInstanceDetailVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *StpInstanceDetailVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *StpInstanceDetailVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


