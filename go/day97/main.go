package main 
import (
	"fmt"
	"net/http"
)
http.handleFunc("/scan", handler)
http.ListenAndServe(":8080", nil )
