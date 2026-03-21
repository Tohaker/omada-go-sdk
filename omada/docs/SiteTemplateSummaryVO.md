# SiteTemplateSummaryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindSiteNum** | Pointer to **int64** | The number of sites to which the template is bound. | [optional] 
**Category** | Pointer to **string** | The category of the site template. | [optional] 
**Id** | Pointer to **string** | Site Template ID | [optional] 
**Name** | Pointer to **string** | Name of the site should contain 1 to 64 characters. | [optional] 
**OmadacId** | Pointer to **string** | Omada ID | [optional] 
**Settings** | Pointer to **[]string** | The settings of site template seleted. | [optional] 
**Status** | Pointer to **int32** | Site template sync sites status. | [optional] 
**Type** | Pointer to **int32** | The type should be a value as follows: 0:basic; 1:pro. | [optional] 

## Methods

### NewSiteTemplateSummaryVO

`func NewSiteTemplateSummaryVO() *SiteTemplateSummaryVO`

NewSiteTemplateSummaryVO instantiates a new SiteTemplateSummaryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTemplateSummaryVOWithDefaults

`func NewSiteTemplateSummaryVOWithDefaults() *SiteTemplateSummaryVO`

NewSiteTemplateSummaryVOWithDefaults instantiates a new SiteTemplateSummaryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindSiteNum

`func (o *SiteTemplateSummaryVO) GetBindSiteNum() int64`

GetBindSiteNum returns the BindSiteNum field if non-nil, zero value otherwise.

### GetBindSiteNumOk

`func (o *SiteTemplateSummaryVO) GetBindSiteNumOk() (*int64, bool)`

GetBindSiteNumOk returns a tuple with the BindSiteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindSiteNum

`func (o *SiteTemplateSummaryVO) SetBindSiteNum(v int64)`

SetBindSiteNum sets BindSiteNum field to given value.

### HasBindSiteNum

`func (o *SiteTemplateSummaryVO) HasBindSiteNum() bool`

HasBindSiteNum returns a boolean if a field has been set.

### GetCategory

`func (o *SiteTemplateSummaryVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *SiteTemplateSummaryVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *SiteTemplateSummaryVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *SiteTemplateSummaryVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetId

`func (o *SiteTemplateSummaryVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SiteTemplateSummaryVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SiteTemplateSummaryVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SiteTemplateSummaryVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SiteTemplateSummaryVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteTemplateSummaryVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteTemplateSummaryVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteTemplateSummaryVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *SiteTemplateSummaryVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *SiteTemplateSummaryVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *SiteTemplateSummaryVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *SiteTemplateSummaryVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetSettings

`func (o *SiteTemplateSummaryVO) GetSettings() []string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SiteTemplateSummaryVO) GetSettingsOk() (*[]string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SiteTemplateSummaryVO) SetSettings(v []string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SiteTemplateSummaryVO) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetStatus

`func (o *SiteTemplateSummaryVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SiteTemplateSummaryVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SiteTemplateSummaryVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SiteTemplateSummaryVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *SiteTemplateSummaryVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteTemplateSummaryVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteTemplateSummaryVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *SiteTemplateSummaryVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


