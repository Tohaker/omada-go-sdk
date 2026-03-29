# PortSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutboundProxy** | Pointer to **string** | When parameter [provider] is 0, parameter [outboundProxy] has a default value of [0.0.0.0]. In other cases, parameter [outboundProxy] is always [0.0.0.0]. | [optional] 
**OutboundProxyPort** | Pointer to **int32** | When parameter [provider] is 0, parameter [outboundProxyPort] has a default value of [5060]. In other cases, parameter [outboundProxyPort] is always 5060. | [optional] 
**RegistrarPort** | Pointer to **int32** | When parameter [provider] is 0, parameter [registrarPort] has a default value of [5060]. In other cases, parameter [registrarPort] is always 5060. | [optional] 
**SipProxy** | Pointer to **string** | Parameter [sipProxy] should not be null when parameter [provider] is 6. When parameter [provider] is 0, parameter [sipProxy] has a default value of [0.0.0.0]. In other cases, parameter [sipProxy] is always [0.0.0.0]. | [optional] 
**SipProxyPort** | Pointer to **int32** | When parameter [provider] is 0, parameter [sipProxyPort] has a default value of [5060]. In other cases, parameter [sipProxyPort] is always 5060. | [optional] 
**ViaOutboundProxy** | Pointer to **bool** | When parameter [provider] is 0, parameter [viaOutboundProxy] has a default value of [true]. In other cases, parameter [viaOutboundProxy] is always true. | [optional] 

## Methods

### NewPortSettingVO

`func NewPortSettingVO() *PortSettingVO`

NewPortSettingVO instantiates a new PortSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortSettingVOWithDefaults

`func NewPortSettingVOWithDefaults() *PortSettingVO`

NewPortSettingVOWithDefaults instantiates a new PortSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutboundProxy

`func (o *PortSettingVO) GetOutboundProxy() string`

GetOutboundProxy returns the OutboundProxy field if non-nil, zero value otherwise.

### GetOutboundProxyOk

`func (o *PortSettingVO) GetOutboundProxyOk() (*string, bool)`

GetOutboundProxyOk returns a tuple with the OutboundProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundProxy

`func (o *PortSettingVO) SetOutboundProxy(v string)`

SetOutboundProxy sets OutboundProxy field to given value.

### HasOutboundProxy

`func (o *PortSettingVO) HasOutboundProxy() bool`

HasOutboundProxy returns a boolean if a field has been set.

### GetOutboundProxyPort

`func (o *PortSettingVO) GetOutboundProxyPort() int32`

GetOutboundProxyPort returns the OutboundProxyPort field if non-nil, zero value otherwise.

### GetOutboundProxyPortOk

`func (o *PortSettingVO) GetOutboundProxyPortOk() (*int32, bool)`

GetOutboundProxyPortOk returns a tuple with the OutboundProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundProxyPort

`func (o *PortSettingVO) SetOutboundProxyPort(v int32)`

SetOutboundProxyPort sets OutboundProxyPort field to given value.

### HasOutboundProxyPort

`func (o *PortSettingVO) HasOutboundProxyPort() bool`

HasOutboundProxyPort returns a boolean if a field has been set.

### GetRegistrarPort

`func (o *PortSettingVO) GetRegistrarPort() int32`

GetRegistrarPort returns the RegistrarPort field if non-nil, zero value otherwise.

### GetRegistrarPortOk

`func (o *PortSettingVO) GetRegistrarPortOk() (*int32, bool)`

GetRegistrarPortOk returns a tuple with the RegistrarPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrarPort

`func (o *PortSettingVO) SetRegistrarPort(v int32)`

SetRegistrarPort sets RegistrarPort field to given value.

### HasRegistrarPort

`func (o *PortSettingVO) HasRegistrarPort() bool`

HasRegistrarPort returns a boolean if a field has been set.

### GetSipProxy

`func (o *PortSettingVO) GetSipProxy() string`

GetSipProxy returns the SipProxy field if non-nil, zero value otherwise.

### GetSipProxyOk

`func (o *PortSettingVO) GetSipProxyOk() (*string, bool)`

GetSipProxyOk returns a tuple with the SipProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipProxy

`func (o *PortSettingVO) SetSipProxy(v string)`

SetSipProxy sets SipProxy field to given value.

### HasSipProxy

`func (o *PortSettingVO) HasSipProxy() bool`

HasSipProxy returns a boolean if a field has been set.

### GetSipProxyPort

`func (o *PortSettingVO) GetSipProxyPort() int32`

GetSipProxyPort returns the SipProxyPort field if non-nil, zero value otherwise.

### GetSipProxyPortOk

`func (o *PortSettingVO) GetSipProxyPortOk() (*int32, bool)`

GetSipProxyPortOk returns a tuple with the SipProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipProxyPort

`func (o *PortSettingVO) SetSipProxyPort(v int32)`

SetSipProxyPort sets SipProxyPort field to given value.

### HasSipProxyPort

`func (o *PortSettingVO) HasSipProxyPort() bool`

HasSipProxyPort returns a boolean if a field has been set.

### GetViaOutboundProxy

`func (o *PortSettingVO) GetViaOutboundProxy() bool`

GetViaOutboundProxy returns the ViaOutboundProxy field if non-nil, zero value otherwise.

### GetViaOutboundProxyOk

`func (o *PortSettingVO) GetViaOutboundProxyOk() (*bool, bool)`

GetViaOutboundProxyOk returns a tuple with the ViaOutboundProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViaOutboundProxy

`func (o *PortSettingVO) SetViaOutboundProxy(v bool)`

SetViaOutboundProxy sets ViaOutboundProxy field to given value.

### HasViaOutboundProxy

`func (o *PortSettingVO) HasViaOutboundProxy() bool`

HasViaOutboundProxy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


