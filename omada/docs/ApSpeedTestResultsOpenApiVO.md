# ApSpeedTestResultsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MainTesting** | Pointer to **bool** | Whether the main AP is currently measuring speed. | [optional] 
**SourceAP** | Pointer to **string** | The MAC address of the AP that initiates the speed measurement when the current device is measuring. | [optional] 
**SpeedTestResult** | Pointer to [**map[string]ApSpeedTestResultOpenApiVO**](ApSpeedTestResultOpenApiVO.md) | Speed measurement results of uplink and downlink devices. | [optional] 
**TargetAP** | Pointer to **string** | The MAC address of the AP whose speed is being measured when the current device is measuring. | [optional] 

## Methods

### NewApSpeedTestResultsOpenApiVO

`func NewApSpeedTestResultsOpenApiVO() *ApSpeedTestResultsOpenApiVO`

NewApSpeedTestResultsOpenApiVO instantiates a new ApSpeedTestResultsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApSpeedTestResultsOpenApiVOWithDefaults

`func NewApSpeedTestResultsOpenApiVOWithDefaults() *ApSpeedTestResultsOpenApiVO`

NewApSpeedTestResultsOpenApiVOWithDefaults instantiates a new ApSpeedTestResultsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMainTesting

`func (o *ApSpeedTestResultsOpenApiVO) GetMainTesting() bool`

GetMainTesting returns the MainTesting field if non-nil, zero value otherwise.

### GetMainTestingOk

`func (o *ApSpeedTestResultsOpenApiVO) GetMainTestingOk() (*bool, bool)`

GetMainTestingOk returns a tuple with the MainTesting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainTesting

`func (o *ApSpeedTestResultsOpenApiVO) SetMainTesting(v bool)`

SetMainTesting sets MainTesting field to given value.

### HasMainTesting

`func (o *ApSpeedTestResultsOpenApiVO) HasMainTesting() bool`

HasMainTesting returns a boolean if a field has been set.

### GetSourceAP

`func (o *ApSpeedTestResultsOpenApiVO) GetSourceAP() string`

GetSourceAP returns the SourceAP field if non-nil, zero value otherwise.

### GetSourceAPOk

`func (o *ApSpeedTestResultsOpenApiVO) GetSourceAPOk() (*string, bool)`

GetSourceAPOk returns a tuple with the SourceAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAP

`func (o *ApSpeedTestResultsOpenApiVO) SetSourceAP(v string)`

SetSourceAP sets SourceAP field to given value.

### HasSourceAP

`func (o *ApSpeedTestResultsOpenApiVO) HasSourceAP() bool`

HasSourceAP returns a boolean if a field has been set.

### GetSpeedTestResult

`func (o *ApSpeedTestResultsOpenApiVO) GetSpeedTestResult() map[string]ApSpeedTestResultOpenApiVO`

GetSpeedTestResult returns the SpeedTestResult field if non-nil, zero value otherwise.

### GetSpeedTestResultOk

`func (o *ApSpeedTestResultsOpenApiVO) GetSpeedTestResultOk() (*map[string]ApSpeedTestResultOpenApiVO, bool)`

GetSpeedTestResultOk returns a tuple with the SpeedTestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeedTestResult

`func (o *ApSpeedTestResultsOpenApiVO) SetSpeedTestResult(v map[string]ApSpeedTestResultOpenApiVO)`

SetSpeedTestResult sets SpeedTestResult field to given value.

### HasSpeedTestResult

`func (o *ApSpeedTestResultsOpenApiVO) HasSpeedTestResult() bool`

HasSpeedTestResult returns a boolean if a field has been set.

### GetTargetAP

`func (o *ApSpeedTestResultsOpenApiVO) GetTargetAP() string`

GetTargetAP returns the TargetAP field if non-nil, zero value otherwise.

### GetTargetAPOk

`func (o *ApSpeedTestResultsOpenApiVO) GetTargetAPOk() (*string, bool)`

GetTargetAPOk returns a tuple with the TargetAP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetAP

`func (o *ApSpeedTestResultsOpenApiVO) SetTargetAP(v string)`

SetTargetAP sets TargetAP field to given value.

### HasTargetAP

`func (o *ApSpeedTestResultsOpenApiVO) HasTargetAP() bool`

HasTargetAP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


