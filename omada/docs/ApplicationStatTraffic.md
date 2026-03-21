# ApplicationStatTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveUser** | Pointer to **int32** | The number of active users under this app. | [optional] 
**ApplicationDescription** | Pointer to **string** | The description of the application. | [optional] 
**ApplicationId** | Pointer to **int32** | The id of the application. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application. | [optional] 
**Block** | Pointer to **int64** | The number of times this app has been blocked. | [optional] 
**Clients** | Pointer to [**[]ClientTrafficWithApplications**](ClientTrafficWithApplications.md) | The clients info. | [optional] 
**ClientsCount** | Pointer to **int32** | The number of clients using the app. | [optional] 
**Download** | Pointer to **int64** | The download traffic used by the app. | [optional] 
**FamilyId** | Pointer to **int32** | The id of the family. | [optional] 
**FamilyName** | Pointer to **string** | The name of the family. | [optional] 
**Traffic** | Pointer to **int64** | The total amount of traffic used by the app. | [optional] 
**TrafficPercent** | Pointer to **float64** | The percentage of traffic used by the app. | [optional] 
**Upload** | Pointer to **int64** | The upload traffic used by the app. | [optional] 

## Methods

### NewApplicationStatTraffic

`func NewApplicationStatTraffic() *ApplicationStatTraffic`

NewApplicationStatTraffic instantiates a new ApplicationStatTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStatTrafficWithDefaults

`func NewApplicationStatTrafficWithDefaults() *ApplicationStatTraffic`

NewApplicationStatTrafficWithDefaults instantiates a new ApplicationStatTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveUser

`func (o *ApplicationStatTraffic) GetActiveUser() int32`

GetActiveUser returns the ActiveUser field if non-nil, zero value otherwise.

### GetActiveUserOk

`func (o *ApplicationStatTraffic) GetActiveUserOk() (*int32, bool)`

GetActiveUserOk returns a tuple with the ActiveUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUser

`func (o *ApplicationStatTraffic) SetActiveUser(v int32)`

SetActiveUser sets ActiveUser field to given value.

### HasActiveUser

`func (o *ApplicationStatTraffic) HasActiveUser() bool`

HasActiveUser returns a boolean if a field has been set.

### GetApplicationDescription

`func (o *ApplicationStatTraffic) GetApplicationDescription() string`

GetApplicationDescription returns the ApplicationDescription field if non-nil, zero value otherwise.

### GetApplicationDescriptionOk

`func (o *ApplicationStatTraffic) GetApplicationDescriptionOk() (*string, bool)`

GetApplicationDescriptionOk returns a tuple with the ApplicationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDescription

`func (o *ApplicationStatTraffic) SetApplicationDescription(v string)`

SetApplicationDescription sets ApplicationDescription field to given value.

### HasApplicationDescription

`func (o *ApplicationStatTraffic) HasApplicationDescription() bool`

HasApplicationDescription returns a boolean if a field has been set.

### GetApplicationId

`func (o *ApplicationStatTraffic) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationStatTraffic) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationStatTraffic) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationStatTraffic) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *ApplicationStatTraffic) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ApplicationStatTraffic) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ApplicationStatTraffic) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *ApplicationStatTraffic) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetBlock

`func (o *ApplicationStatTraffic) GetBlock() int64`

GetBlock returns the Block field if non-nil, zero value otherwise.

### GetBlockOk

`func (o *ApplicationStatTraffic) GetBlockOk() (*int64, bool)`

GetBlockOk returns a tuple with the Block field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlock

`func (o *ApplicationStatTraffic) SetBlock(v int64)`

SetBlock sets Block field to given value.

### HasBlock

`func (o *ApplicationStatTraffic) HasBlock() bool`

HasBlock returns a boolean if a field has been set.

### GetClients

`func (o *ApplicationStatTraffic) GetClients() []ClientTrafficWithApplications`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ApplicationStatTraffic) GetClientsOk() (*[]ClientTrafficWithApplications, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ApplicationStatTraffic) SetClients(v []ClientTrafficWithApplications)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ApplicationStatTraffic) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetClientsCount

`func (o *ApplicationStatTraffic) GetClientsCount() int32`

GetClientsCount returns the ClientsCount field if non-nil, zero value otherwise.

### GetClientsCountOk

`func (o *ApplicationStatTraffic) GetClientsCountOk() (*int32, bool)`

GetClientsCountOk returns a tuple with the ClientsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsCount

`func (o *ApplicationStatTraffic) SetClientsCount(v int32)`

SetClientsCount sets ClientsCount field to given value.

### HasClientsCount

`func (o *ApplicationStatTraffic) HasClientsCount() bool`

HasClientsCount returns a boolean if a field has been set.

### GetDownload

`func (o *ApplicationStatTraffic) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ApplicationStatTraffic) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ApplicationStatTraffic) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ApplicationStatTraffic) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetFamilyId

`func (o *ApplicationStatTraffic) GetFamilyId() int32`

GetFamilyId returns the FamilyId field if non-nil, zero value otherwise.

### GetFamilyIdOk

`func (o *ApplicationStatTraffic) GetFamilyIdOk() (*int32, bool)`

GetFamilyIdOk returns a tuple with the FamilyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyId

`func (o *ApplicationStatTraffic) SetFamilyId(v int32)`

SetFamilyId sets FamilyId field to given value.

### HasFamilyId

`func (o *ApplicationStatTraffic) HasFamilyId() bool`

HasFamilyId returns a boolean if a field has been set.

### GetFamilyName

`func (o *ApplicationStatTraffic) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *ApplicationStatTraffic) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *ApplicationStatTraffic) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.

### HasFamilyName

`func (o *ApplicationStatTraffic) HasFamilyName() bool`

HasFamilyName returns a boolean if a field has been set.

### GetTraffic

`func (o *ApplicationStatTraffic) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *ApplicationStatTraffic) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *ApplicationStatTraffic) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *ApplicationStatTraffic) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetTrafficPercent

`func (o *ApplicationStatTraffic) GetTrafficPercent() float64`

GetTrafficPercent returns the TrafficPercent field if non-nil, zero value otherwise.

### GetTrafficPercentOk

`func (o *ApplicationStatTraffic) GetTrafficPercentOk() (*float64, bool)`

GetTrafficPercentOk returns a tuple with the TrafficPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficPercent

`func (o *ApplicationStatTraffic) SetTrafficPercent(v float64)`

SetTrafficPercent sets TrafficPercent field to given value.

### HasTrafficPercent

`func (o *ApplicationStatTraffic) HasTrafficPercent() bool`

HasTrafficPercent returns a boolean if a field has been set.

### GetUpload

`func (o *ApplicationStatTraffic) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ApplicationStatTraffic) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ApplicationStatTraffic) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ApplicationStatTraffic) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


