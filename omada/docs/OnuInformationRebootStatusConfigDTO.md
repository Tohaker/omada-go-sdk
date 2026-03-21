# OnuInformationRebootStatusConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Identifier of ONU | [optional] 
**OnlineStatus** | Pointer to **string** | Online status should be a value as follows:ONLINE,OFFLINE,UPGRADING | [optional] 

## Methods

### NewOnuInformationRebootStatusConfigDTO

`func NewOnuInformationRebootStatusConfigDTO() *OnuInformationRebootStatusConfigDTO`

NewOnuInformationRebootStatusConfigDTO instantiates a new OnuInformationRebootStatusConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuInformationRebootStatusConfigDTOWithDefaults

`func NewOnuInformationRebootStatusConfigDTOWithDefaults() *OnuInformationRebootStatusConfigDTO`

NewOnuInformationRebootStatusConfigDTOWithDefaults instantiates a new OnuInformationRebootStatusConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *OnuInformationRebootStatusConfigDTO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OnuInformationRebootStatusConfigDTO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OnuInformationRebootStatusConfigDTO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OnuInformationRebootStatusConfigDTO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetOnlineStatus

`func (o *OnuInformationRebootStatusConfigDTO) GetOnlineStatus() string`

GetOnlineStatus returns the OnlineStatus field if non-nil, zero value otherwise.

### GetOnlineStatusOk

`func (o *OnuInformationRebootStatusConfigDTO) GetOnlineStatusOk() (*string, bool)`

GetOnlineStatusOk returns a tuple with the OnlineStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineStatus

`func (o *OnuInformationRebootStatusConfigDTO) SetOnlineStatus(v string)`

SetOnlineStatus sets OnlineStatus field to given value.

### HasOnlineStatus

`func (o *OnuInformationRebootStatusConfigDTO) HasOnlineStatus() bool`

HasOnlineStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


