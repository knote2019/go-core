package main

import (
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Address string
}

type UserResource struct {
	users map[string]User
}

func (us UserResource) createUser(req *restful.Request, res *restful.Response) {
	user := new(User)
	err := req.ReadEntity(&user)
	if err == nil {
		us.users[user.Name] = *user
		res.WriteHeaderAndEntity(http.StatusCreated, user)
	} else {
		res.AddHeader("Content-Type", "text/plain")
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (us UserResource) updateUser(req *restful.Request, res *restful.Response) {
	userName := req.PathParameter("user-name")
	user := User{Name: userName}
	err := req.ReadEntity(&user)
	if err == nil {
		us.users[user.Name] = user
		res.WriteHeaderAndEntity(http.StatusOK, user)
	} else {
		res.AddHeader("Content-Type", "text/plain")
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func (us UserResource) deleteUser(req *restful.Request, res *restful.Response) {
	userName := req.PathParameter("user-name")
	delete(us.users, userName)

}

func (us UserResource) findUser(req *restful.Request, res *restful.Response) {
	userName := req.PathParameter("user-name")
	user, ok := us.users[userName]
	if ok {
		res.WriteEntity(user)
	} else {
		res.AddHeader("Content-Type", "text/plain")
		res.WriteErrorString(http.StatusNotFound, "user not find")
	}
}

func (us UserResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/users").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.POST("").To(us.createUser))
	ws.Route(ws.PUT("/{user-name}").To(us.updateUser))
	ws.Route(ws.DELETE("/{user-name}").To(us.deleteUser))
	ws.Route(ws.GET("/{user-name}").To(us.findUser))
	container.Add(ws)
}

func main() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	us := UserResource{map[string]User{}}
	us.Register(wsContainer)
	log.Printf("start to listen on: 0.0.0.0:5555")
	server := &http.Server{Addr: "0.0.0.0:5555", Handler: wsContainer}
	result := server.ListenAndServe()
	log.Fatal(result)

}
