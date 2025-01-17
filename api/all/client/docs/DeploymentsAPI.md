# \DeploymentsAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeploymentsDeploymentIdGet**](DeploymentsAPI.md#DeploymentsDeploymentIdGet) | **Get** /deployments/{deployment_id} | 
[**DeploymentsDeploymentIdLogsGet**](DeploymentsAPI.md#DeploymentsDeploymentIdLogsGet) | **Get** /deployments/{deployment_id}/logs | 
[**DevicesDeviceIdDeploymentsGet**](DeploymentsAPI.md#DevicesDeviceIdDeploymentsGet) | **Get** /devices/{device_id}/deployments | 
[**WorkspacesWorkspaceIdDeployPost**](DeploymentsAPI.md#WorkspacesWorkspaceIdDeployPost) | **Post** /workspaces/{workspace_id}/deploy | 



## DeploymentsDeploymentIdGet

> Deployment DeploymentsDeploymentIdGet(ctx, deploymentId).Execute()



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
	deploymentId := "deploymentId_example" // string | The unique identifier of the deployment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.DeploymentsDeploymentIdGet(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.DeploymentsDeploymentIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentsDeploymentIdGet`: Deployment
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.DeploymentsDeploymentIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** | The unique identifier of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentsDeploymentIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deployment**](Deployment.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeploymentsDeploymentIdLogsGet

> DeploymentLogList DeploymentsDeploymentIdLogsGet(ctx, deploymentId).Execute()



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
	deploymentId := "deploymentId_example" // string | The unique identifier of the deployment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.DeploymentsDeploymentIdLogsGet(context.Background(), deploymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.DeploymentsDeploymentIdLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeploymentsDeploymentIdLogsGet`: DeploymentLogList
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.DeploymentsDeploymentIdLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deploymentId** | **string** | The unique identifier of the deployment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeploymentsDeploymentIdLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeploymentLogList**](DeploymentLogList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DevicesDeviceIdDeploymentsGet

> DeploymentList DevicesDeviceIdDeploymentsGet(ctx, deviceId).OnDevice(onDevice).Execute()



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
	onDevice := true // bool | Whether to include only on device deployments (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeploymentsAPI.DevicesDeviceIdDeploymentsGet(context.Background(), deviceId).OnDevice(onDevice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.DevicesDeviceIdDeploymentsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DevicesDeviceIdDeploymentsGet`: DeploymentList
	fmt.Fprintf(os.Stdout, "Response from `DeploymentsAPI.DevicesDeviceIdDeploymentsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | The unique identifier of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiDevicesDeviceIdDeploymentsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onDevice** | **bool** | Whether to include only on device deployments | 

### Return type

[**DeploymentList**](DeploymentList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceIdDeployPost

> WorkspacesWorkspaceIdDeployPost(ctx, workspaceId).CreateDeployments(createDeployments).Execute()



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
	workspaceId := "workspaceId_example" // string | The unique identifier of the workspace
	createDeployments := *openapiclient.NewCreateDeployments([]string{"art_1234"}, []string{"device_1234"}) // CreateDeployments | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DeploymentsAPI.WorkspacesWorkspaceIdDeployPost(context.Background(), workspaceId).CreateDeployments(createDeployments).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeploymentsAPI.WorkspacesWorkspaceIdDeployPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdDeployPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDeployments** | [**CreateDeployments**](CreateDeployments.md) |  | 

### Return type

 (empty response body)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

