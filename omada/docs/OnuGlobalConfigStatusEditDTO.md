# OnuGlobalConfigStatusEditDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnuIsolation** | **string** | Whether to enable ONU isolation,onuIsolation should be a value as follows:ENABLE,DISABLE | 

## Methods

### NewOnuGlobalConfigStatusEditDTO

`func NewOnuGlobalConfigStatusEditDTO(onuIsolation string, ) *OnuGlobalConfigStatusEditDTO`

NewOnuGlobalConfigStatusEditDTO instantiates a new OnuGlobalConfigStatusEditDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuGlobalConfigStatusEditDTOWithDefaults

`func NewOnuGlobalConfigStatusEditDTOWithDefaults() *OnuGlobalConfigStatusEditDTO`

NewOnuGlobalConfigStatusEditDTOWithDefaults instantiates a new OnuGlobalConfigStatusEditDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnuIsolation

`func (o *OnuGlobalConfigStatusEditDTO) GetOnuIsolation() string`

GetOnuIsolation returns the OnuIsolation field if non-nil, zero value otherwise.

### GetOnuIsolationOk

`func (o *OnuGlobalConfigStatusEditDTO) GetOnuIsolationOk() (*string, bool)`

GetOnuIsolationOk returns a tuple with the OnuIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuIsolation

`func (o *OnuGlobalConfigStatusEditDTO) SetOnuIsolation(v string)`

SetOnuIsolation sets OnuIsolation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


