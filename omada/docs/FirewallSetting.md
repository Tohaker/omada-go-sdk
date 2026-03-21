# FirewallSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BroadcastPing** | **bool** | Broadcast ping of the firewall setting. | 
**Icmp** | **int32** | ICMP should be within the range of 1–2097151. | 
**Other** | **int32** | Other should be within the range of 1–2097151. | 
**ReceiveRedirects** | **bool** | Receive redirects of the firewall setting. | 
**SendRedirects** | **bool** | Send redirects of the firewall setting. | 
**SynCookies** | **bool** | SYN Cookies of the firewall setting. | 
**TcpClose** | **int32** | TCP close should be within the range of 1–2097151. | 
**TcpCloseWait** | **int32** | TCP close wait should be within the range of 1–2097151. | 
**TcpEstablished** | **int32** | TCP established should be within the range of 1–2097151. | 
**TcpFinWait** | **int32** | TCP FIN wait should be within the range of 1–2097151. | 
**TcpLastAck** | **int32** | TCP last ACK should be within the range of 1–2097151. | 
**TcpSynReceive** | **int32** | TCP SYN receive should be within the range of 1–2097151. | 
**TcpSynSent** | **int32** | TCP SYN sent should be within the range of 1–2097151. | 
**TcpTimeWait** | **int32** | TCP time wait should be within the range of 1–2097151. | 
**UdpOther** | **int32** | UDP other should be within the range of 1–2097151. | 
**UdpStream** | **int32** | UDP stream should be within the range of 1–2097151. | 

## Methods

### NewFirewallSetting

`func NewFirewallSetting(broadcastPing bool, icmp int32, other int32, receiveRedirects bool, sendRedirects bool, synCookies bool, tcpClose int32, tcpCloseWait int32, tcpEstablished int32, tcpFinWait int32, tcpLastAck int32, tcpSynReceive int32, tcpSynSent int32, tcpTimeWait int32, udpOther int32, udpStream int32, ) *FirewallSetting`

NewFirewallSetting instantiates a new FirewallSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSettingWithDefaults

`func NewFirewallSettingWithDefaults() *FirewallSetting`

NewFirewallSettingWithDefaults instantiates a new FirewallSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroadcastPing

`func (o *FirewallSetting) GetBroadcastPing() bool`

GetBroadcastPing returns the BroadcastPing field if non-nil, zero value otherwise.

### GetBroadcastPingOk

`func (o *FirewallSetting) GetBroadcastPingOk() (*bool, bool)`

GetBroadcastPingOk returns a tuple with the BroadcastPing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastPing

`func (o *FirewallSetting) SetBroadcastPing(v bool)`

SetBroadcastPing sets BroadcastPing field to given value.


### GetIcmp

`func (o *FirewallSetting) GetIcmp() int32`

GetIcmp returns the Icmp field if non-nil, zero value otherwise.

### GetIcmpOk

`func (o *FirewallSetting) GetIcmpOk() (*int32, bool)`

GetIcmpOk returns a tuple with the Icmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcmp

`func (o *FirewallSetting) SetIcmp(v int32)`

SetIcmp sets Icmp field to given value.


### GetOther

`func (o *FirewallSetting) GetOther() int32`

GetOther returns the Other field if non-nil, zero value otherwise.

### GetOtherOk

`func (o *FirewallSetting) GetOtherOk() (*int32, bool)`

GetOtherOk returns a tuple with the Other field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther

`func (o *FirewallSetting) SetOther(v int32)`

SetOther sets Other field to given value.


### GetReceiveRedirects

`func (o *FirewallSetting) GetReceiveRedirects() bool`

GetReceiveRedirects returns the ReceiveRedirects field if non-nil, zero value otherwise.

### GetReceiveRedirectsOk

`func (o *FirewallSetting) GetReceiveRedirectsOk() (*bool, bool)`

GetReceiveRedirectsOk returns a tuple with the ReceiveRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveRedirects

`func (o *FirewallSetting) SetReceiveRedirects(v bool)`

SetReceiveRedirects sets ReceiveRedirects field to given value.


### GetSendRedirects

`func (o *FirewallSetting) GetSendRedirects() bool`

GetSendRedirects returns the SendRedirects field if non-nil, zero value otherwise.

### GetSendRedirectsOk

`func (o *FirewallSetting) GetSendRedirectsOk() (*bool, bool)`

GetSendRedirectsOk returns a tuple with the SendRedirects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendRedirects

`func (o *FirewallSetting) SetSendRedirects(v bool)`

SetSendRedirects sets SendRedirects field to given value.


### GetSynCookies

`func (o *FirewallSetting) GetSynCookies() bool`

GetSynCookies returns the SynCookies field if non-nil, zero value otherwise.

### GetSynCookiesOk

`func (o *FirewallSetting) GetSynCookiesOk() (*bool, bool)`

GetSynCookiesOk returns a tuple with the SynCookies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynCookies

`func (o *FirewallSetting) SetSynCookies(v bool)`

SetSynCookies sets SynCookies field to given value.


### GetTcpClose

`func (o *FirewallSetting) GetTcpClose() int32`

