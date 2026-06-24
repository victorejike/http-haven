Go HTTP Server Cheat Sheet
Core Standard Library Functions
1. Routing & Server Management (net/http)
http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))
Registers a handler function for a specific URL path pattern.
http.ListenAndServe(addr string, handler Handler) error
Starts an HTTP server on the specified address (e.g., ":8080"). Passing nil uses the default system router (DefaultServeMux).
http.Error(w ResponseWriter, error string, code int)
A helper that sends a specific error message string and numeric HTTP status code back to the client, automatically ending the request lifecycle safely.
http.Redirect(w ResponseWriter, r *Request, url string, code int)
Sends an HTTP redirect status code (like 301 or 302) forcing the client's browser to jump to a new target URL path.
2. Reading Inputs & Formatting (io, strconv, fmt)
io.ReadAll(r io.Reader) ([]byte, error)
Reads all remaining data from an input stream (like r.Body) until it hits the end of the file/stream (EOF). Returns a byte slice.
strconv.Atoi(s string) (int, error)
Stands for "ASCII to Integer". Converts a text string into a native numeric int. Returns an error if the text contains non-numeric characters.
fmt.Fprintf(w io.Writer, format string, a ...any)
Formats data according to a template string and writes the output directly into an open network socket stream or file (w).



Request Context Breakdown (*http.Request)
When a client hits your server, all incoming information is bundled inside the pointer to the http.Request struct (usually named r):
Struct Field / Method
Purpose / Explanation
Example Usage
r.Method
A string representing the incoming HTTP type. Always use standard library constants for comparisons.
if r.Method == http.MethodPost
r.Body
An open stream containing data uploaded via POST/PUT requests. Always close it after reading to prevent memory leaks.
defer r.Body.Close()
r.URL.Query()
Parses the raw URL query string (everything after the ?) and returns a map-like structure (Values).
queryMap := r.URL.Query()
r.URL.Query().Get(key)
Fetches the value of a specific query parameter. Returns an empty string "" if the key does not exist.
name := r.URL.Query().Get("name")
r.Header.Get(key)
Fetches the value of a specific HTTP request header (case-insensitive). Returns "" if missing.
token := r.Header.Get("X-API-Key")





Response Management (http.ResponseWriter)
The http.ResponseWriter interface (usually named w) is your pipeline to send data back to the client. Keep this execution order in mind:
Modify Headers First: Call w.Header().Set("Key", "Value") before anything else if changing content types or metadata.
Write Status Code Second: Call w.WriteHeader(http.StatusCreated) if returning something other than a standard 200 OK.
Write Body Content Last: Call w.Write([]byte) or fmt.Fprintf(w) to push actual visual content to the user. Writing to the body locks your headers and status code automatically.





















This is a good solution. It satisfies the exercise requirements and also correctly returns a 404 Not Found for other paths.

Let's walk through it:

http.HandleFunc("/", HomeHandle)
This means all requests initially go to HomeHandle.

Then inside the handler:

if r.URL.Path != "/ping" {
	http.Error(w, "404 Not Found", http.StatusNotFound)
	return
}
You manually check whether the path is /ping.

/ping → continue

anything else → return 404

Then:

w.Write([]byte("pong"))
sends "pong" as the response body.

Example Requests
Request:

GET /ping HTTP/1.1
Response:

HTTP/1.1 200 OK

pong
Request:

GET /hello HTTP/1.1
Response:

HTTP/1.1 404 Not Found

404 Not Found
Small improvement
You could register the exact route instead:

http.HandleFunc("/ping", HomeHandle)
and simplify the handler:

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
However, your current approach is actually useful for learning because it teaches you how routers and path validation work. This pattern is very common in projects where you must explicitly check:

r.URL.Path
r.Method
such as your upcoming ascii-art-web project.

One additional thing I'd add is checking the HTTP method:

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/ping" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("pong"))
}
Now the endpoint behaves like a proper HTTP API:

GET /ping → 200 OK

