# PoeRecoverDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailPorts** | Pointer to **map[string]string** | List of ports that PoE recovery failed. | [optional] 
**Mac** | Pointer to **string** | Device MAC address. | [optional] 
**SelectPorts** | Pointer to **[]string** | The selected ports for poe recover. | [optional] 
**Status** | Pointer to **int32** | Device status: 0 - success, 1 - failed. | [optional] 
**SuccessPorts** | Pointer to **[]string** | List of ports that PoE recovery succeeded. | [optional] 

## Methods

### NewPoeRecoverDeviceOpenApiVO

`func NewPoeRecoverDeviceOpenApiVO() *PoeRecoverDeviceOpenApiVO`

NewPoeRecoverDeviceOpenApiVO instantiates a new PoeRecoverDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoeRecoverDeviceOpenApiVOWithDefaults

`func NewPoeRecoverDeviceOpenApiVOWithDefaults() *PoeRecoverDeviceOpenApiVO`

NewPoeRecoverDeviceOpenApiVOWithDefaults instantiates a new PoeRecoverDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailPorts

`func (o *PoeRecoverDeviceOpenApiVO) GetFailPorts() map[string]string`

GetFailPorts returns the FailPorts field if non-nil, zero value otherwise.

### GetFailPortsOk

`func (o *PoeRecoverDeviceOpenApiVO) GetFailPortsOk() (*map[string]string, bool)`

GetFailPortsOk returns a tuple with the FailPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailPorts

`func (o *PoeRecoverDeviceOpenApiVO) SetFailPorts(v map[string]string)`

SetFailPorts sets FailPorts field to given value.

### HasFailPorts

`func (o *PoeRecoverDeviceOpenApiVO) HasFailPorts() bool`

HasFailPorts returns a boolean if a field has been set.

### GetMac

`func (o *PoeRecoverDeviceOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *PoeRecoverDeviceOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *PoeRecoverDeviceOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *PoeRecoverDeviceOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetSelectPorts

`func (o *PoeRecoverDeviceOpenApiVO) GetSelectPorts() []string`

GetSelectPorts returns the SelectPorts field if non-nil, zero value otherwise.

### GetSelectPortsOk

`func (o *PoeRecoverDeviceOpenApiVO) GetSelectPortsOk() (*[]string, bool)`

GetSelectPortsOk returns a tuple with the SelectPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectPorts

`func (o *PoeRecoverDeviceOpenApiVO) SetSelectPorts(v []string)`

SetSelectPorts sets SelectPorts field to given value.

### HasSelectPorts

`func (o *PoeRecoverDeviceOpenApiVO) HasSelectPorts() bool`

HasSelectPorts returns a boolean if a field has been set.

### GetStatus

`func (o *PoeRecoverDeviceOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoeRecoverDeviceOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoeRecoverDeviceOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PoeRecoverDeviceOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuccessPorts

`func (o *PoeRecoverDeviceOpenApiVO) GetSuccessPorts() []string`

GetSuccessPorts returns the SuccessPorts field if non-nil, zero value otherwise.

### GetSuccessPortsOk

`func (o *PoeRecoverDeviceOpenApiVO) GetSuccessPortsOk() (*[]string, bool)`

GetSuccessPortsOk returns a tuple with the SuccessPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessPorts

`func (o *PoeRecoverDeviceOpenApiVO) SetSuccessPorts(v []string)`

SetSuccessPorts sets SuccessPorts field to given value.

### HasSuccessPorts

`func (o *PoeRecoverDeviceOpenApiVO) HasSuccessPorts() bool`

HasSuccessPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


