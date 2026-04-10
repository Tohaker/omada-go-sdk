# GetAlgOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExistFtpAlgConfig** | Pointer to **bool** | Whether to exist custom configuration for FTP ALG | [optional] 
**ExistSipAlgConfig** | Pointer to **bool** | Whether to exist custom configuration for SIP ALG | [optional] 
**Ftp** | **bool** | Whether to enable the FTP ALG | 
**FtpPorts** | Pointer to **[]int32** | The listening port for FTP signaling messages, the default port is 21 | [optional] 
**H323** | **bool** | Whether to enable the H323 ALG | 
**IpSec** | **bool** | Whether to enable the IPSec ALG | 
**Pptp** | **bool** | Whether to enable the PPTP ALG | 
**Sip** | **bool** | Whether to enable the SIP ALG | 
**SipDirectMedia** | Pointer to **bool** | Enable this option if you want SIP media connections to only arrive from registered IP addresses. | [optional] 
**SipDirectSignaling** | Pointer to **bool** | Enable this option if you want SIP signaling connections to only arrive from registered IP addresses, enabled by default. | [optional] 
**SipMediaTimeout** | Pointer to **int32** | Modify this item if you want to adjust the timeout duration for SIP media sessions (1 – 86,400 seconds). | [optional] 
**SipPorts** | Pointer to **[]int32** | The listening ports for SIP signaling messages, the default ports are 5060 and 5061 | [optional] 
**SipSignalingTimeout** | Pointer to **int32** | Modify this item if you want to adjust the timeout duration for SIP signaling sessions (1 – 86,400 seconds). | [optional] 
**SipTcp** | Pointer to **bool** | Enable this option if your SIP signaling uses the TCP protocol and the device needs to modify the IP address and port in SIP messages, enabled by default. | [optional] 
**SipTimeout** | Pointer to **bool** | Enable this option if you want to apply timeout limits to SIP signaling and media connections on the device. | [optional] 
**SipUdp** | Pointer to **bool** | Enable this option if your SIP signaling uses the UDP protocol and the device needs to modify the IP address and port in SIP signaling messages, enabled by default. | [optional] 
**SupportSipAlgConfig** | Pointer to **bool** | Whether to support custom configuration for ALG | [optional] 

## Methods

### NewGetAlgOpenApiVO

`func NewGetAlgOpenApiVO(ftp bool, h323 bool, ipSec bool, pptp bool, sip bool, ) *GetAlgOpenApiVO`

NewGetAlgOpenApiVO instantiates a new GetAlgOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAlgOpenApiVOWithDefaults

`func NewGetAlgOpenApiVOWithDefaults() *GetAlgOpenApiVO`

NewGetAlgOpenApiVOWithDefaults instantiates a new GetAlgOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExistFtpAlgConfig

`func (o *GetAlgOpenApiVO) GetExistFtpAlgConfig() bool`

GetExistFtpAlgConfig returns the ExistFtpAlgConfig field if non-nil, zero value otherwise.

### GetExistFtpAlgConfigOk

`func (o *GetAlgOpenApiVO) GetExistFtpAlgConfigOk() (*bool, bool)`

GetExistFtpAlgConfigOk returns a tuple with the ExistFtpAlgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistFtpAlgConfig

`func (o *GetAlgOpenApiVO) SetExistFtpAlgConfig(v bool)`

SetExistFtpAlgConfig sets ExistFtpAlgConfig field to given value.

### HasExistFtpAlgConfig

`func (o *GetAlgOpenApiVO) HasExistFtpAlgConfig() bool`

HasExistFtpAlgConfig returns a boolean if a field has been set.

### GetExistSipAlgConfig

`func (o *GetAlgOpenApiVO) GetExistSipAlgConfig() bool`

GetExistSipAlgConfig returns the ExistSipAlgConfig field if non-nil, zero value otherwise.

### GetExistSipAlgConfigOk

`func (o *GetAlgOpenApiVO) GetExistSipAlgConfigOk() (*bool, bool)`

GetExistSipAlgConfigOk returns a tuple with the ExistSipAlgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistSipAlgConfig

`func (o *GetAlgOpenApiVO) SetExistSipAlgConfig(v bool)`

SetExistSipAlgConfig sets ExistSipAlgConfig field to given value.

### HasExistSipAlgConfig

`func (o *GetAlgOpenApiVO) HasExistSipAlgConfig() bool`

HasExistSipAlgConfig returns a boolean if a field has been set.

### GetFtp

`func (o *GetAlgOpenApiVO) GetFtp() bool`

GetFtp returns the Ftp field if non-nil, zero value otherwise.

### GetFtpOk

`func (o *GetAlgOpenApiVO) GetFtpOk() (*bool, bool)`

