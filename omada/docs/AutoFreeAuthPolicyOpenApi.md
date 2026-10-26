# AutoFreeAuthPolicyOpenApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourcePortal** | Pointer to **string** | The name of the source portal, required when parameter [sourceType] is 1 | [optional] 
**SourceType** | Pointer to **int32** | The source type of the auto free auth policy, 0: System created; 1: From Portal Created | [optional] 
**Url** | Pointer to **string** | The URL of the automatically generated free auth policy | [optional] 

## Methods

### NewAutoFreeAuthPolicyOpenApi

`func NewAutoFreeAuthPolicyOpenApi() *AutoFreeAuthPolicyOpenApi`

NewAutoFreeAuthPolicyOpenApi instantiates a new AutoFreeAuthPolicyOpenApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoFreeAuthPolicyOpenApiWithDefaults

`func NewAutoFreeAuthPolicyOpenApiWithDefaults() *AutoFreeAuthPolicyOpenApi`

NewAutoFreeAuthPolicyOpenApiWithDefaults instantiates a new AutoFreeAuthPolicyOpenApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourcePortal

`func (o *AutoFreeAuthPolicyOpenApi) GetSourcePortal() string`

GetSourcePortal returns the SourcePortal field if non-nil, zero value otherwise.

### GetSourcePortalOk

`func (o *AutoFreeAuthPolicyOpenApi) GetSourcePortalOk() (*string, bool)`

GetSourcePortalOk returns a tuple with the SourcePortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcePortal

`func (o *AutoFreeAuthPolicyOpenApi) SetSourcePortal(v string)`

SetSourcePortal sets SourcePortal field to given value.

### HasSourcePortal

`func (o *AutoFreeAuthPolicyOpenApi) HasSourcePortal() bool`

HasSourcePortal returns a boolean if a field has been set.

### GetSourceType

`func (o *AutoFreeAuthPolicyOpenApi) GetSourceType() int32`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *AutoFreeAuthPolicyOpenApi) GetSourceTypeOk() (*int32, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *AutoFreeAuthPolicyOpenApi) SetSourceType(v int32)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *AutoFreeAuthPolicyOpenApi) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetUrl

`func (o *AutoFreeAuthPolicyOpenApi) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AutoFreeAuthPolicyOpenApi) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AutoFreeAuthPolicyOpenApi) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AutoFreeAuthPolicyOpenApi) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


