package main

import(
    "fmt"
    "net/http"
    "github.com/gopxl/beep"
    "github.com/gopxl/beep/mp3"
    "github.com/gopxl/beep/speaker"
    "time"
    "log"
    "os"
)

func alert(reason string) {
    tunes, err := os.Open("/home/keg/Music/notification/alarm.mp3")
    if err != nil {
        log.Fatal(err)
    }

    streamer, format, err := mp3.Decode(tunes)
    if err != nil {
        log.Fatal(err)
    }
    defer streamer.Close()

    loop := beep.Loop(-1, streamer)
    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

    speaker.Play(loop)

    fmt.Print("Problem: ", reason)
    fmt.Scanln()
    // fmt.Print("Stopped go see Dad\n")

    speaker.Clear()
    return
}

func alertpage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./html/index.html")
}

func simplealert(w http.ResponseWriter, r *http.Request) {
    alert("somethings up IDK what")
	log.Println("misc")
}

func suction(w http.ResponseWriter, r *http.Request) {
    alert("suction")
	log.Println("suction")
}

func readjust(w http.ResponseWriter, r *http.Request) {
    alert("readjust")
	log.Println("readjust")
}

func pee(w http.ResponseWriter, r *http.Request) {
    alert("needs to pee")
	log.Println("pee")
}

func bathroom(w http.ResponseWriter, r *http.Request) {
    alert("needs to take a shit")
	log.Println("shit")
}

func bedtime(w http.ResponseWriter, r *http.Request) {
    alert("Bedtime")
	log.Println("bedtime")
}

func getup(w http.ResponseWriter, r *http.Request) {
    alert("time to get up")
	log.Println("getup")
}

func chair(w http.ResponseWriter, r *http.Request) {
    alert("chair is having issues")
	log.Println("chair issue")
}

func bed(w http.ResponseWriter, r *http.Request) {
    alert("bed is having issues")
	log.Println("bed issue")
}

func shutoff(w http.ResponseWriter, r *http.Request) {
    speaker.Clear()
}

func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func main() {
	file, err := openLogFile("/home/keg/src/Tobii-Alert/server/logs/serverLogs.log")
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(log.LstdFlags | log.Ltime)

    server := http.FileServer(http.Dir("/home/keg/src/Tobii-Alert/server/static"))
    http.Handle("/static/", http.StripPrefix("/static/", server))

    http.Handle("/", http.StripPrefix("/", server))
    http.HandleFunc("/alert", simplealert)
    http.HandleFunc("/suction", suction)
    http.HandleFunc("/readjust", readjust)
    http.HandleFunc("/pee", pee)
    http.HandleFunc("/bathroom", bathroom)
    http.HandleFunc("/bedtime", bedtime)
    http.HandleFunc("/getup", getup)
    http.HandleFunc("/chair", chair)
    http.HandleFunc("/bed", bed)
    http.HandleFunc("/shutoff", shutoff)

    fmt.Print("Starting Server\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
    // alert()
}


