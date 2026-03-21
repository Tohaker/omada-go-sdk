# UpdateSsidDhcpOptionOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitId** | Pointer to **[]int32** | SSID DHCP Option 82 Circuit-ID config. Circuit-ID is an array formed in the selected order, with each array element corresponding to the following enumeration values: 1: VLAN-ID; 2: AP Radio Mac-Address; 3: SSID-Type; 4: SSID-Name; 5: AP Ethernet MAC address; 6: Site-Name. As in the example [3,1,2,4], the following enumeration values are selected sequentially on the page: SSID-Type, VLAN-ID, AP Radio Mac-Address and SSID-Name. | [optional] 
**Delimiter** | Pointer to **string** | SSID DHCP Option 82 delimiter config (A single arbitrary ASCII character is acceptable). | [optional] 
**DhcpEnable** | **bool** | SSID DHCP Option 82 global config status. True: enable, false: disable. | 
**Format** | Pointer to **int32** | SSID DHCP Option 82 format config; It should be a value as follows: 0: ASCII; 1: Binary. | [optional] 
**RemoteId** | Pointer to **[]int32** | SSID DHCP Option 82 Remote-ID config. Remote-ID is an array formed in the selected order, with each array element corresponding to the following enumeration values: 1: VLAN-ID; 2: AP Radio Mac-Address; 3: SSID-Type; 4: SSID-Name; 5: AP Ethernet MAC address; 6: Site-Name. As in the example [3,1,2,4], the following enumeration values are selected sequentially on the page: SSID-Type, VLAN-ID, AP Radio Mac-Address and SSID-Name. | [optional] 

## Methods

### NewUpdateSsidDhcpOptionOpenApiVO

`func NewUpdateSsidDhcpOptionOpenApiVO(dhcpEnable bool, ) *UpdateSsidDhcpOptionOpenApiVO`

NewUpdateSsidDhcpOptionOpenApiVO instantiates a new UpdateSsidDhcpOptionOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSsidDhcpOptionOpenApiVOWithDefaults

`func NewUpdateSsidDhcpOptionOpenApiVOWithDefaults() *UpdateSsidDhcpOptionOpenApiVO`

NewUpdateSsidDhcpOptionOpenApiVOWithDefaults instantiates a new UpdateSsidDhcpOptionOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitId

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetCircuitId() []int32`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetCircuitIdOk() (*[]int32, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *UpdateSsidDhcpOptionOpenApiVO) SetCircuitId(v []int32)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *UpdateSsidDhcpOptionOpenApiVO) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetDelimiter

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetDelimiter() string`

GetDelimiter returns the Delimiter field if non-nil, zero value otherwise.

### GetDelimiterOk

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetDelimiterOk() (*string, bool)`

GetDelimiterOk returns a tuple with the Delimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelimiter

`func (o *UpdateSsidDhcpOptionOpenApiVO) SetDelimiter(v string)`

SetDelimiter sets Delimiter field to given value.

### HasDelimiter

`func (o *UpdateSsidDhcpOptionOpenApiVO) HasDelimiter() bool`

HasDelimiter returns a boolean if a field has been set.

### GetDhcpEnable

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetDhcpEnable() bool`

GetDhcpEnable returns the DhcpEnable field if non-nil, zero value otherwise.

### GetDhcpEnableOk

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetDhcpEnableOk() (*bool, bool)`

GetDhcpEnableOk returns a tuple with the DhcpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpEnable

`func (o *UpdateSsidDhcpOptionOpenApiVO) SetDhcpEnable(v bool)`

SetDhcpEnable sets DhcpEnable field to given value.


### GetFormat

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *UpdateSsidDhcpOptionOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *UpdateSsidDhcpOptionOpenApiVO) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetRemoteId

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetRemoteId() []int32`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *UpdateSsidDhcpOptionOpenApiVO) GetRemoteIdOk() (*[]int32, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *UpdateSsidDhcpOptionOpenApiVO) SetRemoteId(v []int32)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *UpdateSsidDhcpOptionOpenApiVO) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


