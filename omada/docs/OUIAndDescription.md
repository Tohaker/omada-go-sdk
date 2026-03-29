# OUIAndDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of OUI should contain 1 to 128 characters., could be empty | [optional] 
**Oui** | **string** | Organizationally unique identifier, should be upper MAC | 

## Methods

### NewOUIAndDescription

`func NewOUIAndDescription(oui string, ) *OUIAndDescription`

NewOUIAndDescription instantiates a new OUIAndDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOUIAndDescriptionWithDefaults

`func NewOUIAndDescriptionWithDefaults() *OUIAndDescription`

NewOUIAndDescriptionWithDefaults instantiates a new OUIAndDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *OUIAndDescription) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OUIAndDescription) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OUIAndDescription) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OUIAndDescription) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOui

`func (o *OUIAndDescription) GetOui() string`

GetOui returns the Oui field if non-nil, zero value otherwise.

### GetOuiOk

`func (o *OUIAndDescription) GetOuiOk() (*string, bool)`

GetOuiOk returns a tuple with the Oui field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOui

`func (o *OUIAndDescription) SetOui(v string)`

SetOui sets Oui field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


