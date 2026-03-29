# IPMacBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description should contain 1 to 64 characters. | [optional] 
**EnableExportToDhcpReservation** | Pointer to **bool** | Whether to enable export to the DhcpReservation table. | [optional] 
**ExportToDhcpReservation** | Pointer to **bool** | Whether to export to the DhcpReservation table. This subsection is deprecated. | [optional] 
**InterfaceId** | **string** | Interface ID. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**InterfaceType** | **int32** | Interface type should be a value as follows: 0: WAN; 1: LAN interface. | 
**Ip** | **string** | IP of the IP MAC binding entity. | 
**Mac** | **string** | MAC of the IP MAC binding entity. | 
**Status** | **bool** | Status of the IP MAC binding entity. | 
**SupportExport** | Pointer to **bool** | Whether could export to the DhcpReservation table. This subsection is deprecated. | [optional] 

## Methods

### NewIPMacBinding

`func NewIPMacBinding(interfaceId string, interfaceType int32, ip string, mac string, status bool, ) *IPMacBinding`

NewIPMacBinding instantiates a new IPMacBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPMacBindingWithDefaults

`func NewIPMacBindingWithDefaults() *IPMacBinding`

NewIPMacBindingWithDefaults instantiates a new IPMacBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IPMacBinding) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPMacBinding) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPMacBinding) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPMacBinding) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableExportToDhcpReservation

`func (o *IPMacBinding) GetEnableExportToDhcpReservation() bool`

GetEnableExportToDhcpReservation returns the EnableExportToDhcpReservation field if non-nil, zero value otherwise.

### GetEnableExportToDhcpReservationOk

`func (o *IPMacBinding) GetEnableExportToDhcpReservationOk() (*bool, bool)`

GetEnableExportToDhcpReservationOk returns a tuple with the EnableExportToDhcpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExportToDhcpReservation

`func (o *IPMacBinding) SetEnableExportToDhcpReservation(v bool)`

SetEnableExportToDhcpReservation sets EnableExportToDhcpReservation field to given value.

### HasEnableExportToDhcpReservation

`func (o *IPMacBinding) HasEnableExportToDhcpReservation() bool`

HasEnableExportToDhcpReservation returns a boolean if a field has been set.

### GetExportToDhcpReservation

`func (o *IPMacBinding) GetExportToDhcpReservation() bool`

GetExportToDhcpReservation returns the ExportToDhcpReservation field if non-nil, zero value otherwise.

### GetExportToDhcpReservationOk

`func (o *IPMacBinding) GetExportToDhcpReservationOk() (*bool, bool)`

GetExportToDhcpReservationOk returns a tuple with the ExportToDhcpReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportToDhcpReservation

`func (o *IPMacBinding) SetExportToDhcpReservation(v bool)`

SetExportToDhcpReservation sets ExportToDhcpReservation field to given value.

### HasExportToDhcpReservation

`func (o *IPMacBinding) HasExportToDhcpReservation() bool`

HasExportToDhcpReservation returns a boolean if a field has been set.

### GetInterfaceId

`func (o *IPMacBinding) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *IPMacBinding) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *IPMacBinding) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.


### GetInterfaceType

`func (o *IPMacBinding) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *IPMacBinding) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *IPMacBinding) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.


### GetIp

`func (o *IPMacBinding) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IPMacBinding) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IPMacBinding) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMac

`func (o *IPMacBinding) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *IPMacBinding) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *IPMacBinding) SetMac(v string)`

SetMac sets Mac field to given value.


### GetStatus

`func (o *IPMacBinding) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPMacBinding) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPMacBinding) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetSupportExport

`func (o *IPMacBinding) GetSupportExport() bool`

GetSupportExport returns the SupportExport field if non-nil, zero value otherwise.

### GetSupportExportOk

`func (o *IPMacBinding) GetSupportExportOk() (*bool, bool)`

GetSupportExportOk returns a tuple with the SupportExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportExport

`func (o *IPMacBinding) SetSupportExport(v bool)`

SetSupportExport sets SupportExport field to given value.

### HasSupportExport

`func (o *IPMacBinding) HasSupportExport() bool`

HasSupportExport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


