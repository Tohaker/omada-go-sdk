# ApMdnsRuleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientVlan** | **string** | Client Network VLAN. ClientVlan should be within the range of 1 to 4094. Enter one or multiple VLANs. For example: 1,2-100 | 
**ServiceVlan** | **string** | Services Network VLAN. ServiceVlan should be within the range of 1 to 4094. Enter only one VLAN | 

## Methods

### NewApMdnsRuleOpenApiVO

`func NewApMdnsRuleOpenApiVO(clientVlan string, serviceVlan string, ) *ApMdnsRuleOpenApiVO`

NewApMdnsRuleOpenApiVO instantiates a new ApMdnsRuleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMdnsRuleOpenApiVOWithDefaults

`func NewApMdnsRuleOpenApiVOWithDefaults() *ApMdnsRuleOpenApiVO`

NewApMdnsRuleOpenApiVOWithDefaults instantiates a new ApMdnsRuleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientVlan

`func (o *ApMdnsRuleOpenApiVO) GetClientVlan() string`

GetClientVlan returns the ClientVlan field if non-nil, zero value otherwise.

### GetClientVlanOk

`func (o *ApMdnsRuleOpenApiVO) GetClientVlanOk() (*string, bool)`

GetClientVlanOk returns a tuple with the ClientVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientVlan

`func (o *ApMdnsRuleOpenApiVO) SetClientVlan(v string)`

SetClientVlan sets ClientVlan field to given value.


### GetServiceVlan

`func (o *ApMdnsRuleOpenApiVO) GetServiceVlan() string`

GetServiceVlan returns the ServiceVlan field if non-nil, zero value otherwise.

### GetServiceVlanOk

`func (o *ApMdnsRuleOpenApiVO) GetServiceVlanOk() (*string, bool)`

GetServiceVlanOk returns a tuple with the ServiceVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVlan

`func (o *ApMdnsRuleOpenApiVO) SetServiceVlan(v string)`

SetServiceVlan sets ServiceVlan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


