# OnuIsolationStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsolateAllVlan** | Pointer to **string** | Whether to isolate all ONUs within the VLANs.IsolateAllVlan should be a value as follows:DISABLE,ENABLE.Effective only when onuIsolation is set to ENABLE. | [optional] 
**OnuIsolation** | Pointer to **string** |  | [optional] 
**VlanList** | Pointer to **string** | vlanList should be a value as follows: \&quot;1,3-5,7\&quot;.Effective only when onuIsolation is ENABLE and isolateAllVlan is DISABLE. | [optional] 

## Methods

### NewOnuIsolationStatusDTO

`func NewOnuIsolationStatusDTO() *OnuIsolationStatusDTO`

NewOnuIsolationStatusDTO instantiates a new OnuIsolationStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnuIsolationStatusDTOWithDefaults

`func NewOnuIsolationStatusDTOWithDefaults() *OnuIsolationStatusDTO`

NewOnuIsolationStatusDTOWithDefaults instantiates a new OnuIsolationStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsolateAllVlan

`func (o *OnuIsolationStatusDTO) GetIsolateAllVlan() string`

GetIsolateAllVlan returns the IsolateAllVlan field if non-nil, zero value otherwise.

### GetIsolateAllVlanOk

`func (o *OnuIsolationStatusDTO) GetIsolateAllVlanOk() (*string, bool)`

GetIsolateAllVlanOk returns a tuple with the IsolateAllVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolateAllVlan

`func (o *OnuIsolationStatusDTO) SetIsolateAllVlan(v string)`

SetIsolateAllVlan sets IsolateAllVlan field to given value.

### HasIsolateAllVlan

`func (o *OnuIsolationStatusDTO) HasIsolateAllVlan() bool`

HasIsolateAllVlan returns a boolean if a field has been set.

### GetOnuIsolation

`func (o *OnuIsolationStatusDTO) GetOnuIsolation() string`

GetOnuIsolation returns the OnuIsolation field if non-nil, zero value otherwise.

### GetOnuIsolationOk

`func (o *OnuIsolationStatusDTO) GetOnuIsolationOk() (*string, bool)`

GetOnuIsolationOk returns a tuple with the OnuIsolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuIsolation

`func (o *OnuIsolationStatusDTO) SetOnuIsolation(v string)`

SetOnuIsolation sets OnuIsolation field to given value.

### HasOnuIsolation

`func (o *OnuIsolationStatusDTO) HasOnuIsolation() bool`

HasOnuIsolation returns a boolean if a field has been set.

### GetVlanList

`func (o *OnuIsolationStatusDTO) GetVlanList() string`

GetVlanList returns the VlanList field if non-nil, zero value otherwise.

### GetVlanListOk

`func (o *OnuIsolationStatusDTO) GetVlanListOk() (*string, bool)`

GetVlanListOk returns a tuple with the VlanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanList

`func (o *OnuIsolationStatusDTO) SetVlanList(v string)`

SetVlanList sets VlanList field to given value.

### HasVlanList

`func (o *OnuIsolationStatusDTO) HasVlanList() bool`

HasVlanList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


