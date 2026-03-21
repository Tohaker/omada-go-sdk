# PonPortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoDetectInterval** | Pointer to **int32** | Set the interval time for port traffic ONU detection, unit: minutes.AutoDetectInternal should be within the range of 1 to 100.Default value:15 | [optional] 
**DbaCalculateMode** | Pointer to **string** | Configure the DBA calculation cycle mode. DbaCalculateMode should be a value as follows: MIN_DELAY:In minimum delay mode, the DBA calculation cycle is related to the number of T-CONTs. The more T-CONTs there are, the longer the calculation cycle. Under the premise that the bandwidth can complete the calculation, it is advisable to use as few calculation cycles as possible. When dealing with TDM services, the minimum delay mode must be selected; MAX_BW:In maximum bandwidth utilization mode, the DBA calculation cycle is allocated over multiple frames, with each frame lasting 125 µs. This method is suitable when the service does not have high delay requirements. | [optional] 
**DownLinkSpeed** | Pointer to **int32** | The downstream rate of the PON port. The range of values varies by model. | [optional] 
**DownlinkList** | Pointer to [**[]DeviceVO**](DeviceVO.md) | Down link device list | [optional] 
**DownstreamFEC** | Pointer to **string** | Whether to start downstream FEC function of the port.DownstreamFEC should be a value as follows:DISABLE,ENABLE. | [optional] 
**DuplexLink** | Pointer to **string** | Duplex mode.DuplexLink should be a value as follows:FULL | [optional] 
**KeyExchangePeriod** | Pointer to **int32** | Configure the downstream encryption ONT key update time. KeyExchangePeriod should be a value between 0 to 60 minutes, where 0 means key update is disabled. | [optional] 
**LinkStatus** | Pointer to **string** | Pon port link status.LinkStatus should be a value as follows:INACTIVE,ACTIVE_WORKING,ACTIVE_WORKING_NONE,ACTIVE_STANDBY | [optional] 
**LongLaserOnuAutoDetect** | Pointer to **string** | Whether to enable port rogue ONU automatic detection function.LongLaserOnuAutoDetect should be a value as follows:DISABLE,ENABLE.Default value:DISABLE | [optional] 
**LongLaserOnuAutoIsolate** | Pointer to **string** | Whether to enable port rogue ONU automatic isolation function.LongLaserOnuAutoIsolate should be a value as follows:DISABLE,ENABLE.Default value:DISABLE | [optional] 
**MaxDistance** | Pointer to **int32** | The maximum logical distance of the ONT.MaxDistance should be within the range of 1 to 60 (km). The difference between the maximum logical distance and the minimum logical distance must not exceed 40 km. | [optional] 
**MinDistance** | Pointer to **int32** | The minimum logical distance of the ONT.MinDistance should be within the range of 1 to 40 (km). The difference between the maximum logical distance and the minimum logical distance must not exceed 40 km. | [optional] 
**OnuIsolate** | Pointer to **string** | Set ONU isolation function.(only DS-P7001-04、DS-P7001-08 OLT support this function,other olts should not provide or display this field)When enabled, different ONUs on the same port cannot communicate with each other.OnuIsolate should be a value as follows:DISABLE,ENABLE. | [optional] 
**OnuNum** | Pointer to **int32** | Online ONU number | [optional] 
**PortId** | Pointer to **string** | Port ID | [optional] 
**PortIsolate** | Pointer to **string** | Set port isolation function.(only DS-P7001-04、DS-P7001-08 OLT support this function,other olts should not provide or display this field)When enabled, the port cannot communicate with other ports, and the ONU on this port cannot communicate with ONUs on other ports.PortIsolation should be a value as follows:DISABLE,ENABLE. | [optional] 
**Speed** | Pointer to **string** | Speed configured by customer.Speed should be a value as follows:GPON,XG_PON,XGS_PON,NONE | [optional] 
**Status** | Pointer to **string** | Whether to enable the laser for the PON port.Status should be a value as follows:DISABLE,ENABLE. | [optional] 
**Type** | Pointer to **string** | Pon port type,type should be a value as follows:COPPER,COMBO,SFP | [optional] 
**UpLinkSpeed** | Pointer to **int32** | The upstream rate of the PON port, unit: M. The range of values varies by model. | [optional] 

