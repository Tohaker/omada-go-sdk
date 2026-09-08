# TopSwitchByPacketVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NeedTip** | Pointer to **bool** | Not supported by the firmware on some devices | [optional] 
**TopPacketError** | Pointer to [**[]SwitchPacketErrorVO**](SwitchPacketErrorVO.md) | Top switches by packet error | [optional] 
**TopPacketLoss** | Pointer to [**[]SwitchPacketLossVO**](SwitchPacketLossVO.md) | Top switches by packet loss | [optional] 

## Methods

### NewTopSwitchByPacketVO

`func NewTopSwitchByPacketVO() *TopSwitchByPacketVO`

NewTopSwitchByPacketVO instantiates a new TopSwitchByPacketVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopSwitchByPacketVOWithDefaults

`func NewTopSwitchByPacketVOWithDefaults() *TopSwitchByPacketVO`

NewTopSwitchByPacketVOWithDefaults instantiates a new TopSwitchByPacketVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNeedTip

`func (o *TopSwitchByPacketVO) GetNeedTip() bool`

GetNeedTip returns the NeedTip field if non-nil, zero value otherwise.

### GetNeedTipOk

`func (o *TopSwitchByPacketVO) GetNeedTipOk() (*bool, bool)`

GetNeedTipOk returns a tuple with the NeedTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedTip

`func (o *TopSwitchByPacketVO) SetNeedTip(v bool)`

SetNeedTip sets NeedTip field to given value.

### HasNeedTip

`func (o *TopSwitchByPacketVO) HasNeedTip() bool`

HasNeedTip returns a boolean if a field has been set.

### GetTopPacketError

`func (o *TopSwitchByPacketVO) GetTopPacketError() []SwitchPacketErrorVO`

GetTopPacketError returns the TopPacketError field if non-nil, zero value otherwise.

### GetTopPacketErrorOk

`func (o *TopSwitchByPacketVO) GetTopPacketErrorOk() (*[]SwitchPacketErrorVO, bool)`

GetTopPacketErrorOk returns a tuple with the TopPacketError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPacketError

`func (o *TopSwitchByPacketVO) SetTopPacketError(v []SwitchPacketErrorVO)`

SetTopPacketError sets TopPacketError field to given value.

### HasTopPacketError

`func (o *TopSwitchByPacketVO) HasTopPacketError() bool`

HasTopPacketError returns a boolean if a field has been set.

### GetTopPacketLoss

`func (o *TopSwitchByPacketVO) GetTopPacketLoss() []SwitchPacketLossVO`

GetTopPacketLoss returns the TopPacketLoss field if non-nil, zero value otherwise.

### GetTopPacketLossOk

`func (o *TopSwitchByPacketVO) GetTopPacketLossOk() (*[]SwitchPacketLossVO, bool)`

GetTopPacketLossOk returns a tuple with the TopPacketLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPacketLoss

`func (o *TopSwitchByPacketVO) SetTopPacketLoss(v []SwitchPacketLossVO)`

SetTopPacketLoss sets TopPacketLoss field to given value.

### HasTopPacketLoss

`func (o *TopSwitchByPacketVO) HasTopPacketLoss() bool`

HasTopPacketLoss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


