# glyph-browser

A test of creating a browser developed with Go and running in cli. It can be used with the keyboard only to navigate in the UI. Developed to understand how a browser works by making himself one.

The browser drain many principles from an article written by Paul Irish(Chrome Developer Relations) : https://web.dev/articles/howbrowserswork

---

## How it works

### Engine

- start a new pipe of process flow for each "tab"
- maintain a state of each pipe
- listen for any interrupt (cancel page loading, new page invoked,....)

### NETWORKING
- resolve uri 
- get the content of the page
- pass content to the renderer through a chan

### RENDERER
- instantiate parser for each "language" (HTML, CSS,....)
- build the dom
- pass rendered content to the UI

### UI
to be defined....