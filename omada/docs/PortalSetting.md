# PortalSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTimeout** | [**AuthTimeoutSetting**](AuthTimeoutSetting.md) |  | 
**AuthType** | **int32** | Auth Type, should be a value as follows: &lt;br/&gt;0：No Authentication; 1：Simple Password;&lt;br/&gt;2: External Radius Server; 4：External Portal Server;&lt;br/&gt;11：Hotspot; 15: Ldap; 16: Social Login | 
**Enable** | **bool** | Portal Enable | 
**ExternalPortal** | Pointer to [**ExternalServerPortalSetting**](ExternalServerPortalSetting.md) |  | [optional] 
**ExternalRadius** | Pointer to [**ExternalRadiusSetting**](ExternalRadiusSetting.md) |  | [optional] 
**Google** | Pointer to [**GoogleOAuthSettingOpenApiVO**](GoogleOAuthSettingOpenApiVO.md) |  | [optional] 
**Hotspot** | Pointer to [**HotspotSetting**](HotspotSetting.md) |  | [optional] 
**HotspotRadius** | Pointer to [**HotspotRadiusSetting**](HotspotRadiusSetting.md) |  | [optional] 
**HttpsRedirectEnable** | **bool** | With this option enabled, unauthenticated clients will be redirected to the Portal page when they are trying to browse HTTPS websites. | 
**ImportedPortalPage** | Pointer to [**ImportedPortalPageOpenApiVO**](ImportedPortalPageOpenApiVO.md) |  | [optional] 
**LandingPage** | **int32** | LandingPage enum: 1: Redirect to the original URL, 2: Redirect to Promotional URL, 3: Redirect to Logout Page. &lt;br/&gt;With The Original URL selected, clients are directed to the URL they request for after they pass Portal authentication. &lt;br/&gt;With The Promotional URL selected, clients are directed to the specified URL here after they pass Portal authentication. &lt;br/&gt;With The Logout Page selected, clients are directed to the Portal Logout Page after they pass Portal authentication.  | 
**LandingUrl** | Pointer to **string** | If Parameter [landingPage] is 2(Redirect to Promotional URL),this Parameter is requested. | [optional] 
**LandingUrlScheme** | Pointer to **string** | If Parameter [landingPage] is 2(Redirect to Promotional URL),this Parameter is requested, content is http or https. | [optional] 
**Ldap** | Pointer to [**LdapSetting**](LdapSetting.md) |  | [optional] 
**Name** | **string** | Portal Name, should contain 1 to 128 characters | 
**NetworkList** | Pointer to **[]string** | Lan network ID list bound with this Portal. LAN Network can be created using &#39;Create LAN network&#39; (&#39;Create LAN network template&#39;) interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; (&#39;Get LAN network template list&#39;) interface | [optional] 
**NoAuth** | Pointer to [**NoAuthSetting**](NoAuthSetting.md) |  | [optional] 
**PageType** | Pointer to **int32** | Page type, should be a value as follows: 1: Use default page, 2: use uploaded page. When [pageType] is null, it defaults to 1 | [optional] 
**PortalCustomize** | Pointer to [**PortalCustomizeOpenApiVO**](PortalCustomizeOpenApiVO.md) |  | [optional] 
**PortalFormId** | Pointer to **string** | Portal form ID, required when [authType] is 11 and hotspot [enabledTypes] contains 12. | [optional] 
**SimplePassword** | Pointer to [**SimplePasswordSetting**](SimplePasswordSetting.md) |  | [optional] 
**Sms** | Pointer to [**SmsSetting**](SmsSetting.md) |  | [optional] 
**SocialLogin** | Pointer to [**SocialLoginSettingOpenApiVO**](SocialLoginSettingOpenApiVO.md) |  | [optional] 
**SsidList** | Pointer to **[]string** | SSID ID list bound with this Portal. SSID can be created using &#39;Create new SSID&#39; (&#39;Create new SSID template&#39;) interface, and SSID ID can be obtained from &#39;Get SSID list&#39; (&#39;Get SSID template list&#39;) interface | [optional] 

