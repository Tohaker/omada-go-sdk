# ApRadioAnteGainConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnteGain** | Pointer to **int32** | Antenna Gain. The value of parameter [anteGain] is between 0 and 30 | [optional] 
**RadioId** | Pointer to **int32** | The value of parameter [radioId] is between 0 and 3. 0： 2.4 GHz, 1: 5 GHz, 2: 5 GHz-2, 3: 6 GHz. | [optional] 

## Methods

### NewApRadioAnteGainConfigOpenApiVO

`func NewApRadioAnteGainConfigOpenApiVO() *ApRadioAnteGainConfigOpenApiVO`

NewApRadioAnteGainConfigOpenApiVO instantiates a new ApRadioAnteGainConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRadioAnteGainConfigOpenApiVOWithDefaults

`func NewApRadioAnteGainConfigOpenApiVOWithDefaults() *ApRadioAnteGainConfigOpenApiVO`

NewApRadioAnteGainConfigOpenApiVOWithDefaults instantiates a new ApRadioAnteGainConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnteGain

`func (o *ApRadioAnteGainConfigOpenApiVO) GetAnteGain() int32`

GetAnteGain returns the AnteGain field if non-nil, zero value otherwise.

### GetAnteGainOk

`func (o *ApRadioAnteGainConfigOpenApiVO) GetAnteGainOk() (*int32, bool)`

GetAnteGainOk returns a tuple with the AnteGain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnteGain

`func (o *ApRadioAnteGainConfigOpenApiVO) SetAnteGain(v int32)`

SetAnteGain sets AnteGain field to given value.

### HasAnteGain

`func (o *ApRadioAnteGainConfigOpenApiVO) HasAnteGain() bool`

HasAnteGain returns a boolean if a field has been set.

### GetRadioId

`func (o *ApRadioAnteGainConfigOpenApiVO) GetRadioId() int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *ApRadioAnteGainConfigOpenApiVO) GetRadioIdOk() (*int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *ApRadioAnteGainConfigOpenApiVO) SetRadioId(v int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *ApRadioAnteGainConfigOpenApiVO) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


