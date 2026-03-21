# SslVpnResourceModifyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain of the SSL VPN resource, exists when type is 1. | [optional] 
**DstPortEnd** | Pointer to **int32** | End destination port of the SSL VPN resource, exists when protocol is TCP or UDP. It should be within the range of 0–65535 | [optional] 
**DstPortStart** | Pointer to **int32** | Start destination port of the SSL VPN resource, exists when protocol is TCP or UDP. It should be within the range of 0–65535 | [optional] 
**IcmpCode** | Pointer to **int32** | ICMP code of the SSL VPN resource, exists when protocol is ICMP. It should be within the range of 0–255 | [optional] 
**IcmpType** | Pointer to **int32** | ICMP type of the SSL VPN resource, exists when protocol is ICMP. It should be within the range of 0–255 | [optional] 
**Ip** | Pointer to **string** | IP of the SSL VPN resource, exists when type is 0. | [optional] 
**Mask** | Pointer to **string** | Mask of the SSL VPN resource, exists when type is 0. | [optional] 
**OtherProtocol** | Pointer to **int32** | Other protocol of the SSL VPN resource. It should be within the range of 1–255 | [optional] 
**Protocol** | Pointer to **int32** | Protocol of the SSL VPN resource should be a value as follows: 0:All; 1:TCP; 2:UDP; 3:TCP/UDP; 4:ICMP; 5:Other | [optional] 
**SrcPortEnd** | Pointer to **int32** | End source port of the SSL VPN resource, exists when protocol is TCP or UDP. It should be within the range of 0–65535 | [optional] 
**SrcPortStart** | Pointer to **int32** | Start source port of the SSL VPN resource, exists when protocol is TCP or UDP. It should be within the range of 0–65535 | [optional] 
**Type** | **int32** | Type of the SSL VPN resource should be a value as follows: 0: IP; 1: domain, and it can not be modified. | 

## Methods

### NewSslVpnResourceModifyOpenApiVO

`func NewSslVpnResourceModifyOpenApiVO(type_ int32, ) *SslVpnResourceModifyOpenApiVO`

NewSslVpnResourceModifyOpenApiVO instantiates a new SslVpnResourceModifyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnResourceModifyOpenApiVOWithDefaults

`func NewSslVpnResourceModifyOpenApiVOWithDefaults() *SslVpnResourceModifyOpenApiVO`

NewSslVpnResourceModifyOpenApiVOWithDefaults instantiates a new SslVpnResourceModifyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SslVpnResourceModifyOpenApiVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SslVpnResourceModifyOpenApiVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SslVpnResourceModifyOpenApiVO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SslVpnResourceModifyOpenApiVO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDstPortEnd

`func (o *SslVpnResourceModifyOpenApiVO) GetDstPortEnd() int32`

GetDstPortEnd returns the DstPortEnd field if non-nil, zero value otherwise.

### GetDstPortEndOk

`func (o *SslVpnResourceModifyOpenApiVO) GetDstPortEndOk() (*int32, bool)`

GetDstPortEndOk returns a tuple with the DstPortEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPortEnd

`func (o *SslVpnResourceModifyOpenApiVO) SetDstPortEnd(v int32)`

SetDstPortEnd sets DstPortEnd field to given value.

### HasDstPortEnd

`func (o *SslVpnResourceModifyOpenApiVO) HasDstPortEnd() bool`

HasDstPortEnd returns a boolean if a field has been set.

### GetDstPortStart

`func (o *SslVpnResourceModifyOpenApiVO) GetDstPortStart() int32`

GetDstPortStart returns the DstPortStart field if non-nil, zero value otherwise.

### GetDstPortStartOk

`func (o *SslVpnResourceModifyOpenApiVO) GetDstPortStartOk() (*int32, bool)`

GetDstPortStartOk returns a tuple with the DstPortStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPortStart

`func (o *SslVpnResourceModifyOpenApiVO) SetDstPortStart(v int32)`

