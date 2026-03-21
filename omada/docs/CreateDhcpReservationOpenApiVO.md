# CreateDhcpReservationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmConflict** | Pointer to **bool** | True when creating an entry with IP and IP-MAC Binding conflicts confirmed | [optional] 
**Description** | Pointer to **string** | Description of DHCP reservation. Description should contain 1 to 128 characters | [optional] 
**Ip** | Pointer to **string** | Reserved IP address | [optional] 
**Mac** | **string** | Device MAC address, format: AA-BB-CC-11-22-33 | 
**NetId** | **string** | This field represents LAN Network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface | 
**Options** | Pointer to [**[]CustomDHCPOptions**](CustomDHCPOptions.md) | Advanced DHCP options | [optional] 
**ServerMac** | Pointer to **string** | Dhcp Server Device Mac | [optional] 
**ServerStackId** | Pointer to **string** | Dhcp Server Stack ID | [optional] 
**ServerType** | Pointer to **string** | Dhcp Server Device Type | [optional] 
**Status** | **bool** | DHCP reservation enable status | 

## Methods

### NewCreateDhcpReservationOpenApiVO

`func NewCreateDhcpReservationOpenApiVO(mac string, netId string, status bool, ) *CreateDhcpReservationOpenApiVO`

NewCreateDhcpReservationOpenApiVO instantiates a new CreateDhcpReservationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDhcpReservationOpenApiVOWithDefaults

`func NewCreateDhcpReservationOpenApiVOWithDefaults() *CreateDhcpReservationOpenApiVO`

NewCreateDhcpReservationOpenApiVOWithDefaults instantiates a new CreateDhcpReservationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmConflict

`func (o *CreateDhcpReservationOpenApiVO) GetConfirmConflict() bool`

GetConfirmConflict returns the ConfirmConflict field if non-nil, zero value otherwise.

### GetConfirmConflictOk

`func (o *CreateDhcpReservationOpenApiVO) GetConfirmConflictOk() (*bool, bool)`

GetConfirmConflictOk returns a tuple with the ConfirmConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmConflict

`func (o *CreateDhcpReservationOpenApiVO) SetConfirmConflict(v bool)`

SetConfirmConflict sets ConfirmConflict field to given value.

### HasConfirmConflict

`func (o *CreateDhcpReservationOpenApiVO) HasConfirmConflict() bool`

HasConfirmConflict returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDhcpReservationOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDhcpReservationOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDhcpReservationOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDhcpReservationOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIp

`func (o *CreateDhcpReservationOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CreateDhcpReservationOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CreateDhcpReservationOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CreateDhcpReservationOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *CreateDhcpReservationOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CreateDhcpReservationOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CreateDhcpReservationOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetNetId

`func (o *CreateDhcpReservationOpenApiVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *CreateDhcpReservationOpenApiVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *CreateDhcpReservationOpenApiVO) SetNetId(v string)`

SetNetId sets NetId field to given value.


### GetOptions

`func (o *CreateDhcpReservationOpenApiVO) GetOptions() []CustomDHCPOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateDhcpReservationOpenApiVO) GetOptionsOk() (*[]CustomDHCPOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateDhcpReservationOpenApiVO) SetOptions(v []CustomDHCPOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateDhcpReservationOpenApiVO) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetServerMac

`func (o *CreateDhcpReservationOpenApiVO) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *CreateDhcpReservationOpenApiVO) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *CreateDhcpReservationOpenApiVO) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *CreateDhcpReservationOpenApiVO) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerStackId

`func (o *CreateDhcpReservationOpenApiVO) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *CreateDhcpReservationOpenApiVO) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *CreateDhcpReservationOpenApiVO) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *CreateDhcpReservationOpenApiVO) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetServerType

`func (o *CreateDhcpReservationOpenApiVO) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *CreateDhcpReservationOpenApiVO) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *CreateDhcpReservationOpenApiVO) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *CreateDhcpReservationOpenApiVO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetStatus

`func (o *CreateDhcpReservationOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateDhcpReservationOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateDhcpReservationOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


