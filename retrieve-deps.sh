#!/bin/bash
mkdir -p static/js/
curl -L -o static/js/postcards-html.js "https://unpkg.com/@dotpostcard/postcards-html?module"
