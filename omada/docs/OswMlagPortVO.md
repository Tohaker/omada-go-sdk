# OswMlagPortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** |  | [optional] 
**PortStatus** | Pointer to [**OswPortStatusVO**](OswPortStatusVO.md) |  | [optional] 
**Selectable** | Pointer to **bool** | whether the port can be selected to configure DAD Port or Peer Link Port. | [optional] 
**StandardPort** | Pointer to [**OswStandPortVO**](OswStandPortVO.md) |  | [optional] 

## Methods

### NewOswMlagPortVO

`func NewOswMlagPortVO() *OswMlagPortVO`

NewOswMlagPortVO instantiates a new OswMlagPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswMlagPortVOWithDefaults

`func NewOswMlagPortVOWithDefaults() *OswMlagPortVO`

NewOswMlagPortVOWithDefaults instantiates a new OswMlagPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *OswMlagPortVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswMlagPortVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswMlagPortVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswMlagPortVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortStatus

`func (o *OswMlagPortVO) GetPortStatus() OswPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *OswMlagPortVO) GetPortStatusOk() (*OswPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *OswMlagPortVO) SetPortStatus(v OswPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *OswMlagPortVO) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetSelectable

`func (o *OswMlagPortVO) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *OswMlagPortVO) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *OswMlagPortVO) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *OswMlagPortVO) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswMlagPortVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswMlagPortVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswMlagPortVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswMlagPortVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


