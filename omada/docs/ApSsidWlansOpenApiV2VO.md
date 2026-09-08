# ApSsidWlansOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApGroupId** | Pointer to **string** | AP Group Id | [optional] 
**ApGroupName** | Pointer to **string** | AP Group Name | [optional] 
**SsidOverrides** | [**[]SsidOverrideConfigOpenApiV2VO**](SsidOverrideConfigOpenApiV2VO.md) | SsidOverride Config List | 

## Methods

### NewApSsidWlansOpenApiV2VO

`func NewApSsidWlansOpenApiV2VO(ssidOverrides []SsidOverrideConfigOpenApiV2VO, ) *ApSsidWlansOpenApiV2VO`

NewApSsidWlansOpenApiV2VO instantiates a new ApSsidWlansOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApSsidWlansOpenApiV2VOWithDefaults

`func NewApSsidWlansOpenApiV2VOWithDefaults() *ApSsidWlansOpenApiV2VO`

NewApSsidWlansOpenApiV2VOWithDefaults instantiates a new ApSsidWlansOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApGroupId

`func (o *ApSsidWlansOpenApiV2VO) GetApGroupId() string`

GetApGroupId returns the ApGroupId field if non-nil, zero value otherwise.

### GetApGroupIdOk

`func (o *ApSsidWlansOpenApiV2VO) GetApGroupIdOk() (*string, bool)`

GetApGroupIdOk returns a tuple with the ApGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApGroupId

`func (o *ApSsidWlansOpenApiV2VO) SetApGroupId(v string)`

SetApGroupId sets ApGroupId field to given value.

### HasApGroupId

`func (o *ApSsidWlansOpenApiV2VO) HasApGroupId() bool`

HasApGroupId returns a boolean if a field has been set.

### GetApGroupName

`func (o *ApSsidWlansOpenApiV2VO) GetApGroupName() string`

GetApGroupName returns the ApGroupName field if non-nil, zero value otherwise.

### GetApGroupNameOk

`func (o *ApSsidWlansOpenApiV2VO) GetApGroupNameOk() (*string, bool)`

GetApGroupNameOk returns a tuple with the ApGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApGroupName

`func (o *ApSsidWlansOpenApiV2VO) SetApGroupName(v string)`

SetApGroupName sets ApGroupName field to given value.

### HasApGroupName

`func (o *ApSsidWlansOpenApiV2VO) HasApGroupName() bool`

HasApGroupName returns a boolean if a field has been set.

### GetSsidOverrides

`func (o *ApSsidWlansOpenApiV2VO) GetSsidOverrides() []SsidOverrideConfigOpenApiV2VO`

GetSsidOverrides returns the SsidOverrides field if non-nil, zero value otherwise.

### GetSsidOverridesOk

`func (o *ApSsidWlansOpenApiV2VO) GetSsidOverridesOk() (*[]SsidOverrideConfigOpenApiV2VO, bool)`

GetSsidOverridesOk returns a tuple with the SsidOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidOverrides

`func (o *ApSsidWlansOpenApiV2VO) SetSsidOverrides(v []SsidOverrideConfigOpenApiV2VO)`

SetSsidOverrides sets SsidOverrides field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


