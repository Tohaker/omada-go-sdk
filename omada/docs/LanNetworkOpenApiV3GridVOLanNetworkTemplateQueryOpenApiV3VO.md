# LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CombinedGateway** | Pointer to **bool** | Whether the current site is combined gateway | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]LanNetworkTemplateQueryOpenApiV3VO**](LanNetworkTemplateQueryOpenApiV3VO.md) |  | [optional] 
**DhcpRangePoolSize** | Pointer to **int32** | The size of DHCP range pool supported by the lan network DHCP. | [optional] 
**InterfaceNum** | Pointer to **int32** | The number of interface supported by current gateway. | [optional] 
**OsgSupportIgmpSnooping** | Pointer to **[]string** | The name of the model that supports IgmpSnooping under site. | [optional] 
**SupportArpDetection** | Pointer to **bool** | Whether ARP Detection is supported of the lan network. | [optional] 
**SupportCustomDhcpOption** | Pointer to **bool** | Whether custom DHCP option configuration is supported of the lan netowrk. | [optional] 
**SupportDefault** | Pointer to **bool** | Whether to allow the selection of Default VLAN. | [optional] 
**SupportDhcpNextServer** | Pointer to **bool** | Whether DHCP Next Server is supported of the lan network. | [optional] 
**SupportLanIpv6** | Pointer to **bool** | Whether it supports lan Ipv6. | [optional] 
**SupportLanIpv6PassThrough** | Pointer to **bool** | Whether it supports Pass Through. | [optional] 
**SupportMaxVlanNum** | Pointer to **int32** | The number of vlan supported by current gateway. | [optional] 
**SupportMultiVlan** | Pointer to **bool** | Whether multi vlan configuration is supported of the lan netowrk. | [optional] 
**SupportNetworkIsolation** | Pointer to **bool** | Whether it supports isolate network. | [optional] 
**SupportPdOnDhcp** | Pointer to **bool** | Whether it supports \&quot;Get from Prefix Delegation\&quot; when disable Prefix Delegation on the corresponding WAN interface page. | [optional] 
**SupportRA** | Pointer to **bool** | Whether Router Advertisement configuration is supported of the lan netowrk. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 
**VlanNums** | Pointer to **int32** | The vlan num of the site. | [optional] 

## Methods

### NewLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO

`func NewLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO() *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO`

NewLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO instantiates a new LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VOWithDefaults

`func NewLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VOWithDefaults() *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO`

NewLanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VOWithDefaults instantiates a new LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinedGateway

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetCombinedGateway() bool`

GetCombinedGateway returns the CombinedGateway field if non-nil, zero value otherwise.

### GetCombinedGatewayOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetCombinedGatewayOk() (*bool, bool)`

GetCombinedGatewayOk returns a tuple with the CombinedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinedGateway

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetCombinedGateway(v bool)`

SetCombinedGateway sets CombinedGateway field to given value.

### HasCombinedGateway

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasCombinedGateway() bool`

HasCombinedGateway returns a boolean if a field has been set.

