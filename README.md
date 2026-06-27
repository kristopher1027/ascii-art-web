ascii-art-web
Description
A web-based ASCII art generator built with Go. It serves an HTML interface where users can convert plain text into ASCII art using different banner styles (standard, shadow, thinkertoy). The server handles form submissions, renders the ASCII art using pre-loaded banner files, and returns the result in the same page.

Authors
[Christopher Okoh]
Usage
Run the server
go run main.go
Then open your browser at: http://localhost:8080

Steps
1. Enter your text in the input field.
2. Select a banner style: Standard, Shadow, or Thinkertoy.
3. Click Generate ASCII Art.
4. The result is displayed on the same page below the form.
5. Implementation Details
Algorithm
Banner Loading (loadBanner): Reads a .txt banner file from the banners/ directory. Each file contains 95 characters (ASCII 32–126), where each character occupies 8 lines followed by 1 blank line (9 lines total). The function builds a map[rune][]string mapping each character to its 8-line ASCII representation.

Art Generation (generateArt): Iterates over each line of the input (split by \n). For each line, it builds 8 rows of output by concatenating the corresponding row from each character's ASCII representation. Empty input lines produce a single blank line in the output.

HTTP Endpoints:

GET / — Serves the main HTML page (200 OK). Returns 404 if the path is anything other than /.
POST /ascii-art — Accepts form data (text, banner), validates both, loads the banner, generates the ASCII art, and returns the result rendered in the same template. Returns appropriate HTTP status codes (400, 404, 500) on errors.
Template Rendering: Uses html/template with a single template file (templates/index.html). A PageData struct carries the result, input echo, selected banner, and any error message to the template.

Project Structure
ascii-art-web/
├── main.go
├── README.md
├── Handlers
├── generate.go
├── banners/
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
└── templates/
    └── index.html
    └── home.hmtl
