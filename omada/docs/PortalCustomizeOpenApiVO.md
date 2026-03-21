# PortalCustomizeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advertisement** | Pointer to [**AdvertisementSetting**](AdvertisementSetting.md) |  | [optional] 
**BackgroundPictureId** | Pointer to **string** | Background picture ID | [optional] 
**ButtonColor** | Pointer to **string** | Button color. Hex color code such as: #ffffff. | [optional] 
**ButtonOpacity** | Pointer to **int32** | Button opacity, should be within the range of 0–100. | [optional] 
**ButtonText** | Pointer to **string** | Button text, should contain 0 to 32 characters, default value is \&quot;Log In\&quot;. | [optional] 
**ButtonTextColor** | Pointer to **string** | Button text color. Hex color code such as: #ffffff. | [optional] 
**ButtonTextOpacity** | Pointer to **int32** | Button text opacity, should be within the range of 0–100. | [optional] 
**Copyright** | Pointer to **string** | Copyright text, should contain 0 to 200 characters. | [optional] 
**CopyrightEnable** | **bool** | Whether to dispaly the copyright. | 
**CopyrightTextColor** | Pointer to **string** | Copyright text color. Hex color code such as: #ffffff. | [optional] 
**CopyrightTextFontSize** | Pointer to **int32** | Copyright text font size, should be within the range of 12–18. | [optional] 
**CopyrightTextOpacity** | Pointer to **int32** | Copyright text opacity, should be within the range of 0–100. | [optional] 
**DefaultLanguage** | **int32** | The controller automatically adjusts the language displayed on the Portal page according to the system language of the clients.If the language is not supported, the controller will use the default language specified here.&lt;br/&gt;1: en_US (English); 3: cs_CZ (Český); 4: de_DE (Deutsch); 5: da_DK (Dansk); 6: el_GR (ελληνικά);&lt;br/&gt;7: fr_FR (Français); 8: es_ES (Español); 9: nl_NL (Nederlands); 10: it_IT (Italiano); 11: pl_PL (Polski);&lt;br/&gt;12: pt_PT (Português); 13: ru_RU (Русский); 14: sv_SE (Svenska); 15: tr_TR (Türkçe);&lt;br/&gt;16: ar_SA (لغة عربية); &lt;br/&gt;17: ja_JP (日本語); 18: zh_TW (中文(繁體)); 19: th_TH (ไทย); 20: vi_VN (Tiếng Việt); 21: ko_KR (한국어) | 
**FormAuthButtonText** | Pointer to **string** | Form auth button text, should contain 0 to 32 characters, required when [authType] is 11 and hotspot [enabledTypes] contains 12.Default value is \&quot;Take the Survey\&quot;. | [optional] 
**InputBoxColor** | Pointer to **string** | Input box color. Hex color code such as: #ffffff. | [optional] 
**InputBoxOpacity** | Pointer to **int32** | Input box opacity, should be within the range of 0–100. | [optional] 
**InputTextColor** | Pointer to **string** | Input text color. Hex color code such as: #ffffff. | [optional] 
**InputTextOpacity** | Pointer to **int32** | Input text opacity, should be within the range of 0–100. | [optional] 
**LogoDisplay** | **bool** | Whether to display the default logo. | 
**LogoPictureId** | Pointer to **string** | Logo picture ID | [optional] 
**RedirectionCountDownEnable** | Pointer to **bool** | Whether to show redirection countdown after authorized. | [optional] 
**TermsOfServiceEnable** | **bool** | Whether to display terms of service. | 
**TermsOfServiceFontSize** | Pointer to **int32** | Terms of service text font size, should be within the range of 12–18. | [optional] 
**TermsOfServiceText** | Pointer to **string** | Terms of service text, should contain 0 to 100 characters. | [optional] 
**TermsOfServiceUrlTexts** | Pointer to [**[]TermsOfServiceUrlVO**](TermsOfServiceUrlVO.md) | Terms of service url texts, match the termsOfServiceText and turn the matching characters into an openable link, Up to 3 entries are allowed for the list. | [optional] 
**WelcomeEnable** | **bool** | Whether to display the welcome info | 
**WelcomeInformation** | Pointer to **string** | Welcome Information, should contain 1 to 31 characters. | [optional] 
**WelcomeTextColor** | Pointer to **string** | Welcome text color. Hex color code such as: #ffffff. | [optional] 
**WelcomeTextFontSize** | Pointer to **int32** | Welcome text font size, should be within the range of 12–18. | [optional] 
**WelcomeTextOpacity** | Pointer to **int32** | Welcome text opacity, should be within the range of 0–100. | [optional] 

