// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// MergeUsers Merges interactions (purchases, ratings, bookmarks, detail views ...) of two different users under a single user ID. This is especially useful for online e-commerce applications working with anonymous users identified by unique tokens such as the session ID. In such applications, it may often happen that a user owns a persistent account, yet accesses the system anonymously while, e.g., putting items into a shopping cart. At some point in time, such as when the user wishes to confirm the purchase, (s)he logs into the system using his/her username and password. The interactions made under anonymous session ID then become connected with the persistent account, and merging these two becomes desirable.
// Merging happens between two users referred to as the *target* and the *source*. After the merge, all the interactions of the source user are attributed to the target user, and the source user is **deleted**.
// By default, the *Merge Users* request is only available from server-side integrations for security reasons, to prevent potential abuse.
// If you need to call this request from a client-side environment (such as a web or mobile app), please contact our support and request access to enable this feature for your database.
type MergeUsers struct {
	ApiRequest
	client ApiClient
}

// NewMergeUsers creates MergeUsers request.
// Merges interactions (purchases, ratings, bookmarks, detail views ...) of two different users under a single user ID. This is especially useful for online e-commerce applications working with anonymous users identified by unique tokens such as the session ID. In such applications, it may often happen that a user owns a persistent account, yet accesses the system anonymously while, e.g., putting items into a shopping cart. At some point in time, such as when the user wishes to confirm the purchase, (s)he logs into the system using his/her username and password. The interactions made under anonymous session ID then become connected with the persistent account, and merging these two becomes desirable.
// Merging happens between two users referred to as the *target* and the *source*. After the merge, all the interactions of the source user are attributed to the target user, and the source user is **deleted**.
// By default, the *Merge Users* request is only available from server-side integrations for security reasons, to prevent potential abuse.
// If you need to call this request from a client-side environment (such as a web or mobile app), please contact our support and request access to enable this feature for your database.
func NewMergeUsers(client ApiClient, targetUserId string, sourceUserId string) *MergeUsers {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &MergeUsers{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/users/%s/merge/%s", targetUserId, sourceUserId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the user *targetUserId* should be created if not present in the database.
func (r *MergeUsers) SetCascadeCreate(cascadeCreate bool) *MergeUsers {
	r.QueryParameters["cascadeCreate"] = cascadeCreate
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *MergeUsers) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *MergeUsers) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
