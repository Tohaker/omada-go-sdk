# SetTagResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** | 0: success, 1: failed | [optional] 

## Methods

### NewSetTagResultOpenApiVO

`func NewSetTagResultOpenApiVO() *SetTagResultOpenApiVO`

NewSetTagResultOpenApiVO instantiates a new SetTagResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetTagResultOpenApiVOWithDefaults

`func NewSetTagResultOpenApiVOWithDefaults() *SetTagResultOpenApiVO`

NewSetTagResultOpenApiVOWithDefaults instantiates a new SetTagResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *SetTagResultOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SetTagResultOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SetTagResultOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SetTagResultOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *SetTagResultOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SetTagResultOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SetTagResultOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SetTagResultOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


