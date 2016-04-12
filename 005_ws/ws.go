package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/google/go-github/github"

	"golang.org/x/net/websocket"
	"golang.org/x/oauth2"
	ghoauth "golang.org/x/oauth2/github"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"user:email"},
		Endpoint:     ghoauth.Endpoint,
	}
)

type connInfo struct {
	userID string
	io.Writer
}

type online struct {
	sync.Mutex
	connections map[string]connInfo
}

var on online

func (o *online) add(ws io.Writer) string {
	o.Lock()
	id := connID()
	o.connections[id] = connInfo{"", ws}
	o.Unlock()
	return id
}

func (o *online) remove(id string) {
	o.Lock()
	delete(o.connections, id)
	o.Unlock()
}

func (o *online) broadcast(msg string) {
	for _, ws := range o.connections {
		ws.Write([]byte(msg))
	}

}

func connID() string {
	b := make([]byte, 64)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	return fmt.Sprintf("%x", md5.Sum(b))
}

type oauthURL struct {
	URL string
}

type ghUser struct {
	User string
}

func chat(ws *websocket.Conn) {
	connID := on.add(ws)
	defer func() {
		on.remove(connID)
	}()
	url, _ := json.Marshal(oauthURL{URL: oauthConf.AuthCodeURL(connID, oauth2.AccessTypeOnline)})
	ws.Write(url)
	request := make([]byte, 1000)
	for {
		n, err := ws.Read(request)
		if err != nil {
			fmt.Println("Error", err.Error())
			return
		}
		on.broadcast(string(request[:n]))
	}

}

func ghCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if _, ok := on.connections[state]; !ok {
		return
	}
	code := r.FormValue("code")
	token, err := oauthConf.Exchange(oauth2.NoContext, code)
	if err != nil {
		return
	}

	oauthClient := oauthConf.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get("")
	if err != nil {
		return
	}
	userName, _ := json.Marshal(ghUser{User: *user.Login})

	c := on.connections[state]

	c.Write(userName)
	c.userID = *user.Login
	on.Lock()
	on.connections[state] = c
	on.Unlock()
}

func init() {
	on.connections = make(map[string]connInfo)
}
func main() {
	http.HandleFunc("/githubcb", ghCallback)
	http.Handle("/chat", websocket.Handler(chat))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
