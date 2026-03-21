# DragSortIndexTemplateOpenapiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexes** | **map[string]int32** | The order in which items take effect, this object is a Map, the key is item ID and the value is the index you want to set. | 
**Type** | **string** | \&quot;gateway\&quot; or \&quot;eap\&quot; | 

## Methods

### NewDragSortIndexTemplateOpenapiVO

`func NewDragSortIndexTemplateOpenapiVO(indexes map[string]int32, type_ string, ) *DragSortIndexTemplateOpenapiVO`

NewDragSortIndexTemplateOpenapiVO instantiates a new DragSortIndexTemplateOpenapiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDragSortIndexTemplateOpenapiVOWithDefaults

`func NewDragSortIndexTemplateOpenapiVOWithDefaults() *DragSortIndexTemplateOpenapiVO`

NewDragSortIndexTemplateOpenapiVOWithDefaults instantiates a new DragSortIndexTemplateOpenapiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexes

`func (o *DragSortIndexTemplateOpenapiVO) GetIndexes() map[string]int32`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *DragSortIndexTemplateOpenapiVO) GetIndexesOk() (*map[string]int32, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *DragSortIndexTemplateOpenapiVO) SetIndexes(v map[string]int32)`

SetIndexes sets Indexes field to given value.


### GetType

`func (o *DragSortIndexTemplateOpenapiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DragSortIndexTemplateOpenapiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DragSortIndexTemplateOpenapiVO) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