GetTcpClose returns the TcpClose field if non-nil, zero value otherwise.

### GetTcpCloseOk

`func (o *FirewallSetting) GetTcpCloseOk() (*int32, bool)`

GetTcpCloseOk returns a tuple with the TcpClose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpClose

`func (o *FirewallSetting) SetTcpClose(v int32)`

SetTcpClose sets TcpClose field to given value.


### GetTcpCloseWait

`func (o *FirewallSetting) GetTcpCloseWait() int32`

GetTcpCloseWait returns the TcpCloseWait field if non-nil, zero value otherwise.

### GetTcpCloseWaitOk

`func (o *FirewallSetting) GetTcpCloseWaitOk() (*int32, bool)`

GetTcpCloseWaitOk returns a tuple with the TcpCloseWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpCloseWait

`func (o *FirewallSetting) SetTcpCloseWait(v int32)`

SetTcpCloseWait sets TcpCloseWait field to given value.


### GetTcpEstablished

`func (o *FirewallSetting) GetTcpEstablished() int32`

GetTcpEstablished returns the TcpEstablished field if non-nil, zero value otherwise.

### GetTcpEstablishedOk

`func (o *FirewallSetting) GetTcpEstablishedOk() (*int32, bool)`

GetTcpEstablishedOk returns a tuple with the TcpEstablished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpEstablished

`func (o *FirewallSetting) SetTcpEstablished(v int32)`

SetTcpEstablished sets TcpEstablished field to given value.


### GetTcpFinWait

`func (o *FirewallSetting) GetTcpFinWait() int32`

GetTcpFinWait returns the TcpFinWait field if non-nil, zero value otherwise.

### GetTcpFinWaitOk

`func (o *FirewallSetting) GetTcpFinWaitOk() (*int32, bool)`

GetTcpFinWaitOk returns a tuple with the TcpFinWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpFinWait

`func (o *FirewallSetting) SetTcpFinWait(v int32)`

SetTcpFinWait sets TcpFinWait field to given value.


### GetTcpLastAck

`func (o *FirewallSetting) GetTcpLastAck() int32`

GetTcpLastAck returns the TcpLastAck field if non-nil, zero value otherwise.

### GetTcpLastAckOk

`func (o *FirewallSetting) GetTcpLastAckOk() (*int32, bool)`

GetTcpLastAckOk returns a tuple with the TcpLastAck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpLastAck

`func (o *FirewallSetting) SetTcpLastAck(v int32)`

SetTcpLastAck sets TcpLastAck field to given value.


### GetTcpSynReceive

`func (o *FirewallSetting) GetTcpSynReceive() int32`

GetTcpSynReceive returns the TcpSynReceive field if non-nil, zero value otherwise.

### GetTcpSynReceiveOk

`func (o *FirewallSetting) GetTcpSynReceiveOk() (*int32, bool)`

GetTcpSynReceiveOk returns a tuple with the TcpSynReceive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSynReceive

`func (o *FirewallSetting) SetTcpSynReceive(v int32)`

SetTcpSynReceive sets TcpSynReceive field to given value.


### GetTcpSynSent

`func (o *FirewallSetting) GetTcpSynSent() int32`

GetTcpSynSent returns the TcpSynSent field if non-nil, zero value otherwise.

### GetTcpSynSentOk

`func (o *FirewallSetting) GetTcpSynSentOk() (*int32, bool)`

GetTcpSynSentOk returns a tuple with the TcpSynSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpSynSent

`func (o *FirewallSetting) SetTcpSynSent(v int32)`

SetTcpSynSent sets TcpSynSent field to given value.


### GetTcpTimeWait

`func (o *FirewallSetting) GetTcpTimeWait() int32`

GetTcpTimeWait returns the TcpTimeWait field if non-nil, zero value otherwise.

### GetTcpTimeWaitOk

`func (o *FirewallSetting) GetTcpTimeWaitOk() (*int32, bool)`

GetTcpTimeWaitOk returns a tuple with the TcpTimeWait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpTimeWait

`func (o *FirewallSetting) SetTcpTimeWait(v int32)`

SetTcpTimeWait sets TcpTimeWait field to given value.


### GetUdpOther

`func (o *FirewallSetting) GetUdpOther() int32`

GetUdpOther returns the UdpOther field if non-nil, zero value otherwise.

### GetUdpOtherOk

`func (o *FirewallSetting) GetUdpOtherOk() (*int32, bool)`

GetUdpOtherOk returns a tuple with the UdpOther field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpOther

`func (o *FirewallSetting) SetUdpOther(v int32)`

SetUdpOther sets UdpOther field to given value.


### GetUdpStream

`func (o *FirewallSetting) GetUdpStream() int32`

GetUdpStream returns the UdpStream field if non-nil, zero value otherwise.

### GetUdpStreamOk

`func (o *FirewallSetting) GetUdpStreamOk() (*int32, bool)`

GetUdpStreamOk returns a tuple with the UdpStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpStream

`func (o *FirewallSetting) SetUdpStream(v int32)`

SetUdpStream sets UdpStream field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


