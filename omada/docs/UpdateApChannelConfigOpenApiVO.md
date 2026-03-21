# UpdateApChannelConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Radios-Channel config; The channel list supported by device can be obtained from interface : Get available channel list of ap; If select auto configuration need to enter 0. | [optional] 
**ChannelWidth** | Pointer to **int32** | Radios-Channel width config; The channelWidth list supported by device can be obtained from interface : Get available channel list of ap; If select auto configuration need to enter 0. | [optional] 
**RadioEnable** | **bool** | Enable/Disable radio setting(if false, other params is not required) | 
**RadioId** | **int32** | (Wireless) Radio ID should be a value as follows: 0: 2.4GHz; 1: 5GHz(5GHz-1); 2: 5GHz-2; 3: 6GHz. | 

## Methods

### NewUpdateApChannelConfigOpenApiVO

`func NewUpdateApChannelConfigOpenApiVO(radioEnable bool, radioId int32, ) *UpdateApChannelConfigOpenApiVO`

NewUpdateApChannelConfigOpenApiVO instantiates a new UpdateApChannelConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApChannelConfigOpenApiVOWithDefaults

`func NewUpdateApChannelConfigOpenApiVOWithDefaults() *UpdateApChannelConfigOpenApiVO`

NewUpdateApChannelConfigOpenApiVOWithDefaults instantiates a new UpdateApChannelConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *UpdateApChannelConfigOpenApiVO) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *UpdateApChannelConfigOpenApiVO) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *UpdateApChannelConfigOpenApiVO) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *UpdateApChannelConfigOpenApiVO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetChannelWidth

`func (o *UpdateApChannelConfigOpenApiVO) GetChannelWidth() int32`

GetChannelWidth returns the ChannelWidth field if non-nil, zero value otherwise.

### GetChannelWidthOk

`func (o *UpdateApChannelConfigOpenApiVO) GetChannelWidthOk() (*int32, bool)`

GetChannelWidthOk returns a tuple with the ChannelWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidth

`func (o *UpdateApChannelConfigOpenApiVO) SetChannelWidth(v int32)`

SetChannelWidth sets ChannelWidth field to given value.

### HasChannelWidth

`func (o *UpdateApChannelConfigOpenApiVO) HasChannelWidth() bool`

HasChannelWidth returns a boolean if a field has been set.

### GetRadioEnable

`func (o *UpdateApChannelConfigOpenApiVO) GetRadioEnable() bool`

GetRadioEnable returns the RadioEnable field if non-nil, zero value otherwise.

### GetRadioEnableOk

`func (o *UpdateApChannelConfigOpenApiVO) GetRadioEnableOk() (*bool, bool)`

GetRadioEnableOk returns a tuple with the RadioEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioEnable

`func (o *UpdateApChannelConfigOpenApiVO) SetRadioEnable(v bool)`

SetRadioEnable sets RadioEnable field to given value.


### GetRadioId

`func (o *UpdateApChannelConfigOpenApiVO) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *UpdateApChannelConfigOpenApiVO) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *UpdateApChannelConfigOpenApiVO) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


