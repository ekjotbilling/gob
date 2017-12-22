package main

import (
	"./gob"
	"errors"
	"io"
	"net/http"
	// "fmt"
)

type BaseContext struct {
	ReqCount int
}

func (c *BaseContext) SayHello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello")
}

func (c *BaseContext) SayWorld() (string, int, error) {
	return "World!", 200, nil
}

func (c *BaseContext) SetRandomVal(w http.ResponseWriter, req *http.Request) error {
	c.ReqCount = 21
	return nil
}

type UserProfile struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type UserCreate struct {
	*UserProfile
}

type UserContext struct {
	*BaseContext
	AuthToken string
	writer    http.ResponseWriter
}

func (c *UserContext) SetupContext(w http.ResponseWriter, req *http.Request) error {
	c.writer = w
	return nil
}

func (c *UserContext) RequiresLoggedIn(w http.ResponseWriter, req *http.Request) error {
	authToken := req.Header.Get("x-auth-token")
	if authToken == "FAKE" {
		return errors.New("fake auth token received!!")
	}
	c.AuthToken = authToken
	return nil
}

func (c *UserContext) GetCurrentUser() (*UserProfile, int, error) {
	user := &UserProfile{
		Name:  "John Doe",
		Age:   21,
		Email: "john@doe.co",
	}
	return user, 200, nil
}

func (c *UserContext) CreateUser(req *UserCreate) (*UserProfile, int, error) {

	if req.Age < 18 {
		return nil, 400, errors.New("User age must be atleast 18!")
	}

	user := &UserProfile{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}

	userAlreadyExists := req.Name == "Admin Man"
	if userAlreadyExists {
		c.writer.Header().Set("Location", "https://api.com/user/as98da")

		return user, 201, nil
	}

	return user, 200, nil
}

func (c *BaseContext) Say() (string, int, error) {
  return "Say", 200, nil
}

// -> both are same paths
// /p/:id
// /p/:man
//
// /p/:id/comments
// /p/:user/friends
// -> converts to
//   user
// /p/:/comments
//   user
// /p/:/friends
func main() {
	rootRouter := gob.NewRouter(BaseContext{}, "/api/v2").
		Middleware((*BaseContext).SetRandomVal).
		Route("GET", "/hello", (*BaseContext).SayHello).
		Route("GET", "/world", (*BaseContext).SayWorld).
    Route("GET", "/speak/:language", (*BaseContext).Say)

	userRouter := rootRouter.Subrouter(UserContext{}, "/user").
		Middleware((*UserContext).SetupContext).
		Route("POST", "/", (*UserContext).CreateUser)

	userRouter.Middleware((*UserContext).RequiresLoggedIn).
		Route("GET", "/me", (*UserContext).GetCurrentUser)

	http.ListenAndServe(":3000", rootRouter)
}
