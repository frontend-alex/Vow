#!/usr/bin/env sh
set -eu

migrate -path db/migrations -database "${DATABASE_URL:?DATABASE_URL is required}" down
