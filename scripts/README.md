very basic `connectivity check` that you can run on prospective customer delegate machines to test connectivity to
`S3` or to `Harness App` to see if they need to open firewall ports. 
this is just a beginning, we (or you) can add more to it as necessary

it’s a `bash script` that you can run on their machines. 
if needed, add more `test` steps to verify URL connectivity as done in the script
```
test " - Test connecting to [Harness API server]...                   " "https://api.harness.io/api/version"
```
Specify the decription and the URL in the tuple.