## Methods

### NewPortalCustomizeOpenApiVO

`func NewPortalCustomizeOpenApiVO(copyrightEnable bool, defaultLanguage int32, logoDisplay bool, termsOfServiceEnable bool, welcomeEnable bool, ) *PortalCustomizeOpenApiVO`

NewPortalCustomizeOpenApiVO instantiates a new PortalCustomizeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalCustomizeOpenApiVOWithDefaults

`func NewPortalCustomizeOpenApiVOWithDefaults() *PortalCustomizeOpenApiVO`

NewPortalCustomizeOpenApiVOWithDefaults instantiates a new PortalCustomizeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertisement

`func (o *PortalCustomizeOpenApiVO) GetAdvertisement() AdvertisementSetting`

GetAdvertisement returns the Advertisement field if non-nil, zero value otherwise.

### GetAdvertisementOk

`func (o *PortalCustomizeOpenApiVO) GetAdvertisementOk() (*AdvertisementSetting, bool)`

GetAdvertisementOk returns a tuple with the Advertisement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisement

`func (o *PortalCustomizeOpenApiVO) SetAdvertisement(v AdvertisementSetting)`

SetAdvertisement sets Advertisement field to given value.

### HasAdvertisement

`func (o *PortalCustomizeOpenApiVO) HasAdvertisement() bool`

HasAdvertisement returns a boolean if a field has been set.

### GetBackgroundPictureId

`func (o *PortalCustomizeOpenApiVO) GetBackgroundPictureId() string`

GetBackgroundPictureId returns the BackgroundPictureId field if non-nil, zero value otherwise.

### GetBackgroundPictureIdOk

`func (o *PortalCustomizeOpenApiVO) GetBackgroundPictureIdOk() (*string, bool)`

GetBackgroundPictureIdOk returns a tuple with the BackgroundPictureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPictureId

`func (o *PortalCustomizeOpenApiVO) SetBackgroundPictureId(v string)`

SetBackgroundPictureId sets BackgroundPictureId field to given value.

### HasBackgroundPictureId

`func (o *PortalCustomizeOpenApiVO) HasBackgroundPictureId() bool`

HasBackgroundPictureId returns a boolean if a field has been set.

### GetButtonColor

`func (o *PortalCustomizeOpenApiVO) GetButtonColor() string`

GetButtonColor returns the ButtonColor field if non-nil, zero value otherwise.

### GetButtonColorOk

`func (o *PortalCustomizeOpenApiVO) GetButtonColorOk() (*string, bool)`

GetButtonColorOk returns a tuple with the ButtonColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonColor

`func (o *PortalCustomizeOpenApiVO) SetButtonColor(v string)`

SetButtonColor sets ButtonColor field to given value.

### HasButtonColor

`func (o *PortalCustomizeOpenApiVO) HasButtonColor() bool`

HasButtonColor returns a boolean if a field has been set.

### GetButtonOpacity

`func (o *PortalCustomizeOpenApiVO) GetButtonOpacity() int32`

GetButtonOpacity returns the ButtonOpacity field if non-nil, zero value otherwise.

### GetButtonOpacityOk

`func (o *PortalCustomizeOpenApiVO) GetButtonOpacityOk() (*int32, bool)`

GetButtonOpacityOk returns a tuple with the ButtonOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonOpacity

`func (o *PortalCustomizeOpenApiVO) SetButtonOpacity(v int32)`

SetButtonOpacity sets ButtonOpacity field to given value.

### HasButtonOpacity

`func (o *PortalCustomizeOpenApiVO) HasButtonOpacity() bool`

HasButtonOpacity returns a boolean if a field has been set.

### GetButtonText

`func (o *PortalCustomizeOpenApiVO) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *PortalCustomizeOpenApiVO) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *PortalCustomizeOpenApiVO) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *PortalCustomizeOpenApiVO) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### GetButtonTextColor

`func (o *PortalCustomizeOpenApiVO) GetButtonTextColor() string`

GetButtonTextColor returns the ButtonTextColor field if non-nil, zero value otherwise.

### GetButtonTextColorOk

`func (o *PortalCustomizeOpenApiVO) GetButtonTextColorOk() (*string, bool)`

GetButtonTextColorOk returns a tuple with the ButtonTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTextColor

`func (o *PortalCustomizeOpenApiVO) SetButtonTextColor(v string)`

SetButtonTextColor sets ButtonTextColor field to given value.

### HasButtonTextColor

`func (o *PortalCustomizeOpenApiVO) HasButtonTextColor() bool`

HasButtonTextColor returns a boolean if a field has been set.

### GetButtonTextOpacity

`func (o *PortalCustomizeOpenApiVO) GetButtonTextOpacity() int32`

GetButtonTextOpacity returns the ButtonTextOpacity field if non-nil, zero value otherwise.

### GetButtonTextOpacityOk

`func (o *PortalCustomizeOpenApiVO) GetButtonTextOpacityOk() (*int32, bool)`

GetButtonTextOpacityOk returns a tuple with the ButtonTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTextOpacity

`func (o *PortalCustomizeOpenApiVO) SetButtonTextOpacity(v int32)`

SetButtonTextOpacity sets ButtonTextOpacity field to given value.

### HasButtonTextOpacity

`func (o *PortalCustomizeOpenApiVO) HasButtonTextOpacity() bool`

HasButtonTextOpacity returns a boolean if a field has been set.

### GetCopyright

`func (o *PortalCustomizeOpenApiVO) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *PortalCustomizeOpenApiVO) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *PortalCustomizeOpenApiVO) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *PortalCustomizeOpenApiVO) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### GetCopyrightEnable

`func (o *PortalCustomizeOpenApiVO) GetCopyrightEnable() bool`

GetCopyrightEnable returns the CopyrightEnable field if non-nil, zero value otherwise.

### GetCopyrightEnableOk

`func (o *PortalCustomizeOpenApiVO) GetCopyrightEnableOk() (*bool, bool)`

GetCopyrightEnableOk returns a tuple with the CopyrightEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightEnable

`func (o *PortalCustomizeOpenApiVO) SetCopyrightEnable(v bool)`

SetCopyrightEnable sets CopyrightEnable field to given value.


### GetCopyrightTextColor

`func (o *PortalCustomizeOpenApiVO) GetCopyrightTextColor() string`

GetCopyrightTextColor returns the CopyrightTextColor field if non-nil, zero value otherwise.

### GetCopyrightTextColorOk

`func (o *PortalCustomizeOpenApiVO) GetCopyrightTextColorOk() (*string, bool)`

GetCopyrightTextColorOk returns a tuple with the CopyrightTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightTextColor

`func (o *PortalCustomizeOpenApiVO) SetCopyrightTextColor(v string)`

SetCopyrightTextColor sets CopyrightTextColor field to given value.

### HasCopyrightTextColor

`func (o *PortalCustomizeOpenApiVO) HasCopyrightTextColor() bool`

HasCopyrightTextColor returns a boolean if a field has been set.

### GetCopyrightTextFontSize

`func (o *PortalCustomizeOpenApiVO) GetCopyrightTextFontSize() int32`

GetCopyrightTextFontSize returns the CopyrightTextFontSize field if non-nil, zero value otherwise.