GetFtpOk returns a tuple with the Ftp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtp

`func (o *GetAlgOpenApiVO) SetFtp(v bool)`

SetFtp sets Ftp field to given value.


### GetFtpPorts

`func (o *GetAlgOpenApiVO) GetFtpPorts() []int32`

GetFtpPorts returns the FtpPorts field if non-nil, zero value otherwise.

### GetFtpPortsOk

`func (o *GetAlgOpenApiVO) GetFtpPortsOk() (*[]int32, bool)`

GetFtpPortsOk returns a tuple with the FtpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPorts

`func (o *GetAlgOpenApiVO) SetFtpPorts(v []int32)`

SetFtpPorts sets FtpPorts field to given value.

### HasFtpPorts

`func (o *GetAlgOpenApiVO) HasFtpPorts() bool`

HasFtpPorts returns a boolean if a field has been set.

### GetH323

`func (o *GetAlgOpenApiVO) GetH323() bool`

GetH323 returns the H323 field if non-nil, zero value otherwise.

### GetH323Ok

`func (o *GetAlgOpenApiVO) GetH323Ok() (*bool, bool)`

GetH323Ok returns a tuple with the H323 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH323

`func (o *GetAlgOpenApiVO) SetH323(v bool)`

SetH323 sets H323 field to given value.


### GetIpSec

`func (o *GetAlgOpenApiVO) GetIpSec() bool`

GetIpSec returns the IpSec field if non-nil, zero value otherwise.

### GetIpSecOk

`func (o *GetAlgOpenApiVO) GetIpSecOk() (*bool, bool)`

GetIpSecOk returns a tuple with the IpSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSec

`func (o *GetAlgOpenApiVO) SetIpSec(v bool)`

SetIpSec sets IpSec field to given value.


### GetPptp

`func (o *GetAlgOpenApiVO) GetPptp() bool`

GetPptp returns the Pptp field if non-nil, zero value otherwise.

### GetPptpOk

`func (o *GetAlgOpenApiVO) GetPptpOk() (*bool, bool)`

GetPptpOk returns a tuple with the Pptp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPptp

`func (o *GetAlgOpenApiVO) SetPptp(v bool)`

SetPptp sets Pptp field to given value.


### GetSip

`func (o *GetAlgOpenApiVO) GetSip() bool`

GetSip returns the Sip field if non-nil, zero value otherwise.

### GetSipOk

`func (o *GetAlgOpenApiVO) GetSipOk() (*bool, bool)`

GetSipOk returns a tuple with the Sip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSip

`func (o *GetAlgOpenApiVO) SetSip(v bool)`

SetSip sets Sip field to given value.


### GetSipDirectMedia

`func (o *GetAlgOpenApiVO) GetSipDirectMedia() bool`

GetSipDirectMedia returns the SipDirectMedia field if non-nil, zero value otherwise.

### GetSipDirectMediaOk

`func (o *GetAlgOpenApiVO) GetSipDirectMediaOk() (*bool, bool)`

GetSipDirectMediaOk returns a tuple with the SipDirectMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipDirectMedia

`func (o *GetAlgOpenApiVO) SetSipDirectMedia(v bool)`

SetSipDirectMedia sets SipDirectMedia field to given value.

### HasSipDirectMedia

`func (o *GetAlgOpenApiVO) HasSipDirectMedia() bool`

HasSipDirectMedia returns a boolean if a field has been set.

### GetSipDirectSignaling

`func (o *GetAlgOpenApiVO) GetSipDirectSignaling() bool`

GetSipDirectSignaling returns the SipDirectSignaling field if non-nil, zero value otherwise.

### GetSipDirectSignalingOk

`func (o *GetAlgOpenApiVO) GetSipDirectSignalingOk() (*bool, bool)`

GetSipDirectSignalingOk returns a tuple with the SipDirectSignaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipDirectSignaling

`func (o *GetAlgOpenApiVO) SetSipDirectSignaling(v bool)`

SetSipDirectSignaling sets SipDirectSignaling field to given value.

### HasSipDirectSignaling

`func (o *GetAlgOpenApiVO) HasSipDirectSignaling() bool`

HasSipDirectSignaling returns a boolean if a field has been set.

### GetSipMediaTimeout

`func (o *GetAlgOpenApiVO) GetSipMediaTimeout() int32`

GetSipMediaTimeout returns the SipMediaTimeout field if non-nil, zero value otherwise.

### GetSipMediaTimeoutOk

`func (o *GetAlgOpenApiVO) GetSipMediaTimeoutOk() (*int32, bool)`

GetSipMediaTimeoutOk returns a tuple with the SipMediaTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipMediaTimeout

