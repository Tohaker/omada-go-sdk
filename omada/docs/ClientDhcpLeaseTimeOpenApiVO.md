# ClientDhcpLeaseTimeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaseTime** | Pointer to **int32** | lease time of the client | [optional] 
**Mac** | Pointer to **string** | device mac | [optional] 

## Methods

### NewClientDhcpLeaseTimeOpenApiVO

`func NewClientDhcpLeaseTimeOpenApiVO() *ClientDhcpLeaseTimeOpenApiVO`

NewClientDhcpLeaseTimeOpenApiVO instantiates a new ClientDhcpLeaseTimeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDhcpLeaseTimeOpenApiVOWithDefaults

`func NewClientDhcpLeaseTimeOpenApiVOWithDefaults() *ClientDhcpLeaseTimeOpenApiVO`

NewClientDhcpLeaseTimeOpenApiVOWithDefaults instantiates a new ClientDhcpLeaseTimeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaseTime

`func (o *ClientDhcpLeaseTimeOpenApiVO) GetLeaseTime() int32`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *ClientDhcpLeaseTimeOpenApiVO) GetLeaseTimeOk() (*int32, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *ClientDhcpLeaseTimeOpenApiVO) SetLeaseTime(v int32)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *ClientDhcpLeaseTimeOpenApiVO) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetMac

`func (o *ClientDhcpLeaseTimeOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientDhcpLeaseTimeOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientDhcpLeaseTimeOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientDhcpLeaseTimeOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


