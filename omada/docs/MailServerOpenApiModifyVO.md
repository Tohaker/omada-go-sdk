# MailServerOpenApiModifyVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | **bool** | Enable this feature if the login of the mailbox requires a username and authorization code. | 
**AuthCode** | Pointer to **string** | When Authentication is enabled, enter the authorization code that enables a third party to log in to the mailbox. Note that the authorization code is not the mailbox&#39;s password. | [optional] 
**Receiver** | **string** | The email address of the receiver, which can be the same as or different from the sender&#39;s email address. | 
**Sender** | **string** | The email address of the sender. | 
**SmtpPort** | Pointer to **int32** | Enter the port used by the SMTP server according to the instructions of your email service provider. | [optional] 
**SmtpServer** | Pointer to **string** | Enter the domain name or IP address of the SMTP server. | [optional] 
**Ssl** | **bool** | Enable this feature, and the data will be transmitted based on the SSL protocol. | 
**Username** | Pointer to **string** | When Authentication is enabled, enter your email address as the username. | [optional] 

## Methods

### NewMailServerOpenApiModifyVO

`func NewMailServerOpenApiModifyVO(auth bool, receiver string, sender string, ssl bool, ) *MailServerOpenApiModifyVO`

NewMailServerOpenApiModifyVO instantiates a new MailServerOpenApiModifyVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailServerOpenApiModifyVOWithDefaults

`func NewMailServerOpenApiModifyVOWithDefaults() *MailServerOpenApiModifyVO`

NewMailServerOpenApiModifyVOWithDefaults instantiates a new MailServerOpenApiModifyVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *MailServerOpenApiModifyVO) GetAuth() bool`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *MailServerOpenApiModifyVO) GetAuthOk() (*bool, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *MailServerOpenApiModifyVO) SetAuth(v bool)`

SetAuth sets Auth field to given value.


### GetAuthCode

`func (o *MailServerOpenApiModifyVO) GetAuthCode() string`

GetAuthCode returns the AuthCode field if non-nil, zero value otherwise.

### GetAuthCodeOk

`func (o *MailServerOpenApiModifyVO) GetAuthCodeOk() (*string, bool)`

GetAuthCodeOk returns a tuple with the AuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCode

`func (o *MailServerOpenApiModifyVO) SetAuthCode(v string)`

SetAuthCode sets AuthCode field to given value.

### HasAuthCode

`func (o *MailServerOpenApiModifyVO) HasAuthCode() bool`

HasAuthCode returns a boolean if a field has been set.

### GetReceiver

`func (o *MailServerOpenApiModifyVO) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *MailServerOpenApiModifyVO) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *MailServerOpenApiModifyVO) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetSender

`func (o *MailServerOpenApiModifyVO) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *MailServerOpenApiModifyVO) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *MailServerOpenApiModifyVO) SetSender(v string)`

SetSender sets Sender field to given value.


### GetSmtpPort

`func (o *MailServerOpenApiModifyVO) GetSmtpPort() int32`

GetSmtpPort returns the SmtpPort field if non-nil, zero value otherwise.

### GetSmtpPortOk

`func (o *MailServerOpenApiModifyVO) GetSmtpPortOk() (*int32, bool)`

GetSmtpPortOk returns a tuple with the SmtpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpPort

`func (o *MailServerOpenApiModifyVO) SetSmtpPort(v int32)`

SetSmtpPort sets SmtpPort field to given value.

### HasSmtpPort

`func (o *MailServerOpenApiModifyVO) HasSmtpPort() bool`

HasSmtpPort returns a boolean if a field has been set.

### GetSmtpServer

`func (o *MailServerOpenApiModifyVO) GetSmtpServer() string`

GetSmtpServer returns the SmtpServer field if non-nil, zero value otherwise.

### GetSmtpServerOk

`func (o *MailServerOpenApiModifyVO) GetSmtpServerOk() (*string, bool)`

GetSmtpServerOk returns a tuple with the SmtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmtpServer

`func (o *MailServerOpenApiModifyVO) SetSmtpServer(v string)`

SetSmtpServer sets SmtpServer field to given value.

### HasSmtpServer

`func (o *MailServerOpenApiModifyVO) HasSmtpServer() bool`

HasSmtpServer returns a boolean if a field has been set.

### GetSsl

`func (o *MailServerOpenApiModifyVO) GetSsl() bool`

GetSsl returns the Ssl field if non-nil, zero value otherwise.

### GetSslOk

`func (o *MailServerOpenApiModifyVO) GetSslOk() (*bool, bool)`

GetSslOk returns a tuple with the Ssl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsl

`func (o *MailServerOpenApiModifyVO) SetSsl(v bool)`

SetSsl sets Ssl field to given value.


### GetUsername

`func (o *MailServerOpenApiModifyVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MailServerOpenApiModifyVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MailServerOpenApiModifyVO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *MailServerOpenApiModifyVO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


