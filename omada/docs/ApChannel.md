# ApChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Channel. | [optional] 
**Load** | Pointer to **int32** | The channel load corresponding to the channel. | [optional] 

## Methods

### NewApChannel

`func NewApChannel() *ApChannel`

NewApChannel instantiates a new ApChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApChannelWithDefaults

`func NewApChannelWithDefaults() *ApChannel`

NewApChannelWithDefaults instantiates a new ApChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *ApChannel) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApChannel) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApChannel) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApChannel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetLoad

`func (o *ApChannel) GetLoad() int32`

GetLoad returns the Load field if non-nil, zero value otherwise.

### GetLoadOk

`func (o *ApChannel) GetLoadOk() (*int32, bool)`

GetLoadOk returns a tuple with the Load field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoad

`func (o *ApChannel) SetLoad(v int32)`

SetLoad sets Load field to given value.

### HasLoad

`func (o *ApChannel) HasLoad() bool`

HasLoad returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


