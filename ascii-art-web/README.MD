# ASCII Art Web

## Description

ASCII Art Web is a Go HTTP server that renders text as ASCII art from a browser. The web page accepts user text, lets the user choose one of the `standard`, `shadow`, or `thinkertoy` banners, submits the form to `/ascii-art`, and displays the generated result.

## Authors

- Student project author

## Usage: how to run

Run the server from the project root:

```sh
go run .
```

Then open:

```text
http://localhost:8080
```

Run the tests with:

```sh
go test ./...
```

## Implementation details: algorithm

The server exposes `GET /` for the main HTML page and `POST /ascii-art` for form submissions. Templates are loaded from the root `templates` directory with Go's `html/template` package.

When a request is submitted, the handler validates that the input contains printable ASCII characters or newlines, validates the selected banner, and asks the ASCII generator to render the result. The generator loads the selected banner file from `Banners`, caches the parsed banner lines, normalizes Windows newlines to Unix newlines, and splits the input into lines.

Each printable character maps to an 8-line glyph in the banner file. Since banner files include a separator line for each character, the glyph start index is calculated with:

```text
(character - 32) * 9
```

For every input line, the generator writes glyph row 0 for all characters, then glyph row 1 for all characters, continuing through row 7. Empty input lines are preserved as blank lines in the output.
