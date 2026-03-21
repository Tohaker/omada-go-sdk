# SiteTemplateBriefOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the site template | [optional] 
**Id** | Pointer to **string** | Site Template ID | [optional] 
**Name** | Pointer to **string** | Name of the site template should contain 1 to 64 characters. | [optional] 
**OmadacId** | Pointer to **string** | Omada ID | [optional] 
**Settings** | Pointer to **[]string** | The settings of the site template. | [optional] 
**Type** | Pointer to **int32** | The type should be a value as follows: 0:basic; 1:pro. | [optional] 

## Methods

### NewSiteTemplateBriefOpenApiVO

`func NewSiteTemplateBriefOpenApiVO() *SiteTemplateBriefOpenApiVO`

NewSiteTemplateBriefOpenApiVO instantiates a new SiteTemplateBriefOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTemplateBriefOpenApiVOWithDefaults

`func NewSiteTemplateBriefOpenApiVOWithDefaults() *SiteTemplateBriefOpenApiVO`

NewSiteTemplateBriefOpenApiVOWithDefaults instantiates a new SiteTemplateBriefOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *SiteTemplateBriefOpenApiVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SiteTemplateBriefOpenApiVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SiteTemplateBriefOpenApiVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SiteTemplateBriefOpenApiVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetId

`func (o *SiteTemplateBriefOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteTemplateBriefOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteTemplateBriefOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteTemplateBriefOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SiteTemplateBriefOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteTemplateBriefOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteTemplateBriefOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteTemplateBriefOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *SiteTemplateBriefOpenApiVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *SiteTemplateBriefOpenApiVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *SiteTemplateBriefOpenApiVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *SiteTemplateBriefOpenApiVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetSettings

`func (o *SiteTemplateBriefOpenApiVO) GetSettings() []string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SiteTemplateBriefOpenApiVO) GetSettingsOk() (*[]string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SiteTemplateBriefOpenApiVO) SetSettings(v []string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SiteTemplateBriefOpenApiVO) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetType

`func (o *SiteTemplateBriefOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteTemplateBriefOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteTemplateBriefOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SiteTemplateBriefOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


