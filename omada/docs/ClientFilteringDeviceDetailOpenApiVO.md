# ClientFilteringDeviceDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device MAC | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Stack** | Pointer to **bool** | Does this device belong to a stack group | [optional] 
**StackName** | Pointer to **string** | Stack group name | [optional] 

## Methods

### NewClientFilteringDeviceDetailOpenApiVO

`func NewClientFilteringDeviceDetailOpenApiVO() *ClientFilteringDeviceDetailOpenApiVO`

NewClientFilteringDeviceDetailOpenApiVO instantiates a new ClientFilteringDeviceDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientFilteringDeviceDetailOpenApiVOWithDefaults

`func NewClientFilteringDeviceDetailOpenApiVOWithDefaults() *ClientFilteringDeviceDetailOpenApiVO`

NewClientFilteringDeviceDetailOpenApiVOWithDefaults instantiates a new ClientFilteringDeviceDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *ClientFilteringDeviceDetailOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientFilteringDeviceDetailOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientFilteringDeviceDetailOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientFilteringDeviceDetailOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *ClientFilteringDeviceDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientFilteringDeviceDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientFilteringDeviceDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientFilteringDeviceDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStack

`func (o *ClientFilteringDeviceDetailOpenApiVO) GetStack() bool`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *ClientFilteringDeviceDetailOpenApiVO) GetStackOk() (*bool, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *ClientFilteringDeviceDetailOpenApiVO) SetStack(v bool)`

SetStack sets Stack field to given value.

### HasStack

`func (o *ClientFilteringDeviceDetailOpenApiVO) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetStackName

`func (o *ClientFilteringDeviceDetailOpenApiVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *ClientFilteringDeviceDetailOpenApiVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *ClientFilteringDeviceDetailOpenApiVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *ClientFilteringDeviceDetailOpenApiVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