POST /ping → 405 Method Not Allowed

GET /anything → 404 Not Found

That's the pattern you'll use repeatedly in Go web development.


Exercise 2: Query Parameters & Path Validation
Goal: Create a /hello endpoint that reads a name query parameter (e.g., /hello?name=Alice) and responds with "Hello, Alice!". If the parameter is missing, default to "Hello, Guest!".
Tasks:
Extract query parameters using r.URL.Query().Get("name").
Reject any HTTP method that is not GET by returning an http.StatusMethodNotAllowed status code.


and also explain it so that i can understand  it

This exercise introduces query parameters.

A URL has two main parts:

/hello?name=Alice
Path: /hello

Query parameter: name=Alice

The ? starts the query string, and query parameters are key-value pairs:

name=Alice
key = name

value = Alice

Solution
package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", HelloHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
Understanding it step by step
1. Path Validation
if r.URL.Path != "/hello"
When a request arrives:

http://localhost:8080/hello
then:

r.URL.Path
contains:

/hello
If someone visits:

http://localhost:8080/test
then:

r.URL.Path
contains:

/test
and we return:

http.Error(w, "404 Not Found", http.StatusNotFound)
2. Method Validation
if r.Method != http.MethodGet
The browser usually sends:

GET /hello HTTP/1.1
But someone could send:

POST /hello HTTP/1.1
or

PUT /hello HTTP/1.1
The exercise says:

Reject any method that is not GET.

So we return:

http.StatusMethodNotAllowed
which is:

405 Method Not Allowed
3. Reading Query Parameters
This is the new concept.

name := r.URL.Query().Get("name")
Suppose the user visits:

http://localhost:8080/hello?name=Alice
Go extracts:

name = "Alice"
Suppose:

http://localhost:8080/hello?name=Bob
Then:

name = "Bob"
Visualizing it
URL:

/hello?name=Alice
becomes:

r.URL.Query()
which acts roughly like:

map[string][]string{
	"name": {"Alice"},
}
Then:

Get("name")
returns:

"Alice"
4. Default Value
What if the user visits:

http://localhost:8080/hello
There is no query parameter.

So:

name := r.URL.Query().Get("name")
returns:

""
(an empty string)

That's why we check:

if name == "" {
	name = "Guest"
}
Now:

http://localhost:8080/hello
returns:

Hello, Guest!
5. Sending the Response
fmt.Fprintf(w, "Hello, %s!", name)
If:

name = "Alice"
the output becomes:

Hello, Alice!
If:

name = "Guest"
the output becomes:

Hello, Guest!
Test Cases
Request 1
http://localhost:8080/hello
Response:

Hello, Guest!
Request 2
http://localhost:8080/hello?name=Alice
Response:

Hello, Alice!
Request 3
http://localhost:8080/hello?name=Ejike
Response:

Hello, Ejike!
Request 4
POST /hello
Response:

405 Method Not Allowed
The big idea
This exercise is teaching you that an HTTP request contains:

r.URL.Path      // which endpoint was visited
r.Method        // GET, POST, PUT, DELETE...
r.URL.Query()   // values after the '?'
These three things are the foundation of almost every web API you'll build in Go. Once you're comfortable with them, the next step is usually handling JSON request bodies and responses.


Exercise 2: Query Parameters & Path Validation
Goal: Create a /hello endpoint that reads a name query parameter (e.g., /hello?name=Alice) and responds with "Hello, Alice!". If the parameter is missing, default to "Hello, Guest!".
Tasks:
Extract query parameters using r.URL.Query().Get("name").
Reject any HTTP method that is not GET by returning an http.StatusMethodNotAllowed status code.

Here's a straightforward solution that satisfies all the requirements:

