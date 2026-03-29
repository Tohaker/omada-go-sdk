# Ipv6StaticOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** |  | 
**DefaultGateway** | **string** |  | 
**PrefixLen** | **int32** | It should be within the range of 1–128 | 
**PrimaryDns** | **string** |  | 
**SecondaryDns** | Pointer to **string** |  | [optional] 

## Methods

### NewIpv6StaticOpenApiVO

`func NewIpv6StaticOpenApiVO(address string, defaultGateway string, prefixLen int32, primaryDns string, ) *Ipv6StaticOpenApiVO`

NewIpv6StaticOpenApiVO instantiates a new Ipv6StaticOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6StaticOpenApiVOWithDefaults

`func NewIpv6StaticOpenApiVOWithDefaults() *Ipv6StaticOpenApiVO`

NewIpv6StaticOpenApiVOWithDefaults instantiates a new Ipv6StaticOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Ipv6StaticOpenApiVO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Ipv6StaticOpenApiVO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Ipv6StaticOpenApiVO) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDefaultGateway

`func (o *Ipv6StaticOpenApiVO) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *Ipv6StaticOpenApiVO) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *Ipv6StaticOpenApiVO) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.


### GetPrefixLen

`func (o *Ipv6StaticOpenApiVO) GetPrefixLen() int32`

GetPrefixLen returns the PrefixLen field if non-nil, zero value otherwise.

### GetPrefixLenOk

`func (o *Ipv6StaticOpenApiVO) GetPrefixLenOk() (*int32, bool)`

GetPrefixLenOk returns a tuple with the PrefixLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefixLen

`func (o *Ipv6StaticOpenApiVO) SetPrefixLen(v int32)`

SetPrefixLen sets PrefixLen field to given value.


### GetPrimaryDns

`func (o *Ipv6StaticOpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv6StaticOpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv6StaticOpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.


### GetSecondaryDns

`func (o *Ipv6StaticOpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *Ipv6StaticOpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *Ipv6StaticOpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *Ipv6StaticOpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


