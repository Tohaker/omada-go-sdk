# MspSiteSummaryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Site address | [optional] 
**ConnectedApNum** | Pointer to **int64** | Connected ap num of current site | [optional] 
**ConnectedSwitchNum** | Pointer to **int64** | Connected switch num of current site | [optional] 
**CustomerId** | Pointer to **string** | Customer ID | [optional] 
**CustomerName** | Pointer to **string** | Customer name | [optional] 
**DisconnectedApNum** | Pointer to **int64** | Disconnected ap num of current site | [optional] 
**DisconnectedSwitchNum** | Pointer to **int64** | Disconnected switch num of current site | [optional] 
**IsolatedApNum** | Pointer to **int64** | Isolated ap num of current site | [optional] 
**Latitude** | Pointer to **float64** | Site latitude should be within the range of -90~90 | [optional] 
**Longitude** | Pointer to **float64** | Site longitude should be within the range of -180~180 | [optional] 
**Region** | Pointer to **string** | Country/Region of the site; For the values of region, refer to the abbreviation of the ISO country code; For example, you need to input \&quot;United States\&quot; for the United States of America. | [optional] 
**Scenario** | Pointer to **string** | Site scenario | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**SiteName** | Pointer to **string** | Site name | [optional] 
**SitePublicIp** | Pointer to **string** | Adopted gateway public ip of the site, only useful for cloud based controller and remote management local Controller | [optional] 
**SupportES** | Pointer to **bool** | Whether the site supports adopting Agile Series Switches | [optional] 
**SupportL2** | Pointer to **bool** | Whether the site supports adopting Non-Agile Series Switches | [optional] 
**TagIds** | Pointer to **[]string** | Site tag ID | [optional] 
**TimeZone** | Pointer to **string** | For the values of the timezone of the site, refer to section 5.1 of the Open API Access Guide. | [optional] 
**Type** | Pointer to **int32** | Site type(only for pro controller). It should be a value as follows: 0: Basic Site; 1: Pro Site | [optional] 
**Wan** | Pointer to **bool** | Whether exists gateway is connected of current site | [optional] 

## Methods

### NewMspSiteSummaryInfo

`func NewMspSiteSummaryInfo() *MspSiteSummaryInfo`

NewMspSiteSummaryInfo instantiates a new MspSiteSummaryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspSiteSummaryInfoWithDefaults

`func NewMspSiteSummaryInfoWithDefaults() *MspSiteSummaryInfo`

NewMspSiteSummaryInfoWithDefaults instantiates a new MspSiteSummaryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MspSiteSummaryInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MspSiteSummaryInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MspSiteSummaryInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MspSiteSummaryInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetConnectedApNum

`func (o *MspSiteSummaryInfo) GetConnectedApNum() int64`

GetConnectedApNum returns the ConnectedApNum field if non-nil, zero value otherwise.

### GetConnectedApNumOk

`func (o *MspSiteSummaryInfo) GetConnectedApNumOk() (*int64, bool)`

GetConnectedApNumOk returns a tuple with the ConnectedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedApNum

`func (o *MspSiteSummaryInfo) SetConnectedApNum(v int64)`

SetConnectedApNum sets ConnectedApNum field to given value.

### HasConnectedApNum

`func (o *MspSiteSummaryInfo) HasConnectedApNum() bool`

HasConnectedApNum returns a boolean if a field has been set.

### GetConnectedSwitchNum

`func (o *MspSiteSummaryInfo) GetConnectedSwitchNum() int64`

GetConnectedSwitchNum returns the ConnectedSwitchNum field if non-nil, zero value otherwise.

### GetConnectedSwitchNumOk

`func (o *MspSiteSummaryInfo) GetConnectedSwitchNumOk() (*int64, bool)`

GetConnectedSwitchNumOk returns a tuple with the ConnectedSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedSwitchNum

`func (o *MspSiteSummaryInfo) SetConnectedSwitchNum(v int64)`

SetConnectedSwitchNum sets ConnectedSwitchNum field to given value.

### HasConnectedSwitchNum

`func (o *MspSiteSummaryInfo) HasConnectedSwitchNum() bool`

HasConnectedSwitchNum returns a boolean if a field has been set.

### GetCustomerId

`func (o *MspSiteSummaryInfo) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *MspSiteSummaryInfo) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *MspSiteSummaryInfo) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *MspSiteSummaryInfo) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetCustomerName

`func (o *MspSiteSummaryInfo) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *MspSiteSummaryInfo) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *MspSiteSummaryInfo) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *MspSiteSummaryInfo) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetDisconnectedApNum

`func (o *MspSiteSummaryInfo) GetDisconnectedApNum() int64`

GetDisconnectedApNum returns the DisconnectedApNum field if non-nil, zero value otherwise.

### GetDisconnectedApNumOk

`func (o *MspSiteSummaryInfo) GetDisconnectedApNumOk() (*int64, bool)`

GetDisconnectedApNumOk returns a tuple with the DisconnectedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedApNum

`func (o *MspSiteSummaryInfo) SetDisconnectedApNum(v int64)`

SetDisconnectedApNum sets DisconnectedApNum field to given value.

### HasDisconnectedApNum

`func (o *MspSiteSummaryInfo) HasDisconnectedApNum() bool`

HasDisconnectedApNum returns a boolean if a field has been set.

### GetDisconnectedSwitchNum

`func (o *MspSiteSummaryInfo) GetDisconnectedSwitchNum() int64`

GetDisconnectedSwitchNum returns the DisconnectedSwitchNum field if non-nil, zero value otherwise.

### GetDisconnectedSwitchNumOk

