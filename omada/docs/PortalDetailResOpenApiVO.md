# PortalDetailResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTimeout** | Pointer to [**AuthTimeOpenApiVO**](AuthTimeOpenApiVO.md) |  | [optional] 
**AuthType** | Pointer to **int32** | Auth Type, should be a value as follows: &lt;br/&gt;0：No Authentication; 1：Simple Password;&lt;br/&gt;2: External Radius Server; 4：External Portal Server;&lt;br/&gt;11：Hotspot; 15: Ldap; 16: Social Login. | [optional] 
**Enable** | Pointer to **bool** | Portal enable status | [optional] 
**ExternalPortal** | Pointer to [**ExternalServerPortalSetting**](ExternalServerPortalSetting.md) |  | [optional] 
**ExternalRadius** | Pointer to [**ExternalRadiusSettingResOpenApiVO**](ExternalRadiusSettingResOpenApiVO.md) |  | [optional] 
**Google** | Pointer to [**GoogleOAuthSettingOpenApiVO**](GoogleOAuthSettingOpenApiVO.md) |  | [optional] 
**Hotspot** | Pointer to [**HotspotSetting**](HotspotSetting.md) |  | [optional] 
**HotspotRadius** | Pointer to [**HotspotRadiusSettingResOpenApiVO**](HotspotRadiusSettingResOpenApiVO.md) |  | [optional] 
**HttpsRedirectEnable** | Pointer to **bool** | With this option enabled, unauthenticated clients will be redirected to the Portal page when they are trying to browse HTTPS websites. | [optional] 
**Id** | Pointer to **string** | Portal ID | [optional] 
**LandingPage** | Pointer to **int32** | LandingPage enum, should be a value as follows: 1: Redirect to the original URL, 2: Redirect to Promotional URL. &lt;br/&gt;With The Original URL selected, clients are directed to the URL they request for after they pass Portal authentication. &lt;br/&gt;With The Promotional URL selected, clients are directed to the specified URL here after they pass Portal authentication. | [optional] 
**LandingUrl** | Pointer to **string** | If Parameter [landingPage] is 2(Redirect to Promotional URL),this Parameter is requested. | [optional] 
**LandingUrlScheme** | Pointer to **string** | If Parameter [landingPage] is 2(Redirect to Promotional URL),this Parameter is requested, content is http or https. | [optional] 
**Ldap** | Pointer to [**LdapSetting**](LdapSetting.md) |  | [optional] 
**Name** | Pointer to **string** | Portal name, should contain 1 to 128 characters | [optional] 
**NetworkList** | Pointer to **[]string** | Lan network ID list bound with this Portal. LAN Network can be created using &#39;Create LAN network&#39; (&#39;Create LAN network template&#39;) interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; (&#39;Get LAN network template list&#39;) interface | [optional] 
**NoAuth** | Pointer to [**NoAuthSetting**](NoAuthSetting.md) |  | [optional] 
**PortalFormId** | Pointer to **string** | Portal form ID, required when [authType] is 11 and hotspot [enabledTypes] contains 12. Portal form can be created using &#39;Create a new authentication survey&#39; interface, and Portal form ID can be obtained from &#39;Get authentication survey list&#39; interface | [optional] 
**SimplePassword** | Pointer to [**SimplePasswordSetting**](SimplePasswordSetting.md) |  | [optional] 
**Sms** | Pointer to [**SmsSettingResOpenApiVO**](SmsSettingResOpenApiVO.md) |  | [optional] 
**SocialLogin** | Pointer to [**SocialLoginSettingOpenApiVO**](SocialLoginSettingOpenApiVO.md) |  | [optional] 
**SsidList** | Pointer to **[]string** | SSID ID list bound with this Portal. SSID can be created using &#39;Create new SSID&#39; (&#39;Create new SSID template&#39;) interface, and SSID ID can be obtained from &#39;Get SSID list&#39; (&#39;Get SSID template list&#39;) interface | [optional] 

