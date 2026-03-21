# ChannelLimitConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimitType** | Pointer to **int32** | Channel limit enable status. It should be a value as follows: 0: default, 1: false, 2: true. | [optional] 
**SupportChannelLimit** | Pointer to **bool** | Indicates whether the device supports channel limit | [optional] 

## Methods

### NewChannelLimitConfigOpenApiVO

`func NewChannelLimitConfigOpenApiVO() *ChannelLimitConfigOpenApiVO`

NewChannelLimitConfigOpenApiVO instantiates a new ChannelLimitConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelLimitConfigOpenApiVOWithDefaults

`func NewChannelLimitConfigOpenApiVOWithDefaults() *ChannelLimitConfigOpenApiVO`

NewChannelLimitConfigOpenApiVOWithDefaults instantiates a new ChannelLimitConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLimitType

`func (o *ChannelLimitConfigOpenApiVO) GetChannelLimitType() int32`

GetChannelLimitType returns the ChannelLimitType field if non-nil, zero value otherwise.

### GetChannelLimitTypeOk

`func (o *ChannelLimitConfigOpenApiVO) GetChannelLimitTypeOk() (*int32, bool)`

GetChannelLimitTypeOk returns a tuple with the ChannelLimitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLimitType

`func (o *ChannelLimitConfigOpenApiVO) SetChannelLimitType(v int32)`

SetChannelLimitType sets ChannelLimitType field to given value.

### HasChannelLimitType

`func (o *ChannelLimitConfigOpenApiVO) HasChannelLimitType() bool`

HasChannelLimitType returns a boolean if a field has been set.

### GetSupportChannelLimit

`func (o *ChannelLimitConfigOpenApiVO) GetSupportChannelLimit() bool`

GetSupportChannelLimit returns the SupportChannelLimit field if non-nil, zero value otherwise.

### GetSupportChannelLimitOk

`func (o *ChannelLimitConfigOpenApiVO) GetSupportChannelLimitOk() (*bool, bool)`

GetSupportChannelLimitOk returns a tuple with the SupportChannelLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportChannelLimit

`func (o *ChannelLimitConfigOpenApiVO) SetSupportChannelLimit(v bool)`

SetSupportChannelLimit sets SupportChannelLimit field to given value.

### HasSupportChannelLimit

`func (o *ChannelLimitConfigOpenApiVO) HasSupportChannelLimit() bool`

HasSupportChannelLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


