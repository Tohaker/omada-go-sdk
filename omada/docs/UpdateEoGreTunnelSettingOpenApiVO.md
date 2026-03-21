# UpdateEoGreTunnelSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GreEnable** | **bool** | EoGRE Tunnel global config status; True: enable, false: disable. | 
**Interval** | Pointer to **int32** | EoGRE Tunnel keep interval config(Unit: Seconds); It should be within the range of 10–600. | [optional] 
**MaxCount** | Pointer to **int32** | EoGRE Tunnel max keepalive skip count config; It should be within the range of 3–10. | [optional] 
**PrimaryAddress** | Pointer to **string** | EoGRE Tunnel primary gateway IP address config | [optional] 
**SecondaryAddress** | Pointer to **string** | EoGRE Tunnel secondary gateway IP address config | [optional] 
**TunnelMtu** | **int32** | Maximum Transmission Unit. The value of parameter [tunnelMtu] should be between [850, 1500]. | 

## Methods

### NewUpdateEoGreTunnelSettingOpenApiVO

`func NewUpdateEoGreTunnelSettingOpenApiVO(greEnable bool, tunnelMtu int32, ) *UpdateEoGreTunnelSettingOpenApiVO`

NewUpdateEoGreTunnelSettingOpenApiVO instantiates a new UpdateEoGreTunnelSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEoGreTunnelSettingOpenApiVOWithDefaults

`func NewUpdateEoGreTunnelSettingOpenApiVOWithDefaults() *UpdateEoGreTunnelSettingOpenApiVO`

NewUpdateEoGreTunnelSettingOpenApiVOWithDefaults instantiates a new UpdateEoGreTunnelSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreEnable

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetGreEnable() bool`

GetGreEnable returns the GreEnable field if non-nil, zero value otherwise.

### GetGreEnableOk

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetGreEnableOk() (*bool, bool)`

GetGreEnableOk returns a tuple with the GreEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreEnable

`func (o *UpdateEoGreTunnelSettingOpenApiVO) SetGreEnable(v bool)`

SetGreEnable sets GreEnable field to given value.


### GetInterval

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UpdateEoGreTunnelSettingOpenApiVO) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *UpdateEoGreTunnelSettingOpenApiVO) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMaxCount

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *UpdateEoGreTunnelSettingOpenApiVO) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *UpdateEoGreTunnelSettingOpenApiVO) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetPrimaryAddress

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetPrimaryAddress() string`

GetPrimaryAddress returns the PrimaryAddress field if non-nil, zero value otherwise.

### GetPrimaryAddressOk

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetPrimaryAddressOk() (*string, bool)`

GetPrimaryAddressOk returns a tuple with the PrimaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAddress

`func (o *UpdateEoGreTunnelSettingOpenApiVO) SetPrimaryAddress(v string)`

SetPrimaryAddress sets PrimaryAddress field to given value.

### HasPrimaryAddress

`func (o *UpdateEoGreTunnelSettingOpenApiVO) HasPrimaryAddress() bool`

HasPrimaryAddress returns a boolean if a field has been set.

### GetSecondaryAddress

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetSecondaryAddress() string`

GetSecondaryAddress returns the SecondaryAddress field if non-nil, zero value otherwise.

### GetSecondaryAddressOk

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetSecondaryAddressOk() (*string, bool)`

GetSecondaryAddressOk returns a tuple with the SecondaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAddress

`func (o *UpdateEoGreTunnelSettingOpenApiVO) SetSecondaryAddress(v string)`

SetSecondaryAddress sets SecondaryAddress field to given value.

### HasSecondaryAddress

`func (o *UpdateEoGreTunnelSettingOpenApiVO) HasSecondaryAddress() bool`

HasSecondaryAddress returns a boolean if a field has been set.

### GetTunnelMtu

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetTunnelMtu() int32`

GetTunnelMtu returns the TunnelMtu field if non-nil, zero value otherwise.

### GetTunnelMtuOk

`func (o *UpdateEoGreTunnelSettingOpenApiVO) GetTunnelMtuOk() (*int32, bool)`

GetTunnelMtuOk returns a tuple with the TunnelMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelMtu

`func (o *UpdateEoGreTunnelSettingOpenApiVO) SetTunnelMtu(v int32)`

SetTunnelMtu sets TunnelMtu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