## Methods

### NewPonPortDTO

`func NewPonPortDTO() *PonPortDTO`

NewPonPortDTO instantiates a new PonPortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPonPortDTOWithDefaults

`func NewPonPortDTOWithDefaults() *PonPortDTO`

NewPonPortDTOWithDefaults instantiates a new PonPortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoDetectInterval

`func (o *PonPortDTO) GetAutoDetectInterval() int32`

GetAutoDetectInterval returns the AutoDetectInterval field if non-nil, zero value otherwise.

### GetAutoDetectIntervalOk

`func (o *PonPortDTO) GetAutoDetectIntervalOk() (*int32, bool)`

GetAutoDetectIntervalOk returns a tuple with the AutoDetectInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetectInterval

`func (o *PonPortDTO) SetAutoDetectInterval(v int32)`

SetAutoDetectInterval sets AutoDetectInterval field to given value.

### HasAutoDetectInterval

`func (o *PonPortDTO) HasAutoDetectInterval() bool`

HasAutoDetectInterval returns a boolean if a field has been set.

### GetDbaCalculateMode

`func (o *PonPortDTO) GetDbaCalculateMode() string`

GetDbaCalculateMode returns the DbaCalculateMode field if non-nil, zero value otherwise.

### GetDbaCalculateModeOk

`func (o *PonPortDTO) GetDbaCalculateModeOk() (*string, bool)`

GetDbaCalculateModeOk returns a tuple with the DbaCalculateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaCalculateMode

`func (o *PonPortDTO) SetDbaCalculateMode(v string)`

SetDbaCalculateMode sets DbaCalculateMode field to given value.

### HasDbaCalculateMode

`func (o *PonPortDTO) HasDbaCalculateMode() bool`

HasDbaCalculateMode returns a boolean if a field has been set.

### GetDownLinkSpeed

`func (o *PonPortDTO) GetDownLinkSpeed() int32`

GetDownLinkSpeed returns the DownLinkSpeed field if non-nil, zero value otherwise.

### GetDownLinkSpeedOk

`func (o *PonPortDTO) GetDownLinkSpeedOk() (*int32, bool)`

GetDownLinkSpeedOk returns a tuple with the DownLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLinkSpeed

`func (o *PonPortDTO) SetDownLinkSpeed(v int32)`

SetDownLinkSpeed sets DownLinkSpeed field to given value.

### HasDownLinkSpeed

`func (o *PonPortDTO) HasDownLinkSpeed() bool`

HasDownLinkSpeed returns a boolean if a field has been set.

### GetDownlinkList

`func (o *PonPortDTO) GetDownlinkList() []DeviceVO`

GetDownlinkList returns the DownlinkList field if non-nil, zero value otherwise.

### GetDownlinkListOk

`func (o *PonPortDTO) GetDownlinkListOk() (*[]DeviceVO, bool)`

GetDownlinkListOk returns a tuple with the DownlinkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkList

`func (o *PonPortDTO) SetDownlinkList(v []DeviceVO)`

SetDownlinkList sets DownlinkList field to given value.

### HasDownlinkList

`func (o *PonPortDTO) HasDownlinkList() bool`

HasDownlinkList returns a boolean if a field has been set.

### GetDownstreamFEC

`func (o *PonPortDTO) GetDownstreamFEC() string`

GetDownstreamFEC returns the DownstreamFEC field if non-nil, zero value otherwise.

### GetDownstreamFECOk

`func (o *PonPortDTO) GetDownstreamFECOk() (*string, bool)`

GetDownstreamFECOk returns a tuple with the DownstreamFEC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamFEC

`func (o *PonPortDTO) SetDownstreamFEC(v string)`

