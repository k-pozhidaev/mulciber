package web

import (
	"fmt"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"net/http"
)

type Context struct {
	manager *manage.Manager
	server *server.Server
	clientStore *store.ClientStore
	port uint16
}

var context *Context

/**
@param port int16
*/
func InitContext(port uint16)  {
	manager, clientStore := initManager()
	srv := initServer(manager)
	context = &Context{
		manager: manager,
		clientStore: clientStore,
		server: srv,
		port: port,
	}
}

func GetStringPort() string {
	return fmt.Sprintf(":%d", context.port)
}

func ServerHandleTokenRequest(w http.ResponseWriter, r *http.Request) error {
	return context.server.HandleTokenRequest(w, r)
}

func ServerValidationBearerToken(r *http.Request) (oauth2.TokenInfo, error) {
	return context.server.ValidationBearerToken(r)
}

func ClientStoreSet(id string, clientSecret string) error {
	return context.clientStore.Set(id, &models.Client{
		ID:     id,
		Secret: clientSecret,
		Domain: "http://localhost:9094", // TODO hardcode
	})
}

func initManager() (manager *manage.Manager, clientStore *store.ClientStore ) {
	manager = manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore = store.NewClientStore()

	manager.MapClientStorage(clientStore)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)
	return
}

func initServer(manager *manage.Manager) (srv *server.Server) {
	srv = server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})
	return
}