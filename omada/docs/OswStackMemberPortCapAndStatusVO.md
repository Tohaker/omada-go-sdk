# OswStackMemberPortCapAndStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortStatus** | Pointer to [**OswPortStatusVO**](OswPortStatusVO.md) |  | [optional] 
**StackSupport** | Pointer to **bool** | Stack Support | [optional] 
**StandardPort** | Pointer to [**OswStandPortVO**](OswStandPortVO.md) |  | [optional] 

## Methods

### NewOswStackMemberPortCapAndStatusVO

`func NewOswStackMemberPortCapAndStatusVO() *OswStackMemberPortCapAndStatusVO`

NewOswStackMemberPortCapAndStatusVO instantiates a new OswStackMemberPortCapAndStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackMemberPortCapAndStatusVOWithDefaults

`func NewOswStackMemberPortCapAndStatusVOWithDefaults() *OswStackMemberPortCapAndStatusVO`

NewOswStackMemberPortCapAndStatusVOWithDefaults instantiates a new OswStackMemberPortCapAndStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortStatus

`func (o *OswStackMemberPortCapAndStatusVO) GetPortStatus() OswPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *OswStackMemberPortCapAndStatusVO) GetPortStatusOk() (*OswPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *OswStackMemberPortCapAndStatusVO) SetPortStatus(v OswPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *OswStackMemberPortCapAndStatusVO) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetStackSupport

`func (o *OswStackMemberPortCapAndStatusVO) GetStackSupport() bool`

GetStackSupport returns the StackSupport field if non-nil, zero value otherwise.

### GetStackSupportOk

`func (o *OswStackMemberPortCapAndStatusVO) GetStackSupportOk() (*bool, bool)`

GetStackSupportOk returns a tuple with the StackSupport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackSupport

`func (o *OswStackMemberPortCapAndStatusVO) SetStackSupport(v bool)`

SetStackSupport sets StackSupport field to given value.

### HasStackSupport

`func (o *OswStackMemberPortCapAndStatusVO) HasStackSupport() bool`

HasStackSupport returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswStackMemberPortCapAndStatusVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswStackMemberPortCapAndStatusVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswStackMemberPortCapAndStatusVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswStackMemberPortCapAndStatusVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


