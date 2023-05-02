gen-sql:
	if [ -d "data" ]; then \
    rm -r data; \
  fi;
	sqlc generate -f ./db/sqlc.yaml