`func (o *GetAlgOpenApiVO) SetSipMediaTimeout(v int32)`

SetSipMediaTimeout sets SipMediaTimeout field to given value.

### HasSipMediaTimeout

`func (o *GetAlgOpenApiVO) HasSipMediaTimeout() bool`

HasSipMediaTimeout returns a boolean if a field has been set.

### GetSipPorts

`func (o *GetAlgOpenApiVO) GetSipPorts() []int32`

GetSipPorts returns the SipPorts field if non-nil, zero value otherwise.

### GetSipPortsOk

`func (o *GetAlgOpenApiVO) GetSipPortsOk() (*[]int32, bool)`

GetSipPortsOk returns a tuple with the SipPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipPorts

`func (o *GetAlgOpenApiVO) SetSipPorts(v []int32)`

SetSipPorts sets SipPorts field to given value.

### HasSipPorts

`func (o *GetAlgOpenApiVO) HasSipPorts() bool`

HasSipPorts returns a boolean if a field has been set.

### GetSipSignalingTimeout

`func (o *GetAlgOpenApiVO) GetSipSignalingTimeout() int32`

GetSipSignalingTimeout returns the SipSignalingTimeout field if non-nil, zero value otherwise.

### GetSipSignalingTimeoutOk

`func (o *GetAlgOpenApiVO) GetSipSignalingTimeoutOk() (*int32, bool)`

GetSipSignalingTimeoutOk returns a tuple with the SipSignalingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSignalingTimeout

`func (o *GetAlgOpenApiVO) SetSipSignalingTimeout(v int32)`

SetSipSignalingTimeout sets SipSignalingTimeout field to given value.

### HasSipSignalingTimeout

`func (o *GetAlgOpenApiVO) HasSipSignalingTimeout() bool`

HasSipSignalingTimeout returns a boolean if a field has been set.

### GetSipTcp

`func (o *GetAlgOpenApiVO) GetSipTcp() bool`

GetSipTcp returns the SipTcp field if non-nil, zero value otherwise.

### GetSipTcpOk

`func (o *GetAlgOpenApiVO) GetSipTcpOk() (*bool, bool)`

GetSipTcpOk returns a tuple with the SipTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipTcp

`func (o *GetAlgOpenApiVO) SetSipTcp(v bool)`

SetSipTcp sets SipTcp field to given value.

### HasSipTcp

`func (o *GetAlgOpenApiVO) HasSipTcp() bool`

HasSipTcp returns a boolean if a field has been set.

### GetSipTimeout

`func (o *GetAlgOpenApiVO) GetSipTimeout() bool`

GetSipTimeout returns the SipTimeout field if non-nil, zero value otherwise.

### GetSipTimeoutOk

`func (o *GetAlgOpenApiVO) GetSipTimeoutOk() (*bool, bool)`

GetSipTimeoutOk returns a tuple with the SipTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipTimeout

`func (o *GetAlgOpenApiVO) SetSipTimeout(v bool)`

SetSipTimeout sets SipTimeout field to given value.

### HasSipTimeout

`func (o *GetAlgOpenApiVO) HasSipTimeout() bool`

HasSipTimeout returns a boolean if a field has been set.

### GetSipUdp

`func (o *GetAlgOpenApiVO) GetSipUdp() bool`

GetSipUdp returns the SipUdp field if non-nil, zero value otherwise.

### GetSipUdpOk

`func (o *GetAlgOpenApiVO) GetSipUdpOk() (*bool, bool)`

GetSipUdpOk returns a tuple with the SipUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUdp

`func (o *GetAlgOpenApiVO) SetSipUdp(v bool)`

SetSipUdp sets SipUdp field to given value.

### HasSipUdp

`func (o *GetAlgOpenApiVO) HasSipUdp() bool`

HasSipUdp returns a boolean if a field has been set.

### GetSupportSipAlgConfig

`func (o *GetAlgOpenApiVO) GetSupportSipAlgConfig() bool`

GetSupportSipAlgConfig returns the SupportSipAlgConfig field if non-nil, zero value otherwise.

### GetSupportSipAlgConfigOk

`func (o *GetAlgOpenApiVO) GetSupportSipAlgConfigOk() (*bool, bool)`

GetSupportSipAlgConfigOk returns a tuple with the SupportSipAlgConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSipAlgConfig

`func (o *GetAlgOpenApiVO) SetSupportSipAlgConfig(v bool)`

SetSupportSipAlgConfig sets SupportSipAlgConfig field to given value.

### HasSupportSipAlgConfig

`func (o *GetAlgOpenApiVO) HasSupportSipAlgConfig() bool`

HasSupportSipAlgConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


