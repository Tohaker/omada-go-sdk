# PolicyRoutingOpenApiGridVOPolicyRoutingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]PolicyRoutingInfo**](PolicyRoutingInfo.md) |  | [optional] 
**SupportByDsLiteAndMapE** | Pointer to **bool** | Whether this feature is supported for the DS-Lite or Map-E WAN connection types. | [optional] 
**SupportDomainGroupDest** | Pointer to **bool** | Whether Domain Group is supported as Interface in Policy Routing. | [optional] 
**SupportLocationGroupDest** | Pointer to **bool** | Whether Location Group is supported as Destination in Policy Routing. | [optional] 
**SupportMulti** | Pointer to **bool** | Whether multiple WAN interface is supported in Policy Routing. | [optional] 
**SupportVpnClient** | Pointer to **bool** | Whether Vpn Client is supported as Interface in Policy Routing. | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewPolicyRoutingOpenApiGridVOPolicyRoutingInfo

`func NewPolicyRoutingOpenApiGridVOPolicyRoutingInfo() *PolicyRoutingOpenApiGridVOPolicyRoutingInfo`

NewPolicyRoutingOpenApiGridVOPolicyRoutingInfo instantiates a new PolicyRoutingOpenApiGridVOPolicyRoutingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRoutingOpenApiGridVOPolicyRoutingInfoWithDefaults

`func NewPolicyRoutingOpenApiGridVOPolicyRoutingInfoWithDefaults() *PolicyRoutingOpenApiGridVOPolicyRoutingInfo`

NewPolicyRoutingOpenApiGridVOPolicyRoutingInfoWithDefaults instantiates a new PolicyRoutingOpenApiGridVOPolicyRoutingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetData() []PolicyRoutingInfo`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetDataOk() (*[]PolicyRoutingInfo, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetData(v []PolicyRoutingInfo)`

SetData sets Data field to given value.

### HasData

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportByDsLiteAndMapE

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportByDsLiteAndMapE() bool`

GetSupportByDsLiteAndMapE returns the SupportByDsLiteAndMapE field if non-nil, zero value otherwise.

### GetSupportByDsLiteAndMapEOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportByDsLiteAndMapEOk() (*bool, bool)`

GetSupportByDsLiteAndMapEOk returns a tuple with the SupportByDsLiteAndMapE field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportByDsLiteAndMapE

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetSupportByDsLiteAndMapE(v bool)`

SetSupportByDsLiteAndMapE sets SupportByDsLiteAndMapE field to given value.

### HasSupportByDsLiteAndMapE

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasSupportByDsLiteAndMapE() bool`

HasSupportByDsLiteAndMapE returns a boolean if a field has been set.

### GetSupportDomainGroupDest

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportDomainGroupDest() bool`

GetSupportDomainGroupDest returns the SupportDomainGroupDest field if non-nil, zero value otherwise.

### GetSupportDomainGroupDestOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportDomainGroupDestOk() (*bool, bool)`

GetSupportDomainGroupDestOk returns a tuple with the SupportDomainGroupDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDomainGroupDest

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetSupportDomainGroupDest(v bool)`

SetSupportDomainGroupDest sets SupportDomainGroupDest field to given value.

### HasSupportDomainGroupDest

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasSupportDomainGroupDest() bool`

HasSupportDomainGroupDest returns a boolean if a field has been set.

### GetSupportLocationGroupDest

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportLocationGroupDest() bool`

GetSupportLocationGroupDest returns the SupportLocationGroupDest field if non-nil, zero value otherwise.

### GetSupportLocationGroupDestOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportLocationGroupDestOk() (*bool, bool)`

GetSupportLocationGroupDestOk returns a tuple with the SupportLocationGroupDest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocationGroupDest

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetSupportLocationGroupDest(v bool)`

SetSupportLocationGroupDest sets SupportLocationGroupDest field to given value.

### HasSupportLocationGroupDest

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasSupportLocationGroupDest() bool`

HasSupportLocationGroupDest returns a boolean if a field has been set.

### GetSupportMulti

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportMulti() bool`

GetSupportMulti returns the SupportMulti field if non-nil, zero value otherwise.

### GetSupportMultiOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportMultiOk() (*bool, bool)`

GetSupportMultiOk returns a tuple with the SupportMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMulti

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetSupportMulti(v bool)`

SetSupportMulti sets SupportMulti field to given value.

### HasSupportMulti

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasSupportMulti() bool`

HasSupportMulti returns a boolean if a field has been set.

### GetSupportVpnClient

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportVpnClient() bool`

GetSupportVpnClient returns the SupportVpnClient field if non-nil, zero value otherwise.

### GetSupportVpnClientOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetSupportVpnClientOk() (*bool, bool)`

GetSupportVpnClientOk returns a tuple with the SupportVpnClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVpnClient

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetSupportVpnClient(v bool)`

SetSupportVpnClient sets SupportVpnClient field to given value.

### HasSupportVpnClient

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasSupportVpnClient() bool`

HasSupportVpnClient returns a boolean if a field has been set.

### GetTotalRows

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *PolicyRoutingOpenApiGridVOPolicyRoutingInfo) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


