package orchestrator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Vlad-PyDev/AsyncCalculationWebService/internal/models"
	"github.com/Vlad-PyDev/AsyncCalculationWebService/pkg/ast"
	"github.com/Vlad-PyDev/AsyncCalculationWebService/pkg/crypto/jwt"
	"github.com/Vlad-PyDev/AsyncCalculationWebService/pkg/crypto/password"
)

func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	expressionID := "-1"
	expressionID = strings.TrimPrefix(r.URL.Path, "/api/v1/expressions/")
	if checkId(expressionID) {
		idValue, err := strconv.Atoi(expressionID)
		if err != nil {
			errMsg := fmt.Sprintf("%s", err)
			errorResponse(w, "internal server error", http.StatusInternalServerError)
			log.Printf("Code: %v, %s", http.StatusInternalServerError, errMsg)
			return
		}

		userIDValue := r.Context().Value(userID)
		expressionData, err := db.SelectExprByID(r.Context(), idValue, userIDValue.(int))
		if err != nil {
			errorResponse(w, "expression does not exist", http.StatusNotFound)
			log.Printf("Code: %v, %s", http.StatusNotFound, err)
			return
		}

		var jsonOutput []byte
		jsonOutput, err = json.MarshalIndent(expressionData, "", "  ")
		if err != nil {
			errorResponse(w, fmt.Sprintf("%s", err), http.StatusInternalServerError)
			log.Printf("Code: %v, error marshaling JSON", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonOutput)
		if err != nil {
			errorResponse(w, "error with JSON data", http.StatusInternalServerError)
			log.Printf("Code: %v, internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	expressions, err := db.SelectExpressions(r.Context(), r.Context().Value(userID).(int))
	if err != nil {
		errorResponse(w, "you haven't calculated any expressions yet", http.StatusInternalServerError)
		log.Printf("Code: %v, no expressions for user %v", http.StatusInternalServerError, userID)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write([]byte(expressions))
	if err != nil {
		errorResponse(w, "error with JSON data", http.StatusInternalServerError)
		log.Printf("Code: %v, internal server error", http.StatusInternalServerError)
		return
	}
}

func ExpressionHandler(w http.ResponseWriter, r *http.Request) {
	expressionData := r.Context().Value(ctxKey).(*Expression)

	astNode, err := ast.Build(expressionData.exp)
	if err != nil {
		errorMsg := fmt.Sprintf("%s", err)
		log.Printf("Expression %d: failed to build AST - %s", expressionData.id, errorMsg)
		if err := db.UpdateExpression(context.Background(), expressionData.id, errorMsg, 0.0); err != nil {
			log.Printf("Failed to update expression %d: %v", expressionData.id, err)
		}
		return
	}
	expressionProcessor := NewExpression(astNode)

	resultValue, err := expressionProcessor.calc()
	if err != nil {
		log.Printf("Expression %v: division by zero detected", expressionData.id)
		if err := db.UpdateExpression(context.Background(), expressionData.id, "zero division error", 0.0); err != nil {
			log.Printf("Failed to update expression %d: %v", expressionData.id, err)
		}
		return
	}
	log.Printf("Expression %v computed successfully", expressionData.id)

	if err := db.UpdateExpression(context.Background(), expressionData.id, "done", resultValue); err != nil {
		log.Printf("Failed to update expression %d: %v", expressionData.id, err)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "invalid request method", http.StatusMethodNotAllowed)
		log.Printf("Code: %v, invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		errorResponse(w, "invalid request body", http.StatusBadRequest)
		log.Printf("Code: %v, JSON decoding error", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userData, err := db.SelectUserByLogin(ctx, requestBody.Login)
	if err != nil {
		errorResponse(w, "user not found", http.StatusNotFound)
		log.Printf("Code: %v, user %v not found", http.StatusNotFound, requestBody.Login)
		return
	}
	if err := password.Compare(userData.Password, requestBody.Password); err != nil {
		errorResponse(w, "incorrect password", http.StatusForbidden)
		log.Printf("Code: %v, incorrect password", http.StatusForbidden)
		return
	}

	var response struct {
		Jwt string `json:"jwt"`
	}
	tokenValue, err := jwt.Generate(int(userData.ID))
	if err != nil {
		errorResponse(w, "internal server error", http.StatusInternalServerError)
		log.Printf("Code: %v, token generation error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    tokenValue,
		Expires:  time.Now().Add(10 * time.Minute),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	response.Jwt = tokenValue
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, "invalid request method", http.StatusMethodNotAllowed)
		log.Printf("Code: %v, invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		errorResponse(w, "invalid request body", http.StatusBadRequest)
		log.Printf("Code: %v, JSON decoding error", http.StatusBadRequest)
		return
	}

	if len(requestBody.Password) == 0 {
		errorResponse(w, "password cannot be empty", http.StatusForbidden)
		log.Printf("Code: %v, empty password", http.StatusForbidden)
		return
	}

	hashedPassword, err := password.Generate(requestBody.Password)
	if err != nil {
		errorResponse(w, "internal server error", http.StatusInternalServerError)
		log.Printf("Code: %v, %s", http.StatusInternalServerError, err)
		return
	}

	ctx := r.Context()
	newUser := &models.User{
		Login:    requestBody.Login,
		Password: hashedPassword,
	}
	_, err = db.InsertUser(ctx, newUser)
	if err != nil {
		errorResponse(w, "user already exists", http.StatusConflict)
		log.Printf("Code: %v, user %s already exists", http.StatusConflict, requestBody.Login)
		return
	}

	log.Printf("user: %v registered successfully", newUser.Login)
	w.WriteHeader(http.StatusOK)
}

func databaseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			errorResponse(w, "invalid request method", http.StatusMethodNotAllowed)
			log.Printf("Code: %v, invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var requestBody ExpressionReq
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			errorResponse(w, "internal server error", http.StatusInternalServerError)
			log.Printf("Code: %v, error decoding request body", http.StatusInternalServerError)
			return
		}

		log.Printf("Inserting expression into database")
		expression := &models.Expression{
			Expression: requestBody.Expression,
			Status:     "in process",
			Result:     0.0,
		}
		expressionID, err := db.InsertExpression(r.Context(), expression, r.Context().Value(userID).(int))
		if err != nil {
			errorResponse(w, "internal server error", http.StatusInternalServerError)
			log.Printf("Code: %v, database error", http.StatusInternalServerError)
			return
		}
		responseID := RespID{Id: int(expressionID)}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseID)

		go func() {
			expressionCtx := &Expression{
				exp: requestBody.Expression,
				id:  int(expressionID),
			}
			updatedCtx := context.WithValue(r.Context(), ctxKey, expressionCtx)
			next.ServeHTTP(w, r.WithContext(updatedCtx))
		}()
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authToken string
		cookieData, err := r.Cookie("jwt")
		if checkCookie(cookieData, err) {
			authToken = cookieData.Value
			log.Print("token retrieved from cookie")
		} else {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				errorResponse(w, "authorization required", http.StatusUnauthorized)
				log.Printf("Code: %v, unauthorized access", http.StatusUnauthorized)
				return
			}

			headerParts := strings.Split(authHeader, " ")
			if len(headerParts) != 2 || headerParts[0] != "Bearer" {
				errorResponse(w, "invalid token format", http.StatusUnauthorized)
				log.Printf("Code: %v, invalid token format", http.StatusUnauthorized)
				return
			}
			authToken = headerParts[1]
			log.Print("token retrieved from header")
		}

		isValid, userIDValue := jwt.Verify(authToken)
		if !isValid {
			errorResponse(w, "invalid token", http.StatusUnauthorized)
			log.Printf("Code: %v, invalid token", http.StatusUnauthorized)
			return
		}

		updatedCtx := context.WithValue(r.Context(), userID, userIDValue)
		next.ServeHTTP(w, r.WithContext(updatedCtx))
	})
}

func logsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request Method: %s, URL: %s", r.Method, r.URL)
		requestStart := time.Now()
		next.ServeHTTP(w, r)
		requestDuration := time.Since(requestStart)
		log.Printf("Request Method: %s, completed in: %v", r.Method, requestDuration)
	})
}
