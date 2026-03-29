# CategoryTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]ApplicationTrafficWithClientCount**](ApplicationTrafficWithClientCount.md) | The applications used in this category. | [optional] 
**Clients** | Pointer to [**[]ClientTrafficWithApplications**](ClientTrafficWithApplications.md) | The clients using the app in this category. | [optional] 
**Download** | Pointer to **int64** | The download traffic used by the category. | [optional] 
**FamilyId** | Pointer to **int32** | The id of the category. | [optional] 
**FamilyName** | Pointer to **string** | The name of the category. | [optional] 
**TotalApplications** | Pointer to **int32** | The number of applications used in this category. | [optional] 
**TotalClients** | Pointer to **int32** | The number of clients using the app in this category. | [optional] 
**Traffic** | Pointer to **int64** | The total amount of traffic used by the category. | [optional] 
**TrafficPercent** | Pointer to **float64** | The percentage of traffic used by the category. | [optional] 
**Upload** | Pointer to **int64** | The upload traffic used by the category. | [optional] 

## Methods

### NewCategoryTraffic

`func NewCategoryTraffic() *CategoryTraffic`

NewCategoryTraffic instantiates a new CategoryTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryTrafficWithDefaults

`func NewCategoryTrafficWithDefaults() *CategoryTraffic`

NewCategoryTrafficWithDefaults instantiates a new CategoryTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *CategoryTraffic) GetApplications() []ApplicationTrafficWithClientCount`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *CategoryTraffic) GetApplicationsOk() (*[]ApplicationTrafficWithClientCount, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *CategoryTraffic) SetApplications(v []ApplicationTrafficWithClientCount)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *CategoryTraffic) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetClients

`func (o *CategoryTraffic) GetClients() []ClientTrafficWithApplications`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *CategoryTraffic) GetClientsOk() (*[]ClientTrafficWithApplications, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *CategoryTraffic) SetClients(v []ClientTrafficWithApplications)`

SetClients sets Clients field to given value.

### HasClients

`func (o *CategoryTraffic) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetDownload

`func (o *CategoryTraffic) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *CategoryTraffic) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *CategoryTraffic) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *CategoryTraffic) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetFamilyId

`func (o *CategoryTraffic) GetFamilyId() int32`

GetFamilyId returns the FamilyId field if non-nil, zero value otherwise.

### GetFamilyIdOk

`func (o *CategoryTraffic) GetFamilyIdOk() (*int32, bool)`

GetFamilyIdOk returns a tuple with the FamilyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyId

`func (o *CategoryTraffic) SetFamilyId(v int32)`

SetFamilyId sets FamilyId field to given value.

### HasFamilyId

`func (o *CategoryTraffic) HasFamilyId() bool`

HasFamilyId returns a boolean if a field has been set.

### GetFamilyName

`func (o *CategoryTraffic) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *CategoryTraffic) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *CategoryTraffic) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *CategoryTraffic) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetTotalApplications

`func (o *CategoryTraffic) GetTotalApplications() int32`

GetTotalApplications returns the TotalApplications field if non-nil, zero value otherwise.

### GetTotalApplicationsOk

`func (o *CategoryTraffic) GetTotalApplicationsOk() (*int32, bool)`

GetTotalApplicationsOk returns a tuple with the TotalApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApplications

`func (o *CategoryTraffic) SetTotalApplications(v int32)`

SetTotalApplications sets TotalApplications field to given value.

### HasTotalApplications

`func (o *CategoryTraffic) HasTotalApplications() bool`

HasTotalApplications returns a boolean if a field has been set.

### GetTotalClients

`func (o *CategoryTraffic) GetTotalClients() int32`

GetTotalClients returns the TotalClients field if non-nil, zero value otherwise.

### GetTotalClientsOk

`func (o *CategoryTraffic) GetTotalClientsOk() (*int32, bool)`

GetTotalClientsOk returns a tuple with the TotalClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClients

`func (o *CategoryTraffic) SetTotalClients(v int32)`

SetTotalClients sets TotalClients field to given value.

### HasTotalClients

`func (o *CategoryTraffic) HasTotalClients() bool`

HasTotalClients returns a boolean if a field has been set.

### GetTraffic

`func (o *CategoryTraffic) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *CategoryTraffic) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *CategoryTraffic) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *CategoryTraffic) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetTrafficPercent

`func (o *CategoryTraffic) GetTrafficPercent() float64`

GetTrafficPercent returns the TrafficPercent field if non-nil, zero value otherwise.

### GetTrafficPercentOk

`func (o *CategoryTraffic) GetTrafficPercentOk() (*float64, bool)`

GetTrafficPercentOk returns a tuple with the TrafficPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPercent

`func (o *CategoryTraffic) SetTrafficPercent(v float64)`

SetTrafficPercent sets TrafficPercent field to given value.

### HasTrafficPercent

`func (o *CategoryTraffic) HasTrafficPercent() bool`

HasTrafficPercent returns a boolean if a field has been set.

### GetUpload

`func (o *CategoryTraffic) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *CategoryTraffic) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *CategoryTraffic) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *CategoryTraffic) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


