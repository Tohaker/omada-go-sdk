# OspfDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceName** | Pointer to **string** | Device Name | [optional] 
**IsStack** | Pointer to **bool** | Indicates whether the device is a stack member device. | [optional] 
**Mac** | Pointer to **string** | Device Mac | [optional] 
**SupportDeadInterval** | Pointer to **bool** | Indicates whether the device supports interface dead interval. | [optional] 

## Methods

### NewOspfDeviceOpenApiVO

`func NewOspfDeviceOpenApiVO() *OspfDeviceOpenApiVO`

NewOspfDeviceOpenApiVO instantiates a new OspfDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfDeviceOpenApiVOWithDefaults

`func NewOspfDeviceOpenApiVOWithDefaults() *OspfDeviceOpenApiVO`

NewOspfDeviceOpenApiVOWithDefaults instantiates a new OspfDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceName

`func (o *OspfDeviceOpenApiVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *OspfDeviceOpenApiVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *OspfDeviceOpenApiVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *OspfDeviceOpenApiVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetIsStack

`func (o *OspfDeviceOpenApiVO) GetIsStack() bool`

GetIsStack returns the IsStack field if non-nil, zero value otherwise.

### GetIsStackOk

`func (o *OspfDeviceOpenApiVO) GetIsStackOk() (*bool, bool)`

GetIsStackOk returns a tuple with the IsStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStack

`func (o *OspfDeviceOpenApiVO) SetIsStack(v bool)`

SetIsStack sets IsStack field to given value.

### HasIsStack

`func (o *OspfDeviceOpenApiVO) HasIsStack() bool`

HasIsStack returns a boolean if a field has been set.

### GetMac

`func (o *OspfDeviceOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OspfDeviceOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OspfDeviceOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OspfDeviceOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetSupportDeadInterval

`func (o *OspfDeviceOpenApiVO) GetSupportDeadInterval() bool`

GetSupportDeadInterval returns the SupportDeadInterval field if non-nil, zero value otherwise.

### GetSupportDeadIntervalOk

`func (o *OspfDeviceOpenApiVO) GetSupportDeadIntervalOk() (*bool, bool)`

GetSupportDeadIntervalOk returns a tuple with the SupportDeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDeadInterval

`func (o *OspfDeviceOpenApiVO) SetSupportDeadInterval(v bool)`

SetSupportDeadInterval sets SupportDeadInterval field to given value.

### HasSupportDeadInterval

`func (o *OspfDeviceOpenApiVO) HasSupportDeadInterval() bool`

HasSupportDeadInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


