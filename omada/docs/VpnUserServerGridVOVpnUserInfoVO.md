# VpnUserServerGridVOVpnUserInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **int32** |  | [optional] 
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]VpnUserInfoVO**](VpnUserInfoVO.md) |  | [optional] 
**Expired** | Pointer to **int32** |  | [optional] 
**MaxConcurrentUser** | Pointer to **int32** | The maximum number of concurrent users supported by VPN. | [optional] 
**SslVpnMaxConUserNum** | Pointer to **int32** | The maximum number of concurrent users supported by SSL VPN. | [optional] 
**SubnetsLimitSize** | Pointer to **int32** | The maximum number of subnets allowed for a VPN user. | [optional] 
**SupportL2TP** | Pointer to **bool** | Whether the L2TP configuration is supported by the VPN user. | [optional] 
**SupportLocalIp** | Pointer to **bool** | Whether the local IP address configuration is supported by the VPN user. | [optional] 
**SupportOpenVpn** | Pointer to **bool** | Whether the OpenVPN configuration is supported by the VPN user. | [optional] 
**SupportProtocol** | Pointer to **bool** | Whether protocol configuration is supported of the VPN user | [optional] 
**SupportServerOptional** | Pointer to **bool** | Whether VPN server is optional. | [optional] 
**SupportSslVpn** | Pointer to **bool** | Whether the SSLVPN configuration is supported by the VPN user. | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 
**VpnInfo** | Pointer to [**VpnUserServerInfoVO**](VpnUserServerInfoVO.md) |  | [optional] 

## Methods

### NewVpnUserServerGridVOVpnUserInfoVO

`func NewVpnUserServerGridVOVpnUserInfoVO() *VpnUserServerGridVOVpnUserInfoVO`

NewVpnUserServerGridVOVpnUserInfoVO instantiates a new VpnUserServerGridVOVpnUserInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserServerGridVOVpnUserInfoVOWithDefaults

`func NewVpnUserServerGridVOVpnUserInfoVOWithDefaults() *VpnUserServerGridVOVpnUserInfoVO`

