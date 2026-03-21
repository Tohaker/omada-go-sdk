# VoipDeleteTelephoneBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForceDelete** | **bool** | Other configurations depending on this contact (call forwarding rules for instance) will be deleted together if this field is true. Otherwise, a pre-check will be executed and an error message will be returned when this contact is already used in some other configurations. | 

## Methods

### NewVoipDeleteTelephoneBook

`func NewVoipDeleteTelephoneBook(forceDelete bool, ) *VoipDeleteTelephoneBook`

NewVoipDeleteTelephoneBook instantiates a new VoipDeleteTelephoneBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipDeleteTelephoneBookWithDefaults

`func NewVoipDeleteTelephoneBookWithDefaults() *VoipDeleteTelephoneBook`

NewVoipDeleteTelephoneBookWithDefaults instantiates a new VoipDeleteTelephoneBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForceDelete

`func (o *VoipDeleteTelephoneBook) GetForceDelete() bool`

GetForceDelete returns the ForceDelete field if non-nil, zero value otherwise.

### GetForceDeleteOk

`func (o *VoipDeleteTelephoneBook) GetForceDeleteOk() (*bool, bool)`

GetForceDeleteOk returns a tuple with the ForceDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDelete

`func (o *VoipDeleteTelephoneBook) SetForceDelete(v bool)`

SetForceDelete sets ForceDelete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


