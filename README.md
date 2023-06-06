## nslookup 
In this code, the `Selenium` package is used, which can be used to get the `IP addresses` of any `domain` through the `nslookup.com` site.


#### build static command :
```
go build -ldflags "-linkmode 'external' -extldflags '-static'" -o go_nslookup
```

#### require package install chromedriver on linux :
```
apt install chromium-chromedriver
```