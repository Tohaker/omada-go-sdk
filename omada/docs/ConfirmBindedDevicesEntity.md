# ConfirmBindedDevicesEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileIds** | **[]string** | Provider profile IDs | 
**SkipConfirm** | **bool** | skipConfirm indicates whether to skip the query of the devices bound to provider profiles. false: Not to skip the query. true: Skip the query and delete the provider profiles corresponding to the profileIds. | 

## Methods

### NewConfirmBindedDevicesEntity

`func NewConfirmBindedDevicesEntity(profileIds []string, skipConfirm bool, ) *ConfirmBindedDevicesEntity`

NewConfirmBindedDevicesEntity instantiates a new ConfirmBindedDevicesEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmBindedDevicesEntityWithDefaults

`func NewConfirmBindedDevicesEntityWithDefaults() *ConfirmBindedDevicesEntity`

NewConfirmBindedDevicesEntityWithDefaults instantiates a new ConfirmBindedDevicesEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileIds

`func (o *ConfirmBindedDevicesEntity) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *ConfirmBindedDevicesEntity) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *ConfirmBindedDevicesEntity) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.


### GetSkipConfirm

`func (o *ConfirmBindedDevicesEntity) GetSkipConfirm() bool`

GetSkipConfirm returns the SkipConfirm field if non-nil, zero value otherwise.

### GetSkipConfirmOk

`func (o *ConfirmBindedDevicesEntity) GetSkipConfirmOk() (*bool, bool)`

GetSkipConfirmOk returns a tuple with the SkipConfirm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipConfirm

`func (o *ConfirmBindedDevicesEntity) SetSkipConfirm(v bool)`

SetSkipConfirm sets SkipConfirm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


