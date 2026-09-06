# MspDeviceIncidentCountRequestOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceList** | [**[]MspDeviceIncidentCountItemOpenApiVO**](MspDeviceIncidentCountItemOpenApiVO.md) | Devices grouped by their owning customer. | 

## Methods

### NewMspDeviceIncidentCountRequestOpenApiVO

`func NewMspDeviceIncidentCountRequestOpenApiVO(deviceList []MspDeviceIncidentCountItemOpenApiVO, ) *MspDeviceIncidentCountRequestOpenApiVO`

NewMspDeviceIncidentCountRequestOpenApiVO instantiates a new MspDeviceIncidentCountRequestOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspDeviceIncidentCountRequestOpenApiVOWithDefaults

`func NewMspDeviceIncidentCountRequestOpenApiVOWithDefaults() *MspDeviceIncidentCountRequestOpenApiVO`

NewMspDeviceIncidentCountRequestOpenApiVOWithDefaults instantiates a new MspDeviceIncidentCountRequestOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceList

`func (o *MspDeviceIncidentCountRequestOpenApiVO) GetDeviceList() []MspDeviceIncidentCountItemOpenApiVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *MspDeviceIncidentCountRequestOpenApiVO) GetDeviceListOk() (*[]MspDeviceIncidentCountItemOpenApiVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *MspDeviceIncidentCountRequestOpenApiVO) SetDeviceList(v []MspDeviceIncidentCountItemOpenApiVO)`

SetDeviceList sets DeviceList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


