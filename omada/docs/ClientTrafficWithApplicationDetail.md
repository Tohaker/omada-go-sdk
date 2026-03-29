# ClientTrafficWithApplicationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationList** | Pointer to [**[]ApplicationTrafficWithClientCount**](ApplicationTrafficWithClientCount.md) | The applications info using by the client. | [optional] 
**ClientName** | Pointer to **string** | The name of the client. | [optional] 
**Download** | Pointer to **int64** | The download traffic used by the client. | [optional] 
**Mac** | Pointer to **string** | The mac of the client. | [optional] 
**Manager** | Pointer to **bool** | Whether the client is managed by the controller.. | [optional] 
**TotalApplications** | Pointer to **int32** | The number of applications using by the client. | [optional] 
**Traffic** | Pointer to **int64** | The total amount of traffic used by the client. | [optional] 
**TrafficPercent** | Pointer to **float64** | The percentage of download traffic used by the client. | [optional] 
**Type** | Pointer to **string** | The type of the client. | [optional] 
**Upload** | Pointer to **int64** | The upload traffic used by the client. | [optional] 

## Methods

### NewClientTrafficWithApplicationDetail

`func NewClientTrafficWithApplicationDetail() *ClientTrafficWithApplicationDetail`

NewClientTrafficWithApplicationDetail instantiates a new ClientTrafficWithApplicationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTrafficWithApplicationDetailWithDefaults

`func NewClientTrafficWithApplicationDetailWithDefaults() *ClientTrafficWithApplicationDetail`

NewClientTrafficWithApplicationDetailWithDefaults instantiates a new ClientTrafficWithApplicationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationList

`func (o *ClientTrafficWithApplicationDetail) GetApplicationList() []ApplicationTrafficWithClientCount`

GetApplicationList returns the ApplicationList field if non-nil, zero value otherwise.

### GetApplicationListOk

`func (o *ClientTrafficWithApplicationDetail) GetApplicationListOk() (*[]ApplicationTrafficWithClientCount, bool)`

GetApplicationListOk returns a tuple with the ApplicationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationList

`func (o *ClientTrafficWithApplicationDetail) SetApplicationList(v []ApplicationTrafficWithClientCount)`

SetApplicationList sets ApplicationList field to given value.

### HasApplicationList

`func (o *ClientTrafficWithApplicationDetail) HasApplicationList() bool`

HasApplicationList returns a boolean if a field has been set.

### GetClientName

`func (o *ClientTrafficWithApplicationDetail) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ClientTrafficWithApplicationDetail) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ClientTrafficWithApplicationDetail) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ClientTrafficWithApplicationDetail) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDownload

`func (o *ClientTrafficWithApplicationDetail) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ClientTrafficWithApplicationDetail) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ClientTrafficWithApplicationDetail) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ClientTrafficWithApplicationDetail) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetMac

`func (o *ClientTrafficWithApplicationDetail) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientTrafficWithApplicationDetail) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientTrafficWithApplicationDetail) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientTrafficWithApplicationDetail) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManager

`func (o *ClientTrafficWithApplicationDetail) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ClientTrafficWithApplicationDetail) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ClientTrafficWithApplicationDetail) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *ClientTrafficWithApplicationDetail) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetTotalApplications

`func (o *ClientTrafficWithApplicationDetail) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *ClientTrafficWithApplicationDetail) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *ClientTrafficWithApplicationDetail) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.

### HasTotalApplications

`func (o *ClientTrafficWithApplicationDetail) HasTotalApplications() bool`

HasTotalApplications returns a boolean if a field has been set.

### GetTraffic

`func (o *ClientTrafficWithApplicationDetail) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *ClientTrafficWithApplicationDetail) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *ClientTrafficWithApplicationDetail) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *ClientTrafficWithApplicationDetail) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetTrafficPercent

`func (o *ClientTrafficWithApplicationDetail) GetTrafficPercent() float64`

GetTrafficPercent returns the TrafficPercent field if non-nil, zero value otherwise.

### GetTrafficPercentOk

`func (o *ClientTrafficWithApplicationDetail) GetTrafficPercentOk() (*float64, bool)`

GetTrafficPercentOk returns a tuple with the TrafficPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPercent

`func (o *ClientTrafficWithApplicationDetail) SetTrafficPercent(v float64)`

SetTrafficPercent sets TrafficPercent field to given value.

### HasTrafficPercent

`func (o *ClientTrafficWithApplicationDetail) HasTrafficPercent() bool`

HasTrafficPercent returns a boolean if a field has been set.

### GetType

`func (o *ClientTrafficWithApplicationDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientTrafficWithApplicationDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientTrafficWithApplicationDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientTrafficWithApplicationDetail) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpload

`func (o *ClientTrafficWithApplicationDetail) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ClientTrafficWithApplicationDetail) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ClientTrafficWithApplicationDetail) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ClientTrafficWithApplicationDetail) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


