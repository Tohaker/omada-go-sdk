# WanLoadBalanceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppOptRouting** | **bool** | Application Optimized Routing | 
**BackupMode** | Pointer to **int32** | It is required when [linkBackup] is true. 0: The system will try to forward the traffic via the backup WAN port when primary WAN fails. Even if the primary WAN is recovered, it will not switch back unless the backup WAN fails; 1: Traffic is always forwarded through the primary WAN port unless it fails. The system will try to forward the traffic via the backup WAN port when it fails, and switch back when it recovers. | [optional] 
**BackupWan** | Pointer to **string** | Backup WAN ID. It is required when [linkBackup] is true. | [optional] 
**LinkBackup** | **bool** | Link Backup | 
**Method** | Pointer to **string** | effective only for &#39;linkBackup: true&#39; | [optional] 
**Mode** | Pointer to **int32** | It is required when [linkBackup] is true. 0: Enable backup link when any primary WAN fails. 1: Enable backup link when all primary WANs fail. 2: Timing | [optional] 
**PrimaryWans** | Pointer to **[]string** | Primary WAN port IDs. It is required when [linkBackup] is true. | [optional] 
**TimeRangeId** | Pointer to **string** | Time Range ID. It is required when [mode] is timing. | [optional] 
**VirtualWanWeights** | Pointer to [**[]VirtualWanWeightOpenApiVO**](VirtualWanWeightOpenApiVO.md) | virtual wan load balance | [optional] 
**Weights** | **[]int32** | Load Balancing Weights,item of weights should be within the range of 1 to the max int value(2147483647). It is sorted by port ID. | 

## Methods

### NewWanLoadBalanceOpenApiVO

`func NewWanLoadBalanceOpenApiVO(appOptRouting bool, linkBackup bool, weights []int32, ) *WanLoadBalanceOpenApiVO`

NewWanLoadBalanceOpenApiVO instantiates a new WanLoadBalanceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanLoadBalanceOpenApiVOWithDefaults

`func NewWanLoadBalanceOpenApiVOWithDefaults() *WanLoadBalanceOpenApiVO`

NewWanLoadBalanceOpenApiVOWithDefaults instantiates a new WanLoadBalanceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppOptRouting

`func (o *WanLoadBalanceOpenApiVO) GetAppOptRouting() bool`

GetAppOptRouting returns the AppOptRouting field if non-nil, zero value otherwise.

### GetAppOptRoutingOk

`func (o *WanLoadBalanceOpenApiVO) GetAppOptRoutingOk() (*bool, bool)`

GetAppOptRoutingOk returns a tuple with the AppOptRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppOptRouting

`func (o *WanLoadBalanceOpenApiVO) SetAppOptRouting(v bool)`

SetAppOptRouting sets AppOptRouting field to given value.


### GetBackupMode

`func (o *WanLoadBalanceOpenApiVO) GetBackupMode() int32`

GetBackupMode returns the BackupMode field if non-nil, zero value otherwise.

### GetBackupModeOk

`func (o *WanLoadBalanceOpenApiVO) GetBackupModeOk() (*int32, bool)`

GetBackupModeOk returns a tuple with the BackupMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMode

`func (o *WanLoadBalanceOpenApiVO) SetBackupMode(v int32)`

SetBackupMode sets BackupMode field to given value.

### HasBackupMode

`func (o *WanLoadBalanceOpenApiVO) HasBackupMode() bool`

HasBackupMode returns a boolean if a field has been set.

### GetBackupWan

`func (o *WanLoadBalanceOpenApiVO) GetBackupWan() string`

GetBackupWan returns the BackupWan field if non-nil, zero value otherwise.

### GetBackupWanOk

`func (o *WanLoadBalanceOpenApiVO) GetBackupWanOk() (*string, bool)`

GetBackupWanOk returns a tuple with the BackupWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupWan

`func (o *WanLoadBalanceOpenApiVO) SetBackupWan(v string)`

SetBackupWan sets BackupWan field to given value.

### HasBackupWan

