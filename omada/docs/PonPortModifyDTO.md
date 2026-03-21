# PonPortModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDetectInterval** | **int32** | Set the interval time for port traffic ONU detection, unit: minutes.AutoDetectInternal should be within the range of 1 to 100.Default value:15 | 
**DbaCalculateMode** | Pointer to **string** | Configure the DBA calculation cycle mode. DbaCalculateMode should be a value as follows: MIN_DELAY:In minimum delay mode, the DBA calculation cycle is related to the number of T-CONTs. The more T-CONTs there are, the longer the calculation cycle. Under the premise that the bandwidth can complete the calculation, it is advisable to use as few calculation cycles as possible. When dealing with TDM services, the minimum delay mode must be selected; MAX_BW:In maximum bandwidth utilization mode, the DBA calculation cycle is allocated over multiple frames, with each frame lasting 125 µs. This method is suitable when the service does not have high delay requirements. | [optional] 
**DownstreamFEC** | Pointer to **string** | Whether to start downstream FEC function of the port.DownstreamFEC should be a value as follows:DISABLE,ENABLE. | [optional] 
**KeyExchangePeriod** | **int32** | Configure the downstream encryption ONT key update time. KeyExchangePeriod should be within the range of 0 to 60 minutes, where 0 means key update is disabled. | 
**LongLaserOnuAutoDetect** | Pointer to **string** | Whether to enable port rogue ONU automatic detection function.LongLaserOnuAutoDetect should be a value as follows:DISABLE,ENABLE.Default value:DISABLE | [optional] 
**LongLaserOnuAutoIsolate** | Pointer to **string** | Whether to enable port rogue ONU automatic isolation function.LongLaserOnuAutoIsolate should be a value as follows:DISABLE,ENABLE.Default value:DISABLE | [optional] 
**MaxDistance** | **int32** | The maximum logical distance of the ONT.MaxDistance should be within the range of 1 to 60 (km). The difference between the maximum logical distance and the minimum logical distance must not exceed 40 km. | 
**MinDistance** | **int32** | The minimum logical distance of the ONT.MinDistance should be within the range of 1 to 40 (km). The difference between the maximum logical distance and the minimum logical distance must not exceed 40 km. | 
**PortId** | **string** | Port ID | 
**Status** | Pointer to **string** | Whether to enable the laser for the PON port.Status should be a value as follows:DISABLE,ENABLE. | [optional] 

## Methods

### NewPonPortModifyDTO

`func NewPonPortModifyDTO(autoDetectInterval int32, keyExchangePeriod int32, maxDistance int32, minDistance int32, portId string, ) *PonPortModifyDTO`

NewPonPortModifyDTO instantiates a new PonPortModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPonPortModifyDTOWithDefaults

`func NewPonPortModifyDTOWithDefaults() *PonPortModifyDTO`

NewPonPortModifyDTOWithDefaults instantiates a new PonPortModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDetectInterval

`func (o *PonPortModifyDTO) GetAutoDetectInterval() int32`

GetAutoDetectInterval returns the AutoDetectInterval field if non-nil, zero value otherwise.

### GetAutoDetectIntervalOk

`func (o *PonPortModifyDTO) GetAutoDetectIntervalOk() (*int32, bool)`

GetAutoDetectIntervalOk returns a tuple with the AutoDetectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetectInterval

`func (o *PonPortModifyDTO) SetAutoDetectInterval(v int32)`

SetAutoDetectInterval sets AutoDetectInterval field to given value.


### GetDbaCalculateMode

`func (o *PonPortModifyDTO) GetDbaCalculateMode() string`

GetDbaCalculateMode returns the DbaCalculateMode field if non-nil, zero value otherwise.

### GetDbaCalculateModeOk

`func (o *PonPortModifyDTO) GetDbaCalculateModeOk() (*string, bool)`

GetDbaCalculateModeOk returns a tuple with the DbaCalculateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaCalculateMode

`func (o *PonPortModifyDTO) SetDbaCalculateMode(v string)`

SetDbaCalculateMode sets DbaCalculateMode field to given value.

### HasDbaCalculateMode

`func (o *PonPortModifyDTO) HasDbaCalculateMode() bool`

HasDbaCalculateMode returns a boolean if a field has been set.

### GetDownstreamFEC

`func (o *PonPortModifyDTO) GetDownstreamFEC() string`

