// +build integration
package main

import (
	"testing"
)

var a = authService{Base: "http://localhost:8001"}

func TestWrongUsernamePass(t *testing.T) {
	if a.Login("user1", "wrongpass").Token != "" {
		t.Fail()
	}
}

func TestCorrectCase(t *testing.T) {
	if a.Login("user1", "pass1").Token == "" {
		t.Fail()
	}
}

func TestInvalidReqAuth(t *testing.T) {
	username := "user1"
	lr := a.Login(username, "wrongpass")
	if a.Authenticate(username, lr.Token) {
		t.Fail()
	}
}

// Request should be successed if the user has a valid session token
func TestUserReqAuth(t *testing.T) {
	username := "user1"
	lr := a.Login(username, "pass1")
	if !a.Authenticate(username, lr.Token) {
		t.Fail()
	}
}

// Request should be rejected in the case user has logged out
func TestUserReqAuthAfterLogout(t *testing.T) {
	username := "user1"
	lr := a.Login(username, "pass1")
	if !a.Logout(username, lr.Token) {
		t.Fail()
	}

	if a.Authenticate(username, lr.Token) {
		t.Fail()
	}
}
