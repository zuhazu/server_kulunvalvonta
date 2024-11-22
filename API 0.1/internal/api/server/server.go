package server

import (
	"context"
	"goapi/internal/api/handlers/data"
	"goapi/internal/api/handlers/person"
	"goapi/internal/api/handlers/room"
	"goapi/internal/api/middleware"
	"goapi/internal/api/service"
	"log"
	"net/http"
)

type Server struct {
	ctx        context.Context
	HTTPServer *http.Server
	logger     *log.Logger
}

func NewServer(ctx context.Context, sf *service.ServiceFactory, logger *log.Logger) *Server {

	mux := http.NewServeMux()
	err := setupDataHandlers(mux, sf, logger)
	if err != nil {
		logger.Fatalf("Error setting up data handlers: %v", err)
	}

	middlewares := []middleware.Middleware{
		middleware.BasicAuthenticationMiddleware,
		middleware.CommonMiddleware,
	}

	return &Server{
		ctx:    ctx,
		logger: logger,
		HTTPServer: &http.Server{
			Handler: middleware.ChainMiddleware(mux, middlewares...),
		},
	}
}

func (api *Server) Shutdown() error {
	api.logger.Println("Gracefully shutting down server...")
	return api.HTTPServer.Shutdown(api.ctx)
}

func (api *Server) ListenAndServe(addr string) error {
	api.HTTPServer.Addr = addr
	return api.HTTPServer.ListenAndServe()
}

// * REST API handlers
func setupDataHandlers(mux *http.ServeMux, sf *service.ServiceFactory, logger *log.Logger) error {

	/*ds, err := sf.CreateDataService(service.SQLiteDataService)
	if err != nil {
		return err
	}*/

	//lisää
	personService, err := sf.CreatePersonService(service.SQLitePersonService)
	roomService, err := sf.CreateRoomService(service.SQLiteRoomService)
	if err != nil {
		return err
	}

	mux.HandleFunc("OPTIONS /api/*", func(w http.ResponseWriter, r *http.Request) {
		data.OptionsHandler(w, r)
	})
	mux.HandleFunc("POST /api/person", func(w http.ResponseWriter, r *http.Request) {
		person.PostPersonHandler(w, r, logger, personService)
	})
	mux.HandleFunc("PUT /api/person", func(w http.ResponseWriter, r *http.Request) {
		person.PutPersonHandler(w, r, logger, personService)
	})
	mux.HandleFunc("GET /api/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		person.GetPersonByIDHandler(w, r, logger, personService)
	})
	mux.HandleFunc("DELETE /api/person/{id}", func(w http.ResponseWriter, r *http.Request) {
		person.DeletePersonHandler(w, r, logger, personService)
	})
	mux.HandleFunc("PUT /api/tag", func(w http.ResponseWriter, r *http.Request) {
		person.PostTagHandler(w, r, logger, personService)
	})
	mux.HandleFunc("POST /api/room", func(w http.ResponseWriter, r *http.Request) {
		room.PostRoomHandler(w, r, logger, roomService)
	})
	mux.HandleFunc("GET /api/room/{id}", func(w http.ResponseWriter, r *http.Request) {
		person.GetPersonsByRoomIdHandler(w, r, logger, personService)
	})

	mux.HandleFunc("GET /main", func(w http.ResponseWriter, r *http.Request) {
		room.GetPageHandler(w, r, logger, roomService)
	})

	/*mux.HandleFunc("POST /data", func(w http.ResponseWriter, r *http.Request) {
		data.PostHandler(w, r, logger, ds)
	})
	mux.HandleFunc("PUT /data", func(w http.ResponseWriter, r *http.Request) {
		data.PutHandler(w, r, logger, ds)
	})
	mux.HandleFunc("GET /data", func(w http.ResponseWriter, r *http.Request) {
		data.GetHandler(w, r, logger, ds)
	})
	mux.HandleFunc("GET /data/{id}", func(w http.ResponseWriter, r *http.Request) {
		data.GetByIDHandler(w, r, logger, ds)
	})
	mux.HandleFunc("DELETE /data/{id}", func(w http.ResponseWriter, r *http.Request) {
		data.DeleteHandler(w, r, logger, ds)
	})*/
	return err
}
