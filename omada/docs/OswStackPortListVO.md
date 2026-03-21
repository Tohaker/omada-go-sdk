# OswStackPortListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StackId** | Pointer to **string** | Stack Id | [optional] 
**StandardPorts** | Pointer to **[]string** | Standard port should be in the format of unit/slot/portId. e.g. 1/0/1 | [optional] 

## Methods

### NewOswStackPortListVO

`func NewOswStackPortListVO() *OswStackPortListVO`

NewOswStackPortListVO instantiates a new OswStackPortListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackPortListVOWithDefaults

`func NewOswStackPortListVOWithDefaults() *OswStackPortListVO`

NewOswStackPortListVOWithDefaults instantiates a new OswStackPortListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStackId

`func (o *OswStackPortListVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswStackPortListVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswStackPortListVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswStackPortListVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStandardPorts

`func (o *OswStackPortListVO) GetStandardPorts() []string`

GetStandardPorts returns the StandardPorts field if non-nil, zero value otherwise.

### GetStandardPortsOk

`func (o *OswStackPortListVO) GetStandardPortsOk() (*[]string, bool)`

GetStandardPortsOk returns a tuple with the StandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPorts

`func (o *OswStackPortListVO) SetStandardPorts(v []string)`

SetStandardPorts sets StandardPorts field to given value.

### HasStandardPorts

`func (o *OswStackPortListVO) HasStandardPorts() bool`

HasStandardPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


