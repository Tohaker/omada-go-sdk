# VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VirtualWanInfoOpenApiVO**](VirtualWanInfoOpenApiVO.md) |  | [optional] 
**NumReachLimit** | Pointer to **bool** | Whether the virtual WAN reaches number limit. | [optional] 
**SupportByDsLiteAndMapE** | Pointer to **bool** | Whether this feature is supported for the DS-Lite or Map-E WAN connection types. | [optional] 
**SupportMssClamping** | Pointer to **bool** | Whether the pppoe supports mss clamping. | [optional] 
**SupportPppoeMru** | Pointer to **bool** | Whether the virtual WAN supports configuring pppoe mru. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO

`func NewVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO() *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO`

NewVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO instantiates a new VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVOWithDefaults

`func NewVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVOWithDefaults() *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO`

NewVirtualWanGridOpenApiVOVirtualWanInfoOpenApiVOWithDefaults instantiates a new VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetData() []VirtualWanInfoOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetDataOk() (*[]VirtualWanInfoOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) SetData(v []VirtualWanInfoOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetNumReachLimit

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetNumReachLimit() bool`

GetNumReachLimit returns the NumReachLimit field if non-nil, zero value otherwise.

### GetNumReachLimitOk

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetNumReachLimitOk() (*bool, bool)`

GetNumReachLimitOk returns a tuple with the NumReachLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumReachLimit

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) SetNumReachLimit(v bool)`

SetNumReachLimit sets NumReachLimit field to given value.

### HasNumReachLimit

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) HasNumReachLimit() bool`

HasNumReachLimit returns a boolean if a field has been set.

### GetSupportByDsLiteAndMapE

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetSupportByDsLiteAndMapE() bool`

GetSupportByDsLiteAndMapE returns the SupportByDsLiteAndMapE field if non-nil, zero value otherwise.

### GetSupportByDsLiteAndMapEOk

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetSupportByDsLiteAndMapEOk() (*bool, bool)`

GetSupportByDsLiteAndMapEOk returns a tuple with the SupportByDsLiteAndMapE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportByDsLiteAndMapE

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) SetSupportByDsLiteAndMapE(v bool)`

SetSupportByDsLiteAndMapE sets SupportByDsLiteAndMapE field to given value.

### HasSupportByDsLiteAndMapE

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) HasSupportByDsLiteAndMapE() bool`

HasSupportByDsLiteAndMapE returns a boolean if a field has been set.

### GetSupportMssClamping

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetSupportMssClamping() bool`

GetSupportMssClamping returns the SupportMssClamping field if non-nil, zero value otherwise.

### GetSupportMssClampingOk

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetSupportMssClampingOk() (*bool, bool)`

GetSupportMssClampingOk returns a tuple with the SupportMssClamping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMssClamping

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) SetSupportMssClamping(v bool)`

SetSupportMssClamping sets SupportMssClamping field to given value.

### HasSupportMssClamping

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) HasSupportMssClamping() bool`

HasSupportMssClamping returns a boolean if a field has been set.

### GetSupportPppoeMru

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetSupportPppoeMru() bool`

GetSupportPppoeMru returns the SupportPppoeMru field if non-nil, zero value otherwise.

### GetSupportPppoeMruOk

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetSupportPppoeMruOk() (*bool, bool)`

GetSupportPppoeMruOk returns a tuple with the SupportPppoeMru field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPppoeMru

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) SetSupportPppoeMru(v bool)`

SetSupportPppoeMru sets SupportPppoeMru field to given value.

### HasSupportPppoeMru

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) HasSupportPppoeMru() bool`

HasSupportPppoeMru returns a boolean if a field has been set.

### GetTotalRows

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *VirtualWanGridOpenApiVOVirtualWanInfoOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


