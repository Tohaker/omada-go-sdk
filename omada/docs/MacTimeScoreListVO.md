# MacTimeScoreListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac address | [optional] 
**Scores** | Pointer to [**[]TimeScoreItemVO**](TimeScoreItemVO.md) | Health score with time | [optional] 

## Methods

### NewMacTimeScoreListVO

`func NewMacTimeScoreListVO() *MacTimeScoreListVO`

NewMacTimeScoreListVO instantiates a new MacTimeScoreListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacTimeScoreListVOWithDefaults

`func NewMacTimeScoreListVOWithDefaults() *MacTimeScoreListVO`

NewMacTimeScoreListVOWithDefaults instantiates a new MacTimeScoreListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *MacTimeScoreListVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MacTimeScoreListVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MacTimeScoreListVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MacTimeScoreListVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetScores

`func (o *MacTimeScoreListVO) GetScores() []TimeScoreItemVO`

GetScores returns the Scores field if non-nil, zero value otherwise.

### GetScoresOk

`func (o *MacTimeScoreListVO) GetScoresOk() (*[]TimeScoreItemVO, bool)`

GetScoresOk returns a tuple with the Scores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScores

`func (o *MacTimeScoreListVO) SetScores(v []TimeScoreItemVO)`

SetScores sets Scores field to given value.

### HasScores

`func (o *MacTimeScoreListVO) HasScores() bool`

HasScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


