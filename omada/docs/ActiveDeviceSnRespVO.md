# ActiveDeviceSnRespVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sn** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | Active device status should be a value as follows: 0: success; -1: failed; -2002: device already active; -2003: license not enough; -2011: license category not match device;)  | [optional] 

## Methods

### NewActiveDeviceSnRespVO

`func NewActiveDeviceSnRespVO() *ActiveDeviceSnRespVO`

NewActiveDeviceSnRespVO instantiates a new ActiveDeviceSnRespVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDeviceSnRespVOWithDefaults

`func NewActiveDeviceSnRespVOWithDefaults() *ActiveDeviceSnRespVO`

NewActiveDeviceSnRespVOWithDefaults instantiates a new ActiveDeviceSnRespVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSn

`func (o *ActiveDeviceSnRespVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *ActiveDeviceSnRespVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *ActiveDeviceSnRespVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *ActiveDeviceSnRespVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *ActiveDeviceSnRespVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActiveDeviceSnRespVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActiveDeviceSnRespVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActiveDeviceSnRespVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


