# RadiusProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcctServer** | Pointer to [**[]RadiusAcctServerOpenApiVO**](RadiusAcctServerOpenApiVO.md) | Radius accounting server list, valid when parameter [radiusAccountingEnable] is true | [optional] 
**AuthServer** | Pointer to [**[]RadiusAuthServerOpenApiVO**](RadiusAuthServerOpenApiVO.md) | Radius authentication server list | [optional] 
**BuiltInServer** | Pointer to **bool** | Is this Radius server a built-in server | [optional] 
**BuiltInServerSecret** | Pointer to **string** | Built-in Radius server secret, valid when parameter [builtInServer] is true | [optional] 
**CoaEnable** | Pointer to **bool** | Radius CoA enable status | [optional] 
**CoaPassword** | Pointer to **string** | Radius CoA password, valid when parameter [coaEnable] is true | [optional] 
**CustomIp** | Pointer to **string** | Built-in Radius server custom IP, valid when parameter [builtInServer] is true and [ipType] is 1 | [optional] 
**DomainEnable** | Pointer to **bool** | Domain enable status | [optional] 
**InterimUpdateEnable** | Pointer to **bool** | When radius accounting enables, interval update enable status | [optional] 
**InterimUpdateInterval** | Pointer to **int32** | When interval update enables, interval update duration | [optional] 
**IpType** | Pointer to **int32** | Built-in Radius server IP type, 0: auto，1: custom, valid when parameter [builtInServer] is true | [optional] 
**Name** | Pointer to **string** | Radius profile name | [optional] 
**RadiusAccountingEnable** | Pointer to **bool** | Radius accounting enable status | [optional] 
**RadiusProfileId** | Pointer to **string** | Radius profile ID | [optional] 
**ServerEnable** | Pointer to **bool** | Built-in Radius server enable status, valid when parameter [builtInServer] is true | [optional] 
**TunnelReplyEnable** | Pointer to **bool** | Built-in Radius server tunneled reply enable status, valid when parameter [builtInServer] is true | [optional] 
**WirelessVlanAssignment** | Pointer to **bool** | VLAN assignment for wireless network enable status | [optional] 

## Methods

### NewRadiusProfileOpenApiVO

`func NewRadiusProfileOpenApiVO() *RadiusProfileOpenApiVO`

NewRadiusProfileOpenApiVO instantiates a new RadiusProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusProfileOpenApiVOWithDefaults

`func NewRadiusProfileOpenApiVOWithDefaults() *RadiusProfileOpenApiVO`

NewRadiusProfileOpenApiVOWithDefaults instantiates a new RadiusProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcctServer

`func (o *RadiusProfileOpenApiVO) GetAcctServer() []RadiusAcctServerOpenApiVO`

GetAcctServer returns the AcctServer field if non-nil, zero value otherwise.

### GetAcctServerOk

`func (o *RadiusProfileOpenApiVO) GetAcctServerOk() (*[]RadiusAcctServerOpenApiVO, bool)`

GetAcctServerOk returns a tuple with the AcctServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcctServer

`func (o *RadiusProfileOpenApiVO) SetAcctServer(v []RadiusAcctServerOpenApiVO)`

SetAcctServer sets AcctServer field to given value.

### HasAcctServer

`func (o *RadiusProfileOpenApiVO) HasAcctServer() bool`

HasAcctServer returns a boolean if a field has been set.

### GetAuthServer

`func (o *RadiusProfileOpenApiVO) GetAuthServer() []RadiusAuthServerOpenApiVO`

GetAuthServer returns the AuthServer field if non-nil, zero value otherwise.

### GetAuthServerOk

`func (o *RadiusProfileOpenApiVO) GetAuthServerOk() (*[]RadiusAuthServerOpenApiVO, bool)`

GetAuthServerOk returns a tuple with the AuthServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthServer

`func (o *RadiusProfileOpenApiVO) SetAuthServer(v []RadiusAuthServerOpenApiVO)`

SetAuthServer sets AuthServer field to given value.

### HasAuthServer

`func (o *RadiusProfileOpenApiVO) HasAuthServer() bool`

HasAuthServer returns a boolean if a field has been set.

### GetBuiltInServer

`func (o *RadiusProfileOpenApiVO) GetBuiltInServer() bool`

GetBuiltInServer returns the BuiltInServer field if non-nil, zero value otherwise.

### GetBuiltInServerOk

`func (o *RadiusProfileOpenApiVO) GetBuiltInServerOk() (*bool, bool)`

GetBuiltInServerOk returns a tuple with the BuiltInServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInServer

`func (o *RadiusProfileOpenApiVO) SetBuiltInServer(v bool)`

SetBuiltInServer sets BuiltInServer field to given value.

### HasBuiltInServer

