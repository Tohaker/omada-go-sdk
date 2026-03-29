# ExportCliVarOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | [**[]DeviceExportCliVO**](DeviceExportCliVO.md) | Devices | 
**Type** | **int32** | Export type should be a value as follows: 0:CSV; 1:JSON | 

## Methods

### NewExportCliVarOpenApiVO

`func NewExportCliVarOpenApiVO(devices []DeviceExportCliVO, type_ int32, ) *ExportCliVarOpenApiVO`

NewExportCliVarOpenApiVO instantiates a new ExportCliVarOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportCliVarOpenApiVOWithDefaults

`func NewExportCliVarOpenApiVOWithDefaults() *ExportCliVarOpenApiVO`

NewExportCliVarOpenApiVOWithDefaults instantiates a new ExportCliVarOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *ExportCliVarOpenApiVO) GetDevices() []DeviceExportCliVO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *ExportCliVarOpenApiVO) GetDevicesOk() (*[]DeviceExportCliVO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *ExportCliVarOpenApiVO) SetDevices(v []DeviceExportCliVO)`

SetDevices sets Devices field to given value.


### GetType

`func (o *ExportCliVarOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportCliVarOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportCliVarOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


