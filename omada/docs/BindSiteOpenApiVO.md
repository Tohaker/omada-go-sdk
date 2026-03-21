# BindSiteOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SiteId** | **string** | Site ID to bind. | 
**Switches** | Pointer to [**[]DeviceBindOpenApiVO**](DeviceBindOpenApiVO.md) | Switch choose for binding device template. | [optional] 

## Methods

### NewBindSiteOpenApiVO

`func NewBindSiteOpenApiVO(siteId string, ) *BindSiteOpenApiVO`

NewBindSiteOpenApiVO instantiates a new BindSiteOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindSiteOpenApiVOWithDefaults

`func NewBindSiteOpenApiVOWithDefaults() *BindSiteOpenApiVO`

NewBindSiteOpenApiVOWithDefaults instantiates a new BindSiteOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSiteId

`func (o *BindSiteOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *BindSiteOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *BindSiteOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.


### GetSwitches

`func (o *BindSiteOpenApiVO) GetSwitches() []DeviceBindOpenApiVO`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *BindSiteOpenApiVO) GetSwitchesOk() (*[]DeviceBindOpenApiVO, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *BindSiteOpenApiVO) SetSwitches(v []DeviceBindOpenApiVO)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *BindSiteOpenApiVO) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


