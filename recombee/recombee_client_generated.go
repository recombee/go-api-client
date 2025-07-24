// This file is auto-generated

package recombee

import (
	"github.com/recombee/go-api-client/v5/recombee/requests"
)

// NewAddItem creates AddItem request.
// Adds new item of the given `itemId` to the items catalog.
// All the item properties for the newly created items are set to null.
func (c *RecombeeClient) NewAddItem(itemId string) *requests.AddItem {
	return requests.NewAddItem(c, itemId)
}

// NewDeleteItem creates DeleteItem request.
// Deletes an item of the given `itemId` from the catalog.
// If there are any *purchases*, *ratings*, *bookmarks*, *cart additions*, or *detail views* of the item present in the database, they will be deleted in cascade as well. Also, if the item is present in some *series*, it will be removed from all the *series* where present.
// If an item becomes obsolete/no longer available, it is meaningful to keep it in the catalog (along with all the interaction data, which are very useful), and **only exclude the item from recommendations**. In such a case, use [ReQL filter](https://docs.recombee.com/reql) instead of deleting the item completely.
func (c *RecombeeClient) NewDeleteItem(itemId string) *requests.DeleteItem {
	return requests.NewDeleteItem(c, itemId)
}

// NewGetItemValues creates GetItemValues request.
// Gets all the current property values of the given item.
func (c *RecombeeClient) NewGetItemValues(itemId string) *requests.GetItemValues {
	return requests.NewGetItemValues(c, itemId)
}

// NewListItems creates ListItems request.
// Gets a list of IDs of items currently present in the catalog.
func (c *RecombeeClient) NewListItems() *requests.ListItems {
	return requests.NewListItems(c)
}

// NewAddItemProperty creates AddItemProperty request.
// Adding an item property is somewhat equivalent to adding a column to the table of items. The items may be characterized by various properties of different types.
func (c *RecombeeClient) NewAddItemProperty(propertyName string, propertyType string) *requests.AddItemProperty {
	return requests.NewAddItemProperty(c, propertyName, propertyType)
}

// NewDeleteItemProperty creates DeleteItemProperty request.
// Deleting an item property is roughly equivalent to removing a column from the table of items.
func (c *RecombeeClient) NewDeleteItemProperty(propertyName string) *requests.DeleteItemProperty {
	return requests.NewDeleteItemProperty(c, propertyName)
}

// NewGetItemPropertyInfo creates GetItemPropertyInfo request.
// Gets information about specified item property.
func (c *RecombeeClient) NewGetItemPropertyInfo(propertyName string) *requests.GetItemPropertyInfo {
	return requests.NewGetItemPropertyInfo(c, propertyName)
}

// NewListItemProperties creates ListItemProperties request.
// Gets the list of all the item properties in your database.
func (c *RecombeeClient) NewListItemProperties() *requests.ListItemProperties {
	return requests.NewListItemProperties(c)
}

// NewUpdateMoreItems creates UpdateMoreItems request.
// Updates (some) property values of all the items that pass the filter.
// Example: *Setting all the items that are older than a week as unavailable*
//
//	```json
//	  {
//	    "filter": "'releaseDate' < now() - 7*24*3600",
//	    "changes": {"available": false}
//	  }
//	```
func (c *RecombeeClient) NewUpdateMoreItems(filter string, changes map[string]interface{}) *requests.UpdateMoreItems {
	return requests.NewUpdateMoreItems(c, filter, changes)
}

// NewDeleteMoreItems creates DeleteMoreItems request.
// Deletes all the items that pass the filter.
// If an item becomes obsolete/no longer available, it is meaningful to **keep it in the catalog** (along with all the interaction data, which are very useful) and **only exclude the item from recommendations**. In such a case, use [ReQL filter](https://docs.recombee.com/reql) instead of deleting the item completely.
func (c *RecombeeClient) NewDeleteMoreItems(filter string) *requests.DeleteMoreItems {
	return requests.NewDeleteMoreItems(c, filter)
}

// NewAddSeries creates AddSeries request.
// Creates a new series in the database.
func (c *RecombeeClient) NewAddSeries(seriesId string) *requests.AddSeries {
	return requests.NewAddSeries(c, seriesId)
}

