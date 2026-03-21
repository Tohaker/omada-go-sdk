# DhcpGuardVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpSvr1** | Pointer to **string** |  | [optional] 
**DhcpSvr2** | Pointer to **string** |  | [optional] 
**Enable** | **bool** |  | 

## Methods

### NewDhcpGuardVO

`func NewDhcpGuardVO(enable bool, ) *DhcpGuardVO`

NewDhcpGuardVO instantiates a new DhcpGuardVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpGuardVOWithDefaults

`func NewDhcpGuardVOWithDefaults() *DhcpGuardVO`

NewDhcpGuardVOWithDefaults instantiates a new DhcpGuardVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpSvr1

`func (o *DhcpGuardVO) GetDhcpSvr1() string`

GetDhcpSvr1 returns the DhcpSvr1 field if non-nil, zero value otherwise.

### GetDhcpSvr1Ok

`func (o *DhcpGuardVO) GetDhcpSvr1Ok() (*string, bool)`

GetDhcpSvr1Ok returns a tuple with the DhcpSvr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSvr1

`func (o *DhcpGuardVO) SetDhcpSvr1(v string)`

SetDhcpSvr1 sets DhcpSvr1 field to given value.

### HasDhcpSvr1

`func (o *DhcpGuardVO) HasDhcpSvr1() bool`

HasDhcpSvr1 returns a boolean if a field has been set.

### GetDhcpSvr2

`func (o *DhcpGuardVO) GetDhcpSvr2() string`

GetDhcpSvr2 returns the DhcpSvr2 field if non-nil, zero value otherwise.

### GetDhcpSvr2Ok

`func (o *DhcpGuardVO) GetDhcpSvr2Ok() (*string, bool)`

GetDhcpSvr2Ok returns a tuple with the DhcpSvr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSvr2

`func (o *DhcpGuardVO) SetDhcpSvr2(v string)`

SetDhcpSvr2 sets DhcpSvr2 field to given value.

### HasDhcpSvr2

`func (o *DhcpGuardVO) HasDhcpSvr2() bool`

HasDhcpSvr2 returns a boolean if a field has been set.

### GetEnable

`func (o *DhcpGuardVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DhcpGuardVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DhcpGuardVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


