# NetworkActivityVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityList** | Pointer to [**[]ActivityBaseVO**](ActivityBaseVO.md) |  | [optional] 
**ExistData** | Pointer to **bool** | mark whether the data exists | [optional] 
**TotalClientCount** | Pointer to **int32** | total number of client | [optional] 
**TotalDownload** | Pointer to **int64** | total rx traffic | [optional] 
**TotalTraffic** | Pointer to **int64** | total traffic | [optional] 
**TotalUpload** | Pointer to **int64** | total tx traffic | [optional] 
**TotalWiredCount** | Pointer to **int32** | total number of wired client | [optional] 
**TotalWirelessCount** | Pointer to **int32** | total number of wireless client | [optional] 

## Methods

### NewNetworkActivityVO

`func NewNetworkActivityVO() *NetworkActivityVO`

NewNetworkActivityVO instantiates a new NetworkActivityVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkActivityVOWithDefaults

`func NewNetworkActivityVOWithDefaults() *NetworkActivityVO`

NewNetworkActivityVOWithDefaults instantiates a new NetworkActivityVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityList

`func (o *NetworkActivityVO) GetActivityList() []ActivityBaseVO`

GetActivityList returns the ActivityList field if non-nil, zero value otherwise.

### GetActivityListOk

`func (o *NetworkActivityVO) GetActivityListOk() (*[]ActivityBaseVO, bool)`

GetActivityListOk returns a tuple with the ActivityList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityList

`func (o *NetworkActivityVO) SetActivityList(v []ActivityBaseVO)`

SetActivityList sets ActivityList field to given value.

### HasActivityList

`func (o *NetworkActivityVO) HasActivityList() bool`

HasActivityList returns a boolean if a field has been set.

### GetExistData

`func (o *NetworkActivityVO) GetExistData() bool`

GetExistData returns the ExistData field if non-nil, zero value otherwise.

### GetExistDataOk

`func (o *NetworkActivityVO) GetExistDataOk() (*bool, bool)`

GetExistDataOk returns a tuple with the ExistData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistData

`func (o *NetworkActivityVO) SetExistData(v bool)`

SetExistData sets ExistData field to given value.

### HasExistData

`func (o *NetworkActivityVO) HasExistData() bool`

HasExistData returns a boolean if a field has been set.

### GetTotalClientCount

`func (o *NetworkActivityVO) GetTotalClientCount() int32`

GetTotalClientCount returns the TotalClientCount field if non-nil, zero value otherwise.

### GetTotalClientCountOk

`func (o *NetworkActivityVO) GetTotalClientCountOk() (*int32, bool)`

GetTotalClientCountOk returns a tuple with the TotalClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClientCount

`func (o *NetworkActivityVO) SetTotalClientCount(v int32)`

SetTotalClientCount sets TotalClientCount field to given value.

### HasTotalClientCount

`func (o *NetworkActivityVO) HasTotalClientCount() bool`

HasTotalClientCount returns a boolean if a field has been set.

### GetTotalDownload

`func (o *NetworkActivityVO) GetTotalDownload() int64`

GetTotalDownload returns the TotalDownload field if non-nil, zero value otherwise.

### GetTotalDownloadOk

`func (o *NetworkActivityVO) GetTotalDownloadOk() (*int64, bool)`

GetTotalDownloadOk returns a tuple with the TotalDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDownload

`func (o *NetworkActivityVO) SetTotalDownload(v int64)`

SetTotalDownload sets TotalDownload field to given value.

### HasTotalDownload

`func (o *NetworkActivityVO) HasTotalDownload() bool`

HasTotalDownload returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *NetworkActivityVO) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *NetworkActivityVO) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *NetworkActivityVO) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *NetworkActivityVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.

### GetTotalUpload

`func (o *NetworkActivityVO) GetTotalUpload() int64`

GetTotalUpload returns the TotalUpload field if non-nil, zero value otherwise.

### GetTotalUploadOk

`func (o *NetworkActivityVO) GetTotalUploadOk() (*int64, bool)`

GetTotalUploadOk returns a tuple with the TotalUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUpload

`func (o *NetworkActivityVO) SetTotalUpload(v int64)`

SetTotalUpload sets TotalUpload field to given value.

### HasTotalUpload

`func (o *NetworkActivityVO) HasTotalUpload() bool`

HasTotalUpload returns a boolean if a field has been set.

### GetTotalWiredCount

`func (o *NetworkActivityVO) GetTotalWiredCount() int32`

GetTotalWiredCount returns the TotalWiredCount field if non-nil, zero value otherwise.

### GetTotalWiredCountOk

`func (o *NetworkActivityVO) GetTotalWiredCountOk() (*int32, bool)`

GetTotalWiredCountOk returns a tuple with the TotalWiredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWiredCount

`func (o *NetworkActivityVO) SetTotalWiredCount(v int32)`

SetTotalWiredCount sets TotalWiredCount field to given value.

### HasTotalWiredCount

`func (o *NetworkActivityVO) HasTotalWiredCount() bool`

HasTotalWiredCount returns a boolean if a field has been set.

### GetTotalWirelessCount

`func (o *NetworkActivityVO) GetTotalWirelessCount() int32`

GetTotalWirelessCount returns the TotalWirelessCount field if non-nil, zero value otherwise.

### GetTotalWirelessCountOk

`func (o *NetworkActivityVO) GetTotalWirelessCountOk() (*int32, bool)`

GetTotalWirelessCountOk returns a tuple with the TotalWirelessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWirelessCount

`func (o *NetworkActivityVO) SetTotalWirelessCount(v int32)`

SetTotalWirelessCount sets TotalWirelessCount field to given value.

### HasTotalWirelessCount

`func (o *NetworkActivityVO) HasTotalWirelessCount() bool`

HasTotalWirelessCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


