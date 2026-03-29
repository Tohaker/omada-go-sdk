# DhcpServerForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ranges** | Pointer to [**[]DhcpServerRangeVO**](DhcpServerRangeVO.md) | Dhcp Server Ranges | [optional] 

## Methods

### NewDhcpServerForVlanVO

`func NewDhcpServerForVlanVO() *DhcpServerForVlanVO`

NewDhcpServerForVlanVO instantiates a new DhcpServerForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpServerForVlanVOWithDefaults

`func NewDhcpServerForVlanVOWithDefaults() *DhcpServerForVlanVO`

NewDhcpServerForVlanVOWithDefaults instantiates a new DhcpServerForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRanges

`func (o *DhcpServerForVlanVO) GetRanges() []DhcpServerRangeVO`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *DhcpServerForVlanVO) GetRangesOk() (*[]DhcpServerRangeVO, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *DhcpServerForVlanVO) SetRanges(v []DhcpServerRangeVO)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *DhcpServerForVlanVO) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


