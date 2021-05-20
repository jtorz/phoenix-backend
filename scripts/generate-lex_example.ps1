$db = "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
$schema = "public"
$filterPrefix = ""
go run ./cmd/generate-lex --db=$db --schema=$schema --filterPrefix=$filterPrefix --overwrite=true