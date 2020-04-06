package userinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/database/models"
	"project/database/query"
	"strings"

	"github.com/google/uuid"

	//"project/database/query"
	"project/httpsrv"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		{
			uList := query.GET()
			b, err := json.Marshal(uList)
			if err != nil {
				fmt.Fprintf(w, "Error: %s", err)
				return
			}
			fmt.Fprintf(w, string(b))
			break
		}
	case "POST":
		{
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}
			var userget user.User
			json.Unmarshal(b, &userget)
			if query.POST(userget) {
				fmt.Fprintf(w, "success")
			} else {
				fmt.Fprintf(w, "some has wrong")
			}

			break
		}
	case "PUT":
		{
			args := strings.Split(r.URL.String(), ":")
			if len(args) > 1 {
				uuidUser := args[1]
				if uuidUser == "" {
					fmt.Fprintf(w, "Error: %s", "Arguments not set")
				}
			} else {
				fmt.Fprintf(w, "Error: %s", "Arguments not set")
			}

			break
		}
	case "DELETE":
		{
			fmt.Println("DELETE call")

			args := strings.Split(r.URL.String(), ":")
			if len(args) > 1 {
				uuidUser := args[1]
				if uuidUser == "" {
					fmt.Fprintf(w, "Error: %s", "Arguments not set")
					return
				} else {
					uuidGet, err := uuid.Parse(uuidUser)
					if err != nil {
						fmt.Fprintf(w, "Error: %s", "UUID is wrong")
						return
					} else {
						if query.DELETE(uuidGet) {
							fmt.Fprintf(w, "Success: %s", "DELETED")
						} else {
							fmt.Fprintf(w, "Error: %s", "DELETED")
						}
					}
				}
			} else {
				fmt.Fprintf(w, "Error: %s", "Arguments not set")
				return
			}

			break
		}
	}
}

func GetEP() httpsrv.EPHandler {

	h1 := httpsrv.EPHandler{
		URL:        "/userinfo/",
		HandleFunc: handleRequest,
	}
	return h1
}
