# BatchCopySiteTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the site should contain 1 to 64 characters. | 
**TargetOmadacs** | **[]string** | The target Customer ID needs to be obtained from the interface \&quot;Obtain the customer ID with permission to modify site templates\&quot;. | 

## Methods

### NewBatchCopySiteTemplateOpenApiVO

`func NewBatchCopySiteTemplateOpenApiVO(name string, targetOmadacs []string, ) *BatchCopySiteTemplateOpenApiVO`

NewBatchCopySiteTemplateOpenApiVO instantiates a new BatchCopySiteTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchCopySiteTemplateOpenApiVOWithDefaults

`func NewBatchCopySiteTemplateOpenApiVOWithDefaults() *BatchCopySiteTemplateOpenApiVO`

NewBatchCopySiteTemplateOpenApiVOWithDefaults instantiates a new BatchCopySiteTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BatchCopySiteTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchCopySiteTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchCopySiteTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetTargetOmadacs

`func (o *BatchCopySiteTemplateOpenApiVO) GetTargetOmadacs() []string`

GetTargetOmadacs returns the TargetOmadacs field if non-nil, zero value otherwise.

### GetTargetOmadacsOk

`func (o *BatchCopySiteTemplateOpenApiVO) GetTargetOmadacsOk() (*[]string, bool)`

GetTargetOmadacsOk returns a tuple with the TargetOmadacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetOmadacs

`func (o *BatchCopySiteTemplateOpenApiVO) SetTargetOmadacs(v []string)`

SetTargetOmadacs sets TargetOmadacs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