`func (o *WanLoadBalanceOpenApiVO) HasBackupWan() bool`

HasBackupWan returns a boolean if a field has been set.

### GetLinkBackup

`func (o *WanLoadBalanceOpenApiVO) GetLinkBackup() bool`

GetLinkBackup returns the LinkBackup field if non-nil, zero value otherwise.

### GetLinkBackupOk

`func (o *WanLoadBalanceOpenApiVO) GetLinkBackupOk() (*bool, bool)`

GetLinkBackupOk returns a tuple with the LinkBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkBackup

`func (o *WanLoadBalanceOpenApiVO) SetLinkBackup(v bool)`

SetLinkBackup sets LinkBackup field to given value.


### GetMethod

`func (o *WanLoadBalanceOpenApiVO) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WanLoadBalanceOpenApiVO) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WanLoadBalanceOpenApiVO) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *WanLoadBalanceOpenApiVO) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMode

`func (o *WanLoadBalanceOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *WanLoadBalanceOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *WanLoadBalanceOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *WanLoadBalanceOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPrimaryWans

`func (o *WanLoadBalanceOpenApiVO) GetPrimaryWans() []string`

GetPrimaryWans returns the PrimaryWans field if non-nil, zero value otherwise.

### GetPrimaryWansOk

`func (o *WanLoadBalanceOpenApiVO) GetPrimaryWansOk() (*[]string, bool)`

GetPrimaryWansOk returns a tuple with the PrimaryWans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryWans

`func (o *WanLoadBalanceOpenApiVO) SetPrimaryWans(v []string)`

SetPrimaryWans sets PrimaryWans field to given value.

### HasPrimaryWans

`func (o *WanLoadBalanceOpenApiVO) HasPrimaryWans() bool`

HasPrimaryWans returns a boolean if a field has been set.

### GetTimeRangeId

`func (o *WanLoadBalanceOpenApiVO) GetTimeRangeId() string`

GetTimeRangeId returns the TimeRangeId field if non-nil, zero value otherwise.

### GetTimeRangeIdOk

`func (o *WanLoadBalanceOpenApiVO) GetTimeRangeIdOk() (*string, bool)`

GetTimeRangeIdOk returns a tuple with the TimeRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeId

`func (o *WanLoadBalanceOpenApiVO) SetTimeRangeId(v string)`

SetTimeRangeId sets TimeRangeId field to given value.

### HasTimeRangeId

`func (o *WanLoadBalanceOpenApiVO) HasTimeRangeId() bool`

HasTimeRangeId returns a boolean if a field has been set.

### GetVirtualWanWeights

`func (o *WanLoadBalanceOpenApiVO) GetVirtualWanWeights() []VirtualWanWeightOpenApiVO`

GetVirtualWanWeights returns the VirtualWanWeights field if non-nil, zero value otherwise.

### GetVirtualWanWeightsOk

`func (o *WanLoadBalanceOpenApiVO) GetVirtualWanWeightsOk() (*[]VirtualWanWeightOpenApiVO, bool)`

GetVirtualWanWeightsOk returns a tuple with the VirtualWanWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanWeights

`func (o *WanLoadBalanceOpenApiVO) SetVirtualWanWeights(v []VirtualWanWeightOpenApiVO)`

SetVirtualWanWeights sets VirtualWanWeights field to given value.

### HasVirtualWanWeights

`func (o *WanLoadBalanceOpenApiVO) HasVirtualWanWeights() bool`

HasVirtualWanWeights returns a boolean if a field has been set.

### GetWeights

`func (o *WanLoadBalanceOpenApiVO) GetWeights() []int32`

GetWeights returns the Weights field if non-nil, zero value otherwise.

### GetWeightsOk

`func (o *WanLoadBalanceOpenApiVO) GetWeightsOk() (*[]int32, bool)`

GetWeightsOk returns a tuple with the Weights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeights

`func (o *WanLoadBalanceOpenApiVO) SetWeights(v []int32)`

SetWeights sets Weights field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