// NewDeleteSeries creates DeleteSeries request.
// Deletes the series of the given `seriesId` from the database.
// Deleting a series will only delete assignment of items to it, not the items themselves!
func (c *RecombeeClient) NewDeleteSeries(seriesId string) *requests.DeleteSeries {
	return requests.NewDeleteSeries(c, seriesId)
}

// NewListSeries creates ListSeries request.
// Gets the list of all the series currently present in the database.
func (c *RecombeeClient) NewListSeries() *requests.ListSeries {
	return requests.NewListSeries(c)
}

// NewListSeriesItems creates ListSeriesItems request.
// Lists all the items present in the given series, sorted according to their time index values.
func (c *RecombeeClient) NewListSeriesItems(seriesId string) *requests.ListSeriesItems {
	return requests.NewListSeriesItems(c, seriesId)
}

// NewInsertToSeries creates InsertToSeries request.
// Inserts an existing item/series into a series of the given seriesId at a position determined by time.
func (c *RecombeeClient) NewInsertToSeries(seriesId string, itemType string, itemId string, time float64) *requests.InsertToSeries {
	return requests.NewInsertToSeries(c, seriesId, itemType, itemId, time)
}

// NewRemoveFromSeries creates RemoveFromSeries request.
// Removes an existing series item from the series.
func (c *RecombeeClient) NewRemoveFromSeries(seriesId string, itemType string, itemId string) *requests.RemoveFromSeries {
	return requests.NewRemoveFromSeries(c, seriesId, itemType, itemId)
}

// NewAddUser creates AddUser request.
// Adds a new user to the database.
func (c *RecombeeClient) NewAddUser(userId string) *requests.AddUser {
	return requests.NewAddUser(c, userId)
}

// NewDeleteUser creates DeleteUser request.
// Deletes a user of the given *userId* from the database.
// If there are any purchases, ratings, bookmarks, cart additions or detail views made by the user present in the database, they will be deleted in cascade as well.
func (c *RecombeeClient) NewDeleteUser(userId string) *requests.DeleteUser {
	return requests.NewDeleteUser(c, userId)
}

// NewGetUserValues creates GetUserValues request.
// Gets all the current property values of the given user.
func (c *RecombeeClient) NewGetUserValues(userId string) *requests.GetUserValues {
	return requests.NewGetUserValues(c, userId)
}

// NewMergeUsers creates MergeUsers request.
// Merges interactions (purchases, ratings, bookmarks, detail views ...) of two different users under a single user ID. This is especially useful for online e-commerce applications working with anonymous users identified by unique tokens such as the session ID. In such applications, it may often happen that a user owns a persistent account, yet accesses the system anonymously while, e.g., putting items into a shopping cart. At some point in time, such as when the user wishes to confirm the purchase, (s)he logs into the system using his/her username and password. The interactions made under anonymous session ID then become connected with the persistent account, and merging these two becomes desirable.
// Merging happens between two users referred to as the *target* and the *source*. After the merge, all the interactions of the source user are attributed to the target user, and the source user is **deleted**.
// By default, the *Merge Users* request is only available from server-side integrations for security reasons, to prevent potential abuse.
// If you need to call this request from a client-side environment (such as a web or mobile app), please contact our support and request access to enable this feature for your database.
func (c *RecombeeClient) NewMergeUsers(targetUserId string, sourceUserId string) *requests.MergeUsers {
	return requests.NewMergeUsers(c, targetUserId, sourceUserId)
}

// NewListUsers creates ListUsers request.
// Gets a list of IDs of users currently present in the catalog.
func (c *RecombeeClient) NewListUsers() *requests.ListUsers {
	return requests.NewListUsers(c)
}

// NewAddUserProperty creates AddUserProperty request.
// Adding a user property is somewhat equivalent to adding a column to the table of users. The users may be characterized by various properties of different types.
func (c *RecombeeClient) NewAddUserProperty(propertyName string, propertyType string) *requests.AddUserProperty {
	return requests.NewAddUserProperty(c, propertyName, propertyType)
}

