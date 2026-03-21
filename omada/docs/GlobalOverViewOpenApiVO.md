# GlobalOverViewOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudAccessStatus** | Pointer to **int32** | the status of cloud access. -1: disabled  0: disconneted  1: conneted | [optional] 
**ConnectedApNum** | Pointer to **int32** | total number of connected ap | [optional] 
**ConnectedGatewayNum** | Pointer to **int32** | total number of connected gateways | [optional] 
**ConnectedOltNum** | Pointer to **int32** | total number of connected olt | [optional] 
**ConnectedSwitchNum** | Pointer to **int32** | total number of connected switch | [optional] 
**CountryNum** | Pointer to **int32** | total number of country | [optional] 
**DisconnectedApNum** | Pointer to **int32** | total number of disconnected ap | [optional] 
**DisconnectedGatewayNum** | Pointer to **int32** | total number of disconnected gateways | [optional] 
**DisconnectedOltNum** | Pointer to **int32** | total number of disconnected olt | [optional] 
**DisconnectedSwitchNum** | Pointer to **int32** | total number of disconnected switch | [optional] 
**GuestNum** | Pointer to **int32** | total number of wireless guest | [optional] 
**IsolatedApNum** | Pointer to **int32** | total number of isolated ap | [optional] 
**PreConfigApNum** | Pointer to **int32** | the number of preconfig ap. | [optional] 
**PreConfigOltNum** | Pointer to **int32** | the number of preconfig olt. | [optional] 
**PreConfigOsgNum** | Pointer to **int32** | the number of preconfig gateway. | [optional] 
**PreConfigOswNum** | Pointer to **int32** | the number of preconfig switch. | [optional] 
**SiteNum** | Pointer to **int32** | total number of site | [optional] 
**TotalApNum** | Pointer to **int32** | total number of ap | [optional] 
**TotalClientNum** | Pointer to **int32** | total number of client | [optional] 
**TotalGatewayNum** | Pointer to **int32** | total number of gateways | [optional] 
**TotalOltNum** | Pointer to **int32** | total number of olt | [optional] 
**TotalSwitchNum** | Pointer to **int32** | total number of switch | [optional] 
**WiredClientNum** | Pointer to **int32** | total number of wired client | [optional] 
**WirelessClientNum** | Pointer to **int32** | total number of wireless client | [optional] 

## Methods

### NewGlobalOverViewOpenApiVO

`func NewGlobalOverViewOpenApiVO() *GlobalOverViewOpenApiVO`

NewGlobalOverViewOpenApiVO instantiates a new GlobalOverViewOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalOverViewOpenApiVOWithDefaults

`func NewGlobalOverViewOpenApiVOWithDefaults() *GlobalOverViewOpenApiVO`

NewGlobalOverViewOpenApiVOWithDefaults instantiates a new GlobalOverViewOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudAccessStatus

`func (o *GlobalOverViewOpenApiVO) GetCloudAccessStatus() int32`

GetCloudAccessStatus returns the CloudAccessStatus field if non-nil, zero value otherwise.

### GetCloudAccessStatusOk

`func (o *GlobalOverViewOpenApiVO) GetCloudAccessStatusOk() (*int32, bool)`

GetCloudAccessStatusOk returns a tuple with the CloudAccessStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAccessStatus

`func (o *GlobalOverViewOpenApiVO) SetCloudAccessStatus(v int32)`

SetCloudAccessStatus sets CloudAccessStatus field to given value.

### HasCloudAccessStatus

`func (o *GlobalOverViewOpenApiVO) HasCloudAccessStatus() bool`

HasCloudAccessStatus returns a boolean if a field has been set.

### GetConnectedApNum

`func (o *GlobalOverViewOpenApiVO) GetConnectedApNum() int32`

GetConnectedApNum returns the ConnectedApNum field if non-nil, zero value otherwise.

### GetConnectedApNumOk

`func (o *GlobalOverViewOpenApiVO) GetConnectedApNumOk() (*int32, bool)`

GetConnectedApNumOk returns a tuple with the ConnectedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedApNum

`func (o *GlobalOverViewOpenApiVO) SetConnectedApNum(v int32)`

SetConnectedApNum sets ConnectedApNum field to given value.

### HasConnectedApNum

`func (o *GlobalOverViewOpenApiVO) HasConnectedApNum() bool`

HasConnectedApNum returns a boolean if a field has been set.

