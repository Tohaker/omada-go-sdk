# UrlDragSortIndexOpenapiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexes** | **map[string]int32** | The order in which items take effect, this object is a Map, the key is item ID and the value is the index you want to set. | 
**Type** | **string** | Type must be gateway or ap | 

## Methods

### NewUrlDragSortIndexOpenapiVO

`func NewUrlDragSortIndexOpenapiVO(indexes map[string]int32, type_ string, ) *UrlDragSortIndexOpenapiVO`

NewUrlDragSortIndexOpenapiVO instantiates a new UrlDragSortIndexOpenapiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUrlDragSortIndexOpenapiVOWithDefaults

`func NewUrlDragSortIndexOpenapiVOWithDefaults() *UrlDragSortIndexOpenapiVO`

NewUrlDragSortIndexOpenapiVOWithDefaults instantiates a new UrlDragSortIndexOpenapiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexes

`func (o *UrlDragSortIndexOpenapiVO) GetIndexes() map[string]int32`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *UrlDragSortIndexOpenapiVO) GetIndexesOk() (*map[string]int32, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *UrlDragSortIndexOpenapiVO) SetIndexes(v map[string]int32)`

SetIndexes sets Indexes field to given value.


### GetType

`func (o *UrlDragSortIndexOpenapiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UrlDragSortIndexOpenapiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UrlDragSortIndexOpenapiVO) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


