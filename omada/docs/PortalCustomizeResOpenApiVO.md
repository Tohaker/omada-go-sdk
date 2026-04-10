# PortalCustomizeResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Advertisement** | Pointer to [**AdvertisementSettingResOpenApiVO**](AdvertisementSettingResOpenApiVO.md) |  | [optional] 
**Background** | Pointer to **int32** | Background type, should be a value as follows: 0:library, 1: solid color, 2: picture | [optional] 
**BackgroundColor** | Pointer to **string** | Portal background picture color. Hex color code such as: #ffffff | [optional] 
**BackgroundMaskColor** | Pointer to **string** | Background mask color. Hex color code such as: #ffffff. | [optional] 
**BackgroundMaskEnable** | Pointer to **bool** | Whether to enable multiple language. | [optional] 
**BackgroundMaskOpacity** | Pointer to **int32** | Background mask opacity, should be within the range of 0–100. | [optional] 
**BackgroundOpacity** | Pointer to **int32** | Portal background picture transparency percentage, should be within the range of 0–100 | [optional] 
**BackgroundPicture** | Pointer to [**PortalPictureInfo**](PortalPictureInfo.md) |  | [optional] 
**BackgroundPictureIndex** | Pointer to **int32** | Index of library background picture, should be within the range of 0-5. | [optional] 
**BgPicCoordinatesOfLibrary** | Pointer to [**BgPicCoordinatesOfLibraryOpenApiVO**](BgPicCoordinatesOfLibraryOpenApiVO.md) |  | [optional] 
**BodyContainerBgBlur** | Pointer to **int32** | Body container background blurriness, should be within the range of 0–10. | [optional] 
**BodyContainerBgBlurEnable** | Pointer to **bool** | Whether to enable body container background blur. | [optional] 
**BodyContainerColor** | Pointer to **string** | Body container color. Hex color code such as: #ffffff. | [optional] 
**BodyContainerEnable** | Pointer to **bool** | Whether to enable body container. | [optional] 
**BodyContainerOpacity** | Pointer to **int32** | Body container opacity, should be within the range of 0–100. | [optional] 
**BodyContainerRadius** | Pointer to **int32** | Body container radius, should be within the range of 0–30. | [optional] 
**BodyContainerType** | Pointer to **int32** | Type of body container, 0: none; 1: half; 2: all | [optional] 
**ButtonColor** | Pointer to **string** | Button color. Hex color code such as: #ffffff | [optional] 
**ButtonOpacity** | Pointer to **int32** | Button opacity, should be within the range of 0–100 | [optional] 
**ButtonPositionRatio** | Pointer to **float32** | Button position ratio | [optional] 
**ButtonRadius** | Pointer to **int32** | Button radius, should be within the range of 0–30. | [optional] 
**ButtonText** | Pointer to **string** | Button text, should contain 0 to 32 characters, default value is \&quot;Log In\&quot; | [optional] 
**ButtonTextColor** | Pointer to **string** | Button text color. Hex color code such as: #ffffff | [optional] 
**ButtonTextOpacity** | Pointer to **int32** | Button text opacity, should be within the range of 0–100 | [optional] 
**ButtonTranslate** | Pointer to **int32** | The distance between the position of the button and the top | [optional] 
**Copyright** | Pointer to **string** | Copyright text, should contain 0 to 200 characters | [optional] 
**CopyrightEnable** | Pointer to **bool** | Whether to display the copyright | [optional] 
**CopyrightTextColor** | Pointer to **string** | Copyright text color. Hex color code such as: #ffffff | [optional] 
**CopyrightTextFontSize** | Pointer to **int32** | Copyright text font size, should be within the range of 12–18 | [optional] 
**CopyrightTextOpacity** | Pointer to **int32** | Copyright text opacity, should be within the range of 0–100 | [optional] 
**DefaultLanguage** | Pointer to **int32** | The controller automatically adjusts the language displayed on the Portal page according to the system language of the clients. If the language is not supported, the controller will use the default language specified here. DefaultLanguage should be a value as follows:&lt;br/&gt;1: en_US (English); 3: cs_CZ (Český); 4: de_DE (Deutsch); 5: da_DK (Dansk); 6: el_GR (ελληνικά);&lt;br/&gt;7: fr_FR (Français); 8: es_ES (Español); 9: nl_NL (Nederlands); 10: it_IT (Italiano); 11: pl_PL (Polski);&lt;br/&gt;12: pt_PT (Português); 13: ru_RU (Русский); 14: sv_SE (Svenska); 15: tr_TR (Türkçe);&lt;br/&gt;16: ar_SA (لغة عربية); &lt;br/&gt;17: ja_JP (日本語); 18: zh_TW (中文(繁體)); 19: th_TH (ไทย); 20: vi_VN (Tiếng Việt); 21: ko_KR (한국어) | [optional] 
**DescriptionText** | Pointer to **string** | Description text, should contain 0 to 256 characters. | [optional] 
**DescriptionTextColor** | Pointer to **string** | Description text color. Hex color code such as: #ffffff. | [optional] 
**DescriptionTextFontSize** | Pointer to **int32** | Description text font size, should be within the range of 12–18. | [optional] 
**DescriptionTextOpacity** | Pointer to **int32** | Description text opacity, should be within the range of 0–100. | [optional] 
**EnableDeviceSpecificBg** | Pointer to **bool** | Whether to use different images in mobile and PC devices | [optional] 
**FormAuthButtonText** | Pointer to **string** | Form auth button text, should contain 0 to 32 characters. Required when [authType] is 11 and hotspot [enabledTypes] contains 12. Default value is \&quot;Take the Survey\&quot; | [optional] 
**InputBoxBorderColor** | Pointer to **string** | Input box border color. Hex color code such as: #ffffff. | [optional] 
**InputBoxBorderOpacity** | Pointer to **int32** | Input box border opacity, should be within the range of 0–100. | [optional] 
**InputBoxColor** | Pointer to **string** | Input box color. Hex color code such as: #ffffff | [optional] 
**InputBoxOpacity** | Pointer to **int32** | Input box opacity, should be within the range of 0–100 | [optional] 
**InputBoxRadius** | Pointer to **int32** | Input box radius, should be within the range of 0–30. | [optional] 
**InputTextColor** | Pointer to **string** | Input text color. Hex color code such as: #ffffff | [optional] 
**InputTextOpacity** | Pointer to **int32** | Input text opacity, should be within the range of 0–100 | [optional] 
**LogoDisplay** | Pointer to **bool** | Whether to display the default logo | [optional] 
**LogoHorizontalPosition** | Pointer to **int32** | Position of logo horizontal, 0: left; 1: medium; 2: right | [optional] 
**LogoPicture** | Pointer to [**PortalPictureInfo**](PortalPictureInfo.md) |  | [optional] 
**LogoPosition** | Pointer to **int32** | Logo position, should be a value as follow: 1: up; 2: middle; 3: lower | [optional] 
**LogoPositionRatio** | Pointer to **float32** | Logo position ratio, should be within the range of 1-100, represents from high to low. | [optional] 
**LogoSize** | Pointer to **int32** | Logo size, should be within the range of 1-300 | [optional] 
**LogoTranslate** | Pointer to **int32** | Distance of the logo from the top of the container, range: above 0 | [optional] 
**MobileBackgroundPicture** | Pointer to [**PortalPictureInfo**](PortalPictureInfo.md) |  | [optional] 
**MobileBgPicCoordinatesOfLibrary** | Pointer to [**BgPicCoordinatesOfLibraryOpenApiVO**](BgPicCoordinatesOfLibraryOpenApiVO.md) |  | [optional] 
**PcAlign** | Pointer to **int32** | Position of pc align, 0: left; 1: medium; 2: right | [optional] 
**RedirectionCountDownEnable** | Pointer to **bool** | Whether to show redirection countdown after authorized | [optional] 
**TermsOfService** | Pointer to **string** | Service Terms Content | [optional] 
**TermsOfServiceEnable** | Pointer to **bool** | Whether to display terms of service | [optional] 
**TermsOfServiceFontSize** | Pointer to **int32** | Terms of service text font size, should be within the range of 12–18 | [optional] 
**TermsOfServiceText** | Pointer to **string** | Terms of service text, should contain 0 to 100 characters | [optional] 
**TermsOfServiceTextColor** | Pointer to **string** | Terms of service text color. Hex color code such as: #ffffff. | [optional] 
**TermsOfServiceTextOpacity** | Pointer to **int32** | Terms of service text opacity, should be within the range of 0–100. | [optional] 
**TermsOfServiceUrlTexts** | Pointer to [**[]TermsOfServiceUrlVO**](TermsOfServiceUrlVO.md) | Terms of service URL texts, match the termsOfServiceText and turn the matching characters into an openable link. Up to 3 entries are allowed for the list | [optional] 
**WelcomeEnable** | Pointer to **bool** | Whether to display the welcome info | [optional] 
**WelcomeInformation** | Pointer to **string** | Welcome information, should contain 1 to 31 characters | [optional] 
**WelcomeTextColor** | Pointer to **string** | Welcome text color. Hex color code such as: #ffffff | [optional] 
**WelcomeTextFontSize** | Pointer to **int32** | Welcome text font size, should be within the range of 12–18 | [optional] 
**WelcomeTextOpacity** | Pointer to **int32** | Welcome text opacity, should be within the range of 0–100 | [optional] 

