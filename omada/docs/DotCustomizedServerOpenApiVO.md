# DotCustomizedServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Custom service enable status | 
**Name** | **string** | Custom service name, should contain 1 to 64 characters | 
**Servers** | **[]string** | Custom service IPs, up to 2 IPs can be configured | 

## Methods

### NewDotCustomizedServerOpenApiVO

`func NewDotCustomizedServerOpenApiVO(enable bool, name string, servers []string, ) *DotCustomizedServerOpenApiVO`

NewDotCustomizedServerOpenApiVO instantiates a new DotCustomizedServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDotCustomizedServerOpenApiVOWithDefaults

`func NewDotCustomizedServerOpenApiVOWithDefaults() *DotCustomizedServerOpenApiVO`

NewDotCustomizedServerOpenApiVOWithDefaults instantiates a new DotCustomizedServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *DotCustomizedServerOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DotCustomizedServerOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DotCustomizedServerOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetName

`func (o *DotCustomizedServerOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DotCustomizedServerOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DotCustomizedServerOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetServers

`func (o *DotCustomizedServerOpenApiVO) GetServers() []string`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *DotCustomizedServerOpenApiVO) GetServersOk() (*[]string, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *DotCustomizedServerOpenApiVO) SetServers(v []string)`

SetServers sets Servers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


