# DhcpServerDevCapForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportGetDhcpClientTable** | Pointer to **bool** | Indicates whether the device supports obtaining DHCP client table information through the GET message | [optional] 
**SupportGetDhcpServerInfo** | Pointer to **bool** | Indicates whether the device supports obtaining DHCP server information through the GET message | [optional] 

## Methods

### NewDhcpServerDevCapForVlanVO

`func NewDhcpServerDevCapForVlanVO() *DhcpServerDevCapForVlanVO`

NewDhcpServerDevCapForVlanVO instantiates a new DhcpServerDevCapForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpServerDevCapForVlanVOWithDefaults

`func NewDhcpServerDevCapForVlanVOWithDefaults() *DhcpServerDevCapForVlanVO`

NewDhcpServerDevCapForVlanVOWithDefaults instantiates a new DhcpServerDevCapForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportGetDhcpClientTable

`func (o *DhcpServerDevCapForVlanVO) GetSupportGetDhcpClientTable() bool`

GetSupportGetDhcpClientTable returns the SupportGetDhcpClientTable field if non-nil, zero value otherwise.

### GetSupportGetDhcpClientTableOk

`func (o *DhcpServerDevCapForVlanVO) GetSupportGetDhcpClientTableOk() (*bool, bool)`

GetSupportGetDhcpClientTableOk returns a tuple with the SupportGetDhcpClientTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGetDhcpClientTable

`func (o *DhcpServerDevCapForVlanVO) SetSupportGetDhcpClientTable(v bool)`

SetSupportGetDhcpClientTable sets SupportGetDhcpClientTable field to given value.

### HasSupportGetDhcpClientTable

`func (o *DhcpServerDevCapForVlanVO) HasSupportGetDhcpClientTable() bool`

HasSupportGetDhcpClientTable returns a boolean if a field has been set.

### GetSupportGetDhcpServerInfo

`func (o *DhcpServerDevCapForVlanVO) GetSupportGetDhcpServerInfo() bool`

GetSupportGetDhcpServerInfo returns the SupportGetDhcpServerInfo field if non-nil, zero value otherwise.

### GetSupportGetDhcpServerInfoOk

`func (o *DhcpServerDevCapForVlanVO) GetSupportGetDhcpServerInfoOk() (*bool, bool)`

GetSupportGetDhcpServerInfoOk returns a tuple with the SupportGetDhcpServerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGetDhcpServerInfo

`func (o *DhcpServerDevCapForVlanVO) SetSupportGetDhcpServerInfo(v bool)`

SetSupportGetDhcpServerInfo sets SupportGetDhcpServerInfo field to given value.

### HasSupportGetDhcpServerInfo

`func (o *DhcpServerDevCapForVlanVO) HasSupportGetDhcpServerInfo() bool`

HasSupportGetDhcpServerInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