## Methods

### NewPortalCustomizeResOpenApiVO

`func NewPortalCustomizeResOpenApiVO() *PortalCustomizeResOpenApiVO`

NewPortalCustomizeResOpenApiVO instantiates a new PortalCustomizeResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalCustomizeResOpenApiVOWithDefaults

`func NewPortalCustomizeResOpenApiVOWithDefaults() *PortalCustomizeResOpenApiVO`

NewPortalCustomizeResOpenApiVOWithDefaults instantiates a new PortalCustomizeResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertisement

`func (o *PortalCustomizeResOpenApiVO) GetAdvertisement() AdvertisementSettingResOpenApiVO`

GetAdvertisement returns the Advertisement field if non-nil, zero value otherwise.

### GetAdvertisementOk

`func (o *PortalCustomizeResOpenApiVO) GetAdvertisementOk() (*AdvertisementSettingResOpenApiVO, bool)`

GetAdvertisementOk returns a tuple with the Advertisement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisement

`func (o *PortalCustomizeResOpenApiVO) SetAdvertisement(v AdvertisementSettingResOpenApiVO)`

SetAdvertisement sets Advertisement field to given value.

### HasAdvertisement

`func (o *PortalCustomizeResOpenApiVO) HasAdvertisement() bool`

HasAdvertisement returns a boolean if a field has been set.

### GetBackground

`func (o *PortalCustomizeResOpenApiVO) GetBackground() int32`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundOk() (*int32, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *PortalCustomizeResOpenApiVO) SetBackground(v int32)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *PortalCustomizeResOpenApiVO) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *PortalCustomizeResOpenApiVO) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *PortalCustomizeResOpenApiVO) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetBackgroundMaskColor

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundMaskColor() string`

GetBackgroundMaskColor returns the BackgroundMaskColor field if non-nil, zero value otherwise.

### GetBackgroundMaskColorOk

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundMaskColorOk() (*string, bool)`

GetBackgroundMaskColorOk returns a tuple with the BackgroundMaskColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMaskColor

`func (o *PortalCustomizeResOpenApiVO) SetBackgroundMaskColor(v string)`

SetBackgroundMaskColor sets BackgroundMaskColor field to given value.

### HasBackgroundMaskColor

`func (o *PortalCustomizeResOpenApiVO) HasBackgroundMaskColor() bool`

HasBackgroundMaskColor returns a boolean if a field has been set.

### GetBackgroundMaskEnable

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundMaskEnable() bool`

GetBackgroundMaskEnable returns the BackgroundMaskEnable field if non-nil, zero value otherwise.

### GetBackgroundMaskEnableOk

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundMaskEnableOk() (*bool, bool)`

GetBackgroundMaskEnableOk returns a tuple with the BackgroundMaskEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMaskEnable

`func (o *PortalCustomizeResOpenApiVO) SetBackgroundMaskEnable(v bool)`

SetBackgroundMaskEnable sets BackgroundMaskEnable field to given value.

### HasBackgroundMaskEnable

`func (o *PortalCustomizeResOpenApiVO) HasBackgroundMaskEnable() bool`

HasBackgroundMaskEnable returns a boolean if a field has been set.

### GetBackgroundMaskOpacity

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundMaskOpacity() int32`

GetBackgroundMaskOpacity returns the BackgroundMaskOpacity field if non-nil, zero value otherwise.

### GetBackgroundMaskOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundMaskOpacityOk() (*int32, bool)`

GetBackgroundMaskOpacityOk returns a tuple with the BackgroundMaskOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundMaskOpacity

`func (o *PortalCustomizeResOpenApiVO) SetBackgroundMaskOpacity(v int32)`

SetBackgroundMaskOpacity sets BackgroundMaskOpacity field to given value.

### HasBackgroundMaskOpacity

`func (o *PortalCustomizeResOpenApiVO) HasBackgroundMaskOpacity() bool`

HasBackgroundMaskOpacity returns a boolean if a field has been set.

### GetBackgroundOpacity

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundOpacity() int32`

GetBackgroundOpacity returns the BackgroundOpacity field if non-nil, zero value otherwise.

### GetBackgroundOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundOpacityOk() (*int32, bool)`

GetBackgroundOpacityOk returns a tuple with the BackgroundOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundOpacity

`func (o *PortalCustomizeResOpenApiVO) SetBackgroundOpacity(v int32)`

SetBackgroundOpacity sets BackgroundOpacity field to given value.

### HasBackgroundOpacity

`func (o *PortalCustomizeResOpenApiVO) HasBackgroundOpacity() bool`

HasBackgroundOpacity returns a boolean if a field has been set.

### GetBackgroundPicture

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundPicture() PortalPictureInfo`

GetBackgroundPicture returns the BackgroundPicture field if non-nil, zero value otherwise.

### GetBackgroundPictureOk

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundPictureOk() (*PortalPictureInfo, bool)`

GetBackgroundPictureOk returns a tuple with the BackgroundPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPicture

`func (o *PortalCustomizeResOpenApiVO) SetBackgroundPicture(v PortalPictureInfo)`

SetBackgroundPicture sets BackgroundPicture field to given value.

### HasBackgroundPicture

`func (o *PortalCustomizeResOpenApiVO) HasBackgroundPicture() bool`

HasBackgroundPicture returns a boolean if a field has been set.

