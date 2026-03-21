# UpgradeScheduleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMacs** | **[]string** | MAC address of the selected devices. E.g. AA-BB-CC-DD-11-22 | 
**Name** | **string** | Upgrade schedule name should contain 1 to 128 characters. | 
**NextExecuteTime** | Pointer to **int64** | Execution timeStamp(ms). Required when type is 0. | [optional] 
**OccurrenceTime** | Pointer to [**BaseScheduleTimeOpenApiVO**](BaseScheduleTimeOpenApiVO.md) |  | [optional] 
**Status** | **bool** | Upgrade schedule status. | 
**Type** | **int32** | Type should be a value as follows: 0: execute only once; 1: repeat | 

## Methods

### NewUpgradeScheduleOpenApiVO

`func NewUpgradeScheduleOpenApiVO(deviceMacs []string, name string, status bool, type_ int32, ) *UpgradeScheduleOpenApiVO`

NewUpgradeScheduleOpenApiVO instantiates a new UpgradeScheduleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeScheduleOpenApiVOWithDefaults

`func NewUpgradeScheduleOpenApiVOWithDefaults() *UpgradeScheduleOpenApiVO`

NewUpgradeScheduleOpenApiVOWithDefaults instantiates a new UpgradeScheduleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMacs

`func (o *UpgradeScheduleOpenApiVO) GetDeviceMacs() []string`

GetDeviceMacs returns the DeviceMacs field if non-nil, zero value otherwise.

### GetDeviceMacsOk

`func (o *UpgradeScheduleOpenApiVO) GetDeviceMacsOk() (*[]string, bool)`

GetDeviceMacsOk returns a tuple with the DeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacs

`func (o *UpgradeScheduleOpenApiVO) SetDeviceMacs(v []string)`

SetDeviceMacs sets DeviceMacs field to given value.


### GetName

`func (o *UpgradeScheduleOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpgradeScheduleOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpgradeScheduleOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNextExecuteTime

`func (o *UpgradeScheduleOpenApiVO) GetNextExecuteTime() int64`

GetNextExecuteTime returns the NextExecuteTime field if non-nil, zero value otherwise.

### GetNextExecuteTimeOk

`func (o *UpgradeScheduleOpenApiVO) GetNextExecuteTimeOk() (*int64, bool)`

GetNextExecuteTimeOk returns a tuple with the NextExecuteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecuteTime

`func (o *UpgradeScheduleOpenApiVO) SetNextExecuteTime(v int64)`

SetNextExecuteTime sets NextExecuteTime field to given value.

### HasNextExecuteTime

`func (o *UpgradeScheduleOpenApiVO) HasNextExecuteTime() bool`

HasNextExecuteTime returns a boolean if a field has been set.

### GetOccurrenceTime

`func (o *UpgradeScheduleOpenApiVO) GetOccurrenceTime() BaseScheduleTimeOpenApiVO`

GetOccurrenceTime returns the OccurrenceTime field if non-nil, zero value otherwise.

### GetOccurrenceTimeOk

`func (o *UpgradeScheduleOpenApiVO) GetOccurrenceTimeOk() (*BaseScheduleTimeOpenApiVO, bool)`

GetOccurrenceTimeOk returns a tuple with the OccurrenceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceTime

`func (o *UpgradeScheduleOpenApiVO) SetOccurrenceTime(v BaseScheduleTimeOpenApiVO)`

SetOccurrenceTime sets OccurrenceTime field to given value.

### HasOccurrenceTime

`func (o *UpgradeScheduleOpenApiVO) HasOccurrenceTime() bool`

HasOccurrenceTime returns a boolean if a field has been set.

### GetStatus

`func (o *UpgradeScheduleOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpgradeScheduleOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpgradeScheduleOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetType

`func (o *UpgradeScheduleOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpgradeScheduleOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpgradeScheduleOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


