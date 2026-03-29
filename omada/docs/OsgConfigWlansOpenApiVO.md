# OsgConfigWlansOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidOverrides** | [**[]OsgSsidOverrideOpenApiVO**](OsgSsidOverrideOpenApiVO.md) | Overrided SSID List | 
**WlanId** | Pointer to **string** | WLAN ID | [optional] 
**WlansResource** | Pointer to **int32** | Wlans Resource | [optional] 

## Methods

### NewOsgConfigWlansOpenApiVO

`func NewOsgConfigWlansOpenApiVO(ssidOverrides []OsgSsidOverrideOpenApiVO, ) *OsgConfigWlansOpenApiVO`

NewOsgConfigWlansOpenApiVO instantiates a new OsgConfigWlansOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgConfigWlansOpenApiVOWithDefaults

`func NewOsgConfigWlansOpenApiVOWithDefaults() *OsgConfigWlansOpenApiVO`

NewOsgConfigWlansOpenApiVOWithDefaults instantiates a new OsgConfigWlansOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidOverrides

`func (o *OsgConfigWlansOpenApiVO) GetSsidOverrides() []OsgSsidOverrideOpenApiVO`

GetSsidOverrides returns the SsidOverrides field if non-nil, zero value otherwise.

### GetSsidOverridesOk

`func (o *OsgConfigWlansOpenApiVO) GetSsidOverridesOk() (*[]OsgSsidOverrideOpenApiVO, bool)`

GetSsidOverridesOk returns a tuple with the SsidOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidOverrides

`func (o *OsgConfigWlansOpenApiVO) SetSsidOverrides(v []OsgSsidOverrideOpenApiVO)`

SetSsidOverrides sets SsidOverrides field to given value.


### GetWlanId

`func (o *OsgConfigWlansOpenApiVO) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *OsgConfigWlansOpenApiVO) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *OsgConfigWlansOpenApiVO) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *OsgConfigWlansOpenApiVO) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.

### GetWlansResource

`func (o *OsgConfigWlansOpenApiVO) GetWlansResource() int32`

GetWlansResource returns the WlansResource field if non-nil, zero value otherwise.

### GetWlansResourceOk

`func (o *OsgConfigWlansOpenApiVO) GetWlansResourceOk() (*int32, bool)`

GetWlansResourceOk returns a tuple with the WlansResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlansResource

`func (o *OsgConfigWlansOpenApiVO) SetWlansResource(v int32)`

SetWlansResource sets WlansResource field to given value.

### HasWlansResource

`func (o *OsgConfigWlansOpenApiVO) HasWlansResource() bool`

HasWlansResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


