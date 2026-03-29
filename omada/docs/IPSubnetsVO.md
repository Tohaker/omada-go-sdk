# IPSubnetsVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description should contain 1 to 512 characters. | [optional] 
**Ip** | **string** | IP address, should be a valid IP format | 
**Mask** | **int32** | IP mask, mask should be within the range of 1-32 | 

## Methods

### NewIPSubnetsVO

`func NewIPSubnetsVO(ip string, mask int32, ) *IPSubnetsVO`

NewIPSubnetsVO instantiates a new IPSubnetsVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSubnetsVOWithDefaults

`func NewIPSubnetsVOWithDefaults() *IPSubnetsVO`

NewIPSubnetsVOWithDefaults instantiates a new IPSubnetsVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IPSubnetsVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPSubnetsVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPSubnetsVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPSubnetsVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIp

`func (o *IPSubnetsVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPSubnetsVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPSubnetsVO) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMask

`func (o *IPSubnetsVO) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *IPSubnetsVO) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *IPSubnetsVO) SetMask(v int32)`

SetMask sets Mask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


