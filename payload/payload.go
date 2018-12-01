package main

import (
	"github.com/reujab/wallpaper"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/user"
	"strconv"
	"time"
)

const (
	logfile = "payload.log"
)

var (
	logger    *log.Logger
	gopherdir string
	sleeptime int
)

func init() {
	// Args
	if len(os.Args) < 2 {
		sleeptime = 60
	} else {
		sleeptime, err := strconv.Atoi(os.Args[1])
		if err != nil {
			logger.Println("Failed to convert args[1] to int")
			sleeptime = 60
		}
		logger.Println("Sleeptime set to " + strconv.Itoa(sleeptime))
	}

	// Get user information
	logger.Println("Getting current user...")
	usr, err := user.Current()
	must(err)
	// set the gophers directory
	gopherdir = usr.HomeDir + `\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup\gophers\`

	// Create logger
	file, err := os.OpenFile(logfile,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger = log.New(file, "", 2)
}

func main() {
	// Create a list "gophers" and store all the gopher files inside it
	gophers, err := ioutil.ReadDir(gopherdir)
	must(err)

	// Every sleeptime seconds pick a random gopher and change the background picture
	var choice os.FileInfo
	for {
		// Pick a random one and change the desktop background
		rand.Seed(time.Now().Unix())
		choice = gophers[rand.Intn(len(gophers))]
		wallpaper.SetFromFile(gopherdir + choice.Name())

		// Wait 60 seconds or what arg[1] was
		time.Sleep(time.Duration(sleeptime) * time.Second)
	}
}

// since the window is hidden i'll write the error to a file
func must(err error) {
	if err != nil {
		logger.Print("[FATAL] ")
		logger.Println(err.Error())
		os.Exit(1)
	}
}

// since the window is hidden i'll write the error to a file
func handle(err error) {
	if err != nil {
		logger.Println(err.Error())
	}
}
