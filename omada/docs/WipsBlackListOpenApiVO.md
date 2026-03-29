# WipsBlackListOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectMac** | Pointer to **string** | The Mac address of the AP that detects the event. | [optional] 
**StationMac** | Pointer to **string** | MAC address of the device that triggered the event. | [optional] 
**TimeStamp** | Pointer to **int64** | Timestamp of the detected event. | [optional] 

## Methods

### NewWipsBlackListOpenApiVO

`func NewWipsBlackListOpenApiVO() *WipsBlackListOpenApiVO`

NewWipsBlackListOpenApiVO instantiates a new WipsBlackListOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWipsBlackListOpenApiVOWithDefaults

`func NewWipsBlackListOpenApiVOWithDefaults() *WipsBlackListOpenApiVO`

NewWipsBlackListOpenApiVOWithDefaults instantiates a new WipsBlackListOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectMac

`func (o *WipsBlackListOpenApiVO) GetDetectMac() string`

GetDetectMac returns the DetectMac field if non-nil, zero value otherwise.

### GetDetectMacOk

`func (o *WipsBlackListOpenApiVO) GetDetectMacOk() (*string, bool)`

GetDetectMacOk returns a tuple with the DetectMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectMac

`func (o *WipsBlackListOpenApiVO) SetDetectMac(v string)`

SetDetectMac sets DetectMac field to given value.

### HasDetectMac

`func (o *WipsBlackListOpenApiVO) HasDetectMac() bool`

HasDetectMac returns a boolean if a field has been set.

### GetStationMac

`func (o *WipsBlackListOpenApiVO) GetStationMac() string`

GetStationMac returns the StationMac field if non-nil, zero value otherwise.

### GetStationMacOk

`func (o *WipsBlackListOpenApiVO) GetStationMacOk() (*string, bool)`

GetStationMacOk returns a tuple with the StationMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationMac

`func (o *WipsBlackListOpenApiVO) SetStationMac(v string)`

SetStationMac sets StationMac field to given value.

### HasStationMac

`func (o *WipsBlackListOpenApiVO) HasStationMac() bool`

HasStationMac returns a boolean if a field has been set.

### GetTimeStamp

`func (o *WipsBlackListOpenApiVO) GetTimeStamp() int64`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *WipsBlackListOpenApiVO) GetTimeStampOk() (*int64, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *WipsBlackListOpenApiVO) SetTimeStamp(v int64)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *WipsBlackListOpenApiVO) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


