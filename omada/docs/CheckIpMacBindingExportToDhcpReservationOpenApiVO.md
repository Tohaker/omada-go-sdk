# CheckIpMacBindingExportToDhcpReservationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | **string** | Interface ID. WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | 
**Ip** | **string** | IP of the IP MAC binding entity. | 
**Mac** | **string** | MAC of the IP MAC binding entity. | 

## Methods

### NewCheckIpMacBindingExportToDhcpReservationOpenApiVO

`func NewCheckIpMacBindingExportToDhcpReservationOpenApiVO(interfaceId string, ip string, mac string, ) *CheckIpMacBindingExportToDhcpReservationOpenApiVO`

NewCheckIpMacBindingExportToDhcpReservationOpenApiVO instantiates a new CheckIpMacBindingExportToDhcpReservationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckIpMacBindingExportToDhcpReservationOpenApiVOWithDefaults

`func NewCheckIpMacBindingExportToDhcpReservationOpenApiVOWithDefaults() *CheckIpMacBindingExportToDhcpReservationOpenApiVO`

NewCheckIpMacBindingExportToDhcpReservationOpenApiVOWithDefaults instantiates a new CheckIpMacBindingExportToDhcpReservationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.


### GetIp

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMac

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CheckIpMacBindingExportToDhcpReservationOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


