# SslVpnResourceEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** | Domain of the SSL VPN resource, exists when type is 1. | [optional] 
**DstPortEnd** | Pointer to **int32** | End destination port of the SSL VPN resource, exists when protocol is TCP or UDP. It should be within the range of 0–65535 | [optional] 
**DstPortStart** | Pointer to **int32** | Start destination port of the SSL VPN resource, exists when protocol is TCP or UDP. It should be within the range of 0–65535 | [optional] 
**IcmpCode** | Pointer to **int32** | ICMP code of the SSL VPN resource, exists when protocol is ICMP. It should be within the range of 0–255 | [optional] 
**IcmpType** | Pointer to **int32** | ICMP type of the SSL VPN resource, exists when protocol is ICMP. It should be within the range of 0–255 | [optional] 
**Id** | Pointer to **string** | ID of the SSL VPN resource. | [optional] 
**Ip** | Pointer to **string** | IP of the SSL VPN resource, exists when type is 0. | [optional] 
**Mask** | Pointer to **string** | Mask of the SSL VPN resource, exists when type is 0. | [optional] 
**Name** | **string** | Name of the SSL VPN resource. | 
**OtherProtocol** | Pointer to **int32** | Other protocol of the SSL VPN resource. It should be within the range of 1–255 | [optional] 
**Protocol** | Pointer to **int32** | Protocol of the SSL VPN resource should be a value as follows: 0:All; 1:TCP; 2:UDP; 3:TCP/UDP; 4:ICMP; 5:Other | [optional] 
**ResourceGroupList** | Pointer to [**[]SslVpnResourceGroupBriefInfo**](SslVpnResourceGroupBriefInfo.md) | Resource group list of the SSL VPN resource. | [optional] 
**SrcPortEnd** | Pointer to **int32** | End source port of the SSL VPN resource, exists when protocol is TCP or UDP. It should be within the range of 0–65535 | [optional] 
**SrcPortStart** | Pointer to **int32** | Start source port of the SSL VPN resource, exists when protocol is TCP or UDP. It should be within the range of 0–65535 | [optional] 
**Type** | **int32** | Type of the SSL VPN resource should be a value as follows: 0: IP; 1: domain. | 

## Methods

### NewSslVpnResourceEntity

`func NewSslVpnResourceEntity(name string, type_ int32, ) *SslVpnResourceEntity`

NewSslVpnResourceEntity instantiates a new SslVpnResourceEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnResourceEntityWithDefaults

`func NewSslVpnResourceEntityWithDefaults() *SslVpnResourceEntity`

NewSslVpnResourceEntityWithDefaults instantiates a new SslVpnResourceEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *SslVpnResourceEntity) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *SslVpnResourceEntity) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *SslVpnResourceEntity) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *SslVpnResourceEntity) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetDstPortEnd

`func (o *SslVpnResourceEntity) GetDstPortEnd() int32`

GetDstPortEnd returns the DstPortEnd field if non-nil, zero value otherwise.

### GetDstPortEndOk

`func (o *SslVpnResourceEntity) GetDstPortEndOk() (*int32, bool)`

GetDstPortEndOk returns a tuple with the DstPortEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPortEnd

`func (o *SslVpnResourceEntity) SetDstPortEnd(v int32)`

SetDstPortEnd sets DstPortEnd field to given value.

### HasDstPortEnd

`func (o *SslVpnResourceEntity) HasDstPortEnd() bool`

HasDstPortEnd returns a boolean if a field has been set.

### GetDstPortStart

`func (o *SslVpnResourceEntity) GetDstPortStart() int32`

GetDstPortStart returns the DstPortStart field if non-nil, zero value otherwise.

### GetDstPortStartOk

`func (o *SslVpnResourceEntity) GetDstPortStartOk() (*int32, bool)`

GetDstPortStartOk returns a tuple with the DstPortStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstPortStart

`func (o *SslVpnResourceEntity) SetDstPortStart(v int32)`

SetDstPortStart sets DstPortStart field to given value.

### HasDstPortStart

`func (o *SslVpnResourceEntity) HasDstPortStart() bool`

HasDstPortStart returns a boolean if a field has been set.

### GetIcmpCode

`func (o *SslVpnResourceEntity) GetIcmpCode() int32`

GetIcmpCode returns the IcmpCode field if non-nil, zero value otherwise.

### GetIcmpCodeOk

`func (o *SslVpnResourceEntity) GetIcmpCodeOk() (*int32, bool)`

GetIcmpCodeOk returns a tuple with the IcmpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpCode

`func (o *SslVpnResourceEntity) SetIcmpCode(v int32)`

SetIcmpCode sets IcmpCode field to given value.

### HasIcmpCode

`func (o *SslVpnResourceEntity) HasIcmpCode() bool`

