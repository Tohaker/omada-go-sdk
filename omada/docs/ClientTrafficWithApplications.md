# ClientTrafficWithApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationList** | Pointer to [**[]ApplicationBasicInfo**](ApplicationBasicInfo.md) | The applications info using by the client. | [optional] 
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

### NewClientTrafficWithApplications

`func NewClientTrafficWithApplications() *ClientTrafficWithApplications`

NewClientTrafficWithApplications instantiates a new ClientTrafficWithApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTrafficWithApplicationsWithDefaults

`func NewClientTrafficWithApplicationsWithDefaults() *ClientTrafficWithApplications`

NewClientTrafficWithApplicationsWithDefaults instantiates a new ClientTrafficWithApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationList

`func (o *ClientTrafficWithApplications) GetApplicationList() []ApplicationBasicInfo`

GetApplicationList returns the ApplicationList field if non-nil, zero value otherwise.

### GetApplicationListOk

`func (o *ClientTrafficWithApplications) GetApplicationListOk() (*[]ApplicationBasicInfo, bool)`

GetApplicationListOk returns a tuple with the ApplicationList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationList

`func (o *ClientTrafficWithApplications) SetApplicationList(v []ApplicationBasicInfo)`

SetApplicationList sets ApplicationList field to given value.

### HasApplicationList

`func (o *ClientTrafficWithApplications) HasApplicationList() bool`

HasApplicationList returns a boolean if a field has been set.

### GetClientName

`func (o *ClientTrafficWithApplications) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ClientTrafficWithApplications) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ClientTrafficWithApplications) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ClientTrafficWithApplications) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDownload

`func (o *ClientTrafficWithApplications) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ClientTrafficWithApplications) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ClientTrafficWithApplications) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ClientTrafficWithApplications) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetMac

`func (o *ClientTrafficWithApplications) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientTrafficWithApplications) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientTrafficWithApplications) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientTrafficWithApplications) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetManager

`func (o *ClientTrafficWithApplications) GetManager() bool`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *ClientTrafficWithApplications) GetManagerOk() (*bool, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *ClientTrafficWithApplications) SetManager(v bool)`

SetManager sets Manager field to given value.

### HasManager

`func (o *ClientTrafficWithApplications) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetTotalApplications

`func (o *ClientTrafficWithApplications) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *ClientTrafficWithApplications) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *ClientTrafficWithApplications) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.

### HasTotalApplications

`func (o *ClientTrafficWithApplications) HasTotalApplications() bool`

HasTotalApplications returns a boolean if a field has been set.

### GetTraffic

`func (o *ClientTrafficWithApplications) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *ClientTrafficWithApplications) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *ClientTrafficWithApplications) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *ClientTrafficWithApplications) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetTrafficPercent

`func (o *ClientTrafficWithApplications) GetTrafficPercent() float64`

GetTrafficPercent returns the TrafficPercent field if non-nil, zero value otherwise.

### GetTrafficPercentOk

`func (o *ClientTrafficWithApplications) GetTrafficPercentOk() (*float64, bool)`

GetTrafficPercentOk returns a tuple with the TrafficPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPercent

`func (o *ClientTrafficWithApplications) SetTrafficPercent(v float64)`

SetTrafficPercent sets TrafficPercent field to given value.

### HasTrafficPercent

`func (o *ClientTrafficWithApplications) HasTrafficPercent() bool`

HasTrafficPercent returns a boolean if a field has been set.

### GetType

`func (o *ClientTrafficWithApplications) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientTrafficWithApplications) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientTrafficWithApplications) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClientTrafficWithApplications) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpload

`func (o *ClientTrafficWithApplications) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ClientTrafficWithApplications) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ClientTrafficWithApplications) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ClientTrafficWithApplications) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


