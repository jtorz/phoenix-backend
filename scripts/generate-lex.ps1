$db = $env:PHOENIX_DB_MAIN_CONNECTION
$schema = "public"
$filterPrefix = ""
$testPkg = "github.com/jtorz/phoenix-backend/app/config/testconfig"
go run ./cmd/generate-lex --db=$db --schema=$schema --filterPrefix=$filterPrefix --overwrite=true --testPkg=$testPkg