## Methods

### NewPortalSetting

`func NewPortalSetting(authTimeout AuthTimeoutSetting, authType int32, enable bool, httpsRedirectEnable bool, landingPage int32, name string, ) *PortalSetting`

NewPortalSetting instantiates a new PortalSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalSettingWithDefaults

`func NewPortalSettingWithDefaults() *PortalSetting`

NewPortalSettingWithDefaults instantiates a new PortalSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTimeout

`func (o *PortalSetting) GetAuthTimeout() AuthTimeoutSetting`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *PortalSetting) GetAuthTimeoutOk() (*AuthTimeoutSetting, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *PortalSetting) SetAuthTimeout(v AuthTimeoutSetting)`

SetAuthTimeout sets AuthTimeout field to given value.


### GetAuthType

`func (o *PortalSetting) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *PortalSetting) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *PortalSetting) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.


### GetEnable

`func (o *PortalSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PortalSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PortalSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetExternalPortal

`func (o *PortalSetting) GetExternalPortal() ExternalServerPortalSetting`

GetExternalPortal returns the ExternalPortal field if non-nil, zero value otherwise.

### GetExternalPortalOk

`func (o *PortalSetting) GetExternalPortalOk() (*ExternalServerPortalSetting, bool)`

GetExternalPortalOk returns a tuple with the ExternalPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortal

`func (o *PortalSetting) SetExternalPortal(v ExternalServerPortalSetting)`

SetExternalPortal sets ExternalPortal field to given value.

### HasExternalPortal

`func (o *PortalSetting) HasExternalPortal() bool`

HasExternalPortal returns a boolean if a field has been set.

### GetExternalRadius

`func (o *PortalSetting) GetExternalRadius() ExternalRadiusSetting`

GetExternalRadius returns the ExternalRadius field if non-nil, zero value otherwise.

### GetExternalRadiusOk

`func (o *PortalSetting) GetExternalRadiusOk() (*ExternalRadiusSetting, bool)`

GetExternalRadiusOk returns a tuple with the ExternalRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalRadius

`func (o *PortalSetting) SetExternalRadius(v ExternalRadiusSetting)`

SetExternalRadius sets ExternalRadius field to given value.

### HasExternalRadius

`func (o *PortalSetting) HasExternalRadius() bool`

HasExternalRadius returns a boolean if a field has been set.

### GetGoogle

`func (o *PortalSetting) GetGoogle() GoogleOAuthSettingOpenApiVO`

GetGoogle returns the Google field if non-nil, zero value otherwise.

### GetGoogleOk

`func (o *PortalSetting) GetGoogleOk() (*GoogleOAuthSettingOpenApiVO, bool)`

GetGoogleOk returns a tuple with the Google field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogle

`func (o *PortalSetting) SetGoogle(v GoogleOAuthSettingOpenApiVO)`

SetGoogle sets Google field to given value.

### HasGoogle

`func (o *PortalSetting) HasGoogle() bool`

HasGoogle returns a boolean if a field has been set.

### GetHotspot

`func (o *PortalSetting) GetHotspot() HotspotSetting`

GetHotspot returns the Hotspot field if non-nil, zero value otherwise.

### GetHotspotOk

`func (o *PortalSetting) GetHotspotOk() (*HotspotSetting, bool)`

GetHotspotOk returns a tuple with the Hotspot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspot

`func (o *PortalSetting) SetHotspot(v HotspotSetting)`

SetHotspot sets Hotspot field to given value.

### HasHotspot

`func (o *PortalSetting) HasHotspot() bool`

HasHotspot returns a boolean if a field has been set.

### GetHotspotRadius

`func (o *PortalSetting) GetHotspotRadius() HotspotRadiusSetting`

GetHotspotRadius returns the HotspotRadius field if non-nil, zero value otherwise.

### GetHotspotRadiusOk

`func (o *PortalSetting) GetHotspotRadiusOk() (*HotspotRadiusSetting, bool)`

GetHotspotRadiusOk returns a tuple with the HotspotRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotspotRadius

`func (o *PortalSetting) SetHotspotRadius(v HotspotRadiusSetting)`

SetHotspotRadius sets HotspotRadius field to given value.

### HasHotspotRadius

`func (o *PortalSetting) HasHotspotRadius() bool`

HasHotspotRadius returns a boolean if a field has been set.

### GetHttpsRedirectEnable

`func (o *PortalSetting) GetHttpsRedirectEnable() bool`

GetHttpsRedirectEnable returns the HttpsRedirectEnable field if non-nil, zero value otherwise.

### GetHttpsRedirectEnableOk

`func (o *PortalSetting) GetHttpsRedirectEnableOk() (*bool, bool)`

GetHttpsRedirectEnableOk returns a tuple with the HttpsRedirectEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsRedirectEnable

`func (o *PortalSetting) SetHttpsRedirectEnable(v bool)`

SetHttpsRedirectEnable sets HttpsRedirectEnable field to given value.


### GetImportedPortalPage

`func (o *PortalSetting) GetImportedPortalPage() ImportedPortalPageOpenApiVO`

GetImportedPortalPage returns the ImportedPortalPage field if non-nil, zero value otherwise.

### GetImportedPortalPageOk

`func (o *PortalSetting) GetImportedPortalPageOk() (*ImportedPortalPageOpenApiVO, bool)`

GetImportedPortalPageOk returns a tuple with the ImportedPortalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedPortalPage

`func (o *PortalSetting) SetImportedPortalPage(v ImportedPortalPageOpenApiVO)`

SetImportedPortalPage sets ImportedPortalPage field to given value.

### HasImportedPortalPage

`func (o *PortalSetting) HasImportedPortalPage() bool`

HasImportedPortalPage returns a boolean if a field has been set.

### GetLandingPage

`func (o *PortalSetting) GetLandingPage() int32`

GetLandingPage returns the LandingPage field if non-nil, zero value otherwise.

### GetLandingPageOk

`func (o *PortalSetting) GetLandingPageOk() (*int32, bool)`

GetLandingPageOk returns a tuple with the LandingPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingPage

`func (o *PortalSetting) SetLandingPage(v int32)`

SetLandingPage sets LandingPage field to given value.


### GetLandingUrl

`func (o *PortalSetting) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *PortalSetting) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *PortalSetting) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *PortalSetting) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### GetLandingUrlScheme

