# OnuGlobalConfigStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OnuIsolation** | Pointer to **string** | Whether to enable ONU isolation,onuIsolation should be a value as follows:ENABLE,DISABLE | [optional] 
**SupportOnuIsolation** | Pointer to **string** | Whether the ONU support isolation.SupportOnuIsolation should be a value as follows:ENABLE,DISABLE | [optional] 

## Methods

### NewOnuGlobalConfigStatusDTO

`func NewOnuGlobalConfigStatusDTO() *OnuGlobalConfigStatusDTO`

NewOnuGlobalConfigStatusDTO instantiates a new OnuGlobalConfigStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuGlobalConfigStatusDTOWithDefaults

`func NewOnuGlobalConfigStatusDTOWithDefaults() *OnuGlobalConfigStatusDTO`

NewOnuGlobalConfigStatusDTOWithDefaults instantiates a new OnuGlobalConfigStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnuIsolation

`func (o *OnuGlobalConfigStatusDTO) GetOnuIsolation() string`

GetOnuIsolation returns the OnuIsolation field if non-nil, zero value otherwise.

### GetOnuIsolationOk

`func (o *OnuGlobalConfigStatusDTO) GetOnuIsolationOk() (*string, bool)`

GetOnuIsolationOk returns a tuple with the OnuIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuIsolation

`func (o *OnuGlobalConfigStatusDTO) SetOnuIsolation(v string)`

SetOnuIsolation sets OnuIsolation field to given value.

### HasOnuIsolation

`func (o *OnuGlobalConfigStatusDTO) HasOnuIsolation() bool`

HasOnuIsolation returns a boolean if a field has been set.

### GetSupportOnuIsolation

`func (o *OnuGlobalConfigStatusDTO) GetSupportOnuIsolation() string`

GetSupportOnuIsolation returns the SupportOnuIsolation field if non-nil, zero value otherwise.

### GetSupportOnuIsolationOk

`func (o *OnuGlobalConfigStatusDTO) GetSupportOnuIsolationOk() (*string, bool)`

GetSupportOnuIsolationOk returns a tuple with the SupportOnuIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportOnuIsolation

`func (o *OnuGlobalConfigStatusDTO) SetSupportOnuIsolation(v string)`

SetSupportOnuIsolation sets SupportOnuIsolation field to given value.

### HasSupportOnuIsolation

`func (o *OnuGlobalConfigStatusDTO) HasSupportOnuIsolation() bool`

HasSupportOnuIsolation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


