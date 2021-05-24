$db = "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
$schema = "public"
$table = "fnd_user"
$svc = "foundation"
$svcAbbr = "fnd"
$component = "A"
go run ./cmd/generate-code --db=$db --schema=$schema --table=$table --svc=$svc --svcAbbr=$svcAbbr --component=$component

