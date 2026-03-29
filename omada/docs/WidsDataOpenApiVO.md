# WidsDataOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | Pointer to **int32** | Channel band of the AP detecting the event (0:2.4g,1:5g, 2:5G,3:6g). | [optional] 
**Channel** | Pointer to **int32** | The Channel of the AP that detects the event. | [optional] 
**DetectMac** | Pointer to **string** | The Mac address of the AP that detects the event. | [optional] 
**Level** | Pointer to **int32** | Detection event level (0:Alert,1:Major,2:Normal). | [optional] 
**StartTime** | Pointer to **int64** | Detection time. | [optional] 
**StationMac** | Pointer to **string** | MAC address of the device that triggered the event. | [optional] 
**StationMacList** | Pointer to **[]string** | List of MAC addresses of the device that triggered the event. | [optional] 
**Type** | Pointer to **int32** | Detected event type. | [optional] 

## Methods

### NewWidsDataOpenApiVO

`func NewWidsDataOpenApiVO() *WidsDataOpenApiVO`

NewWidsDataOpenApiVO instantiates a new WidsDataOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidsDataOpenApiVOWithDefaults

`func NewWidsDataOpenApiVOWithDefaults() *WidsDataOpenApiVO`

NewWidsDataOpenApiVOWithDefaults instantiates a new WidsDataOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *WidsDataOpenApiVO) GetBand() int32`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *WidsDataOpenApiVO) GetBandOk() (*int32, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *WidsDataOpenApiVO) SetBand(v int32)`

SetBand sets Band field to given value.

### HasBand

`func (o *WidsDataOpenApiVO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetChannel

`func (o *WidsDataOpenApiVO) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *WidsDataOpenApiVO) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *WidsDataOpenApiVO) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *WidsDataOpenApiVO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDetectMac

`func (o *WidsDataOpenApiVO) GetDetectMac() string`

GetDetectMac returns the DetectMac field if non-nil, zero value otherwise.

### GetDetectMacOk

`func (o *WidsDataOpenApiVO) GetDetectMacOk() (*string, bool)`

GetDetectMacOk returns a tuple with the DetectMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectMac

`func (o *WidsDataOpenApiVO) SetDetectMac(v string)`

SetDetectMac sets DetectMac field to given value.

### HasDetectMac

`func (o *WidsDataOpenApiVO) HasDetectMac() bool`

HasDetectMac returns a boolean if a field has been set.

### GetLevel

`func (o *WidsDataOpenApiVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *WidsDataOpenApiVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *WidsDataOpenApiVO) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *WidsDataOpenApiVO) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetStartTime

`func (o *WidsDataOpenApiVO) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *WidsDataOpenApiVO) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *WidsDataOpenApiVO) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *WidsDataOpenApiVO) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStationMac

`func (o *WidsDataOpenApiVO) GetStationMac() string`

GetStationMac returns the StationMac field if non-nil, zero value otherwise.

### GetStationMacOk

`func (o *WidsDataOpenApiVO) GetStationMacOk() (*string, bool)`

GetStationMacOk returns a tuple with the StationMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationMac

`func (o *WidsDataOpenApiVO) SetStationMac(v string)`

SetStationMac sets StationMac field to given value.

### HasStationMac

`func (o *WidsDataOpenApiVO) HasStationMac() bool`

HasStationMac returns a boolean if a field has been set.

### GetStationMacList

`func (o *WidsDataOpenApiVO) GetStationMacList() []string`

GetStationMacList returns the StationMacList field if non-nil, zero value otherwise.

### GetStationMacListOk

`func (o *WidsDataOpenApiVO) GetStationMacListOk() (*[]string, bool)`

GetStationMacListOk returns a tuple with the StationMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStationMacList

`func (o *WidsDataOpenApiVO) SetStationMacList(v []string)`

SetStationMacList sets StationMacList field to given value.

### HasStationMacList

`func (o *WidsDataOpenApiVO) HasStationMacList() bool`

HasStationMacList returns a boolean if a field has been set.

### GetType

`func (o *WidsDataOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WidsDataOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WidsDataOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *WidsDataOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


