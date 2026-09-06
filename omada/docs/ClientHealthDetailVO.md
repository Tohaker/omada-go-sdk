# ClientHealthDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationTime** | Pointer to [**CommonSubHealthInfoDetailVOInteger**](CommonSubHealthInfoDetailVOInteger.md) |  | [optional] 
**ConnectScore** | Pointer to [**CommonSubHealthInfoDetailVOInteger**](CommonSubHealthInfoDetailVOInteger.md) |  | [optional] 
**Incident** | Pointer to [**IncidentSubHealthInfoDetailVO**](IncidentSubHealthInfoDetailVO.md) |  | [optional] 
**LinkErrorScore** | Pointer to [**CommonSubHealthInfoDetailVOInteger**](CommonSubHealthInfoDetailVOInteger.md) |  | [optional] 
**Rate** | Pointer to [**CommonSubHealthInfoDetailVOLong**](CommonSubHealthInfoDetailVOLong.md) |  | [optional] 
**Rssi** | Pointer to [**CommonSubHealthInfoDetailVOInteger**](CommonSubHealthInfoDetailVOInteger.md) |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Snr** | Pointer to [**CommonSubHealthInfoDetailVOInteger**](CommonSubHealthInfoDetailVOInteger.md) |  | [optional] 

## Methods

### NewClientHealthDetailVO

`func NewClientHealthDetailVO() *ClientHealthDetailVO`

NewClientHealthDetailVO instantiates a new ClientHealthDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientHealthDetailVOWithDefaults

`func NewClientHealthDetailVOWithDefaults() *ClientHealthDetailVO`

NewClientHealthDetailVOWithDefaults instantiates a new ClientHealthDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationTime

`func (o *ClientHealthDetailVO) GetAssociationTime() CommonSubHealthInfoDetailVOInteger`

GetAssociationTime returns the AssociationTime field if non-nil, zero value otherwise.

### GetAssociationTimeOk

`func (o *ClientHealthDetailVO) GetAssociationTimeOk() (*CommonSubHealthInfoDetailVOInteger, bool)`

GetAssociationTimeOk returns a tuple with the AssociationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTime

`func (o *ClientHealthDetailVO) SetAssociationTime(v CommonSubHealthInfoDetailVOInteger)`

SetAssociationTime sets AssociationTime field to given value.

### HasAssociationTime

`func (o *ClientHealthDetailVO) HasAssociationTime() bool`

HasAssociationTime returns a boolean if a field has been set.

### GetConnectScore

`func (o *ClientHealthDetailVO) GetConnectScore() CommonSubHealthInfoDetailVOInteger`

GetConnectScore returns the ConnectScore field if non-nil, zero value otherwise.

### GetConnectScoreOk

`func (o *ClientHealthDetailVO) GetConnectScoreOk() (*CommonSubHealthInfoDetailVOInteger, bool)`

GetConnectScoreOk returns a tuple with the ConnectScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectScore

`func (o *ClientHealthDetailVO) SetConnectScore(v CommonSubHealthInfoDetailVOInteger)`

SetConnectScore sets ConnectScore field to given value.

### HasConnectScore

`func (o *ClientHealthDetailVO) HasConnectScore() bool`

HasConnectScore returns a boolean if a field has been set.

### GetIncident

`func (o *ClientHealthDetailVO) GetIncident() IncidentSubHealthInfoDetailVO`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *ClientHealthDetailVO) GetIncidentOk() (*IncidentSubHealthInfoDetailVO, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *ClientHealthDetailVO) SetIncident(v IncidentSubHealthInfoDetailVO)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *ClientHealthDetailVO) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetLinkErrorScore

`func (o *ClientHealthDetailVO) GetLinkErrorScore() CommonSubHealthInfoDetailVOInteger`

GetLinkErrorScore returns the LinkErrorScore field if non-nil, zero value otherwise.

### GetLinkErrorScoreOk

`func (o *ClientHealthDetailVO) GetLinkErrorScoreOk() (*CommonSubHealthInfoDetailVOInteger, bool)`

GetLinkErrorScoreOk returns a tuple with the LinkErrorScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkErrorScore

`func (o *ClientHealthDetailVO) SetLinkErrorScore(v CommonSubHealthInfoDetailVOInteger)`

SetLinkErrorScore sets LinkErrorScore field to given value.

### HasLinkErrorScore

`func (o *ClientHealthDetailVO) HasLinkErrorScore() bool`

HasLinkErrorScore returns a boolean if a field has been set.

### GetRate

`func (o *ClientHealthDetailVO) GetRate() CommonSubHealthInfoDetailVOLong`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ClientHealthDetailVO) GetRateOk() (*CommonSubHealthInfoDetailVOLong, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ClientHealthDetailVO) SetRate(v CommonSubHealthInfoDetailVOLong)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ClientHealthDetailVO) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRssi

`func (o *ClientHealthDetailVO) GetRssi() CommonSubHealthInfoDetailVOInteger`

GetRssi returns the Rssi field if non-nil, zero value otherwise.

### GetRssiOk

`func (o *ClientHealthDetailVO) GetRssiOk() (*CommonSubHealthInfoDetailVOInteger, bool)`

GetRssiOk returns a tuple with the Rssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRssi

`func (o *ClientHealthDetailVO) SetRssi(v CommonSubHealthInfoDetailVOInteger)`

SetRssi sets Rssi field to given value.

### HasRssi

`func (o *ClientHealthDetailVO) HasRssi() bool`

HasRssi returns a boolean if a field has been set.

### GetScore

`func (o *ClientHealthDetailVO) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ClientHealthDetailVO) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ClientHealthDetailVO) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ClientHealthDetailVO) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSnr

`func (o *ClientHealthDetailVO) GetSnr() CommonSubHealthInfoDetailVOInteger`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *ClientHealthDetailVO) GetSnrOk() (*CommonSubHealthInfoDetailVOInteger, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *ClientHealthDetailVO) SetSnr(v CommonSubHealthInfoDetailVOInteger)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *ClientHealthDetailVO) HasSnr() bool`

HasSnr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