SetDstPortStart sets DstPortStart field to given value.

### HasDstPortStart

`func (o *SslVpnResourceModifyOpenApiVO) HasDstPortStart() bool`

HasDstPortStart returns a boolean if a field has been set.

### GetIcmpCode

`func (o *SslVpnResourceModifyOpenApiVO) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *SslVpnResourceModifyOpenApiVO) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *SslVpnResourceModifyOpenApiVO) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.

### HasIcmpCode

`func (o *SslVpnResourceModifyOpenApiVO) HasIcmpCode() bool`

HasIcmpCode returns a boolean if a field has been set.

### GetIcmpType

`func (o *SslVpnResourceModifyOpenApiVO) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *SslVpnResourceModifyOpenApiVO) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *SslVpnResourceModifyOpenApiVO) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *SslVpnResourceModifyOpenApiVO) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.

### GetIp

`func (o *SslVpnResourceModifyOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SslVpnResourceModifyOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SslVpnResourceModifyOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SslVpnResourceModifyOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMask

`func (o *SslVpnResourceModifyOpenApiVO) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *SslVpnResourceModifyOpenApiVO) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *SslVpnResourceModifyOpenApiVO) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *SslVpnResourceModifyOpenApiVO) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetOtherProtocol

`func (o *SslVpnResourceModifyOpenApiVO) GetOtherProtocol() int32`

GetOtherProtocol returns the OtherProtocol field if non-nil, zero value otherwise.

### GetOtherProtocolOk

`func (o *SslVpnResourceModifyOpenApiVO) GetOtherProtocolOk() (*int32, bool)`

GetOtherProtocolOk returns a tuple with the OtherProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherProtocol

`func (o *SslVpnResourceModifyOpenApiVO) SetOtherProtocol(v int32)`

SetOtherProtocol sets OtherProtocol field to given value.

### HasOtherProtocol

`func (o *SslVpnResourceModifyOpenApiVO) HasOtherProtocol() bool`

HasOtherProtocol returns a boolean if a field has been set.

### GetProtocol

`func (o *SslVpnResourceModifyOpenApiVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SslVpnResourceModifyOpenApiVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SslVpnResourceModifyOpenApiVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SslVpnResourceModifyOpenApiVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSrcPortEnd

`func (o *SslVpnResourceModifyOpenApiVO) GetSrcPortEnd() int32`

GetSrcPortEnd returns the SrcPortEnd field if non-nil, zero value otherwise.

### GetSrcPortEndOk

`func (o *SslVpnResourceModifyOpenApiVO) GetSrcPortEndOk() (*int32, bool)`

GetSrcPortEndOk returns a tuple with the SrcPortEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPortEnd

`func (o *SslVpnResourceModifyOpenApiVO) SetSrcPortEnd(v int32)`

SetSrcPortEnd sets SrcPortEnd field to given value.

### HasSrcPortEnd

`func (o *SslVpnResourceModifyOpenApiVO) HasSrcPortEnd() bool`

HasSrcPortEnd returns a boolean if a field has been set.

### GetSrcPortStart

`func (o *SslVpnResourceModifyOpenApiVO) GetSrcPortStart() int32`

GetSrcPortStart returns the SrcPortStart field if non-nil, zero value otherwise.

### GetSrcPortStartOk

`func (o *SslVpnResourceModifyOpenApiVO) GetSrcPortStartOk() (*int32, bool)`

GetSrcPortStartOk returns a tuple with the SrcPortStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPortStart

`func (o *SslVpnResourceModifyOpenApiVO) SetSrcPortStart(v int32)`

SetSrcPortStart sets SrcPortStart field to given value.

### HasSrcPortStart

`func (o *SslVpnResourceModifyOpenApiVO) HasSrcPortStart() bool`

HasSrcPortStart returns a boolean if a field has been set.

### GetType

`func (o *SslVpnResourceModifyOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SslVpnResourceModifyOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SslVpnResourceModifyOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


