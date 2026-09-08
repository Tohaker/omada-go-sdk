# ClientHealthDetailV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandRssi** | Pointer to [**ChannelSubHealthInfoDetailVO**](ChannelSubHealthInfoDetailVO.md) |  | [optional] 
**ConnectScore** | Pointer to [**ConnectScoreSubHealthInfoDetailVO**](ConnectScoreSubHealthInfoDetailVO.md) |  | [optional] 
**DataRate** | Pointer to [**ClientDataRateSubHealthDetailVO**](ClientDataRateSubHealthDetailVO.md) |  | [optional] 
**Incident** | Pointer to [**IncidentSubHealthInfoDetailVO**](IncidentSubHealthInfoDetailVO.md) |  | [optional] 
**LinkErrorScore** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 
**OnboardingTime** | Pointer to [**OnBoardingTimeSubHealthDetailVO**](OnBoardingTimeSubHealthDetailVO.md) |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Snr** | Pointer to [**CommonSubHealthInfoDetailVO**](CommonSubHealthInfoDetailVO.md) |  | [optional] 

## Methods

### NewClientHealthDetailV2VO

`func NewClientHealthDetailV2VO() *ClientHealthDetailV2VO`

NewClientHealthDetailV2VO instantiates a new ClientHealthDetailV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientHealthDetailV2VOWithDefaults

`func NewClientHealthDetailV2VOWithDefaults() *ClientHealthDetailV2VO`

NewClientHealthDetailV2VOWithDefaults instantiates a new ClientHealthDetailV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandRssi

`func (o *ClientHealthDetailV2VO) GetBandRssi() ChannelSubHealthInfoDetailVO`

GetBandRssi returns the BandRssi field if non-nil, zero value otherwise.

### GetBandRssiOk

`func (o *ClientHealthDetailV2VO) GetBandRssiOk() (*ChannelSubHealthInfoDetailVO, bool)`

GetBandRssiOk returns a tuple with the BandRssi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandRssi

`func (o *ClientHealthDetailV2VO) SetBandRssi(v ChannelSubHealthInfoDetailVO)`

SetBandRssi sets BandRssi field to given value.

### HasBandRssi

`func (o *ClientHealthDetailV2VO) HasBandRssi() bool`

HasBandRssi returns a boolean if a field has been set.

### GetConnectScore

`func (o *ClientHealthDetailV2VO) GetConnectScore() ConnectScoreSubHealthInfoDetailVO`

GetConnectScore returns the ConnectScore field if non-nil, zero value otherwise.

### GetConnectScoreOk

`func (o *ClientHealthDetailV2VO) GetConnectScoreOk() (*ConnectScoreSubHealthInfoDetailVO, bool)`

GetConnectScoreOk returns a tuple with the ConnectScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectScore

`func (o *ClientHealthDetailV2VO) SetConnectScore(v ConnectScoreSubHealthInfoDetailVO)`

SetConnectScore sets ConnectScore field to given value.

### HasConnectScore

`func (o *ClientHealthDetailV2VO) HasConnectScore() bool`

HasConnectScore returns a boolean if a field has been set.

### GetDataRate

`func (o *ClientHealthDetailV2VO) GetDataRate() ClientDataRateSubHealthDetailVO`

GetDataRate returns the DataRate field if non-nil, zero value otherwise.

### GetDataRateOk

`func (o *ClientHealthDetailV2VO) GetDataRateOk() (*ClientDataRateSubHealthDetailVO, bool)`

GetDataRateOk returns a tuple with the DataRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRate

`func (o *ClientHealthDetailV2VO) SetDataRate(v ClientDataRateSubHealthDetailVO)`

SetDataRate sets DataRate field to given value.

### HasDataRate

`func (o *ClientHealthDetailV2VO) HasDataRate() bool`

HasDataRate returns a boolean if a field has been set.

### GetIncident

`func (o *ClientHealthDetailV2VO) GetIncident() IncidentSubHealthInfoDetailVO`

GetIncident returns the Incident field if non-nil, zero value otherwise.

### GetIncidentOk

`func (o *ClientHealthDetailV2VO) GetIncidentOk() (*IncidentSubHealthInfoDetailVO, bool)`

GetIncidentOk returns a tuple with the Incident field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncident

`func (o *ClientHealthDetailV2VO) SetIncident(v IncidentSubHealthInfoDetailVO)`

SetIncident sets Incident field to given value.

### HasIncident

`func (o *ClientHealthDetailV2VO) HasIncident() bool`

HasIncident returns a boolean if a field has been set.

### GetLinkErrorScore

`func (o *ClientHealthDetailV2VO) GetLinkErrorScore() CommonSubHealthInfoDetailVO`

GetLinkErrorScore returns the LinkErrorScore field if non-nil, zero value otherwise.

### GetLinkErrorScoreOk

`func (o *ClientHealthDetailV2VO) GetLinkErrorScoreOk() (*CommonSubHealthInfoDetailVO, bool)`

GetLinkErrorScoreOk returns a tuple with the LinkErrorScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkErrorScore

`func (o *ClientHealthDetailV2VO) SetLinkErrorScore(v CommonSubHealthInfoDetailVO)`

SetLinkErrorScore sets LinkErrorScore field to given value.

### HasLinkErrorScore

`func (o *ClientHealthDetailV2VO) HasLinkErrorScore() bool`

HasLinkErrorScore returns a boolean if a field has been set.

### GetOnboardingTime

`func (o *ClientHealthDetailV2VO) GetOnboardingTime() OnBoardingTimeSubHealthDetailVO`

GetOnboardingTime returns the OnboardingTime field if non-nil, zero value otherwise.

### GetOnboardingTimeOk

`func (o *ClientHealthDetailV2VO) GetOnboardingTimeOk() (*OnBoardingTimeSubHealthDetailVO, bool)`

GetOnboardingTimeOk returns a tuple with the OnboardingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboardingTime

`func (o *ClientHealthDetailV2VO) SetOnboardingTime(v OnBoardingTimeSubHealthDetailVO)`

SetOnboardingTime sets OnboardingTime field to given value.

### HasOnboardingTime

`func (o *ClientHealthDetailV2VO) HasOnboardingTime() bool`

HasOnboardingTime returns a boolean if a field has been set.

### GetScore

`func (o *ClientHealthDetailV2VO) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ClientHealthDetailV2VO) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ClientHealthDetailV2VO) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ClientHealthDetailV2VO) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSnr

`func (o *ClientHealthDetailV2VO) GetSnr() CommonSubHealthInfoDetailVO`

GetSnr returns the Snr field if non-nil, zero value otherwise.

### GetSnrOk

`func (o *ClientHealthDetailV2VO) GetSnrOk() (*CommonSubHealthInfoDetailVO, bool)`

GetSnrOk returns a tuple with the Snr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnr

`func (o *ClientHealthDetailV2VO) SetSnr(v CommonSubHealthInfoDetailVO)`

SetSnr sets Snr field to given value.

### HasSnr

`func (o *ClientHealthDetailV2VO) HasSnr() bool`

HasSnr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


