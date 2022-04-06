package tests

import (
	"github.com/revel/revel/testing"
)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// Check main page is kinda there
func (t *AppTest) TestHelloLandingPage() {
	t.Get("/App/Hello")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

//func (t *AppTest) TestForUsersMyName() {
//	data := url.Values{}
//	data.Add("myName", "testuser")
//
//	t.PostForm("/App/Hello", data)
//t.AssertOk()
//t.AssertContains("login successful")
//t.AssertOk()
//}

func (t *AppTest) TestValidateUsersCreated() {
	t.Get("/App/Hello?myName=Jason")
	t.Get("/App/Hello?myName=User1")
	t.Get("/App/Hello?myName=User2")
	t.Get("/App/Hello?myName=User3")
	t.AssertOk()
}

func (t *AppTest) After() {
	println("Tear down")
}
