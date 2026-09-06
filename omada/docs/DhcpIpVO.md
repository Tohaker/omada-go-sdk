# DhcpIpVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | DHCP Server IP, like 192.168.0.1. | [optional] 
**Netmask** | Pointer to **string** | Parameter [netmask] should not within the range of 1-30 | [optional] 

## Methods

### NewDhcpIpVO

`func NewDhcpIpVO() *DhcpIpVO`

NewDhcpIpVO instantiates a new DhcpIpVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpIpVOWithDefaults

`func NewDhcpIpVOWithDefaults() *DhcpIpVO`

NewDhcpIpVOWithDefaults instantiates a new DhcpIpVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *DhcpIpVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DhcpIpVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DhcpIpVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DhcpIpVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetmask

`func (o *DhcpIpVO) GetNetmask() string`

GetNetmask returns the Netmask field if non-nil, zero value otherwise.

### GetNetmaskOk

`func (o *DhcpIpVO) GetNetmaskOk() (*string, bool)`

GetNetmaskOk returns a tuple with the Netmask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetmask

`func (o *DhcpIpVO) SetNetmask(v string)`

SetNetmask sets Netmask field to given value.

### HasNetmask

`func (o *DhcpIpVO) HasNetmask() bool`

HasNetmask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


