# ChannelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Channel that device supports. | [optional] 
**DfsFlag** | Pointer to **bool** | Whether the channel is DFS channel. | [optional] 

## Methods

### NewChannelInfo

`func NewChannelInfo() *ChannelInfo`

NewChannelInfo instantiates a new ChannelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelInfoWithDefaults

`func NewChannelInfoWithDefaults() *ChannelInfo`

NewChannelInfoWithDefaults instantiates a new ChannelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ChannelInfo) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ChannelInfo) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ChannelInfo) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ChannelInfo) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDfsFlag

`func (o *ChannelInfo) GetDfsFlag() bool`

GetDfsFlag returns the DfsFlag field if non-nil, zero value otherwise.

### GetDfsFlagOk

`func (o *ChannelInfo) GetDfsFlagOk() (*bool, bool)`

GetDfsFlagOk returns a tuple with the DfsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfsFlag

`func (o *ChannelInfo) SetDfsFlag(v bool)`

SetDfsFlag sets DfsFlag field to given value.

### HasDfsFlag

`func (o *ChannelInfo) HasDfsFlag() bool`

HasDfsFlag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


