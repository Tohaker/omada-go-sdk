# QueryDhcpLeaseTimeParamVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | **string** | the ipAddress you want to get dhcp-lease-time. | 
**Mac** | **string** | the mac you want to get dhcp-lease-time. | 

## Methods

### NewQueryDhcpLeaseTimeParamVO

`func NewQueryDhcpLeaseTimeParamVO(ipAddress string, mac string, ) *QueryDhcpLeaseTimeParamVO`

NewQueryDhcpLeaseTimeParamVO instantiates a new QueryDhcpLeaseTimeParamVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDhcpLeaseTimeParamVOWithDefaults

`func NewQueryDhcpLeaseTimeParamVOWithDefaults() *QueryDhcpLeaseTimeParamVO`

NewQueryDhcpLeaseTimeParamVOWithDefaults instantiates a new QueryDhcpLeaseTimeParamVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *QueryDhcpLeaseTimeParamVO) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *QueryDhcpLeaseTimeParamVO) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *QueryDhcpLeaseTimeParamVO) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetMac

`func (o *QueryDhcpLeaseTimeParamVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *QueryDhcpLeaseTimeParamVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *QueryDhcpLeaseTimeParamVO) SetMac(v string)`

SetMac sets Mac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


