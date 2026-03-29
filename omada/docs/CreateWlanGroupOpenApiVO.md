# CreateWlanGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clone** | **bool** | Whether to clone SSID list from other WLAN group. | 
**ClonedWlanId** | Pointer to **string** | WLAN group ID that needs to be cloned. Parameter [clonedWlanId] should not be null when Parameter [clone] is true. | [optional] 
**Name** | **string** | WLAN group name should contain 1 to 128 characters. | 

## Methods

### NewCreateWlanGroupOpenApiVO

`func NewCreateWlanGroupOpenApiVO(clone bool, name string, ) *CreateWlanGroupOpenApiVO`

NewCreateWlanGroupOpenApiVO instantiates a new CreateWlanGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWlanGroupOpenApiVOWithDefaults

`func NewCreateWlanGroupOpenApiVOWithDefaults() *CreateWlanGroupOpenApiVO`

NewCreateWlanGroupOpenApiVOWithDefaults instantiates a new CreateWlanGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClone

`func (o *CreateWlanGroupOpenApiVO) GetClone() bool`

GetClone returns the Clone field if non-nil, zero value otherwise.

### GetCloneOk

`func (o *CreateWlanGroupOpenApiVO) GetCloneOk() (*bool, bool)`

GetCloneOk returns a tuple with the Clone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClone

`func (o *CreateWlanGroupOpenApiVO) SetClone(v bool)`

SetClone sets Clone field to given value.


### GetClonedWlanId

`func (o *CreateWlanGroupOpenApiVO) GetClonedWlanId() string`

GetClonedWlanId returns the ClonedWlanId field if non-nil, zero value otherwise.

### GetClonedWlanIdOk

`func (o *CreateWlanGroupOpenApiVO) GetClonedWlanIdOk() (*string, bool)`

GetClonedWlanIdOk returns a tuple with the ClonedWlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedWlanId

`func (o *CreateWlanGroupOpenApiVO) SetClonedWlanId(v string)`

SetClonedWlanId sets ClonedWlanId field to given value.

### HasClonedWlanId

`func (o *CreateWlanGroupOpenApiVO) HasClonedWlanId() bool`

HasClonedWlanId returns a boolean if a field has been set.

### GetName

`func (o *CreateWlanGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWlanGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWlanGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


