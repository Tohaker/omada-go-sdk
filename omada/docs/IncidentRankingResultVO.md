# IncidentRankingResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandList** | Pointer to [**[]IncidentRankingBandItemVO**](IncidentRankingBandItemVO.md) | Top frequency bands ranked by incident count (e.g. 0&#x3D;2.4GHz, 1&#x3D;5GHz, 2&#x3D;5GHz-2, 3&#x3D;6GHz) | [optional] 
**ClientList** | Pointer to [**[]IncidentRankingClientItemVO**](IncidentRankingClientItemVO.md) | Top clients ranked by incident count | [optional] 
**DeviceList** | Pointer to [**[]IncidentRankingDeviceItemVO**](IncidentRankingDeviceItemVO.md) | Top devices ranked by incident count | [optional] 
**SsidList** | Pointer to [**[]IncidentRankingSsidItemVO**](IncidentRankingSsidItemVO.md) | Top SSIDs ranked by incident count | [optional] 

## Methods

### NewIncidentRankingResultVO

`func NewIncidentRankingResultVO() *IncidentRankingResultVO`

NewIncidentRankingResultVO instantiates a new IncidentRankingResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentRankingResultVOWithDefaults

`func NewIncidentRankingResultVOWithDefaults() *IncidentRankingResultVO`

NewIncidentRankingResultVOWithDefaults instantiates a new IncidentRankingResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandList

`func (o *IncidentRankingResultVO) GetBandList() []IncidentRankingBandItemVO`

GetBandList returns the BandList field if non-nil, zero value otherwise.

### GetBandListOk

`func (o *IncidentRankingResultVO) GetBandListOk() (*[]IncidentRankingBandItemVO, bool)`

GetBandListOk returns a tuple with the BandList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandList

`func (o *IncidentRankingResultVO) SetBandList(v []IncidentRankingBandItemVO)`

SetBandList sets BandList field to given value.

### HasBandList

`func (o *IncidentRankingResultVO) HasBandList() bool`

HasBandList returns a boolean if a field has been set.

### GetClientList

`func (o *IncidentRankingResultVO) GetClientList() []IncidentRankingClientItemVO`

GetClientList returns the ClientList field if non-nil, zero value otherwise.

### GetClientListOk

`func (o *IncidentRankingResultVO) GetClientListOk() (*[]IncidentRankingClientItemVO, bool)`

GetClientListOk returns a tuple with the ClientList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientList

`func (o *IncidentRankingResultVO) SetClientList(v []IncidentRankingClientItemVO)`

SetClientList sets ClientList field to given value.

### HasClientList

`func (o *IncidentRankingResultVO) HasClientList() bool`

HasClientList returns a boolean if a field has been set.

### GetDeviceList

`func (o *IncidentRankingResultVO) GetDeviceList() []IncidentRankingDeviceItemVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *IncidentRankingResultVO) GetDeviceListOk() (*[]IncidentRankingDeviceItemVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *IncidentRankingResultVO) SetDeviceList(v []IncidentRankingDeviceItemVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *IncidentRankingResultVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetSsidList

`func (o *IncidentRankingResultVO) GetSsidList() []IncidentRankingSsidItemVO`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *IncidentRankingResultVO) GetSsidListOk() (*[]IncidentRankingSsidItemVO, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *IncidentRankingResultVO) SetSsidList(v []IncidentRankingSsidItemVO)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *IncidentRankingResultVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


