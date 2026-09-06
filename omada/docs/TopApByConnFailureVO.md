# TopApByConnFailureVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeedTip** | Pointer to **bool** | Not supported by the firmware on some devices | [optional] 
**TopAp2gFailList** | Pointer to [**[]ConnFailureApVO**](ConnFailureApVO.md) | Top AP by connection failure on 2G band | [optional] 
**TopAp5gFailList** | Pointer to [**[]ConnFailureApVO**](ConnFailureApVO.md) | Top AP by connection failure on 5G band | [optional] 
**TopAp6gFailList** | Pointer to [**[]ConnFailureApVO**](ConnFailureApVO.md) | Top AP by connection failure on 6G band | [optional] 

## Methods

### NewTopApByConnFailureVO

`func NewTopApByConnFailureVO() *TopApByConnFailureVO`

NewTopApByConnFailureVO instantiates a new TopApByConnFailureVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopApByConnFailureVOWithDefaults

`func NewTopApByConnFailureVOWithDefaults() *TopApByConnFailureVO`

NewTopApByConnFailureVOWithDefaults instantiates a new TopApByConnFailureVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeedTip

`func (o *TopApByConnFailureVO) GetNeedTip() bool`

GetNeedTip returns the NeedTip field if non-nil, zero value otherwise.

### GetNeedTipOk

`func (o *TopApByConnFailureVO) GetNeedTipOk() (*bool, bool)`

GetNeedTipOk returns a tuple with the NeedTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedTip

`func (o *TopApByConnFailureVO) SetNeedTip(v bool)`

SetNeedTip sets NeedTip field to given value.

### HasNeedTip

`func (o *TopApByConnFailureVO) HasNeedTip() bool`

HasNeedTip returns a boolean if a field has been set.

### GetTopAp2gFailList

`func (o *TopApByConnFailureVO) GetTopAp2gFailList() []ConnFailureApVO`

GetTopAp2gFailList returns the TopAp2gFailList field if non-nil, zero value otherwise.

### GetTopAp2gFailListOk

`func (o *TopApByConnFailureVO) GetTopAp2gFailListOk() (*[]ConnFailureApVO, bool)`

GetTopAp2gFailListOk returns a tuple with the TopAp2gFailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAp2gFailList

`func (o *TopApByConnFailureVO) SetTopAp2gFailList(v []ConnFailureApVO)`

SetTopAp2gFailList sets TopAp2gFailList field to given value.

### HasTopAp2gFailList

`func (o *TopApByConnFailureVO) HasTopAp2gFailList() bool`

HasTopAp2gFailList returns a boolean if a field has been set.

### GetTopAp5gFailList

`func (o *TopApByConnFailureVO) GetTopAp5gFailList() []ConnFailureApVO`

GetTopAp5gFailList returns the TopAp5gFailList field if non-nil, zero value otherwise.

### GetTopAp5gFailListOk

`func (o *TopApByConnFailureVO) GetTopAp5gFailListOk() (*[]ConnFailureApVO, bool)`

GetTopAp5gFailListOk returns a tuple with the TopAp5gFailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAp5gFailList

`func (o *TopApByConnFailureVO) SetTopAp5gFailList(v []ConnFailureApVO)`

SetTopAp5gFailList sets TopAp5gFailList field to given value.

### HasTopAp5gFailList

`func (o *TopApByConnFailureVO) HasTopAp5gFailList() bool`

HasTopAp5gFailList returns a boolean if a field has been set.

### GetTopAp6gFailList

`func (o *TopApByConnFailureVO) GetTopAp6gFailList() []ConnFailureApVO`

GetTopAp6gFailList returns the TopAp6gFailList field if non-nil, zero value otherwise.

### GetTopAp6gFailListOk

`func (o *TopApByConnFailureVO) GetTopAp6gFailListOk() (*[]ConnFailureApVO, bool)`

GetTopAp6gFailListOk returns a tuple with the TopAp6gFailList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopAp6gFailList

`func (o *TopApByConnFailureVO) SetTopAp6gFailList(v []ConnFailureApVO)`

SetTopAp6gFailList sets TopAp6gFailList field to given value.

### HasTopAp6gFailList

`func (o *TopApByConnFailureVO) HasTopAp6gFailList() bool`

HasTopAp6gFailList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


