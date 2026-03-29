# OnuInformationDescriptionEditConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Identifier of ONU | 
**OnuDescription** | **string** | ONU description should contain 1 to 32 characters, including uppercase and lowercase letters, numbers, and the symbols -@_:/.  | 

## Methods

### NewOnuInformationDescriptionEditConfigDTO

`func NewOnuInformationDescriptionEditConfigDTO(key string, onuDescription string, ) *OnuInformationDescriptionEditConfigDTO`

NewOnuInformationDescriptionEditConfigDTO instantiates a new OnuInformationDescriptionEditConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuInformationDescriptionEditConfigDTOWithDefaults

`func NewOnuInformationDescriptionEditConfigDTOWithDefaults() *OnuInformationDescriptionEditConfigDTO`

NewOnuInformationDescriptionEditConfigDTOWithDefaults instantiates a new OnuInformationDescriptionEditConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *OnuInformationDescriptionEditConfigDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OnuInformationDescriptionEditConfigDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OnuInformationDescriptionEditConfigDTO) SetKey(v string)`

SetKey sets Key field to given value.


### GetOnuDescription

`func (o *OnuInformationDescriptionEditConfigDTO) GetOnuDescription() string`

GetOnuDescription returns the OnuDescription field if non-nil, zero value otherwise.

### GetOnuDescriptionOk

`func (o *OnuInformationDescriptionEditConfigDTO) GetOnuDescriptionOk() (*string, bool)`

GetOnuDescriptionOk returns a tuple with the OnuDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuDescription

`func (o *OnuInformationDescriptionEditConfigDTO) SetOnuDescription(v string)`

SetOnuDescription sets OnuDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


