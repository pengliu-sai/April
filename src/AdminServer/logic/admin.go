package logic

import (
	"github.com/julienschmidt/httprouter"
	"github.com/serialx/hashring"
	"net"
	"net/http"
	"os"
	"sync"
	. "tools"
	"tools/cfg"
	"tools/db"
	"tools/http_api"
	"tools/redis"
	"tools/util"
)

var (
	waitGroup util.WaitGroupWrapper
	_lock     sync.Mutex
)

var (
	gameServers     []string
	gameServersRing *hashring.HashRing

	g_userID_GameServerAddr map[uint64]string

	developMode bool
)

func init() {

	developMode = cfg.IsDevelopMode()

	loadGameServers()

	g_userID_GameServerAddr = make(map[uint64]string)
}

func InitDB() {
	err := db.Connect("admin_db")
	checkError(err)
}

func InitRedis() {
	err := redis.Connect("admin_redis")
	checkError(err)
}

func StartAdminServer() {
	if developMode {
		WARN("admin server in develop mode")
	}

	InitDB()

	InitRedis()

	adminPort := cfg.GetServerConfig().AdminServer.Port
	if adminPort == "" {
		ERR("admin port not setting")
		os.Exit(1)
	}

	adminAddr := "127.0.0.1:" + adminPort

	httpListener, err := net.Listen("tcp", adminAddr)
	if err != nil {
		os.Exit(1)
	}

	httpServer := newHttpServer()
	waitGroup.Wrap(func() {
		http_api.Serve(httpListener, httpServer, "AdminServer")
	})

	waitGroup.Wait()
}

func checkError(err error) {
	if err != nil {
		ERR("Fatal error: %v", err)
		os.Exit(-1)
	}
}

func loadGameServers() {
	serverList := cfg.GetServerConfig().GameServerList
	serverLen := len(serverList)
	gameServers = make([]string, serverLen)

	for i := 0; i < serverLen; i++ {
		gameServers[i] = serverList[i].IP + ":" + serverList[i].Port
	}

	gameServersRing = hashring.New(gameServers)
}

type httpServer struct {
	router http.Handler
}

func newHttpServer() *httpServer {
	router := httprouter.New()
	router.HandleMethodNotAllowed = true
	router.PanicHandler = http_api.LogPanicHandler()
	router.NotFound = http_api.LogNotFoundHandler()
	router.MethodNotAllowed = http_api.LogMethodNotAllowedHandler()
	s := &httpServer{
		router: router,
	}

	router.Handle("GET", "/", http_api.Decorate(s.doPing, http_api.V1))

	/*-------------[[From Gate]]-----------------------------------------*/
	router.Handle("POST", "/gate/userLoginByUserName", http_api.Decorate(s.doGate_UserLoginByUserName, http_api.V1))
	router.Handle("POST", "/gate/userLoginByUserID", http_api.Decorate(s.doGate_UserLoginByUserID, http_api.V1))
	router.Handle("POST", "/gate/userRegister", http_api.Decorate(s.doGate_UserRegister, http_api.V1))
	router.Handle("GET", "/gate/userExit", http_api.Decorate(s.doGate_UserExit, http_api.V1))

	/*-------------[[From Admin]]----------------------------------------*/
	router.Handle("GET", "/admin/userInfo", http_api.Decorate(s.doAdmin_UserInfo, http_api.V1))
	router.Handle("DELETE", "/admin/userDelete", http_api.Decorate(s.doAdmin_UserDelete, http_api.V1))
	router.Handle("PUT", "/admin/userMarkDeleted", http_api.Decorate(s.doAdmin_UserMarkDeleted, http_api.V1))
	router.Handle("PUT", "/admin/userRebirth", http_api.Decorate(s.doAdmin_UserRebirth, http_api.V1))
	router.Handle("GET", "/admin/gameServerOffline", http_api.Decorate(s.doAdmin_GameServerOffline, http_api.V1))

	/*-------------[[From Game]]----------------------------------------*/
	//router.Handle("PUT", "/game/userCheckLogin", http_api.Decorate(s.doGame_UserCheckLogin, http_api.V1))
	router.Handle("GET", "/game/userOffline", http_api.Decorate(s.doGame_UserOffline, http_api.V1))
	router.Handle("GET", "/game/gameServerOffline", http_api.Decorate(s.doGame_GameServerOffline, http_api.V1))

	/*------------[[From Other Admin]]--------------------------------------------*/
	router.Handle("GET", "/otheradmin/userOffline", http_api.Decorate(s.doOtherAdmin_UserOffline, http_api.V1))
	return s
}

func (s *httpServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *httpServer) doPing(w http.ResponseWriter, req *http.Request, ps httprouter.Params) (interface{}, error) {
	reqParams, err := http_api.NewReqParams(req)

	if err != nil {
		return nil, http_api.Err{400, "INVALID_REQUEST"}
	}

	_, err = reqParams.Get("ping")
	if err != nil {
		return nil, http_api.Err{400, "MISSING_ARG_PING"}
	}

	return "pong", nil
}
