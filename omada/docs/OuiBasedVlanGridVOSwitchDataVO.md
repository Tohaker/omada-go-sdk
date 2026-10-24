# OuiBasedVlanGridVOSwitchDataVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]SwitchDataVO**](SwitchDataVO.md) |  | [optional] 
**SupportAutoAddVlanRows** | Pointer to **int64** | Total rows of all items support auto add vlan. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 
**UnSupportAutoAddVlanRows** | Pointer to **int64** | Total rows of all items does not support auto add vlan. | [optional] 

## Methods

### NewOuiBasedVlanGridVOSwitchDataVO

`func NewOuiBasedVlanGridVOSwitchDataVO() *OuiBasedVlanGridVOSwitchDataVO`

NewOuiBasedVlanGridVOSwitchDataVO instantiates a new OuiBasedVlanGridVOSwitchDataVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanGridVOSwitchDataVOWithDefaults

`func NewOuiBasedVlanGridVOSwitchDataVOWithDefaults() *OuiBasedVlanGridVOSwitchDataVO`

NewOuiBasedVlanGridVOSwitchDataVOWithDefaults instantiates a new OuiBasedVlanGridVOSwitchDataVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *OuiBasedVlanGridVOSwitchDataVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *OuiBasedVlanGridVOSwitchDataVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *OuiBasedVlanGridVOSwitchDataVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *OuiBasedVlanGridVOSwitchDataVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetData() []SwitchDataVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetDataOk() (*[]SwitchDataVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OuiBasedVlanGridVOSwitchDataVO) SetData(v []SwitchDataVO)`

SetData sets Data field to given value.

### HasData

`func (o *OuiBasedVlanGridVOSwitchDataVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportAutoAddVlanRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetSupportAutoAddVlanRows() int64`

GetSupportAutoAddVlanRows returns the SupportAutoAddVlanRows field if non-nil, zero value otherwise.

### GetSupportAutoAddVlanRowsOk

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetSupportAutoAddVlanRowsOk() (*int64, bool)`

GetSupportAutoAddVlanRowsOk returns a tuple with the SupportAutoAddVlanRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAutoAddVlanRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) SetSupportAutoAddVlanRows(v int64)`

SetSupportAutoAddVlanRows sets SupportAutoAddVlanRows field to given value.

### HasSupportAutoAddVlanRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) HasSupportAutoAddVlanRows() bool`

HasSupportAutoAddVlanRows returns a boolean if a field has been set.

### GetTotalRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.

### GetUnSupportAutoAddVlanRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetUnSupportAutoAddVlanRows() int64`

GetUnSupportAutoAddVlanRows returns the UnSupportAutoAddVlanRows field if non-nil, zero value otherwise.

### GetUnSupportAutoAddVlanRowsOk

`func (o *OuiBasedVlanGridVOSwitchDataVO) GetUnSupportAutoAddVlanRowsOk() (*int64, bool)`

GetUnSupportAutoAddVlanRowsOk returns a tuple with the UnSupportAutoAddVlanRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnSupportAutoAddVlanRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) SetUnSupportAutoAddVlanRows(v int64)`

SetUnSupportAutoAddVlanRows sets UnSupportAutoAddVlanRows field to given value.

### HasUnSupportAutoAddVlanRows

`func (o *OuiBasedVlanGridVOSwitchDataVO) HasUnSupportAutoAddVlanRows() bool`

HasUnSupportAutoAddVlanRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


