# \ContainersAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DevicesDeviceIdContainersGet**](ContainersAPI.md#DevicesDeviceIdContainersGet) | **Get** /devices/{device_id}/containers | 



## DevicesDeviceIdContainersGet

> ContainerList DevicesDeviceIdContainersGet(ctx, deviceId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	deviceId := "deviceId_example" // string | The unique identifier of the device

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainersAPI.DevicesDeviceIdContainersGet(context.Background(), deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainersAPI.DevicesDeviceIdContainersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesDeviceIdContainersGet`: ContainerList
	fmt.Fprintf(os.Stdout, "Response from `ContainersAPI.DevicesDeviceIdContainersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique identifier of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceIdContainersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerList**](ContainerList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

