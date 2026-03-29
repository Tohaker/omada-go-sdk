# EoGreTunnelSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GreEnable** | Pointer to **bool** | EoGRE Tunnel global config status; True: enable, false: disable. | [optional] 
**Interval** | Pointer to **int32** | EoGRE Tunnel keep interval config(Unit: Seconds); It should be within the range of 10–600. | [optional] 
**MaxCount** | Pointer to **int32** | EoGRE Tunnel max keepalive skip count config; It should be within the range of 3–10. | [optional] 
**PrimaryAddress** | Pointer to **string** | EoGRE Tunnel primary gateway IP address config | [optional] 
**SecondaryAddress** | Pointer to **string** | EoGRE Tunnel secondary gateway IP address config | [optional] 
**TunnelMtu** | Pointer to **int32** | Maximum Transmission Unit. The value of parameter [tunnelMtu] should be between [850, 1500]. | [optional] 

## Methods

### NewEoGreTunnelSettingOpenApiVO

`func NewEoGreTunnelSettingOpenApiVO() *EoGreTunnelSettingOpenApiVO`

NewEoGreTunnelSettingOpenApiVO instantiates a new EoGreTunnelSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEoGreTunnelSettingOpenApiVOWithDefaults

`func NewEoGreTunnelSettingOpenApiVOWithDefaults() *EoGreTunnelSettingOpenApiVO`

NewEoGreTunnelSettingOpenApiVOWithDefaults instantiates a new EoGreTunnelSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreEnable

`func (o *EoGreTunnelSettingOpenApiVO) GetGreEnable() bool`

GetGreEnable returns the GreEnable field if non-nil, zero value otherwise.

### GetGreEnableOk

`func (o *EoGreTunnelSettingOpenApiVO) GetGreEnableOk() (*bool, bool)`

GetGreEnableOk returns a tuple with the GreEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreEnable

`func (o *EoGreTunnelSettingOpenApiVO) SetGreEnable(v bool)`

SetGreEnable sets GreEnable field to given value.

### HasGreEnable

`func (o *EoGreTunnelSettingOpenApiVO) HasGreEnable() bool`

HasGreEnable returns a boolean if a field has been set.

### GetInterval

`func (o *EoGreTunnelSettingOpenApiVO) GetInterval() int32`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *EoGreTunnelSettingOpenApiVO) GetIntervalOk() (*int32, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *EoGreTunnelSettingOpenApiVO) SetInterval(v int32)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *EoGreTunnelSettingOpenApiVO) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetMaxCount

`func (o *EoGreTunnelSettingOpenApiVO) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *EoGreTunnelSettingOpenApiVO) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *EoGreTunnelSettingOpenApiVO) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *EoGreTunnelSettingOpenApiVO) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetPrimaryAddress

`func (o *EoGreTunnelSettingOpenApiVO) GetPrimaryAddress() string`

GetPrimaryAddress returns the PrimaryAddress field if non-nil, zero value otherwise.

### GetPrimaryAddressOk

`func (o *EoGreTunnelSettingOpenApiVO) GetPrimaryAddressOk() (*string, bool)`

GetPrimaryAddressOk returns a tuple with the PrimaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAddress

`func (o *EoGreTunnelSettingOpenApiVO) SetPrimaryAddress(v string)`

SetPrimaryAddress sets PrimaryAddress field to given value.

### HasPrimaryAddress

`func (o *EoGreTunnelSettingOpenApiVO) HasPrimaryAddress() bool`

HasPrimaryAddress returns a boolean if a field has been set.

### GetSecondaryAddress

`func (o *EoGreTunnelSettingOpenApiVO) GetSecondaryAddress() string`

GetSecondaryAddress returns the SecondaryAddress field if non-nil, zero value otherwise.

### GetSecondaryAddressOk

`func (o *EoGreTunnelSettingOpenApiVO) GetSecondaryAddressOk() (*string, bool)`

GetSecondaryAddressOk returns a tuple with the SecondaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryAddress

`func (o *EoGreTunnelSettingOpenApiVO) SetSecondaryAddress(v string)`

SetSecondaryAddress sets SecondaryAddress field to given value.

### HasSecondaryAddress

`func (o *EoGreTunnelSettingOpenApiVO) HasSecondaryAddress() bool`

HasSecondaryAddress returns a boolean if a field has been set.

### GetTunnelMtu

`func (o *EoGreTunnelSettingOpenApiVO) GetTunnelMtu() int32`

GetTunnelMtu returns the TunnelMtu field if non-nil, zero value otherwise.

### GetTunnelMtuOk

`func (o *EoGreTunnelSettingOpenApiVO) GetTunnelMtuOk() (*int32, bool)`

GetTunnelMtuOk returns a tuple with the TunnelMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelMtu

`func (o *EoGreTunnelSettingOpenApiVO) SetTunnelMtu(v int32)`

SetTunnelMtu sets TunnelMtu field to given value.

### HasTunnelMtu

`func (o *EoGreTunnelSettingOpenApiVO) HasTunnelMtu() bool`

HasTunnelMtu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