## Methods

### NewPortalDetailResOpenApiVO

`func NewPortalDetailResOpenApiVO() *PortalDetailResOpenApiVO`

NewPortalDetailResOpenApiVO instantiates a new PortalDetailResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalDetailResOpenApiVOWithDefaults

`func NewPortalDetailResOpenApiVOWithDefaults() *PortalDetailResOpenApiVO`

NewPortalDetailResOpenApiVOWithDefaults instantiates a new PortalDetailResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTimeout

`func (o *PortalDetailResOpenApiVO) GetAuthTimeout() AuthTimeOpenApiVO`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *PortalDetailResOpenApiVO) GetAuthTimeoutOk() (*AuthTimeOpenApiVO, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *PortalDetailResOpenApiVO) SetAuthTimeout(v AuthTimeOpenApiVO)`

SetAuthTimeout sets AuthTimeout field to given value.

### HasAuthTimeout

`func (o *PortalDetailResOpenApiVO) HasAuthTimeout() bool`

HasAuthTimeout returns a boolean if a field has been set.

### GetAuthType

`func (o *PortalDetailResOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *PortalDetailResOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *PortalDetailResOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *PortalDetailResOpenApiVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetEnable

`func (o *PortalDetailResOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PortalDetailResOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PortalDetailResOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PortalDetailResOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetExternalPortal

`func (o *PortalDetailResOpenApiVO) GetExternalPortal() ExternalServerPortalSetting`

GetExternalPortal returns the ExternalPortal field if non-nil, zero value otherwise.

### GetExternalPortalOk

`func (o *PortalDetailResOpenApiVO) GetExternalPortalOk() (*ExternalServerPortalSetting, bool)`

GetExternalPortalOk returns a tuple with the ExternalPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortal

`func (o *PortalDetailResOpenApiVO) SetExternalPortal(v ExternalServerPortalSetting)`

SetExternalPortal sets ExternalPortal field to given value.

### HasExternalPortal

`func (o *PortalDetailResOpenApiVO) HasExternalPortal() bool`

HasExternalPortal returns a boolean if a field has been set.

### GetExternalRadius

`func (o *PortalDetailResOpenApiVO) GetExternalRadius() ExternalRadiusSettingResOpenApiVO`

GetExternalRadius returns the ExternalRadius field if non-nil, zero value otherwise.

### GetExternalRadiusOk

`func (o *PortalDetailResOpenApiVO) GetExternalRadiusOk() (*ExternalRadiusSettingResOpenApiVO, bool)`

GetExternalRadiusOk returns a tuple with the ExternalRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRadius

`func (o *PortalDetailResOpenApiVO) SetExternalRadius(v ExternalRadiusSettingResOpenApiVO)`

SetExternalRadius sets ExternalRadius field to given value.

### HasExternalRadius

`func (o *PortalDetailResOpenApiVO) HasExternalRadius() bool`

HasExternalRadius returns a boolean if a field has been set.

### GetGoogle

`func (o *PortalDetailResOpenApiVO) GetGoogle() GoogleOAuthSettingOpenApiVO`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *PortalDetailResOpenApiVO) GetGoogleOk() (*GoogleOAuthSettingOpenApiVO, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *PortalDetailResOpenApiVO) SetGoogle(v GoogleOAuthSettingOpenApiVO)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *PortalDetailResOpenApiVO) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### GetHotspot

`func (o *PortalDetailResOpenApiVO) GetHotspot() HotspotSetting`

GetHotspot returns the Hotspot field if non-nil, zero value otherwise.

### GetHotspotOk

`func (o *PortalDetailResOpenApiVO) GetHotspotOk() (*HotspotSetting, bool)`

GetHotspotOk returns a tuple with the Hotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspot

`func (o *PortalDetailResOpenApiVO) SetHotspot(v HotspotSetting)`

SetHotspot sets Hotspot field to given value.

### HasHotspot

`func (o *PortalDetailResOpenApiVO) HasHotspot() bool`

HasHotspot returns a boolean if a field has been set.

### GetHotspotRadius

`func (o *PortalDetailResOpenApiVO) GetHotspotRadius() HotspotRadiusSettingResOpenApiVO`

GetHotspotRadius returns the HotspotRadius field if non-nil, zero value otherwise.

### GetHotspotRadiusOk

`func (o *PortalDetailResOpenApiVO) GetHotspotRadiusOk() (*HotspotRadiusSettingResOpenApiVO, bool)`

GetHotspotRadiusOk returns a tuple with the HotspotRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspotRadius

`func (o *PortalDetailResOpenApiVO) SetHotspotRadius(v HotspotRadiusSettingResOpenApiVO)`

SetHotspotRadius sets HotspotRadius field to given value.

### HasHotspotRadius

`func (o *PortalDetailResOpenApiVO) HasHotspotRadius() bool`

HasHotspotRadius returns a boolean if a field has been set.

### GetHttpsRedirectEnable

`func (o *PortalDetailResOpenApiVO) GetHttpsRedirectEnable() bool`

GetHttpsRedirectEnable returns the HttpsRedirectEnable field if non-nil, zero value otherwise.

### GetHttpsRedirectEnableOk

`func (o *PortalDetailResOpenApiVO) GetHttpsRedirectEnableOk() (*bool, bool)`

GetHttpsRedirectEnableOk returns a tuple with the HttpsRedirectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRedirectEnable

`func (o *PortalDetailResOpenApiVO) SetHttpsRedirectEnable(v bool)`

SetHttpsRedirectEnable sets HttpsRedirectEnable field to given value.

### HasHttpsRedirectEnable

`func (o *PortalDetailResOpenApiVO) HasHttpsRedirectEnable() bool`

HasHttpsRedirectEnable returns a boolean if a field has been set.

### GetId

`func (o *PortalDetailResOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalDetailResOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalDetailResOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortalDetailResOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLandingPage

`func (o *PortalDetailResOpenApiVO) GetLandingPage() int32`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *PortalDetailResOpenApiVO) GetLandingPageOk() (*int32, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *PortalDetailResOpenApiVO) SetLandingPage(v int32)`

SetLandingPage sets LandingPage field to given value.

### HasLandingPage

`func (o *PortalDetailResOpenApiVO) HasLandingPage() bool`

HasLandingPage returns a boolean if a field has been set.

### GetLandingUrl

`func (o *PortalDetailResOpenApiVO) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *PortalDetailResOpenApiVO) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *PortalDetailResOpenApiVO) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *PortalDetailResOpenApiVO) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### GetLandingUrlScheme

