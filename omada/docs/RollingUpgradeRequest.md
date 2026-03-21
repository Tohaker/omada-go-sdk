# RollingUpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | MAC list of devices. E.g. AA-BB-CC-DD-11-22 | [optional] 

## Methods

### NewRollingUpgradeRequest

`func NewRollingUpgradeRequest() *RollingUpgradeRequest`

NewRollingUpgradeRequest instantiates a new RollingUpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRollingUpgradeRequestWithDefaults

`func NewRollingUpgradeRequestWithDefaults() *RollingUpgradeRequest`

NewRollingUpgradeRequestWithDefaults instantiates a new RollingUpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *RollingUpgradeRequest) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *RollingUpgradeRequest) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *RollingUpgradeRequest) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *RollingUpgradeRequest) HasMacs() bool`

HasMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


