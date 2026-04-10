# VpnUserOpenApiGridVOVpnUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VpnUserResponse**](VpnUserResponse.md) |  | [optional] 
**MaxConcurrentUser** | Pointer to **int32** | The maximum number of concurrent users supported by VPN. | [optional] 
**SslVpnMaxConUserNum** | Pointer to **int32** | The maximum number of concurrent users supported by SSL VPN. | [optional] 
**SubnetsLimitSize** | Pointer to **int32** | The maximum number of subnets allowed for a VPN user. | [optional] 
**SupportL2TP** | Pointer to **bool** | Whether the L2TP configuration is supported by the VPN user. | [optional] 
**SupportLocalIp** | Pointer to **bool** | Whether the local IP address configuration is supported by the VPN user. | [optional] 
**SupportOpenVpn** | Pointer to **bool** | Whether the OpenVPN configuration is supported by the VPN user. | [optional] 
**SupportProtocol** | Pointer to **bool** | Whether protocol configuration is supported of the VPN user | [optional] 
**SupportServerOptional** | Pointer to **bool** | Whether VPN server is optional. | [optional] 
**SupportSslVpn** | Pointer to **bool** | Whether the SSLVPN configuration is supported by the VPN user. | [optional] 
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

### GetMaxConcurrentUser

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetMaxConcurrentUser() int32`

GetMaxConcurrentUser returns the MaxConcurrentUser field if non-nil, zero value otherwise.

### GetMaxConcurrentUserOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetMaxConcurrentUserOk() (*int32, bool)`

GetMaxConcurrentUserOk returns a tuple with the MaxConcurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentUser

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetMaxConcurrentUser(v int32)`

SetMaxConcurrentUser sets MaxConcurrentUser field to given value.

### HasMaxConcurrentUser

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasMaxConcurrentUser() bool`

HasMaxConcurrentUser returns a boolean if a field has been set.

### GetSslVpnMaxConUserNum

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSslVpnMaxConUserNum() int32`

GetSslVpnMaxConUserNum returns the SslVpnMaxConUserNum field if non-nil, zero value otherwise.

### GetSslVpnMaxConUserNumOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSslVpnMaxConUserNumOk() (*int32, bool)`

GetSslVpnMaxConUserNumOk returns a tuple with the SslVpnMaxConUserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVpnMaxConUserNum

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetSslVpnMaxConUserNum(v int32)`

SetSslVpnMaxConUserNum sets SslVpnMaxConUserNum field to given value.

### HasSslVpnMaxConUserNum

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasSslVpnMaxConUserNum() bool`

HasSslVpnMaxConUserNum returns a boolean if a field has been set.

### GetSubnetsLimitSize

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSubnetsLimitSize() int32`

GetSubnetsLimitSize returns the SubnetsLimitSize field if non-nil, zero value otherwise.

### GetSubnetsLimitSizeOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSubnetsLimitSizeOk() (*int32, bool)`

GetSubnetsLimitSizeOk returns a tuple with the SubnetsLimitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsLimitSize

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetSubnetsLimitSize(v int32)`

SetSubnetsLimitSize sets SubnetsLimitSize field to given value.

### HasSubnetsLimitSize

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasSubnetsLimitSize() bool`

HasSubnetsLimitSize returns a boolean if a field has been set.

### GetSupportL2TP

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportL2TP() bool`

GetSupportL2TP returns the SupportL2TP field if non-nil, zero value otherwise.

### GetSupportL2TPOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportL2TPOk() (*bool, bool)`

GetSupportL2TPOk returns a tuple with the SupportL2TP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2TP

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetSupportL2TP(v bool)`

SetSupportL2TP sets SupportL2TP field to given value.

### HasSupportL2TP

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasSupportL2TP() bool`

HasSupportL2TP returns a boolean if a field has been set.

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

### GetSupportOpenVpn

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportOpenVpn() bool`

GetSupportOpenVpn returns the SupportOpenVpn field if non-nil, zero value otherwise.

### GetSupportOpenVpnOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportOpenVpnOk() (*bool, bool)`

GetSupportOpenVpnOk returns a tuple with the SupportOpenVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOpenVpn

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetSupportOpenVpn(v bool)`

SetSupportOpenVpn sets SupportOpenVpn field to given value.

### HasSupportOpenVpn

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasSupportOpenVpn() bool`

HasSupportOpenVpn returns a boolean if a field has been set.

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

### GetSupportServerOptional

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportServerOptional() bool`

GetSupportServerOptional returns the SupportServerOptional field if non-nil, zero value otherwise.

### GetSupportServerOptionalOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportServerOptionalOk() (*bool, bool)`

GetSupportServerOptionalOk returns a tuple with the SupportServerOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportServerOptional

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetSupportServerOptional(v bool)`

SetSupportServerOptional sets SupportServerOptional field to given value.

### HasSupportServerOptional

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasSupportServerOptional() bool`

HasSupportServerOptional returns a boolean if a field has been set.

### GetSupportSslVpn

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportSslVpn() bool`

GetSupportSslVpn returns the SupportSslVpn field if non-nil, zero value otherwise.

### GetSupportSslVpnOk

`func (o *VpnUserOpenApiGridVOVpnUserResponse) GetSupportSslVpnOk() (*bool, bool)`

GetSupportSslVpnOk returns a tuple with the SupportSslVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSslVpn

`func (o *VpnUserOpenApiGridVOVpnUserResponse) SetSupportSslVpn(v bool)`

SetSupportSslVpn sets SupportSslVpn field to given value.

### HasSupportSslVpn

`func (o *VpnUserOpenApiGridVOVpnUserResponse) HasSupportSslVpn() bool`

HasSupportSslVpn returns a boolean if a field has been set.

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


