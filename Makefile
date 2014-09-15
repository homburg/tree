benchmark:
	go test -bench .

time:
	cat data.csv | time -p tree , > /dev/null
