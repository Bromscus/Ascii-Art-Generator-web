# ASCII Art Web

## Description

ASCII Art Web is a Go web application that converts text into ASCII art. It supports the `standard`, `shadow`, and `thinkertoy` banner styles and displays the generated result in a browser.

## Authors

- ahmsalem
- asami

## Usage: how to run

1. From the root project directory, start the server:

   ```bash
   go run .
   ```

2. Open http://localhost:8080 in a browser.

3. Enter text, select a banner style, and submit the form to generate ASCII art.

The application listens on port `8080`.

## Implementation details: algorithm

1. At startup, the application loads the three banner files: `standard.txt`, `shadow.txt`, and `thinkertoy.txt`.
2. The submitted text and selected style are received by the `/ascii-art` HTTP handler.
3. The selected banner file is split into lines. Each printable character is represented by 8 consecutive rows, with 9-line blocks used to locate the next character in the font data.
4. The input is processed as runes so characters are handled as Unicode code points. The escaped sequence `\\n` starts a new output line.
5. For each character, its 8 banner rows are appended to an output table.
6. The completed rows are passed to the printer, which joins them into the final ASCII-art result and returns it to the client.