### GetBackgroundPictureIndex

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundPictureIndex() int32`

GetBackgroundPictureIndex returns the BackgroundPictureIndex field if non-nil, zero value otherwise.

### GetBackgroundPictureIndexOk

`func (o *PortalCustomizeResOpenApiVO) GetBackgroundPictureIndexOk() (*int32, bool)`

GetBackgroundPictureIndexOk returns a tuple with the BackgroundPictureIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundPictureIndex

`func (o *PortalCustomizeResOpenApiVO) SetBackgroundPictureIndex(v int32)`

SetBackgroundPictureIndex sets BackgroundPictureIndex field to given value.

### HasBackgroundPictureIndex

`func (o *PortalCustomizeResOpenApiVO) HasBackgroundPictureIndex() bool`

HasBackgroundPictureIndex returns a boolean if a field has been set.

### GetBgPicCoordinatesOfLibrary

`func (o *PortalCustomizeResOpenApiVO) GetBgPicCoordinatesOfLibrary() BgPicCoordinatesOfLibraryOpenApiVO`

GetBgPicCoordinatesOfLibrary returns the BgPicCoordinatesOfLibrary field if non-nil, zero value otherwise.

### GetBgPicCoordinatesOfLibraryOk

`func (o *PortalCustomizeResOpenApiVO) GetBgPicCoordinatesOfLibraryOk() (*BgPicCoordinatesOfLibraryOpenApiVO, bool)`

GetBgPicCoordinatesOfLibraryOk returns a tuple with the BgPicCoordinatesOfLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgPicCoordinatesOfLibrary

`func (o *PortalCustomizeResOpenApiVO) SetBgPicCoordinatesOfLibrary(v BgPicCoordinatesOfLibraryOpenApiVO)`

SetBgPicCoordinatesOfLibrary sets BgPicCoordinatesOfLibrary field to given value.

### HasBgPicCoordinatesOfLibrary

`func (o *PortalCustomizeResOpenApiVO) HasBgPicCoordinatesOfLibrary() bool`

HasBgPicCoordinatesOfLibrary returns a boolean if a field has been set.

### GetBodyContainerBgBlur

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerBgBlur() int32`

GetBodyContainerBgBlur returns the BodyContainerBgBlur field if non-nil, zero value otherwise.

### GetBodyContainerBgBlurOk

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerBgBlurOk() (*int32, bool)`

GetBodyContainerBgBlurOk returns a tuple with the BodyContainerBgBlur field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContainerBgBlur

`func (o *PortalCustomizeResOpenApiVO) SetBodyContainerBgBlur(v int32)`

SetBodyContainerBgBlur sets BodyContainerBgBlur field to given value.

### HasBodyContainerBgBlur

`func (o *PortalCustomizeResOpenApiVO) HasBodyContainerBgBlur() bool`

HasBodyContainerBgBlur returns a boolean if a field has been set.

### GetBodyContainerBgBlurEnable

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerBgBlurEnable() bool`

GetBodyContainerBgBlurEnable returns the BodyContainerBgBlurEnable field if non-nil, zero value otherwise.

### GetBodyContainerBgBlurEnableOk

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerBgBlurEnableOk() (*bool, bool)`

GetBodyContainerBgBlurEnableOk returns a tuple with the BodyContainerBgBlurEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContainerBgBlurEnable

`func (o *PortalCustomizeResOpenApiVO) SetBodyContainerBgBlurEnable(v bool)`

SetBodyContainerBgBlurEnable sets BodyContainerBgBlurEnable field to given value.

### HasBodyContainerBgBlurEnable

`func (o *PortalCustomizeResOpenApiVO) HasBodyContainerBgBlurEnable() bool`

HasBodyContainerBgBlurEnable returns a boolean if a field has been set.

### GetBodyContainerColor

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerColor() string`

GetBodyContainerColor returns the BodyContainerColor field if non-nil, zero value otherwise.

### GetBodyContainerColorOk

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerColorOk() (*string, bool)`

GetBodyContainerColorOk returns a tuple with the BodyContainerColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContainerColor

`func (o *PortalCustomizeResOpenApiVO) SetBodyContainerColor(v string)`

SetBodyContainerColor sets BodyContainerColor field to given value.

### HasBodyContainerColor

`func (o *PortalCustomizeResOpenApiVO) HasBodyContainerColor() bool`

HasBodyContainerColor returns a boolean if a field has been set.

### GetBodyContainerEnable

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerEnable() bool`

GetBodyContainerEnable returns the BodyContainerEnable field if non-nil, zero value otherwise.

### GetBodyContainerEnableOk

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerEnableOk() (*bool, bool)`

GetBodyContainerEnableOk returns a tuple with the BodyContainerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContainerEnable

`func (o *PortalCustomizeResOpenApiVO) SetBodyContainerEnable(v bool)`

SetBodyContainerEnable sets BodyContainerEnable field to given value.

### HasBodyContainerEnable

`func (o *PortalCustomizeResOpenApiVO) HasBodyContainerEnable() bool`

HasBodyContainerEnable returns a boolean if a field has been set.

### GetBodyContainerOpacity

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerOpacity() int32`

GetBodyContainerOpacity returns the BodyContainerOpacity field if non-nil, zero value otherwise.

### GetBodyContainerOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerOpacityOk() (*int32, bool)`

GetBodyContainerOpacityOk returns a tuple with the BodyContainerOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContainerOpacity

`func (o *PortalCustomizeResOpenApiVO) SetBodyContainerOpacity(v int32)`

SetBodyContainerOpacity sets BodyContainerOpacity field to given value.

### HasBodyContainerOpacity

`func (o *PortalCustomizeResOpenApiVO) HasBodyContainerOpacity() bool`

HasBodyContainerOpacity returns a boolean if a field has been set.

### GetBodyContainerRadius

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerRadius() int32`

GetBodyContainerRadius returns the BodyContainerRadius field if non-nil, zero value otherwise.

### GetBodyContainerRadiusOk

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerRadiusOk() (*int32, bool)`

GetBodyContainerRadiusOk returns a tuple with the BodyContainerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContainerRadius

`func (o *PortalCustomizeResOpenApiVO) SetBodyContainerRadius(v int32)`

SetBodyContainerRadius sets BodyContainerRadius field to given value.

### HasBodyContainerRadius

`func (o *PortalCustomizeResOpenApiVO) HasBodyContainerRadius() bool`

HasBodyContainerRadius returns a boolean if a field has been set.

### GetBodyContainerType

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerType() int32`

GetBodyContainerType returns the BodyContainerType field if non-nil, zero value otherwise.

### GetBodyContainerTypeOk

`func (o *PortalCustomizeResOpenApiVO) GetBodyContainerTypeOk() (*int32, bool)`

GetBodyContainerTypeOk returns a tuple with the BodyContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyContainerType

`func (o *PortalCustomizeResOpenApiVO) SetBodyContainerType(v int32)`

SetBodyContainerType sets BodyContainerType field to given value.

### HasBodyContainerType

`func (o *PortalCustomizeResOpenApiVO) HasBodyContainerType() bool`

HasBodyContainerType returns a boolean if a field has been set.

### GetButtonColor

`func (o *PortalCustomizeResOpenApiVO) GetButtonColor() string`

GetButtonColor returns the ButtonColor field if non-nil, zero value otherwise.

### GetButtonColorOk

`func (o *PortalCustomizeResOpenApiVO) GetButtonColorOk() (*string, bool)`

GetButtonColorOk returns a tuple with the ButtonColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonColor

`func (o *PortalCustomizeResOpenApiVO) SetButtonColor(v string)`

SetButtonColor sets ButtonColor field to given value.

### HasButtonColor

`func (o *PortalCustomizeResOpenApiVO) HasButtonColor() bool`

HasButtonColor returns a boolean if a field has been set.

### GetButtonOpacity

`func (o *PortalCustomizeResOpenApiVO) GetButtonOpacity() int32`

GetButtonOpacity returns the ButtonOpacity field if non-nil, zero value otherwise.

### GetButtonOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetButtonOpacityOk() (*int32, bool)`

