# OnBoardingTimeSubHealthDetailVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssocTime** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | Client association cost time list. | [optional] 
**AuthTime** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | Client authorization cost time list. | [optional] 
**AverageAssocTime** | Pointer to **int32** | Average association time. | [optional] 
**AverageAuthTime** | Pointer to **int32** | Average authorization time. | [optional] 
**AverageDhcpTime** | Pointer to **int32** | Average DHCP time. | [optional] 
**AverageDnsTime** | Pointer to **int32** | Average DNS time. | [optional] 
**DhcpTime** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | Client DHCP cost time list. | [optional] 
**DnsTime** | Pointer to [**[]TimeValueItemVO**](TimeValueItemVO.md) | Client dns cost time list. | [optional] 
**Incidents** | Pointer to [**[]AnomalyBriefCountVO**](AnomalyBriefCountVO.md) | Incident information for this health dimension, null if no incidents | [optional] 
**SummaryScore** | Pointer to **int32** | Sub dimension health score | [optional] 
**Support** | Pointer to **bool** | Sub dimension support | [optional] 

## Methods

### NewOnBoardingTimeSubHealthDetailVO

`func NewOnBoardingTimeSubHealthDetailVO() *OnBoardingTimeSubHealthDetailVO`

NewOnBoardingTimeSubHealthDetailVO instantiates a new OnBoardingTimeSubHealthDetailVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnBoardingTimeSubHealthDetailVOWithDefaults

`func NewOnBoardingTimeSubHealthDetailVOWithDefaults() *OnBoardingTimeSubHealthDetailVO`

NewOnBoardingTimeSubHealthDetailVOWithDefaults instantiates a new OnBoardingTimeSubHealthDetailVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssocTime

`func (o *OnBoardingTimeSubHealthDetailVO) GetAssocTime() []TimeValueItemVO`

GetAssocTime returns the AssocTime field if non-nil, zero value otherwise.

### GetAssocTimeOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetAssocTimeOk() (*[]TimeValueItemVO, bool)`

GetAssocTimeOk returns a tuple with the AssocTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssocTime

`func (o *OnBoardingTimeSubHealthDetailVO) SetAssocTime(v []TimeValueItemVO)`

SetAssocTime sets AssocTime field to given value.

### HasAssocTime

`func (o *OnBoardingTimeSubHealthDetailVO) HasAssocTime() bool`

HasAssocTime returns a boolean if a field has been set.

### GetAuthTime

`func (o *OnBoardingTimeSubHealthDetailVO) GetAuthTime() []TimeValueItemVO`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetAuthTimeOk() (*[]TimeValueItemVO, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *OnBoardingTimeSubHealthDetailVO) SetAuthTime(v []TimeValueItemVO)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *OnBoardingTimeSubHealthDetailVO) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetAverageAssocTime

`func (o *OnBoardingTimeSubHealthDetailVO) GetAverageAssocTime() int32`

GetAverageAssocTime returns the AverageAssocTime field if non-nil, zero value otherwise.

### GetAverageAssocTimeOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetAverageAssocTimeOk() (*int32, bool)`

GetAverageAssocTimeOk returns a tuple with the AverageAssocTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageAssocTime

`func (o *OnBoardingTimeSubHealthDetailVO) SetAverageAssocTime(v int32)`

SetAverageAssocTime sets AverageAssocTime field to given value.

### HasAverageAssocTime

`func (o *OnBoardingTimeSubHealthDetailVO) HasAverageAssocTime() bool`

HasAverageAssocTime returns a boolean if a field has been set.

### GetAverageAuthTime

`func (o *OnBoardingTimeSubHealthDetailVO) GetAverageAuthTime() int32`

GetAverageAuthTime returns the AverageAuthTime field if non-nil, zero value otherwise.

### GetAverageAuthTimeOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetAverageAuthTimeOk() (*int32, bool)`

GetAverageAuthTimeOk returns a tuple with the AverageAuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageAuthTime

`func (o *OnBoardingTimeSubHealthDetailVO) SetAverageAuthTime(v int32)`

SetAverageAuthTime sets AverageAuthTime field to given value.

### HasAverageAuthTime

`func (o *OnBoardingTimeSubHealthDetailVO) HasAverageAuthTime() bool`

HasAverageAuthTime returns a boolean if a field has been set.

### GetAverageDhcpTime

`func (o *OnBoardingTimeSubHealthDetailVO) GetAverageDhcpTime() int32`

GetAverageDhcpTime returns the AverageDhcpTime field if non-nil, zero value otherwise.

### GetAverageDhcpTimeOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetAverageDhcpTimeOk() (*int32, bool)`

GetAverageDhcpTimeOk returns a tuple with the AverageDhcpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDhcpTime

`func (o *OnBoardingTimeSubHealthDetailVO) SetAverageDhcpTime(v int32)`

SetAverageDhcpTime sets AverageDhcpTime field to given value.

### HasAverageDhcpTime

`func (o *OnBoardingTimeSubHealthDetailVO) HasAverageDhcpTime() bool`

HasAverageDhcpTime returns a boolean if a field has been set.

### GetAverageDnsTime

`func (o *OnBoardingTimeSubHealthDetailVO) GetAverageDnsTime() int32`

GetAverageDnsTime returns the AverageDnsTime field if non-nil, zero value otherwise.

### GetAverageDnsTimeOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetAverageDnsTimeOk() (*int32, bool)`

GetAverageDnsTimeOk returns a tuple with the AverageDnsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageDnsTime

`func (o *OnBoardingTimeSubHealthDetailVO) SetAverageDnsTime(v int32)`

SetAverageDnsTime sets AverageDnsTime field to given value.

### HasAverageDnsTime

`func (o *OnBoardingTimeSubHealthDetailVO) HasAverageDnsTime() bool`

HasAverageDnsTime returns a boolean if a field has been set.

### GetDhcpTime

`func (o *OnBoardingTimeSubHealthDetailVO) GetDhcpTime() []TimeValueItemVO`

GetDhcpTime returns the DhcpTime field if non-nil, zero value otherwise.

### GetDhcpTimeOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetDhcpTimeOk() (*[]TimeValueItemVO, bool)`

GetDhcpTimeOk returns a tuple with the DhcpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpTime

`func (o *OnBoardingTimeSubHealthDetailVO) SetDhcpTime(v []TimeValueItemVO)`

SetDhcpTime sets DhcpTime field to given value.

### HasDhcpTime

`func (o *OnBoardingTimeSubHealthDetailVO) HasDhcpTime() bool`

HasDhcpTime returns a boolean if a field has been set.

### GetDnsTime

`func (o *OnBoardingTimeSubHealthDetailVO) GetDnsTime() []TimeValueItemVO`

GetDnsTime returns the DnsTime field if non-nil, zero value otherwise.

### GetDnsTimeOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetDnsTimeOk() (*[]TimeValueItemVO, bool)`

GetDnsTimeOk returns a tuple with the DnsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTime

`func (o *OnBoardingTimeSubHealthDetailVO) SetDnsTime(v []TimeValueItemVO)`

SetDnsTime sets DnsTime field to given value.

### HasDnsTime

`func (o *OnBoardingTimeSubHealthDetailVO) HasDnsTime() bool`

HasDnsTime returns a boolean if a field has been set.

### GetIncidents

`func (o *OnBoardingTimeSubHealthDetailVO) GetIncidents() []AnomalyBriefCountVO`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetIncidentsOk() (*[]AnomalyBriefCountVO, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *OnBoardingTimeSubHealthDetailVO) SetIncidents(v []AnomalyBriefCountVO)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *OnBoardingTimeSubHealthDetailVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetSummaryScore

`func (o *OnBoardingTimeSubHealthDetailVO) GetSummaryScore() int32`

GetSummaryScore returns the SummaryScore field if non-nil, zero value otherwise.

### GetSummaryScoreOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetSummaryScoreOk() (*int32, bool)`

GetSummaryScoreOk returns a tuple with the SummaryScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryScore

`func (o *OnBoardingTimeSubHealthDetailVO) SetSummaryScore(v int32)`

SetSummaryScore sets SummaryScore field to given value.

### HasSummaryScore

`func (o *OnBoardingTimeSubHealthDetailVO) HasSummaryScore() bool`

HasSummaryScore returns a boolean if a field has been set.

### GetSupport

`func (o *OnBoardingTimeSubHealthDetailVO) GetSupport() bool`

GetSupport returns the Support field if non-nil, zero value otherwise.

### GetSupportOk

`func (o *OnBoardingTimeSubHealthDetailVO) GetSupportOk() (*bool, bool)`

GetSupportOk returns a tuple with the Support field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport

`func (o *OnBoardingTimeSubHealthDetailVO) SetSupport(v bool)`

SetSupport sets Support field to given value.

### HasSupport

`func (o *OnBoardingTimeSubHealthDetailVO) HasSupport() bool`

HasSupport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


