# UpdateApRadioAnteGainConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnteGain** | **int32** | Antenna Gain. The value of [anteGain] should be between 0 and 30.  | 
**RadioId** | **int32** | The value of parameter [radioId] should be between 0 and 3. 0: 2.4 GHz, 1: 5 GHz, 2: 5 GHz-2, 3: 6 GHz. | 

## Methods

### NewUpdateApRadioAnteGainConfigOpenApiVO

`func NewUpdateApRadioAnteGainConfigOpenApiVO(anteGain int32, radioId int32, ) *UpdateApRadioAnteGainConfigOpenApiVO`

NewUpdateApRadioAnteGainConfigOpenApiVO instantiates a new UpdateApRadioAnteGainConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateApRadioAnteGainConfigOpenApiVOWithDefaults

`func NewUpdateApRadioAnteGainConfigOpenApiVOWithDefaults() *UpdateApRadioAnteGainConfigOpenApiVO`

NewUpdateApRadioAnteGainConfigOpenApiVOWithDefaults instantiates a new UpdateApRadioAnteGainConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnteGain

`func (o *UpdateApRadioAnteGainConfigOpenApiVO) GetAnteGain() int32`

GetAnteGain returns the AnteGain field if non-nil, zero value otherwise.

### GetAnteGainOk

`func (o *UpdateApRadioAnteGainConfigOpenApiVO) GetAnteGainOk() (*int32, bool)`

GetAnteGainOk returns a tuple with the AnteGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnteGain

`func (o *UpdateApRadioAnteGainConfigOpenApiVO) SetAnteGain(v int32)`

SetAnteGain sets AnteGain field to given value.


### GetRadioId

`func (o *UpdateApRadioAnteGainConfigOpenApiVO) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *UpdateApRadioAnteGainConfigOpenApiVO) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *UpdateApRadioAnteGainConfigOpenApiVO) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