`func (o *PortalDetailResOpenApiVO) GetLandingUrlScheme() string`

GetLandingUrlScheme returns the LandingUrlScheme field if non-nil, zero value otherwise.

### GetLandingUrlSchemeOk

`func (o *PortalDetailResOpenApiVO) GetLandingUrlSchemeOk() (*string, bool)`

GetLandingUrlSchemeOk returns a tuple with the LandingUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrlScheme

`func (o *PortalDetailResOpenApiVO) SetLandingUrlScheme(v string)`

SetLandingUrlScheme sets LandingUrlScheme field to given value.

### HasLandingUrlScheme

`func (o *PortalDetailResOpenApiVO) HasLandingUrlScheme() bool`

HasLandingUrlScheme returns a boolean if a field has been set.

### GetLdap

`func (o *PortalDetailResOpenApiVO) GetLdap() LdapSetting`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *PortalDetailResOpenApiVO) GetLdapOk() (*LdapSetting, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *PortalDetailResOpenApiVO) SetLdap(v LdapSetting)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *PortalDetailResOpenApiVO) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetName

`func (o *PortalDetailResOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortalDetailResOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortalDetailResOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortalDetailResOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkList

`func (o *PortalDetailResOpenApiVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *PortalDetailResOpenApiVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *PortalDetailResOpenApiVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *PortalDetailResOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNoAuth

`func (o *PortalDetailResOpenApiVO) GetNoAuth() NoAuthSetting`

GetNoAuth returns the NoAuth field if non-nil, zero value otherwise.

### GetNoAuthOk

`func (o *PortalDetailResOpenApiVO) GetNoAuthOk() (*NoAuthSetting, bool)`

GetNoAuthOk returns a tuple with the NoAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAuth

`func (o *PortalDetailResOpenApiVO) SetNoAuth(v NoAuthSetting)`

SetNoAuth sets NoAuth field to given value.

### HasNoAuth

`func (o *PortalDetailResOpenApiVO) HasNoAuth() bool`

HasNoAuth returns a boolean if a field has been set.

### GetPortalFormId

`func (o *PortalDetailResOpenApiVO) GetPortalFormId() string`

GetPortalFormId returns the PortalFormId field if non-nil, zero value otherwise.

### GetPortalFormIdOk

`func (o *PortalDetailResOpenApiVO) GetPortalFormIdOk() (*string, bool)`

GetPortalFormIdOk returns a tuple with the PortalFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalFormId

`func (o *PortalDetailResOpenApiVO) SetPortalFormId(v string)`

SetPortalFormId sets PortalFormId field to given value.

### HasPortalFormId

`func (o *PortalDetailResOpenApiVO) HasPortalFormId() bool`

HasPortalFormId returns a boolean if a field has been set.

### GetSimplePassword

`func (o *PortalDetailResOpenApiVO) GetSimplePassword() SimplePasswordSetting`

GetSimplePassword returns the SimplePassword field if non-nil, zero value otherwise.

### GetSimplePasswordOk

`func (o *PortalDetailResOpenApiVO) GetSimplePasswordOk() (*SimplePasswordSetting, bool)`

GetSimplePasswordOk returns a tuple with the SimplePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePassword

`func (o *PortalDetailResOpenApiVO) SetSimplePassword(v SimplePasswordSetting)`

SetSimplePassword sets SimplePassword field to given value.

### HasSimplePassword

`func (o *PortalDetailResOpenApiVO) HasSimplePassword() bool`

HasSimplePassword returns a boolean if a field has been set.

### GetSms

`func (o *PortalDetailResOpenApiVO) GetSms() SmsSettingResOpenApiVO`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *PortalDetailResOpenApiVO) GetSmsOk() (*SmsSettingResOpenApiVO, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *PortalDetailResOpenApiVO) SetSms(v SmsSettingResOpenApiVO)`

SetSms sets Sms field to given value.

### HasSms

`func (o *PortalDetailResOpenApiVO) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetSocialLogin

`func (o *PortalDetailResOpenApiVO) GetSocialLogin() SocialLoginSettingOpenApiVO`

GetSocialLogin returns the SocialLogin field if non-nil, zero value otherwise.

### GetSocialLoginOk

`func (o *PortalDetailResOpenApiVO) GetSocialLoginOk() (*SocialLoginSettingOpenApiVO, bool)`

GetSocialLoginOk returns a tuple with the SocialLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLogin

`func (o *PortalDetailResOpenApiVO) SetSocialLogin(v SocialLoginSettingOpenApiVO)`

SetSocialLogin sets SocialLogin field to given value.

### HasSocialLogin

`func (o *PortalDetailResOpenApiVO) HasSocialLogin() bool`

HasSocialLogin returns a boolean if a field has been set.

### GetSsidList

`func (o *PortalDetailResOpenApiVO) GetSsidList() []string`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *PortalDetailResOpenApiVO) GetSsidListOk() (*[]string, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *PortalDetailResOpenApiVO) SetSsidList(v []string)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *PortalDetailResOpenApiVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


