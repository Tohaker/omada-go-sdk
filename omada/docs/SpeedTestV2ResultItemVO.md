# SpeedTestV2ResultItemVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Down** | Pointer to **int64** | The download speed of this wan port speed test result. | [optional] 
**Isp** | Pointer to **string** | Name of isp. | [optional] 
**Latency** | Pointer to **int32** | The latency of this wan port speed test result. | [optional] 
**PortId** | Pointer to **int32** | The id of physical wan port. | [optional] 
**PortName** | Pointer to **string** | The name of wan or virtual wan. | [optional] 
**Progress** | Pointer to **float64** | The progress of this wan port speed test result. | [optional] 
**ServerLocation** | Pointer to **string** | Location of speed test server. | [optional] 
**ServerName** | Pointer to **string** | Name of speed test server. | [optional] 
**Status** | Pointer to **int32** | The status of this wan port speed test result. | [optional] 
**Time** | Pointer to **int64** | Time of speed test. | [optional] 
**Up** | Pointer to **int64** | The upload speed of this wan port speed test result. | [optional] 
**VirtualWanEntryId** | Pointer to **int32** | The entry id of virtual wan. | [optional] 

## Methods

### NewSpeedTestV2ResultItemVO

`func NewSpeedTestV2ResultItemVO() *SpeedTestV2ResultItemVO`

NewSpeedTestV2ResultItemVO instantiates a new SpeedTestV2ResultItemVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeedTestV2ResultItemVOWithDefaults

`func NewSpeedTestV2ResultItemVOWithDefaults() *SpeedTestV2ResultItemVO`

NewSpeedTestV2ResultItemVOWithDefaults instantiates a new SpeedTestV2ResultItemVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDown

`func (o *SpeedTestV2ResultItemVO) GetDown() int64`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *SpeedTestV2ResultItemVO) GetDownOk() (*int64, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *SpeedTestV2ResultItemVO) SetDown(v int64)`

SetDown sets Down field to given value.

### HasDown

`func (o *SpeedTestV2ResultItemVO) HasDown() bool`

HasDown returns a boolean if a field has been set.

### GetIsp

`func (o *SpeedTestV2ResultItemVO) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *SpeedTestV2ResultItemVO) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *SpeedTestV2ResultItemVO) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *SpeedTestV2ResultItemVO) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetLatency

`func (o *SpeedTestV2ResultItemVO) GetLatency() int32`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *SpeedTestV2ResultItemVO) GetLatencyOk() (*int32, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *SpeedTestV2ResultItemVO) SetLatency(v int32)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *SpeedTestV2ResultItemVO) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetPortId

`func (o *SpeedTestV2ResultItemVO) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *SpeedTestV2ResultItemVO) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *SpeedTestV2ResultItemVO) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *SpeedTestV2ResultItemVO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortName

`func (o *SpeedTestV2ResultItemVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *SpeedTestV2ResultItemVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *SpeedTestV2ResultItemVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *SpeedTestV2ResultItemVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetProgress

`func (o *SpeedTestV2ResultItemVO) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *SpeedTestV2ResultItemVO) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *SpeedTestV2ResultItemVO) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *SpeedTestV2ResultItemVO) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetServerLocation

`func (o *SpeedTestV2ResultItemVO) GetServerLocation() string`

GetServerLocation returns the ServerLocation field if non-nil, zero value otherwise.

### GetServerLocationOk

`func (o *SpeedTestV2ResultItemVO) GetServerLocationOk() (*string, bool)`

GetServerLocationOk returns a tuple with the ServerLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLocation

`func (o *SpeedTestV2ResultItemVO) SetServerLocation(v string)`

SetServerLocation sets ServerLocation field to given value.

### HasServerLocation

`func (o *SpeedTestV2ResultItemVO) HasServerLocation() bool`

HasServerLocation returns a boolean if a field has been set.

### GetServerName

`func (o *SpeedTestV2ResultItemVO) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *SpeedTestV2ResultItemVO) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *SpeedTestV2ResultItemVO) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *SpeedTestV2ResultItemVO) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetStatus

`func (o *SpeedTestV2ResultItemVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpeedTestV2ResultItemVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpeedTestV2ResultItemVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpeedTestV2ResultItemVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *SpeedTestV2ResultItemVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *SpeedTestV2ResultItemVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *SpeedTestV2ResultItemVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *SpeedTestV2ResultItemVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUp

`func (o *SpeedTestV2ResultItemVO) GetUp() int64`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *SpeedTestV2ResultItemVO) GetUpOk() (*int64, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *SpeedTestV2ResultItemVO) SetUp(v int64)`

SetUp sets Up field to given value.

### HasUp

`func (o *SpeedTestV2ResultItemVO) HasUp() bool`

HasUp returns a boolean if a field has been set.

### GetVirtualWanEntryId

`func (o *SpeedTestV2ResultItemVO) GetVirtualWanEntryId() int32`

GetVirtualWanEntryId returns the VirtualWanEntryId field if non-nil, zero value otherwise.

### GetVirtualWanEntryIdOk

`func (o *SpeedTestV2ResultItemVO) GetVirtualWanEntryIdOk() (*int32, bool)`

GetVirtualWanEntryIdOk returns a tuple with the VirtualWanEntryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanEntryId

`func (o *SpeedTestV2ResultItemVO) SetVirtualWanEntryId(v int32)`

SetVirtualWanEntryId sets VirtualWanEntryId field to given value.

### HasVirtualWanEntryId

`func (o *SpeedTestV2ResultItemVO) HasVirtualWanEntryId() bool`

HasVirtualWanEntryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


