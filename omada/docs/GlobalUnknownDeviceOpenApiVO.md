# GlobalUnknownDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSeriesType** | Pointer to **int32** | Device series type. 0 means basic, 1 means pro. | [optional] 
**Ip** | Pointer to **string** | Device IP | [optional] 
**LastSeen** | Pointer to **int64** | Device lastSeen, unit: ms | [optional] 
**Mac** | Pointer to **string** | Device MAC | [optional] 
**Model** | Pointer to **string** | Device model name | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Status** | Pointer to **int32** | Device status should be a value as follows: 0: Disconnected; 1: Connected; 2: Pending; 3: Heartbeat Missed; 4: Isolated | [optional] 
**Subtype** | Pointer to **string** | Switch subtype should be a value as follows: smart: Non-Agile Series Switch; es: Agile Series Switch. | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 
**Uptime** | Pointer to **string** | Device uptime | [optional] 

## Methods

### NewGlobalUnknownDeviceOpenApiVO

`func NewGlobalUnknownDeviceOpenApiVO() *GlobalUnknownDeviceOpenApiVO`

NewGlobalUnknownDeviceOpenApiVO instantiates a new GlobalUnknownDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalUnknownDeviceOpenApiVOWithDefaults

`func NewGlobalUnknownDeviceOpenApiVOWithDefaults() *GlobalUnknownDeviceOpenApiVO`

NewGlobalUnknownDeviceOpenApiVOWithDefaults instantiates a new GlobalUnknownDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSeriesType

`func (o *GlobalUnknownDeviceOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *GlobalUnknownDeviceOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *GlobalUnknownDeviceOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetIp

`func (o *GlobalUnknownDeviceOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GlobalUnknownDeviceOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GlobalUnknownDeviceOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLastSeen

`func (o *GlobalUnknownDeviceOpenApiVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GlobalUnknownDeviceOpenApiVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GlobalUnknownDeviceOpenApiVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *GlobalUnknownDeviceOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GlobalUnknownDeviceOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GlobalUnknownDeviceOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *GlobalUnknownDeviceOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GlobalUnknownDeviceOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GlobalUnknownDeviceOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *GlobalUnknownDeviceOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalUnknownDeviceOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalUnknownDeviceOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *GlobalUnknownDeviceOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GlobalUnknownDeviceOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GlobalUnknownDeviceOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtype

`func (o *GlobalUnknownDeviceOpenApiVO) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *GlobalUnknownDeviceOpenApiVO) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *GlobalUnknownDeviceOpenApiVO) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetType

`func (o *GlobalUnknownDeviceOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalUnknownDeviceOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GlobalUnknownDeviceOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *GlobalUnknownDeviceOpenApiVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GlobalUnknownDeviceOpenApiVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GlobalUnknownDeviceOpenApiVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GlobalUnknownDeviceOpenApiVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


