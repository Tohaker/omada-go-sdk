# ApIPv6Setting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DynamicIpv6Setting** | Pointer to [**DynamicIpv6SettingEntity**](DynamicIpv6SettingEntity.md) |  | [optional] 
**Ipv6Enable** | **bool** | ipv6 enable | 
**Mode** | Pointer to **string** | Mode should be a value as follows: static; dynamic | [optional] 
**StaticIpv6Setting** | Pointer to [**StaticIpv6SettingEntity**](StaticIpv6SettingEntity.md) |  | [optional] 

## Methods

### NewApIPv6Setting

`func NewApIPv6Setting(ipv6Enable bool, ) *ApIPv6Setting`

NewApIPv6Setting instantiates a new ApIPv6Setting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApIPv6SettingWithDefaults

`func NewApIPv6SettingWithDefaults() *ApIPv6Setting`

NewApIPv6SettingWithDefaults instantiates a new ApIPv6Setting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDynamicIpv6Setting

`func (o *ApIPv6Setting) GetDynamicIpv6Setting() DynamicIpv6SettingEntity`

GetDynamicIpv6Setting returns the DynamicIpv6Setting field if non-nil, zero value otherwise.

### GetDynamicIpv6SettingOk

`func (o *ApIPv6Setting) GetDynamicIpv6SettingOk() (*DynamicIpv6SettingEntity, bool)`

GetDynamicIpv6SettingOk returns a tuple with the DynamicIpv6Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicIpv6Setting

`func (o *ApIPv6Setting) SetDynamicIpv6Setting(v DynamicIpv6SettingEntity)`

SetDynamicIpv6Setting sets DynamicIpv6Setting field to given value.

### HasDynamicIpv6Setting

`func (o *ApIPv6Setting) HasDynamicIpv6Setting() bool`

HasDynamicIpv6Setting returns a boolean if a field has been set.

### GetIpv6Enable

`func (o *ApIPv6Setting) GetIpv6Enable() bool`

GetIpv6Enable returns the Ipv6Enable field if non-nil, zero value otherwise.

### GetIpv6EnableOk

`func (o *ApIPv6Setting) GetIpv6EnableOk() (*bool, bool)`

GetIpv6EnableOk returns a tuple with the Ipv6Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Enable

`func (o *ApIPv6Setting) SetIpv6Enable(v bool)`

SetIpv6Enable sets Ipv6Enable field to given value.


### GetMode

`func (o *ApIPv6Setting) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ApIPv6Setting) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ApIPv6Setting) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *ApIPv6Setting) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetStaticIpv6Setting

`func (o *ApIPv6Setting) GetStaticIpv6Setting() StaticIpv6SettingEntity`

GetStaticIpv6Setting returns the StaticIpv6Setting field if non-nil, zero value otherwise.

### GetStaticIpv6SettingOk

`func (o *ApIPv6Setting) GetStaticIpv6SettingOk() (*StaticIpv6SettingEntity, bool)`

GetStaticIpv6SettingOk returns a tuple with the StaticIpv6Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIpv6Setting

`func (o *ApIPv6Setting) SetStaticIpv6Setting(v StaticIpv6SettingEntity)`

SetStaticIpv6Setting sets StaticIpv6Setting field to given value.

### HasStaticIpv6Setting

`func (o *ApIPv6Setting) HasStaticIpv6Setting() bool`

HasStaticIpv6Setting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