### GetConnectedGatewayNum

`func (o *GlobalOverViewOpenApiVO) GetConnectedGatewayNum() int32`

GetConnectedGatewayNum returns the ConnectedGatewayNum field if non-nil, zero value otherwise.

### GetConnectedGatewayNumOk

`func (o *GlobalOverViewOpenApiVO) GetConnectedGatewayNumOk() (*int32, bool)`

GetConnectedGatewayNumOk returns a tuple with the ConnectedGatewayNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedGatewayNum

`func (o *GlobalOverViewOpenApiVO) SetConnectedGatewayNum(v int32)`

SetConnectedGatewayNum sets ConnectedGatewayNum field to given value.

### HasConnectedGatewayNum

`func (o *GlobalOverViewOpenApiVO) HasConnectedGatewayNum() bool`

HasConnectedGatewayNum returns a boolean if a field has been set.

### GetConnectedOltNum

`func (o *GlobalOverViewOpenApiVO) GetConnectedOltNum() int32`

GetConnectedOltNum returns the ConnectedOltNum field if non-nil, zero value otherwise.

### GetConnectedOltNumOk

`func (o *GlobalOverViewOpenApiVO) GetConnectedOltNumOk() (*int32, bool)`

GetConnectedOltNumOk returns a tuple with the ConnectedOltNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedOltNum

`func (o *GlobalOverViewOpenApiVO) SetConnectedOltNum(v int32)`

SetConnectedOltNum sets ConnectedOltNum field to given value.

### HasConnectedOltNum

`func (o *GlobalOverViewOpenApiVO) HasConnectedOltNum() bool`

HasConnectedOltNum returns a boolean if a field has been set.

### GetConnectedSwitchNum

`func (o *GlobalOverViewOpenApiVO) GetConnectedSwitchNum() int32`

GetConnectedSwitchNum returns the ConnectedSwitchNum field if non-nil, zero value otherwise.

### GetConnectedSwitchNumOk

`func (o *GlobalOverViewOpenApiVO) GetConnectedSwitchNumOk() (*int32, bool)`

GetConnectedSwitchNumOk returns a tuple with the ConnectedSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSwitchNum

`func (o *GlobalOverViewOpenApiVO) SetConnectedSwitchNum(v int32)`

SetConnectedSwitchNum sets ConnectedSwitchNum field to given value.

### HasConnectedSwitchNum

`func (o *GlobalOverViewOpenApiVO) HasConnectedSwitchNum() bool`

HasConnectedSwitchNum returns a boolean if a field has been set.

### GetCountryNum

`func (o *GlobalOverViewOpenApiVO) GetCountryNum() int32`

GetCountryNum returns the CountryNum field if non-nil, zero value otherwise.

### GetCountryNumOk

`func (o *GlobalOverViewOpenApiVO) GetCountryNumOk() (*int32, bool)`

GetCountryNumOk returns a tuple with the CountryNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryNum

`func (o *GlobalOverViewOpenApiVO) SetCountryNum(v int32)`

SetCountryNum sets CountryNum field to given value.

### HasCountryNum

`func (o *GlobalOverViewOpenApiVO) HasCountryNum() bool`

HasCountryNum returns a boolean if a field has been set.

### GetDisconnectedApNum

`func (o *GlobalOverViewOpenApiVO) GetDisconnectedApNum() int32`

GetDisconnectedApNum returns the DisconnectedApNum field if non-nil, zero value otherwise.

### GetDisconnectedApNumOk

`func (o *GlobalOverViewOpenApiVO) GetDisconnectedApNumOk() (*int32, bool)`

GetDisconnectedApNumOk returns a tuple with the DisconnectedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedApNum

`func (o *GlobalOverViewOpenApiVO) SetDisconnectedApNum(v int32)`

SetDisconnectedApNum sets DisconnectedApNum field to given value.

### HasDisconnectedApNum

`func (o *GlobalOverViewOpenApiVO) HasDisconnectedApNum() bool`

HasDisconnectedApNum returns a boolean if a field has been set.

### GetDisconnectedGatewayNum

`func (o *GlobalOverViewOpenApiVO) GetDisconnectedGatewayNum() int32`

GetDisconnectedGatewayNum returns the DisconnectedGatewayNum field if non-nil, zero value otherwise.

### GetDisconnectedGatewayNumOk

`func (o *GlobalOverViewOpenApiVO) GetDisconnectedGatewayNumOk() (*int32, bool)`

