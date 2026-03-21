# ExportDhcpReservationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | **string** | This field represents LAN Network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface | 

## Methods

### NewExportDhcpReservationOpenApiVO

`func NewExportDhcpReservationOpenApiVO(interfaceId string, ) *ExportDhcpReservationOpenApiVO`

NewExportDhcpReservationOpenApiVO instantiates a new ExportDhcpReservationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportDhcpReservationOpenApiVOWithDefaults

`func NewExportDhcpReservationOpenApiVOWithDefaults() *ExportDhcpReservationOpenApiVO`

NewExportDhcpReservationOpenApiVOWithDefaults instantiates a new ExportDhcpReservationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *ExportDhcpReservationOpenApiVO) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *ExportDhcpReservationOpenApiVO) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *ExportDhcpReservationOpenApiVO) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


