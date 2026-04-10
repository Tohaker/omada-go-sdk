# StaticRoutingOpenApiGridVOStaticRoutingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]StaticRoutingInfo**](StaticRoutingInfo.md) |  | [optional] 
**SupportByDsLiteAndMapE** | Pointer to **bool** | Whether this feature is supported for the DS-Lite or Map-E WAN connection types. | [optional] 
**SupportVirtualWan** | Pointer to **bool** | Whether Virtual Wan is supported in Static Route. | [optional] 
**SupportVpnClient** | Pointer to **bool** | Whether VPN Client is supported in Static Route. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewStaticRoutingOpenApiGridVOStaticRoutingInfo

`func NewStaticRoutingOpenApiGridVOStaticRoutingInfo() *StaticRoutingOpenApiGridVOStaticRoutingInfo`

NewStaticRoutingOpenApiGridVOStaticRoutingInfo instantiates a new StaticRoutingOpenApiGridVOStaticRoutingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRoutingOpenApiGridVOStaticRoutingInfoWithDefaults

`func NewStaticRoutingOpenApiGridVOStaticRoutingInfoWithDefaults() *StaticRoutingOpenApiGridVOStaticRoutingInfo`

NewStaticRoutingOpenApiGridVOStaticRoutingInfoWithDefaults instantiates a new StaticRoutingOpenApiGridVOStaticRoutingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetData() []StaticRoutingInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetDataOk() (*[]StaticRoutingInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) SetData(v []StaticRoutingInfo)`

SetData sets Data field to given value.

### HasData

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportByDsLiteAndMapE

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetSupportByDsLiteAndMapE() bool`

GetSupportByDsLiteAndMapE returns the SupportByDsLiteAndMapE field if non-nil, zero value otherwise.

### GetSupportByDsLiteAndMapEOk

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetSupportByDsLiteAndMapEOk() (*bool, bool)`

GetSupportByDsLiteAndMapEOk returns a tuple with the SupportByDsLiteAndMapE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportByDsLiteAndMapE

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) SetSupportByDsLiteAndMapE(v bool)`

SetSupportByDsLiteAndMapE sets SupportByDsLiteAndMapE field to given value.

### HasSupportByDsLiteAndMapE

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) HasSupportByDsLiteAndMapE() bool`

HasSupportByDsLiteAndMapE returns a boolean if a field has been set.

### GetSupportVirtualWan

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetSupportVirtualWan() bool`

GetSupportVirtualWan returns the SupportVirtualWan field if non-nil, zero value otherwise.

### GetSupportVirtualWanOk

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetSupportVirtualWanOk() (*bool, bool)`

GetSupportVirtualWanOk returns a tuple with the SupportVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVirtualWan

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) SetSupportVirtualWan(v bool)`

SetSupportVirtualWan sets SupportVirtualWan field to given value.

### HasSupportVirtualWan

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) HasSupportVirtualWan() bool`

HasSupportVirtualWan returns a boolean if a field has been set.

### GetSupportVpnClient

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetSupportVpnClient() bool`

GetSupportVpnClient returns the SupportVpnClient field if non-nil, zero value otherwise.

### GetSupportVpnClientOk

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetSupportVpnClientOk() (*bool, bool)`

GetSupportVpnClientOk returns a tuple with the SupportVpnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVpnClient

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) SetSupportVpnClient(v bool)`

SetSupportVpnClient sets SupportVpnClient field to given value.

### HasSupportVpnClient

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) HasSupportVpnClient() bool`

HasSupportVpnClient returns a boolean if a field has been set.

### GetTotalRows

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *StaticRoutingOpenApiGridVOStaticRoutingInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


