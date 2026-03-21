# Ipv4Connection2OpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGateway** | Pointer to **string** | (Optional) It is required when [secondaryProtoType] is 0. | [optional] 
**IpAddress** | Pointer to **string** | (Optional) It is required when [secondaryProtoType] is 0. | [optional] 
**PrimaryDns** | Pointer to **string** | (Optional) It is required when [secondaryProtoType] is 0. | [optional] 
**SecondaryDns** | Pointer to **string** | (Optional) It is required when [secondaryProtoType] is 0. | [optional] 
**SecondaryProtoType** | **int32** | It should be a value as follows: 0:Static IP; 1:Dynamic IP; 2: None(Only for PPPoE). | 
**Server** | Pointer to **string** | (Optional) VPN Server/Domain Name. It is required for L2TP/PPTP. | [optional] 
**SubnetMask** | Pointer to **string** | (Optional) It is required when [secondaryProtoType] is 0. | [optional] 

## Methods

### NewIpv4Connection2OpenApiVO

`func NewIpv4Connection2OpenApiVO(secondaryProtoType int32, ) *Ipv4Connection2OpenApiVO`

NewIpv4Connection2OpenApiVO instantiates a new Ipv4Connection2OpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4Connection2OpenApiVOWithDefaults

`func NewIpv4Connection2OpenApiVOWithDefaults() *Ipv4Connection2OpenApiVO`

NewIpv4Connection2OpenApiVOWithDefaults instantiates a new Ipv4Connection2OpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultGateway

`func (o *Ipv4Connection2OpenApiVO) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *Ipv4Connection2OpenApiVO) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *Ipv4Connection2OpenApiVO) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *Ipv4Connection2OpenApiVO) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetIpAddress

`func (o *Ipv4Connection2OpenApiVO) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *Ipv4Connection2OpenApiVO) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *Ipv4Connection2OpenApiVO) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *Ipv4Connection2OpenApiVO) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetPrimaryDns

`func (o *Ipv4Connection2OpenApiVO) GetPrimaryDns() string`

GetPrimaryDns returns the PrimaryDns field if non-nil, zero value otherwise.

### GetPrimaryDnsOk

`func (o *Ipv4Connection2OpenApiVO) GetPrimaryDnsOk() (*string, bool)`

GetPrimaryDnsOk returns a tuple with the PrimaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryDns

`func (o *Ipv4Connection2OpenApiVO) SetPrimaryDns(v string)`

SetPrimaryDns sets PrimaryDns field to given value.

### HasPrimaryDns

`func (o *Ipv4Connection2OpenApiVO) HasPrimaryDns() bool`

HasPrimaryDns returns a boolean if a field has been set.

### GetSecondaryDns

`func (o *Ipv4Connection2OpenApiVO) GetSecondaryDns() string`

GetSecondaryDns returns the SecondaryDns field if non-nil, zero value otherwise.

### GetSecondaryDnsOk

`func (o *Ipv4Connection2OpenApiVO) GetSecondaryDnsOk() (*string, bool)`

GetSecondaryDnsOk returns a tuple with the SecondaryDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDns

`func (o *Ipv4Connection2OpenApiVO) SetSecondaryDns(v string)`

SetSecondaryDns sets SecondaryDns field to given value.

### HasSecondaryDns

`func (o *Ipv4Connection2OpenApiVO) HasSecondaryDns() bool`

HasSecondaryDns returns a boolean if a field has been set.

### GetSecondaryProtoType

`func (o *Ipv4Connection2OpenApiVO) GetSecondaryProtoType() int32`

GetSecondaryProtoType returns the SecondaryProtoType field if non-nil, zero value otherwise.

### GetSecondaryProtoTypeOk

`func (o *Ipv4Connection2OpenApiVO) GetSecondaryProtoTypeOk() (*int32, bool)`

GetSecondaryProtoTypeOk returns a tuple with the SecondaryProtoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryProtoType

`func (o *Ipv4Connection2OpenApiVO) SetSecondaryProtoType(v int32)`

SetSecondaryProtoType sets SecondaryProtoType field to given value.


### GetServer

`func (o *Ipv4Connection2OpenApiVO) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *Ipv4Connection2OpenApiVO) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *Ipv4Connection2OpenApiVO) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *Ipv4Connection2OpenApiVO) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSubnetMask

`func (o *Ipv4Connection2OpenApiVO) GetSubnetMask() string`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *Ipv4Connection2OpenApiVO) GetSubnetMaskOk() (*string, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *Ipv4Connection2OpenApiVO) SetSubnetMask(v string)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *Ipv4Connection2OpenApiVO) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


