# OswSdmTemplateVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InUse** | Pointer to **string** | Sdm template currently in use. | [optional] 
**Templates** | Pointer to [**[]OswSdmBriefVO**](OswSdmBriefVO.md) | All sdm templates supported by the device. | [optional] 

## Methods

### NewOswSdmTemplateVO

`func NewOswSdmTemplateVO() *OswSdmTemplateVO`

NewOswSdmTemplateVO instantiates a new OswSdmTemplateVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswSdmTemplateVOWithDefaults

`func NewOswSdmTemplateVOWithDefaults() *OswSdmTemplateVO`

NewOswSdmTemplateVOWithDefaults instantiates a new OswSdmTemplateVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInUse

`func (o *OswSdmTemplateVO) GetInUse() string`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *OswSdmTemplateVO) GetInUseOk() (*string, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *OswSdmTemplateVO) SetInUse(v string)`

SetInUse sets InUse field to given value.

### HasInUse

`func (o *OswSdmTemplateVO) HasInUse() bool`

HasInUse returns a boolean if a field has been set.

### GetTemplates

`func (o *OswSdmTemplateVO) GetTemplates() []OswSdmBriefVO`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *OswSdmTemplateVO) GetTemplatesOk() (*[]OswSdmBriefVO, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *OswSdmTemplateVO) SetTemplates(v []OswSdmBriefVO)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *OswSdmTemplateVO) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


