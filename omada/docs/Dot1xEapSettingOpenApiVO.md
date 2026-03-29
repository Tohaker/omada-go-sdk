# Dot1xEapSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dot1xPorts** | Pointer to **[]string** | EAP 802.1x enabled ports | [optional] 
**Mac** | **string** | MAC address of the EAP | 

## Methods

### NewDot1xEapSettingOpenApiVO

`func NewDot1xEapSettingOpenApiVO(mac string, ) *Dot1xEapSettingOpenApiVO`

NewDot1xEapSettingOpenApiVO instantiates a new Dot1xEapSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xEapSettingOpenApiVOWithDefaults

`func NewDot1xEapSettingOpenApiVOWithDefaults() *Dot1xEapSettingOpenApiVO`

NewDot1xEapSettingOpenApiVOWithDefaults instantiates a new Dot1xEapSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDot1xPorts

`func (o *Dot1xEapSettingOpenApiVO) GetDot1xPorts() []string`

GetDot1xPorts returns the Dot1xPorts field if non-nil, zero value otherwise.

### GetDot1xPortsOk

`func (o *Dot1xEapSettingOpenApiVO) GetDot1xPortsOk() (*[]string, bool)`

GetDot1xPortsOk returns a tuple with the Dot1xPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1xPorts

`func (o *Dot1xEapSettingOpenApiVO) SetDot1xPorts(v []string)`

SetDot1xPorts sets Dot1xPorts field to given value.

### HasDot1xPorts

`func (o *Dot1xEapSettingOpenApiVO) HasDot1xPorts() bool`

HasDot1xPorts returns a boolean if a field has been set.

### GetMac

`func (o *Dot1xEapSettingOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Dot1xEapSettingOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Dot1xEapSettingOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


