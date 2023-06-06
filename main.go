package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {

	if len(os.Args) <= 1{
		log.Fatalln("Please enter argument.")
	}

	// Run Chrome browser
	service, err := selenium.NewChromeDriverService("/usr/bin/chromedriver", 9515)
	if err != nil {
		log.Fatalln(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless",
	}})

	driver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub",9515))
	if err != nil {
		log.Fatalln(err)
	}
	defer driver.Quit()

	driver.Get(fmt.Sprintf("https://www.nslookup.io/domains/%s/dns-records/",os.Args[1]))

	links, err := driver.FindElements(selenium.ByClassName, "py-1")
	if err != nil {
		log.Fatalln(err)
	}

	for _, link := range links {
		if ip, err := link.Text(); err == nil {
			parsedIP := net.ParseIP(ip)
			if parsedIP.To4() != nil {
				fmt.Printf("A records %s : %s \n", os.Args[1],ip)
				break
			}
		}
	}
}