GetButtonOpacityOk returns a tuple with the ButtonOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonOpacity

`func (o *PortalCustomizeResOpenApiVO) SetButtonOpacity(v int32)`

SetButtonOpacity sets ButtonOpacity field to given value.

### HasButtonOpacity

`func (o *PortalCustomizeResOpenApiVO) HasButtonOpacity() bool`

HasButtonOpacity returns a boolean if a field has been set.

### GetButtonPositionRatio

`func (o *PortalCustomizeResOpenApiVO) GetButtonPositionRatio() float32`

GetButtonPositionRatio returns the ButtonPositionRatio field if non-nil, zero value otherwise.

### GetButtonPositionRatioOk

`func (o *PortalCustomizeResOpenApiVO) GetButtonPositionRatioOk() (*float32, bool)`

GetButtonPositionRatioOk returns a tuple with the ButtonPositionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonPositionRatio

`func (o *PortalCustomizeResOpenApiVO) SetButtonPositionRatio(v float32)`

SetButtonPositionRatio sets ButtonPositionRatio field to given value.

### HasButtonPositionRatio

`func (o *PortalCustomizeResOpenApiVO) HasButtonPositionRatio() bool`

HasButtonPositionRatio returns a boolean if a field has been set.

### GetButtonRadius

`func (o *PortalCustomizeResOpenApiVO) GetButtonRadius() int32`

GetButtonRadius returns the ButtonRadius field if non-nil, zero value otherwise.

### GetButtonRadiusOk

`func (o *PortalCustomizeResOpenApiVO) GetButtonRadiusOk() (*int32, bool)`

GetButtonRadiusOk returns a tuple with the ButtonRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonRadius

`func (o *PortalCustomizeResOpenApiVO) SetButtonRadius(v int32)`

SetButtonRadius sets ButtonRadius field to given value.

### HasButtonRadius

`func (o *PortalCustomizeResOpenApiVO) HasButtonRadius() bool`

HasButtonRadius returns a boolean if a field has been set.

### GetButtonText

`func (o *PortalCustomizeResOpenApiVO) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *PortalCustomizeResOpenApiVO) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *PortalCustomizeResOpenApiVO) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *PortalCustomizeResOpenApiVO) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### GetButtonTextColor

`func (o *PortalCustomizeResOpenApiVO) GetButtonTextColor() string`

GetButtonTextColor returns the ButtonTextColor field if non-nil, zero value otherwise.

### GetButtonTextColorOk

`func (o *PortalCustomizeResOpenApiVO) GetButtonTextColorOk() (*string, bool)`

GetButtonTextColorOk returns a tuple with the ButtonTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTextColor

`func (o *PortalCustomizeResOpenApiVO) SetButtonTextColor(v string)`

SetButtonTextColor sets ButtonTextColor field to given value.

### HasButtonTextColor

`func (o *PortalCustomizeResOpenApiVO) HasButtonTextColor() bool`

HasButtonTextColor returns a boolean if a field has been set.

### GetButtonTextOpacity

`func (o *PortalCustomizeResOpenApiVO) GetButtonTextOpacity() int32`

GetButtonTextOpacity returns the ButtonTextOpacity field if non-nil, zero value otherwise.

### GetButtonTextOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetButtonTextOpacityOk() (*int32, bool)`

GetButtonTextOpacityOk returns a tuple with the ButtonTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTextOpacity

`func (o *PortalCustomizeResOpenApiVO) SetButtonTextOpacity(v int32)`

SetButtonTextOpacity sets ButtonTextOpacity field to given value.

### HasButtonTextOpacity

`func (o *PortalCustomizeResOpenApiVO) HasButtonTextOpacity() bool`

HasButtonTextOpacity returns a boolean if a field has been set.

### GetButtonTranslate

`func (o *PortalCustomizeResOpenApiVO) GetButtonTranslate() int32`

GetButtonTranslate returns the ButtonTranslate field if non-nil, zero value otherwise.

### GetButtonTranslateOk

`func (o *PortalCustomizeResOpenApiVO) GetButtonTranslateOk() (*int32, bool)`

GetButtonTranslateOk returns a tuple with the ButtonTranslate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTranslate

`func (o *PortalCustomizeResOpenApiVO) SetButtonTranslate(v int32)`

SetButtonTranslate sets ButtonTranslate field to given value.

### HasButtonTranslate

`func (o *PortalCustomizeResOpenApiVO) HasButtonTranslate() bool`

HasButtonTranslate returns a boolean if a field has been set.

### GetCopyright

`func (o *PortalCustomizeResOpenApiVO) GetCopyright() string`

GetCopyright returns the Copyright field if non-nil, zero value otherwise.

### GetCopyrightOk

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightOk() (*string, bool)`

GetCopyrightOk returns a tuple with the Copyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyright

`func (o *PortalCustomizeResOpenApiVO) SetCopyright(v string)`

SetCopyright sets Copyright field to given value.

### HasCopyright

`func (o *PortalCustomizeResOpenApiVO) HasCopyright() bool`

HasCopyright returns a boolean if a field has been set.

### GetCopyrightEnable

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightEnable() bool`

GetCopyrightEnable returns the CopyrightEnable field if non-nil, zero value otherwise.

### GetCopyrightEnableOk

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightEnableOk() (*bool, bool)`

GetCopyrightEnableOk returns a tuple with the CopyrightEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightEnable

`func (o *PortalCustomizeResOpenApiVO) SetCopyrightEnable(v bool)`

SetCopyrightEnable sets CopyrightEnable field to given value.

### HasCopyrightEnable

`func (o *PortalCustomizeResOpenApiVO) HasCopyrightEnable() bool`

HasCopyrightEnable returns a boolean if a field has been set.

### GetCopyrightTextColor

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightTextColor() string`

GetCopyrightTextColor returns the CopyrightTextColor field if non-nil, zero value otherwise.

### GetCopyrightTextColorOk

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightTextColorOk() (*string, bool)`

GetCopyrightTextColorOk returns a tuple with the CopyrightTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightTextColor

`func (o *PortalCustomizeResOpenApiVO) SetCopyrightTextColor(v string)`

SetCopyrightTextColor sets CopyrightTextColor field to given value.

### HasCopyrightTextColor

`func (o *PortalCustomizeResOpenApiVO) HasCopyrightTextColor() bool`

HasCopyrightTextColor returns a boolean if a field has been set.

### GetCopyrightTextFontSize

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightTextFontSize() int32`

GetCopyrightTextFontSize returns the CopyrightTextFontSize field if non-nil, zero value otherwise.

### GetCopyrightTextFontSizeOk

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightTextFontSizeOk() (*int32, bool)`

GetCopyrightTextFontSizeOk returns a tuple with the CopyrightTextFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightTextFontSize

`func (o *PortalCustomizeResOpenApiVO) SetCopyrightTextFontSize(v int32)`

SetCopyrightTextFontSize sets CopyrightTextFontSize field to given value.

### HasCopyrightTextFontSize

`func (o *PortalCustomizeResOpenApiVO) HasCopyrightTextFontSize() bool`

HasCopyrightTextFontSize returns a boolean if a field has been set.

### GetCopyrightTextOpacity

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightTextOpacity() int32`

