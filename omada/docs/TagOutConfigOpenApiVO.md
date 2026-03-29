# TagOutConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagConfigList** | [**[]TagOutItemOpenApiVO**](TagOutItemOpenApiVO.md) | The Tag Outbound configuration of class type, must contains class 1, class 2, class 3 and others. | 

## Methods

### NewTagOutConfigOpenApiVO

`func NewTagOutConfigOpenApiVO(tagConfigList []TagOutItemOpenApiVO, ) *TagOutConfigOpenApiVO`

NewTagOutConfigOpenApiVO instantiates a new TagOutConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagOutConfigOpenApiVOWithDefaults

`func NewTagOutConfigOpenApiVOWithDefaults() *TagOutConfigOpenApiVO`

NewTagOutConfigOpenApiVOWithDefaults instantiates a new TagOutConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagConfigList

`func (o *TagOutConfigOpenApiVO) GetTagConfigList() []TagOutItemOpenApiVO`

GetTagConfigList returns the TagConfigList field if non-nil, zero value otherwise.

### GetTagConfigListOk

`func (o *TagOutConfigOpenApiVO) GetTagConfigListOk() (*[]TagOutItemOpenApiVO, bool)`

GetTagConfigListOk returns a tuple with the TagConfigList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagConfigList

`func (o *TagOutConfigOpenApiVO) SetTagConfigList(v []TagOutItemOpenApiVO)`

SetTagConfigList sets TagConfigList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


