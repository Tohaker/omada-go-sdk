# SsidOverrideConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | Pointer to **[]int32** | SSID band should be a value as follows: 0: 2.4GHz; 1: 5GHz; 2: 6GHz | [optional] 
**Security** | Pointer to **int32** | SSID security mode should be a value as follows: 0: None; 2: WPA-Enterprise; 3: WPA-Personal;4: PPSK without RADIUS; 5: PPSK with RADIUS. | [optional] 
**SsidEnable** | Pointer to **bool** | Enable/disable SSID | [optional] 
**SsidId** | **int32** | SSID Entry ID | 
**SsidName** | **string** | SSID name | 
**SsidPassword** | Pointer to **string** | SSID password | [optional] 
**VlanEnable** | Pointer to **bool** | Enable/disable VLAN | [optional] 
**VlanId** | Pointer to **int32** | VLAN ID | [optional] 

## Methods

### NewSsidOverrideConfigOpenApiVO

`func NewSsidOverrideConfigOpenApiVO(ssidId int32, ssidName string, ) *SsidOverrideConfigOpenApiVO`

NewSsidOverrideConfigOpenApiVO instantiates a new SsidOverrideConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidOverrideConfigOpenApiVOWithDefaults

`func NewSsidOverrideConfigOpenApiVOWithDefaults() *SsidOverrideConfigOpenApiVO`

NewSsidOverrideConfigOpenApiVOWithDefaults instantiates a new SsidOverrideConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *SsidOverrideConfigOpenApiVO) GetBand() []int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *SsidOverrideConfigOpenApiVO) GetBandOk() (*[]int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *SsidOverrideConfigOpenApiVO) SetBand(v []int32)`

SetBand sets Band field to given value.

### HasBand

`func (o *SsidOverrideConfigOpenApiVO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetSecurity

`func (o *SsidOverrideConfigOpenApiVO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SsidOverrideConfigOpenApiVO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SsidOverrideConfigOpenApiVO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SsidOverrideConfigOpenApiVO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSsidEnable

`func (o *SsidOverrideConfigOpenApiVO) GetSsidEnable() bool`

GetSsidEnable returns the SsidEnable field if non-nil, zero value otherwise.

### GetSsidEnableOk

`func (o *SsidOverrideConfigOpenApiVO) GetSsidEnableOk() (*bool, bool)`

GetSsidEnableOk returns a tuple with the SsidEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidEnable

`func (o *SsidOverrideConfigOpenApiVO) SetSsidEnable(v bool)`

SetSsidEnable sets SsidEnable field to given value.

### HasSsidEnable

`func (o *SsidOverrideConfigOpenApiVO) HasSsidEnable() bool`

HasSsidEnable returns a boolean if a field has been set.

### GetSsidId

`func (o *SsidOverrideConfigOpenApiVO) GetSsidId() int32`

GetSsidId returns the SsidId field if non-nil, zero value otherwise.

### GetSsidIdOk

`func (o *SsidOverrideConfigOpenApiVO) GetSsidIdOk() (*int32, bool)`

GetSsidIdOk returns a tuple with the SsidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidId

`func (o *SsidOverrideConfigOpenApiVO) SetSsidId(v int32)`

SetSsidId sets SsidId field to given value.


### GetSsidName

`func (o *SsidOverrideConfigOpenApiVO) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *SsidOverrideConfigOpenApiVO) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *SsidOverrideConfigOpenApiVO) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.


### GetSsidPassword

`func (o *SsidOverrideConfigOpenApiVO) GetSsidPassword() string`

GetSsidPassword returns the SsidPassword field if non-nil, zero value otherwise.

### GetSsidPasswordOk

`func (o *SsidOverrideConfigOpenApiVO) GetSsidPasswordOk() (*string, bool)`

GetSsidPasswordOk returns a tuple with the SsidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidPassword

`func (o *SsidOverrideConfigOpenApiVO) SetSsidPassword(v string)`

SetSsidPassword sets SsidPassword field to given value.

### HasSsidPassword

`func (o *SsidOverrideConfigOpenApiVO) HasSsidPassword() bool`

HasSsidPassword returns a boolean if a field has been set.

### GetVlanEnable

`func (o *SsidOverrideConfigOpenApiVO) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *SsidOverrideConfigOpenApiVO) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *SsidOverrideConfigOpenApiVO) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.

### HasVlanEnable

`func (o *SsidOverrideConfigOpenApiVO) HasVlanEnable() bool`

HasVlanEnable returns a boolean if a field has been set.

### GetVlanId

`func (o *SsidOverrideConfigOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SsidOverrideConfigOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SsidOverrideConfigOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SsidOverrideConfigOpenApiVO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


