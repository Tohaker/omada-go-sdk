# WanPortMacSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | It is required when [method] is 1. | [optional] 
**Method** | **int32** | Method of obtaining MAC address. 0：Use Default MAC address; 1: Customize MAC Address. | 

## Methods

### NewWanPortMacSettingOpenApiVO

`func NewWanPortMacSettingOpenApiVO(method int32, ) *WanPortMacSettingOpenApiVO`

NewWanPortMacSettingOpenApiVO instantiates a new WanPortMacSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanPortMacSettingOpenApiVOWithDefaults

`func NewWanPortMacSettingOpenApiVOWithDefaults() *WanPortMacSettingOpenApiVO`

NewWanPortMacSettingOpenApiVOWithDefaults instantiates a new WanPortMacSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *WanPortMacSettingOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *WanPortMacSettingOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *WanPortMacSettingOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *WanPortMacSettingOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMethod

`func (o *WanPortMacSettingOpenApiVO) GetMethod() int32`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WanPortMacSettingOpenApiVO) GetMethodOk() (*int32, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WanPortMacSettingOpenApiVO) SetMethod(v int32)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


