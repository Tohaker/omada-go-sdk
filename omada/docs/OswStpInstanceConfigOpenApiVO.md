# OswStpInstanceConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | ID | 
**Priority** | **int32** | Priority | 
**Vlan** | **string** | Parameter [vlan] should be between 1 and 4094. | 

## Methods

### NewOswStpInstanceConfigOpenApiVO

`func NewOswStpInstanceConfigOpenApiVO(id int32, priority int32, vlan string, ) *OswStpInstanceConfigOpenApiVO`

NewOswStpInstanceConfigOpenApiVO instantiates a new OswStpInstanceConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStpInstanceConfigOpenApiVOWithDefaults

`func NewOswStpInstanceConfigOpenApiVOWithDefaults() *OswStpInstanceConfigOpenApiVO`

NewOswStpInstanceConfigOpenApiVOWithDefaults instantiates a new OswStpInstanceConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OswStpInstanceConfigOpenApiVO) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswStpInstanceConfigOpenApiVO) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswStpInstanceConfigOpenApiVO) SetId(v int32)`

SetId sets Id field to given value.


### GetPriority

`func (o *OswStpInstanceConfigOpenApiVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *OswStpInstanceConfigOpenApiVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *OswStpInstanceConfigOpenApiVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetVlan

`func (o *OswStpInstanceConfigOpenApiVO) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *OswStpInstanceConfigOpenApiVO) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *OswStpInstanceConfigOpenApiVO) SetVlan(v string)`

SetVlan sets Vlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


