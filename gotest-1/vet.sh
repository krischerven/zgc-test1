#!/bin/bash -e
echo "Running go vet:"
x=$(go vet ./...  2>&1 | perl -wln -M'Term::ANSIColor' -e 'print "\e[1;91m", "$_", "\e[0m"')
if [ -z "$x" ]; then
	echo "  No issues found ✔️"
else
	echo "$x"
fi
