# ApP2pInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildAps** | Pointer to [**[]P2pInfoOpenApiVO**](P2pInfoOpenApiVO.md) | Child aps info | [optional] 
**MainAp** | Pointer to [**P2pInfoOpenApiVO**](P2pInfoOpenApiVO.md) |  | [optional] 

## Methods

### NewApP2pInfo

`func NewApP2pInfo() *ApP2pInfo`

NewApP2pInfo instantiates a new ApP2pInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApP2pInfoWithDefaults

`func NewApP2pInfoWithDefaults() *ApP2pInfo`

NewApP2pInfoWithDefaults instantiates a new ApP2pInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildAps

`func (o *ApP2pInfo) GetChildAps() []P2pInfoOpenApiVO`

GetChildAps returns the ChildAps field if non-nil, zero value otherwise.

### GetChildApsOk

`func (o *ApP2pInfo) GetChildApsOk() (*[]P2pInfoOpenApiVO, bool)`

GetChildApsOk returns a tuple with the ChildAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildAps

`func (o *ApP2pInfo) SetChildAps(v []P2pInfoOpenApiVO)`

SetChildAps sets ChildAps field to given value.

### HasChildAps

`func (o *ApP2pInfo) HasChildAps() bool`

HasChildAps returns a boolean if a field has been set.

### GetMainAp

`func (o *ApP2pInfo) GetMainAp() P2pInfoOpenApiVO`

GetMainAp returns the MainAp field if non-nil, zero value otherwise.

### GetMainApOk

`func (o *ApP2pInfo) GetMainApOk() (*P2pInfoOpenApiVO, bool)`

GetMainApOk returns a tuple with the MainAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainAp

`func (o *ApP2pInfo) SetMainAp(v P2pInfoOpenApiVO)`

SetMainAp sets MainAp field to given value.

### HasMainAp

`func (o *ApP2pInfo) HasMainAp() bool`

HasMainAp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


