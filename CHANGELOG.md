## v0.3.0 (2023-01-30)

### Feat

- **manga**: add endpoint for volume creation
- **manga**: add list and get support
- **utils**: ensure all tables are present
- **models**: add relationships between tables
- **utils**: add function for setting up the db
- **utils**: add function for creating uuids
- **manga**: add bare minimums for manga endpoint
- **zerolog**: set log format based on environment
- **models**: add types for storing manga in the database

### Fix

- **chi**: specify use of id parameter
- **manga**: get manga by id using path instead of url query
- **models**: use the right type for ids
- **utils**: set uuid octets as constant

## v0.2.0 (2023-01-27)

### Feat

- **chi**: add simple heartbeat endpoint
- implement json logging with httplog
- **swagger**: implement swaggo

### Fix

- **utils**: use envfile variable correctly
- **utils**: fix a bug where file would be closed before it was scanned
- **utils**: fix a bug where file would be closed before it was scanned

## v0.1.0 (2023-01-27)

### Feat

- serve new health endpoint
- **api**: improve healthcheck endpoint
- **utils**: add function for setting environment
