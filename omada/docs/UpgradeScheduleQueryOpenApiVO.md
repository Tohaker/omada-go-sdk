# UpgradeScheduleQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMacs** | **[]string** | MAC address of the selected devices. | 
**Id** | Pointer to **string** | Reboot Schedule ID. | [optional] 
**Name** | **string** | Reboot Schedule name should contain 1 to 128 characters. | 
**NextExecuteTime** | Pointer to **int64** | Execution timeStamp(ms). Required when type is 0. | [optional] 
**OccurrenceTime** | Pointer to [**BaseScheduleTimeOpenApiVO**](BaseScheduleTimeOpenApiVO.md) |  | [optional] 
**Status** | **bool** | Reboot Schedule status. | 
**Type** | **int32** | Type should be a value as follows: 0: execute only once; 1: repeat | 

## Methods

### NewUpgradeScheduleQueryOpenApiVO

`func NewUpgradeScheduleQueryOpenApiVO(deviceMacs []string, name string, status bool, type_ int32, ) *UpgradeScheduleQueryOpenApiVO`

NewUpgradeScheduleQueryOpenApiVO instantiates a new UpgradeScheduleQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeScheduleQueryOpenApiVOWithDefaults

`func NewUpgradeScheduleQueryOpenApiVOWithDefaults() *UpgradeScheduleQueryOpenApiVO`

NewUpgradeScheduleQueryOpenApiVOWithDefaults instantiates a new UpgradeScheduleQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMacs

`func (o *UpgradeScheduleQueryOpenApiVO) GetDeviceMacs() []string`

GetDeviceMacs returns the DeviceMacs field if non-nil, zero value otherwise.

### GetDeviceMacsOk

`func (o *UpgradeScheduleQueryOpenApiVO) GetDeviceMacsOk() (*[]string, bool)`

GetDeviceMacsOk returns a tuple with the DeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacs

`func (o *UpgradeScheduleQueryOpenApiVO) SetDeviceMacs(v []string)`

SetDeviceMacs sets DeviceMacs field to given value.


### GetId

`func (o *UpgradeScheduleQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpgradeScheduleQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpgradeScheduleQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpgradeScheduleQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpgradeScheduleQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpgradeScheduleQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpgradeScheduleQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNextExecuteTime

`func (o *UpgradeScheduleQueryOpenApiVO) GetNextExecuteTime() int64`

GetNextExecuteTime returns the NextExecuteTime field if non-nil, zero value otherwise.

### GetNextExecuteTimeOk

`func (o *UpgradeScheduleQueryOpenApiVO) GetNextExecuteTimeOk() (*int64, bool)`

GetNextExecuteTimeOk returns a tuple with the NextExecuteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecuteTime

`func (o *UpgradeScheduleQueryOpenApiVO) SetNextExecuteTime(v int64)`

SetNextExecuteTime sets NextExecuteTime field to given value.

### HasNextExecuteTime

`func (o *UpgradeScheduleQueryOpenApiVO) HasNextExecuteTime() bool`

HasNextExecuteTime returns a boolean if a field has been set.

### GetOccurrenceTime

`func (o *UpgradeScheduleQueryOpenApiVO) GetOccurrenceTime() BaseScheduleTimeOpenApiVO`

GetOccurrenceTime returns the OccurrenceTime field if non-nil, zero value otherwise.

### GetOccurrenceTimeOk

`func (o *UpgradeScheduleQueryOpenApiVO) GetOccurrenceTimeOk() (*BaseScheduleTimeOpenApiVO, bool)`

GetOccurrenceTimeOk returns a tuple with the OccurrenceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceTime

`func (o *UpgradeScheduleQueryOpenApiVO) SetOccurrenceTime(v BaseScheduleTimeOpenApiVO)`

SetOccurrenceTime sets OccurrenceTime field to given value.

### HasOccurrenceTime

`func (o *UpgradeScheduleQueryOpenApiVO) HasOccurrenceTime() bool`

HasOccurrenceTime returns a boolean if a field has been set.

### GetStatus

`func (o *UpgradeScheduleQueryOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpgradeScheduleQueryOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpgradeScheduleQueryOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetType

`func (o *UpgradeScheduleQueryOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpgradeScheduleQueryOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpgradeScheduleQueryOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


