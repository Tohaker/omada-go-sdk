# TopModelBaseVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCount** | Pointer to **int64** | Count of clients | [optional] 
**Device** | Pointer to **string** | Device name | [optional] 
**Health** | Pointer to **int32** | Device health core | [optional] 
**Ip** | Pointer to **string** | Device IP | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelType** | Pointer to **string** |  | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**PoePower** | Pointer to **int64** | POE power | [optional] 
**Ratio** | Pointer to **int64** | Percent | [optional] 
**Status** | Pointer to **int32** | Device status should be a value as follows: 0:DISCONNECTED, 1:CONNECTED, 2:PENDING, 3:HEARTBEAT MISSED, 4: ISOLATED | [optional] 
**Traffic** | Pointer to **int64** | Total traffic | [optional] 
**Type** | Pointer to **string** | Device type should be a value as follows: ap; switch | [optional] 
**Util** | Pointer to **int64** | Utilization | [optional] 

## Methods

### NewTopModelBaseVO

`func NewTopModelBaseVO() *TopModelBaseVO`

NewTopModelBaseVO instantiates a new TopModelBaseVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopModelBaseVOWithDefaults

`func NewTopModelBaseVOWithDefaults() *TopModelBaseVO`

NewTopModelBaseVOWithDefaults instantiates a new TopModelBaseVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCount

`func (o *TopModelBaseVO) GetClientCount() int64`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *TopModelBaseVO) GetClientCountOk() (*int64, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *TopModelBaseVO) SetClientCount(v int64)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *TopModelBaseVO) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.

### GetDevice

`func (o *TopModelBaseVO) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *TopModelBaseVO) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *TopModelBaseVO) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *TopModelBaseVO) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetHealth

`func (o *TopModelBaseVO) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *TopModelBaseVO) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *TopModelBaseVO) SetHealth(v int32)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *TopModelBaseVO) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetIp

`func (o *TopModelBaseVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TopModelBaseVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TopModelBaseVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TopModelBaseVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *TopModelBaseVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopModelBaseVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopModelBaseVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopModelBaseVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *TopModelBaseVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopModelBaseVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopModelBaseVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopModelBaseVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelType

`func (o *TopModelBaseVO) GetModelType() string`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *TopModelBaseVO) GetModelTypeOk() (*string, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *TopModelBaseVO) SetModelType(v string)`

SetModelType sets ModelType field to given value.

### HasModelType

`func (o *TopModelBaseVO) HasModelType() bool`

HasModelType returns a boolean if a field has been set.

### GetModelVersion

`func (o *TopModelBaseVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *TopModelBaseVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *TopModelBaseVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *TopModelBaseVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetPoePower

`func (o *TopModelBaseVO) GetPoePower() int64`

GetPoePower returns the PoePower field if non-nil, zero value otherwise.

### GetPoePowerOk

`func (o *TopModelBaseVO) GetPoePowerOk() (*int64, bool)`

GetPoePowerOk returns a tuple with the PoePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePower

`func (o *TopModelBaseVO) SetPoePower(v int64)`

SetPoePower sets PoePower field to given value.

### HasPoePower

`func (o *TopModelBaseVO) HasPoePower() bool`

HasPoePower returns a boolean if a field has been set.

### GetRatio

`func (o *TopModelBaseVO) GetRatio() int64`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *TopModelBaseVO) GetRatioOk() (*int64, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *TopModelBaseVO) SetRatio(v int64)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *TopModelBaseVO) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetStatus

`func (o *TopModelBaseVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TopModelBaseVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TopModelBaseVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TopModelBaseVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTraffic

`func (o *TopModelBaseVO) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *TopModelBaseVO) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *TopModelBaseVO) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *TopModelBaseVO) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.

### GetType

`func (o *TopModelBaseVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopModelBaseVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopModelBaseVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopModelBaseVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUtil

`func (o *TopModelBaseVO) GetUtil() int64`

GetUtil returns the Util field if non-nil, zero value otherwise.

### GetUtilOk

`func (o *TopModelBaseVO) GetUtilOk() (*int64, bool)`

GetUtilOk returns a tuple with the Util field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtil

`func (o *TopModelBaseVO) SetUtil(v int64)`

SetUtil sets Util field to given value.

### HasUtil

`func (o *TopModelBaseVO) HasUtil() bool`

HasUtil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


