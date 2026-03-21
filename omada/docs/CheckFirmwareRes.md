# CheckFirmwareRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMacList** | Pointer to **[]string** | List of the ap MAC address with firmware updates. E.g. AA-BB-CC-DD-11-22 | [optional] 
**Finished** | Pointer to **bool** | Whether the task is complete | [optional] 
**GatewayMacList** | Pointer to **[]string** | List of gateway MAC address with firmware updates. E.g. AA-BB-CC-DD-11-22 | [optional] 
**SwitchMacList** | Pointer to **[]string** | List of switch MAC address with firmware updates. E.g. AA-BB-CC-DD-11-22 | [optional] 

## Methods

### NewCheckFirmwareRes

`func NewCheckFirmwareRes() *CheckFirmwareRes`

NewCheckFirmwareRes instantiates a new CheckFirmwareRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckFirmwareResWithDefaults

`func NewCheckFirmwareResWithDefaults() *CheckFirmwareRes`

NewCheckFirmwareResWithDefaults instantiates a new CheckFirmwareRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMacList

`func (o *CheckFirmwareRes) GetApMacList() []string`

GetApMacList returns the ApMacList field if non-nil, zero value otherwise.

### GetApMacListOk

`func (o *CheckFirmwareRes) GetApMacListOk() (*[]string, bool)`

GetApMacListOk returns a tuple with the ApMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMacList

`func (o *CheckFirmwareRes) SetApMacList(v []string)`

SetApMacList sets ApMacList field to given value.

### HasApMacList

`func (o *CheckFirmwareRes) HasApMacList() bool`

HasApMacList returns a boolean if a field has been set.

### GetFinished

`func (o *CheckFirmwareRes) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *CheckFirmwareRes) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *CheckFirmwareRes) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *CheckFirmwareRes) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetGatewayMacList

`func (o *CheckFirmwareRes) GetGatewayMacList() []string`

GetGatewayMacList returns the GatewayMacList field if non-nil, zero value otherwise.

### GetGatewayMacListOk

`func (o *CheckFirmwareRes) GetGatewayMacListOk() (*[]string, bool)`

GetGatewayMacListOk returns a tuple with the GatewayMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayMacList

`func (o *CheckFirmwareRes) SetGatewayMacList(v []string)`

SetGatewayMacList sets GatewayMacList field to given value.

### HasGatewayMacList

`func (o *CheckFirmwareRes) HasGatewayMacList() bool`

HasGatewayMacList returns a boolean if a field has been set.

### GetSwitchMacList

`func (o *CheckFirmwareRes) GetSwitchMacList() []string`

GetSwitchMacList returns the SwitchMacList field if non-nil, zero value otherwise.

### GetSwitchMacListOk

`func (o *CheckFirmwareRes) GetSwitchMacListOk() (*[]string, bool)`

GetSwitchMacListOk returns a tuple with the SwitchMacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchMacList

`func (o *CheckFirmwareRes) SetSwitchMacList(v []string)`

SetSwitchMacList sets SwitchMacList field to given value.

### HasSwitchMacList

`func (o *CheckFirmwareRes) HasSwitchMacList() bool`

HasSwitchMacList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


