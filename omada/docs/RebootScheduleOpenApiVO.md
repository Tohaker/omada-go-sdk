# RebootScheduleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMacs** | **[]string** | MAC address of the selected devices. E.g. AA-BB-CC-DD-11-22 | 
**Id** | Pointer to **string** | Reboot Schedule Id. | [optional] 
**Name** | **string** | Reboot Schedule name should contain 1 to 128 characters. | 
**Status** | **bool** | Reboot Schedule status. | 
**Time** | [**RebootScheduleTimeOpenApiVO**](RebootScheduleTimeOpenApiVO.md) |  | 

## Methods

### NewRebootScheduleOpenApiVO

`func NewRebootScheduleOpenApiVO(deviceMacs []string, name string, status bool, time RebootScheduleTimeOpenApiVO, ) *RebootScheduleOpenApiVO`

NewRebootScheduleOpenApiVO instantiates a new RebootScheduleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootScheduleOpenApiVOWithDefaults

`func NewRebootScheduleOpenApiVOWithDefaults() *RebootScheduleOpenApiVO`

NewRebootScheduleOpenApiVOWithDefaults instantiates a new RebootScheduleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMacs

`func (o *RebootScheduleOpenApiVO) GetDeviceMacs() []string`

GetDeviceMacs returns the DeviceMacs field if non-nil, zero value otherwise.

### GetDeviceMacsOk

`func (o *RebootScheduleOpenApiVO) GetDeviceMacsOk() (*[]string, bool)`

GetDeviceMacsOk returns a tuple with the DeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacs

`func (o *RebootScheduleOpenApiVO) SetDeviceMacs(v []string)`

SetDeviceMacs sets DeviceMacs field to given value.


### GetId

`func (o *RebootScheduleOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RebootScheduleOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RebootScheduleOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RebootScheduleOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RebootScheduleOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RebootScheduleOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RebootScheduleOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *RebootScheduleOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RebootScheduleOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RebootScheduleOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTime

`func (o *RebootScheduleOpenApiVO) GetTime() RebootScheduleTimeOpenApiVO`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RebootScheduleOpenApiVO) GetTimeOk() (*RebootScheduleTimeOpenApiVO, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RebootScheduleOpenApiVO) SetTime(v RebootScheduleTimeOpenApiVO)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