SetDownstreamFEC sets DownstreamFEC field to given value.

### HasDownstreamFEC

`func (o *PonPortDTO) HasDownstreamFEC() bool`

HasDownstreamFEC returns a boolean if a field has been set.

### GetDuplexLink

`func (o *PonPortDTO) GetDuplexLink() string`

GetDuplexLink returns the DuplexLink field if non-nil, zero value otherwise.

### GetDuplexLinkOk

`func (o *PonPortDTO) GetDuplexLinkOk() (*string, bool)`

GetDuplexLinkOk returns a tuple with the DuplexLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplexLink

`func (o *PonPortDTO) SetDuplexLink(v string)`

SetDuplexLink sets DuplexLink field to given value.

### HasDuplexLink

`func (o *PonPortDTO) HasDuplexLink() bool`

HasDuplexLink returns a boolean if a field has been set.

### GetKeyExchangePeriod

`func (o *PonPortDTO) GetKeyExchangePeriod() int32`

GetKeyExchangePeriod returns the KeyExchangePeriod field if non-nil, zero value otherwise.

### GetKeyExchangePeriodOk

`func (o *PonPortDTO) GetKeyExchangePeriodOk() (*int32, bool)`

GetKeyExchangePeriodOk returns a tuple with the KeyExchangePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchangePeriod

`func (o *PonPortDTO) SetKeyExchangePeriod(v int32)`

SetKeyExchangePeriod sets KeyExchangePeriod field to given value.

### HasKeyExchangePeriod

`func (o *PonPortDTO) HasKeyExchangePeriod() bool`

HasKeyExchangePeriod returns a boolean if a field has been set.

### GetLinkStatus

`func (o *PonPortDTO) GetLinkStatus() string`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *PonPortDTO) GetLinkStatusOk() (*string, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *PonPortDTO) SetLinkStatus(v string)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *PonPortDTO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetLongLaserOnuAutoDetect

`func (o *PonPortDTO) GetLongLaserOnuAutoDetect() string`

GetLongLaserOnuAutoDetect returns the LongLaserOnuAutoDetect field if non-nil, zero value otherwise.

### GetLongLaserOnuAutoDetectOk

`func (o *PonPortDTO) GetLongLaserOnuAutoDetectOk() (*string, bool)`

GetLongLaserOnuAutoDetectOk returns a tuple with the LongLaserOnuAutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongLaserOnuAutoDetect

`func (o *PonPortDTO) SetLongLaserOnuAutoDetect(v string)`

SetLongLaserOnuAutoDetect sets LongLaserOnuAutoDetect field to given value.

### HasLongLaserOnuAutoDetect

`func (o *PonPortDTO) HasLongLaserOnuAutoDetect() bool`

HasLongLaserOnuAutoDetect returns a boolean if a field has been set.

### GetLongLaserOnuAutoIsolate

`func (o *PonPortDTO) GetLongLaserOnuAutoIsolate() string`

GetLongLaserOnuAutoIsolate returns the LongLaserOnuAutoIsolate field if non-nil, zero value otherwise.

### GetLongLaserOnuAutoIsolateOk

`func (o *PonPortDTO) GetLongLaserOnuAutoIsolateOk() (*string, bool)`

GetLongLaserOnuAutoIsolateOk returns a tuple with the LongLaserOnuAutoIsolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongLaserOnuAutoIsolate

`func (o *PonPortDTO) SetLongLaserOnuAutoIsolate(v string)`

SetLongLaserOnuAutoIsolate sets LongLaserOnuAutoIsolate field to given value.

### HasLongLaserOnuAutoIsolate

`func (o *PonPortDTO) HasLongLaserOnuAutoIsolate() bool`

HasLongLaserOnuAutoIsolate returns a boolean if a field has been set.

### GetMaxDistance

`func (o *PonPortDTO) GetMaxDistance() int32`

GetMaxDistance returns the MaxDistance field if non-nil, zero value otherwise.

### GetMaxDistanceOk

