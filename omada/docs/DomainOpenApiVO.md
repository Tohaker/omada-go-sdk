# DomainOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Domain address, should be a valid domain address | 
**Description** | Pointer to **string** | Domain description, description should contain 1 to 512 characters. | [optional] 
**Port** | Pointer to **string** | Domain port, port should be within the range of 0-65535 or empty, e.g. 80,80-100 | [optional] 

## Methods

### NewDomainOpenApiVO

`func NewDomainOpenApiVO(address string, ) *DomainOpenApiVO`

NewDomainOpenApiVO instantiates a new DomainOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainOpenApiVOWithDefaults

`func NewDomainOpenApiVOWithDefaults() *DomainOpenApiVO`

NewDomainOpenApiVOWithDefaults instantiates a new DomainOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *DomainOpenApiVO) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DomainOpenApiVO) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DomainOpenApiVO) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetDescription

`func (o *DomainOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPort

`func (o *DomainOpenApiVO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DomainOpenApiVO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DomainOpenApiVO) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *DomainOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