GetCopyrightTextOpacity returns the CopyrightTextOpacity field if non-nil, zero value otherwise.

### GetCopyrightTextOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetCopyrightTextOpacityOk() (*int32, bool)`

GetCopyrightTextOpacityOk returns a tuple with the CopyrightTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightTextOpacity

`func (o *PortalCustomizeResOpenApiVO) SetCopyrightTextOpacity(v int32)`

SetCopyrightTextOpacity sets CopyrightTextOpacity field to given value.

### HasCopyrightTextOpacity

`func (o *PortalCustomizeResOpenApiVO) HasCopyrightTextOpacity() bool`

HasCopyrightTextOpacity returns a boolean if a field has been set.

### GetDefaultLanguage

`func (o *PortalCustomizeResOpenApiVO) GetDefaultLanguage() int32`

GetDefaultLanguage returns the DefaultLanguage field if non-nil, zero value otherwise.

### GetDefaultLanguageOk

`func (o *PortalCustomizeResOpenApiVO) GetDefaultLanguageOk() (*int32, bool)`

GetDefaultLanguageOk returns a tuple with the DefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguage

`func (o *PortalCustomizeResOpenApiVO) SetDefaultLanguage(v int32)`

SetDefaultLanguage sets DefaultLanguage field to given value.

### HasDefaultLanguage

`func (o *PortalCustomizeResOpenApiVO) HasDefaultLanguage() bool`

HasDefaultLanguage returns a boolean if a field has been set.

### GetDescriptionText

`func (o *PortalCustomizeResOpenApiVO) GetDescriptionText() string`

GetDescriptionText returns the DescriptionText field if non-nil, zero value otherwise.

### GetDescriptionTextOk

`func (o *PortalCustomizeResOpenApiVO) GetDescriptionTextOk() (*string, bool)`

GetDescriptionTextOk returns a tuple with the DescriptionText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionText

`func (o *PortalCustomizeResOpenApiVO) SetDescriptionText(v string)`

SetDescriptionText sets DescriptionText field to given value.

### HasDescriptionText

`func (o *PortalCustomizeResOpenApiVO) HasDescriptionText() bool`

HasDescriptionText returns a boolean if a field has been set.

### GetDescriptionTextColor

`func (o *PortalCustomizeResOpenApiVO) GetDescriptionTextColor() string`

GetDescriptionTextColor returns the DescriptionTextColor field if non-nil, zero value otherwise.

### GetDescriptionTextColorOk

`func (o *PortalCustomizeResOpenApiVO) GetDescriptionTextColorOk() (*string, bool)`

GetDescriptionTextColorOk returns a tuple with the DescriptionTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionTextColor

`func (o *PortalCustomizeResOpenApiVO) SetDescriptionTextColor(v string)`

SetDescriptionTextColor sets DescriptionTextColor field to given value.

### HasDescriptionTextColor

`func (o *PortalCustomizeResOpenApiVO) HasDescriptionTextColor() bool`

HasDescriptionTextColor returns a boolean if a field has been set.

### GetDescriptionTextFontSize

`func (o *PortalCustomizeResOpenApiVO) GetDescriptionTextFontSize() int32`

GetDescriptionTextFontSize returns the DescriptionTextFontSize field if non-nil, zero value otherwise.

### GetDescriptionTextFontSizeOk

`func (o *PortalCustomizeResOpenApiVO) GetDescriptionTextFontSizeOk() (*int32, bool)`

GetDescriptionTextFontSizeOk returns a tuple with the DescriptionTextFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionTextFontSize

`func (o *PortalCustomizeResOpenApiVO) SetDescriptionTextFontSize(v int32)`

SetDescriptionTextFontSize sets DescriptionTextFontSize field to given value.

### HasDescriptionTextFontSize

`func (o *PortalCustomizeResOpenApiVO) HasDescriptionTextFontSize() bool`

HasDescriptionTextFontSize returns a boolean if a field has been set.

### GetDescriptionTextOpacity

`func (o *PortalCustomizeResOpenApiVO) GetDescriptionTextOpacity() int32`

GetDescriptionTextOpacity returns the DescriptionTextOpacity field if non-nil, zero value otherwise.

### GetDescriptionTextOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetDescriptionTextOpacityOk() (*int32, bool)`

GetDescriptionTextOpacityOk returns a tuple with the DescriptionTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionTextOpacity

`func (o *PortalCustomizeResOpenApiVO) SetDescriptionTextOpacity(v int32)`

SetDescriptionTextOpacity sets DescriptionTextOpacity field to given value.

### HasDescriptionTextOpacity

`func (o *PortalCustomizeResOpenApiVO) HasDescriptionTextOpacity() bool`

HasDescriptionTextOpacity returns a boolean if a field has been set.

### GetEnableDeviceSpecificBg

`func (o *PortalCustomizeResOpenApiVO) GetEnableDeviceSpecificBg() bool`

GetEnableDeviceSpecificBg returns the EnableDeviceSpecificBg field if non-nil, zero value otherwise.

### GetEnableDeviceSpecificBgOk

`func (o *PortalCustomizeResOpenApiVO) GetEnableDeviceSpecificBgOk() (*bool, bool)`

GetEnableDeviceSpecificBgOk returns a tuple with the EnableDeviceSpecificBg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDeviceSpecificBg

`func (o *PortalCustomizeResOpenApiVO) SetEnableDeviceSpecificBg(v bool)`

SetEnableDeviceSpecificBg sets EnableDeviceSpecificBg field to given value.

### HasEnableDeviceSpecificBg

`func (o *PortalCustomizeResOpenApiVO) HasEnableDeviceSpecificBg() bool`

HasEnableDeviceSpecificBg returns a boolean if a field has been set.

### GetFormAuthButtonText

`func (o *PortalCustomizeResOpenApiVO) GetFormAuthButtonText() string`

GetFormAuthButtonText returns the FormAuthButtonText field if non-nil, zero value otherwise.

### GetFormAuthButtonTextOk

`func (o *PortalCustomizeResOpenApiVO) GetFormAuthButtonTextOk() (*string, bool)`

GetFormAuthButtonTextOk returns a tuple with the FormAuthButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormAuthButtonText

`func (o *PortalCustomizeResOpenApiVO) SetFormAuthButtonText(v string)`

SetFormAuthButtonText sets FormAuthButtonText field to given value.

### HasFormAuthButtonText

`func (o *PortalCustomizeResOpenApiVO) HasFormAuthButtonText() bool`

HasFormAuthButtonText returns a boolean if a field has been set.

### GetInputBoxBorderColor

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxBorderColor() string`

GetInputBoxBorderColor returns the InputBoxBorderColor field if non-nil, zero value otherwise.

### GetInputBoxBorderColorOk

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxBorderColorOk() (*string, bool)`

GetInputBoxBorderColorOk returns a tuple with the InputBoxBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBoxBorderColor

`func (o *PortalCustomizeResOpenApiVO) SetInputBoxBorderColor(v string)`

SetInputBoxBorderColor sets InputBoxBorderColor field to given value.

### HasInputBoxBorderColor

