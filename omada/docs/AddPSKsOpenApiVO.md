# AddPSKsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PpskList** | [**[]PpskSetting**](PpskSetting.md) | PSK entries that need to be added to the PPSK Profile | 

## Methods

### NewAddPSKsOpenApiVO

`func NewAddPSKsOpenApiVO(ppskList []PpskSetting, ) *AddPSKsOpenApiVO`

NewAddPSKsOpenApiVO instantiates a new AddPSKsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPSKsOpenApiVOWithDefaults

`func NewAddPSKsOpenApiVOWithDefaults() *AddPSKsOpenApiVO`

NewAddPSKsOpenApiVOWithDefaults instantiates a new AddPSKsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPpskList

`func (o *AddPSKsOpenApiVO) GetPpskList() []PpskSetting`

GetPpskList returns the PpskList field if non-nil, zero value otherwise.

### GetPpskListOk

`func (o *AddPSKsOpenApiVO) GetPpskListOk() (*[]PpskSetting, bool)`

GetPpskListOk returns a tuple with the PpskList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpskList

`func (o *AddPSKsOpenApiVO) SetPpskList(v []PpskSetting)`

SetPpskList sets PpskList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


