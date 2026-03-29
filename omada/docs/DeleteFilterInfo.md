# DeleteFilterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int64** | End timestamp, in seconds, such as 1682000000. | 
**SearchKey** | Pointer to **string** | Searching key for clients to delete, Searching by: client mac, client name, ssid name, network name. | [optional] 
**Start** | **int64** | Start timestamp, in seconds, such as 1682000000. | 

## Methods

### NewDeleteFilterInfo

`func NewDeleteFilterInfo(end int64, start int64, ) *DeleteFilterInfo`

NewDeleteFilterInfo instantiates a new DeleteFilterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteFilterInfoWithDefaults

`func NewDeleteFilterInfoWithDefaults() *DeleteFilterInfo`

NewDeleteFilterInfoWithDefaults instantiates a new DeleteFilterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *DeleteFilterInfo) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *DeleteFilterInfo) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *DeleteFilterInfo) SetEnd(v int64)`

SetEnd sets End field to given value.


### GetSearchKey

`func (o *DeleteFilterInfo) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *DeleteFilterInfo) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *DeleteFilterInfo) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *DeleteFilterInfo) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetStart

`func (o *DeleteFilterInfo) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *DeleteFilterInfo) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *DeleteFilterInfo) SetStart(v int64)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


