# CreateSiteTemplateEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the site template should contain 1 to 64 characters. | 
**Settings** | Pointer to **[]string** | Settings of the site template support. Optional settings are as follows: AIRFPlaning, firewall, wirelessIdsIps(only for pro controller), siteCLI, abnormal(only for pro controller), log, auditLogs, sim, applicationControl, servicesRebootSchedule, servicesPortSchedule | [optional] 
**Type** | Pointer to **int32** | The type should be a value as follows: 0:basic; 1:pro. This field applies to the Omada Pro Controller only. Please do not use it for non-Pro controllers. | [optional] 

## Methods

### NewCreateSiteTemplateEntity

`func NewCreateSiteTemplateEntity(name string, ) *CreateSiteTemplateEntity`

NewCreateSiteTemplateEntity instantiates a new CreateSiteTemplateEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSiteTemplateEntityWithDefaults

`func NewCreateSiteTemplateEntityWithDefaults() *CreateSiteTemplateEntity`

NewCreateSiteTemplateEntityWithDefaults instantiates a new CreateSiteTemplateEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSiteTemplateEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSiteTemplateEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSiteTemplateEntity) SetName(v string)`

SetName sets Name field to given value.


### GetSettings

`func (o *CreateSiteTemplateEntity) GetSettings() []string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CreateSiteTemplateEntity) GetSettingsOk() (*[]string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CreateSiteTemplateEntity) SetSettings(v []string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CreateSiteTemplateEntity) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetType

`func (o *CreateSiteTemplateEntity) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSiteTemplateEntity) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSiteTemplateEntity) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSiteTemplateEntity) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


