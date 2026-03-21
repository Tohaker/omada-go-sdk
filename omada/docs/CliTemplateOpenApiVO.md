# CliTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apply** | Pointer to **bool** | Whether the CLI configuration can be applied. Only the CLI configuration in inactive state can be applied. | [optional] 
**Description** | Pointer to **string** | CLI configuration description, it should be within the range of 0 - 256 characters. | [optional] 
**Id** | Pointer to **string** | CLI configuration ID | [optional] 
**Name** | Pointer to **string** | CLI configuration name | [optional] 
**Status** | Pointer to **int32** | CLI configuration status, it should be a value as follows: 0: active, 1: inactive | [optional] 

## Methods

### NewCliTemplateOpenApiVO

`func NewCliTemplateOpenApiVO() *CliTemplateOpenApiVO`

NewCliTemplateOpenApiVO instantiates a new CliTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCliTemplateOpenApiVOWithDefaults

`func NewCliTemplateOpenApiVOWithDefaults() *CliTemplateOpenApiVO`

NewCliTemplateOpenApiVOWithDefaults instantiates a new CliTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApply

`func (o *CliTemplateOpenApiVO) GetApply() bool`

GetApply returns the Apply field if non-nil, zero value otherwise.

### GetApplyOk

`func (o *CliTemplateOpenApiVO) GetApplyOk() (*bool, bool)`

GetApplyOk returns a tuple with the Apply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApply

`func (o *CliTemplateOpenApiVO) SetApply(v bool)`

SetApply sets Apply field to given value.

### HasApply

`func (o *CliTemplateOpenApiVO) HasApply() bool`

HasApply returns a boolean if a field has been set.

### GetDescription

`func (o *CliTemplateOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CliTemplateOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CliTemplateOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CliTemplateOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CliTemplateOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CliTemplateOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CliTemplateOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CliTemplateOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CliTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CliTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CliTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CliTemplateOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *CliTemplateOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CliTemplateOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CliTemplateOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CliTemplateOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


