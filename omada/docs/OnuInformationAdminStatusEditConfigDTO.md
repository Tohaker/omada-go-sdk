# OnuInformationAdminStatusEditConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | **string** | Admin status should be a value as follows:ACTIVATE,DEACTIVATE | 
**Keys** | **[]string** | ONU identifier list | 

## Methods

### NewOnuInformationAdminStatusEditConfigDTO

`func NewOnuInformationAdminStatusEditConfigDTO(adminStatus string, keys []string, ) *OnuInformationAdminStatusEditConfigDTO`

NewOnuInformationAdminStatusEditConfigDTO instantiates a new OnuInformationAdminStatusEditConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuInformationAdminStatusEditConfigDTOWithDefaults

`func NewOnuInformationAdminStatusEditConfigDTOWithDefaults() *OnuInformationAdminStatusEditConfigDTO`

NewOnuInformationAdminStatusEditConfigDTOWithDefaults instantiates a new OnuInformationAdminStatusEditConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *OnuInformationAdminStatusEditConfigDTO) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *OnuInformationAdminStatusEditConfigDTO) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *OnuInformationAdminStatusEditConfigDTO) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.


### GetKeys

`func (o *OnuInformationAdminStatusEditConfigDTO) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *OnuInformationAdminStatusEditConfigDTO) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *OnuInformationAdminStatusEditConfigDTO) SetKeys(v []string)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


