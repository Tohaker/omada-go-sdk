# VirtualWanMacSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device MAC. When [method] is set, parameter [mac] should not be empty. | [optional] 
**Method** | **string** | Parameter [method] should be set or recover. Set: Customize MAC address; Recover: Use default MAC address. | 

## Methods

### NewVirtualWanMacSettingOpenApiVO

`func NewVirtualWanMacSettingOpenApiVO(method string, ) *VirtualWanMacSettingOpenApiVO`

NewVirtualWanMacSettingOpenApiVO instantiates a new VirtualWanMacSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanMacSettingOpenApiVOWithDefaults

`func NewVirtualWanMacSettingOpenApiVOWithDefaults() *VirtualWanMacSettingOpenApiVO`

NewVirtualWanMacSettingOpenApiVOWithDefaults instantiates a new VirtualWanMacSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *VirtualWanMacSettingOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *VirtualWanMacSettingOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *VirtualWanMacSettingOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *VirtualWanMacSettingOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMethod

`func (o *VirtualWanMacSettingOpenApiVO) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *VirtualWanMacSettingOpenApiVO) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *VirtualWanMacSettingOpenApiVO) SetMethod(v string)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