// NewDeleteUserProperty creates DeleteUserProperty request.
// Deleting a user property is roughly equivalent to removing a column from the table of users.
func (c *RecombeeClient) NewDeleteUserProperty(propertyName string) *requests.DeleteUserProperty {
	return requests.NewDeleteUserProperty(c, propertyName)
}

// NewGetUserPropertyInfo creates GetUserPropertyInfo request.
// Gets information about specified user property.
func (c *RecombeeClient) NewGetUserPropertyInfo(propertyName string) *requests.GetUserPropertyInfo {
	return requests.NewGetUserPropertyInfo(c, propertyName)
}

// NewListUserProperties creates ListUserProperties request.
// Gets the list of all the user properties in your database.
func (c *RecombeeClient) NewListUserProperties() *requests.ListUserProperties {
	return requests.NewListUserProperties(c)
}

// NewAddDetailView creates AddDetailView request.
// Adds a detail view of the given item made by the given user.
func (c *RecombeeClient) NewAddDetailView(userId string, itemId string) *requests.AddDetailView {
	return requests.NewAddDetailView(c, userId, itemId)
}

// NewDeleteDetailView creates DeleteDetailView request.
// Deletes an existing detail view uniquely specified by (`userId`, `itemId`, and `timestamp`) or all the detail views with the given `userId` and `itemId` if `timestamp` is omitted.
func (c *RecombeeClient) NewDeleteDetailView(userId string, itemId string) *requests.DeleteDetailView {
	return requests.NewDeleteDetailView(c, userId, itemId)
}

// NewListItemDetailViews creates ListItemDetailViews request.
// Lists all the detail views of the given item ever made by different users.
func (c *RecombeeClient) NewListItemDetailViews(itemId string) *requests.ListItemDetailViews {
	return requests.NewListItemDetailViews(c, itemId)
}

// NewListUserDetailViews creates ListUserDetailViews request.
// Lists all the detail views of different items ever made by the given user.
func (c *RecombeeClient) NewListUserDetailViews(userId string) *requests.ListUserDetailViews {
	return requests.NewListUserDetailViews(c, userId)
}

// NewAddPurchase creates AddPurchase request.
// Adds a purchase of the given item made by the given user.
func (c *RecombeeClient) NewAddPurchase(userId string, itemId string) *requests.AddPurchase {
	return requests.NewAddPurchase(c, userId, itemId)
}

// NewDeletePurchase creates DeletePurchase request.
// Deletes an existing purchase uniquely specified by `userId`, `itemId`, and `timestamp` or all the purchases with the given `userId` and `itemId` if `timestamp` is omitted.
func (c *RecombeeClient) NewDeletePurchase(userId string, itemId string) *requests.DeletePurchase {
	return requests.NewDeletePurchase(c, userId, itemId)
}

// NewListItemPurchases creates ListItemPurchases request.
// Lists all the ever-made purchases of the given item.
func (c *RecombeeClient) NewListItemPurchases(itemId string) *requests.ListItemPurchases {
	return requests.NewListItemPurchases(c, itemId)
}

// NewListUserPurchases creates ListUserPurchases request.
// Lists all the purchases ever made by the given user.
func (c *RecombeeClient) NewListUserPurchases(userId string) *requests.ListUserPurchases {
	return requests.NewListUserPurchases(c, userId)
}

// NewAddRating creates AddRating request.
// Adds a rating of the given item made by the given user.
func (c *RecombeeClient) NewAddRating(userId string, itemId string, rating float64) *requests.AddRating {
	return requests.NewAddRating(c, userId, itemId, rating)
}

// NewDeleteRating creates DeleteRating request.
// Deletes an existing rating specified by (`userId`, `itemId`, `timestamp`) from the database or all the ratings with the given `userId` and `itemId` if `timestamp` is omitted.
func (c *RecombeeClient) NewDeleteRating(userId string, itemId string) *requests.DeleteRating {
	return requests.NewDeleteRating(c, userId, itemId)
}

// NewListItemRatings creates ListItemRatings request.
// Lists all the ratings of an item ever submitted by different users.
func (c *RecombeeClient) NewListItemRatings(itemId string) *requests.ListItemRatings {
	return requests.NewListItemRatings(c, itemId)
}

