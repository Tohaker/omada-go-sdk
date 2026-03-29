# LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]LanNetworkSplitOpenApiVO**](LanNetworkSplitOpenApiVO.md) |  | [optional] 
**DhcpRangePoolSize** | Pointer to **int32** | The size of DHCP range pool supported by the lan network DHCP. | [optional] 
**SupportArpDetection** | Pointer to **bool** | Whether ARP Detection is supported of the lan network. | [optional] 
**SupportCustomDhcpOption** | Pointer to **bool** | Whether custom DHCP option configuration is supported of the lan netowrk. | [optional] 
**SupportDhcpNextServer** | Pointer to **bool** | Whether DHCP Next Server is supported of the lan network. | [optional] 
**SupportMultiVlan** | Pointer to **bool** | Whether multi vlan configuration is supported of the lan netowrk. | [optional] 
**SupportRA** | Pointer to **bool** | Whether Router Advertisement configuration is supported of the lan netowrk. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO

`func NewLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO() *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO`

NewLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO instantiates a new LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVOWithDefaults

`func NewLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVOWithDefaults() *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO`

NewLanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVOWithDefaults instantiates a new LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetData() []LanNetworkSplitOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetDataOk() (*[]LanNetworkSplitOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetData(v []LanNetworkSplitOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDhcpRangePoolSize

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetDhcpRangePoolSize() int32`

GetDhcpRangePoolSize returns the DhcpRangePoolSize field if non-nil, zero value otherwise.

### GetDhcpRangePoolSizeOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetDhcpRangePoolSizeOk() (*int32, bool)`

GetDhcpRangePoolSizeOk returns a tuple with the DhcpRangePoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRangePoolSize

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetDhcpRangePoolSize(v int32)`

SetDhcpRangePoolSize sets DhcpRangePoolSize field to given value.

### HasDhcpRangePoolSize

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasDhcpRangePoolSize() bool`

HasDhcpRangePoolSize returns a boolean if a field has been set.

### GetSupportArpDetection

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportArpDetection() bool`

GetSupportArpDetection returns the SupportArpDetection field if non-nil, zero value otherwise.

### GetSupportArpDetectionOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportArpDetectionOk() (*bool, bool)`

GetSupportArpDetectionOk returns a tuple with the SupportArpDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportArpDetection

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetSupportArpDetection(v bool)`

SetSupportArpDetection sets SupportArpDetection field to given value.

### HasSupportArpDetection

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasSupportArpDetection() bool`

HasSupportArpDetection returns a boolean if a field has been set.

### GetSupportCustomDhcpOption

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportCustomDhcpOption() bool`

GetSupportCustomDhcpOption returns the SupportCustomDhcpOption field if non-nil, zero value otherwise.

### GetSupportCustomDhcpOptionOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportCustomDhcpOptionOk() (*bool, bool)`

GetSupportCustomDhcpOptionOk returns a tuple with the SupportCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDhcpOption

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetSupportCustomDhcpOption(v bool)`

SetSupportCustomDhcpOption sets SupportCustomDhcpOption field to given value.

### HasSupportCustomDhcpOption

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasSupportCustomDhcpOption() bool`

HasSupportCustomDhcpOption returns a boolean if a field has been set.

### GetSupportDhcpNextServer

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportDhcpNextServer() bool`

GetSupportDhcpNextServer returns the SupportDhcpNextServer field if non-nil, zero value otherwise.

### GetSupportDhcpNextServerOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportDhcpNextServerOk() (*bool, bool)`

GetSupportDhcpNextServerOk returns a tuple with the SupportDhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpNextServer

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetSupportDhcpNextServer(v bool)`

SetSupportDhcpNextServer sets SupportDhcpNextServer field to given value.

### HasSupportDhcpNextServer

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasSupportDhcpNextServer() bool`

HasSupportDhcpNextServer returns a boolean if a field has been set.

### GetSupportMultiVlan

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportMultiVlan() bool`

GetSupportMultiVlan returns the SupportMultiVlan field if non-nil, zero value otherwise.

### GetSupportMultiVlanOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportMultiVlanOk() (*bool, bool)`

GetSupportMultiVlanOk returns a tuple with the SupportMultiVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMultiVlan

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetSupportMultiVlan(v bool)`

SetSupportMultiVlan sets SupportMultiVlan field to given value.

### HasSupportMultiVlan

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasSupportMultiVlan() bool`

HasSupportMultiVlan returns a boolean if a field has been set.

### GetSupportRA

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportRA() bool`

GetSupportRA returns the SupportRA field if non-nil, zero value otherwise.

### GetSupportRAOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetSupportRAOk() (*bool, bool)`

GetSupportRAOk returns a tuple with the SupportRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRA

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetSupportRA(v bool)`

SetSupportRA sets SupportRA field to given value.

### HasSupportRA

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasSupportRA() bool`

HasSupportRA returns a boolean if a field has been set.

### GetTotalRows

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *LanNetworkOpenApiV2GridVOLanNetworkSplitOpenApiVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


