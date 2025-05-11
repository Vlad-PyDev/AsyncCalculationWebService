package orchestrator

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"sync"

	"github.com/Vlad-PyDev/AsyncCalculationWebService/pkg/database"
)

const port = ":8080"

type (
	Orchestrator struct {
	}

	ExpressionReq struct {
		Expression string `json:"expression"`
	}

	RespID struct {
		Id int `json:"id"`
	}

	ErrorResponse struct {
		Res string `json:"error" example:"Internal server error"`
	}

	Expression struct {
		exp string
		id  int
	}

	contextKey string
	userid     string
)

var (
	db     *database.SqlDB
	mu     sync.Mutex
	ctxKey contextKey = "expression id"
	userID userid     = "user id"
)

func (o *Orchestrator) Run() {
	db = database.NewDB()
	defer db.Store.Close()

	StartManager()
	go runGRPC()

	router := http.NewServeMux()

	registerHandler := http.HandlerFunc(RegisterHandler)
	loginHandler := http.HandlerFunc(LoginHandler)
	expressionHandler := http.HandlerFunc(ExpressionHandler)
	dataHandler := http.HandlerFunc(GetDataHandler)

	router.Handle("/api/v1/register", logsMiddleware(registerHandler))
	router.Handle("/api/v1/login", logsMiddleware(loginHandler))
	router.Handle("/api/v1/calculate", logsMiddleware(authMiddleware(databaseMiddleware(expressionHandler))))
	router.Handle("/api/v1/expressions/", logsMiddleware(authMiddleware(dataHandler)))

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func checkId(id string) bool {
	if id == "-1" || id == "" {
		return false
	}

	pattern := "^[0-9]+$"
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(id)
}

func errorResponse(w http.ResponseWriter, err string, statusCode int) {
	w.WriteHeader(statusCode)
	response := ErrorResponse{Res: err}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func checkCookie(cookie *http.Cookie, err error) bool {
	if err != nil {
		return false
	}

	tokenValue := cookie.Value
	return len(tokenValue) != 0
}

func New() *Orchestrator {
	return &Orchestrator{}
}
