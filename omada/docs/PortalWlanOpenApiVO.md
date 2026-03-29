# PortalWlanOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SsidList** | Pointer to [**[]PortalSsidOpenApiVO**](PortalSsidOpenApiVO.md) | SSID list of WLAN | [optional] 
**WlanId** | Pointer to **string** | WLAN ID | [optional] 
**WlanName** | Pointer to **string** | WLAN name | [optional] 

## Methods

### NewPortalWlanOpenApiVO

`func NewPortalWlanOpenApiVO() *PortalWlanOpenApiVO`

NewPortalWlanOpenApiVO instantiates a new PortalWlanOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalWlanOpenApiVOWithDefaults

`func NewPortalWlanOpenApiVOWithDefaults() *PortalWlanOpenApiVO`

NewPortalWlanOpenApiVOWithDefaults instantiates a new PortalWlanOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSsidList

`func (o *PortalWlanOpenApiVO) GetSsidList() []PortalSsidOpenApiVO`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *PortalWlanOpenApiVO) GetSsidListOk() (*[]PortalSsidOpenApiVO, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *PortalWlanOpenApiVO) SetSsidList(v []PortalSsidOpenApiVO)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *PortalWlanOpenApiVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.

### GetWlanId

`func (o *PortalWlanOpenApiVO) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *PortalWlanOpenApiVO) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *PortalWlanOpenApiVO) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *PortalWlanOpenApiVO) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.

### GetWlanName

`func (o *PortalWlanOpenApiVO) GetWlanName() string`

GetWlanName returns the WlanName field if non-nil, zero value otherwise.

### GetWlanNameOk

`func (o *PortalWlanOpenApiVO) GetWlanNameOk() (*string, bool)`

GetWlanNameOk returns a tuple with the WlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanName

`func (o *PortalWlanOpenApiVO) SetWlanName(v string)`

SetWlanName sets WlanName field to given value.

### HasWlanName

`func (o *PortalWlanOpenApiVO) HasWlanName() bool`

HasWlanName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