`func (o *PortalSetting) GetLandingUrlScheme() string`

GetLandingUrlScheme returns the LandingUrlScheme field if non-nil, zero value otherwise.

### GetLandingUrlSchemeOk

`func (o *PortalSetting) GetLandingUrlSchemeOk() (*string, bool)`

GetLandingUrlSchemeOk returns a tuple with the LandingUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrlScheme

`func (o *PortalSetting) SetLandingUrlScheme(v string)`

SetLandingUrlScheme sets LandingUrlScheme field to given value.

### HasLandingUrlScheme

`func (o *PortalSetting) HasLandingUrlScheme() bool`

HasLandingUrlScheme returns a boolean if a field has been set.

### GetLdap

`func (o *PortalSetting) GetLdap() LdapSetting`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *PortalSetting) GetLdapOk() (*LdapSetting, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *PortalSetting) SetLdap(v LdapSetting)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *PortalSetting) HasLdap() bool`

HasLdap returns a boolean if a field has been set.

### GetName

`func (o *PortalSetting) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortalSetting) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortalSetting) SetName(v string)`

SetName sets Name field to given value.


### GetNetworkList

`func (o *PortalSetting) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *PortalSetting) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *PortalSetting) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *PortalSetting) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetNoAuth

`func (o *PortalSetting) GetNoAuth() NoAuthSetting`

GetNoAuth returns the NoAuth field if non-nil, zero value otherwise.

### GetNoAuthOk

`func (o *PortalSetting) GetNoAuthOk() (*NoAuthSetting, bool)`

GetNoAuthOk returns a tuple with the NoAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAuth

`func (o *PortalSetting) SetNoAuth(v NoAuthSetting)`