`func (o *PortalCustomizeResOpenApiVO) HasInputBoxBorderColor() bool`

HasInputBoxBorderColor returns a boolean if a field has been set.

### GetInputBoxBorderOpacity

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxBorderOpacity() int32`

GetInputBoxBorderOpacity returns the InputBoxBorderOpacity field if non-nil, zero value otherwise.

### GetInputBoxBorderOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxBorderOpacityOk() (*int32, bool)`

GetInputBoxBorderOpacityOk returns a tuple with the InputBoxBorderOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBoxBorderOpacity

`func (o *PortalCustomizeResOpenApiVO) SetInputBoxBorderOpacity(v int32)`

SetInputBoxBorderOpacity sets InputBoxBorderOpacity field to given value.

### HasInputBoxBorderOpacity

`func (o *PortalCustomizeResOpenApiVO) HasInputBoxBorderOpacity() bool`

HasInputBoxBorderOpacity returns a boolean if a field has been set.

### GetInputBoxColor

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxColor() string`

GetInputBoxColor returns the InputBoxColor field if non-nil, zero value otherwise.

### GetInputBoxColorOk

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxColorOk() (*string, bool)`

GetInputBoxColorOk returns a tuple with the InputBoxColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBoxColor

`func (o *PortalCustomizeResOpenApiVO) SetInputBoxColor(v string)`

SetInputBoxColor sets InputBoxColor field to given value.

### HasInputBoxColor

`func (o *PortalCustomizeResOpenApiVO) HasInputBoxColor() bool`

HasInputBoxColor returns a boolean if a field has been set.

### GetInputBoxOpacity

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxOpacity() int32`

GetInputBoxOpacity returns the InputBoxOpacity field if non-nil, zero value otherwise.

### GetInputBoxOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxOpacityOk() (*int32, bool)`

GetInputBoxOpacityOk returns a tuple with the InputBoxOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBoxOpacity

`func (o *PortalCustomizeResOpenApiVO) SetInputBoxOpacity(v int32)`

SetInputBoxOpacity sets InputBoxOpacity field to given value.

### HasInputBoxOpacity

`func (o *PortalCustomizeResOpenApiVO) HasInputBoxOpacity() bool`

HasInputBoxOpacity returns a boolean if a field has been set.

### GetInputBoxRadius

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxRadius() int32`

GetInputBoxRadius returns the InputBoxRadius field if non-nil, zero value otherwise.

### GetInputBoxRadiusOk

`func (o *PortalCustomizeResOpenApiVO) GetInputBoxRadiusOk() (*int32, bool)`

GetInputBoxRadiusOk returns a tuple with the InputBoxRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBoxRadius

`func (o *PortalCustomizeResOpenApiVO) SetInputBoxRadius(v int32)`

SetInputBoxRadius sets InputBoxRadius field to given value.

### HasInputBoxRadius

`func (o *PortalCustomizeResOpenApiVO) HasInputBoxRadius() bool`

HasInputBoxRadius returns a boolean if a field has been set.

### GetInputTextColor

`func (o *PortalCustomizeResOpenApiVO) GetInputTextColor() string`

GetInputTextColor returns the InputTextColor field if non-nil, zero value otherwise.

### GetInputTextColorOk

`func (o *PortalCustomizeResOpenApiVO) GetInputTextColorOk() (*string, bool)`

GetInputTextColorOk returns a tuple with the InputTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTextColor

`func (o *PortalCustomizeResOpenApiVO) SetInputTextColor(v string)`

SetInputTextColor sets InputTextColor field to given value.

### HasInputTextColor

`func (o *PortalCustomizeResOpenApiVO) HasInputTextColor() bool`

HasInputTextColor returns a boolean if a field has been set.

### GetInputTextOpacity

`func (o *PortalCustomizeResOpenApiVO) GetInputTextOpacity() int32`

GetInputTextOpacity returns the InputTextOpacity field if non-nil, zero value otherwise.

### GetInputTextOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetInputTextOpacityOk() (*int32, bool)`

GetInputTextOpacityOk returns a tuple with the InputTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputTextOpacity

`func (o *PortalCustomizeResOpenApiVO) SetInputTextOpacity(v int32)`

SetInputTextOpacity sets InputTextOpacity field to given value.

### HasInputTextOpacity

`func (o *PortalCustomizeResOpenApiVO) HasInputTextOpacity() bool`

HasInputTextOpacity returns a boolean if a field has been set.

### GetLogoDisplay

`func (o *PortalCustomizeResOpenApiVO) GetLogoDisplay() bool`

GetLogoDisplay returns the LogoDisplay field if non-nil, zero value otherwise.

### GetLogoDisplayOk

`func (o *PortalCustomizeResOpenApiVO) GetLogoDisplayOk() (*bool, bool)`

GetLogoDisplayOk returns a tuple with the LogoDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoDisplay

`func (o *PortalCustomizeResOpenApiVO) SetLogoDisplay(v bool)`

SetLogoDisplay sets LogoDisplay field to given value.

### HasLogoDisplay

`func (o *PortalCustomizeResOpenApiVO) HasLogoDisplay() bool`

HasLogoDisplay returns a boolean if a field has been set.

### GetLogoHorizontalPosition

`func (o *PortalCustomizeResOpenApiVO) GetLogoHorizontalPosition() int32`

GetLogoHorizontalPosition returns the LogoHorizontalPosition field if non-nil, zero value otherwise.

### GetLogoHorizontalPositionOk

`func (o *PortalCustomizeResOpenApiVO) GetLogoHorizontalPositionOk() (*int32, bool)`

GetLogoHorizontalPositionOk returns a tuple with the LogoHorizontalPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoHorizontalPosition

`func (o *PortalCustomizeResOpenApiVO) SetLogoHorizontalPosition(v int32)`

SetLogoHorizontalPosition sets LogoHorizontalPosition field to given value.

### HasLogoHorizontalPosition

`func (o *PortalCustomizeResOpenApiVO) HasLogoHorizontalPosition() bool`

HasLogoHorizontalPosition returns a boolean if a field has been set.

### GetLogoPicture

`func (o *PortalCustomizeResOpenApiVO) GetLogoPicture() PortalPictureInfo`

GetLogoPicture returns the LogoPicture field if non-nil, zero value otherwise.

### GetLogoPictureOk

`func (o *PortalCustomizeResOpenApiVO) GetLogoPictureOk() (*PortalPictureInfo, bool)`

GetLogoPictureOk returns a tuple with the LogoPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPicture

`func (o *PortalCustomizeResOpenApiVO) SetLogoPicture(v PortalPictureInfo)`

SetLogoPicture sets LogoPicture field to given value.

### HasLogoPicture

`func (o *PortalCustomizeResOpenApiVO) HasLogoPicture() bool`

HasLogoPicture returns a boolean if a field has been set.

### GetLogoPosition

`func (o *PortalCustomizeResOpenApiVO) GetLogoPosition() int32`

GetLogoPosition returns the LogoPosition field if non-nil, zero value otherwise.

### GetLogoPositionOk

`func (o *PortalCustomizeResOpenApiVO) GetLogoPositionOk() (*int32, bool)`

GetLogoPositionOk returns a tuple with the LogoPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPosition

`func (o *PortalCustomizeResOpenApiVO) SetLogoPosition(v int32)`

SetLogoPosition sets LogoPosition field to given value.

### HasLogoPosition

`func (o *PortalCustomizeResOpenApiVO) HasLogoPosition() bool`