### GetCopyrightTextFontSizeOk

`func (o *PortalCustomizeOpenApiVO) GetCopyrightTextFontSizeOk() (*int32, bool)`

GetCopyrightTextFontSizeOk returns a tuple with the CopyrightTextFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightTextFontSize

`func (o *PortalCustomizeOpenApiVO) SetCopyrightTextFontSize(v int32)`

SetCopyrightTextFontSize sets CopyrightTextFontSize field to given value.

### HasCopyrightTextFontSize

`func (o *PortalCustomizeOpenApiVO) HasCopyrightTextFontSize() bool`

HasCopyrightTextFontSize returns a boolean if a field has been set.

### GetCopyrightTextOpacity

`func (o *PortalCustomizeOpenApiVO) GetCopyrightTextOpacity() int32`

GetCopyrightTextOpacity returns the CopyrightTextOpacity field if non-nil, zero value otherwise.

### GetCopyrightTextOpacityOk

`func (o *PortalCustomizeOpenApiVO) GetCopyrightTextOpacityOk() (*int32, bool)`

GetCopyrightTextOpacityOk returns a tuple with the CopyrightTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightTextOpacity

`func (o *PortalCustomizeOpenApiVO) SetCopyrightTextOpacity(v int32)`

SetCopyrightTextOpacity sets CopyrightTextOpacity field to given value.

### HasCopyrightTextOpacity

`func (o *PortalCustomizeOpenApiVO) HasCopyrightTextOpacity() bool`

HasCopyrightTextOpacity returns a boolean if a field has been set.

### GetDefaultLanguage

`func (o *PortalCustomizeOpenApiVO) GetDefaultLanguage() int32`

GetDefaultLanguage returns the DefaultLanguage field if non-nil, zero value otherwise.

### GetDefaultLanguageOk

`func (o *PortalCustomizeOpenApiVO) GetDefaultLanguageOk() (*int32, bool)`

GetDefaultLanguageOk returns a tuple with the DefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguage

`func (o *PortalCustomizeOpenApiVO) SetDefaultLanguage(v int32)`

SetDefaultLanguage sets DefaultLanguage field to given value.


### GetFormAuthButtonText

`func (o *PortalCustomizeOpenApiVO) GetFormAuthButtonText() string`

GetFormAuthButtonText returns the FormAuthButtonText field if non-nil, zero value otherwise.

### GetFormAuthButtonTextOk

`func (o *PortalCustomizeOpenApiVO) GetFormAuthButtonTextOk() (*string, bool)`

GetFormAuthButtonTextOk returns a tuple with the FormAuthButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormAuthButtonText

`func (o *PortalCustomizeOpenApiVO) SetFormAuthButtonText(v string)`

SetFormAuthButtonText sets FormAuthButtonText field to given value.

### HasFormAuthButtonText

`func (o *PortalCustomizeOpenApiVO) HasFormAuthButtonText() bool`

HasFormAuthButtonText returns a boolean if a field has been set.

### GetInputBoxColor

`func (o *PortalCustomizeOpenApiVO) GetInputBoxColor() string`

GetInputBoxColor returns the InputBoxColor field if non-nil, zero value otherwise.

### GetInputBoxColorOk

`func (o *PortalCustomizeOpenApiVO) GetInputBoxColorOk() (*string, bool)`

GetInputBoxColorOk returns a tuple with the InputBoxColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBoxColor

`func (o *PortalCustomizeOpenApiVO) SetInputBoxColor(v string)`

SetInputBoxColor sets InputBoxColor field to given value.

### HasInputBoxColor

`func (o *PortalCustomizeOpenApiVO) HasInputBoxColor() bool`

HasInputBoxColor returns a boolean if a field has been set.

### GetInputBoxOpacity

`func (o *PortalCustomizeOpenApiVO) GetInputBoxOpacity() int32`

GetInputBoxOpacity returns the InputBoxOpacity field if non-nil, zero value otherwise.

