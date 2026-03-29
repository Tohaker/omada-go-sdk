# OswRankingCardsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TopActiveSwitches** | Pointer to [**[]ActiveOswHealthOpenApiVO**](ActiveOswHealthOpenApiVO.md) | Active switches information list | [optional] 
**TopCpu** | Pointer to [**[]TopCpuUsageHealthOpenApiVO**](TopCpuUsageHealthOpenApiVO.md) | Cpu information list | [optional] 
**TopMem** | Pointer to [**[]TopMemUsageHealthOpenApiVO**](TopMemUsageHealthOpenApiVO.md) | Memory information list | [optional] 
**TopPacketsError** | Pointer to [**[]TopOswErrorPacketOpenApiVO**](TopOswErrorPacketOpenApiVO.md) | Packet error information list | [optional] 
**TopPacketsLoss** | Pointer to [**[]TopOswPacketLossOpenApiVO**](TopOswPacketLossOpenApiVO.md) | Packet loss information list | [optional] 

## Methods

### NewOswRankingCardsOpenApiVO

`func NewOswRankingCardsOpenApiVO() *OswRankingCardsOpenApiVO`

NewOswRankingCardsOpenApiVO instantiates a new OswRankingCardsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswRankingCardsOpenApiVOWithDefaults

`func NewOswRankingCardsOpenApiVOWithDefaults() *OswRankingCardsOpenApiVO`

NewOswRankingCardsOpenApiVOWithDefaults instantiates a new OswRankingCardsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopActiveSwitches

`func (o *OswRankingCardsOpenApiVO) GetTopActiveSwitches() []ActiveOswHealthOpenApiVO`

GetTopActiveSwitches returns the TopActiveSwitches field if non-nil, zero value otherwise.

### GetTopActiveSwitchesOk

`func (o *OswRankingCardsOpenApiVO) GetTopActiveSwitchesOk() (*[]ActiveOswHealthOpenApiVO, bool)`

GetTopActiveSwitchesOk returns a tuple with the TopActiveSwitches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopActiveSwitches

`func (o *OswRankingCardsOpenApiVO) SetTopActiveSwitches(v []ActiveOswHealthOpenApiVO)`

SetTopActiveSwitches sets TopActiveSwitches field to given value.

### HasTopActiveSwitches

`func (o *OswRankingCardsOpenApiVO) HasTopActiveSwitches() bool`

HasTopActiveSwitches returns a boolean if a field has been set.

### GetTopCpu

`func (o *OswRankingCardsOpenApiVO) GetTopCpu() []TopCpuUsageHealthOpenApiVO`

GetTopCpu returns the TopCpu field if non-nil, zero value otherwise.

### GetTopCpuOk

`func (o *OswRankingCardsOpenApiVO) GetTopCpuOk() (*[]TopCpuUsageHealthOpenApiVO, bool)`

GetTopCpuOk returns a tuple with the TopCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopCpu

`func (o *OswRankingCardsOpenApiVO) SetTopCpu(v []TopCpuUsageHealthOpenApiVO)`

SetTopCpu sets TopCpu field to given value.

### HasTopCpu

`func (o *OswRankingCardsOpenApiVO) HasTopCpu() bool`

HasTopCpu returns a boolean if a field has been set.

### GetTopMem

`func (o *OswRankingCardsOpenApiVO) GetTopMem() []TopMemUsageHealthOpenApiVO`

GetTopMem returns the TopMem field if non-nil, zero value otherwise.

### GetTopMemOk

`func (o *OswRankingCardsOpenApiVO) GetTopMemOk() (*[]TopMemUsageHealthOpenApiVO, bool)`

GetTopMemOk returns a tuple with the TopMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopMem

`func (o *OswRankingCardsOpenApiVO) SetTopMem(v []TopMemUsageHealthOpenApiVO)`

SetTopMem sets TopMem field to given value.

### HasTopMem

`func (o *OswRankingCardsOpenApiVO) HasTopMem() bool`

HasTopMem returns a boolean if a field has been set.

### GetTopPacketsError

`func (o *OswRankingCardsOpenApiVO) GetTopPacketsError() []TopOswErrorPacketOpenApiVO`

GetTopPacketsError returns the TopPacketsError field if non-nil, zero value otherwise.

### GetTopPacketsErrorOk

`func (o *OswRankingCardsOpenApiVO) GetTopPacketsErrorOk() (*[]TopOswErrorPacketOpenApiVO, bool)`

GetTopPacketsErrorOk returns a tuple with the TopPacketsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPacketsError

`func (o *OswRankingCardsOpenApiVO) SetTopPacketsError(v []TopOswErrorPacketOpenApiVO)`

SetTopPacketsError sets TopPacketsError field to given value.

### HasTopPacketsError

`func (o *OswRankingCardsOpenApiVO) HasTopPacketsError() bool`

HasTopPacketsError returns a boolean if a field has been set.

### GetTopPacketsLoss

`func (o *OswRankingCardsOpenApiVO) GetTopPacketsLoss() []TopOswPacketLossOpenApiVO`

GetTopPacketsLoss returns the TopPacketsLoss field if non-nil, zero value otherwise.

### GetTopPacketsLossOk

`func (o *OswRankingCardsOpenApiVO) GetTopPacketsLossOk() (*[]TopOswPacketLossOpenApiVO, bool)`

GetTopPacketsLossOk returns a tuple with the TopPacketsLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPacketsLoss

`func (o *OswRankingCardsOpenApiVO) SetTopPacketsLoss(v []TopOswPacketLossOpenApiVO)`

SetTopPacketsLoss sets TopPacketsLoss field to given value.

### HasTopPacketsLoss

`func (o *OswRankingCardsOpenApiVO) HasTopPacketsLoss() bool`

HasTopPacketsLoss returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