// NewListUserRatings creates ListUserRatings request.
// Lists all the ratings ever submitted by the given user.
func (c *RecombeeClient) NewListUserRatings(userId string) *requests.ListUserRatings {
	return requests.NewListUserRatings(c, userId)
}

// NewAddCartAddition creates AddCartAddition request.
// Adds a cart addition of the given item made by the given user.
func (c *RecombeeClient) NewAddCartAddition(userId string, itemId string) *requests.AddCartAddition {
	return requests.NewAddCartAddition(c, userId, itemId)
}

// NewDeleteCartAddition creates DeleteCartAddition request.
// Deletes an existing cart addition uniquely specified by `userId`, `itemId`, and `timestamp` or all the cart additions with the given `userId` and `itemId` if `timestamp` is omitted.
func (c *RecombeeClient) NewDeleteCartAddition(userId string, itemId string) *requests.DeleteCartAddition {
	return requests.NewDeleteCartAddition(c, userId, itemId)
}

// NewListItemCartAdditions creates ListItemCartAdditions request.
// Lists all the ever-made cart additions of the given item.
func (c *RecombeeClient) NewListItemCartAdditions(itemId string) *requests.ListItemCartAdditions {
	return requests.NewListItemCartAdditions(c, itemId)
}

// NewListUserCartAdditions creates ListUserCartAdditions request.
// Lists all the cart additions ever made by the given user.
func (c *RecombeeClient) NewListUserCartAdditions(userId string) *requests.ListUserCartAdditions {
	return requests.NewListUserCartAdditions(c, userId)
}

// NewAddBookmark creates AddBookmark request.
// Adds a bookmark of the given item made by the given user.
func (c *RecombeeClient) NewAddBookmark(userId string, itemId string) *requests.AddBookmark {
	return requests.NewAddBookmark(c, userId, itemId)
}

// NewDeleteBookmark creates DeleteBookmark request.
// Deletes a bookmark uniquely specified by `userId`, `itemId`, and `timestamp` or all the bookmarks with the given `userId` and `itemId` if `timestamp` is omitted.
func (c *RecombeeClient) NewDeleteBookmark(userId string, itemId string) *requests.DeleteBookmark {
	return requests.NewDeleteBookmark(c, userId, itemId)
}

// NewListItemBookmarks creates ListItemBookmarks request.
// Lists all the ever-made bookmarks of the given item.
func (c *RecombeeClient) NewListItemBookmarks(itemId string) *requests.ListItemBookmarks {
	return requests.NewListItemBookmarks(c, itemId)
}

// NewListUserBookmarks creates ListUserBookmarks request.
// Lists all the bookmarks ever made by the given user.
func (c *RecombeeClient) NewListUserBookmarks(userId string) *requests.ListUserBookmarks {
	return requests.NewListUserBookmarks(c, userId)
}

// NewSetViewPortion creates SetViewPortion request.
// Sets viewed portion of an item (for example a video or article) by a user (at a session).
// If you send a new request with the same (`userId`, `itemId`, `sessionId`), the portion gets updated.
func (c *RecombeeClient) NewSetViewPortion(userId string, itemId string, portion float64) *requests.SetViewPortion {
	return requests.NewSetViewPortion(c, userId, itemId, portion)
}

// NewDeleteViewPortion creates DeleteViewPortion request.
// Deletes an existing view portion specified by (`userId`, `itemId`, `sessionId`) from the database.
func (c *RecombeeClient) NewDeleteViewPortion(userId string, itemId string) *requests.DeleteViewPortion {
	return requests.NewDeleteViewPortion(c, userId, itemId)
}

// NewListItemViewPortions creates ListItemViewPortions request.
// Lists all the view portions of an item ever submitted by different users.
func (c *RecombeeClient) NewListItemViewPortions(itemId string) *requests.ListItemViewPortions {
	return requests.NewListItemViewPortions(c, itemId)
}