`func (o *RadiusProfileOpenApiVO) HasBuiltInServer() bool`

HasBuiltInServer returns a boolean if a field has been set.

### GetBuiltInServerSecret

`func (o *RadiusProfileOpenApiVO) GetBuiltInServerSecret() string`

GetBuiltInServerSecret returns the BuiltInServerSecret field if non-nil, zero value otherwise.

### GetBuiltInServerSecretOk

`func (o *RadiusProfileOpenApiVO) GetBuiltInServerSecretOk() (*string, bool)`

GetBuiltInServerSecretOk returns a tuple with the BuiltInServerSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInServerSecret

`func (o *RadiusProfileOpenApiVO) SetBuiltInServerSecret(v string)`

SetBuiltInServerSecret sets BuiltInServerSecret field to given value.

### HasBuiltInServerSecret

`func (o *RadiusProfileOpenApiVO) HasBuiltInServerSecret() bool`

HasBuiltInServerSecret returns a boolean if a field has been set.

### GetCoaEnable

`func (o *RadiusProfileOpenApiVO) GetCoaEnable() bool`

GetCoaEnable returns the CoaEnable field if non-nil, zero value otherwise.

### GetCoaEnableOk

`func (o *RadiusProfileOpenApiVO) GetCoaEnableOk() (*bool, bool)`

GetCoaEnableOk returns a tuple with the CoaEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaEnable

`func (o *RadiusProfileOpenApiVO) SetCoaEnable(v bool)`

SetCoaEnable sets CoaEnable field to given value.

### HasCoaEnable

`func (o *RadiusProfileOpenApiVO) HasCoaEnable() bool`

HasCoaEnable returns a boolean if a field has been set.

### GetCoaPassword

`func (o *RadiusProfileOpenApiVO) GetCoaPassword() string`

GetCoaPassword returns the CoaPassword field if non-nil, zero value otherwise.

### GetCoaPasswordOk

`func (o *RadiusProfileOpenApiVO) GetCoaPasswordOk() (*string, bool)`

GetCoaPasswordOk returns a tuple with the CoaPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoaPassword

`func (o *RadiusProfileOpenApiVO) SetCoaPassword(v string)`

SetCoaPassword sets CoaPassword field to given value.

### HasCoaPassword

`func (o *RadiusProfileOpenApiVO) HasCoaPassword() bool`

HasCoaPassword returns a boolean if a field has been set.

### GetCustomIp

`func (o *RadiusProfileOpenApiVO) GetCustomIp() string`

GetCustomIp returns the CustomIp field if non-nil, zero value otherwise.

### GetCustomIpOk

`func (o *RadiusProfileOpenApiVO) GetCustomIpOk() (*string, bool)`

GetCustomIpOk returns a tuple with the CustomIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomIp

`func (o *RadiusProfileOpenApiVO) SetCustomIp(v string)`

SetCustomIp sets CustomIp field to given value.

### HasCustomIp

`func (o *RadiusProfileOpenApiVO) HasCustomIp() bool`

HasCustomIp returns a boolean if a field has been set.

### GetDomainEnable

`func (o *RadiusProfileOpenApiVO) GetDomainEnable() bool`

GetDomainEnable returns the DomainEnable field if non-nil, zero value otherwise.

### GetDomainEnableOk

`func (o *RadiusProfileOpenApiVO) GetDomainEnableOk() (*bool, bool)`

GetDomainEnableOk returns a tuple with the DomainEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEnable

`func (o *RadiusProfileOpenApiVO) SetDomainEnable(v bool)`

SetDomainEnable sets DomainEnable field to given value.

### HasDomainEnable

`func (o *RadiusProfileOpenApiVO) HasDomainEnable() bool`

HasDomainEnable returns a boolean if a field has been set.

### GetInterimUpdateEnable

`func (o *RadiusProfileOpenApiVO) GetInterimUpdateEnable() bool`

GetInterimUpdateEnable returns the InterimUpdateEnable field if non-nil, zero value otherwise.

### GetInterimUpdateEnableOk

`func (o *RadiusProfileOpenApiVO) GetInterimUpdateEnableOk() (*bool, bool)`

GetInterimUpdateEnableOk returns a tuple with the InterimUpdateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimUpdateEnable

`func (o *RadiusProfileOpenApiVO) SetInterimUpdateEnable(v bool)`

SetInterimUpdateEnable sets InterimUpdateEnable field to given value.

### HasInterimUpdateEnable

`func (o *RadiusProfileOpenApiVO) HasInterimUpdateEnable() bool`

HasInterimUpdateEnable returns a boolean if a field has been set.

### GetInterimUpdateInterval

`func (o *RadiusProfileOpenApiVO) GetInterimUpdateInterval() int32`