GetDisconnectedGatewayNumOk returns a tuple with the DisconnectedGatewayNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedGatewayNum

`func (o *GlobalOverViewOpenApiVO) SetDisconnectedGatewayNum(v int32)`

SetDisconnectedGatewayNum sets DisconnectedGatewayNum field to given value.

### HasDisconnectedGatewayNum

`func (o *GlobalOverViewOpenApiVO) HasDisconnectedGatewayNum() bool`

HasDisconnectedGatewayNum returns a boolean if a field has been set.

### GetDisconnectedOltNum

`func (o *GlobalOverViewOpenApiVO) GetDisconnectedOltNum() int32`

GetDisconnectedOltNum returns the DisconnectedOltNum field if non-nil, zero value otherwise.

### GetDisconnectedOltNumOk

`func (o *GlobalOverViewOpenApiVO) GetDisconnectedOltNumOk() (*int32, bool)`

GetDisconnectedOltNumOk returns a tuple with the DisconnectedOltNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedOltNum

`func (o *GlobalOverViewOpenApiVO) SetDisconnectedOltNum(v int32)`

SetDisconnectedOltNum sets DisconnectedOltNum field to given value.

### HasDisconnectedOltNum

`func (o *GlobalOverViewOpenApiVO) HasDisconnectedOltNum() bool`

HasDisconnectedOltNum returns a boolean if a field has been set.

### GetDisconnectedSwitchNum

`func (o *GlobalOverViewOpenApiVO) GetDisconnectedSwitchNum() int32`

GetDisconnectedSwitchNum returns the DisconnectedSwitchNum field if non-nil, zero value otherwise.

### GetDisconnectedSwitchNumOk

`func (o *GlobalOverViewOpenApiVO) GetDisconnectedSwitchNumOk() (*int32, bool)`

GetDisconnectedSwitchNumOk returns a tuple with the DisconnectedSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSwitchNum

`func (o *GlobalOverViewOpenApiVO) SetDisconnectedSwitchNum(v int32)`

SetDisconnectedSwitchNum sets DisconnectedSwitchNum field to given value.

### HasDisconnectedSwitchNum

`func (o *GlobalOverViewOpenApiVO) HasDisconnectedSwitchNum() bool`

HasDisconnectedSwitchNum returns a boolean if a field has been set.

### GetGuestNum

`func (o *GlobalOverViewOpenApiVO) GetGuestNum() int32`

GetGuestNum returns the GuestNum field if non-nil, zero value otherwise.

### GetGuestNumOk

`func (o *GlobalOverViewOpenApiVO) GetGuestNumOk() (*int32, bool)`

GetGuestNumOk returns a tuple with the GuestNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestNum

`func (o *GlobalOverViewOpenApiVO) SetGuestNum(v int32)`

SetGuestNum sets GuestNum field to given value.

### HasGuestNum

`func (o *GlobalOverViewOpenApiVO) HasGuestNum() bool`

HasGuestNum returns a boolean if a field has been set.

### GetIsolatedApNum

`func (o *GlobalOverViewOpenApiVO) GetIsolatedApNum() int32`

GetIsolatedApNum returns the IsolatedApNum field if non-nil, zero value otherwise.

### GetIsolatedApNumOk

`func (o *GlobalOverViewOpenApiVO) GetIsolatedApNumOk() (*int32, bool)`

GetIsolatedApNumOk returns a tuple with the IsolatedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedApNum

`func (o *GlobalOverViewOpenApiVO) SetIsolatedApNum(v int32)`

SetIsolatedApNum sets IsolatedApNum field to given value.

### HasIsolatedApNum

`func (o *GlobalOverViewOpenApiVO) HasIsolatedApNum() bool`

HasIsolatedApNum returns a boolean if a field has been set.

### GetPreConfigApNum

`func (o *GlobalOverViewOpenApiVO) GetPreConfigApNum() int32`

GetPreConfigApNum returns the PreConfigApNum field if non-nil, zero value otherwise.

### GetPreConfigApNumOk

`func (o *GlobalOverViewOpenApiVO) GetPreConfigApNumOk() (*int32, bool)`

GetPreConfigApNumOk returns a tuple with the PreConfigApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigApNum

`func (o *GlobalOverViewOpenApiVO) SetPreConfigApNum(v int32)`

