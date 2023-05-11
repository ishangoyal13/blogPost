package routers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ishangoyal13/blogPost/pkg/config"
)

var (
	serverState ServerState
)

func runServer(app config.App) {

	server := &http.Server{
		Addr:    ":" + app.Config.Server.Port,
		Handler: app.Router,
	}

	graceful := Graceful{
		Server:          server,
		ShutdownTimeout: time.Duration(5 * time.Second),
		State:           &serverState,
	}
	displayService()
	graceful.ListenAndServe()
}

// You can generate ASCI art here
// https://patorjk.com/software/taag/#p=display&f=ANSI%20Regular&t=BLOG%20POST
func displayService() {
	fmt.Println(`

	██████  ██       ██████   ██████      ██████   ██████  ███████ ████████ 
	██   ██ ██      ██    ██ ██           ██   ██ ██    ██ ██         ██    
	██████  ██      ██    ██ ██   ███     ██████  ██    ██ ███████    ██    
	██   ██ ██      ██    ██ ██    ██     ██      ██    ██      ██    ██    
	██████  ███████  ██████   ██████      ██       ██████  ███████    ██    
                                                                                                                                                                                                  
	`)
}
