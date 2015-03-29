#
# a script to use with entr in development
# find app lib db config ../../qor/qor | entr -r scripts/run.sh
#

set -e

echo "----------------------------"

# echo go-bindata -o ui.go -ignore .module-cache -dev ui/...
# go-bindata -o ui.go -ignore .module-cache -dev ui/...

echo go install
go install

app=dbocean
if pgrep $app > /dev/null; then
	echo kill $app
	pkill $app
fi

echo running $app
echo "----------------------------"
$app