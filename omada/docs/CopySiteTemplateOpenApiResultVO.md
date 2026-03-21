# CopySiteTemplateOpenApiResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCustomers** | Pointer to [**[]SiteCopyResultVO**](SiteCopyResultVO.md) | Copy failed customer related information. | [optional] 
**FailNum** | Pointer to **int32** | Number of failed copies. | [optional] 
**SuccessNum** | Pointer to **int32** | Number of successful copies. | [optional] 

## Methods

### NewCopySiteTemplateOpenApiResultVO

`func NewCopySiteTemplateOpenApiResultVO() *CopySiteTemplateOpenApiResultVO`

NewCopySiteTemplateOpenApiResultVO instantiates a new CopySiteTemplateOpenApiResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopySiteTemplateOpenApiResultVOWithDefaults

`func NewCopySiteTemplateOpenApiResultVOWithDefaults() *CopySiteTemplateOpenApiResultVO`

NewCopySiteTemplateOpenApiResultVOWithDefaults instantiates a new CopySiteTemplateOpenApiResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCustomers

`func (o *CopySiteTemplateOpenApiResultVO) GetErrorCustomers() []SiteCopyResultVO`

GetErrorCustomers returns the ErrorCustomers field if non-nil, zero value otherwise.

### GetErrorCustomersOk

`func (o *CopySiteTemplateOpenApiResultVO) GetErrorCustomersOk() (*[]SiteCopyResultVO, bool)`

GetErrorCustomersOk returns a tuple with the ErrorCustomers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCustomers

`func (o *CopySiteTemplateOpenApiResultVO) SetErrorCustomers(v []SiteCopyResultVO)`

SetErrorCustomers sets ErrorCustomers field to given value.

### HasErrorCustomers

`func (o *CopySiteTemplateOpenApiResultVO) HasErrorCustomers() bool`

HasErrorCustomers returns a boolean if a field has been set.

### GetFailNum

`func (o *CopySiteTemplateOpenApiResultVO) GetFailNum() int32`

GetFailNum returns the FailNum field if non-nil, zero value otherwise.

### GetFailNumOk

`func (o *CopySiteTemplateOpenApiResultVO) GetFailNumOk() (*int32, bool)`

GetFailNumOk returns a tuple with the FailNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailNum

`func (o *CopySiteTemplateOpenApiResultVO) SetFailNum(v int32)`

SetFailNum sets FailNum field to given value.

### HasFailNum

`func (o *CopySiteTemplateOpenApiResultVO) HasFailNum() bool`

HasFailNum returns a boolean if a field has been set.

### GetSuccessNum

`func (o *CopySiteTemplateOpenApiResultVO) GetSuccessNum() int32`

GetSuccessNum returns the SuccessNum field if non-nil, zero value otherwise.

### GetSuccessNumOk

`func (o *CopySiteTemplateOpenApiResultVO) GetSuccessNumOk() (*int32, bool)`

GetSuccessNumOk returns a tuple with the SuccessNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessNum

`func (o *CopySiteTemplateOpenApiResultVO) SetSuccessNum(v int32)`

SetSuccessNum sets SuccessNum field to given value.

### HasSuccessNum

`func (o *CopySiteTemplateOpenApiResultVO) HasSuccessNum() bool`

HasSuccessNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


