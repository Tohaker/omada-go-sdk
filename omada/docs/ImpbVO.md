# ImpbVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **string** | clientName | [optional] 
**Id** | Pointer to **string** | id | [optional] 
**Ip** | Pointer to **string** | ip | [optional] 
**Mac** | Pointer to **string** | mac | [optional] 
**Port** | Pointer to [**PortVO**](PortVO.md) |  | [optional] 
**Vlan** | Pointer to **int32** | vlan | [optional] 

## Methods

### NewImpbVO

`func NewImpbVO() *ImpbVO`

NewImpbVO instantiates a new ImpbVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImpbVOWithDefaults

`func NewImpbVOWithDefaults() *ImpbVO`

NewImpbVOWithDefaults instantiates a new ImpbVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ImpbVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ImpbVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ImpbVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ImpbVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetId

`func (o *ImpbVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImpbVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImpbVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ImpbVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *ImpbVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ImpbVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ImpbVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ImpbVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *ImpbVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ImpbVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ImpbVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ImpbVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPort

`func (o *ImpbVO) GetPort() PortVO`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ImpbVO) GetPortOk() (*PortVO, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ImpbVO) SetPort(v PortVO)`

SetPort sets Port field to given value.

### HasPort

`func (o *ImpbVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetVlan

`func (o *ImpbVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *ImpbVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *ImpbVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *ImpbVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


