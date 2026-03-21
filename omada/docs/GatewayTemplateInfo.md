# GatewayTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of gateway template | [optional] 
**PortConfigs** | Pointer to [**[]GatewayPortConfig**](GatewayPortConfig.md) | Gateway port configs | [optional] 
**ShowModel** | Pointer to **string** | Gateway model description | [optional] 

## Methods

### NewGatewayTemplateInfo

`func NewGatewayTemplateInfo() *GatewayTemplateInfo`

NewGatewayTemplateInfo instantiates a new GatewayTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTemplateInfoWithDefaults

`func NewGatewayTemplateInfoWithDefaults() *GatewayTemplateInfo`

NewGatewayTemplateInfoWithDefaults instantiates a new GatewayTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayTemplateInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayTemplateInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayTemplateInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayTemplateInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortConfigs

`func (o *GatewayTemplateInfo) GetPortConfigs() []GatewayPortConfig`

GetPortConfigs returns the PortConfigs field if non-nil, zero value otherwise.

### GetPortConfigsOk

`func (o *GatewayTemplateInfo) GetPortConfigsOk() (*[]GatewayPortConfig, bool)`

GetPortConfigsOk returns a tuple with the PortConfigs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortConfigs

`func (o *GatewayTemplateInfo) SetPortConfigs(v []GatewayPortConfig)`

SetPortConfigs sets PortConfigs field to given value.

### HasPortConfigs

`func (o *GatewayTemplateInfo) HasPortConfigs() bool`

HasPortConfigs returns a boolean if a field has been set.

### GetShowModel

`func (o *GatewayTemplateInfo) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *GatewayTemplateInfo) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *GatewayTemplateInfo) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *GatewayTemplateInfo) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