SetPreConfigApNum sets PreConfigApNum field to given value.

### HasPreConfigApNum

`func (o *GlobalOverViewOpenApiVO) HasPreConfigApNum() bool`

HasPreConfigApNum returns a boolean if a field has been set.

### GetPreConfigOltNum

`func (o *GlobalOverViewOpenApiVO) GetPreConfigOltNum() int32`

GetPreConfigOltNum returns the PreConfigOltNum field if non-nil, zero value otherwise.

### GetPreConfigOltNumOk

`func (o *GlobalOverViewOpenApiVO) GetPreConfigOltNumOk() (*int32, bool)`

GetPreConfigOltNumOk returns a tuple with the PreConfigOltNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigOltNum

`func (o *GlobalOverViewOpenApiVO) SetPreConfigOltNum(v int32)`

SetPreConfigOltNum sets PreConfigOltNum field to given value.

### HasPreConfigOltNum

`func (o *GlobalOverViewOpenApiVO) HasPreConfigOltNum() bool`

HasPreConfigOltNum returns a boolean if a field has been set.

### GetPreConfigOsgNum

`func (o *GlobalOverViewOpenApiVO) GetPreConfigOsgNum() int32`

GetPreConfigOsgNum returns the PreConfigOsgNum field if non-nil, zero value otherwise.

### GetPreConfigOsgNumOk

`func (o *GlobalOverViewOpenApiVO) GetPreConfigOsgNumOk() (*int32, bool)`

GetPreConfigOsgNumOk returns a tuple with the PreConfigOsgNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigOsgNum

`func (o *GlobalOverViewOpenApiVO) SetPreConfigOsgNum(v int32)`

SetPreConfigOsgNum sets PreConfigOsgNum field to given value.

### HasPreConfigOsgNum

`func (o *GlobalOverViewOpenApiVO) HasPreConfigOsgNum() bool`

HasPreConfigOsgNum returns a boolean if a field has been set.

### GetPreConfigOswNum

`func (o *GlobalOverViewOpenApiVO) GetPreConfigOswNum() int32`

GetPreConfigOswNum returns the PreConfigOswNum field if non-nil, zero value otherwise.

### GetPreConfigOswNumOk

`func (o *GlobalOverViewOpenApiVO) GetPreConfigOswNumOk() (*int32, bool)`

GetPreConfigOswNumOk returns a tuple with the PreConfigOswNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreConfigOswNum

`func (o *GlobalOverViewOpenApiVO) SetPreConfigOswNum(v int32)`

SetPreConfigOswNum sets PreConfigOswNum field to given value.

### HasPreConfigOswNum

`func (o *GlobalOverViewOpenApiVO) HasPreConfigOswNum() bool`

HasPreConfigOswNum returns a boolean if a field has been set.

### GetSiteNum

`func (o *GlobalOverViewOpenApiVO) GetSiteNum() int32`

GetSiteNum returns the SiteNum field if non-nil, zero value otherwise.

### GetSiteNumOk

`func (o *GlobalOverViewOpenApiVO) GetSiteNumOk() (*int32, bool)`

GetSiteNumOk returns a tuple with the SiteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNum

`func (o *GlobalOverViewOpenApiVO) SetSiteNum(v int32)`

SetSiteNum sets SiteNum field to given value.

### HasSiteNum

`func (o *GlobalOverViewOpenApiVO) HasSiteNum() bool`

HasSiteNum returns a boolean if a field has been set.

### GetTotalApNum

`func (o *GlobalOverViewOpenApiVO) GetTotalApNum() int32`

GetTotalApNum returns the TotalApNum field if non-nil, zero value otherwise.

### GetTotalApNumOk

`func (o *GlobalOverViewOpenApiVO) GetTotalApNumOk() (*int32, bool)`

GetTotalApNumOk returns a tuple with the TotalApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalApNum

`func (o *GlobalOverViewOpenApiVO) SetTotalApNum(v int32)`

SetTotalApNum sets TotalApNum field to given value.

### HasTotalApNum

`func (o *GlobalOverViewOpenApiVO) HasTotalApNum() bool`

HasTotalApNum returns a boolean if a field has been set.

### GetTotalClientNum

`func (o *GlobalOverViewOpenApiVO) GetTotalClientNum() int32`

GetTotalClientNum returns the TotalClientNum field if non-nil, zero value otherwise.

### GetTotalClientNumOk

