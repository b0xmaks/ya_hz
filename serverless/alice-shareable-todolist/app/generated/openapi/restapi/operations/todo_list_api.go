// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewTodoListAPI creates a new TodoList instance
func NewTodoListAPI(spec *loads.Document) *TodoListAPI {
	return &TodoListAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		AcceptInvitationHandler: AcceptInvitationHandlerFunc(func(params AcceptInvitationParams) middleware.Responder {
			return middleware.NotImplemented("operation AcceptInvitation has not yet been implemented")
		}),
		AddItemHandler: AddItemHandlerFunc(func(params AddItemParams) middleware.Responder {
			return middleware.NotImplemented("operation AddItem has not yet been implemented")
		}),
		CreateListHandler: CreateListHandlerFunc(func(params CreateListParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateList has not yet been implemented")
		}),
		DeleteItemHandler: DeleteItemHandlerFunc(func(params DeleteItemParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteItem has not yet been implemented")
		}),
		DeleteListHandler: DeleteListHandlerFunc(func(params DeleteListParams) middleware.Responder {
			return middleware.NotImplemented("operation DeleteList has not yet been implemented")
		}),
		GetListHandler: GetListHandlerFunc(func(params GetListParams) middleware.Responder {
			return middleware.NotImplemented("operation GetList has not yet been implemented")
		}),
		GetListUsersHandler: GetListUsersHandlerFunc(func(params GetListUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation GetListUsers has not yet been implemented")
		}),
		InviteUserHandler: InviteUserHandlerFunc(func(params InviteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation InviteUser has not yet been implemented")
		}),
		ListListsHandler: ListListsHandlerFunc(func(params ListListsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListLists has not yet been implemented")
		}),
		PageLoginHandler: PageLoginHandlerFunc(func(params PageLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation PageLogin has not yet been implemented")
		}),
		PageReceiveTokenHandler: PageReceiveTokenHandlerFunc(func(params PageReceiveTokenParams) middleware.Responder {
			return middleware.NotImplemented("operation PageReceiveToken has not yet been implemented")
		}),
		RevokeInvitationHandler: RevokeInvitationHandlerFunc(func(params RevokeInvitationParams) middleware.Responder {
			return middleware.NotImplemented("operation RevokeInvitation has not yet been implemented")
		}),
		UserInfoHandler: UserInfoHandlerFunc(func(params UserInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation UserInfo has not yet been implemented")
		}),
	}
}

/*TodoListAPI the todo list API */
type TodoListAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// AcceptInvitationHandler sets the operation handler for the accept invitation operation
	AcceptInvitationHandler AcceptInvitationHandler
	// AddItemHandler sets the operation handler for the add item operation
	AddItemHandler AddItemHandler
	// CreateListHandler sets the operation handler for the create list operation
	CreateListHandler CreateListHandler
	// DeleteItemHandler sets the operation handler for the delete item operation
	DeleteItemHandler DeleteItemHandler
	// DeleteListHandler sets the operation handler for the delete list operation
	DeleteListHandler DeleteListHandler
	// GetListHandler sets the operation handler for the get list operation
	GetListHandler GetListHandler
	// GetListUsersHandler sets the operation handler for the get list users operation
	GetListUsersHandler GetListUsersHandler
	// InviteUserHandler sets the operation handler for the invite user operation
	InviteUserHandler InviteUserHandler
	// ListListsHandler sets the operation handler for the list lists operation
	ListListsHandler ListListsHandler
	// PageLoginHandler sets the operation handler for the page login operation
	PageLoginHandler PageLoginHandler
	// PageReceiveTokenHandler sets the operation handler for the page receive token operation
	PageReceiveTokenHandler PageReceiveTokenHandler
	// RevokeInvitationHandler sets the operation handler for the revoke invitation operation
	RevokeInvitationHandler RevokeInvitationHandler
	// UserInfoHandler sets the operation handler for the user info operation
	UserInfoHandler UserInfoHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *TodoListAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *TodoListAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *TodoListAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *TodoListAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *TodoListAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *TodoListAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *TodoListAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *TodoListAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *TodoListAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the TodoListAPI
func (o *TodoListAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AcceptInvitationHandler == nil {
		unregistered = append(unregistered, "AcceptInvitationHandler")
	}
	if o.AddItemHandler == nil {
		unregistered = append(unregistered, "AddItemHandler")
	}
	if o.CreateListHandler == nil {
		unregistered = append(unregistered, "CreateListHandler")
	}
	if o.DeleteItemHandler == nil {
		unregistered = append(unregistered, "DeleteItemHandler")
	}
	if o.DeleteListHandler == nil {
		unregistered = append(unregistered, "DeleteListHandler")
	}
	if o.GetListHandler == nil {
		unregistered = append(unregistered, "GetListHandler")
	}
	if o.GetListUsersHandler == nil {
		unregistered = append(unregistered, "GetListUsersHandler")
	}
	if o.InviteUserHandler == nil {
		unregistered = append(unregistered, "InviteUserHandler")
	}
	if o.ListListsHandler == nil {
		unregistered = append(unregistered, "ListListsHandler")
	}
	if o.PageLoginHandler == nil {
		unregistered = append(unregistered, "PageLoginHandler")
	}
	if o.PageReceiveTokenHandler == nil {
		unregistered = append(unregistered, "PageReceiveTokenHandler")
	}
	if o.RevokeInvitationHandler == nil {
		unregistered = append(unregistered, "RevokeInvitationHandler")
	}
	if o.UserInfoHandler == nil {
		unregistered = append(unregistered, "UserInfoHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *TodoListAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *TodoListAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *TodoListAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *TodoListAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *TodoListAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *TodoListAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the todo list API
func (o *TodoListAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *TodoListAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/todo-lists/{list_id}/accept_invitation"] = NewAcceptInvitation(o.context, o.AcceptInvitationHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/todo-lists/{list_id}/items"] = NewAddItem(o.context, o.AddItemHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/todo-lists"] = NewCreateList(o.context, o.CreateListHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/api/todo-lists/{list_id}/items/{item_id}/wa"] = NewDeleteItem(o.context, o.DeleteItemHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/api/todo-lists/{list_id}/wa"] = NewDeleteList(o.context, o.DeleteListHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/todo-lists/{list_id}/wa"] = NewGetList(o.context, o.GetListHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/todo-lists/{list_id}/users"] = NewGetListUsers(o.context, o.GetListUsersHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/todo-lists/{list_id}/inviteUser"] = NewInviteUser(o.context, o.InviteUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/todo-lists"] = NewListLists(o.context, o.ListListsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/login"] = NewPageLogin(o.context, o.PageLoginHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/receive-token"] = NewPageReceiveToken(o.context, o.PageReceiveTokenHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/api/todo-lists/{list_id}/users/{user_id}/wa"] = NewRevokeInvitation(o.context, o.RevokeInvitationHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/me"] = NewUserInfo(o.context, o.UserInfoHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *TodoListAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *TodoListAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *TodoListAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *TodoListAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *TodoListAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
