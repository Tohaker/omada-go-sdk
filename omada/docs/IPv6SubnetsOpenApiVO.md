# IPv6SubnetsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | IPv6 description, description should contain 1 to 512 characters. | [optional] 
**Ip** | **string** | IPv6 address, should be a valid IPv6 format | 
**Prefix** | **int32** | IPv6 prefix, prefix should be within the range of 1-128 | 

## Methods

### NewIPv6SubnetsOpenApiVO

`func NewIPv6SubnetsOpenApiVO(ip string, prefix int32, ) *IPv6SubnetsOpenApiVO`

NewIPv6SubnetsOpenApiVO instantiates a new IPv6SubnetsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPv6SubnetsOpenApiVOWithDefaults

`func NewIPv6SubnetsOpenApiVOWithDefaults() *IPv6SubnetsOpenApiVO`

NewIPv6SubnetsOpenApiVOWithDefaults instantiates a new IPv6SubnetsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IPv6SubnetsOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPv6SubnetsOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPv6SubnetsOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPv6SubnetsOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIp

`func (o *IPv6SubnetsOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPv6SubnetsOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPv6SubnetsOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPrefix

`func (o *IPv6SubnetsOpenApiVO) GetPrefix() int32`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *IPv6SubnetsOpenApiVO) GetPrefixOk() (*int32, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *IPv6SubnetsOpenApiVO) SetPrefix(v int32)`

SetPrefix sets Prefix field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


