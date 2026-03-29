# ApIPSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpIpSetting** | Pointer to [**DhcpIpSettingEntity**](DhcpIpSettingEntity.md) |  | [optional] 
**Mode** | **string** | Mode should be a value as follows: Static; DHCP | 
**StaticIpSetting** | Pointer to [**StaticIpSettingEntity**](StaticIpSettingEntity.md) |  | [optional] 

## Methods

### NewApIPSetting

`func NewApIPSetting(mode string, ) *ApIPSetting`

NewApIPSetting instantiates a new ApIPSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApIPSettingWithDefaults

`func NewApIPSettingWithDefaults() *ApIPSetting`

NewApIPSettingWithDefaults instantiates a new ApIPSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpIpSetting

`func (o *ApIPSetting) GetDhcpIpSetting() DhcpIpSettingEntity`

GetDhcpIpSetting returns the DhcpIpSetting field if non-nil, zero value otherwise.

### GetDhcpIpSettingOk

`func (o *ApIPSetting) GetDhcpIpSettingOk() (*DhcpIpSettingEntity, bool)`

GetDhcpIpSettingOk returns a tuple with the DhcpIpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpIpSetting

`func (o *ApIPSetting) SetDhcpIpSetting(v DhcpIpSettingEntity)`

SetDhcpIpSetting sets DhcpIpSetting field to given value.

### HasDhcpIpSetting

`func (o *ApIPSetting) HasDhcpIpSetting() bool`

HasDhcpIpSetting returns a boolean if a field has been set.

### GetMode

`func (o *ApIPSetting) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApIPSetting) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApIPSetting) SetMode(v string)`

SetMode sets Mode field to given value.


### GetStaticIpSetting

`func (o *ApIPSetting) GetStaticIpSetting() StaticIpSettingEntity`

GetStaticIpSetting returns the StaticIpSetting field if non-nil, zero value otherwise.

### GetStaticIpSettingOk

`func (o *ApIPSetting) GetStaticIpSettingOk() (*StaticIpSettingEntity, bool)`

GetStaticIpSettingOk returns a tuple with the StaticIpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpSetting

`func (o *ApIPSetting) SetStaticIpSetting(v StaticIpSettingEntity)`

SetStaticIpSetting sets StaticIpSetting field to given value.

### HasStaticIpSetting

`func (o *ApIPSetting) HasStaticIpSetting() bool`

HasStaticIpSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