// NewListUserViewPortions creates ListUserViewPortions request.
// Lists all the view portions ever submitted by the given user.
func (c *RecombeeClient) NewListUserViewPortions(userId string) *requests.ListUserViewPortions {
	return requests.NewListUserViewPortions(c, userId)
}

// NewRecommendItemsToUser creates RecommendItemsToUser request.
// Based on the user's past interactions (purchases, ratings, etc.) with the items, recommends top-N items that are most likely to be of high value for the given user.
// The most typical use cases are recommendations on the homepage, in some "Picked just for you" section, or in email.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this recommendation was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui#reported-metrics).
// - Get subsequent recommended items when the user scrolls down (*infinite scroll*) or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
func (c *RecombeeClient) NewRecommendItemsToUser(userId string, count int) *requests.RecommendItemsToUser {
	return requests.NewRecommendItemsToUser(c, userId, count)
}

// NewRecommendItemsToItem creates RecommendItemsToItem request.
// Recommends a set of items that are somehow related to one given item, *X*. A typical scenario is when the user *A* is viewing *X*. Then you may display items to the user that he might also be interested in. Recommend items to item request gives you Top-N such items, optionally taking the target user *A* into account.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this recommendation was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui#reported-metrics).
// - Get subsequent recommended items when the user scrolls down (*infinite scroll*) or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
func (c *RecombeeClient) NewRecommendItemsToItem(itemId string, targetUserId string, count int) *requests.RecommendItemsToItem {
	return requests.NewRecommendItemsToItem(c, itemId, targetUserId, count)
}

// NewRecommendItemsToItemSegment creates RecommendItemsToItemSegment request.
// Recommends Items that are the most relevant to a particular Segment from a context [Segmentation](https://docs.recombee.com/segmentations).
// Based on the used Segmentation, this endpoint can be used for example for:
// - Recommending articles related to a particular topic
// - Recommending songs belonging to a particular genre
// - Recommending products produced by a particular brand
// You need to set the used context Segmentation in the Admin UI in the [Scenario settings](https://docs.recombee.com/scenarios) prior to using this endpoint.
// The returned items are sorted by relevance (the first item being the most relevant).
// It is also possible to use the POST HTTP method (for example, in the case of a very long ReQL filter) — query parameters then become body parameters.
func (c *RecombeeClient) NewRecommendItemsToItemSegment(contextSegmentId string, targetUserId string, count int) *requests.RecommendItemsToItemSegment {
	return requests.NewRecommendItemsToItemSegment(c, contextSegmentId, targetUserId, count)
}

// NewRecommendNextItems creates RecommendNextItems request.
// Returns items that shall be shown to a user as next recommendations when the user e.g. scrolls the page down (*infinite scroll*) or goes to the next page.
// It accepts `recommId` of a base recommendation request (e.g., request from the first page) and the number of items that shall be returned (`count`).
// The base request can be one of:
//   - [Recommend Items to Item](https://docs.recombee.com/api#recommend-items-to-item)
//   - [Recommend Items to User](https://docs.recombee.com/api#recommend-items-to-user)
//   - [Recommend Items to Item Segment](https://docs.recombee.com/api#recommend-items-to-item-segment)
//   - [Search Items](https://docs.recombee.com/api#search-items)
//
// All the other parameters are inherited from the base request.
// *Recommend next items* can be called many times for a single `recommId` and each call returns different (previously not recommended) items.
// The number of *Recommend next items* calls performed so far is returned in the `numberNextRecommsCalls` field.
// *Recommend next items* can be requested up to 30 minutes after the base request or a previous *Recommend next items* call.
// For billing purposes, each call to *Recommend next items* is counted as a separate recommendation request.
func (c *RecombeeClient) NewRecommendNextItems(recommId string, count int) *requests.RecommendNextItems {
	return requests.NewRecommendNextItems(c, recommId, count)
}

// NewRecommendUsersToUser creates RecommendUsersToUser request.
// Gets users similar to the given user, based on the user's past interactions (purchases, ratings, etc.) and values of properties.
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
// The returned users are sorted by similarity (the first user being the most similar).
func (c *RecombeeClient) NewRecommendUsersToUser(userId string, count int) *requests.RecommendUsersToUser {
	return requests.NewRecommendUsersToUser(c, userId, count)
}