### GetInputBoxOpacityOk

`func (o *PortalCustomizeOpenApiVO) GetInputBoxOpacityOk() (*int32, bool)`

GetInputBoxOpacityOk returns a tuple with the InputBoxOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBoxOpacity

`func (o *PortalCustomizeOpenApiVO) SetInputBoxOpacity(v int32)`

SetInputBoxOpacity sets InputBoxOpacity field to given value.

### HasInputBoxOpacity

`func (o *PortalCustomizeOpenApiVO) HasInputBoxOpacity() bool`

HasInputBoxOpacity returns a boolean if a field has been set.

### GetInputTextColor

`func (o *PortalCustomizeOpenApiVO) GetInputTextColor() string`

GetInputTextColor returns the InputTextColor field if non-nil, zero value otherwise.

### GetInputTextColorOk

`func (o *PortalCustomizeOpenApiVO) GetInputTextColorOk() (*string, bool)`

GetInputTextColorOk returns a tuple with the InputTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTextColor

`func (o *PortalCustomizeOpenApiVO) SetInputTextColor(v string)`

SetInputTextColor sets InputTextColor field to given value.

### HasInputTextColor

`func (o *PortalCustomizeOpenApiVO) HasInputTextColor() bool`

HasInputTextColor returns a boolean if a field has been set.

### GetInputTextOpacity

`func (o *PortalCustomizeOpenApiVO) GetInputTextOpacity() int32`

GetInputTextOpacity returns the InputTextOpacity field if non-nil, zero value otherwise.

### GetInputTextOpacityOk

`func (o *PortalCustomizeOpenApiVO) GetInputTextOpacityOk() (*int32, bool)`

GetInputTextOpacityOk returns a tuple with the InputTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTextOpacity

`func (o *PortalCustomizeOpenApiVO) SetInputTextOpacity(v int32)`

SetInputTextOpacity sets InputTextOpacity field to given value.

### HasInputTextOpacity

`func (o *PortalCustomizeOpenApiVO) HasInputTextOpacity() bool`

HasInputTextOpacity returns a boolean if a field has been set.

### GetLogoDisplay

`func (o *PortalCustomizeOpenApiVO) GetLogoDisplay() bool`

GetLogoDisplay returns the LogoDisplay field if non-nil, zero value otherwise.

### GetLogoDisplayOk

`func (o *PortalCustomizeOpenApiVO) GetLogoDisplayOk() (*bool, bool)`

GetLogoDisplayOk returns a tuple with the LogoDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDisplay

`func (o *PortalCustomizeOpenApiVO) SetLogoDisplay(v bool)`

SetLogoDisplay sets LogoDisplay field to given value.


### GetLogoPictureId

`func (o *PortalCustomizeOpenApiVO) GetLogoPictureId() string`

GetLogoPictureId returns the LogoPictureId field if non-nil, zero value otherwise.

### GetLogoPictureIdOk

`func (o *PortalCustomizeOpenApiVO) GetLogoPictureIdOk() (*string, bool)`

GetLogoPictureIdOk returns a tuple with the LogoPictureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPictureId

`func (o *PortalCustomizeOpenApiVO) SetLogoPictureId(v string)`

SetLogoPictureId sets LogoPictureId field to given value.

### HasLogoPictureId

`func (o *PortalCustomizeOpenApiVO) HasLogoPictureId() bool`

HasLogoPictureId returns a boolean if a field has been set.

### GetRedirectionCountDownEnable

`func (o *PortalCustomizeOpenApiVO) GetRedirectionCountDownEnable() bool`

GetRedirectionCountDownEnable returns the RedirectionCountDownEnable field if non-nil, zero value otherwise.

### GetRedirectionCountDownEnableOk

`func (o *PortalCustomizeOpenApiVO) GetRedirectionCountDownEnableOk() (*bool, bool)`

GetRedirectionCountDownEnableOk returns a tuple with the RedirectionCountDownEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionCountDownEnable

