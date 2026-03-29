# ServicePortProfileDeleteDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** | Service port profile Id. The service port profile ID should be within the range of 1 to 512 and it can be obtained from \&quot;Get service port profile list\&quot; interface. Up to 50 entries are allowed for the ids list. | [optional] 

## Methods

### NewServicePortProfileDeleteDTO

`func NewServicePortProfileDeleteDTO() *ServicePortProfileDeleteDTO`

NewServicePortProfileDeleteDTO instantiates a new ServicePortProfileDeleteDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortProfileDeleteDTOWithDefaults

`func NewServicePortProfileDeleteDTOWithDefaults() *ServicePortProfileDeleteDTO`

NewServicePortProfileDeleteDTOWithDefaults instantiates a new ServicePortProfileDeleteDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *ServicePortProfileDeleteDTO) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *ServicePortProfileDeleteDTO) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *ServicePortProfileDeleteDTO) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *ServicePortProfileDeleteDTO) HasKeys() bool`

HasKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


