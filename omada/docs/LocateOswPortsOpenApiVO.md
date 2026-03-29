# LocateOswPortsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocateEnable** | **bool** | Whether locate function is enabled | 
**PortListVOs** | Pointer to [**[]OswNormalPortListVO**](OswNormalPortListVO.md) | List of Switch MAC and ports. | [optional] 
**StackPortListVOs** | Pointer to [**[]OswStackPortListVO**](OswStackPortListVO.md) | List of Stack ID and Standard Ports. | [optional] 

## Methods

### NewLocateOswPortsOpenApiVO

`func NewLocateOswPortsOpenApiVO(locateEnable bool, ) *LocateOswPortsOpenApiVO`

NewLocateOswPortsOpenApiVO instantiates a new LocateOswPortsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocateOswPortsOpenApiVOWithDefaults

`func NewLocateOswPortsOpenApiVOWithDefaults() *LocateOswPortsOpenApiVO`

NewLocateOswPortsOpenApiVOWithDefaults instantiates a new LocateOswPortsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocateEnable

`func (o *LocateOswPortsOpenApiVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *LocateOswPortsOpenApiVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *LocateOswPortsOpenApiVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.


### GetPortListVOs

`func (o *LocateOswPortsOpenApiVO) GetPortListVOs() []OswNormalPortListVO`

GetPortListVOs returns the PortListVOs field if non-nil, zero value otherwise.

### GetPortListVOsOk

`func (o *LocateOswPortsOpenApiVO) GetPortListVOsOk() (*[]OswNormalPortListVO, bool)`

GetPortListVOsOk returns a tuple with the PortListVOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortListVOs

`func (o *LocateOswPortsOpenApiVO) SetPortListVOs(v []OswNormalPortListVO)`

SetPortListVOs sets PortListVOs field to given value.

### HasPortListVOs

`func (o *LocateOswPortsOpenApiVO) HasPortListVOs() bool`

HasPortListVOs returns a boolean if a field has been set.

### GetStackPortListVOs

`func (o *LocateOswPortsOpenApiVO) GetStackPortListVOs() []OswStackPortListVO`

GetStackPortListVOs returns the StackPortListVOs field if non-nil, zero value otherwise.

### GetStackPortListVOsOk

`func (o *LocateOswPortsOpenApiVO) GetStackPortListVOsOk() (*[]OswStackPortListVO, bool)`

GetStackPortListVOsOk returns a tuple with the StackPortListVOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortListVOs

`func (o *LocateOswPortsOpenApiVO) SetStackPortListVOs(v []OswStackPortListVO)`

SetStackPortListVOs sets StackPortListVOs field to given value.

### HasStackPortListVOs

`func (o *LocateOswPortsOpenApiVO) HasStackPortListVOs() bool`

HasStackPortListVOs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