package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	http.HandleFunc("/", HelloHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
How to test it
Run:

go run main.go
Open:

http://localhost:8080/hello
Response:

Hello, Guest!
Open:

http://localhost:8080/hello?name=Alice
Response:

Hello, Alice!
Open:

http://localhost:8080/hello?name=Ejike
Response:

Hello, Ejike!
Key line to understand
name := r.URL.Query().Get("name")
For this URL:

/hello?name=Alice
Path = /hello

Query parameter = name=Alice

So:

r.URL.Query().Get("name")
returns:

"Alice"
If the URL is just:

/hello
then there is no name parameter, so it returns:

""
(an empty string), which is why we replace it with "Guest".

This exercise is mainly teaching you how to access:

r.URL.Path → the route (/hello)

r.Method → the HTTP method (GET, POST, etc.)

r.URL.Query() → data after the ? in the URL.


Exercise 3: Text Counter (URL Variables & Methods)
Goal: Build a server with a /count route. If a user sends a GET request, return the text "Send a POST request with text to count words". If they send a POST request, read the text body and return the number of characters.
Key Tasks:
Differentiate between GET and POST methods using r.Method.
Read the entire request body using io.ReadAll(r.Body).
Return the character length as a string.

This exercise introduces a very important concept:

Reading data sent by the client in the request body.

In the previous exercise, data came from the URL:

/hello?name=Alice
Now the data comes from the request body itself.

Solution
package main

import (
	"fmt"
	"io"
	"net/http"
)

func CountHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/count" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}

	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		length := len(body)

		fmt.Fprintf(w, "%d", length)
		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/", CountHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
Understanding the Request Body
When someone sends:

POST /count HTTP/1.1

Hello World
The text:

Hello World
is stored in:

r.Body
Think of r.Body as a stream of data coming from the client.

Reading the Body
body, err := io.ReadAll(r.Body)
This reads everything from the stream and returns:

[]byte("Hello World")
The variable body is a slice of bytes:

[]byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd'}
Counting Characters
length := len(body)
For:

Hello
length becomes:

5
For:

Hello World
length becomes:

11
because the space counts as a character.

Returning the Result
fmt.Fprintf(w, "%d", length)
If:

length = 11
the response becomes:

11
Testing
GET Request
Visit:

http://localhost:8080/count
Response:

Send a POST request with text to count words
POST Request with curl
curl -X POST -d "Hello World" http://localhost:8080/count
Response:

11
Another Example
curl -X POST -d "Go is awesome" http://localhost:8080/count
Characters:

G o _ i s _ a w e s o m e
Total:

13
Response:

13
What this exercise is teaching
By the end of it, you'll know how to use:

r.URL.Path      // Which endpoint was called
r.Method        // GET, POST, PUT, DELETE...
r.Body          // Data sent by the client
io.ReadAll()    // Read the body into memory
len()           // Count bytes/characters
These are the building blocks for handling form submissions, JSON APIs, login requests, and most web applications you'll build in Go.


Exercise 4: Basic Math API (Multiple Query Parameters)
Goal: Create a /calculate route that accepts three query parameters: op (operation), a, and b. For example, /calculate?op=add&a=10&b=5 should respond with Result: 15.
Key Tasks:
Parse string query variables using r.URL.Query().Get().
Convert string inputs to integers using strconv.Atoi().
Support add, subtract, and multiply. Return an HTTP 400 Bad Request if the operation is unknown or parsing fails.

This exercise combines everything you've learned so far:

r.URL.Path

r.Method

Query parameters

Converting strings to integers

Returning proper HTTP errors

Solution
package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/calculate" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	op := r.URL.Query().Get("op")
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")

	a, err := strconv.Atoi(aStr)
	if err != nil {
		http.Error(w, "Invalid value for a", http.StatusBadRequest)
		return
	}

	b, err := strconv.Atoi(bStr)
	if err != nil {
		http.Error(w, "Invalid value for b", http.StatusBadRequest)
		return
	}

	var result int

	switch op {
	case "add":
		result = a + b
	case "subtract":
		result = a - b
	case "multiply":
		result = a * b
	default:
		http.Error(w, "Unknown operation", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Result: %d", result)
}

func main() {
	http.HandleFunc("/", CalculateHandler)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
Understanding it
Step 1: Read the query parameters
Suppose the URL is:

/calculate?op=add&a=10&b=5
Then:

op := r.URL.Query().Get("op")
becomes:

op = "add"
and:

aStr := r.URL.Query().Get("a")
becomes:

aStr = "10"
and:

bStr := r.URL.Query().Get("b")
becomes:

bStr = "5"
Notice these are all strings.

Step 2: Convert strings to integers
You can't do math with strings.

This won't work:

"10" + "5"
So we convert them:

a, err := strconv.Atoi(aStr)
Atoi means:

ASCII TO Integer

Example:

strconv.Atoi("10")
returns:

10, nil
Example:

strconv.Atoi("hello")
returns an error.

That's why we check:

if err != nil {
	http.Error(w, "Invalid value for a", http.StatusBadRequest)
	return
}
Step 3: Choose the operation
switch op {
This is like:

if op == "add"
else if op == "subtract"
else if op == "multiply"
Example:

op = "add"
runs:

result = a + b
Step 4: Return the answer
fmt.Fprintf(w, "Result: %d", result)
If:

result = 15
the client receives:

Result: 15
Test Cases
Addition
URL:

http://localhost:8080/calculate?op=add&a=10&b=5
Response:

Result: 15
Subtraction
URL:

http://localhost:8080/calculate?op=subtract&a=10&b=5
Response:

Result: 5
Multiplication
URL:

http://localhost:8080/calculate?op=multiply&a=10&b=5
Response:

Result: 50
Invalid Number
URL:

http://localhost:8080/calculate?op=add&a=hello&b=5
Response:

Invalid value for a
Status Code:

400 Bad Request
Unknown Operation
URL:

http://localhost:8080/calculate?op=divide&a=10&b=5
Response:

Unknown operation
Status Code:

400 Bad Request
The most important lesson
Query parameters always arrive as strings:

aStr := r.URL.Query().Get("a")
Even if the URL contains:

a=10
Go still gives you:

"10"
not:

10
Whenever you need to perform arithmetic, you'll usually convert with:

strconv.Atoi()
This pattern appears constantly in Go web APIs.


Exercise 5: User-Agent Echo (Reading Headers)
Goal: Create an /agent route that reads the incoming browser or client header details and echoes it back in plain text: "You are visiting us using: [User-Agent Info]".
Key Tasks:
Inspect request headers using r.Header.Get("User-Agent").
Handle instances where the header might be missing or empty.

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
This exercise introduces the 301 status code, which tells clients:

"This resource has permanently moved to a new location."


Go HTTP Server Cheat Sheet
Core Standard Library Functions
1. Routing & Server Management (net/http)
http.HandleFunc(pattern string, handler func(ResponseWriter, *Request))
Registers a handler function for a specific URL path pattern.
http.ListenAndServe(addr string, handler Handler) error
Starts an HTTP server on the specified address (e.g., ":8080"). Passing nil uses the default system router (DefaultServeMux).
http.Error(w ResponseWriter, error string, code int)
A helper that sends a specific error message string and numeric HTTP status code back to the client, automatically ending the request lifecycle safely.
http.Redirect(w ResponseWriter, r *Request, url string, code int)
Sends an HTTP redirect status code (like 301 or 302) forcing the client's browser to jump to a new target URL path.
2. Reading Inputs & Formatting (io, strconv, fmt)
io.ReadAll(r io.Reader) ([]byte, error)
Reads all remaining data from an input stream (like r.Body) until it hits the end of the file/stream (EOF). Returns a byte slice.
strconv.Atoi(s string) (int, error)
Stands for "ASCII to Integer". Converts a text string into a native numeric int. Returns an error if the text contains non-numeric characters.
fmt.Fprintf(w io.Writer, format string, a ...any)
Formats data according to a template string and writes the output directly into an open network socket stream or file (w).



Request Context Breakdown (*http.Request)
When a client hits your server, all incoming information is bundled inside the pointer to the http.Request struct (usually named r):
Struct Field / Method
Purpose / Explanation
Example Usage
r.Method
A string representing the incoming HTTP type. Always use standard library constants for comparisons.
if r.Method == http.MethodPost
r.Body
An open stream containing data uploaded via POST/PUT requests. Always close it after reading to prevent memory leaks.
defer r.Body.Close()
r.URL.Query()
Parses the raw URL query string (everything after the ?) and returns a map-like structure (Values).
queryMap := r.URL.Query()
r.URL.Query().Get(key)
Fetches the value of a specific query parameter. Returns an empty string "" if the key does not exist.
name := r.URL.Query().Get("name")
r.Header.Get(key)
Fetches the value of a specific HTTP request header (case-insensitive). Returns "" if missing.
token := r.Header.Get("X-API-Key")





Response Management (http.ResponseWriter)
The http.ResponseWriter interface (usually named w) is your pipeline to send data back to the client. Keep this execution order in mind:
Modify Headers First: Call w.Header().Set("Key", "Value") before anything else if changing content types or metadata.
Write Status Code Second: Call w.WriteHeader(http.StatusCreated) if returning something other than a standard 200 OK.
Write Body Content Last: Call w.Write([]byte) or fmt.Fprintf(w) to push actual visual content to the user. Writing to the body locks your headers and status code automatically.


This is a solid cheat sheet. There are just a few points I'd refine so your mental model matches how Go's HTTP server actually behaves.

1. r.Body does not always need Close()
You wrote:

defer r.Body.Close()
In an HTTP server handler, Go's server closes the request body for you after the handler returns.

So this is usually unnecessary:

func Handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	_ = body
	_ = err
}
You'll see defer r.Body.Close() more often when making HTTP requests as a client with http.Get() or http.Client.

2. Writing the body automatically sends 200 OK
These two are equivalent:

w.Write([]byte("Hello"))
and

w.WriteHeader(http.StatusOK)
w.Write([]byte("Hello"))
because Write() automatically sends:

HTTP/1.1 200 OK
if no status code has been written yet.

3. Once you write the body, headers are locked
This is extremely important.

Wrong:

w.Write([]byte("Hello"))
w.Header().Set("Content-Type", "application/json")
The header change won't take effect.

Correct:

w.Header().Set("Content-Type", "application/json")
w.Write([]byte("Hello"))
Think:

Headers
↓
Status Code
↓
Body
4. http.Error() ends the handler logically, not physically
This is a common beginner mistake.

Wrong:

if r.Method != http.MethodGet {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

fmt.Fprint(w, "Hello")
The response may become corrupted because execution continues.

Correct:

if r.Method != http.MethodGet {
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	return
}

fmt.Fprint(w, "Hello")
Always:

http.Error(...)
return
5. Most-used status codes
You'll use these constantly:

http.StatusOK                  // 200
http.StatusCreated             // 201
http.StatusMovedPermanently    // 301
http.StatusBadRequest          // 400
http.StatusUnauthorized        // 401
http.StatusForbidden           // 403
http.StatusNotFound            // 404
http.StatusMethodNotAllowed    // 405
http.StatusInternalServerError // 500
6. The request lifecycle
This is the model I'd memorize:

Client Request
     ↓
Router (HandleFunc)
     ↓
Handler Function
     ↓
Check Path
     ↓
Check Method
     ↓
Read Query/Header/Body
     ↓
Process Data
     ↓
Set Headers
     ↓
Write Status Code
     ↓
Write Response Body
     ↓
Response Sent
7. The five things you'll use in almost every handler
r.URL.Path
r.Method
r.URL.Query().Get(...)
r.Header.Get(...)
r.Body
And on the response side:

w.Header().Set(...)
w.WriteHeader(...)
w.Write(...)
If you can comfortably use those eight things, you're already equipped to build the majority of beginner-to-intermediate Go web APIs, including projects like ascii-art-web and simple REST services.


