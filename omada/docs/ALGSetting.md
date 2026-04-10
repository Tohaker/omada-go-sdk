# ALGSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

## Methods

### NewALGSetting

`func NewALGSetting(ftp bool, h323 bool, ipSec bool, pptp bool, sip bool, ) *ALGSetting`

NewALGSetting instantiates a new ALGSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewALGSettingWithDefaults

`func NewALGSettingWithDefaults() *ALGSetting`

NewALGSettingWithDefaults instantiates a new ALGSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFtp

`func (o *ALGSetting) GetFtp() bool`

GetFtp returns the Ftp field if non-nil, zero value otherwise.

### GetFtpOk

`func (o *ALGSetting) GetFtpOk() (*bool, bool)`

GetFtpOk returns a tuple with the Ftp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtp

`func (o *ALGSetting) SetFtp(v bool)`

SetFtp sets Ftp field to given value.


### GetFtpPorts

`func (o *ALGSetting) GetFtpPorts() []int32`

GetFtpPorts returns the FtpPorts field if non-nil, zero value otherwise.

### GetFtpPortsOk

`func (o *ALGSetting) GetFtpPortsOk() (*[]int32, bool)`

GetFtpPortsOk returns a tuple with the FtpPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFtpPorts

`func (o *ALGSetting) SetFtpPorts(v []int32)`

SetFtpPorts sets FtpPorts field to given value.

### HasFtpPorts

`func (o *ALGSetting) HasFtpPorts() bool`

HasFtpPorts returns a boolean if a field has been set.

### GetH323

`func (o *ALGSetting) GetH323() bool`

GetH323 returns the H323 field if non-nil, zero value otherwise.

### GetH323Ok

`func (o *ALGSetting) GetH323Ok() (*bool, bool)`

GetH323Ok returns a tuple with the H323 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetH323

`func (o *ALGSetting) SetH323(v bool)`

SetH323 sets H323 field to given value.


### GetIpSec

`func (o *ALGSetting) GetIpSec() bool`

GetIpSec returns the IpSec field if non-nil, zero value otherwise.

### GetIpSecOk

`func (o *ALGSetting) GetIpSecOk() (*bool, bool)`

GetIpSecOk returns a tuple with the IpSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSec

`func (o *ALGSetting) SetIpSec(v bool)`

SetIpSec sets IpSec field to given value.


### GetPptp

`func (o *ALGSetting) GetPptp() bool`

GetPptp returns the Pptp field if non-nil, zero value otherwise.

### GetPptpOk

`func (o *ALGSetting) GetPptpOk() (*bool, bool)`

GetPptpOk returns a tuple with the Pptp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPptp

`func (o *ALGSetting) SetPptp(v bool)`

SetPptp sets Pptp field to given value.


### GetSip

`func (o *ALGSetting) GetSip() bool`

GetSip returns the Sip field if non-nil, zero value otherwise.

### GetSipOk

`func (o *ALGSetting) GetSipOk() (*bool, bool)`

GetSipOk returns a tuple with the Sip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSip

`func (o *ALGSetting) SetSip(v bool)`

SetSip sets Sip field to given value.


### GetSipDirectMedia

`func (o *ALGSetting) GetSipDirectMedia() bool`

GetSipDirectMedia returns the SipDirectMedia field if non-nil, zero value otherwise.

### GetSipDirectMediaOk

`func (o *ALGSetting) GetSipDirectMediaOk() (*bool, bool)`

GetSipDirectMediaOk returns a tuple with the SipDirectMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipDirectMedia

`func (o *ALGSetting) SetSipDirectMedia(v bool)`

SetSipDirectMedia sets SipDirectMedia field to given value.

### HasSipDirectMedia

`func (o *ALGSetting) HasSipDirectMedia() bool`

HasSipDirectMedia returns a boolean if a field has been set.

### GetSipDirectSignaling

`func (o *ALGSetting) GetSipDirectSignaling() bool`

GetSipDirectSignaling returns the SipDirectSignaling field if non-nil, zero value otherwise.

### GetSipDirectSignalingOk

`func (o *ALGSetting) GetSipDirectSignalingOk() (*bool, bool)`

GetSipDirectSignalingOk returns a tuple with the SipDirectSignaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipDirectSignaling

