# SocialLoginSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledTypes** | **[]int32** | Social login enabled types, should be a value as follows: 17: Google. | 

## Methods

### NewSocialLoginSettingOpenApiVO

`func NewSocialLoginSettingOpenApiVO(enabledTypes []int32, ) *SocialLoginSettingOpenApiVO`

NewSocialLoginSettingOpenApiVO instantiates a new SocialLoginSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialLoginSettingOpenApiVOWithDefaults

`func NewSocialLoginSettingOpenApiVOWithDefaults() *SocialLoginSettingOpenApiVO`

NewSocialLoginSettingOpenApiVOWithDefaults instantiates a new SocialLoginSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledTypes

`func (o *SocialLoginSettingOpenApiVO) GetEnabledTypes() []int32`

GetEnabledTypes returns the EnabledTypes field if non-nil, zero value otherwise.

### GetEnabledTypesOk

`func (o *SocialLoginSettingOpenApiVO) GetEnabledTypesOk() (*[]int32, bool)`

GetEnabledTypesOk returns a tuple with the EnabledTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledTypes

`func (o *SocialLoginSettingOpenApiVO) SetEnabledTypes(v []int32)`

SetEnabledTypes sets EnabledTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


