# OswLoopbackInterfaceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP address. | 
**LoopbackId** | **int32** | Loopback ID. | 
**Name** | **string** | Loopback Interface name. | 
**Status** | **int32** | Interface status.0: disable, 1: enable. | 
**VrfId** | Pointer to **string** | VRF ID. This field is non-empty when the device supports VRF. | [optional] 

## Methods

### NewOswLoopbackInterfaceVO

`func NewOswLoopbackInterfaceVO(ip string, loopbackId int32, name string, status int32, ) *OswLoopbackInterfaceVO`

NewOswLoopbackInterfaceVO instantiates a new OswLoopbackInterfaceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswLoopbackInterfaceVOWithDefaults

`func NewOswLoopbackInterfaceVOWithDefaults() *OswLoopbackInterfaceVO`

NewOswLoopbackInterfaceVOWithDefaults instantiates a new OswLoopbackInterfaceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *OswLoopbackInterfaceVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswLoopbackInterfaceVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswLoopbackInterfaceVO) SetIp(v string)`

SetIp sets Ip field to given value.


### GetLoopbackId

`func (o *OswLoopbackInterfaceVO) GetLoopbackId() int32`

GetLoopbackId returns the LoopbackId field if non-nil, zero value otherwise.

### GetLoopbackIdOk

`func (o *OswLoopbackInterfaceVO) GetLoopbackIdOk() (*int32, bool)`

GetLoopbackIdOk returns a tuple with the LoopbackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackId

`func (o *OswLoopbackInterfaceVO) SetLoopbackId(v int32)`

SetLoopbackId sets LoopbackId field to given value.


### GetName

`func (o *OswLoopbackInterfaceVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswLoopbackInterfaceVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswLoopbackInterfaceVO) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *OswLoopbackInterfaceVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswLoopbackInterfaceVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswLoopbackInterfaceVO) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetVrfId

`func (o *OswLoopbackInterfaceVO) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *OswLoopbackInterfaceVO) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *OswLoopbackInterfaceVO) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *OswLoopbackInterfaceVO) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


