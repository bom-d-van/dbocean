# jsx --watch ui/js ui/ &
# find ui | entr -d go-bindata -o ui.go -ignore .module-cache -dev ui/... &
# find . -iname '*.go' -o -iname '*.js' -not -path './node_modules/*' -not -path './ui/.module-cache/*' | entr -dr ./run.sh

function run_gulp {
	while sleep 1; do
		gulp
	done
}

run_gulp &

gulp_pid=$(echo $!)
trap "kill $gulp_pid; pkill gulp; exit" SIGINT

while sleep 1; do
	# ls src/*.rb | entr -d rake
	find . ../sqlkungfu -iname '*.go' | entr -dr rungoapp.sh dbocean
done
