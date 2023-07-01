package main

/*
type APIServer struct {
	listenAddr string
	db         PriceStorage
}

func NewAPIServer(listenAddr string, pd PriceStorage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		db:         pd,
	}
}

// adapter to allow ordinary functions to be used as HTTP Handlers
// will only really be used in the makeHTTPHandler function below
type apifunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandler(f apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			fmt.Println(err)
		}
	}
}

func (s *APIServer) Run() {

	r := mux.NewRouter()

	r.Handle("/api/warehouse/pricing", makeHTTPHandler(s.handleAddItem))
	r.Handle("/api/warehouse/pricing/{id}", makeHTTPHandler(s.handleGetPrice))

	log.Println("Starting API server on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, r)
}
*/
