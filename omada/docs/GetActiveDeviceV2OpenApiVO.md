# GetActiveDeviceV2OpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveAps** | Pointer to [**[]ActiveDevice**](ActiveDevice.md) | Most Active APs devices by traffic | [optional] 
**TotalTraffic** | Pointer to **float64** | The sum of the traffic of all AP devices under the site | [optional] 

## Methods

### NewGetActiveDeviceV2OpenApiVO

`func NewGetActiveDeviceV2OpenApiVO() *GetActiveDeviceV2OpenApiVO`

NewGetActiveDeviceV2OpenApiVO instantiates a new GetActiveDeviceV2OpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActiveDeviceV2OpenApiVOWithDefaults

`func NewGetActiveDeviceV2OpenApiVOWithDefaults() *GetActiveDeviceV2OpenApiVO`

NewGetActiveDeviceV2OpenApiVOWithDefaults instantiates a new GetActiveDeviceV2OpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveAps

`func (o *GetActiveDeviceV2OpenApiVO) GetActiveAps() []ActiveDevice`

GetActiveAps returns the ActiveAps field if non-nil, zero value otherwise.

### GetActiveApsOk

`func (o *GetActiveDeviceV2OpenApiVO) GetActiveApsOk() (*[]ActiveDevice, bool)`

GetActiveApsOk returns a tuple with the ActiveAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveAps

`func (o *GetActiveDeviceV2OpenApiVO) SetActiveAps(v []ActiveDevice)`

SetActiveAps sets ActiveAps field to given value.

### HasActiveAps

`func (o *GetActiveDeviceV2OpenApiVO) HasActiveAps() bool`

HasActiveAps returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *GetActiveDeviceV2OpenApiVO) GetTotalTraffic() float64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *GetActiveDeviceV2OpenApiVO) GetTotalTrafficOk() (*float64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *GetActiveDeviceV2OpenApiVO) SetTotalTraffic(v float64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *GetActiveDeviceV2OpenApiVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


