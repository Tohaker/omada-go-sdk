# GatewayACLConfigModeEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | **int32** | Gateway acl config mode should be a value as follows: 0: Through profiles; 1: Custom | 

## Methods

### NewGatewayACLConfigModeEntity

`func NewGatewayACLConfigModeEntity(mode int32, ) *GatewayACLConfigModeEntity`

NewGatewayACLConfigModeEntity instantiates a new GatewayACLConfigModeEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayACLConfigModeEntityWithDefaults

`func NewGatewayACLConfigModeEntityWithDefaults() *GatewayACLConfigModeEntity`

NewGatewayACLConfigModeEntityWithDefaults instantiates a new GatewayACLConfigModeEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *GatewayACLConfigModeEntity) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GatewayACLConfigModeEntity) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GatewayACLConfigModeEntity) SetMode(v int32)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