`func (o *MspSiteSummaryInfo) GetDisconnectedSwitchNumOk() (*int64, bool)`

GetDisconnectedSwitchNumOk returns a tuple with the DisconnectedSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSwitchNum

`func (o *MspSiteSummaryInfo) SetDisconnectedSwitchNum(v int64)`

SetDisconnectedSwitchNum sets DisconnectedSwitchNum field to given value.

### HasDisconnectedSwitchNum

`func (o *MspSiteSummaryInfo) HasDisconnectedSwitchNum() bool`

HasDisconnectedSwitchNum returns a boolean if a field has been set.

### GetIsolatedApNum

`func (o *MspSiteSummaryInfo) GetIsolatedApNum() int64`

GetIsolatedApNum returns the IsolatedApNum field if non-nil, zero value otherwise.

### GetIsolatedApNumOk

`func (o *MspSiteSummaryInfo) GetIsolatedApNumOk() (*int64, bool)`

GetIsolatedApNumOk returns a tuple with the IsolatedApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsolatedApNum

`func (o *MspSiteSummaryInfo) SetIsolatedApNum(v int64)`

SetIsolatedApNum sets IsolatedApNum field to given value.

### HasIsolatedApNum

`func (o *MspSiteSummaryInfo) HasIsolatedApNum() bool`

HasIsolatedApNum returns a boolean if a field has been set.

### GetLatitude

`func (o *MspSiteSummaryInfo) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *MspSiteSummaryInfo) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *MspSiteSummaryInfo) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *MspSiteSummaryInfo) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *MspSiteSummaryInfo) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *MspSiteSummaryInfo) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *MspSiteSummaryInfo) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *MspSiteSummaryInfo) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetRegion

`func (o *MspSiteSummaryInfo) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MspSiteSummaryInfo) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MspSiteSummaryInfo) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MspSiteSummaryInfo) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetScenario

`func (o *MspSiteSummaryInfo) GetScenario() string`

GetScenario returns the Scenario field if non-nil, zero value otherwise.

### GetScenarioOk

`func (o *MspSiteSummaryInfo) GetScenarioOk() (*string, bool)`

GetScenarioOk returns a tuple with the Scenario field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScenario

`func (o *MspSiteSummaryInfo) SetScenario(v string)`

SetScenario sets Scenario field to given value.

### HasScenario

`func (o *MspSiteSummaryInfo) HasScenario() bool`

HasScenario returns a boolean if a field has been set.

### GetSiteId

`func (o *MspSiteSummaryInfo) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MspSiteSummaryInfo) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MspSiteSummaryInfo) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MspSiteSummaryInfo) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *MspSiteSummaryInfo) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *MspSiteSummaryInfo) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *MspSiteSummaryInfo) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *MspSiteSummaryInfo) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSitePublicIp

`func (o *MspSiteSummaryInfo) GetSitePublicIp() string`

GetSitePublicIp returns the SitePublicIp field if non-nil, zero value otherwise.

### GetSitePublicIpOk

`func (o *MspSiteSummaryInfo) GetSitePublicIpOk() (*string, bool)`

GetSitePublicIpOk returns a tuple with the SitePublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitePublicIp

`func (o *MspSiteSummaryInfo) SetSitePublicIp(v string)`

SetSitePublicIp sets SitePublicIp field to given value.

### HasSitePublicIp

`func (o *MspSiteSummaryInfo) HasSitePublicIp() bool`

HasSitePublicIp returns a boolean if a field has been set.

### GetSupportES

`func (o *MspSiteSummaryInfo) GetSupportES() bool`

GetSupportES returns the SupportES field if non-nil, zero value otherwise.

### GetSupportESOk

`func (o *MspSiteSummaryInfo) GetSupportESOk() (*bool, bool)`

GetSupportESOk returns a tuple with the SupportES field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportES

`func (o *MspSiteSummaryInfo) SetSupportES(v bool)`

SetSupportES sets SupportES field to given value.

### HasSupportES

`func (o *MspSiteSummaryInfo) HasSupportES() bool`

HasSupportES returns a boolean if a field has been set.

### GetSupportL2

`func (o *MspSiteSummaryInfo) GetSupportL2() bool`

GetSupportL2 returns the SupportL2 field if non-nil, zero value otherwise.

### GetSupportL2Ok

`func (o *MspSiteSummaryInfo) GetSupportL2Ok() (*bool, bool)`

GetSupportL2Ok returns a tuple with the SupportL2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportL2

`func (o *MspSiteSummaryInfo) SetSupportL2(v bool)`

SetSupportL2 sets SupportL2 field to given value.

### HasSupportL2

`func (o *MspSiteSummaryInfo) HasSupportL2() bool`

HasSupportL2 returns a boolean if a field has been set.

### GetTagIds

`func (o *MspSiteSummaryInfo) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *MspSiteSummaryInfo) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *MspSiteSummaryInfo) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *MspSiteSummaryInfo) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.

### GetTimeZone

`func (o *MspSiteSummaryInfo) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *MspSiteSummaryInfo) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *MspSiteSummaryInfo) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *MspSiteSummaryInfo) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetType

`func (o *MspSiteSummaryInfo) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MspSiteSummaryInfo) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MspSiteSummaryInfo) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *MspSiteSummaryInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWan

`func (o *MspSiteSummaryInfo) GetWan() bool`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *MspSiteSummaryInfo) GetWanOk() (*bool, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *MspSiteSummaryInfo) SetWan(v bool)`

SetWan sets Wan field to given value.

### HasWan

`func (o *MspSiteSummaryInfo) HasWan() bool`

HasWan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