`func (o *PortalCustomizeOpenApiVO) SetRedirectionCountDownEnable(v bool)`

SetRedirectionCountDownEnable sets RedirectionCountDownEnable field to given value.

### HasRedirectionCountDownEnable

`func (o *PortalCustomizeOpenApiVO) HasRedirectionCountDownEnable() bool`

HasRedirectionCountDownEnable returns a boolean if a field has been set.

### GetTermsOfServiceEnable

`func (o *PortalCustomizeOpenApiVO) GetTermsOfServiceEnable() bool`

GetTermsOfServiceEnable returns the TermsOfServiceEnable field if non-nil, zero value otherwise.

### GetTermsOfServiceEnableOk

`func (o *PortalCustomizeOpenApiVO) GetTermsOfServiceEnableOk() (*bool, bool)`

GetTermsOfServiceEnableOk returns a tuple with the TermsOfServiceEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceEnable

`func (o *PortalCustomizeOpenApiVO) SetTermsOfServiceEnable(v bool)`

SetTermsOfServiceEnable sets TermsOfServiceEnable field to given value.


### GetTermsOfServiceFontSize

`func (o *PortalCustomizeOpenApiVO) GetTermsOfServiceFontSize() int32`

GetTermsOfServiceFontSize returns the TermsOfServiceFontSize field if non-nil, zero value otherwise.

### GetTermsOfServiceFontSizeOk

`func (o *PortalCustomizeOpenApiVO) GetTermsOfServiceFontSizeOk() (*int32, bool)`

GetTermsOfServiceFontSizeOk returns a tuple with the TermsOfServiceFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceFontSize

`func (o *PortalCustomizeOpenApiVO) SetTermsOfServiceFontSize(v int32)`

SetTermsOfServiceFontSize sets TermsOfServiceFontSize field to given value.

### HasTermsOfServiceFontSize

`func (o *PortalCustomizeOpenApiVO) HasTermsOfServiceFontSize() bool`

HasTermsOfServiceFontSize returns a boolean if a field has been set.

### GetTermsOfServiceText

`func (o *PortalCustomizeOpenApiVO) GetTermsOfServiceText() string`

GetTermsOfServiceText returns the TermsOfServiceText field if non-nil, zero value otherwise.

### GetTermsOfServiceTextOk

`func (o *PortalCustomizeOpenApiVO) GetTermsOfServiceTextOk() (*string, bool)`

GetTermsOfServiceTextOk returns a tuple with the TermsOfServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceText

`func (o *PortalCustomizeOpenApiVO) SetTermsOfServiceText(v string)`

SetTermsOfServiceText sets TermsOfServiceText field to given value.

### HasTermsOfServiceText

`func (o *PortalCustomizeOpenApiVO) HasTermsOfServiceText() bool`

HasTermsOfServiceText returns a boolean if a field has been set.

### GetTermsOfServiceUrlTexts

`func (o *PortalCustomizeOpenApiVO) GetTermsOfServiceUrlTexts() []TermsOfServiceUrlVO`

GetTermsOfServiceUrlTexts returns the TermsOfServiceUrlTexts field if non-nil, zero value otherwise.

### GetTermsOfServiceUrlTextsOk

`func (o *PortalCustomizeOpenApiVO) GetTermsOfServiceUrlTextsOk() (*[]TermsOfServiceUrlVO, bool)`

GetTermsOfServiceUrlTextsOk returns a tuple with the TermsOfServiceUrlTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceUrlTexts

`func (o *PortalCustomizeOpenApiVO) SetTermsOfServiceUrlTexts(v []TermsOfServiceUrlVO)`

SetTermsOfServiceUrlTexts sets TermsOfServiceUrlTexts field to given value.

### HasTermsOfServiceUrlTexts

`func (o *PortalCustomizeOpenApiVO) HasTermsOfServiceUrlTexts() bool`

HasTermsOfServiceUrlTexts returns a boolean if a field has been set.

