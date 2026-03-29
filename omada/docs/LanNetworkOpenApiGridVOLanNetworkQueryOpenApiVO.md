# LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]LanNetworkQueryOpenApiVO**](LanNetworkQueryOpenApiVO.md) |  | [optional] 
**SupportCustomDhcpOption** | Pointer to **bool** | Whether custom DHCP option configuration is supported of the lan netowrk. | [optional] 
**SupportMultiVlan** | Pointer to **bool** | Whether multi vlan configuration is supported of the lan netowrk. | [optional] 
**SupportRA** | Pointer to **bool** | Whether Router Advertisement configuration is supported of the lan netowrk. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO

`func NewLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO() *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO`

NewLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO instantiates a new LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVOWithDefaults

`func NewLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVOWithDefaults() *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO`

NewLanNetworkOpenApiGridVOLanNetworkQueryOpenApiVOWithDefaults instantiates a new LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetData() []LanNetworkQueryOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetDataOk() (*[]LanNetworkQueryOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) SetData(v []LanNetworkQueryOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportCustomDhcpOption

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetSupportCustomDhcpOption() bool`

GetSupportCustomDhcpOption returns the SupportCustomDhcpOption field if non-nil, zero value otherwise.

### GetSupportCustomDhcpOptionOk

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetSupportCustomDhcpOptionOk() (*bool, bool)`

GetSupportCustomDhcpOptionOk returns a tuple with the SupportCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDhcpOption

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) SetSupportCustomDhcpOption(v bool)`

SetSupportCustomDhcpOption sets SupportCustomDhcpOption field to given value.

### HasSupportCustomDhcpOption

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) HasSupportCustomDhcpOption() bool`

HasSupportCustomDhcpOption returns a boolean if a field has been set.

### GetSupportMultiVlan

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetSupportMultiVlan() bool`

GetSupportMultiVlan returns the SupportMultiVlan field if non-nil, zero value otherwise.

### GetSupportMultiVlanOk

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetSupportMultiVlanOk() (*bool, bool)`

GetSupportMultiVlanOk returns a tuple with the SupportMultiVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMultiVlan

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) SetSupportMultiVlan(v bool)`

SetSupportMultiVlan sets SupportMultiVlan field to given value.

### HasSupportMultiVlan

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) HasSupportMultiVlan() bool`

HasSupportMultiVlan returns a boolean if a field has been set.

### GetSupportRA

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetSupportRA() bool`

GetSupportRA returns the SupportRA field if non-nil, zero value otherwise.

### GetSupportRAOk

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetSupportRAOk() (*bool, bool)`

GetSupportRAOk returns a tuple with the SupportRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRA

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) SetSupportRA(v bool)`

SetSupportRA sets SupportRA field to given value.

### HasSupportRA

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) HasSupportRA() bool`

HasSupportRA returns a boolean if a field has been set.

### GetTotalRows

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *LanNetworkOpenApiGridVOLanNetworkQueryOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


