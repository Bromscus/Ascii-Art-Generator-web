# ASCII Art Web

## Description

The project main idea is to switch the input to a stylized format called ascii format in which we have different banner styles which are supported: `standard`, `thinkertoy` and `shadow`.

## Authors

- ahmsalem
- asami

## Usage: how to run

### Option 1: Run directly with Go

1. From the root project directory, start the server:

```bash
   go run .
```

2. Open http://localhost:8080 in a browser.

3. Enter text, select a banner style, and submit the form to generate ASCII art.

### Option 2: Run with Docker

1. Build the image:

```bash
   docker build -t ascii-art-web .
```

2. Run a container from the image:

```bash
   docker run -p 8080:8080 ascii-art-web
```

3. Open http://localhost:8080 in a browser.

## Implementation details: algorithm

1. Load the three banner files.
2. Get the user's text and selected style.
3. Split the selected banner into lines.
4. Convert each character into its 8 line ASCII art.
5. Join the characters together vertically in a horizontal line if there is no new lines.
6. New lines (Enter key) are accepted.
7. Return the final ASCII-art result to the user.