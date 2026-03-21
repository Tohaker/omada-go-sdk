# VoipTelephoneBookBatchSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactIds** | **[]string** | Delete the contact ID of contact person. | 
**ForceDelete** | **bool** | Other configurations depending on selected contacts (call forwarding rules for instance) will be deleted together if this field is true. Otherwise, a pre-check will be executed and an error message will be returned when this contact is already used in some other configurations. | 
**SelectType** | **string** | SpeedDialNumberType should be a value as follows: all, include, exclude. | 

## Methods

### NewVoipTelephoneBookBatchSetting

`func NewVoipTelephoneBookBatchSetting(contactIds []string, forceDelete bool, selectType string, ) *VoipTelephoneBookBatchSetting`

NewVoipTelephoneBookBatchSetting instantiates a new VoipTelephoneBookBatchSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipTelephoneBookBatchSettingWithDefaults

`func NewVoipTelephoneBookBatchSettingWithDefaults() *VoipTelephoneBookBatchSetting`

NewVoipTelephoneBookBatchSettingWithDefaults instantiates a new VoipTelephoneBookBatchSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactIds

`func (o *VoipTelephoneBookBatchSetting) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *VoipTelephoneBookBatchSetting) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *VoipTelephoneBookBatchSetting) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.


### GetForceDelete

`func (o *VoipTelephoneBookBatchSetting) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *VoipTelephoneBookBatchSetting) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *VoipTelephoneBookBatchSetting) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.


### GetSelectType

`func (o *VoipTelephoneBookBatchSetting) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *VoipTelephoneBookBatchSetting) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *VoipTelephoneBookBatchSetting) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


