# QuerySort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SortDirection** | Pointer to **string** | Sort direction should be a value as follows:ASC,DESC | [optional] 
**SortKey** | Pointer to **string** | The key used to sort | [optional] 

## Methods

### NewQuerySort

`func NewQuerySort() *QuerySort`

NewQuerySort instantiates a new QuerySort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuerySortWithDefaults

`func NewQuerySortWithDefaults() *QuerySort`

NewQuerySortWithDefaults instantiates a new QuerySort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSortDirection

`func (o *QuerySort) GetSortDirection() string`

GetSortDirection returns the SortDirection field if non-nil, zero value otherwise.

### GetSortDirectionOk

`func (o *QuerySort) GetSortDirectionOk() (*string, bool)`

GetSortDirectionOk returns a tuple with the SortDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortDirection

`func (o *QuerySort) SetSortDirection(v string)`

SetSortDirection sets SortDirection field to given value.

### HasSortDirection

`func (o *QuerySort) HasSortDirection() bool`

HasSortDirection returns a boolean if a field has been set.

### GetSortKey

`func (o *QuerySort) GetSortKey() string`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *QuerySort) GetSortKeyOk() (*string, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *QuerySort) SetSortKey(v string)`

SetSortKey sets SortKey field to given value.

### HasSortKey

`func (o *QuerySort) HasSortKey() bool`

HasSortKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


