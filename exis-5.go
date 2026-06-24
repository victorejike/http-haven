This exercise introduces HTTP headers.

Until now you've worked with:

r.URL.Path
r.URL.Query()
r.Method

Now you'll access:

r.Header

which contains extra information sent by the client.

Solution
package main

import (
	"fmt"
	"net/http"
)

func AgentHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/agent" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	userAgent := r.Header.Get("User-Agent")

	if userAgent == "" {
		userAgent = "Unknown Client"
	}

	fmt.Fprintf(w, "You are visiting us using: %s", userAgent)
}

func main() {
	http.HandleFunc("/", AgentHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
What is a Header?

When a browser visits a website, it sends more than just the URL.

For example:

GET /agent HTTP/1.1
Host: localhost:8080
User-Agent: Mozilla/5.0
Accept: text/html

The lines below the request line are called headers.

Think of them as metadata about the request.

Accessing Headers

Go stores them in:

r.Header

You can get a specific header with:

r.Header.Get("User-Agent")

For a browser request, this might return:

Mozilla/5.0 (Windows NT 10.0; Win64; x64)

or

Mozilla/5.0 (X11; Linux x86_64)

depending on the browser and operating system.

Why Check for Empty?

Some clients may not send a User-Agent header.

Then:

userAgent := r.Header.Get("User-Agent")

returns:

""

(an empty string)

So we handle it:

if userAgent == "" {
	userAgent = "Unknown Client"
}
Example Responses
Browser Request

Visit:

http://localhost:8080/agent

Response might be:

You are visiting us using: Mozilla/5.0 (X11; Linux x86_64)
Curl Request
curl http://localhost:8080/agent

Response might be:

You are visiting us using: curl/8.5.0
No User-Agent Header
curl -H "User-Agent:" http://localhost:8080/agent

Response:

You are visiting us using: Unknown Client
Important Concept

HTTP requests have several parts:

r.Method         // GET, POST...
r.URL.Path       // /agent
r.URL.Query()    // ?name=Alice
r.Header         // User-Agent, Content-Type, Accept...
r.Body           // POST data