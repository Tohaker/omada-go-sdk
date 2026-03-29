# OsgWanPortIpv6ConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **int32** | IPv6 enable should be a value as follows: 0: disconnected; 1: connected | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**InternetState** | Pointer to **int32** | It should be a value as follows: 0: disconnected; 1: connected | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**PriDns** | Pointer to **string** |  | [optional] 
**Proto** | Pointer to **string** |  | [optional] 
**SndDns** | Pointer to **string** |  | [optional] 

## Methods

### NewOsgWanPortIpv6ConfigVO

`func NewOsgWanPortIpv6ConfigVO() *OsgWanPortIpv6ConfigVO`

NewOsgWanPortIpv6ConfigVO instantiates a new OsgWanPortIpv6ConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgWanPortIpv6ConfigVOWithDefaults

`func NewOsgWanPortIpv6ConfigVOWithDefaults() *OsgWanPortIpv6ConfigVO`

NewOsgWanPortIpv6ConfigVOWithDefaults instantiates a new OsgWanPortIpv6ConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *OsgWanPortIpv6ConfigVO) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *OsgWanPortIpv6ConfigVO) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *OsgWanPortIpv6ConfigVO) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *OsgWanPortIpv6ConfigVO) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetEnable

`func (o *OsgWanPortIpv6ConfigVO) GetEnable() int32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OsgWanPortIpv6ConfigVO) GetEnableOk() (*int32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OsgWanPortIpv6ConfigVO) SetEnable(v int32)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OsgWanPortIpv6ConfigVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGateway

`func (o *OsgWanPortIpv6ConfigVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *OsgWanPortIpv6ConfigVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *OsgWanPortIpv6ConfigVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *OsgWanPortIpv6ConfigVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInternetState

`func (o *OsgWanPortIpv6ConfigVO) GetInternetState() int32`

GetInternetState returns the InternetState field if non-nil, zero value otherwise.

### GetInternetStateOk

`func (o *OsgWanPortIpv6ConfigVO) GetInternetStateOk() (*int32, bool)`

GetInternetStateOk returns a tuple with the InternetState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetState

`func (o *OsgWanPortIpv6ConfigVO) SetInternetState(v int32)`

SetInternetState sets InternetState field to given value.

### HasInternetState

`func (o *OsgWanPortIpv6ConfigVO) HasInternetState() bool`

HasInternetState returns a boolean if a field has been set.

### GetMac

`func (o *OsgWanPortIpv6ConfigVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OsgWanPortIpv6ConfigVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OsgWanPortIpv6ConfigVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OsgWanPortIpv6ConfigVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPriDns

`func (o *OsgWanPortIpv6ConfigVO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *OsgWanPortIpv6ConfigVO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *OsgWanPortIpv6ConfigVO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *OsgWanPortIpv6ConfigVO) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetProto

`func (o *OsgWanPortIpv6ConfigVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *OsgWanPortIpv6ConfigVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *OsgWanPortIpv6ConfigVO) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *OsgWanPortIpv6ConfigVO) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetSndDns

`func (o *OsgWanPortIpv6ConfigVO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *OsgWanPortIpv6ConfigVO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *OsgWanPortIpv6ConfigVO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *OsgWanPortIpv6ConfigVO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


