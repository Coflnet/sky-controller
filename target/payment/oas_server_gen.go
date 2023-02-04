// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
	"time"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// ApplyPost implements POST /Apply operation.
	//
	// Brings all products, groups and roles into the given state
	// will disable/delete anything not present so use carefully.
	//
	// POST /Apply
	ApplyPost(ctx context.Context, req *SystemState) error
	// CallbackPaypalPost implements POST /Callback/paypal operation.
	//
	// Accept callbacks from paypal.
	//
	// POST /Callback/paypal
	CallbackPaypalPost(ctx context.Context) error
	// CallbackStripePost implements POST /Callback/stripe operation.
	//
	// Webhook callback for stripe.
	//
	// POST /Callback/stripe
	CallbackStripePost(ctx context.Context) error
	// GroupGet implements GET /Group operation.
	//
	// GET /Group
	GroupGet(ctx context.Context, params GroupGetParams) ([]Group, error)
	// GroupGroupSlugDelete implements DELETE /Group/{groupSlug} operation.
	//
	// DELETE /Group/{groupSlug}
	GroupGroupSlugDelete(ctx context.Context, params GroupGroupSlugDeleteParams) (*Group, error)
	// GroupGroupSlugGet implements GET /Group/{groupSlug} operation.
	//
	// GET /Group/{groupSlug}
	GroupGroupSlugGet(ctx context.Context, params GroupGroupSlugGetParams) (*Group, error)
	// GroupGroupSlugProductsDelete implements DELETE /Group/{groupSlug}/products operation.
	//
	// DELETE /Group/{groupSlug}/products
	GroupGroupSlugProductsDelete(ctx context.Context, req []string, params GroupGroupSlugProductsDeleteParams) (*Group, error)
	// GroupGroupSlugProductsPost implements POST /Group/{groupSlug}/products operation.
	//
	// POST /Group/{groupSlug}/products
	GroupGroupSlugProductsPost(ctx context.Context, req []string, params GroupGroupSlugProductsPostParams) (*Group, error)
	// GroupGroupSlugPut implements PUT /Group/{groupSlug} operation.
	//
	// PUT /Group/{groupSlug}
	GroupGroupSlugPut(ctx context.Context, req *Group, params GroupGroupSlugPutParams) (*Group, error)
	// GroupPost implements POST /Group operation.
	//
	// POST /Group
	GroupPost(ctx context.Context, req *Group) (*Group, error)
	// ProductsGet implements GET /Products operation.
	//
	// Get all products.
	//
	// GET /Products
	ProductsGet(ctx context.Context, params ProductsGetParams) ([]PurchaseableProduct, error)
	// ProductsPProductSlugGet implements GET /Products/p/{productSlug} operation.
	//
	// Get the details of a product.
	//
	// GET /Products/p/{productSlug}
	ProductsPProductSlugGet(ctx context.Context, params ProductsPProductSlugGetParams) (*PurchaseableProduct, error)
	// ProductsPut implements PUT /Products operation.
	//
	// Updates a product by replacing it with a new one.
	// Old products can not be deleted to furfill accounting needs.
	//
	// PUT /Products
	ProductsPut(ctx context.Context, req *PurchaseableProduct) (*PurchaseableProduct, error)
	// ProductsServiceServiceSlugCountGet implements GET /Products/service/{serviceSlug}/count operation.
	//
	// Count of users owning a service.
	//
	// GET /Products/service/{serviceSlug}/count
	ProductsServiceServiceSlugCountGet(ctx context.Context, params ProductsServiceServiceSlugCountGetParams) (int32, error)
	// ProductsServiceServiceSlugIdsGet implements GET /Products/service/{serviceSlug}/ids operation.
	//
	// Gets all userIds owning a service.
	//
	// GET /Products/service/{serviceSlug}/ids
	ProductsServiceServiceSlugIdsGet(ctx context.Context, params ProductsServiceServiceSlugIdsGetParams) ([]string, error)
	// ProductsServicesGet implements GET /Products/services operation.
	//
	// Get services.
	//
	// GET /Products/services
	ProductsServicesGet(ctx context.Context, params ProductsServicesGetParams) ([]PurchaseableProduct, error)
	// ProductsTopupGet implements GET /Products/topup operation.
	//
	// Get topup options.
	//
	// GET /Products/topup
	ProductsTopupGet(ctx context.Context, params ProductsTopupGetParams) ([]TopUpProduct, error)
	// ProductsTopupPut implements PUT /Products/topup operation.
	//
	// Updates a topup option by replacing it with a new one.
	// Old options get a new slug and are marked as disabled.
	//
	// PUT /Products/topup
	ProductsTopupPut(ctx context.Context, req *TopUpProduct) (*TopUpProduct, error)
	// ProductsUserUserIdGet implements GET /Products/user/{userId} operation.
	//
	// Get adjusted prices for a user.
	//
	// GET /Products/user/{userId}
	ProductsUserUserIdGet(ctx context.Context, params ProductsUserUserIdGetParams) ([]RuleResult, error)
	// RulesGet implements GET /Rules operation.
	//
	// Returns all rules.
	//
	// GET /Rules
	RulesGet(ctx context.Context, params RulesGetParams) ([]Rule, error)
	// RulesPost implements POST /Rules operation.
	//
	// Creates a new rule.
	//
	// POST /Rules
	RulesPost(ctx context.Context, req *RuleCreate) (*Rule, error)
	// RulesRuleSlugDelete implements DELETE /Rules/{ruleSlug} operation.
	//
	// Deletes a rule.
	//
	// DELETE /Rules/{ruleSlug}
	RulesRuleSlugDelete(ctx context.Context, params RulesRuleSlugDeleteParams) (*Rule, error)
	// RulesRuleSlugGet implements GET /Rules/{ruleSlug} operation.
	//
	// Returns a rule by slug.
	//
	// GET /Rules/{ruleSlug}
	RulesRuleSlugGet(ctx context.Context, params RulesRuleSlugGetParams) (*Rule, error)
	// TopUpCompensatePost implements POST /TopUp/compensate operation.
	//
	// Compensates users of a service for something.
	//
	// POST /TopUp/compensate
	TopUpCompensatePost(ctx context.Context, req *Compensation) (*Int32Int32ValueTuple, error)
	// TopUpCustomPost implements POST /TopUp/custom operation.
	//
	// Creates a custom topup that is instantly credited.
	//
	// POST /TopUp/custom
	TopUpCustomPost(ctx context.Context, req *CustomTopUp, params TopUpCustomPostParams) (*TopUpIdResponse, error)
	// TopUpOptionsGet implements GET /TopUp/options operation.
	//
	// All available topup options.
	//
	// GET /TopUp/options
	TopUpOptionsGet(ctx context.Context) ([]TopUpProduct, error)
	// TopUpPaypalPost implements POST /TopUp/paypal operation.
	//
	// Creates a payment session with paypal.
	//
	// POST /TopUp/paypal
	TopUpPaypalPost(ctx context.Context, req *TopUpOptions, params TopUpPaypalPostParams) (*TopUpIdResponse, error)
	// TopUpStripePost implements POST /TopUp/stripe operation.
	//
	// Creates a payment session with stripe.
	//
	// POST /TopUp/stripe
	TopUpStripePost(ctx context.Context, req *TopUpOptions, params TopUpStripePostParams) (*TopUpIdResponse, error)
	// TransactionPlanedUUserIdGet implements GET /Transaction/planed/u/{userId} operation.
	//
	// GET /Transaction/planed/u/{userId}
	TransactionPlanedUUserIdGet(ctx context.Context, params TransactionPlanedUUserIdGetParams) ([]ExternalTransaction, error)
	// TransactionPlanedUUserIdPost implements POST /Transaction/planed/u/{userId} operation.
	//
	// POST /Transaction/planed/u/{userId}
	TransactionPlanedUUserIdPost(ctx context.Context, req *ExternalTransaction, params TransactionPlanedUUserIdPostParams) (*PlanedTransaction, error)
	// TransactionPlanedUUserIdTTransactionIdDelete implements DELETE /Transaction/planed/u/{userId}/t/{transactionId} operation.
	//
	// DELETE /Transaction/planed/u/{userId}/t/{transactionId}
	TransactionPlanedUUserIdTTransactionIdDelete(ctx context.Context, params TransactionPlanedUUserIdTTransactionIdDeleteParams) (*PlanedTransaction, error)
	// TransactionPlanedUUserIdTTransactionIdPut implements PUT /Transaction/planed/u/{userId}/t/{transactionId} operation.
	//
	// PUT /Transaction/planed/u/{userId}/t/{transactionId}
	TransactionPlanedUUserIdTTransactionIdPut(ctx context.Context, req *ExternalTransaction, params TransactionPlanedUUserIdTTransactionIdPutParams) (*PlanedTransaction, error)
	// TransactionSendPost implements POST /Transaction/send operation.
	//
	// POST /Transaction/send
	TransactionSendPost(ctx context.Context, req *TransactionEvent) error
	// TransactionUUserIdGet implements GET /Transaction/u/{userId} operation.
	//
	// GET /Transaction/u/{userId}
	TransactionUUserIdGet(ctx context.Context, params TransactionUUserIdGetParams) ([]ExternalTransaction, error)
	// UserUserIdGet implements GET /User/{userId} operation.
	//
	// Gets the user with the given id.
	//
	// GET /User/{userId}
	UserUserIdGet(ctx context.Context, params UserUserIdGetParams) (*User, error)
	// UserUserIdOwnsLongestPost implements POST /User/{userId}/owns/longest operation.
	//
	// Returns the bigest time out of a list of product ids.
	//
	// POST /User/{userId}/owns/longest
	UserUserIdOwnsLongestPost(ctx context.Context, req []string, params UserUserIdOwnsLongestPostParams) (time.Time, error)
	// UserUserIdOwnsPost implements POST /User/{userId}/owns operation.
	//
	// Returns all ownership data for an user out of a list of interested.
	//
	// Deprecated: schema marks this operation as deprecated.
	//
	// POST /User/{userId}/owns
	UserUserIdOwnsPost(ctx context.Context, req []string, params UserUserIdOwnsPostParams) ([]OwnerShip, error)
	// UserUserIdOwnsProductSlugUntilGet implements GET /User/{userId}/owns/{productSlug}/until operation.
	//
	// Returns the time for how long a user owns a given product.
	//
	// GET /User/{userId}/owns/{productSlug}/until
	UserUserIdOwnsProductSlugUntilGet(ctx context.Context, params UserUserIdOwnsProductSlugUntilGetParams) (time.Time, error)
	// UserUserIdOwnsUntilPost implements POST /User/{userId}/owns/until operation.
	//
	// Returns all ownership data for an user out of a list of interested.
	//
	// POST /User/{userId}/owns/until
	UserUserIdOwnsUntilPost(ctx context.Context, req []string, params UserUserIdOwnsUntilPostParams) (UserUserIdOwnsUntilPostOKApplicationJSON, error)
	// UserUserIdPost implements POST /User/{userId} operation.
	//
	// Creates a new user with the given id.
	//
	// POST /User/{userId}
	UserUserIdPost(ctx context.Context, params UserUserIdPostParams) (*User, error)
	// UserUserIdPurchaseProductSlugPost implements POST /User/{userId}/purchase/{productSlug} operation.
	//
	// Purchase a new product if enough funds are available.
	//
	// POST /User/{userId}/purchase/{productSlug}
	UserUserIdPurchaseProductSlugPost(ctx context.Context, params UserUserIdPurchaseProductSlugPostParams) (*User, error)
	// UserUserIdServicePurchaseProductSlugPost implements POST /User/{userId}/service/purchase/{productSlug} operation.
	//
	// Purchase/extends a service if enough funds are available.
	//
	// POST /User/{userId}/service/purchase/{productSlug}
	UserUserIdServicePurchaseProductSlugPost(ctx context.Context, params UserUserIdServicePurchaseProductSlugPostParams) (*User, error)
	// UserUserIdTransactionIdDelete implements DELETE /User/{userId}/{transactionId} operation.
	//
	// Undo the purchase of a service.
	//
	// DELETE /User/{userId}/{transactionId}
	UserUserIdTransactionIdDelete(ctx context.Context, params UserUserIdTransactionIdDeleteParams) (*User, error)
	// UserUserIdTransferPost implements POST /User/{userId}/transfer operation.
	//
	// Transfers coins to another user.
	//
	// POST /User/{userId}/transfer
	UserUserIdTransferPost(ctx context.Context, req *TransferRequest, params UserUserIdTransferPostParams) (*TransactionEvent, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}