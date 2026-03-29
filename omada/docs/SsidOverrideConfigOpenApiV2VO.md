# SsidOverrideConfigOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | Pointer to **[]int32** | SSID band should be a value as follows: 0: 2.4GHz; 1: 5GHz; 2: 6GHz | [optional] 
**Security** | Pointer to **int32** | SSID security mode should be a value as follows: 0: None; 2: WPA-Enterprise; 3: WPA-Personal;4: PPSK without RADIUS; 5: PPSK with RADIUS. | [optional] 
**SsidEnable** | Pointer to **bool** | Enable/disable SSID | [optional] 
**SsidEntryId** | Pointer to **int32** | SSID Entry ID | [optional] 
**SsidId** | **string** | SSID ID | 
**SsidName** | **string** | SSID name. It should contain 1 to 32 UTF-8 characters. | 
**SsidPassword** | Pointer to **string** | SSID password. It should contain 8-63 printable ASCII characters or 8-63 hexadecimal digits. | [optional] 
**VlanEnable** | Pointer to **bool** | Enable/disable VLAN | [optional] 
**VlanId** | Pointer to **int32** | VLAN ID | [optional] 
**VlanPoolIds** | Pointer to **string** | VLAN POOL IDS | [optional] 

## Methods

### NewSsidOverrideConfigOpenApiV2VO

`func NewSsidOverrideConfigOpenApiV2VO(ssidId string, ssidName string, ) *SsidOverrideConfigOpenApiV2VO`

NewSsidOverrideConfigOpenApiV2VO instantiates a new SsidOverrideConfigOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidOverrideConfigOpenApiV2VOWithDefaults

`func NewSsidOverrideConfigOpenApiV2VOWithDefaults() *SsidOverrideConfigOpenApiV2VO`

NewSsidOverrideConfigOpenApiV2VOWithDefaults instantiates a new SsidOverrideConfigOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *SsidOverrideConfigOpenApiV2VO) GetBand() []int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetBandOk() (*[]int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *SsidOverrideConfigOpenApiV2VO) SetBand(v []int32)`

SetBand sets Band field to given value.

### HasBand

`func (o *SsidOverrideConfigOpenApiV2VO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetSecurity

`func (o *SsidOverrideConfigOpenApiV2VO) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *SsidOverrideConfigOpenApiV2VO) SetSecurity(v int32)`

SetSecurity sets Security field to given value.

### HasSecurity

`func (o *SsidOverrideConfigOpenApiV2VO) HasSecurity() bool`

HasSecurity returns a boolean if a field has been set.

### GetSsidEnable

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidEnable() bool`

GetSsidEnable returns the SsidEnable field if non-nil, zero value otherwise.

### GetSsidEnableOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidEnableOk() (*bool, bool)`

GetSsidEnableOk returns a tuple with the SsidEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidEnable

`func (o *SsidOverrideConfigOpenApiV2VO) SetSsidEnable(v bool)`

SetSsidEnable sets SsidEnable field to given value.

### HasSsidEnable

`func (o *SsidOverrideConfigOpenApiV2VO) HasSsidEnable() bool`

HasSsidEnable returns a boolean if a field has been set.

### GetSsidEntryId

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidEntryId() int32`

GetSsidEntryId returns the SsidEntryId field if non-nil, zero value otherwise.

### GetSsidEntryIdOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidEntryIdOk() (*int32, bool)`

GetSsidEntryIdOk returns a tuple with the SsidEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidEntryId

`func (o *SsidOverrideConfigOpenApiV2VO) SetSsidEntryId(v int32)`

SetSsidEntryId sets SsidEntryId field to given value.

### HasSsidEntryId

`func (o *SsidOverrideConfigOpenApiV2VO) HasSsidEntryId() bool`

HasSsidEntryId returns a boolean if a field has been set.

### GetSsidId

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidId() string`

GetSsidId returns the SsidId field if non-nil, zero value otherwise.

### GetSsidIdOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidIdOk() (*string, bool)`

GetSsidIdOk returns a tuple with the SsidId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidId

`func (o *SsidOverrideConfigOpenApiV2VO) SetSsidId(v string)`

SetSsidId sets SsidId field to given value.


### GetSsidName

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidName() string`

GetSsidName returns the SsidName field if non-nil, zero value otherwise.

### GetSsidNameOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidNameOk() (*string, bool)`

GetSsidNameOk returns a tuple with the SsidName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidName

`func (o *SsidOverrideConfigOpenApiV2VO) SetSsidName(v string)`

SetSsidName sets SsidName field to given value.


### GetSsidPassword

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidPassword() string`

GetSsidPassword returns the SsidPassword field if non-nil, zero value otherwise.

### GetSsidPasswordOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetSsidPasswordOk() (*string, bool)`

GetSsidPasswordOk returns a tuple with the SsidPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidPassword

`func (o *SsidOverrideConfigOpenApiV2VO) SetSsidPassword(v string)`

SetSsidPassword sets SsidPassword field to given value.

### HasSsidPassword

`func (o *SsidOverrideConfigOpenApiV2VO) HasSsidPassword() bool`

HasSsidPassword returns a boolean if a field has been set.

### GetVlanEnable

`func (o *SsidOverrideConfigOpenApiV2VO) GetVlanEnable() bool`

GetVlanEnable returns the VlanEnable field if non-nil, zero value otherwise.

### GetVlanEnableOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetVlanEnableOk() (*bool, bool)`

GetVlanEnableOk returns a tuple with the VlanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanEnable

`func (o *SsidOverrideConfigOpenApiV2VO) SetVlanEnable(v bool)`

SetVlanEnable sets VlanEnable field to given value.

### HasVlanEnable

`func (o *SsidOverrideConfigOpenApiV2VO) HasVlanEnable() bool`

HasVlanEnable returns a boolean if a field has been set.

### GetVlanId

`func (o *SsidOverrideConfigOpenApiV2VO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *SsidOverrideConfigOpenApiV2VO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *SsidOverrideConfigOpenApiV2VO) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanPoolIds

`func (o *SsidOverrideConfigOpenApiV2VO) GetVlanPoolIds() string`

GetVlanPoolIds returns the VlanPoolIds field if non-nil, zero value otherwise.

### GetVlanPoolIdsOk

`func (o *SsidOverrideConfigOpenApiV2VO) GetVlanPoolIdsOk() (*string, bool)`

GetVlanPoolIdsOk returns a tuple with the VlanPoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPoolIds

`func (o *SsidOverrideConfigOpenApiV2VO) SetVlanPoolIds(v string)`

SetVlanPoolIds sets VlanPoolIds field to given value.

### HasVlanPoolIds

`func (o *SsidOverrideConfigOpenApiV2VO) HasVlanPoolIds() bool`

HasVlanPoolIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


