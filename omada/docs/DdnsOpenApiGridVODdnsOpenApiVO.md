# DdnsOpenApiGridVODdnsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]DdnsOpenApiVO**](DdnsOpenApiVO.md) |  | [optional] 
**SupportByDsLiteAndMapE** | Pointer to **bool** | Whether this feature is supported for the DS-Lite or Map-E WAN connection types. | [optional] 
**SupportCustomDdns** | Pointer to **bool** | Whether Custom Service Provider is supported in Dynamic DNS. | [optional] 
**SupportCustomInterval** | Pointer to **bool** | Whether Custom Update Interval is supported in Dynamic DNS. | [optional] 
**SupportTpLinkddns** | Pointer to **int32** | Whether Dynamic DNS supports TP-Link as a service provider. 0: Not supported, 1: Supported and not configured, 2: Supported and configured. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewDdnsOpenApiGridVODdnsOpenApiVO

`func NewDdnsOpenApiGridVODdnsOpenApiVO() *DdnsOpenApiGridVODdnsOpenApiVO`

NewDdnsOpenApiGridVODdnsOpenApiVO instantiates a new DdnsOpenApiGridVODdnsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdnsOpenApiGridVODdnsOpenApiVOWithDefaults

`func NewDdnsOpenApiGridVODdnsOpenApiVOWithDefaults() *DdnsOpenApiGridVODdnsOpenApiVO`

NewDdnsOpenApiGridVODdnsOpenApiVOWithDefaults instantiates a new DdnsOpenApiGridVODdnsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetData() []DdnsOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetDataOk() (*[]DdnsOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) SetData(v []DdnsOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportByDsLiteAndMapE

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetSupportByDsLiteAndMapE() bool`

GetSupportByDsLiteAndMapE returns the SupportByDsLiteAndMapE field if non-nil, zero value otherwise.

### GetSupportByDsLiteAndMapEOk

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetSupportByDsLiteAndMapEOk() (*bool, bool)`

GetSupportByDsLiteAndMapEOk returns a tuple with the SupportByDsLiteAndMapE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportByDsLiteAndMapE

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) SetSupportByDsLiteAndMapE(v bool)`

SetSupportByDsLiteAndMapE sets SupportByDsLiteAndMapE field to given value.

### HasSupportByDsLiteAndMapE

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) HasSupportByDsLiteAndMapE() bool`

HasSupportByDsLiteAndMapE returns a boolean if a field has been set.

### GetSupportCustomDdns

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetSupportCustomDdns() bool`

GetSupportCustomDdns returns the SupportCustomDdns field if non-nil, zero value otherwise.

### GetSupportCustomDdnsOk

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetSupportCustomDdnsOk() (*bool, bool)`

GetSupportCustomDdnsOk returns a tuple with the SupportCustomDdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDdns

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) SetSupportCustomDdns(v bool)`

SetSupportCustomDdns sets SupportCustomDdns field to given value.

### HasSupportCustomDdns

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) HasSupportCustomDdns() bool`

HasSupportCustomDdns returns a boolean if a field has been set.

### GetSupportCustomInterval

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetSupportCustomInterval() bool`

GetSupportCustomInterval returns the SupportCustomInterval field if non-nil, zero value otherwise.

### GetSupportCustomIntervalOk

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetSupportCustomIntervalOk() (*bool, bool)`

GetSupportCustomIntervalOk returns a tuple with the SupportCustomInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomInterval

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) SetSupportCustomInterval(v bool)`

SetSupportCustomInterval sets SupportCustomInterval field to given value.

### HasSupportCustomInterval

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) HasSupportCustomInterval() bool`

HasSupportCustomInterval returns a boolean if a field has been set.

### GetSupportTpLinkddns

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetSupportTpLinkddns() int32`

GetSupportTpLinkddns returns the SupportTpLinkddns field if non-nil, zero value otherwise.

### GetSupportTpLinkddnsOk

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetSupportTpLinkddnsOk() (*int32, bool)`

GetSupportTpLinkddnsOk returns a tuple with the SupportTpLinkddns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportTpLinkddns

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) SetSupportTpLinkddns(v int32)`

SetSupportTpLinkddns sets SupportTpLinkddns field to given value.

### HasSupportTpLinkddns

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) HasSupportTpLinkddns() bool`

HasSupportTpLinkddns returns a boolean if a field has been set.

### GetTotalRows

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *DdnsOpenApiGridVODdnsOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


