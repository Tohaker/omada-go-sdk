# PortForwardingOpenApiGridVOPortForwardingInfoTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]PortForwardingInfoTemplate**](PortForwardingInfoTemplate.md) |  | [optional] 
**SupportByDsLiteAndMapE** | Pointer to **bool** | Whether this feature is supported for the DS-Lite or Map-E WAN connection types. | [optional] 
**SupportVirtualWan** | Pointer to **bool** | Whether Virtual Wan configuration is supported of port forwarding. | [optional] 
**SupportWanIp** | Pointer to **bool** | Whether Wan Alias Ip configuration is supported of port forwarding. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewPortForwardingOpenApiGridVOPortForwardingInfoTemplate

`func NewPortForwardingOpenApiGridVOPortForwardingInfoTemplate() *PortForwardingOpenApiGridVOPortForwardingInfoTemplate`

NewPortForwardingOpenApiGridVOPortForwardingInfoTemplate instantiates a new PortForwardingOpenApiGridVOPortForwardingInfoTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortForwardingOpenApiGridVOPortForwardingInfoTemplateWithDefaults

`func NewPortForwardingOpenApiGridVOPortForwardingInfoTemplateWithDefaults() *PortForwardingOpenApiGridVOPortForwardingInfoTemplate`

NewPortForwardingOpenApiGridVOPortForwardingInfoTemplateWithDefaults instantiates a new PortForwardingOpenApiGridVOPortForwardingInfoTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetData() []PortForwardingInfoTemplate`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetDataOk() (*[]PortForwardingInfoTemplate, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) SetData(v []PortForwardingInfoTemplate)`

SetData sets Data field to given value.

### HasData

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportByDsLiteAndMapE

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetSupportByDsLiteAndMapE() bool`

GetSupportByDsLiteAndMapE returns the SupportByDsLiteAndMapE field if non-nil, zero value otherwise.

### GetSupportByDsLiteAndMapEOk

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetSupportByDsLiteAndMapEOk() (*bool, bool)`

GetSupportByDsLiteAndMapEOk returns a tuple with the SupportByDsLiteAndMapE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportByDsLiteAndMapE

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) SetSupportByDsLiteAndMapE(v bool)`

SetSupportByDsLiteAndMapE sets SupportByDsLiteAndMapE field to given value.

### HasSupportByDsLiteAndMapE

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) HasSupportByDsLiteAndMapE() bool`

HasSupportByDsLiteAndMapE returns a boolean if a field has been set.

### GetSupportVirtualWan

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetSupportVirtualWan() bool`

GetSupportVirtualWan returns the SupportVirtualWan field if non-nil, zero value otherwise.

### GetSupportVirtualWanOk

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetSupportVirtualWanOk() (*bool, bool)`

GetSupportVirtualWanOk returns a tuple with the SupportVirtualWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVirtualWan

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) SetSupportVirtualWan(v bool)`

SetSupportVirtualWan sets SupportVirtualWan field to given value.

### HasSupportVirtualWan

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) HasSupportVirtualWan() bool`

HasSupportVirtualWan returns a boolean if a field has been set.

### GetSupportWanIp

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetSupportWanIp() bool`

GetSupportWanIp returns the SupportWanIp field if non-nil, zero value otherwise.

### GetSupportWanIpOk

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetSupportWanIpOk() (*bool, bool)`

GetSupportWanIpOk returns a tuple with the SupportWanIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWanIp

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) SetSupportWanIp(v bool)`

SetSupportWanIp sets SupportWanIp field to given value.

### HasSupportWanIp

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) HasSupportWanIp() bool`

HasSupportWanIp returns a boolean if a field has been set.

### GetTotalRows

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *PortForwardingOpenApiGridVOPortForwardingInfoTemplate) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


