# DhcpReservationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abnormal** | Pointer to **int32** | Abnormal cause | [optional] 
**ClientName** | Pointer to **string** | Client name | [optional] 
**Description** | Pointer to **string** | Description of DHCP reservation | [optional] 
**ExistOptions** | Pointer to **[]int32** | Options configured in the current DHCP Reservation entry. | [optional] 
**ExportToIpMacBinding** | Pointer to **bool** | Indicates whether DHCP reservation can be exported to the IP-MAC Binding list | [optional] 
**Id** | Pointer to **string** | DHCP reservation ID | [optional] 
**Ip** | Pointer to **string** | Reserved IP address | [optional] 
**Mac** | Pointer to **string** | Device MAC address | [optional] 
**Model** | Pointer to **string** | Device Model | [optional] 
**ModelVersion** | Pointer to **string** | Device Model Version | [optional] 
**Name** | Pointer to **string** | Device Name | [optional] 
**NetId** | Pointer to **string** | ID of the configured LAN Network | [optional] 
**NetName** | Pointer to **string** | Name of the configured LAN Network | [optional] 
**Options** | Pointer to [**[]CustomDHCPOptions**](CustomDHCPOptions.md) | Advanced DHCP options | [optional] 
**ServerMac** | Pointer to **string** | Dhcp Server Device Mac | [optional] 
**ServerName** | Pointer to **string** | Dhcp Server Device Name | [optional] 
**ServerStackId** | Pointer to **string** | Dhcp Server Stack ID | [optional] 
**ServerType** | Pointer to **string** | Dhcp Server Device Type | [optional] 
**ShowingType** | Pointer to **string** | Client Type or Device Type | [optional] 
**Status** | Pointer to **bool** | DHCP reservation enable status | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 0: Device; 1: Client | [optional] 

## Methods

### NewDhcpReservationOpenApiVO

`func NewDhcpReservationOpenApiVO() *DhcpReservationOpenApiVO`

NewDhcpReservationOpenApiVO instantiates a new DhcpReservationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpReservationOpenApiVOWithDefaults

`func NewDhcpReservationOpenApiVOWithDefaults() *DhcpReservationOpenApiVO`

NewDhcpReservationOpenApiVOWithDefaults instantiates a new DhcpReservationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbnormal

`func (o *DhcpReservationOpenApiVO) GetAbnormal() int32`

GetAbnormal returns the Abnormal field if non-nil, zero value otherwise.

### GetAbnormalOk

`func (o *DhcpReservationOpenApiVO) GetAbnormalOk() (*int32, bool)`

GetAbnormalOk returns a tuple with the Abnormal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormal

`func (o *DhcpReservationOpenApiVO) SetAbnormal(v int32)`

SetAbnormal sets Abnormal field to given value.

### HasAbnormal

`func (o *DhcpReservationOpenApiVO) HasAbnormal() bool`

HasAbnormal returns a boolean if a field has been set.

### GetClientName

`func (o *DhcpReservationOpenApiVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *DhcpReservationOpenApiVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *DhcpReservationOpenApiVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *DhcpReservationOpenApiVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDescription

`func (o *DhcpReservationOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DhcpReservationOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DhcpReservationOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DhcpReservationOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExistOptions

`func (o *DhcpReservationOpenApiVO) GetExistOptions() []int32`

GetExistOptions returns the ExistOptions field if non-nil, zero value otherwise.

### GetExistOptionsOk

`func (o *DhcpReservationOpenApiVO) GetExistOptionsOk() (*[]int32, bool)`

GetExistOptionsOk returns a tuple with the ExistOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistOptions

`func (o *DhcpReservationOpenApiVO) SetExistOptions(v []int32)`

SetExistOptions sets ExistOptions field to given value.

### HasExistOptions

`func (o *DhcpReservationOpenApiVO) HasExistOptions() bool`

HasExistOptions returns a boolean if a field has been set.

### GetExportToIpMacBinding

`func (o *DhcpReservationOpenApiVO) GetExportToIpMacBinding() bool`

GetExportToIpMacBinding returns the ExportToIpMacBinding field if non-nil, zero value otherwise.

### GetExportToIpMacBindingOk

`func (o *DhcpReservationOpenApiVO) GetExportToIpMacBindingOk() (*bool, bool)`

GetExportToIpMacBindingOk returns a tuple with the ExportToIpMacBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportToIpMacBinding

`func (o *DhcpReservationOpenApiVO) SetExportToIpMacBinding(v bool)`

SetExportToIpMacBinding sets ExportToIpMacBinding field to given value.

### HasExportToIpMacBinding

`func (o *DhcpReservationOpenApiVO) HasExportToIpMacBinding() bool`

HasExportToIpMacBinding returns a boolean if a field has been set.

### GetId

`func (o *DhcpReservationOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DhcpReservationOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DhcpReservationOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DhcpReservationOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *DhcpReservationOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DhcpReservationOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DhcpReservationOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DhcpReservationOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *DhcpReservationOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DhcpReservationOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DhcpReservationOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DhcpReservationOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *DhcpReservationOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DhcpReservationOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DhcpReservationOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DhcpReservationOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DhcpReservationOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DhcpReservationOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DhcpReservationOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DhcpReservationOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DhcpReservationOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DhcpReservationOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DhcpReservationOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DhcpReservationOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetId

`func (o *DhcpReservationOpenApiVO) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *DhcpReservationOpenApiVO) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *DhcpReservationOpenApiVO) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *DhcpReservationOpenApiVO) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetNetName

`func (o *DhcpReservationOpenApiVO) GetNetName() string`

GetNetName returns the NetName field if non-nil, zero value otherwise.

### GetNetNameOk

`func (o *DhcpReservationOpenApiVO) GetNetNameOk() (*string, bool)`

GetNetNameOk returns a tuple with the NetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetName

`func (o *DhcpReservationOpenApiVO) SetNetName(v string)`

SetNetName sets NetName field to given value.

### HasNetName

`func (o *DhcpReservationOpenApiVO) HasNetName() bool`

HasNetName returns a boolean if a field has been set.

### GetOptions

`func (o *DhcpReservationOpenApiVO) GetOptions() []CustomDHCPOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DhcpReservationOpenApiVO) GetOptionsOk() (*[]CustomDHCPOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DhcpReservationOpenApiVO) SetOptions(v []CustomDHCPOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DhcpReservationOpenApiVO) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetServerMac

`func (o *DhcpReservationOpenApiVO) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *DhcpReservationOpenApiVO) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *DhcpReservationOpenApiVO) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *DhcpReservationOpenApiVO) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpReservationOpenApiVO) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpReservationOpenApiVO) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpReservationOpenApiVO) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpReservationOpenApiVO) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerStackId

`func (o *DhcpReservationOpenApiVO) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *DhcpReservationOpenApiVO) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *DhcpReservationOpenApiVO) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *DhcpReservationOpenApiVO) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpReservationOpenApiVO) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpReservationOpenApiVO) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpReservationOpenApiVO) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpReservationOpenApiVO) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetShowingType

`func (o *DhcpReservationOpenApiVO) GetShowingType() string`

GetShowingType returns the ShowingType field if non-nil, zero value otherwise.

### GetShowingTypeOk

`func (o *DhcpReservationOpenApiVO) GetShowingTypeOk() (*string, bool)`

GetShowingTypeOk returns a tuple with the ShowingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowingType

`func (o *DhcpReservationOpenApiVO) SetShowingType(v string)`

SetShowingType sets ShowingType field to given value.

### HasShowingType

`func (o *DhcpReservationOpenApiVO) HasShowingType() bool`

HasShowingType returns a boolean if a field has been set.

### GetStatus

`func (o *DhcpReservationOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DhcpReservationOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DhcpReservationOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DhcpReservationOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *DhcpReservationOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DhcpReservationOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DhcpReservationOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DhcpReservationOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


