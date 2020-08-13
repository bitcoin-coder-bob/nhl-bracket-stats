This tool can be used to download the NHL.com Stanley Cup Playoff Brackets and run basic stats on them.

In the directory create a folder called `brackets`. Run `go run main.go` and the `brackets` folder will begin
populating with 458223 brackets, each saved into their own file of the form `brackets/<bracket-number>.txt`.

This has taken over 12 hours to run and will occupy approximately 2GB.

Running `go run analyze.go` will print to the console the stats. As an example, I saved the output to 
`ROUND-OF-16.txt` and `STANLEY-CUP-FINALS.txt` when analyzing on 300k brackets.