// NewRecommendUsersToItem creates RecommendUsersToItem request.
// Recommends users that are likely to be interested in the given item.
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
// The returned users are sorted by predicted interest in the item (the first user being the most interested).
func (c *RecombeeClient) NewRecommendUsersToItem(itemId string, count int) *requests.RecommendUsersToItem {
	return requests.NewRecommendUsersToItem(c, itemId, count)
}

// NewRecommendItemSegmentsToUser creates RecommendItemSegmentsToUser request.
// Recommends the top Segments from a [Segmentation](https://docs.recombee.com/segmentations) for a particular user, based on the user's past interactions.
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Recommending the top categories for the user
//   - Recommending the top genres for the user
//   - Recommending the top brands for the user
//   - Recommending the top artists for the user
//
// You need to set the used Segmentation the Admin UI in the [Scenario settings](https://docs.recombee.com/scenarios) prior to using this endpoint.
// The returned segments are sorted by relevance (first segment being the most relevant).
// It is also possible to use POST HTTP method (for example in case of very long ReQL filter) - query parameters then become body parameters.
func (c *RecombeeClient) NewRecommendItemSegmentsToUser(userId string, count int) *requests.RecommendItemSegmentsToUser {
	return requests.NewRecommendItemSegmentsToUser(c, userId, count)
}

// NewRecommendItemSegmentsToItem creates RecommendItemSegmentsToItem request.
// Recommends Segments from a [Segmentation](https://docs.recombee.com/segmentations) that are the most relevant to a particular item.
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Recommending the related categories
//   - Recommending the related genres
//   - Recommending the related brands
//   - Recommending the related artists
//
// You need to set the used Segmentation the Admin UI in the [Scenario settings](https://docs.recombee.com/scenarios) prior to using this endpoint.
// The returned segments are sorted by relevance (first segment being the most relevant).
// It is also possible to use POST HTTP method (for example in case of very long ReQL filter) - query parameters then become body parameters.
func (c *RecombeeClient) NewRecommendItemSegmentsToItem(itemId string, targetUserId string, count int) *requests.RecommendItemSegmentsToItem {
	return requests.NewRecommendItemSegmentsToItem(c, itemId, targetUserId, count)
}

// NewRecommendItemSegmentsToItemSegment creates RecommendItemSegmentsToItemSegment request.
// Recommends Segments from a result [Segmentation](https://docs.recombee.com/segmentations) that are the most relevant to a particular Segment from a context Segmentation.
// Based on the used Segmentations, this endpoint can be used for example for:
//   - Recommending the related brands to particular brand
//   - Recommending the related brands to particular category
//   - Recommending the related artists to a particular genre (assuming songs are the Items)
//
// You need to set the used context and result Segmentation the Admin UI in the [Scenario settings](https://docs.recombee.com/scenarios) prior to using this endpoint.
// The returned segments are sorted by relevance (first segment being the most relevant).
// It is also possible to use POST HTTP method (for example in case of very long ReQL filter) - query parameters then become body parameters.
func (c *RecombeeClient) NewRecommendItemSegmentsToItemSegment(contextSegmentId string, targetUserId string, count int) *requests.RecommendItemSegmentsToItemSegment {
	return requests.NewRecommendItemSegmentsToItemSegment(c, contextSegmentId, targetUserId, count)
}

// NewSearchItems creates SearchItems request.
// Full-text personalized search. The results are based on the provided `searchQuery` and also on the user's past interactions (purchases, ratings, etc.) with the items (items more suitable for the user are preferred in the results).
// All the string and set item properties are indexed by the search engine.
// This endpoint should be used in a search box on your website/app. It can be called multiple times as the user is typing the query in order to get the most viable suggestions based on the current state of the query, or once after submitting the whole query.
// The returned items are sorted by relevance (the first item being the most relevant).
// Besides the recommended items, also a unique `recommId` is returned in the response. It can be used to:
// - Let Recombee know that this search was successful (e.g., user clicked one of the recommended items). See [Reported metrics](https://docs.recombee.com/admin_ui#reported-metrics).
// - Get subsequent search results when the user scrolls down or goes to the next page. See [Recommend Next Items](https://docs.recombee.com/api#recommend-next-items).
// It is also possible to use POST HTTP method (for example in the case of a very long ReQL filter) - query parameters then become body parameters.
func (c *RecombeeClient) NewSearchItems(userId string, searchQuery string, count int) *requests.SearchItems {
	return requests.NewSearchItems(c, userId, searchQuery, count)
}

