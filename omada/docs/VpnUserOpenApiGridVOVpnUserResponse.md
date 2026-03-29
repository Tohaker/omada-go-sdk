# VpnUserOpenApiGridVOVpnUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VpnUserResponse**](VpnUserResponse.md) |  | [optional] 
**SupportLocalIp** | Pointer to **bool** | Whether local ip configuration is supported of the VPN user | [optional] 
**SupportProtocol** | Pointer to **bool** | Whether protocol configuration is supported of the VPN user | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewVpnUserOpenApiGridVOVpnUserResponse

`func NewVpnUserOpenApiGridVOVpnUserResponse() *VpnUserOpenApiGridVOVpnUserResponse`

NewVpnUserOpenApiGridVOVpnUserResponse instantiates a new VpnUserOpenApiGridVOVpnUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserOpenApiGridVOVpnUserResponseWithDefaults

`func NewVpnUserOpenApiGridVOVpnUserResponseWithDefaults() *VpnUserOpenApiGridVOVpnUserResponse`

NewVpnUserOpenApiGridVOVpnUserResponseWithDefaults instantiates a new VpnUserOpenApiGridVOVpnUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetData() []VpnUserResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetDataOk() (*[]VpnUserResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetData(v []VpnUserResponse)`

SetData sets Data field to given value.

### HasData

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportLocalIp

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportLocalIp() bool`

GetSupportLocalIp returns the SupportLocalIp field if non-nil, zero value otherwise.

### GetSupportLocalIpOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportLocalIpOk() (*bool, bool)`

GetSupportLocalIpOk returns a tuple with the SupportLocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocalIp

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetSupportLocalIp(v bool)`

SetSupportLocalIp sets SupportLocalIp field to given value.

### HasSupportLocalIp

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasSupportLocalIp() bool`

HasSupportLocalIp returns a boolean if a field has been set.

### GetSupportProtocol

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportProtocol() bool`

GetSupportProtocol returns the SupportProtocol field if non-nil, zero value otherwise.

### GetSupportProtocolOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportProtocolOk() (*bool, bool)`

GetSupportProtocolOk returns a tuple with the SupportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportProtocol

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetSupportProtocol(v bool)`

SetSupportProtocol sets SupportProtocol field to given value.

### HasSupportProtocol

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasSupportProtocol() bool`

HasSupportProtocol returns a boolean if a field has been set.

### GetTotalRows

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