HasLogoPosition returns a boolean if a field has been set.

### GetLogoPositionRatio

`func (o *PortalCustomizeResOpenApiVO) GetLogoPositionRatio() float32`

GetLogoPositionRatio returns the LogoPositionRatio field if non-nil, zero value otherwise.

### GetLogoPositionRatioOk

`func (o *PortalCustomizeResOpenApiVO) GetLogoPositionRatioOk() (*float32, bool)`

GetLogoPositionRatioOk returns a tuple with the LogoPositionRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPositionRatio

`func (o *PortalCustomizeResOpenApiVO) SetLogoPositionRatio(v float32)`

SetLogoPositionRatio sets LogoPositionRatio field to given value.

### HasLogoPositionRatio

`func (o *PortalCustomizeResOpenApiVO) HasLogoPositionRatio() bool`

HasLogoPositionRatio returns a boolean if a field has been set.

### GetLogoSize

`func (o *PortalCustomizeResOpenApiVO) GetLogoSize() int32`

GetLogoSize returns the LogoSize field if non-nil, zero value otherwise.

### GetLogoSizeOk

`func (o *PortalCustomizeResOpenApiVO) GetLogoSizeOk() (*int32, bool)`

GetLogoSizeOk returns a tuple with the LogoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSize

`func (o *PortalCustomizeResOpenApiVO) SetLogoSize(v int32)`

SetLogoSize sets LogoSize field to given value.

### HasLogoSize

`func (o *PortalCustomizeResOpenApiVO) HasLogoSize() bool`

HasLogoSize returns a boolean if a field has been set.

### GetLogoTranslate

`func (o *PortalCustomizeResOpenApiVO) GetLogoTranslate() int32`

GetLogoTranslate returns the LogoTranslate field if non-nil, zero value otherwise.

### GetLogoTranslateOk

`func (o *PortalCustomizeResOpenApiVO) GetLogoTranslateOk() (*int32, bool)`

GetLogoTranslateOk returns a tuple with the LogoTranslate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoTranslate

`func (o *PortalCustomizeResOpenApiVO) SetLogoTranslate(v int32)`

SetLogoTranslate sets LogoTranslate field to given value.

### HasLogoTranslate

`func (o *PortalCustomizeResOpenApiVO) HasLogoTranslate() bool`

HasLogoTranslate returns a boolean if a field has been set.

### GetMobileBackgroundPicture

`func (o *PortalCustomizeResOpenApiVO) GetMobileBackgroundPicture() PortalPictureInfo`

GetMobileBackgroundPicture returns the MobileBackgroundPicture field if non-nil, zero value otherwise.

### GetMobileBackgroundPictureOk

`func (o *PortalCustomizeResOpenApiVO) GetMobileBackgroundPictureOk() (*PortalPictureInfo, bool)`

GetMobileBackgroundPictureOk returns a tuple with the MobileBackgroundPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileBackgroundPicture

`func (o *PortalCustomizeResOpenApiVO) SetMobileBackgroundPicture(v PortalPictureInfo)`

SetMobileBackgroundPicture sets MobileBackgroundPicture field to given value.

### HasMobileBackgroundPicture

`func (o *PortalCustomizeResOpenApiVO) HasMobileBackgroundPicture() bool`

HasMobileBackgroundPicture returns a boolean if a field has been set.

### GetMobileBgPicCoordinatesOfLibrary

`func (o *PortalCustomizeResOpenApiVO) GetMobileBgPicCoordinatesOfLibrary() BgPicCoordinatesOfLibraryOpenApiVO`

GetMobileBgPicCoordinatesOfLibrary returns the MobileBgPicCoordinatesOfLibrary field if non-nil, zero value otherwise.

### GetMobileBgPicCoordinatesOfLibraryOk

`func (o *PortalCustomizeResOpenApiVO) GetMobileBgPicCoordinatesOfLibraryOk() (*BgPicCoordinatesOfLibraryOpenApiVO, bool)`

GetMobileBgPicCoordinatesOfLibraryOk returns a tuple with the MobileBgPicCoordinatesOfLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileBgPicCoordinatesOfLibrary

`func (o *PortalCustomizeResOpenApiVO) SetMobileBgPicCoordinatesOfLibrary(v BgPicCoordinatesOfLibraryOpenApiVO)`

SetMobileBgPicCoordinatesOfLibrary sets MobileBgPicCoordinatesOfLibrary field to given value.

### HasMobileBgPicCoordinatesOfLibrary

`func (o *PortalCustomizeResOpenApiVO) HasMobileBgPicCoordinatesOfLibrary() bool`

HasMobileBgPicCoordinatesOfLibrary returns a boolean if a field has been set.

### GetPcAlign

`func (o *PortalCustomizeResOpenApiVO) GetPcAlign() int32`

GetPcAlign returns the PcAlign field if non-nil, zero value otherwise.

### GetPcAlignOk

`func (o *PortalCustomizeResOpenApiVO) GetPcAlignOk() (*int32, bool)`

GetPcAlignOk returns a tuple with the PcAlign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcAlign

`func (o *PortalCustomizeResOpenApiVO) SetPcAlign(v int32)`

SetPcAlign sets PcAlign field to given value.

### HasPcAlign

`func (o *PortalCustomizeResOpenApiVO) HasPcAlign() bool`

HasPcAlign returns a boolean if a field has been set.

### GetRedirectionCountDownEnable

`func (o *PortalCustomizeResOpenApiVO) GetRedirectionCountDownEnable() bool`

GetRedirectionCountDownEnable returns the RedirectionCountDownEnable field if non-nil, zero value otherwise.

### GetRedirectionCountDownEnableOk

`func (o *PortalCustomizeResOpenApiVO) GetRedirectionCountDownEnableOk() (*bool, bool)`

GetRedirectionCountDownEnableOk returns a tuple with the RedirectionCountDownEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectionCountDownEnable

`func (o *PortalCustomizeResOpenApiVO) SetRedirectionCountDownEnable(v bool)`

SetRedirectionCountDownEnable sets RedirectionCountDownEnable field to given value.

### HasRedirectionCountDownEnable

`func (o *PortalCustomizeResOpenApiVO) HasRedirectionCountDownEnable() bool`

HasRedirectionCountDownEnable returns a boolean if a field has been set.

### GetTermsOfService

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *PortalCustomizeResOpenApiVO) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.

### HasTermsOfService

`func (o *PortalCustomizeResOpenApiVO) HasTermsOfService() bool`

HasTermsOfService returns a boolean if a field has been set.

### GetTermsOfServiceEnable

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceEnable() bool`

GetTermsOfServiceEnable returns the TermsOfServiceEnable field if non-nil, zero value otherwise.

### GetTermsOfServiceEnableOk

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceEnableOk() (*bool, bool)`

GetTermsOfServiceEnableOk returns a tuple with the TermsOfServiceEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceEnable

`func (o *PortalCustomizeResOpenApiVO) SetTermsOfServiceEnable(v bool)`

SetTermsOfServiceEnable sets TermsOfServiceEnable field to given value.

### HasTermsOfServiceEnable

`func (o *PortalCustomizeResOpenApiVO) HasTermsOfServiceEnable() bool`

HasTermsOfServiceEnable returns a boolean if a field has been set.

### GetTermsOfServiceFontSize

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceFontSize() int32`

GetTermsOfServiceFontSize returns the TermsOfServiceFontSize field if non-nil, zero value otherwise.

### GetTermsOfServiceFontSizeOk

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceFontSizeOk() (*int32, bool)`

