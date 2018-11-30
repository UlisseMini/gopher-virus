// TODO (36): Figure out the logging module and make it log to a file

package main

import (
	"github.com/reujab/wallpaper"
	"io/ioutil"
	"math/rand"
	"os"
	"os/user"
	"strconv"
	"time"
)

func main() {
	// Args
	var sleeptime int
	if len(os.Args) < 2 {
		sleeptime = 60
	} else {
		sleeptime, err := strconv.Atoi(os.Args[1])
		if err != nil {
			println("Failed to convert args[1] to int")
			sleeptime = 60
		}
		println("Sleeptime set to " + strconv.Itoa(sleeptime))
	}
	// Get user information
	log("Getting current user...")
	usr, err := user.Current()
	handle(err)

	// Change directory to the gophers directory
	gopherdir := usr.HomeDir + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\gophers\\"

	// Create a list "gophers" and store all the gopher files inside it
	gophers, err := ioutil.ReadDir(gopherdir)
	handle(err)

	// Every 5 minutes pick a random gopher and change the background picture
	var choice os.FileInfo

	for {
		// Pick a random one and change the desktop background
		rand.Seed(time.Now().Unix())
		choice = gophers[rand.Intn(len(gophers))]
		log(choice.Name())
		wallpaper.SetFromFile(gopherdir + choice.Name())

		// Wait 60 seconds or what arg[1] was
		time.Sleep(time.Duration(sleeptime) * time.Second)
	}
}

// since the window is hidden i'll write the error to a file
func handle(err error) {
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

// make this go to a file later
func log(msg string) {
	println(msg)
}
