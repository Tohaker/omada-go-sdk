# StatisticInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OctetsRx** | Pointer to **int64** | Octets received | [optional] 
**OctetsTx** | Pointer to **int64** | Octets transmitted | [optional] 
**PacketsRx** | Pointer to **int64** | Packets received | [optional] 
**PacketsTx** | Pointer to **int64** | Packets transmitted | [optional] 
**ServicePortIndex** | **int32** | ID of service port.ServicePortIndex should be within the range of 1 to 8100 | 

## Methods

### NewStatisticInfoDTO

`func NewStatisticInfoDTO(servicePortIndex int32, ) *StatisticInfoDTO`

NewStatisticInfoDTO instantiates a new StatisticInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticInfoDTOWithDefaults

`func NewStatisticInfoDTOWithDefaults() *StatisticInfoDTO`

NewStatisticInfoDTOWithDefaults instantiates a new StatisticInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOctetsRx

`func (o *StatisticInfoDTO) GetOctetsRx() int64`

GetOctetsRx returns the OctetsRx field if non-nil, zero value otherwise.

### GetOctetsRxOk

`func (o *StatisticInfoDTO) GetOctetsRxOk() (*int64, bool)`

GetOctetsRxOk returns a tuple with the OctetsRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctetsRx

`func (o *StatisticInfoDTO) SetOctetsRx(v int64)`

SetOctetsRx sets OctetsRx field to given value.

### HasOctetsRx

`func (o *StatisticInfoDTO) HasOctetsRx() bool`

HasOctetsRx returns a boolean if a field has been set.

### GetOctetsTx

`func (o *StatisticInfoDTO) GetOctetsTx() int64`

GetOctetsTx returns the OctetsTx field if non-nil, zero value otherwise.

### GetOctetsTxOk

`func (o *StatisticInfoDTO) GetOctetsTxOk() (*int64, bool)`

GetOctetsTxOk returns a tuple with the OctetsTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOctetsTx

`func (o *StatisticInfoDTO) SetOctetsTx(v int64)`

SetOctetsTx sets OctetsTx field to given value.

### HasOctetsTx

`func (o *StatisticInfoDTO) HasOctetsTx() bool`

HasOctetsTx returns a boolean if a field has been set.

### GetPacketsRx

`func (o *StatisticInfoDTO) GetPacketsRx() int64`

GetPacketsRx returns the PacketsRx field if non-nil, zero value otherwise.

### GetPacketsRxOk

`func (o *StatisticInfoDTO) GetPacketsRxOk() (*int64, bool)`

GetPacketsRxOk returns a tuple with the PacketsRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsRx

`func (o *StatisticInfoDTO) SetPacketsRx(v int64)`

SetPacketsRx sets PacketsRx field to given value.

### HasPacketsRx

`func (o *StatisticInfoDTO) HasPacketsRx() bool`

HasPacketsRx returns a boolean if a field has been set.

### GetPacketsTx

`func (o *StatisticInfoDTO) GetPacketsTx() int64`

GetPacketsTx returns the PacketsTx field if non-nil, zero value otherwise.

### GetPacketsTxOk

`func (o *StatisticInfoDTO) GetPacketsTxOk() (*int64, bool)`

GetPacketsTxOk returns a tuple with the PacketsTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketsTx

`func (o *StatisticInfoDTO) SetPacketsTx(v int64)`

SetPacketsTx sets PacketsTx field to given value.

### HasPacketsTx

`func (o *StatisticInfoDTO) HasPacketsTx() bool`

HasPacketsTx returns a boolean if a field has been set.

### GetServicePortIndex

`func (o *StatisticInfoDTO) GetServicePortIndex() int32`

GetServicePortIndex returns the ServicePortIndex field if non-nil, zero value otherwise.

### GetServicePortIndexOk

`func (o *StatisticInfoDTO) GetServicePortIndexOk() (*int32, bool)`

GetServicePortIndexOk returns a tuple with the ServicePortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortIndex

`func (o *StatisticInfoDTO) SetServicePortIndex(v int32)`

SetServicePortIndex sets ServicePortIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


