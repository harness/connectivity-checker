# connectivity-checker
Checks basic connectivity to a list of URLs from inside customer delegate machines

Edit the `urls.yaml` file to specify additional URLs to check.

The `yaml` file with list of URLs can be named something different, just run the checker with 
`./connectivity-checker -list <yourfile.yaml>`

To build:
`go build .`

If you add additional non-standard Go dependencies, just make sure they are included in the `vendor` directory that you can do with `govendor`
https://github.com/kardianos/govendor

