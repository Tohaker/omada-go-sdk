# ApChannelDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableChannelWidthList** | Pointer to **[]int32** | Available bandwidth list for the channel configuration; For example: [20, 40, 80, 160, 240, 320]. | [optional] 
**Channel** | Pointer to **int32** | channel value; For example, 1 in 1/2412MHz. | [optional] 
**Freq** | Pointer to **int32** | channel frequency; For example, 2412 in 1/2412MHz. | [optional] 
**Index** | Pointer to **int32** | channel index; For example, if channel 36 is the first channel of the device&#39;s 5 GHz frequency band, its index is 1. | [optional] 

## Methods

### NewApChannelDetailOpenApiVO

`func NewApChannelDetailOpenApiVO() *ApChannelDetailOpenApiVO`

NewApChannelDetailOpenApiVO instantiates a new ApChannelDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApChannelDetailOpenApiVOWithDefaults

`func NewApChannelDetailOpenApiVOWithDefaults() *ApChannelDetailOpenApiVO`

NewApChannelDetailOpenApiVOWithDefaults instantiates a new ApChannelDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableChannelWidthList

`func (o *ApChannelDetailOpenApiVO) GetAvailableChannelWidthList() []int32`

GetAvailableChannelWidthList returns the AvailableChannelWidthList field if non-nil, zero value otherwise.

### GetAvailableChannelWidthListOk

`func (o *ApChannelDetailOpenApiVO) GetAvailableChannelWidthListOk() (*[]int32, bool)`

GetAvailableChannelWidthListOk returns a tuple with the AvailableChannelWidthList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableChannelWidthList

`func (o *ApChannelDetailOpenApiVO) SetAvailableChannelWidthList(v []int32)`

SetAvailableChannelWidthList sets AvailableChannelWidthList field to given value.

### HasAvailableChannelWidthList

`func (o *ApChannelDetailOpenApiVO) HasAvailableChannelWidthList() bool`

HasAvailableChannelWidthList returns a boolean if a field has been set.

### GetChannel

`func (o *ApChannelDetailOpenApiVO) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *ApChannelDetailOpenApiVO) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *ApChannelDetailOpenApiVO) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *ApChannelDetailOpenApiVO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetFreq

`func (o *ApChannelDetailOpenApiVO) GetFreq() int32`

GetFreq returns the Freq field if non-nil, zero value otherwise.

### GetFreqOk

`func (o *ApChannelDetailOpenApiVO) GetFreqOk() (*int32, bool)`

GetFreqOk returns a tuple with the Freq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreq

`func (o *ApChannelDetailOpenApiVO) SetFreq(v int32)`

SetFreq sets Freq field to given value.

### HasFreq

`func (o *ApChannelDetailOpenApiVO) HasFreq() bool`

HasFreq returns a boolean if a field has been set.

### GetIndex

`func (o *ApChannelDetailOpenApiVO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ApChannelDetailOpenApiVO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ApChannelDetailOpenApiVO) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ApChannelDetailOpenApiVO) HasIndex() bool`

HasIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


