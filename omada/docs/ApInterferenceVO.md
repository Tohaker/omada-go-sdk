# ApInterferenceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device** | Pointer to **string** | Device name | [optional] 
**InterUtil** | Pointer to **int32** | Channel interference rate | [optional] 
**Interference** | Pointer to **int32** | Degree of interference. interference &#x3D; interUtil/(interUtil+txUtil+rxUtil) | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**RxUtil** | Pointer to **int32** | Receive channel utilization | [optional] 
**TxUtil** | Pointer to **int32** | Transmit channel utilization | [optional] 
**Type** | Pointer to **string** | Device type should be a value as follows: ap. | [optional] 

## Methods

### NewApInterferenceVO

`func NewApInterferenceVO() *ApInterferenceVO`

NewApInterferenceVO instantiates a new ApInterferenceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApInterferenceVOWithDefaults

`func NewApInterferenceVOWithDefaults() *ApInterferenceVO`

NewApInterferenceVOWithDefaults instantiates a new ApInterferenceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice

`func (o *ApInterferenceVO) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ApInterferenceVO) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ApInterferenceVO) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ApInterferenceVO) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetInterUtil

`func (o *ApInterferenceVO) GetInterUtil() int32`

GetInterUtil returns the InterUtil field if non-nil, zero value otherwise.

### GetInterUtilOk

`func (o *ApInterferenceVO) GetInterUtilOk() (*int32, bool)`

GetInterUtilOk returns a tuple with the InterUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterUtil

`func (o *ApInterferenceVO) SetInterUtil(v int32)`

SetInterUtil sets InterUtil field to given value.

### HasInterUtil

`func (o *ApInterferenceVO) HasInterUtil() bool`

HasInterUtil returns a boolean if a field has been set.

### GetInterference

`func (o *ApInterferenceVO) GetInterference() int32`

GetInterference returns the Interference field if non-nil, zero value otherwise.

### GetInterferenceOk

`func (o *ApInterferenceVO) GetInterferenceOk() (*int32, bool)`

GetInterferenceOk returns a tuple with the Interference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterference

`func (o *ApInterferenceVO) SetInterference(v int32)`

SetInterference sets Interference field to given value.

### HasInterference

`func (o *ApInterferenceVO) HasInterference() bool`

HasInterference returns a boolean if a field has been set.

### GetModel

`func (o *ApInterferenceVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApInterferenceVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApInterferenceVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApInterferenceVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApInterferenceVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApInterferenceVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApInterferenceVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApInterferenceVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetRxUtil

`func (o *ApInterferenceVO) GetRxUtil() int32`

GetRxUtil returns the RxUtil field if non-nil, zero value otherwise.

### GetRxUtilOk

`func (o *ApInterferenceVO) GetRxUtilOk() (*int32, bool)`

GetRxUtilOk returns a tuple with the RxUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRxUtil

`func (o *ApInterferenceVO) SetRxUtil(v int32)`

SetRxUtil sets RxUtil field to given value.

### HasRxUtil

`func (o *ApInterferenceVO) HasRxUtil() bool`

HasRxUtil returns a boolean if a field has been set.

### GetTxUtil

`func (o *ApInterferenceVO) GetTxUtil() int32`

GetTxUtil returns the TxUtil field if non-nil, zero value otherwise.

### GetTxUtilOk

`func (o *ApInterferenceVO) GetTxUtilOk() (*int32, bool)`

GetTxUtilOk returns a tuple with the TxUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxUtil

`func (o *ApInterferenceVO) SetTxUtil(v int32)`

SetTxUtil sets TxUtil field to given value.

### HasTxUtil

`func (o *ApInterferenceVO) HasTxUtil() bool`

HasTxUtil returns a boolean if a field has been set.

### GetType

`func (o *ApInterferenceVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApInterferenceVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApInterferenceVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApInterferenceVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


