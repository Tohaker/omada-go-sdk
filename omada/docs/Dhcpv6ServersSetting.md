# Dhcpv6ServersSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dhcpv6Svr1** | Pointer to **string** | DHCPv6 Server IP1 | [optional] 
**Dhcpv6Svr2** | Pointer to **string** | DHCPv6 Server IP2 | [optional] 
**Enable** | **bool** | The switch of DHCPv6 Guarding | 

## Methods

### NewDhcpv6ServersSetting

`func NewDhcpv6ServersSetting(enable bool, ) *Dhcpv6ServersSetting`

NewDhcpv6ServersSetting instantiates a new Dhcpv6ServersSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpv6ServersSettingWithDefaults

`func NewDhcpv6ServersSettingWithDefaults() *Dhcpv6ServersSetting`

NewDhcpv6ServersSettingWithDefaults instantiates a new Dhcpv6ServersSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpv6Svr1

`func (o *Dhcpv6ServersSetting) GetDhcpv6Svr1() string`

GetDhcpv6Svr1 returns the Dhcpv6Svr1 field if non-nil, zero value otherwise.

### GetDhcpv6Svr1Ok

`func (o *Dhcpv6ServersSetting) GetDhcpv6Svr1Ok() (*string, bool)`

GetDhcpv6Svr1Ok returns a tuple with the Dhcpv6Svr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Svr1

`func (o *Dhcpv6ServersSetting) SetDhcpv6Svr1(v string)`

SetDhcpv6Svr1 sets Dhcpv6Svr1 field to given value.

### HasDhcpv6Svr1

`func (o *Dhcpv6ServersSetting) HasDhcpv6Svr1() bool`

HasDhcpv6Svr1 returns a boolean if a field has been set.

### GetDhcpv6Svr2

`func (o *Dhcpv6ServersSetting) GetDhcpv6Svr2() string`

GetDhcpv6Svr2 returns the Dhcpv6Svr2 field if non-nil, zero value otherwise.

### GetDhcpv6Svr2Ok

`func (o *Dhcpv6ServersSetting) GetDhcpv6Svr2Ok() (*string, bool)`

GetDhcpv6Svr2Ok returns a tuple with the Dhcpv6Svr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpv6Svr2

`func (o *Dhcpv6ServersSetting) SetDhcpv6Svr2(v string)`

SetDhcpv6Svr2 sets Dhcpv6Svr2 field to given value.

### HasDhcpv6Svr2

`func (o *Dhcpv6ServersSetting) HasDhcpv6Svr2() bool`

HasDhcpv6Svr2 returns a boolean if a field has been set.

### GetEnable

`func (o *Dhcpv6ServersSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Dhcpv6ServersSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Dhcpv6ServersSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


