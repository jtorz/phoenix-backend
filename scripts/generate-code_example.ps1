$db = $env:PHOENIX_DB_MAIN_CONNECTION
$schema = "public"
$table = "core_user"
$svc = "core"
$svcAbbr = "core"
$component = "A"
go run ./cmd/generate-code --db=$db --schema=$schema --table=$table --svc=$svc --svcAbbr=$svcAbbr --component=$component --write=true


