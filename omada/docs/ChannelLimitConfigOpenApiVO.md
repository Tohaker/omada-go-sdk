# ChannelLimitConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLimitType** | Pointer to **int32** | Channel limit enable status. It should be a value as follows: 0: default, 1: false, 2: true. | [optional] 
**DefaultInstType5g** | Pointer to **bool** | default mode in 5g radio. true: outdoor; false: indoor | [optional] 
**DefaultInstType6g** | Pointer to **bool** | default mode in 6g radio. true: outdoor; false: indoor | [optional] 
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

### GetDefaultInstType5g

`func (o *ChannelLimitConfigOpenApiVO) GetDefaultInstType5g() bool`

GetDefaultInstType5g returns the DefaultInstType5g field if non-nil, zero value otherwise.

### GetDefaultInstType5gOk

`func (o *ChannelLimitConfigOpenApiVO) GetDefaultInstType5gOk() (*bool, bool)`

GetDefaultInstType5gOk returns a tuple with the DefaultInstType5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstType5g

`func (o *ChannelLimitConfigOpenApiVO) SetDefaultInstType5g(v bool)`

SetDefaultInstType5g sets DefaultInstType5g field to given value.

### HasDefaultInstType5g

`func (o *ChannelLimitConfigOpenApiVO) HasDefaultInstType5g() bool`

HasDefaultInstType5g returns a boolean if a field has been set.

### GetDefaultInstType6g

`func (o *ChannelLimitConfigOpenApiVO) GetDefaultInstType6g() bool`

GetDefaultInstType6g returns the DefaultInstType6g field if non-nil, zero value otherwise.

### GetDefaultInstType6gOk

`func (o *ChannelLimitConfigOpenApiVO) GetDefaultInstType6gOk() (*bool, bool)`

GetDefaultInstType6gOk returns a tuple with the DefaultInstType6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultInstType6g

`func (o *ChannelLimitConfigOpenApiVO) SetDefaultInstType6g(v bool)`

SetDefaultInstType6g sets DefaultInstType6g field to given value.

### HasDefaultInstType6g

`func (o *ChannelLimitConfigOpenApiVO) HasDefaultInstType6g() bool`

HasDefaultInstType6g returns a boolean if a field has been set.

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


