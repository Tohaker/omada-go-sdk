# RequestActionEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | OpenAPI request body, same as regular OpenAPI. Should be a JSON object but not a string. | [optional] 
**Method** | **string** | OpenAPI request method. Same as regular OpenAPI, it should be a value as follows: POST, PATCH, PUT, DELETE. | 
**Path** | **string** | OpenAPI request path. Same as regular OpenAPI, file upload and download are not supported. | 
**Query** | Pointer to **string** | OpenAPI request query of the path, same as regular OpenAPI. | [optional] 

## Methods

### NewRequestActionEntity

`func NewRequestActionEntity(method string, path string, ) *RequestActionEntity`

NewRequestActionEntity instantiates a new RequestActionEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestActionEntityWithDefaults

`func NewRequestActionEntityWithDefaults() *RequestActionEntity`

NewRequestActionEntityWithDefaults instantiates a new RequestActionEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *RequestActionEntity) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RequestActionEntity) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RequestActionEntity) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *RequestActionEntity) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMethod

`func (o *RequestActionEntity) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *RequestActionEntity) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *RequestActionEntity) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPath

`func (o *RequestActionEntity) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RequestActionEntity) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RequestActionEntity) SetPath(v string)`

SetPath sets Path field to given value.


### GetQuery

`func (o *RequestActionEntity) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RequestActionEntity) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RequestActionEntity) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *RequestActionEntity) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