SetNoAuth sets NoAuth field to given value.

### HasNoAuth

`func (o *PortalSetting) HasNoAuth() bool`

HasNoAuth returns a boolean if a field has been set.

### GetPageType

`func (o *PortalSetting) GetPageType() int32`

GetPageType returns the PageType field if non-nil, zero value otherwise.

### GetPageTypeOk

`func (o *PortalSetting) GetPageTypeOk() (*int32, bool)`

GetPageTypeOk returns a tuple with the PageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageType

`func (o *PortalSetting) SetPageType(v int32)`

SetPageType sets PageType field to given value.

### HasPageType

`func (o *PortalSetting) HasPageType() bool`

HasPageType returns a boolean if a field has been set.

### GetPortalCustomize

`func (o *PortalSetting) GetPortalCustomize() PortalCustomizeOpenApiVO`

GetPortalCustomize returns the PortalCustomize field if non-nil, zero value otherwise.

### GetPortalCustomizeOk

`func (o *PortalSetting) GetPortalCustomizeOk() (*PortalCustomizeOpenApiVO, bool)`

GetPortalCustomizeOk returns a tuple with the PortalCustomize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalCustomize

`func (o *PortalSetting) SetPortalCustomize(v PortalCustomizeOpenApiVO)`

SetPortalCustomize sets PortalCustomize field to given value.

### HasPortalCustomize

`func (o *PortalSetting) HasPortalCustomize() bool`

HasPortalCustomize returns a boolean if a field has been set.

### GetPortalFormId

`func (o *PortalSetting) GetPortalFormId() string`

GetPortalFormId returns the PortalFormId field if non-nil, zero value otherwise.

### GetPortalFormIdOk

`func (o *PortalSetting) GetPortalFormIdOk() (*string, bool)`

GetPortalFormIdOk returns a tuple with the PortalFormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalFormId

`func (o *PortalSetting) SetPortalFormId(v string)`

SetPortalFormId sets PortalFormId field to given value.

### HasPortalFormId

`func (o *PortalSetting) HasPortalFormId() bool`

HasPortalFormId returns a boolean if a field has been set.

### GetSimplePassword

`func (o *PortalSetting) GetSimplePassword() SimplePasswordSetting`

GetSimplePassword returns the SimplePassword field if non-nil, zero value otherwise.

### GetSimplePasswordOk

`func (o *PortalSetting) GetSimplePasswordOk() (*SimplePasswordSetting, bool)`

GetSimplePasswordOk returns a tuple with the SimplePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimplePassword

`func (o *PortalSetting) SetSimplePassword(v SimplePasswordSetting)`

SetSimplePassword sets SimplePassword field to given value.

### HasSimplePassword

`func (o *PortalSetting) HasSimplePassword() bool`

HasSimplePassword returns a boolean if a field has been set.

### GetSms

`func (o *PortalSetting) GetSms() SmsSetting`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *PortalSetting) GetSmsOk() (*SmsSetting, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *PortalSetting) SetSms(v SmsSetting)`

SetSms sets Sms field to given value.

### HasSms

`func (o *PortalSetting) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetSocialLogin

`func (o *PortalSetting) GetSocialLogin() SocialLoginSettingOpenApiVO`

GetSocialLogin returns the SocialLogin field if non-nil, zero value otherwise.

### GetSocialLoginOk

`func (o *PortalSetting) GetSocialLoginOk() (*SocialLoginSettingOpenApiVO, bool)`

GetSocialLoginOk returns a tuple with the SocialLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialLogin

`func (o *PortalSetting) SetSocialLogin(v SocialLoginSettingOpenApiVO)`

SetSocialLogin sets SocialLogin field to given value.

### HasSocialLogin

`func (o *PortalSetting) HasSocialLogin() bool`

HasSocialLogin returns a boolean if a field has been set.

### GetSsidList

`func (o *PortalSetting) GetSsidList() []string`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *PortalSetting) GetSsidListOk() (*[]string, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *PortalSetting) SetSsidList(v []string)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *PortalSetting) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