`func (o *GlobalOverViewOpenApiVO) GetTotalClientNumOk() (*int32, bool)`

GetTotalClientNumOk returns a tuple with the TotalClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalClientNum

`func (o *GlobalOverViewOpenApiVO) SetTotalClientNum(v int32)`

SetTotalClientNum sets TotalClientNum field to given value.

### HasTotalClientNum

`func (o *GlobalOverViewOpenApiVO) HasTotalClientNum() bool`

HasTotalClientNum returns a boolean if a field has been set.

### GetTotalGatewayNum

`func (o *GlobalOverViewOpenApiVO) GetTotalGatewayNum() int32`

GetTotalGatewayNum returns the TotalGatewayNum field if non-nil, zero value otherwise.

### GetTotalGatewayNumOk

`func (o *GlobalOverViewOpenApiVO) GetTotalGatewayNumOk() (*int32, bool)`

GetTotalGatewayNumOk returns a tuple with the TotalGatewayNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGatewayNum

`func (o *GlobalOverViewOpenApiVO) SetTotalGatewayNum(v int32)`

SetTotalGatewayNum sets TotalGatewayNum field to given value.

### HasTotalGatewayNum

`func (o *GlobalOverViewOpenApiVO) HasTotalGatewayNum() bool`

HasTotalGatewayNum returns a boolean if a field has been set.

### GetTotalOltNum

`func (o *GlobalOverViewOpenApiVO) GetTotalOltNum() int32`

GetTotalOltNum returns the TotalOltNum field if non-nil, zero value otherwise.

### GetTotalOltNumOk

`func (o *GlobalOverViewOpenApiVO) GetTotalOltNumOk() (*int32, bool)`

GetTotalOltNumOk returns a tuple with the TotalOltNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOltNum

`func (o *GlobalOverViewOpenApiVO) SetTotalOltNum(v int32)`

SetTotalOltNum sets TotalOltNum field to given value.

### HasTotalOltNum

`func (o *GlobalOverViewOpenApiVO) HasTotalOltNum() bool`

HasTotalOltNum returns a boolean if a field has been set.

### GetTotalSwitchNum

`func (o *GlobalOverViewOpenApiVO) GetTotalSwitchNum() int32`

GetTotalSwitchNum returns the TotalSwitchNum field if non-nil, zero value otherwise.

### GetTotalSwitchNumOk

`func (o *GlobalOverViewOpenApiVO) GetTotalSwitchNumOk() (*int32, bool)`

GetTotalSwitchNumOk returns a tuple with the TotalSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSwitchNum

`func (o *GlobalOverViewOpenApiVO) SetTotalSwitchNum(v int32)`

SetTotalSwitchNum sets TotalSwitchNum field to given value.

### HasTotalSwitchNum

`func (o *GlobalOverViewOpenApiVO) HasTotalSwitchNum() bool`

HasTotalSwitchNum returns a boolean if a field has been set.

### GetWiredClientNum

`func (o *GlobalOverViewOpenApiVO) GetWiredClientNum() int32`

GetWiredClientNum returns the WiredClientNum field if non-nil, zero value otherwise.

### GetWiredClientNumOk

`func (o *GlobalOverViewOpenApiVO) GetWiredClientNumOk() (*int32, bool)`

GetWiredClientNumOk returns a tuple with the WiredClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClientNum

`func (o *GlobalOverViewOpenApiVO) SetWiredClientNum(v int32)`

SetWiredClientNum sets WiredClientNum field to given value.

### HasWiredClientNum

`func (o *GlobalOverViewOpenApiVO) HasWiredClientNum() bool`

HasWiredClientNum returns a boolean if a field has been set.

### GetWirelessClientNum

`func (o *GlobalOverViewOpenApiVO) GetWirelessClientNum() int32`

GetWirelessClientNum returns the WirelessClientNum field if non-nil, zero value otherwise.

### GetWirelessClientNumOk

`func (o *GlobalOverViewOpenApiVO) GetWirelessClientNumOk() (*int32, bool)`

GetWirelessClientNumOk returns a tuple with the WirelessClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessClientNum

`func (o *GlobalOverViewOpenApiVO) SetWirelessClientNum(v int32)`

SetWirelessClientNum sets WirelessClientNum field to given value.

### HasWirelessClientNum

`func (o *GlobalOverViewOpenApiVO) HasWirelessClientNum() bool`

HasWirelessClientNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


