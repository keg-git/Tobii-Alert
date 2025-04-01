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

    fmt.Print("Dad Currently needs you for ", reason, "!\nGet off your ass now")
    fmt.Scanln()
    fmt.Print("Stopped go see Dad\n")

    speaker.Clear()
    return
}

func alertpage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./pages/index.html")
}

func simplealert(w http.ResponseWriter, r *http.Request) {
    alert("something")
}

func suction(w http.ResponseWriter, r *http.Request) {
    alert("suction")
}

func readjust(w http.ResponseWriter, r *http.Request) {
    alert("readjusting")
}

func bathroom(w http.ResponseWriter, r *http.Request) {
    alert("taking a shit")
}

func bedtime(w http.ResponseWriter, r *http.Request) {
    alert("going to bed")
}

func getup(w http.ResponseWriter, r *http.Request) {
    alert("getting up")
}

func shutoff(w http.ResponseWriter, r *http.Request) {
    speaker.Clear()
}
func main() {

    server := http.FileServer(http.Dir("./pages"))
    http.Handle("/pages/", http.StripPrefix("/pages/", server))

    http.HandleFunc("/", alertpage)
    http.HandleFunc("/alert", simplealert)
    http.HandleFunc("/suction", suction)
    http.HandleFunc("/readjust", readjust)
    http.HandleFunc("/bathroom", bathroom)
    http.HandleFunc("/bedtime", bedtime)
    http.HandleFunc("/getup", getup)
    http.HandleFunc("/shutoff", shutoff)

    fmt.Print("Starting Server\n")
    log.Fatal(http.ListenAndServe(":8081", nil))
    // alert()
}

