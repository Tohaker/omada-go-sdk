# OsgLanStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientNum** | Pointer to **int32** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**LanName** | Pointer to **string** |  | [optional] 
**LanPortIpv6Config** | Pointer to [**OsgLanPortIpv6ConfigVO**](OsgLanPortIpv6ConfigVO.md) |  | [optional] 
**Rx** | Pointer to **int64** |  | [optional] 
**Tx** | Pointer to **int64** |  | [optional] 
**Vlan** | Pointer to **int32** |  | [optional] 

## Methods

### NewOsgLanStatVO

`func NewOsgLanStatVO() *OsgLanStatVO`

NewOsgLanStatVO instantiates a new OsgLanStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgLanStatVOWithDefaults

`func NewOsgLanStatVOWithDefaults() *OsgLanStatVO`

NewOsgLanStatVOWithDefaults instantiates a new OsgLanStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientNum

`func (o *OsgLanStatVO) GetClientNum() int32`

GetClientNum returns the ClientNum field if non-nil, zero value otherwise.

### GetClientNumOk

`func (o *OsgLanStatVO) GetClientNumOk() (*int32, bool)`

GetClientNumOk returns a tuple with the ClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNum

`func (o *OsgLanStatVO) SetClientNum(v int32)`

SetClientNum sets ClientNum field to given value.

### HasClientNum

`func (o *OsgLanStatVO) HasClientNum() bool`

HasClientNum returns a boolean if a field has been set.

### GetIp

`func (o *OsgLanStatVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OsgLanStatVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OsgLanStatVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OsgLanStatVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLanName

`func (o *OsgLanStatVO) GetLanName() string`

GetLanName returns the LanName field if non-nil, zero value otherwise.

### GetLanNameOk

`func (o *OsgLanStatVO) GetLanNameOk() (*string, bool)`

GetLanNameOk returns a tuple with the LanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanName

`func (o *OsgLanStatVO) SetLanName(v string)`

SetLanName sets LanName field to given value.

### HasLanName

`func (o *OsgLanStatVO) HasLanName() bool`

HasLanName returns a boolean if a field has been set.

### GetLanPortIpv6Config

`func (o *OsgLanStatVO) GetLanPortIpv6Config() OsgLanPortIpv6ConfigVO`

GetLanPortIpv6Config returns the LanPortIpv6Config field if non-nil, zero value otherwise.

### GetLanPortIpv6ConfigOk

`func (o *OsgLanStatVO) GetLanPortIpv6ConfigOk() (*OsgLanPortIpv6ConfigVO, bool)`

GetLanPortIpv6ConfigOk returns a tuple with the LanPortIpv6Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPortIpv6Config

`func (o *OsgLanStatVO) SetLanPortIpv6Config(v OsgLanPortIpv6ConfigVO)`

SetLanPortIpv6Config sets LanPortIpv6Config field to given value.

### HasLanPortIpv6Config

`func (o *OsgLanStatVO) HasLanPortIpv6Config() bool`

HasLanPortIpv6Config returns a boolean if a field has been set.

### GetRx

`func (o *OsgLanStatVO) GetRx() int64`

GetRx returns the Rx field if non-nil, zero value otherwise.

### GetRxOk

`func (o *OsgLanStatVO) GetRxOk() (*int64, bool)`

GetRxOk returns a tuple with the Rx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRx

`func (o *OsgLanStatVO) SetRx(v int64)`

SetRx sets Rx field to given value.

### HasRx

`func (o *OsgLanStatVO) HasRx() bool`

HasRx returns a boolean if a field has been set.

### GetTx

`func (o *OsgLanStatVO) GetTx() int64`

GetTx returns the Tx field if non-nil, zero value otherwise.

### GetTxOk

`func (o *OsgLanStatVO) GetTxOk() (*int64, bool)`

GetTxOk returns a tuple with the Tx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTx

`func (o *OsgLanStatVO) SetTx(v int64)`

SetTx sets Tx field to given value.

### HasTx

`func (o *OsgLanStatVO) HasTx() bool`

HasTx returns a boolean if a field has been set.

### GetVlan

`func (o *OsgLanStatVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OsgLanStatVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OsgLanStatVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *OsgLanStatVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


