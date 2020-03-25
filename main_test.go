package main

import  (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

const (
	BadRequestCode = 400
	SuccessRequestCode = 200
)

type TestStruct struct {
	requestBody string
	expectedStatusCode int
	responseBody string
	observedStatusCode int
}

func TestMain(t *testing.T) {
	url := "http://localhost:10000"
}