// NewSearchItemSegments creates SearchItemSegments request.
// Full-text personalized search that returns Segments from a Segmentation. The results are based on the provided `searchQuery` and also on the user's past interactions (purchases, ratings, etc.).
// Based on the used Segmentation, this endpoint can be used for example for:
//   - Searching within categories or brands
//   - Searching within genres or artists
//
// For example if the user is searching for "iPhone" this endpoint can return "cell phones" category.
// You need to set the used Segmentation the Admin UI in the Scenario settings prior to using this endpoint.
// The returned segments are sorted by relevance (first segment being the most relevant).
// It is also possible to use POST HTTP method (for example in case of very long ReQL filter) - query parameters then become body parameters.
func (c *RecombeeClient) NewSearchItemSegments(userId string, searchQuery string, count int) *requests.SearchItemSegments {
	return requests.NewSearchItemSegments(c, userId, searchQuery, count)
}

// NewAddSearchSynonym creates AddSearchSynonym request.
// Adds a new synonym for the [Search items](https://docs.recombee.com/api#search-items).
// When the `term` is used in the search query, the `synonym` is also used for the full-text search.
// Unless `oneWay=true`, it works also in the opposite way (`synonym` -> `term`).
// An example of a synonym can be `science fiction` for the term `sci-fi`.
func (c *RecombeeClient) NewAddSearchSynonym(term string, synonym string) *requests.AddSearchSynonym {
	return requests.NewAddSearchSynonym(c, term, synonym)
}

// NewListSearchSynonyms creates ListSearchSynonyms request.
// Gives the list of synonyms defined in the database.
func (c *RecombeeClient) NewListSearchSynonyms() *requests.ListSearchSynonyms {
	return requests.NewListSearchSynonyms(c)
}

// NewDeleteAllSearchSynonyms creates DeleteAllSearchSynonyms request.
// Deletes all synonyms defined in the database.
func (c *RecombeeClient) NewDeleteAllSearchSynonyms() *requests.DeleteAllSearchSynonyms {
	return requests.NewDeleteAllSearchSynonyms(c)
}

// NewDeleteSearchSynonym creates DeleteSearchSynonym request.
// Deletes synonym of the given `id`. This synonym is no longer taken into account in the [Search items](https://docs.recombee.com/api#search-items).
func (c *RecombeeClient) NewDeleteSearchSynonym(id string) *requests.DeleteSearchSynonym {
	return requests.NewDeleteSearchSynonym(c, id)
}

// NewCreatePropertyBasedSegmentation creates CreatePropertyBasedSegmentation request.
// Creates a Segmentation that splits the items into segments based on values of a particular item property.
// A segment is created for each unique value of the property.
// In case of `set` properties, a segment is created for each value in the set. Item belongs to all these segments.
func (c *RecombeeClient) NewCreatePropertyBasedSegmentation(segmentationId string, sourceType string, propertyName string) *requests.CreatePropertyBasedSegmentation {
	return requests.NewCreatePropertyBasedSegmentation(c, segmentationId, sourceType, propertyName)
}

// NewUpdatePropertyBasedSegmentation creates UpdatePropertyBasedSegmentation request.
// Updates a Property Based Segmentation
func (c *RecombeeClient) NewUpdatePropertyBasedSegmentation(segmentationId string) *requests.UpdatePropertyBasedSegmentation {
	return requests.NewUpdatePropertyBasedSegmentation(c, segmentationId)
}

// NewCreateAutoReqlSegmentation creates CreateAutoReqlSegmentation request.
// Segment the items using a [ReQL](https://docs.recombee.com/reql) expression.
// For each item, the expression should return a set that contains IDs of segments to which the item belongs to.
func (c *RecombeeClient) NewCreateAutoReqlSegmentation(segmentationId string, sourceType string, expression string) *requests.CreateAutoReqlSegmentation {
	return requests.NewCreateAutoReqlSegmentation(c, segmentationId, sourceType, expression)
}

