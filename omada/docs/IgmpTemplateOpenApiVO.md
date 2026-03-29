# IgmpTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** |  | 
**Version** | **int32** | Version should be a value as follows: 2:v2; 3:v3. | 
**WanPortId** | **string** | wan port ID, can be obtained from &#39;Get site template internet basic info&#39; interface. | 

## Methods

### NewIgmpTemplateOpenApiVO

`func NewIgmpTemplateOpenApiVO(enable bool, version int32, wanPortId string, ) *IgmpTemplateOpenApiVO`

NewIgmpTemplateOpenApiVO instantiates a new IgmpTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIgmpTemplateOpenApiVOWithDefaults

`func NewIgmpTemplateOpenApiVOWithDefaults() *IgmpTemplateOpenApiVO`

NewIgmpTemplateOpenApiVOWithDefaults instantiates a new IgmpTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *IgmpTemplateOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IgmpTemplateOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IgmpTemplateOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetVersion

`func (o *IgmpTemplateOpenApiVO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IgmpTemplateOpenApiVO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IgmpTemplateOpenApiVO) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWanPortId

`func (o *IgmpTemplateOpenApiVO) GetWanPortId() string`

GetWanPortId returns the WanPortId field if non-nil, zero value otherwise.

### GetWanPortIdOk

`func (o *IgmpTemplateOpenApiVO) GetWanPortIdOk() (*string, bool)`

GetWanPortIdOk returns a tuple with the WanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortId

`func (o *IgmpTemplateOpenApiVO) SetWanPortId(v string)`

SetWanPortId sets WanPortId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


