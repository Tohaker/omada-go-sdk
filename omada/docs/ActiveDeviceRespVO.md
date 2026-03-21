# ActiveDeviceRespVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | active device status(0: success;-2002: device already active;-2003: license not enough;-2011: license category not match device;) | [optional] 

## Methods

### NewActiveDeviceRespVO

`func NewActiveDeviceRespVO() *ActiveDeviceRespVO`

NewActiveDeviceRespVO instantiates a new ActiveDeviceRespVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDeviceRespVOWithDefaults

`func NewActiveDeviceRespVOWithDefaults() *ActiveDeviceRespVO`

NewActiveDeviceRespVOWithDefaults instantiates a new ActiveDeviceRespVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *ActiveDeviceRespVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ActiveDeviceRespVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ActiveDeviceRespVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ActiveDeviceRespVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *ActiveDeviceRespVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ActiveDeviceRespVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ActiveDeviceRespVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ActiveDeviceRespVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


