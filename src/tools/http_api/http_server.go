package http_api

import (
	"net"
	"net/http"
	"strings"

	"fmt"
	. "tools"
)

func Serve(listener net.Listener, handler http.Handler, proto string) {
	INFO(fmt.Sprintf("%s: listening on %s", proto, listener.Addr()))

	server := &http.Server{
		Handler: handler,
	}
	err := server.Serve(listener)

	if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
		ERR("ERROR: http.Serve() - %s", err)
	}

	INFO("%s: closing %s", proto, listener.Addr())
}
