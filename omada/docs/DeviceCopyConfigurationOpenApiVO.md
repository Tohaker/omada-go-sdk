# DeviceCopyConfigurationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceMac** | **string** | Source device MAC, like AA-BB-CC-DD-EE-FF | 
**TargetMac** | **string** | Target device MAC, like AA-BB-CC-DD-EE-FF | 

## Methods

### NewDeviceCopyConfigurationOpenApiVO

`func NewDeviceCopyConfigurationOpenApiVO(sourceMac string, targetMac string, ) *DeviceCopyConfigurationOpenApiVO`

NewDeviceCopyConfigurationOpenApiVO instantiates a new DeviceCopyConfigurationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCopyConfigurationOpenApiVOWithDefaults

`func NewDeviceCopyConfigurationOpenApiVOWithDefaults() *DeviceCopyConfigurationOpenApiVO`

NewDeviceCopyConfigurationOpenApiVOWithDefaults instantiates a new DeviceCopyConfigurationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceMac

`func (o *DeviceCopyConfigurationOpenApiVO) GetSourceMac() string`

GetSourceMac returns the SourceMac field if non-nil, zero value otherwise.

### GetSourceMacOk

`func (o *DeviceCopyConfigurationOpenApiVO) GetSourceMacOk() (*string, bool)`

GetSourceMacOk returns a tuple with the SourceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMac

`func (o *DeviceCopyConfigurationOpenApiVO) SetSourceMac(v string)`

SetSourceMac sets SourceMac field to given value.


### GetTargetMac

`func (o *DeviceCopyConfigurationOpenApiVO) GetTargetMac() string`

GetTargetMac returns the TargetMac field if non-nil, zero value otherwise.

### GetTargetMacOk

`func (o *DeviceCopyConfigurationOpenApiVO) GetTargetMacOk() (*string, bool)`

GetTargetMacOk returns a tuple with the TargetMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMac

`func (o *DeviceCopyConfigurationOpenApiVO) SetTargetMac(v string)`

SetTargetMac sets TargetMac field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


