# DhcpServersSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpSvr1** | Pointer to **string** | DHCP Server IP1 | [optional] 
**DhcpSvr2** | Pointer to **string** | DHCP Server IP2 | [optional] 
**Enable** | **bool** | The switch of DHCP Guarding | 

## Methods

### NewDhcpServersSetting

`func NewDhcpServersSetting(enable bool, ) *DhcpServersSetting`

NewDhcpServersSetting instantiates a new DhcpServersSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpServersSettingWithDefaults

`func NewDhcpServersSettingWithDefaults() *DhcpServersSetting`

NewDhcpServersSettingWithDefaults instantiates a new DhcpServersSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpSvr1

`func (o *DhcpServersSetting) GetDhcpSvr1() string`

GetDhcpSvr1 returns the DhcpSvr1 field if non-nil, zero value otherwise.

### GetDhcpSvr1Ok

`func (o *DhcpServersSetting) GetDhcpSvr1Ok() (*string, bool)`

GetDhcpSvr1Ok returns a tuple with the DhcpSvr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSvr1

`func (o *DhcpServersSetting) SetDhcpSvr1(v string)`

SetDhcpSvr1 sets DhcpSvr1 field to given value.

### HasDhcpSvr1

`func (o *DhcpServersSetting) HasDhcpSvr1() bool`

HasDhcpSvr1 returns a boolean if a field has been set.

### GetDhcpSvr2

`func (o *DhcpServersSetting) GetDhcpSvr2() string`

GetDhcpSvr2 returns the DhcpSvr2 field if non-nil, zero value otherwise.

### GetDhcpSvr2Ok

`func (o *DhcpServersSetting) GetDhcpSvr2Ok() (*string, bool)`

GetDhcpSvr2Ok returns a tuple with the DhcpSvr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSvr2

`func (o *DhcpServersSetting) SetDhcpSvr2(v string)`

SetDhcpSvr2 sets DhcpSvr2 field to given value.

### HasDhcpSvr2

`func (o *DhcpServersSetting) HasDhcpSvr2() bool`

HasDhcpSvr2 returns a boolean if a field has been set.

### GetEnable

`func (o *DhcpServersSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DhcpServersSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DhcpServersSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


