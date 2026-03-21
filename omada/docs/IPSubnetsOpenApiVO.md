# IPSubnetsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | IP description, description should contain 1 to 512 characters. | [optional] 
**Ip** | **string** | IP address, should be a valid IP format | 
**Mask** | **int32** | IP mask, mask should be within the range of 1-32 | 

## Methods

### NewIPSubnetsOpenApiVO

`func NewIPSubnetsOpenApiVO(ip string, mask int32, ) *IPSubnetsOpenApiVO`

NewIPSubnetsOpenApiVO instantiates a new IPSubnetsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSubnetsOpenApiVOWithDefaults

`func NewIPSubnetsOpenApiVOWithDefaults() *IPSubnetsOpenApiVO`

NewIPSubnetsOpenApiVOWithDefaults instantiates a new IPSubnetsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IPSubnetsOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPSubnetsOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPSubnetsOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPSubnetsOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIp

`func (o *IPSubnetsOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPSubnetsOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPSubnetsOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMask

`func (o *IPSubnetsOpenApiVO) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *IPSubnetsOpenApiVO) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *IPSubnetsOpenApiVO) SetMask(v int32)`

SetMask sets Mask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


