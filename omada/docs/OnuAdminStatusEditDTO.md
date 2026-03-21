# OnuAdminStatusEditDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | **string** | Admin status should be a value as follows:ACTIVATE,DEACTIVATE | 
**Keys** | **[]string** | Entry identifier list | 

## Methods

### NewOnuAdminStatusEditDTO

`func NewOnuAdminStatusEditDTO(adminStatus string, keys []string, ) *OnuAdminStatusEditDTO`

NewOnuAdminStatusEditDTO instantiates a new OnuAdminStatusEditDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuAdminStatusEditDTOWithDefaults

`func NewOnuAdminStatusEditDTOWithDefaults() *OnuAdminStatusEditDTO`

NewOnuAdminStatusEditDTOWithDefaults instantiates a new OnuAdminStatusEditDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *OnuAdminStatusEditDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *OnuAdminStatusEditDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *OnuAdminStatusEditDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.


### GetKeys

`func (o *OnuAdminStatusEditDTO) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *OnuAdminStatusEditDTO) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *OnuAdminStatusEditDTO) SetKeys(v []string)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


