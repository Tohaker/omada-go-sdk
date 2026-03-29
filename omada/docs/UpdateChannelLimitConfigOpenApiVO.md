# UpdateChannelLimitConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimitType** | **int32** | Channel limit enable status. It should be a value as follows: 0: default, 1: false, 2: true. | 

## Methods

### NewUpdateChannelLimitConfigOpenApiVO

`func NewUpdateChannelLimitConfigOpenApiVO(channelLimitType int32, ) *UpdateChannelLimitConfigOpenApiVO`

NewUpdateChannelLimitConfigOpenApiVO instantiates a new UpdateChannelLimitConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateChannelLimitConfigOpenApiVOWithDefaults

`func NewUpdateChannelLimitConfigOpenApiVOWithDefaults() *UpdateChannelLimitConfigOpenApiVO`

NewUpdateChannelLimitConfigOpenApiVOWithDefaults instantiates a new UpdateChannelLimitConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLimitType

`func (o *UpdateChannelLimitConfigOpenApiVO) GetChannelLimitType() int32`

GetChannelLimitType returns the ChannelLimitType field if non-nil, zero value otherwise.

### GetChannelLimitTypeOk

`func (o *UpdateChannelLimitConfigOpenApiVO) GetChannelLimitTypeOk() (*int32, bool)`

GetChannelLimitTypeOk returns a tuple with the ChannelLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimitType

`func (o *UpdateChannelLimitConfigOpenApiVO) SetChannelLimitType(v int32)`

SetChannelLimitType sets ChannelLimitType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