GetInterimUpdateInterval returns the InterimUpdateInterval field if non-nil, zero value otherwise.

### GetInterimUpdateIntervalOk

`func (o *RadiusProfileOpenApiVO) GetInterimUpdateIntervalOk() (*int32, bool)`

GetInterimUpdateIntervalOk returns a tuple with the InterimUpdateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterimUpdateInterval

`func (o *RadiusProfileOpenApiVO) SetInterimUpdateInterval(v int32)`

SetInterimUpdateInterval sets InterimUpdateInterval field to given value.

### HasInterimUpdateInterval

`func (o *RadiusProfileOpenApiVO) HasInterimUpdateInterval() bool`

HasInterimUpdateInterval returns a boolean if a field has been set.

### GetIpType

`func (o *RadiusProfileOpenApiVO) GetIpType() int32`

GetIpType returns the IpType field if non-nil, zero value otherwise.

### GetIpTypeOk

`func (o *RadiusProfileOpenApiVO) GetIpTypeOk() (*int32, bool)`

GetIpTypeOk returns a tuple with the IpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpType

`func (o *RadiusProfileOpenApiVO) SetIpType(v int32)`

SetIpType sets IpType field to given value.

### HasIpType

`func (o *RadiusProfileOpenApiVO) HasIpType() bool`

HasIpType returns a boolean if a field has been set.

### GetName

`func (o *RadiusProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadiusProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadiusProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RadiusProfileOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRadiusAccountingEnable

`func (o *RadiusProfileOpenApiVO) GetRadiusAccountingEnable() bool`

GetRadiusAccountingEnable returns the RadiusAccountingEnable field if non-nil, zero value otherwise.

### GetRadiusAccountingEnableOk

`func (o *RadiusProfileOpenApiVO) GetRadiusAccountingEnableOk() (*bool, bool)`

GetRadiusAccountingEnableOk returns a tuple with the RadiusAccountingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusAccountingEnable

`func (o *RadiusProfileOpenApiVO) SetRadiusAccountingEnable(v bool)`

SetRadiusAccountingEnable sets RadiusAccountingEnable field to given value.

### HasRadiusAccountingEnable

`func (o *RadiusProfileOpenApiVO) HasRadiusAccountingEnable() bool`

HasRadiusAccountingEnable returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *RadiusProfileOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *RadiusProfileOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *RadiusProfileOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *RadiusProfileOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.

### GetServerEnable

`func (o *RadiusProfileOpenApiVO) GetServerEnable() bool`

GetServerEnable returns the ServerEnable field if non-nil, zero value otherwise.

### GetServerEnableOk

`func (o *RadiusProfileOpenApiVO) GetServerEnableOk() (*bool, bool)`

GetServerEnableOk returns a tuple with the ServerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerEnable

`func (o *RadiusProfileOpenApiVO) SetServerEnable(v bool)`

SetServerEnable sets ServerEnable field to given value.

### HasServerEnable

`func (o *RadiusProfileOpenApiVO) HasServerEnable() bool`

HasServerEnable returns a boolean if a field has been set.

### GetTunnelReplyEnable

`func (o *RadiusProfileOpenApiVO) GetTunnelReplyEnable() bool`

GetTunnelReplyEnable returns the TunnelReplyEnable field if non-nil, zero value otherwise.

### GetTunnelReplyEnableOk

`func (o *RadiusProfileOpenApiVO) GetTunnelReplyEnableOk() (*bool, bool)`

GetTunnelReplyEnableOk returns a tuple with the TunnelReplyEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTunnelReplyEnable

`func (o *RadiusProfileOpenApiVO) SetTunnelReplyEnable(v bool)`

SetTunnelReplyEnable sets TunnelReplyEnable field to given value.

### HasTunnelReplyEnable

`func (o *RadiusProfileOpenApiVO) HasTunnelReplyEnable() bool`

HasTunnelReplyEnable returns a boolean if a field has been set.

### GetWirelessVlanAssignment

`func (o *RadiusProfileOpenApiVO) GetWirelessVlanAssignment() bool`

GetWirelessVlanAssignment returns the WirelessVlanAssignment field if non-nil, zero value otherwise.

### GetWirelessVlanAssignmentOk

`func (o *RadiusProfileOpenApiVO) GetWirelessVlanAssignmentOk() (*bool, bool)`

GetWirelessVlanAssignmentOk returns a tuple with the WirelessVlanAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessVlanAssignment

`func (o *RadiusProfileOpenApiVO) SetWirelessVlanAssignment(v bool)`

SetWirelessVlanAssignment sets WirelessVlanAssignment field to given value.

### HasWirelessVlanAssignment

`func (o *RadiusProfileOpenApiVO) HasWirelessVlanAssignment() bool`

HasWirelessVlanAssignment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


