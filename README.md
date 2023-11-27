<img align="left" width="50" height="50" src="anubis.png" alt="Anubis">

# Anubis
```
go mod init anubis
go mod tidy
go run range_get.go httpx_get.go  nuclei_get.go  anubis.go
```

## Used

* https://github.com/projectdiscovery/subfinder/blob/main/v2/examples/main.go
* https://github.com/projectdiscovery/httpx/blob/main/examples/example.go
* https://github.com/projectdiscovery/nuclei/blob/dev/examples/advanced/advanced.go


## List of tools tools

* https://github.com/projectdiscovery/httpx
* https://pkg.go.dev/github.com/projectdiscovery/mapcidr
* https://github.com/projectdiscovery/dnsx
* https://github.com/maurosoria/dirsearch
* https://mxtoolbox.com/SuperTool.aspx
* https://github.com/projectdiscovery/nuclei
* https://github.com/projectdiscovery/nuclei-templates

## Full command
* After installing the go package, the binary is in `~/go/bin`

``````
./subfinder -silent -d [url] | ./httpx | ./nuclei -t nuclei-templates/
```