HasIcmpCode returns a boolean if a field has been set.

### GetIcmpType

`func (o *SslVpnResourceEntity) GetIcmpType() int32`

GetIcmpType returns the IcmpType field if non-nil, zero value otherwise.

### GetIcmpTypeOk

`func (o *SslVpnResourceEntity) GetIcmpTypeOk() (*int32, bool)`

GetIcmpTypeOk returns a tuple with the IcmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmpType

`func (o *SslVpnResourceEntity) SetIcmpType(v int32)`

SetIcmpType sets IcmpType field to given value.

### HasIcmpType

`func (o *SslVpnResourceEntity) HasIcmpType() bool`

HasIcmpType returns a boolean if a field has been set.

### GetId

`func (o *SslVpnResourceEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SslVpnResourceEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SslVpnResourceEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SslVpnResourceEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *SslVpnResourceEntity) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SslVpnResourceEntity) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SslVpnResourceEntity) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SslVpnResourceEntity) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMask

`func (o *SslVpnResourceEntity) GetMask() string`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *SslVpnResourceEntity) GetMaskOk() (*string, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *SslVpnResourceEntity) SetMask(v string)`

SetMask sets Mask field to given value.

### HasMask

`func (o *SslVpnResourceEntity) HasMask() bool`

HasMask returns a boolean if a field has been set.

### GetName

`func (o *SslVpnResourceEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslVpnResourceEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslVpnResourceEntity) SetName(v string)`

SetName sets Name field to given value.


### GetOtherProtocol

`func (o *SslVpnResourceEntity) GetOtherProtocol() int32`

GetOtherProtocol returns the OtherProtocol field if non-nil, zero value otherwise.

### GetOtherProtocolOk

`func (o *SslVpnResourceEntity) GetOtherProtocolOk() (*int32, bool)`

GetOtherProtocolOk returns a tuple with the OtherProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherProtocol

`func (o *SslVpnResourceEntity) SetOtherProtocol(v int32)`

SetOtherProtocol sets OtherProtocol field to given value.

### HasOtherProtocol

`func (o *SslVpnResourceEntity) HasOtherProtocol() bool`

HasOtherProtocol returns a boolean if a field has been set.

### GetProtocol

`func (o *SslVpnResourceEntity) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *SslVpnResourceEntity) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *SslVpnResourceEntity) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *SslVpnResourceEntity) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetResourceGroupList

`func (o *SslVpnResourceEntity) GetResourceGroupList() []SslVpnResourceGroupBriefInfo`

GetResourceGroupList returns the ResourceGroupList field if non-nil, zero value otherwise.

### GetResourceGroupListOk

`func (o *SslVpnResourceEntity) GetResourceGroupListOk() (*[]SslVpnResourceGroupBriefInfo, bool)`

GetResourceGroupListOk returns a tuple with the ResourceGroupList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceGroupList

`func (o *SslVpnResourceEntity) SetResourceGroupList(v []SslVpnResourceGroupBriefInfo)`

SetResourceGroupList sets ResourceGroupList field to given value.

### HasResourceGroupList

`func (o *SslVpnResourceEntity) HasResourceGroupList() bool`

HasResourceGroupList returns a boolean if a field has been set.

### GetSrcPortEnd

`func (o *SslVpnResourceEntity) GetSrcPortEnd() int32`

GetSrcPortEnd returns the SrcPortEnd field if non-nil, zero value otherwise.

### GetSrcPortEndOk

`func (o *SslVpnResourceEntity) GetSrcPortEndOk() (*int32, bool)`

GetSrcPortEndOk returns a tuple with the SrcPortEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPortEnd

`func (o *SslVpnResourceEntity) SetSrcPortEnd(v int32)`

SetSrcPortEnd sets SrcPortEnd field to given value.

### HasSrcPortEnd

`func (o *SslVpnResourceEntity) HasSrcPortEnd() bool`

HasSrcPortEnd returns a boolean if a field has been set.

### GetSrcPortStart

`func (o *SslVpnResourceEntity) GetSrcPortStart() int32`

GetSrcPortStart returns the SrcPortStart field if non-nil, zero value otherwise.

### GetSrcPortStartOk

`func (o *SslVpnResourceEntity) GetSrcPortStartOk() (*int32, bool)`

GetSrcPortStartOk returns a tuple with the SrcPortStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcPortStart

`func (o *SslVpnResourceEntity) SetSrcPortStart(v int32)`

SetSrcPortStart sets SrcPortStart field to given value.

### HasSrcPortStart

`func (o *SslVpnResourceEntity) HasSrcPortStart() bool`

HasSrcPortStart returns a boolean if a field has been set.

### GetType

`func (o *SslVpnResourceEntity) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SslVpnResourceEntity) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SslVpnResourceEntity) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