// NewUpdateAutoReqlSegmentation creates UpdateAutoReqlSegmentation request.
// Update an existing Segmentation.
func (c *RecombeeClient) NewUpdateAutoReqlSegmentation(segmentationId string) *requests.UpdateAutoReqlSegmentation {
	return requests.NewUpdateAutoReqlSegmentation(c, segmentationId)
}

// NewCreateManualReqlSegmentation creates CreateManualReqlSegmentation request.
// Segment the items using multiple [ReQL](https://docs.recombee.com/reql) filters.
// Use the Add Manual ReQL Items Segment endpoint to create the individual segments.
func (c *RecombeeClient) NewCreateManualReqlSegmentation(segmentationId string, sourceType string) *requests.CreateManualReqlSegmentation {
	return requests.NewCreateManualReqlSegmentation(c, segmentationId, sourceType)
}

// NewUpdateManualReqlSegmentation creates UpdateManualReqlSegmentation request.
// Update an existing Segmentation.
func (c *RecombeeClient) NewUpdateManualReqlSegmentation(segmentationId string) *requests.UpdateManualReqlSegmentation {
	return requests.NewUpdateManualReqlSegmentation(c, segmentationId)
}

// NewAddManualReqlSegment creates AddManualReqlSegment request.
// Adds a new Segment into a Manual ReQL Segmentation.
// The new Segment is defined by a [ReQL](https://docs.recombee.com/reql) filter that returns `true` for an item in case that this item belongs to the segment.
func (c *RecombeeClient) NewAddManualReqlSegment(segmentationId string, segmentId string, filter string) *requests.AddManualReqlSegment {
	return requests.NewAddManualReqlSegment(c, segmentationId, segmentId, filter)
}

// NewUpdateManualReqlSegment creates UpdateManualReqlSegment request.
// Update definition of the Segment.
func (c *RecombeeClient) NewUpdateManualReqlSegment(segmentationId string, segmentId string, filter string) *requests.UpdateManualReqlSegment {
	return requests.NewUpdateManualReqlSegment(c, segmentationId, segmentId, filter)
}

// NewDeleteManualReqlSegment creates DeleteManualReqlSegment request.
// Delete a Segment from a Manual ReQL Segmentation.
func (c *RecombeeClient) NewDeleteManualReqlSegment(segmentationId string, segmentId string) *requests.DeleteManualReqlSegment {
	return requests.NewDeleteManualReqlSegment(c, segmentationId, segmentId)
}

// NewListSegmentations creates ListSegmentations request.
// Return all existing items Segmentations.
func (c *RecombeeClient) NewListSegmentations(sourceType string) *requests.ListSegmentations {
	return requests.NewListSegmentations(c, sourceType)
}

// NewGetSegmentation creates GetSegmentation request.
// Get existing Segmentation.
func (c *RecombeeClient) NewGetSegmentation(segmentationId string) *requests.GetSegmentation {
	return requests.NewGetSegmentation(c, segmentationId)
}

// NewDeleteSegmentation creates DeleteSegmentation request.
// Delete existing Segmentation.
func (c *RecombeeClient) NewDeleteSegmentation(segmentationId string) *requests.DeleteSegmentation {
	return requests.NewDeleteSegmentation(c, segmentationId)
}

// NewListScenarios creates ListScenarios request.
// Get all [Scenarios](https://docs.recombee.com/scenarios) of the given database.
func (c *RecombeeClient) NewListScenarios() *requests.ListScenarios {
	return requests.NewListScenarios(c)
}

// NewResetDatabase creates ResetDatabase request.
// Completely erases all your data, including items, item properties, series, user database, purchases, ratings, detail views, and bookmarks. Make sure the request is never executed in the production environment! Resetting your database is irreversible.
func (c *RecombeeClient) NewResetDatabase() *requests.ResetDatabase {
	return requests.NewResetDatabase(c)
}
