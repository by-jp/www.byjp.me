#!/bin/bash
mkdir -p static/js/
curl -L -o static/js/postcards-html.js "https://unpkg.com/@dotpostcard/postcards-html?module"
curl -L -o static/js/mermaid.js "https://unpkg.com/mermaid@10/dist/mermaid.min.js"
