#!/bin/bash

MIGRATION_PATH=resources/migrations/

migrate -path "${MIGRATION_PATH}" -database mysql://"${DATABASE_DSN}" up
