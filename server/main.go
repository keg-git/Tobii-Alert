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

func shutoff(w http.ResponseWriter, r *http.Request) {
    speaker.Clear()
}
func main() {

    http.HandleFunc("/", simplealert)
    http.HandleFunc("/suction", suction)
    http.HandleFunc("/bathroom", bathroom)
    http.HandleFunc("/readjust", readjust)
    http.HandleFunc("/shutoff", shutoff)

    fmt.Print("Starting Server\n")
    log.Fatal(http.ListenAndServe(":8080", nil))
    // alert()
}
