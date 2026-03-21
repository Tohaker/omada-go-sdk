# OswStackSwitchVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdoptedSwitches** | Pointer to [**[]OswVO**](OswVO.md) | Adopted Switches | [optional] 
**DetectedSwitches** | Pointer to [**[]OswVO**](OswVO.md) | Detected Switches | [optional] 

## Methods

### NewOswStackSwitchVO

`func NewOswStackSwitchVO() *OswStackSwitchVO`

NewOswStackSwitchVO instantiates a new OswStackSwitchVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackSwitchVOWithDefaults

`func NewOswStackSwitchVOWithDefaults() *OswStackSwitchVO`

NewOswStackSwitchVOWithDefaults instantiates a new OswStackSwitchVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdoptedSwitches

`func (o *OswStackSwitchVO) GetAdoptedSwitches() []OswVO`

GetAdoptedSwitches returns the AdoptedSwitches field if non-nil, zero value otherwise.

### GetAdoptedSwitchesOk

`func (o *OswStackSwitchVO) GetAdoptedSwitchesOk() (*[]OswVO, bool)`

GetAdoptedSwitchesOk returns a tuple with the AdoptedSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptedSwitches

`func (o *OswStackSwitchVO) SetAdoptedSwitches(v []OswVO)`

SetAdoptedSwitches sets AdoptedSwitches field to given value.

### HasAdoptedSwitches

`func (o *OswStackSwitchVO) HasAdoptedSwitches() bool`

HasAdoptedSwitches returns a boolean if a field has been set.

### GetDetectedSwitches

`func (o *OswStackSwitchVO) GetDetectedSwitches() []OswVO`

GetDetectedSwitches returns the DetectedSwitches field if non-nil, zero value otherwise.

### GetDetectedSwitchesOk

`func (o *OswStackSwitchVO) GetDetectedSwitchesOk() (*[]OswVO, bool)`

GetDetectedSwitchesOk returns a tuple with the DetectedSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedSwitches

`func (o *OswStackSwitchVO) SetDetectedSwitches(v []OswVO)`

SetDetectedSwitches sets DetectedSwitches field to given value.

### HasDetectedSwitches

`func (o *OswStackSwitchVO) HasDetectedSwitches() bool`

HasDetectedSwitches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


