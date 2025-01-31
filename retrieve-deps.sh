#!/bin/bash
mkdir -p static/js/
curl -L -o static/js/mermaid.js "https://unpkg.com/mermaid@10/dist/mermaid.min.js"
curl -L -o static/js/mathjax.js "https://unpkg.com/mathjax@3/dist/mathjax.min.js"
curl -L -o static/js/shikwasa.js "https://unpkg.com/shikwasa@2/dist/shikwasa.min.js"
curl -L -o static/css/shikwasa.css "https://unpkg.com/shikwasa@2/dist/style.css"