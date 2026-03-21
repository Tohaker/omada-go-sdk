# GlobalSearchResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Devices** | Pointer to [**[]DeviceGlobalSearchSummary**](DeviceGlobalSearchSummary.md) | device list result | [optional] 
**SiteNames** | Pointer to **map[string]string** | site name | [optional] 

## Methods

### NewGlobalSearchResultOpenApiVO

`func NewGlobalSearchResultOpenApiVO() *GlobalSearchResultOpenApiVO`

NewGlobalSearchResultOpenApiVO instantiates a new GlobalSearchResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSearchResultOpenApiVOWithDefaults

`func NewGlobalSearchResultOpenApiVOWithDefaults() *GlobalSearchResultOpenApiVO`

NewGlobalSearchResultOpenApiVOWithDefaults instantiates a new GlobalSearchResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevices

`func (o *GlobalSearchResultOpenApiVO) GetDevices() []DeviceGlobalSearchSummary`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *GlobalSearchResultOpenApiVO) GetDevicesOk() (*[]DeviceGlobalSearchSummary, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *GlobalSearchResultOpenApiVO) SetDevices(v []DeviceGlobalSearchSummary)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *GlobalSearchResultOpenApiVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetSiteNames

`func (o *GlobalSearchResultOpenApiVO) GetSiteNames() map[string]string`

GetSiteNames returns the SiteNames field if non-nil, zero value otherwise.

### GetSiteNamesOk

`func (o *GlobalSearchResultOpenApiVO) GetSiteNamesOk() (*map[string]string, bool)`

GetSiteNamesOk returns a tuple with the SiteNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNames

`func (o *GlobalSearchResultOpenApiVO) SetSiteNames(v map[string]string)`

SetSiteNames sets SiteNames field to given value.

### HasSiteNames

`func (o *GlobalSearchResultOpenApiVO) HasSiteNames() bool`

HasSiteNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