`func (o *PonPortDTO) GetMaxDistanceOk() (*int32, bool)`

GetMaxDistanceOk returns a tuple with the MaxDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDistance

`func (o *PonPortDTO) SetMaxDistance(v int32)`

SetMaxDistance sets MaxDistance field to given value.

### HasMaxDistance

`func (o *PonPortDTO) HasMaxDistance() bool`

HasMaxDistance returns a boolean if a field has been set.

### GetMinDistance

`func (o *PonPortDTO) GetMinDistance() int32`

GetMinDistance returns the MinDistance field if non-nil, zero value otherwise.

### GetMinDistanceOk

`func (o *PonPortDTO) GetMinDistanceOk() (*int32, bool)`

GetMinDistanceOk returns a tuple with the MinDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDistance

`func (o *PonPortDTO) SetMinDistance(v int32)`

SetMinDistance sets MinDistance field to given value.

### HasMinDistance

`func (o *PonPortDTO) HasMinDistance() bool`

HasMinDistance returns a boolean if a field has been set.

### GetOnuIsolate

`func (o *PonPortDTO) GetOnuIsolate() string`

GetOnuIsolate returns the OnuIsolate field if non-nil, zero value otherwise.

### GetOnuIsolateOk

`func (o *PonPortDTO) GetOnuIsolateOk() (*string, bool)`

GetOnuIsolateOk returns a tuple with the OnuIsolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuIsolate

`func (o *PonPortDTO) SetOnuIsolate(v string)`

SetOnuIsolate sets OnuIsolate field to given value.

### HasOnuIsolate

`func (o *PonPortDTO) HasOnuIsolate() bool`

HasOnuIsolate returns a boolean if a field has been set.

### GetOnuNum

`func (o *PonPortDTO) GetOnuNum() int32`

GetOnuNum returns the OnuNum field if non-nil, zero value otherwise.

### GetOnuNumOk

`func (o *PonPortDTO) GetOnuNumOk() (*int32, bool)`

GetOnuNumOk returns a tuple with the OnuNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuNum

`func (o *PonPortDTO) SetOnuNum(v int32)`

SetOnuNum sets OnuNum field to given value.

### HasOnuNum

`func (o *PonPortDTO) HasOnuNum() bool`

HasOnuNum returns a boolean if a field has been set.

### GetPortId

`func (o *PonPortDTO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PonPortDTO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PonPortDTO) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PonPortDTO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortIsolate

`func (o *PonPortDTO) GetPortIsolate() string`

GetPortIsolate returns the PortIsolate field if non-nil, zero value otherwise.

### GetPortIsolateOk

`func (o *PonPortDTO) GetPortIsolateOk() (*string, bool)`

GetPortIsolateOk returns a tuple with the PortIsolate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIsolate

`func (o *PonPortDTO) SetPortIsolate(v string)`

SetPortIsolate sets PortIsolate field to given value.

### HasPortIsolate

`func (o *PonPortDTO) HasPortIsolate() bool`

HasPortIsolate returns a boolean if a field has been set.

### GetSpeed

`func (o *PonPortDTO) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *PonPortDTO) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *PonPortDTO) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *PonPortDTO) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *PonPortDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PonPortDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PonPortDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PonPortDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *PonPortDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PonPortDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PonPortDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PonPortDTO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpLinkSpeed

`func (o *PonPortDTO) GetUpLinkSpeed() int32`

GetUpLinkSpeed returns the UpLinkSpeed field if non-nil, zero value otherwise.

### GetUpLinkSpeedOk

`func (o *PonPortDTO) GetUpLinkSpeedOk() (*int32, bool)`

GetUpLinkSpeedOk returns a tuple with the UpLinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLinkSpeed

`func (o *PonPortDTO) SetUpLinkSpeed(v int32)`

SetUpLinkSpeed sets UpLinkSpeed field to given value.

### HasUpLinkSpeed

`func (o *PonPortDTO) HasUpLinkSpeed() bool`

HasUpLinkSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