`func (o *ALGSetting) SetSipDirectSignaling(v bool)`

SetSipDirectSignaling sets SipDirectSignaling field to given value.

### HasSipDirectSignaling

`func (o *ALGSetting) HasSipDirectSignaling() bool`

HasSipDirectSignaling returns a boolean if a field has been set.

### GetSipMediaTimeout

`func (o *ALGSetting) GetSipMediaTimeout() int32`

GetSipMediaTimeout returns the SipMediaTimeout field if non-nil, zero value otherwise.

### GetSipMediaTimeoutOk

`func (o *ALGSetting) GetSipMediaTimeoutOk() (*int32, bool)`

GetSipMediaTimeoutOk returns a tuple with the SipMediaTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipMediaTimeout

`func (o *ALGSetting) SetSipMediaTimeout(v int32)`

SetSipMediaTimeout sets SipMediaTimeout field to given value.

### HasSipMediaTimeout

`func (o *ALGSetting) HasSipMediaTimeout() bool`

HasSipMediaTimeout returns a boolean if a field has been set.

### GetSipPorts

`func (o *ALGSetting) GetSipPorts() []int32`

GetSipPorts returns the SipPorts field if non-nil, zero value otherwise.

### GetSipPortsOk

`func (o *ALGSetting) GetSipPortsOk() (*[]int32, bool)`

GetSipPortsOk returns a tuple with the SipPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipPorts

`func (o *ALGSetting) SetSipPorts(v []int32)`

SetSipPorts sets SipPorts field to given value.

### HasSipPorts

`func (o *ALGSetting) HasSipPorts() bool`

HasSipPorts returns a boolean if a field has been set.

### GetSipSignalingTimeout

`func (o *ALGSetting) GetSipSignalingTimeout() int32`

GetSipSignalingTimeout returns the SipSignalingTimeout field if non-nil, zero value otherwise.

### GetSipSignalingTimeoutOk

`func (o *ALGSetting) GetSipSignalingTimeoutOk() (*int32, bool)`

GetSipSignalingTimeoutOk returns a tuple with the SipSignalingTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipSignalingTimeout

`func (o *ALGSetting) SetSipSignalingTimeout(v int32)`

SetSipSignalingTimeout sets SipSignalingTimeout field to given value.

### HasSipSignalingTimeout

`func (o *ALGSetting) HasSipSignalingTimeout() bool`

HasSipSignalingTimeout returns a boolean if a field has been set.

### GetSipTcp

`func (o *ALGSetting) GetSipTcp() bool`

GetSipTcp returns the SipTcp field if non-nil, zero value otherwise.

### GetSipTcpOk

`func (o *ALGSetting) GetSipTcpOk() (*bool, bool)`

GetSipTcpOk returns a tuple with the SipTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipTcp

`func (o *ALGSetting) SetSipTcp(v bool)`

SetSipTcp sets SipTcp field to given value.

### HasSipTcp

`func (o *ALGSetting) HasSipTcp() bool`

HasSipTcp returns a boolean if a field has been set.

### GetSipTimeout

`func (o *ALGSetting) GetSipTimeout() bool`

GetSipTimeout returns the SipTimeout field if non-nil, zero value otherwise.

### GetSipTimeoutOk

`func (o *ALGSetting) GetSipTimeoutOk() (*bool, bool)`

GetSipTimeoutOk returns a tuple with the SipTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipTimeout

`func (o *ALGSetting) SetSipTimeout(v bool)`

SetSipTimeout sets SipTimeout field to given value.

### HasSipTimeout

`func (o *ALGSetting) HasSipTimeout() bool`

HasSipTimeout returns a boolean if a field has been set.

### GetSipUdp

`func (o *ALGSetting) GetSipUdp() bool`

GetSipUdp returns the SipUdp field if non-nil, zero value otherwise.

### GetSipUdpOk

`func (o *ALGSetting) GetSipUdpOk() (*bool, bool)`

GetSipUdpOk returns a tuple with the SipUdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUdp

`func (o *ALGSetting) SetSipUdp(v bool)`

SetSipUdp sets SipUdp field to given value.

### HasSipUdp

`func (o *ALGSetting) HasSipUdp() bool`

HasSipUdp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