NewVpnUserServerGridVOVpnUserInfoVOWithDefaults instantiates a new VpnUserServerGridVOVpnUserInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetAvailable() int32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetAvailableOk() (*int32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetAvailable(v int32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCurrentPage

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetData() []VpnUserInfoVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetDataOk() (*[]VpnUserInfoVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetData(v []VpnUserInfoVO)`

SetData sets Data field to given value.

### HasData

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetExpired

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetExpired() int32`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetExpiredOk() (*int32, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetExpired(v int32)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetMaxConcurrentUser

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetMaxConcurrentUser() int32`

GetMaxConcurrentUser returns the MaxConcurrentUser field if non-nil, zero value otherwise.

### GetMaxConcurrentUserOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetMaxConcurrentUserOk() (*int32, bool)`

GetMaxConcurrentUserOk returns a tuple with the MaxConcurrentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentUser

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetMaxConcurrentUser(v int32)`

SetMaxConcurrentUser sets MaxConcurrentUser field to given value.

### HasMaxConcurrentUser

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasMaxConcurrentUser() bool`

HasMaxConcurrentUser returns a boolean if a field has been set.

### GetSslVpnMaxConUserNum

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSslVpnMaxConUserNum() int32`

GetSslVpnMaxConUserNum returns the SslVpnMaxConUserNum field if non-nil, zero value otherwise.

### GetSslVpnMaxConUserNumOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSslVpnMaxConUserNumOk() (*int32, bool)`

GetSslVpnMaxConUserNumOk returns a tuple with the SslVpnMaxConUserNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslVpnMaxConUserNum

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetSslVpnMaxConUserNum(v int32)`

SetSslVpnMaxConUserNum sets SslVpnMaxConUserNum field to given value.

### HasSslVpnMaxConUserNum

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasSslVpnMaxConUserNum() bool`

HasSslVpnMaxConUserNum returns a boolean if a field has been set.

### GetSubnetsLimitSize

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSubnetsLimitSize() int32`

GetSubnetsLimitSize returns the SubnetsLimitSize field if non-nil, zero value otherwise.

### GetSubnetsLimitSizeOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSubnetsLimitSizeOk() (*int32, bool)`

GetSubnetsLimitSizeOk returns a tuple with the SubnetsLimitSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetsLimitSize

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetSubnetsLimitSize(v int32)`

SetSubnetsLimitSize sets SubnetsLimitSize field to given value.

### HasSubnetsLimitSize

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasSubnetsLimitSize() bool`

HasSubnetsLimitSize returns a boolean if a field has been set.

### GetSupportL2TP

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportL2TP() bool`

GetSupportL2TP returns the SupportL2TP field if non-nil, zero value otherwise.

### GetSupportL2TPOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportL2TPOk() (*bool, bool)`

GetSupportL2TPOk returns a tuple with the SupportL2TP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2TP

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetSupportL2TP(v bool)`

SetSupportL2TP sets SupportL2TP field to given value.

### HasSupportL2TP

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasSupportL2TP() bool`

HasSupportL2TP returns a boolean if a field has been set.

### GetSupportLocalIp

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportLocalIp() bool`

GetSupportLocalIp returns the SupportLocalIp field if non-nil, zero value otherwise.

### GetSupportLocalIpOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportLocalIpOk() (*bool, bool)`

GetSupportLocalIpOk returns a tuple with the SupportLocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportLocalIp

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetSupportLocalIp(v bool)`

SetSupportLocalIp sets SupportLocalIp field to given value.

### HasSupportLocalIp

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasSupportLocalIp() bool`

HasSupportLocalIp returns a boolean if a field has been set.

### GetSupportOpenVpn

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportOpenVpn() bool`

GetSupportOpenVpn returns the SupportOpenVpn field if non-nil, zero value otherwise.

### GetSupportOpenVpnOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportOpenVpnOk() (*bool, bool)`

GetSupportOpenVpnOk returns a tuple with the SupportOpenVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOpenVpn

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetSupportOpenVpn(v bool)`

SetSupportOpenVpn sets SupportOpenVpn field to given value.

### HasSupportOpenVpn

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasSupportOpenVpn() bool`

HasSupportOpenVpn returns a boolean if a field has been set.

### GetSupportProtocol

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportProtocol() bool`

GetSupportProtocol returns the SupportProtocol field if non-nil, zero value otherwise.

### GetSupportProtocolOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportProtocolOk() (*bool, bool)`

GetSupportProtocolOk returns a tuple with the SupportProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportProtocol

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetSupportProtocol(v bool)`

SetSupportProtocol sets SupportProtocol field to given value.

### HasSupportProtocol

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasSupportProtocol() bool`

HasSupportProtocol returns a boolean if a field has been set.

### GetSupportServerOptional

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportServerOptional() bool`

GetSupportServerOptional returns the SupportServerOptional field if non-nil, zero value otherwise.

### GetSupportServerOptionalOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportServerOptionalOk() (*bool, bool)`

GetSupportServerOptionalOk returns a tuple with the SupportServerOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportServerOptional

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetSupportServerOptional(v bool)`

SetSupportServerOptional sets SupportServerOptional field to given value.

### HasSupportServerOptional

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasSupportServerOptional() bool`

HasSupportServerOptional returns a boolean if a field has been set.

### GetSupportSslVpn

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportSslVpn() bool`

GetSupportSslVpn returns the SupportSslVpn field if non-nil, zero value otherwise.

### GetSupportSslVpnOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetSupportSslVpnOk() (*bool, bool)`

GetSupportSslVpnOk returns a tuple with the SupportSslVpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSslVpn

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetSupportSslVpn(v bool)`

SetSupportSslVpn sets SupportSslVpn field to given value.

### HasSupportSslVpn

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasSupportSslVpn() bool`

HasSupportSslVpn returns a boolean if a field has been set.

### GetTotal

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetTotalRows

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.

### GetVpnInfo

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetVpnInfo() VpnUserServerInfoVO`

GetVpnInfo returns the VpnInfo field if non-nil, zero value otherwise.

### GetVpnInfoOk

`func (o *VpnUserServerGridVOVpnUserInfoVO) GetVpnInfoOk() (*VpnUserServerInfoVO, bool)`

GetVpnInfoOk returns a tuple with the VpnInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnInfo

`func (o *VpnUserServerGridVOVpnUserInfoVO) SetVpnInfo(v VpnUserServerInfoVO)`

SetVpnInfo sets VpnInfo field to given value.

### HasVpnInfo

`func (o *VpnUserServerGridVOVpnUserInfoVO) HasVpnInfo() bool`

HasVpnInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


