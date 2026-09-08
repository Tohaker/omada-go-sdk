# SiteLagHashAlgOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** | It should be a value as follows: 0: SRC MAC; 1: DST MAC; 2: SRC MAC + DST MAC; 3: SRC IP; 4: DST IP; 5: SRC IP + DST IP. | 

## Methods

### NewSiteLagHashAlgOpenApiVO

`func NewSiteLagHashAlgOpenApiVO(type_ int32, ) *SiteLagHashAlgOpenApiVO`

NewSiteLagHashAlgOpenApiVO instantiates a new SiteLagHashAlgOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteLagHashAlgOpenApiVOWithDefaults

`func NewSiteLagHashAlgOpenApiVOWithDefaults() *SiteLagHashAlgOpenApiVO`

NewSiteLagHashAlgOpenApiVOWithDefaults instantiates a new SiteLagHashAlgOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SiteLagHashAlgOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SiteLagHashAlgOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SiteLagHashAlgOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


