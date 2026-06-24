This exercise builds on what you learned about headers.

Instead of reading the standard User-Agent header, you'll read a custom header:

X-API-Key

Many APIs use headers to authenticate requests.

Solution
package main

import (
	"fmt"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dashboard" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	apiKey := r.Header.Get("X-API-Key")

	if apiKey != "secret123" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Fprint(w, "Welcome to the dashboard!")
}

func main() {
	http.HandleFunc("/", DashboardHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
How it works
Step 1: Read the Header
apiKey := r.Header.Get("X-API-Key")

Suppose the client sends:

GET /dashboard HTTP/1.1
X-API-Key: secret123

Then:

apiKey == "secret123"
Step 2: Validate the Key
if apiKey != "secret123" {

If the key is:

wrongkey

or

""

(missing header)

then:

http.Error(w, "Unauthorized", http.StatusUnauthorized)

returns:

401 Unauthorized
Step 3: Allow Access

If the key matches:

secret123

the request continues:

fmt.Fprint(w, "Welcome to the dashboard!")

Response:

Welcome to the dashboard!
Testing
No API Key

Browser:

http://localhost:8080/dashboard

Response:

401 Unauthorized
Wrong API Key
curl -H "X-API-Key: wrongkey" http://localhost:8080/dashboard

Response:

401 Unauthorized
Correct API Key
curl -H "X-API-Key: secret123" http://localhost:8080/dashboard

Response:

Welcome to the dashboard!
What this teaches

Headers are often used for authentication.

A request can contain:

GET /dashboard HTTP/1.1
X-API-Key: secret123

and in Go you access it with:

r.Header.Get("X-API-Key")

This exercise introduces a basic authentication pattern:

Client sends a credential (API key).
Server checks it.
If valid → allow access.
If invalid → return 401 Unauthorized.