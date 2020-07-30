// setup.go

package main

import (
	"github.com/vugu/vugu"
)

var AuthenticationData AuthenticationDataType

func setupAuthentication() {
	AuthenticationData.ClientID = "6fst7hjms26vsahdp3c1pu95bc"
	AuthenticationData.ClientName = "todo-api-client"
	AuthenticationData.RestEndpoint = "https://2wwyvmz2zd.execute-api.eu-west-2.amazonaws.com/prod"
	AuthenticationData.RedirectURI = "http://localhost:8844"

}

// OVERALL APPLICATION WIRING IN vuguSetup
func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	// CREATE THE ROOT COMPONENT
	root := &Root{events: eventEnv}
	root.Body = &ToDoList{} // A COMPONENT WITH PAGE CONTENTS

	return root
}
