# ApFullChannelScanStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelLoadStatus** | Pointer to **int32** | Channel load detection status of the AP device. Status should be a value as follows: 0: In the unscanned state but the scan result is displayed;1: In the unscanned state, and there is no scan result;3: Scanning. | [optional] 
**FromGlobal** | Pointer to **bool** | If the status from batch full channel detect | [optional] 
**InterfStatus** | Pointer to **int32** | Interference detection status of the AP device. Status should be a value as follows: 0: In the unscanned state but the scan result is displayed;1: In the unscanned state, and there is no scan result;3: Scanning. | [optional] 
**LastSeen** | Pointer to **int64** | The last time full channel detection was triggered. | [optional] 
**Status** | Pointer to **int32** | Full channel detection status of the AP device. Status should be a value as follows: 0: In the unscanned state but the scan result is displayed;1: In the unscanned state, and there is no scan result;3: Scanning. | [optional] 
**WifiInterfStatus** | Pointer to **int32** | Wifi interference detection status of the AP device. Status should be a value as follows: 0: In the unscanned state but the scan result is displayed;1: In the unscanned state, and there is no scan result;3: Scanning. | [optional] 

## Methods

### NewApFullChannelScanStatus

`func NewApFullChannelScanStatus() *ApFullChannelScanStatus`

NewApFullChannelScanStatus instantiates a new ApFullChannelScanStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApFullChannelScanStatusWithDefaults

`func NewApFullChannelScanStatusWithDefaults() *ApFullChannelScanStatus`

NewApFullChannelScanStatusWithDefaults instantiates a new ApFullChannelScanStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelLoadStatus

`func (o *ApFullChannelScanStatus) GetChannelLoadStatus() int32`

GetChannelLoadStatus returns the ChannelLoadStatus field if non-nil, zero value otherwise.

### GetChannelLoadStatusOk

`func (o *ApFullChannelScanStatus) GetChannelLoadStatusOk() (*int32, bool)`

GetChannelLoadStatusOk returns a tuple with the ChannelLoadStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLoadStatus

`func (o *ApFullChannelScanStatus) SetChannelLoadStatus(v int32)`

SetChannelLoadStatus sets ChannelLoadStatus field to given value.

### HasChannelLoadStatus

`func (o *ApFullChannelScanStatus) HasChannelLoadStatus() bool`

HasChannelLoadStatus returns a boolean if a field has been set.

### GetFromGlobal

`func (o *ApFullChannelScanStatus) GetFromGlobal() bool`

GetFromGlobal returns the FromGlobal field if non-nil, zero value otherwise.

### GetFromGlobalOk

`func (o *ApFullChannelScanStatus) GetFromGlobalOk() (*bool, bool)`

GetFromGlobalOk returns a tuple with the FromGlobal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromGlobal

`func (o *ApFullChannelScanStatus) SetFromGlobal(v bool)`

SetFromGlobal sets FromGlobal field to given value.

### HasFromGlobal

`func (o *ApFullChannelScanStatus) HasFromGlobal() bool`

HasFromGlobal returns a boolean if a field has been set.

### GetInterfStatus

`func (o *ApFullChannelScanStatus) GetInterfStatus() int32`

GetInterfStatus returns the InterfStatus field if non-nil, zero value otherwise.

### GetInterfStatusOk

`func (o *ApFullChannelScanStatus) GetInterfStatusOk() (*int32, bool)`

GetInterfStatusOk returns a tuple with the InterfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfStatus

`func (o *ApFullChannelScanStatus) SetInterfStatus(v int32)`

SetInterfStatus sets InterfStatus field to given value.

### HasInterfStatus

`func (o *ApFullChannelScanStatus) HasInterfStatus() bool`

HasInterfStatus returns a boolean if a field has been set.

### GetLastSeen

`func (o *ApFullChannelScanStatus) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ApFullChannelScanStatus) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ApFullChannelScanStatus) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ApFullChannelScanStatus) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetStatus

`func (o *ApFullChannelScanStatus) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApFullChannelScanStatus) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApFullChannelScanStatus) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApFullChannelScanStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWifiInterfStatus

`func (o *ApFullChannelScanStatus) GetWifiInterfStatus() int32`

GetWifiInterfStatus returns the WifiInterfStatus field if non-nil, zero value otherwise.

### GetWifiInterfStatusOk

`func (o *ApFullChannelScanStatus) GetWifiInterfStatusOk() (*int32, bool)`

GetWifiInterfStatusOk returns a tuple with the WifiInterfStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiInterfStatus

`func (o *ApFullChannelScanStatus) SetWifiInterfStatus(v int32)`

SetWifiInterfStatus sets WifiInterfStatus field to given value.

### HasWifiInterfStatus

`func (o *ApFullChannelScanStatus) HasWifiInterfStatus() bool`

HasWifiInterfStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


