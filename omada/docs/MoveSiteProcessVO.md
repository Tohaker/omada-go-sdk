# MoveSiteProcessVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finished** | Pointer to **bool** | Process status of site moving | [optional] 
**MoveSiteId** | Pointer to **string** | ID of moved site | [optional] 
**MoveSiteResult** | Pointer to [**[]MoveSiteInfoVO**](MoveSiteInfoVO.md) | MoveSite results for each device. | [optional] 

## Methods

### NewMoveSiteProcessVO

`func NewMoveSiteProcessVO() *MoveSiteProcessVO`

NewMoveSiteProcessVO instantiates a new MoveSiteProcessVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveSiteProcessVOWithDefaults

`func NewMoveSiteProcessVOWithDefaults() *MoveSiteProcessVO`

NewMoveSiteProcessVOWithDefaults instantiates a new MoveSiteProcessVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinished

`func (o *MoveSiteProcessVO) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *MoveSiteProcessVO) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *MoveSiteProcessVO) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *MoveSiteProcessVO) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetMoveSiteId

`func (o *MoveSiteProcessVO) GetMoveSiteId() string`

GetMoveSiteId returns the MoveSiteId field if non-nil, zero value otherwise.

### GetMoveSiteIdOk

`func (o *MoveSiteProcessVO) GetMoveSiteIdOk() (*string, bool)`

GetMoveSiteIdOk returns a tuple with the MoveSiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteId

`func (o *MoveSiteProcessVO) SetMoveSiteId(v string)`

SetMoveSiteId sets MoveSiteId field to given value.

### HasMoveSiteId

`func (o *MoveSiteProcessVO) HasMoveSiteId() bool`

HasMoveSiteId returns a boolean if a field has been set.

### GetMoveSiteResult

`func (o *MoveSiteProcessVO) GetMoveSiteResult() []MoveSiteInfoVO`

GetMoveSiteResult returns the MoveSiteResult field if non-nil, zero value otherwise.

### GetMoveSiteResultOk

`func (o *MoveSiteProcessVO) GetMoveSiteResultOk() (*[]MoveSiteInfoVO, bool)`

GetMoveSiteResultOk returns a tuple with the MoveSiteResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveSiteResult

`func (o *MoveSiteProcessVO) SetMoveSiteResult(v []MoveSiteInfoVO)`

SetMoveSiteResult sets MoveSiteResult field to given value.

### HasMoveSiteResult

`func (o *MoveSiteProcessVO) HasMoveSiteResult() bool`

HasMoveSiteResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


