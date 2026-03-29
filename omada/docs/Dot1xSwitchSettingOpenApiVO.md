# Dot1xSwitchSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountVrfId** | Pointer to **string** | Account VRF ID | [optional] 
**AuthVrfId** | Pointer to **string** | Auth VRF ID | [optional] 
**Dot1xPorts** | Pointer to **[]int32** | Switch 802.1x enabled ports | [optional] 
**MabPorts** | Pointer to **[]int32** | Switch MAB enabled ports | [optional] 
**Mac** | **string** | MAC address of the switch | 

## Methods

### NewDot1xSwitchSettingOpenApiVO

`func NewDot1xSwitchSettingOpenApiVO(mac string, ) *Dot1xSwitchSettingOpenApiVO`

NewDot1xSwitchSettingOpenApiVO instantiates a new Dot1xSwitchSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xSwitchSettingOpenApiVOWithDefaults

`func NewDot1xSwitchSettingOpenApiVOWithDefaults() *Dot1xSwitchSettingOpenApiVO`

NewDot1xSwitchSettingOpenApiVOWithDefaults instantiates a new Dot1xSwitchSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountVrfId

`func (o *Dot1xSwitchSettingOpenApiVO) GetAccountVrfId() string`

GetAccountVrfId returns the AccountVrfId field if non-nil, zero value otherwise.

### GetAccountVrfIdOk

`func (o *Dot1xSwitchSettingOpenApiVO) GetAccountVrfIdOk() (*string, bool)`

GetAccountVrfIdOk returns a tuple with the AccountVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountVrfId

`func (o *Dot1xSwitchSettingOpenApiVO) SetAccountVrfId(v string)`

SetAccountVrfId sets AccountVrfId field to given value.

### HasAccountVrfId

`func (o *Dot1xSwitchSettingOpenApiVO) HasAccountVrfId() bool`

HasAccountVrfId returns a boolean if a field has been set.

### GetAuthVrfId

`func (o *Dot1xSwitchSettingOpenApiVO) GetAuthVrfId() string`

GetAuthVrfId returns the AuthVrfId field if non-nil, zero value otherwise.

### GetAuthVrfIdOk

`func (o *Dot1xSwitchSettingOpenApiVO) GetAuthVrfIdOk() (*string, bool)`

GetAuthVrfIdOk returns a tuple with the AuthVrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthVrfId

`func (o *Dot1xSwitchSettingOpenApiVO) SetAuthVrfId(v string)`

SetAuthVrfId sets AuthVrfId field to given value.

### HasAuthVrfId

`func (o *Dot1xSwitchSettingOpenApiVO) HasAuthVrfId() bool`

HasAuthVrfId returns a boolean if a field has been set.

### GetDot1xPorts

`func (o *Dot1xSwitchSettingOpenApiVO) GetDot1xPorts() []int32`

GetDot1xPorts returns the Dot1xPorts field if non-nil, zero value otherwise.

### GetDot1xPortsOk

`func (o *Dot1xSwitchSettingOpenApiVO) GetDot1xPortsOk() (*[]int32, bool)`

GetDot1xPortsOk returns a tuple with the Dot1xPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1xPorts

`func (o *Dot1xSwitchSettingOpenApiVO) SetDot1xPorts(v []int32)`

SetDot1xPorts sets Dot1xPorts field to given value.

### HasDot1xPorts

`func (o *Dot1xSwitchSettingOpenApiVO) HasDot1xPorts() bool`

HasDot1xPorts returns a boolean if a field has been set.

### GetMabPorts

`func (o *Dot1xSwitchSettingOpenApiVO) GetMabPorts() []int32`

GetMabPorts returns the MabPorts field if non-nil, zero value otherwise.

### GetMabPortsOk

`func (o *Dot1xSwitchSettingOpenApiVO) GetMabPortsOk() (*[]int32, bool)`

GetMabPortsOk returns a tuple with the MabPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMabPorts

`func (o *Dot1xSwitchSettingOpenApiVO) SetMabPorts(v []int32)`

SetMabPorts sets MabPorts field to given value.

### HasMabPorts

`func (o *Dot1xSwitchSettingOpenApiVO) HasMabPorts() bool`

HasMabPorts returns a boolean if a field has been set.

### GetMac

`func (o *Dot1xSwitchSettingOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Dot1xSwitchSettingOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Dot1xSwitchSettingOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


