# \GroupsAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsGroupIdGet**](GroupsAPI.md#GroupsGroupIdGet) | **Get** /groups/{group_id} | 



## GroupsGroupIdGet

> Group GroupsGroupIdGet(ctx, groupId).Execute()



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
	groupId := "groupId_example" // string | The unique identifier of the group

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GroupsGroupIdGet(context.Background(), groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GroupsGroupIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdGet`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GroupsGroupIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique identifier of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Group**](Group.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

