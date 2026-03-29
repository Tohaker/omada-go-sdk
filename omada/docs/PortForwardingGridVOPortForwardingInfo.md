# PortForwardingGridVOPortForwardingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]PortForwardingInfo**](PortForwardingInfo.md) |  | [optional] 
**ExistVirtualWan** | Pointer to **bool** |  | [optional] 
**ExistWanIp** | Pointer to **bool** |  | [optional] 
**SupportVirtualWan** | Pointer to **bool** |  | [optional] 
**SupportWanIp** | Pointer to **bool** |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewPortForwardingGridVOPortForwardingInfo

`func NewPortForwardingGridVOPortForwardingInfo() *PortForwardingGridVOPortForwardingInfo`

NewPortForwardingGridVOPortForwardingInfo instantiates a new PortForwardingGridVOPortForwardingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortForwardingGridVOPortForwardingInfoWithDefaults

`func NewPortForwardingGridVOPortForwardingInfoWithDefaults() *PortForwardingGridVOPortForwardingInfo`

NewPortForwardingGridVOPortForwardingInfoWithDefaults instantiates a new PortForwardingGridVOPortForwardingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *PortForwardingGridVOPortForwardingInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PortForwardingGridVOPortForwardingInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PortForwardingGridVOPortForwardingInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *PortForwardingGridVOPortForwardingInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *PortForwardingGridVOPortForwardingInfo) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *PortForwardingGridVOPortForwardingInfo) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *PortForwardingGridVOPortForwardingInfo) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *PortForwardingGridVOPortForwardingInfo) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *PortForwardingGridVOPortForwardingInfo) GetData() []PortForwardingInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PortForwardingGridVOPortForwardingInfo) GetDataOk() (*[]PortForwardingInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PortForwardingGridVOPortForwardingInfo) SetData(v []PortForwardingInfo)`

SetData sets Data field to given value.

### HasData

`func (o *PortForwardingGridVOPortForwardingInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExistVirtualWan

`func (o *PortForwardingGridVOPortForwardingInfo) GetExistVirtualWan() bool`

GetExistVirtualWan returns the ExistVirtualWan field if non-nil, zero value otherwise.

### GetExistVirtualWanOk

`func (o *PortForwardingGridVOPortForwardingInfo) GetExistVirtualWanOk() (*bool, bool)`

GetExistVirtualWanOk returns a tuple with the ExistVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistVirtualWan

`func (o *PortForwardingGridVOPortForwardingInfo) SetExistVirtualWan(v bool)`

SetExistVirtualWan sets ExistVirtualWan field to given value.

### HasExistVirtualWan

`func (o *PortForwardingGridVOPortForwardingInfo) HasExistVirtualWan() bool`

HasExistVirtualWan returns a boolean if a field has been set.

### GetExistWanIp

`func (o *PortForwardingGridVOPortForwardingInfo) GetExistWanIp() bool`

GetExistWanIp returns the ExistWanIp field if non-nil, zero value otherwise.

### GetExistWanIpOk

`func (o *PortForwardingGridVOPortForwardingInfo) GetExistWanIpOk() (*bool, bool)`

GetExistWanIpOk returns a tuple with the ExistWanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistWanIp

`func (o *PortForwardingGridVOPortForwardingInfo) SetExistWanIp(v bool)`

SetExistWanIp sets ExistWanIp field to given value.

### HasExistWanIp

`func (o *PortForwardingGridVOPortForwardingInfo) HasExistWanIp() bool`

HasExistWanIp returns a boolean if a field has been set.

### GetSupportVirtualWan

`func (o *PortForwardingGridVOPortForwardingInfo) GetSupportVirtualWan() bool`

GetSupportVirtualWan returns the SupportVirtualWan field if non-nil, zero value otherwise.

### GetSupportVirtualWanOk

`func (o *PortForwardingGridVOPortForwardingInfo) GetSupportVirtualWanOk() (*bool, bool)`

GetSupportVirtualWanOk returns a tuple with the SupportVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVirtualWan

`func (o *PortForwardingGridVOPortForwardingInfo) SetSupportVirtualWan(v bool)`

SetSupportVirtualWan sets SupportVirtualWan field to given value.

### HasSupportVirtualWan

`func (o *PortForwardingGridVOPortForwardingInfo) HasSupportVirtualWan() bool`

HasSupportVirtualWan returns a boolean if a field has been set.

### GetSupportWanIp

`func (o *PortForwardingGridVOPortForwardingInfo) GetSupportWanIp() bool`

GetSupportWanIp returns the SupportWanIp field if non-nil, zero value otherwise.

### GetSupportWanIpOk

`func (o *PortForwardingGridVOPortForwardingInfo) GetSupportWanIpOk() (*bool, bool)`

GetSupportWanIpOk returns a tuple with the SupportWanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWanIp

`func (o *PortForwardingGridVOPortForwardingInfo) SetSupportWanIp(v bool)`

SetSupportWanIp sets SupportWanIp field to given value.

### HasSupportWanIp

`func (o *PortForwardingGridVOPortForwardingInfo) HasSupportWanIp() bool`

HasSupportWanIp returns a boolean if a field has been set.

### GetTotalRows

`func (o *PortForwardingGridVOPortForwardingInfo) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *PortForwardingGridVOPortForwardingInfo) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *PortForwardingGridVOPortForwardingInfo) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *PortForwardingGridVOPortForwardingInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


