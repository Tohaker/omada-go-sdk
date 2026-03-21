# DohCustomizedServerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Custom service enable status | 
**Name** | **string** | Custom service name, should contain 1 to 64 characters | 
**Server** | **string** | Custom service IP or domain | 

## Methods

### NewDohCustomizedServerOpenApiVO

`func NewDohCustomizedServerOpenApiVO(enable bool, name string, server string, ) *DohCustomizedServerOpenApiVO`

NewDohCustomizedServerOpenApiVO instantiates a new DohCustomizedServerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDohCustomizedServerOpenApiVOWithDefaults

`func NewDohCustomizedServerOpenApiVOWithDefaults() *DohCustomizedServerOpenApiVO`

NewDohCustomizedServerOpenApiVOWithDefaults instantiates a new DohCustomizedServerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *DohCustomizedServerOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DohCustomizedServerOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DohCustomizedServerOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetName

`func (o *DohCustomizedServerOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DohCustomizedServerOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DohCustomizedServerOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetServer

`func (o *DohCustomizedServerOpenApiVO) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *DohCustomizedServerOpenApiVO) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *DohCustomizedServerOpenApiVO) SetServer(v string)`

SetServer sets Server field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


