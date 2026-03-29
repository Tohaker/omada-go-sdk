# PreAuthAccessPolicyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdInt** | Pointer to **int32** | Entry ID of the policy. Except for newly added policies, this parameter should be retained | [optional] 
**Ip** | Pointer to **string** | IP Address of Pre-Authentication Access | [optional] 
**SubnetMask** | Pointer to **int32** | Subnet mask of Pre-Authentication Access. It should be within the range of 1-32 | [optional] 
**Type** | **int32** | Type of the policy. It should be a value as follows: 1: Destination IP Range, and parameter [ip] and [subnetMask] is needed. 2: URL, and parameter [url] is needed | 
**Url** | Pointer to **string** | URL of Pre-Authentication Access | [optional] 

## Methods

### NewPreAuthAccessPolicyOpenApiVO

`func NewPreAuthAccessPolicyOpenApiVO(type_ int32, ) *PreAuthAccessPolicyOpenApiVO`

NewPreAuthAccessPolicyOpenApiVO instantiates a new PreAuthAccessPolicyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreAuthAccessPolicyOpenApiVOWithDefaults

`func NewPreAuthAccessPolicyOpenApiVOWithDefaults() *PreAuthAccessPolicyOpenApiVO`

NewPreAuthAccessPolicyOpenApiVOWithDefaults instantiates a new PreAuthAccessPolicyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdInt

`func (o *PreAuthAccessPolicyOpenApiVO) GetIdInt() int32`

GetIdInt returns the IdInt field if non-nil, zero value otherwise.

### GetIdIntOk

`func (o *PreAuthAccessPolicyOpenApiVO) GetIdIntOk() (*int32, bool)`

GetIdIntOk returns a tuple with the IdInt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdInt

`func (o *PreAuthAccessPolicyOpenApiVO) SetIdInt(v int32)`

SetIdInt sets IdInt field to given value.

### HasIdInt

`func (o *PreAuthAccessPolicyOpenApiVO) HasIdInt() bool`

HasIdInt returns a boolean if a field has been set.

### GetIp

`func (o *PreAuthAccessPolicyOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *PreAuthAccessPolicyOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *PreAuthAccessPolicyOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *PreAuthAccessPolicyOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSubnetMask

`func (o *PreAuthAccessPolicyOpenApiVO) GetSubnetMask() int32`

GetSubnetMask returns the SubnetMask field if non-nil, zero value otherwise.

### GetSubnetMaskOk

`func (o *PreAuthAccessPolicyOpenApiVO) GetSubnetMaskOk() (*int32, bool)`

GetSubnetMaskOk returns a tuple with the SubnetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetMask

`func (o *PreAuthAccessPolicyOpenApiVO) SetSubnetMask(v int32)`

SetSubnetMask sets SubnetMask field to given value.

### HasSubnetMask

`func (o *PreAuthAccessPolicyOpenApiVO) HasSubnetMask() bool`

HasSubnetMask returns a boolean if a field has been set.

### GetType

`func (o *PreAuthAccessPolicyOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PreAuthAccessPolicyOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PreAuthAccessPolicyOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.


### GetUrl

`func (o *PreAuthAccessPolicyOpenApiVO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PreAuthAccessPolicyOpenApiVO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PreAuthAccessPolicyOpenApiVO) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PreAuthAccessPolicyOpenApiVO) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


