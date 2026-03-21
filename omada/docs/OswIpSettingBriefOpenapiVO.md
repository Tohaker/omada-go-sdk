# OswIpSettingBriefOpenapiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Static IP for mode 0, like 192.168.0.1 | [optional] 
**Mode** | **int32** | IP Setting mode. Static:0, DHCP:1 | 
**Netmask** | Pointer to **string** | IP Mask, like 255.255.255.0 | [optional] 
**Option12** | Pointer to **string** | option12 | [optional] 

## Methods

### NewOswIpSettingBriefOpenapiVO

`func NewOswIpSettingBriefOpenapiVO(mode int32, ) *OswIpSettingBriefOpenapiVO`

NewOswIpSettingBriefOpenapiVO instantiates a new OswIpSettingBriefOpenapiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswIpSettingBriefOpenapiVOWithDefaults

`func NewOswIpSettingBriefOpenapiVOWithDefaults() *OswIpSettingBriefOpenapiVO`

NewOswIpSettingBriefOpenapiVOWithDefaults instantiates a new OswIpSettingBriefOpenapiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *OswIpSettingBriefOpenapiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OswIpSettingBriefOpenapiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OswIpSettingBriefOpenapiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OswIpSettingBriefOpenapiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMode

`func (o *OswIpSettingBriefOpenapiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OswIpSettingBriefOpenapiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OswIpSettingBriefOpenapiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetNetmask

`func (o *OswIpSettingBriefOpenapiVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *OswIpSettingBriefOpenapiVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *OswIpSettingBriefOpenapiVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *OswIpSettingBriefOpenapiVO) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.

### GetOption12

`func (o *OswIpSettingBriefOpenapiVO) GetOption12() string`

GetOption12 returns the Option12 field if non-nil, zero value otherwise.

### GetOption12Ok

`func (o *OswIpSettingBriefOpenapiVO) GetOption12Ok() (*string, bool)`

GetOption12Ok returns a tuple with the Option12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption12

`func (o *OswIpSettingBriefOpenapiVO) SetOption12(v string)`

SetOption12 sets Option12 field to given value.

### HasOption12

`func (o *OswIpSettingBriefOpenapiVO) HasOption12() bool`

HasOption12 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