### GetWelcomeEnable

`func (o *PortalCustomizeOpenApiVO) GetWelcomeEnable() bool`

GetWelcomeEnable returns the WelcomeEnable field if non-nil, zero value otherwise.

### GetWelcomeEnableOk

`func (o *PortalCustomizeOpenApiVO) GetWelcomeEnableOk() (*bool, bool)`

GetWelcomeEnableOk returns a tuple with the WelcomeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeEnable

`func (o *PortalCustomizeOpenApiVO) SetWelcomeEnable(v bool)`

SetWelcomeEnable sets WelcomeEnable field to given value.


### GetWelcomeInformation

`func (o *PortalCustomizeOpenApiVO) GetWelcomeInformation() string`

GetWelcomeInformation returns the WelcomeInformation field if non-nil, zero value otherwise.

### GetWelcomeInformationOk

`func (o *PortalCustomizeOpenApiVO) GetWelcomeInformationOk() (*string, bool)`

GetWelcomeInformationOk returns a tuple with the WelcomeInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeInformation

`func (o *PortalCustomizeOpenApiVO) SetWelcomeInformation(v string)`

SetWelcomeInformation sets WelcomeInformation field to given value.

### HasWelcomeInformation

`func (o *PortalCustomizeOpenApiVO) HasWelcomeInformation() bool`

HasWelcomeInformation returns a boolean if a field has been set.

### GetWelcomeTextColor

`func (o *PortalCustomizeOpenApiVO) GetWelcomeTextColor() string`

GetWelcomeTextColor returns the WelcomeTextColor field if non-nil, zero value otherwise.

### GetWelcomeTextColorOk

`func (o *PortalCustomizeOpenApiVO) GetWelcomeTextColorOk() (*string, bool)`

GetWelcomeTextColorOk returns a tuple with the WelcomeTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTextColor

`func (o *PortalCustomizeOpenApiVO) SetWelcomeTextColor(v string)`

SetWelcomeTextColor sets WelcomeTextColor field to given value.

### HasWelcomeTextColor

`func (o *PortalCustomizeOpenApiVO) HasWelcomeTextColor() bool`

HasWelcomeTextColor returns a boolean if a field has been set.

### GetWelcomeTextFontSize

`func (o *PortalCustomizeOpenApiVO) GetWelcomeTextFontSize() int32`

GetWelcomeTextFontSize returns the WelcomeTextFontSize field if non-nil, zero value otherwise.

### GetWelcomeTextFontSizeOk

`func (o *PortalCustomizeOpenApiVO) GetWelcomeTextFontSizeOk() (*int32, bool)`

GetWelcomeTextFontSizeOk returns a tuple with the WelcomeTextFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTextFontSize

`func (o *PortalCustomizeOpenApiVO) SetWelcomeTextFontSize(v int32)`

SetWelcomeTextFontSize sets WelcomeTextFontSize field to given value.

### HasWelcomeTextFontSize

`func (o *PortalCustomizeOpenApiVO) HasWelcomeTextFontSize() bool`

HasWelcomeTextFontSize returns a boolean if a field has been set.

### GetWelcomeTextOpacity

`func (o *PortalCustomizeOpenApiVO) GetWelcomeTextOpacity() int32`

GetWelcomeTextOpacity returns the WelcomeTextOpacity field if non-nil, zero value otherwise.

### GetWelcomeTextOpacityOk

`func (o *PortalCustomizeOpenApiVO) GetWelcomeTextOpacityOk() (*int32, bool)`

GetWelcomeTextOpacityOk returns a tuple with the WelcomeTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTextOpacity

`func (o *PortalCustomizeOpenApiVO) SetWelcomeTextOpacity(v int32)`

SetWelcomeTextOpacity sets WelcomeTextOpacity field to given value.

### HasWelcomeTextOpacity

`func (o *PortalCustomizeOpenApiVO) HasWelcomeTextOpacity() bool`

HasWelcomeTextOpacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


