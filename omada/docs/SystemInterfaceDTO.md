# SystemInterfaceDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | Pointer to **int32** | Management System Interface ID, its value should be within the range of 0-3. | [optional] 
**InterfaceType** | **string** | Interface Type should be a value as follows: MANAGEMENT, NONE, VLAN, ROUTED_PORT, PORT_CHANNEL. | 
**InterfaceValue** | Pointer to **string** | Interface Value | [optional] 

## Methods

### NewSystemInterfaceDTO

`func NewSystemInterfaceDTO(interfaceType string, ) *SystemInterfaceDTO`

NewSystemInterfaceDTO instantiates a new SystemInterfaceDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInterfaceDTOWithDefaults

`func NewSystemInterfaceDTOWithDefaults() *SystemInterfaceDTO`

NewSystemInterfaceDTOWithDefaults instantiates a new SystemInterfaceDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *SystemInterfaceDTO) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *SystemInterfaceDTO) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *SystemInterfaceDTO) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *SystemInterfaceDTO) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *SystemInterfaceDTO) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *SystemInterfaceDTO) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *SystemInterfaceDTO) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.


### GetInterfaceValue

`func (o *SystemInterfaceDTO) GetInterfaceValue() string`

GetInterfaceValue returns the InterfaceValue field if non-nil, zero value otherwise.

### GetInterfaceValueOk

`func (o *SystemInterfaceDTO) GetInterfaceValueOk() (*string, bool)`

GetInterfaceValueOk returns a tuple with the InterfaceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceValue

`func (o *SystemInterfaceDTO) SetInterfaceValue(v string)`

SetInterfaceValue sets InterfaceValue field to given value.

### HasInterfaceValue

`func (o *SystemInterfaceDTO) HasInterfaceValue() bool`

HasInterfaceValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


