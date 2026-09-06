# DragSortIndexVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Indexes** | **map[string]int32** | DragSort indexes. The key corresponds to the rule id and the value corresponds to the new index value. | 
**OmadacId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewDragSortIndexVO

`func NewDragSortIndexVO(indexes map[string]int32, ) *DragSortIndexVO`

NewDragSortIndexVO instantiates a new DragSortIndexVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDragSortIndexVOWithDefaults

`func NewDragSortIndexVOWithDefaults() *DragSortIndexVO`

NewDragSortIndexVOWithDefaults instantiates a new DragSortIndexVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndexes

`func (o *DragSortIndexVO) GetIndexes() map[string]int32`

GetIndexes returns the Indexes field if non-nil, zero value otherwise.

### GetIndexesOk

`func (o *DragSortIndexVO) GetIndexesOk() (*map[string]int32, bool)`

GetIndexesOk returns a tuple with the Indexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexes

`func (o *DragSortIndexVO) SetIndexes(v map[string]int32)`

SetIndexes sets Indexes field to given value.


### GetOmadacId

`func (o *DragSortIndexVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *DragSortIndexVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *DragSortIndexVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *DragSortIndexVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetSiteId

`func (o *DragSortIndexVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DragSortIndexVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DragSortIndexVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DragSortIndexVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetType

`func (o *DragSortIndexVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DragSortIndexVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DragSortIndexVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DragSortIndexVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


