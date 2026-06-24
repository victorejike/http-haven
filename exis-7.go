This exercise teaches HTTP redirects.

Sometimes an endpoint moves to a new location:

/legacy  →  /v2

Instead of showing content on /legacy, you tell the browser:

"This page has moved. Go to /v2 instead."

Solution
package main

import (
	"fmt"
	"net/http"
)

func LegacyHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func V2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}

func main() {
	http.HandleFunc("/legacy", LegacyHandler)
	http.HandleFunc("/v2", V2Handler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
Understanding http.Redirect

The important line is:

http.Redirect(w, r, "/v2", http.StatusMovedPermanently)

It takes four arguments:

http.Redirect(
	w,
	r,
	"/v2",
	http.StatusMovedPermanently,
)
Argument	Meaning
w	Response writer
r	Current request
"/v2"	Where to send the client
http.StatusMovedPermanently	Status code 301
What Happens Behind the Scenes

User visits:

http://localhost:8080/legacy

Server responds:

HTTP/1.1 301 Moved Permanently
Location: /v2

The browser sees the Location header and automatically requests:

http://localhost:8080/v2

Then your V2Handler responds:

Welcome to version 2
Why 301?

The exercise specifically asks for:

http.StatusMovedPermanently

which equals:

301

This tells browsers and search engines:

"This route has permanently moved."

Other redirects exist:

Status Code	Meaning
301	Moved Permanently
302	Found (temporary redirect)
307	Temporary Redirect
308	Permanent Redirect

For this exercise, use 301.

Testing

Visit:

http://localhost:8080/legacy

The browser will automatically end up at:

http://localhost:8080/v2

and show:

Welcome to version 2

Using curl:

curl -i http://localhost:8080/legacy

You should see something like:

HTTP/1.1 301 Moved Permanently
Location: /v2

To follow the redirect automatically:

curl -L http://localhost:8080/legacy

Response:

Welcome to version 2
What this exercise is teaching

So far you've learned how to work with:

r.URL.Path
r.URL.Query()
r.Method
r.Header
r.Body

This exercise introduces another important HTTP concept:

http.Redirect(...)

which lets your server tell clients:

"The resource you're looking for is somewhere else."

yes

If you're ready to make Exercise 6 a little more realistic, try implementing it with path validation and method checking, like you've been doing in previous exercises:

package main

import (
	"fmt"
	"net/http"
)

func LegacyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func V2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "Welcome to version 2")
}

func main() {
	http.HandleFunc("/legacy", LegacyHandler)
	http.HandleFunc("/v2", V2Handler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
What http.Redirect actually sends

When this line runs:

http.Redirect(w, r, "/v2", http.StatusMovedPermanently)

Go automatically sends:

HTTP/1.1 301 Moved Permanently
Location: /v2

The important part is:

Location: /v2

The browser sees that header and immediately makes another request:

GET /v2 HTTP/1.1

That's why users end up on /v2 without manually typing it.

Interview-style question

If a user visits:

http://localhost:8080/legacy

What handlers run?

Answer:

LegacyHandler runs first.
It sends a 301 redirect to /v2.
The browser makes a new request to /v2.
V2Handler runs.
The user sees:
Welcome to version 2

This is an important concept because many real websites use redirects when pages are renamed, APIs are versioned (/v1 → /v2), or domains change.

Exercise 6: Simple Redirector (Status Codes)
Goal: Create a /legacy route. Whenever a user hits this endpoint, permanently redirect them to a new route /v2 with a friendly "Welcome to version 2" message.
Key Tasks:
Redirect traffic using the http.Redirect helper function.
Use the proper status code for a permanent move (http.StatusMovedPermanently).Exercise 6: Simple Redirector (Status Codes)
Goal: Create a /legacy route. Whenever a user hits this endpoint, permanently redirect them to a new route /v2 with a friendly "Welcome to version 2" message.
Key Tasks:
Redirect traffic using the http.Redirect helper function.
Use the proper status code for a permanent move (http.StatusMovedPermanently).

For this exercise, the simplest correct solution is:

package main

import (
	"fmt"
	"net/http"
)

func LegacyHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/v2", http.StatusMovedPermanently)
}

func V2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to version 2")
}

func main() {
	http.HandleFunc("/legacy", LegacyHandler)
	http.HandleFunc("/v2", V2Handler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
How it works

When a user visits:

http://localhost:8080/legacy

the handler runs:

http.Redirect(w, r, "/v2", http.StatusMovedPermanently)

This sends a response like:

HTTP/1.1 301 Moved Permanently
Location: /v2

The browser reads the Location header and automatically goes to:

http://localhost:8080/v2

Then V2Handler responds with:

Welcome to version 2
The important new concept

http.Redirect does not send the user directly to another handler.

Instead:

Client requests /legacy
Server responds with 301 Moved Permanently
Browser makes a new request to /v2
V2Handler runs
User sees "Welcome to version 2"
Testing

Visit:

http://localhost:8080/legacy

You should end up at:

http://localhost:8080/v2

and see:

Welcome to version 2