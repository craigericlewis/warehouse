#!/bin/sh
gunicorn builder:app -w 2 --threads 2 -b 0.0.0.0:5000