GetTermsOfServiceFontSizeOk returns a tuple with the TermsOfServiceFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceFontSize

`func (o *PortalCustomizeResOpenApiVO) SetTermsOfServiceFontSize(v int32)`

SetTermsOfServiceFontSize sets TermsOfServiceFontSize field to given value.

### HasTermsOfServiceFontSize

`func (o *PortalCustomizeResOpenApiVO) HasTermsOfServiceFontSize() bool`

HasTermsOfServiceFontSize returns a boolean if a field has been set.

### GetTermsOfServiceText

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceText() string`

GetTermsOfServiceText returns the TermsOfServiceText field if non-nil, zero value otherwise.

### GetTermsOfServiceTextOk

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceTextOk() (*string, bool)`

GetTermsOfServiceTextOk returns a tuple with the TermsOfServiceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceText

`func (o *PortalCustomizeResOpenApiVO) SetTermsOfServiceText(v string)`

SetTermsOfServiceText sets TermsOfServiceText field to given value.

### HasTermsOfServiceText

`func (o *PortalCustomizeResOpenApiVO) HasTermsOfServiceText() bool`

HasTermsOfServiceText returns a boolean if a field has been set.

### GetTermsOfServiceTextColor

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceTextColor() string`

GetTermsOfServiceTextColor returns the TermsOfServiceTextColor field if non-nil, zero value otherwise.

### GetTermsOfServiceTextColorOk

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceTextColorOk() (*string, bool)`

GetTermsOfServiceTextColorOk returns a tuple with the TermsOfServiceTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceTextColor

`func (o *PortalCustomizeResOpenApiVO) SetTermsOfServiceTextColor(v string)`

SetTermsOfServiceTextColor sets TermsOfServiceTextColor field to given value.

### HasTermsOfServiceTextColor

`func (o *PortalCustomizeResOpenApiVO) HasTermsOfServiceTextColor() bool`

HasTermsOfServiceTextColor returns a boolean if a field has been set.

### GetTermsOfServiceTextOpacity

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceTextOpacity() int32`

GetTermsOfServiceTextOpacity returns the TermsOfServiceTextOpacity field if non-nil, zero value otherwise.

### GetTermsOfServiceTextOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceTextOpacityOk() (*int32, bool)`

GetTermsOfServiceTextOpacityOk returns a tuple with the TermsOfServiceTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceTextOpacity

`func (o *PortalCustomizeResOpenApiVO) SetTermsOfServiceTextOpacity(v int32)`

SetTermsOfServiceTextOpacity sets TermsOfServiceTextOpacity field to given value.

### HasTermsOfServiceTextOpacity

`func (o *PortalCustomizeResOpenApiVO) HasTermsOfServiceTextOpacity() bool`

HasTermsOfServiceTextOpacity returns a boolean if a field has been set.

### GetTermsOfServiceUrlTexts

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceUrlTexts() []TermsOfServiceUrlVO`

GetTermsOfServiceUrlTexts returns the TermsOfServiceUrlTexts field if non-nil, zero value otherwise.

### GetTermsOfServiceUrlTextsOk

`func (o *PortalCustomizeResOpenApiVO) GetTermsOfServiceUrlTextsOk() (*[]TermsOfServiceUrlVO, bool)`

GetTermsOfServiceUrlTextsOk returns a tuple with the TermsOfServiceUrlTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfServiceUrlTexts

`func (o *PortalCustomizeResOpenApiVO) SetTermsOfServiceUrlTexts(v []TermsOfServiceUrlVO)`

SetTermsOfServiceUrlTexts sets TermsOfServiceUrlTexts field to given value.

### HasTermsOfServiceUrlTexts

`func (o *PortalCustomizeResOpenApiVO) HasTermsOfServiceUrlTexts() bool`

HasTermsOfServiceUrlTexts returns a boolean if a field has been set.

### GetWelcomeEnable

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeEnable() bool`

GetWelcomeEnable returns the WelcomeEnable field if non-nil, zero value otherwise.

### GetWelcomeEnableOk

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeEnableOk() (*bool, bool)`

GetWelcomeEnableOk returns a tuple with the WelcomeEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeEnable

`func (o *PortalCustomizeResOpenApiVO) SetWelcomeEnable(v bool)`

SetWelcomeEnable sets WelcomeEnable field to given value.

### HasWelcomeEnable

`func (o *PortalCustomizeResOpenApiVO) HasWelcomeEnable() bool`

HasWelcomeEnable returns a boolean if a field has been set.

### GetWelcomeInformation

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeInformation() string`

GetWelcomeInformation returns the WelcomeInformation field if non-nil, zero value otherwise.

### GetWelcomeInformationOk

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeInformationOk() (*string, bool)`

GetWelcomeInformationOk returns a tuple with the WelcomeInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeInformation

`func (o *PortalCustomizeResOpenApiVO) SetWelcomeInformation(v string)`

SetWelcomeInformation sets WelcomeInformation field to given value.

### HasWelcomeInformation

`func (o *PortalCustomizeResOpenApiVO) HasWelcomeInformation() bool`

HasWelcomeInformation returns a boolean if a field has been set.

### GetWelcomeTextColor

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeTextColor() string`

GetWelcomeTextColor returns the WelcomeTextColor field if non-nil, zero value otherwise.

### GetWelcomeTextColorOk

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeTextColorOk() (*string, bool)`

GetWelcomeTextColorOk returns a tuple with the WelcomeTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTextColor

`func (o *PortalCustomizeResOpenApiVO) SetWelcomeTextColor(v string)`

SetWelcomeTextColor sets WelcomeTextColor field to given value.

### HasWelcomeTextColor

`func (o *PortalCustomizeResOpenApiVO) HasWelcomeTextColor() bool`

HasWelcomeTextColor returns a boolean if a field has been set.

### GetWelcomeTextFontSize

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeTextFontSize() int32`

GetWelcomeTextFontSize returns the WelcomeTextFontSize field if non-nil, zero value otherwise.

### GetWelcomeTextFontSizeOk

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeTextFontSizeOk() (*int32, bool)`

GetWelcomeTextFontSizeOk returns a tuple with the WelcomeTextFontSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTextFontSize

`func (o *PortalCustomizeResOpenApiVO) SetWelcomeTextFontSize(v int32)`

SetWelcomeTextFontSize sets WelcomeTextFontSize field to given value.

### HasWelcomeTextFontSize

`func (o *PortalCustomizeResOpenApiVO) HasWelcomeTextFontSize() bool`

HasWelcomeTextFontSize returns a boolean if a field has been set.

### GetWelcomeTextOpacity

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeTextOpacity() int32`

GetWelcomeTextOpacity returns the WelcomeTextOpacity field if non-nil, zero value otherwise.

### GetWelcomeTextOpacityOk

`func (o *PortalCustomizeResOpenApiVO) GetWelcomeTextOpacityOk() (*int32, bool)`

GetWelcomeTextOpacityOk returns a tuple with the WelcomeTextOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTextOpacity

`func (o *PortalCustomizeResOpenApiVO) SetWelcomeTextOpacity(v int32)`

SetWelcomeTextOpacity sets WelcomeTextOpacity field to given value.

### HasWelcomeTextOpacity

`func (o *PortalCustomizeResOpenApiVO) HasWelcomeTextOpacity() bool`

HasWelcomeTextOpacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


