# UploadLocalUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**HotspotPortalsOpenApiVO**](HotspotPortalsOpenApiVO.md) |  | 
**File** | ***os.File** |  | 

## Methods

### NewUploadLocalUsersRequest

`func NewUploadLocalUsersRequest(config HotspotPortalsOpenApiVO, file *os.File, ) *UploadLocalUsersRequest`

NewUploadLocalUsersRequest instantiates a new UploadLocalUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadLocalUsersRequestWithDefaults

`func NewUploadLocalUsersRequestWithDefaults() *UploadLocalUsersRequest`

NewUploadLocalUsersRequestWithDefaults instantiates a new UploadLocalUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UploadLocalUsersRequest) GetConfig() HotspotPortalsOpenApiVO`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UploadLocalUsersRequest) GetConfigOk() (*HotspotPortalsOpenApiVO, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UploadLocalUsersRequest) SetConfig(v HotspotPortalsOpenApiVO)`

SetConfig sets Config field to given value.


### GetFile

`func (o *UploadLocalUsersRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *UploadLocalUsersRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *UploadLocalUsersRequest) SetFile(v *os.File)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


