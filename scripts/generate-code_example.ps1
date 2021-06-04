$db = $env:PHOENIX_DB_MAIN_CONNECTION
$schema = "public"
$table = "fnd_user"
$svc = "foundation"
$svcAbbr = "fnd"
$component = "A"
go run ./cmd/generate-code --db=$db --schema=$schema --table=$table --svc=$svc --svcAbbr=$svcAbbr --component=$component --write=true


