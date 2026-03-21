# ServicePortProfileDetailDeleteDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePortId** | Pointer to **[]string** | Service port ID List. It can be obtained from\&quot;Get service port profile detail list\&quot; interface. Up to 50 entries are allowed for the ids list.  | [optional] 
**ServicePortProfileId** | Pointer to **string** | Service port profile ID. The servicePortProfileId should be within the range of 1 to 127. It can be obtained from \&quot;Get service port profile detail\&quot;. | [optional] 

## Methods

### NewServicePortProfileDetailDeleteDTO

`func NewServicePortProfileDetailDeleteDTO() *ServicePortProfileDetailDeleteDTO`

NewServicePortProfileDetailDeleteDTO instantiates a new ServicePortProfileDetailDeleteDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePortProfileDetailDeleteDTOWithDefaults

`func NewServicePortProfileDetailDeleteDTOWithDefaults() *ServicePortProfileDetailDeleteDTO`

NewServicePortProfileDetailDeleteDTOWithDefaults instantiates a new ServicePortProfileDetailDeleteDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePortId

`func (o *ServicePortProfileDetailDeleteDTO) GetServicePortId() []string`

GetServicePortId returns the ServicePortId field if non-nil, zero value otherwise.

### GetServicePortIdOk

`func (o *ServicePortProfileDetailDeleteDTO) GetServicePortIdOk() (*[]string, bool)`

GetServicePortIdOk returns a tuple with the ServicePortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortId

`func (o *ServicePortProfileDetailDeleteDTO) SetServicePortId(v []string)`

SetServicePortId sets ServicePortId field to given value.

### HasServicePortId

`func (o *ServicePortProfileDetailDeleteDTO) HasServicePortId() bool`

HasServicePortId returns a boolean if a field has been set.

### GetServicePortProfileId

`func (o *ServicePortProfileDetailDeleteDTO) GetServicePortProfileId() string`

GetServicePortProfileId returns the ServicePortProfileId field if non-nil, zero value otherwise.

### GetServicePortProfileIdOk

`func (o *ServicePortProfileDetailDeleteDTO) GetServicePortProfileIdOk() (*string, bool)`

GetServicePortProfileIdOk returns a tuple with the ServicePortProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortProfileId

`func (o *ServicePortProfileDetailDeleteDTO) SetServicePortProfileId(v string)`

SetServicePortProfileId sets ServicePortProfileId field to given value.

### HasServicePortProfileId

`func (o *ServicePortProfileDetailDeleteDTO) HasServicePortProfileId() bool`

HasServicePortProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


