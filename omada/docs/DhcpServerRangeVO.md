# DhcpServerRangeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableIpNum** | Pointer to **int32** | The number of available IP addresses in the server address pool. | [optional] 
**EndIp** | Pointer to **string** | End IP | [optional] 
**StartIp** | Pointer to **string** | Start IP | [optional] 

## Methods

### NewDhcpServerRangeVO

`func NewDhcpServerRangeVO() *DhcpServerRangeVO`

NewDhcpServerRangeVO instantiates a new DhcpServerRangeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpServerRangeVOWithDefaults

`func NewDhcpServerRangeVOWithDefaults() *DhcpServerRangeVO`

NewDhcpServerRangeVOWithDefaults instantiates a new DhcpServerRangeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableIpNum

`func (o *DhcpServerRangeVO) GetAvailableIpNum() int32`

GetAvailableIpNum returns the AvailableIpNum field if non-nil, zero value otherwise.

### GetAvailableIpNumOk

`func (o *DhcpServerRangeVO) GetAvailableIpNumOk() (*int32, bool)`

GetAvailableIpNumOk returns a tuple with the AvailableIpNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableIpNum

`func (o *DhcpServerRangeVO) SetAvailableIpNum(v int32)`

SetAvailableIpNum sets AvailableIpNum field to given value.

### HasAvailableIpNum

`func (o *DhcpServerRangeVO) HasAvailableIpNum() bool`

HasAvailableIpNum returns a boolean if a field has been set.

### GetEndIp

`func (o *DhcpServerRangeVO) GetEndIp() string`

GetEndIp returns the EndIp field if non-nil, zero value otherwise.

### GetEndIpOk

`func (o *DhcpServerRangeVO) GetEndIpOk() (*string, bool)`

GetEndIpOk returns a tuple with the EndIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIp

`func (o *DhcpServerRangeVO) SetEndIp(v string)`

SetEndIp sets EndIp field to given value.

### HasEndIp

`func (o *DhcpServerRangeVO) HasEndIp() bool`

HasEndIp returns a boolean if a field has been set.

### GetStartIp

`func (o *DhcpServerRangeVO) GetStartIp() string`

GetStartIp returns the StartIp field if non-nil, zero value otherwise.

### GetStartIpOk

`func (o *DhcpServerRangeVO) GetStartIpOk() (*string, bool)`

GetStartIpOk returns a tuple with the StartIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIp

`func (o *DhcpServerRangeVO) SetStartIp(v string)`

SetStartIp sets StartIp field to given value.

### HasStartIp

`func (o *DhcpServerRangeVO) HasStartIp() bool`

HasStartIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


