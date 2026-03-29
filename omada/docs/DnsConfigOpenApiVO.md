# DnsConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | **string** | Primary DNS | 
**Secondary** | Pointer to **string** | Secondary DNS | [optional] 

## Methods

### NewDnsConfigOpenApiVO

`func NewDnsConfigOpenApiVO(primary string, ) *DnsConfigOpenApiVO`

NewDnsConfigOpenApiVO instantiates a new DnsConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsConfigOpenApiVOWithDefaults

`func NewDnsConfigOpenApiVOWithDefaults() *DnsConfigOpenApiVO`

NewDnsConfigOpenApiVOWithDefaults instantiates a new DnsConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *DnsConfigOpenApiVO) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *DnsConfigOpenApiVO) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *DnsConfigOpenApiVO) SetPrimary(v string)`

SetPrimary sets Primary field to given value.


### GetSecondary

`func (o *DnsConfigOpenApiVO) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *DnsConfigOpenApiVO) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *DnsConfigOpenApiVO) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *DnsConfigOpenApiVO) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


