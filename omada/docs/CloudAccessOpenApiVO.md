# CloudAccessOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | enable | [optional] 
**Owner** | Pointer to **string** | owner | [optional] 
**Status** | Pointer to **int32** | status, DISABLED &#x3D; -1; CONNECTED &#x3D; 1; DISCONNECTED &#x3D; 0 | [optional] 

## Methods

### NewCloudAccessOpenApiVO

`func NewCloudAccessOpenApiVO() *CloudAccessOpenApiVO`

NewCloudAccessOpenApiVO instantiates a new CloudAccessOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccessOpenApiVOWithDefaults

`func NewCloudAccessOpenApiVOWithDefaults() *CloudAccessOpenApiVO`

NewCloudAccessOpenApiVOWithDefaults instantiates a new CloudAccessOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *CloudAccessOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CloudAccessOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CloudAccessOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CloudAccessOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetOwner

`func (o *CloudAccessOpenApiVO) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudAccessOpenApiVO) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudAccessOpenApiVO) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudAccessOpenApiVO) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAccessOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAccessOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAccessOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAccessOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


