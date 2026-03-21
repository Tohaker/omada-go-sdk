# CheckFailedDeviceInfoForVlanVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to [**ResErrorCodeVO**](ResErrorCodeVO.md) |  | [optional] 
**Mac** | Pointer to **string** | check failed device mac | [optional] 
**StackId** | Pointer to **string** | check failed stack id | [optional] 

## Methods

### NewCheckFailedDeviceInfoForVlanVO

`func NewCheckFailedDeviceInfoForVlanVO() *CheckFailedDeviceInfoForVlanVO`

NewCheckFailedDeviceInfoForVlanVO instantiates a new CheckFailedDeviceInfoForVlanVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckFailedDeviceInfoForVlanVOWithDefaults

`func NewCheckFailedDeviceInfoForVlanVOWithDefaults() *CheckFailedDeviceInfoForVlanVO`

NewCheckFailedDeviceInfoForVlanVOWithDefaults instantiates a new CheckFailedDeviceInfoForVlanVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *CheckFailedDeviceInfoForVlanVO) GetErrorCode() ResErrorCodeVO`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CheckFailedDeviceInfoForVlanVO) GetErrorCodeOk() (*ResErrorCodeVO, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CheckFailedDeviceInfoForVlanVO) SetErrorCode(v ResErrorCodeVO)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CheckFailedDeviceInfoForVlanVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMac

`func (o *CheckFailedDeviceInfoForVlanVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CheckFailedDeviceInfoForVlanVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CheckFailedDeviceInfoForVlanVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *CheckFailedDeviceInfoForVlanVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStackId

`func (o *CheckFailedDeviceInfoForVlanVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CheckFailedDeviceInfoForVlanVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CheckFailedDeviceInfoForVlanVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *CheckFailedDeviceInfoForVlanVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


