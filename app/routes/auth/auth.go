package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/SantiagoZuluaga/drawflowapi/app/database"
	"github.com/SantiagoZuluaga/drawflowapi/app/models"
	"github.com/SantiagoZuluaga/drawflowapi/app/utils"
	"github.com/dgraph-io/dgo/v200/protos/api"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type QueryResponse struct {
	User []models.User `json:"user"`
}

func GetUserByID(id string) (models.User, error) {
	db, err := database.GetDatabase()
	if err != nil {
		return models.User{}, err
	}

	const query = `
		query Me($id: string){
			user(func: uid($id)) {
				uid
				fullname
				email
				password
				blocked
				createdAt
				updatedAt
			}
		}
	`
	response, err := db.NewTxn().QueryWithVars(
		context.Background(),
		query,
		map[string]string{"$id": id},
	)
	if err != nil {
		return models.User{}, err
	}

	var queryResponse QueryResponse
	if err := json.Unmarshal(response.Json, &queryResponse); err != nil {
		return models.User{}, err
	}

	if len(queryResponse.User) == 0 {
		return models.User{}, nil
	}

	return queryResponse.User[0], nil
}

func GetUserByEmail(email string) (models.User, error) {
	db, err := database.GetDatabase()
	if err != nil {
		return models.User{}, err
	}

	const query = `
		query Me($email: string){
			user(func: eq(email, $email)) {
				uid
				fullname
				email
				password
				blocked
				createdAt
				updatedAt
			}
		}
	`
	response, err := db.NewTxn().QueryWithVars(
		context.Background(),
		query,
		map[string]string{"$email": email},
	)
	if err != nil {
		return models.User{}, err
	}

	var queryResponse QueryResponse
	if err := json.Unmarshal(response.Json, &queryResponse); err != nil {
		return models.User{}, err
	}

	if len(queryResponse.User) == 0 {
		return models.User{}, nil
	}

	return queryResponse.User[0], nil
}

func Validate(res http.ResponseWriter, req *http.Request) {

	authorization := req.Header.Get("Authorization")
	if authorization == "" {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("Invalid jwt token"))
		return
	}

	token := strings.Replace(authorization, "Bearer ", "", 1)

	id, err := utils.ValidateToken(token)
	if err != nil {
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte("Invalid jwt token"))
		return
	}

	user, err := GetUserByID(id)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	if user == (models.User{}) {
		fmt.Println(err)
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("User not found."))
		return
	}

	userParsed, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	res.Write([]byte(userParsed))
}

func Login(res http.ResponseWriter, req *http.Request) {

	var input models.LoginInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	if input.Email == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Please provider a identifier"))
		return
	}

	if input.Password == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Please provider a password"))
		return
	}

	user, err := GetUserByEmail(input.Email)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	if user == (models.User{}) {
		fmt.Println(err)
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte("User not found."))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Password invalid."))
		return
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	userParsed, err := json.Marshal(models.UserPayload{
		JWT:  token,
		User: user,
	})
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	res.Write(userParsed)
}

func Register(res http.ResponseWriter, req *http.Request) {

	var input models.RegisterInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	if input.Fullname == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Please provider a fullname"))
		return
	}

	if input.Email == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Please provider a email"))
		return
	}

	if input.Password == "" {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Please provider a password"))
		return
	}

	if input.ConfirmPassword == "" || input.Password != input.ConfirmPassword {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Password must match"))
		return
	}

	db, err := database.GetDatabase()
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	user, err := GetUserByEmail(input.Email)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	if user != (models.User{}) {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Email is already taken"))
		return
	}

	now := time.Now()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	newUser := models.User{
		Fullname:  input.Fullname,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Blocked:   false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	userParsed, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}

	mutation := &api.Mutation{
		CommitNow: true,
		SetJson:   userParsed,
	}

	resp, err := db.NewTxn().Mutate(context.Background(), mutation)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte("Error in server, try again."))
		return
	}
	res.Write([]byte(resp.Json))
}

func Routes(route chi.Router) {
	route.Get("/validate", Validate)
	route.Post("/login", Login)
	route.Post("/register", Register)
}