### GetCurrentPage

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetData() []LanNetworkTemplateQueryOpenApiV3VO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetDataOk() (*[]LanNetworkTemplateQueryOpenApiV3VO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetData(v []LanNetworkTemplateQueryOpenApiV3VO)`

SetData sets Data field to given value.

### HasData

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDhcpRangePoolSize

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetDhcpRangePoolSize() int32`

GetDhcpRangePoolSize returns the DhcpRangePoolSize field if non-nil, zero value otherwise.

### GetDhcpRangePoolSizeOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetDhcpRangePoolSizeOk() (*int32, bool)`

GetDhcpRangePoolSizeOk returns a tuple with the DhcpRangePoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRangePoolSize

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetDhcpRangePoolSize(v int32)`

SetDhcpRangePoolSize sets DhcpRangePoolSize field to given value.

### HasDhcpRangePoolSize

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasDhcpRangePoolSize() bool`

HasDhcpRangePoolSize returns a boolean if a field has been set.

### GetInterfaceNum

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetInterfaceNum() int32`

GetInterfaceNum returns the InterfaceNum field if non-nil, zero value otherwise.

### GetInterfaceNumOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetInterfaceNumOk() (*int32, bool)`

GetInterfaceNumOk returns a tuple with the InterfaceNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceNum

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetInterfaceNum(v int32)`

SetInterfaceNum sets InterfaceNum field to given value.

### HasInterfaceNum

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasInterfaceNum() bool`

HasInterfaceNum returns a boolean if a field has been set.

### GetOsgSupportIgmpSnooping

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetOsgSupportIgmpSnooping() []string`

GetOsgSupportIgmpSnooping returns the OsgSupportIgmpSnooping field if non-nil, zero value otherwise.

### GetOsgSupportIgmpSnoopingOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetOsgSupportIgmpSnoopingOk() (*[]string, bool)`

GetOsgSupportIgmpSnoopingOk returns a tuple with the OsgSupportIgmpSnooping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsgSupportIgmpSnooping

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetOsgSupportIgmpSnooping(v []string)`

SetOsgSupportIgmpSnooping sets OsgSupportIgmpSnooping field to given value.

### HasOsgSupportIgmpSnooping

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasOsgSupportIgmpSnooping() bool`

HasOsgSupportIgmpSnooping returns a boolean if a field has been set.

### GetSupportArpDetection

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportArpDetection() bool`

GetSupportArpDetection returns the SupportArpDetection field if non-nil, zero value otherwise.

### GetSupportArpDetectionOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportArpDetectionOk() (*bool, bool)`

GetSupportArpDetectionOk returns a tuple with the SupportArpDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportArpDetection

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportArpDetection(v bool)`

SetSupportArpDetection sets SupportArpDetection field to given value.

### HasSupportArpDetection

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportArpDetection() bool`

HasSupportArpDetection returns a boolean if a field has been set.

### GetSupportCustomDhcpOption

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportCustomDhcpOption() bool`

GetSupportCustomDhcpOption returns the SupportCustomDhcpOption field if non-nil, zero value otherwise.

### GetSupportCustomDhcpOptionOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportCustomDhcpOptionOk() (*bool, bool)`

GetSupportCustomDhcpOptionOk returns a tuple with the SupportCustomDhcpOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCustomDhcpOption

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportCustomDhcpOption(v bool)`

SetSupportCustomDhcpOption sets SupportCustomDhcpOption field to given value.

### HasSupportCustomDhcpOption

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportCustomDhcpOption() bool`

HasSupportCustomDhcpOption returns a boolean if a field has been set.

### GetSupportDefault

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportDefault() bool`

GetSupportDefault returns the SupportDefault field if non-nil, zero value otherwise.

### GetSupportDefaultOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportDefaultOk() (*bool, bool)`

GetSupportDefaultOk returns a tuple with the SupportDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDefault

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportDefault(v bool)`

SetSupportDefault sets SupportDefault field to given value.

### HasSupportDefault

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportDefault() bool`

HasSupportDefault returns a boolean if a field has been set.

### GetSupportDhcpNextServer

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportDhcpNextServer() bool`

GetSupportDhcpNextServer returns the SupportDhcpNextServer field if non-nil, zero value otherwise.

### GetSupportDhcpNextServerOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportDhcpNextServerOk() (*bool, bool)`

GetSupportDhcpNextServerOk returns a tuple with the SupportDhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDhcpNextServer

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportDhcpNextServer(v bool)`

SetSupportDhcpNextServer sets SupportDhcpNextServer field to given value.

### HasSupportDhcpNextServer

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportDhcpNextServer() bool`

HasSupportDhcpNextServer returns a boolean if a field has been set.

### GetSupportLanIpv6

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportLanIpv6() bool`

GetSupportLanIpv6 returns the SupportLanIpv6 field if non-nil, zero value otherwise.

### GetSupportLanIpv6Ok

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportLanIpv6Ok() (*bool, bool)`

GetSupportLanIpv6Ok returns a tuple with the SupportLanIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLanIpv6

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportLanIpv6(v bool)`

SetSupportLanIpv6 sets SupportLanIpv6 field to given value.

### HasSupportLanIpv6

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportLanIpv6() bool`

HasSupportLanIpv6 returns a boolean if a field has been set.

### GetSupportLanIpv6PassThrough

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportLanIpv6PassThrough() bool`

GetSupportLanIpv6PassThrough returns the SupportLanIpv6PassThrough field if non-nil, zero value otherwise.

### GetSupportLanIpv6PassThroughOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportLanIpv6PassThroughOk() (*bool, bool)`

GetSupportLanIpv6PassThroughOk returns a tuple with the SupportLanIpv6PassThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLanIpv6PassThrough

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportLanIpv6PassThrough(v bool)`

SetSupportLanIpv6PassThrough sets SupportLanIpv6PassThrough field to given value.

### HasSupportLanIpv6PassThrough

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportLanIpv6PassThrough() bool`

HasSupportLanIpv6PassThrough returns a boolean if a field has been set.

### GetSupportMaxVlanNum

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportMaxVlanNum() int32`

GetSupportMaxVlanNum returns the SupportMaxVlanNum field if non-nil, zero value otherwise.

### GetSupportMaxVlanNumOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportMaxVlanNumOk() (*int32, bool)`

GetSupportMaxVlanNumOk returns a tuple with the SupportMaxVlanNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMaxVlanNum

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportMaxVlanNum(v int32)`

SetSupportMaxVlanNum sets SupportMaxVlanNum field to given value.

### HasSupportMaxVlanNum

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportMaxVlanNum() bool`

HasSupportMaxVlanNum returns a boolean if a field has been set.

### GetSupportMultiVlan

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportMultiVlan() bool`

GetSupportMultiVlan returns the SupportMultiVlan field if non-nil, zero value otherwise.

### GetSupportMultiVlanOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportMultiVlanOk() (*bool, bool)`

GetSupportMultiVlanOk returns a tuple with the SupportMultiVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMultiVlan

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportMultiVlan(v bool)`

SetSupportMultiVlan sets SupportMultiVlan field to given value.

### HasSupportMultiVlan

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportMultiVlan() bool`

HasSupportMultiVlan returns a boolean if a field has been set.

### GetSupportNetworkIsolation

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportNetworkIsolation() bool`

GetSupportNetworkIsolation returns the SupportNetworkIsolation field if non-nil, zero value otherwise.

### GetSupportNetworkIsolationOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportNetworkIsolationOk() (*bool, bool)`

GetSupportNetworkIsolationOk returns a tuple with the SupportNetworkIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportNetworkIsolation

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportNetworkIsolation(v bool)`

SetSupportNetworkIsolation sets SupportNetworkIsolation field to given value.

### HasSupportNetworkIsolation

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportNetworkIsolation() bool`

HasSupportNetworkIsolation returns a boolean if a field has been set.

### GetSupportPdOnDhcp

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportPdOnDhcp() bool`

GetSupportPdOnDhcp returns the SupportPdOnDhcp field if non-nil, zero value otherwise.

### GetSupportPdOnDhcpOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportPdOnDhcpOk() (*bool, bool)`

GetSupportPdOnDhcpOk returns a tuple with the SupportPdOnDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPdOnDhcp

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportPdOnDhcp(v bool)`

SetSupportPdOnDhcp sets SupportPdOnDhcp field to given value.

### HasSupportPdOnDhcp

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportPdOnDhcp() bool`

HasSupportPdOnDhcp returns a boolean if a field has been set.

### GetSupportRA

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportRA() bool`

GetSupportRA returns the SupportRA field if non-nil, zero value otherwise.

### GetSupportRAOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetSupportRAOk() (*bool, bool)`

GetSupportRAOk returns a tuple with the SupportRA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportRA

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetSupportRA(v bool)`

SetSupportRA sets SupportRA field to given value.

### HasSupportRA

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasSupportRA() bool`

HasSupportRA returns a boolean if a field has been set.

### GetTotalRows

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.

### GetVlanNums

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetVlanNums() int32`

GetVlanNums returns the VlanNums field if non-nil, zero value otherwise.

### GetVlanNumsOk

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) GetVlanNumsOk() (*int32, bool)`

GetVlanNumsOk returns a tuple with the VlanNums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNums

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) SetVlanNums(v int32)`

SetVlanNums sets VlanNums field to given value.

### HasVlanNums

`func (o *LanNetworkOpenApiV3GridVOLanNetworkTemplateQueryOpenApiV3VO) HasVlanNums() bool`

HasVlanNums returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


