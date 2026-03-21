# AdoptedDeviceGridVODeviceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]DeviceInfo**](DeviceInfo.md) |  | [optional] 
**MlagData** | Pointer to [**[]OswMlagDataVODeviceInfo**](OswMlagDataVODeviceInfo.md) |  | [optional] 
**P2pGroupTotalRows** | Pointer to **int64** |  | [optional] 
**StackData** | Pointer to [**[]OswStackDataVODeviceInfo**](OswStackDataVODeviceInfo.md) |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewAdoptedDeviceGridVODeviceInfo

`func NewAdoptedDeviceGridVODeviceInfo() *AdoptedDeviceGridVODeviceInfo`

NewAdoptedDeviceGridVODeviceInfo instantiates a new AdoptedDeviceGridVODeviceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdoptedDeviceGridVODeviceInfoWithDefaults

`func NewAdoptedDeviceGridVODeviceInfoWithDefaults() *AdoptedDeviceGridVODeviceInfo`

NewAdoptedDeviceGridVODeviceInfoWithDefaults instantiates a new AdoptedDeviceGridVODeviceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *AdoptedDeviceGridVODeviceInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *AdoptedDeviceGridVODeviceInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *AdoptedDeviceGridVODeviceInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *AdoptedDeviceGridVODeviceInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *AdoptedDeviceGridVODeviceInfo) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *AdoptedDeviceGridVODeviceInfo) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *AdoptedDeviceGridVODeviceInfo) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *AdoptedDeviceGridVODeviceInfo) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *AdoptedDeviceGridVODeviceInfo) GetData() []DeviceInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdoptedDeviceGridVODeviceInfo) GetDataOk() (*[]DeviceInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdoptedDeviceGridVODeviceInfo) SetData(v []DeviceInfo)`

SetData sets Data field to given value.

### HasData

`func (o *AdoptedDeviceGridVODeviceInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMlagData

`func (o *AdoptedDeviceGridVODeviceInfo) GetMlagData() []OswMlagDataVODeviceInfo`

GetMlagData returns the MlagData field if non-nil, zero value otherwise.

### GetMlagDataOk

`func (o *AdoptedDeviceGridVODeviceInfo) GetMlagDataOk() (*[]OswMlagDataVODeviceInfo, bool)`

GetMlagDataOk returns a tuple with the MlagData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagData

`func (o *AdoptedDeviceGridVODeviceInfo) SetMlagData(v []OswMlagDataVODeviceInfo)`

SetMlagData sets MlagData field to given value.

### HasMlagData

`func (o *AdoptedDeviceGridVODeviceInfo) HasMlagData() bool`

HasMlagData returns a boolean if a field has been set.

### GetP2pGroupTotalRows

`func (o *AdoptedDeviceGridVODeviceInfo) GetP2pGroupTotalRows() int64`

GetP2pGroupTotalRows returns the P2pGroupTotalRows field if non-nil, zero value otherwise.

### GetP2pGroupTotalRowsOk

`func (o *AdoptedDeviceGridVODeviceInfo) GetP2pGroupTotalRowsOk() (*int64, bool)`

GetP2pGroupTotalRowsOk returns a tuple with the P2pGroupTotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetP2pGroupTotalRows

`func (o *AdoptedDeviceGridVODeviceInfo) SetP2pGroupTotalRows(v int64)`

SetP2pGroupTotalRows sets P2pGroupTotalRows field to given value.

### HasP2pGroupTotalRows

`func (o *AdoptedDeviceGridVODeviceInfo) HasP2pGroupTotalRows() bool`

HasP2pGroupTotalRows returns a boolean if a field has been set.

### GetStackData

`func (o *AdoptedDeviceGridVODeviceInfo) GetStackData() []OswStackDataVODeviceInfo`

GetStackData returns the StackData field if non-nil, zero value otherwise.

### GetStackDataOk

`func (o *AdoptedDeviceGridVODeviceInfo) GetStackDataOk() (*[]OswStackDataVODeviceInfo, bool)`

GetStackDataOk returns a tuple with the StackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackData

`func (o *AdoptedDeviceGridVODeviceInfo) SetStackData(v []OswStackDataVODeviceInfo)`

SetStackData sets StackData field to given value.

### HasStackData

`func (o *AdoptedDeviceGridVODeviceInfo) HasStackData() bool`

HasStackData returns a boolean if a field has been set.

### GetTotalRows

`func (o *AdoptedDeviceGridVODeviceInfo) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *AdoptedDeviceGridVODeviceInfo) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *AdoptedDeviceGridVODeviceInfo) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *AdoptedDeviceGridVODeviceInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


