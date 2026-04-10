# Ipv4DsLiteOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AftrName** | Pointer to **string** | It is required when [dslite] is 1. Address Family Transition Router which should be a domain name or IPV6 address. | [optional] 
**Dslite** | **int32** | Parameter [dslite] should be one of the following values: 0:Auto; 1:Manual; 2:Transix; 3:Xpass; 4:V6 connect. | 

## Methods

### NewIpv4DsLiteOpenApiVO

`func NewIpv4DsLiteOpenApiVO(dslite int32, ) *Ipv4DsLiteOpenApiVO`

NewIpv4DsLiteOpenApiVO instantiates a new Ipv4DsLiteOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4DsLiteOpenApiVOWithDefaults

`func NewIpv4DsLiteOpenApiVOWithDefaults() *Ipv4DsLiteOpenApiVO`

NewIpv4DsLiteOpenApiVOWithDefaults instantiates a new Ipv4DsLiteOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAftrName

`func (o *Ipv4DsLiteOpenApiVO) GetAftrName() string`

GetAftrName returns the AftrName field if non-nil, zero value otherwise.

### GetAftrNameOk

`func (o *Ipv4DsLiteOpenApiVO) GetAftrNameOk() (*string, bool)`

GetAftrNameOk returns a tuple with the AftrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAftrName

`func (o *Ipv4DsLiteOpenApiVO) SetAftrName(v string)`

SetAftrName sets AftrName field to given value.

### HasAftrName

`func (o *Ipv4DsLiteOpenApiVO) HasAftrName() bool`

HasAftrName returns a boolean if a field has been set.

### GetDslite

`func (o *Ipv4DsLiteOpenApiVO) GetDslite() int32`

GetDslite returns the Dslite field if non-nil, zero value otherwise.

### GetDsliteOk

`func (o *Ipv4DsLiteOpenApiVO) GetDsliteOk() (*int32, bool)`

GetDsliteOk returns a tuple with the Dslite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslite

`func (o *Ipv4DsLiteOpenApiVO) SetDslite(v int32)`

SetDslite sets Dslite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