GetDownstreamFEC returns the DownstreamFEC field if non-nil, zero value otherwise.

### GetDownstreamFECOk

`func (o *PonPortModifyDTO) GetDownstreamFECOk() (*string, bool)`

GetDownstreamFECOk returns a tuple with the DownstreamFEC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamFEC

`func (o *PonPortModifyDTO) SetDownstreamFEC(v string)`

SetDownstreamFEC sets DownstreamFEC field to given value.

### HasDownstreamFEC

`func (o *PonPortModifyDTO) HasDownstreamFEC() bool`

HasDownstreamFEC returns a boolean if a field has been set.

### GetKeyExchangePeriod

`func (o *PonPortModifyDTO) GetKeyExchangePeriod() int32`

GetKeyExchangePeriod returns the KeyExchangePeriod field if non-nil, zero value otherwise.

### GetKeyExchangePeriodOk

`func (o *PonPortModifyDTO) GetKeyExchangePeriodOk() (*int32, bool)`

GetKeyExchangePeriodOk returns a tuple with the KeyExchangePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchangePeriod

`func (o *PonPortModifyDTO) SetKeyExchangePeriod(v int32)`

SetKeyExchangePeriod sets KeyExchangePeriod field to given value.


### GetLongLaserOnuAutoDetect

`func (o *PonPortModifyDTO) GetLongLaserOnuAutoDetect() string`

GetLongLaserOnuAutoDetect returns the LongLaserOnuAutoDetect field if non-nil, zero value otherwise.

### GetLongLaserOnuAutoDetectOk

`func (o *PonPortModifyDTO) GetLongLaserOnuAutoDetectOk() (*string, bool)`

GetLongLaserOnuAutoDetectOk returns a tuple with the LongLaserOnuAutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongLaserOnuAutoDetect

`func (o *PonPortModifyDTO) SetLongLaserOnuAutoDetect(v string)`

SetLongLaserOnuAutoDetect sets LongLaserOnuAutoDetect field to given value.

### HasLongLaserOnuAutoDetect

`func (o *PonPortModifyDTO) HasLongLaserOnuAutoDetect() bool`

HasLongLaserOnuAutoDetect returns a boolean if a field has been set.

### GetLongLaserOnuAutoIsolate

`func (o *PonPortModifyDTO) GetLongLaserOnuAutoIsolate() string`

GetLongLaserOnuAutoIsolate returns the LongLaserOnuAutoIsolate field if non-nil, zero value otherwise.

### GetLongLaserOnuAutoIsolateOk

`func (o *PonPortModifyDTO) GetLongLaserOnuAutoIsolateOk() (*string, bool)`

GetLongLaserOnuAutoIsolateOk returns a tuple with the LongLaserOnuAutoIsolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongLaserOnuAutoIsolate

`func (o *PonPortModifyDTO) SetLongLaserOnuAutoIsolate(v string)`

SetLongLaserOnuAutoIsolate sets LongLaserOnuAutoIsolate field to given value.

### HasLongLaserOnuAutoIsolate

`func (o *PonPortModifyDTO) HasLongLaserOnuAutoIsolate() bool`

HasLongLaserOnuAutoIsolate returns a boolean if a field has been set.

### GetMaxDistance

`func (o *PonPortModifyDTO) GetMaxDistance() int32`

GetMaxDistance returns the MaxDistance field if non-nil, zero value otherwise.

### GetMaxDistanceOk

`func (o *PonPortModifyDTO) GetMaxDistanceOk() (*int32, bool)`

GetMaxDistanceOk returns a tuple with the MaxDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistance

`func (o *PonPortModifyDTO) SetMaxDistance(v int32)`

SetMaxDistance sets MaxDistance field to given value.


### GetMinDistance

`func (o *PonPortModifyDTO) GetMinDistance() int32`

GetMinDistance returns the MinDistance field if non-nil, zero value otherwise.

### GetMinDistanceOk

`func (o *PonPortModifyDTO) GetMinDistanceOk() (*int32, bool)`

GetMinDistanceOk returns a tuple with the MinDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDistance

`func (o *PonPortModifyDTO) SetMinDistance(v int32)`

SetMinDistance sets MinDistance field to given value.


### GetPortId

`func (o *PonPortModifyDTO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PonPortModifyDTO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PonPortModifyDTO) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetStatus

`func (o *PonPortModifyDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PonPortModifyDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PonPortModifyDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PonPortModifyDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


