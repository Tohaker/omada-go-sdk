# PrivacyPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **int32** | Status should be a value as follows: 0: disable; 1: enable | 

## Methods

### NewPrivacyPolicy

`func NewPrivacyPolicy(status int32, ) *PrivacyPolicy`

NewPrivacyPolicy instantiates a new PrivacyPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivacyPolicyWithDefaults

`func NewPrivacyPolicyWithDefaults() *PrivacyPolicy`

NewPrivacyPolicyWithDefaults instantiates a new PrivacyPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PrivacyPolicy) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivacyPolicy) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivacyPolicy) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


