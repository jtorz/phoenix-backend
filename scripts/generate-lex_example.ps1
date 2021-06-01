$db = "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
$schema = "public"
$filterPrefix = ""
$testPkg = "github.com/jtorz/phoenix-backend/app/config/testconfig"
go run ./cmd/generate-lex --db=$db --schema=$schema --filterPrefix=$filterPrefix --overwrite=true --testPkg=$testPkg