# WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]WifiCallingTrafficOpenApiVO**](WifiCallingTrafficOpenApiVO.md) |  | [optional] 
**TopKEPDGs** | Pointer to [**[]WifiCallingTrafficOpenApiVO**](WifiCallingTrafficOpenApiVO.md) | Top k ePDGs based on voice call traffic statistics. | [optional] 
**TopKSsids** | Pointer to [**[]WifiCallingTrafficOpenApiVO**](WifiCallingTrafficOpenApiVO.md) | Top k SSIDs based on voice call traffic statistics. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO

`func NewWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO() *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO`

NewWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO instantiates a new WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVOWithDefaults

`func NewWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVOWithDefaults() *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO`

NewWifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVOWithDefaults instantiates a new WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetData() []WifiCallingTrafficOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetDataOk() (*[]WifiCallingTrafficOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) SetData(v []WifiCallingTrafficOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTopKEPDGs

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetTopKEPDGs() []WifiCallingTrafficOpenApiVO`

GetTopKEPDGs returns the TopKEPDGs field if non-nil, zero value otherwise.

### GetTopKEPDGsOk

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetTopKEPDGsOk() (*[]WifiCallingTrafficOpenApiVO, bool)`

GetTopKEPDGsOk returns a tuple with the TopKEPDGs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopKEPDGs

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) SetTopKEPDGs(v []WifiCallingTrafficOpenApiVO)`

SetTopKEPDGs sets TopKEPDGs field to given value.

### HasTopKEPDGs

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) HasTopKEPDGs() bool`

HasTopKEPDGs returns a boolean if a field has been set.

### GetTopKSsids

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetTopKSsids() []WifiCallingTrafficOpenApiVO`

GetTopKSsids returns the TopKSsids field if non-nil, zero value otherwise.

### GetTopKSsidsOk

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetTopKSsidsOk() (*[]WifiCallingTrafficOpenApiVO, bool)`

GetTopKSsidsOk returns a tuple with the TopKSsids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopKSsids

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) SetTopKSsids(v []WifiCallingTrafficOpenApiVO)`

SetTopKSsids sets TopKSsids field to given value.

### HasTopKSsids

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) HasTopKSsids() bool`

HasTopKSsids returns a boolean if a field has been set.

### GetTotalRows

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *WifiCallingTrafficGridOpenApiVOWifiCallingTrafficOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


