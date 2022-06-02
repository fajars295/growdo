C=$(env | grep MIGRATION_DATABASE | grep -oe '[^=]*$');
DATABASE="${C}?sslmode=disable"

if [ "up" = $1 ]; then
 migrate -path db/migration -database $DATABASE -verbose goto $2
fi

if [ "down" = $1 ]; then 
 migrate -path db/migration -database $DATABASE -verbose down $2
fi

if [ "create" = $1 ]; then 
 migrate create -ext sql -dir db/migration -seq $2
fi

if [ "fix" = $1 ]; then 
 migrate -path db/migration -database $DATABASE force $2
fi


if [ "migrate" = $1 ]; then
cd db/migration
search_dir=$(pwd)
count=0
for entry in "$search_dir"/*.up.sql
do
  count=$((count + 1))
  migrate -path ./ -database $DATABASE -verbose goto $count
done
fi