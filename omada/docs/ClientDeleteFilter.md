# ClientDeleteFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Block** | Pointer to **bool** | Block, filter by block field if provided. | [optional] 
**ConnectSuccess** | Pointer to **bool** | Connect success, filter by connectSuccess field if provided. | [optional] 
**End** | Pointer to **int64** | End timestamp, in seconds, such as 1682000000. | [optional] 
**Guest** | Pointer to **bool** | Guest, filter by guest field if provided. | [optional] 
**RateLimit** | Pointer to **bool** | RateLimit, filter by rateLimit field if provided. | [optional] 
**SearchKey** | Pointer to **string** | Searching for clients to delete, Searching by client mac, client name. | [optional] 
**Start** | Pointer to **int64** | Start timestamp, in seconds, such as 1682000000. | [optional] 
**Wireless** | Pointer to **bool** | Wireless, filter by wireless field if provided. | [optional] 

## Methods

### NewClientDeleteFilter

`func NewClientDeleteFilter() *ClientDeleteFilter`

NewClientDeleteFilter instantiates a new ClientDeleteFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientDeleteFilterWithDefaults

`func NewClientDeleteFilterWithDefaults() *ClientDeleteFilter`

NewClientDeleteFilterWithDefaults instantiates a new ClientDeleteFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlock

`func (o *ClientDeleteFilter) GetBlock() bool`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ClientDeleteFilter) GetBlockOk() (*bool, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ClientDeleteFilter) SetBlock(v bool)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *ClientDeleteFilter) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetConnectSuccess

`func (o *ClientDeleteFilter) GetConnectSuccess() bool`

GetConnectSuccess returns the ConnectSuccess field if non-nil, zero value otherwise.

### GetConnectSuccessOk

`func (o *ClientDeleteFilter) GetConnectSuccessOk() (*bool, bool)`

GetConnectSuccessOk returns a tuple with the ConnectSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectSuccess

`func (o *ClientDeleteFilter) SetConnectSuccess(v bool)`

SetConnectSuccess sets ConnectSuccess field to given value.

### HasConnectSuccess

`func (o *ClientDeleteFilter) HasConnectSuccess() bool`

HasConnectSuccess returns a boolean if a field has been set.

### GetEnd

`func (o *ClientDeleteFilter) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ClientDeleteFilter) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ClientDeleteFilter) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ClientDeleteFilter) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetGuest

`func (o *ClientDeleteFilter) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *ClientDeleteFilter) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *ClientDeleteFilter) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *ClientDeleteFilter) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetRateLimit

`func (o *ClientDeleteFilter) GetRateLimit() bool`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ClientDeleteFilter) GetRateLimitOk() (*bool, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ClientDeleteFilter) SetRateLimit(v bool)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *ClientDeleteFilter) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetSearchKey

`func (o *ClientDeleteFilter) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *ClientDeleteFilter) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *ClientDeleteFilter) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *ClientDeleteFilter) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetStart

`func (o *ClientDeleteFilter) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ClientDeleteFilter) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ClientDeleteFilter) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *ClientDeleteFilter) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetWireless

`func (o *ClientDeleteFilter) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *ClientDeleteFilter) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *ClientDeleteFilter) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *ClientDeleteFilter) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


