package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fogleman/gg"
)

var webPort int

// Args ...
type Args []interface{}

// ImgCommand ...
type ImgCommand struct {
	Cmd  string `json:"cmd"`
	Args Args   `json:"args,omitempty"`
}

// ImageRequest ...
type ImageRequest struct {
	CanvasWidth  int          `json:"canvas_width"`
	CanvasHeight int          `json:"canvas_height"`
	ImgCommands  []ImgCommand `json:"img_commands"`
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	imgReq := ImageRequest{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&imgReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("bad request"))
		return
	}

	b, _ := json.Marshal(imgReq)
	log.Println("got request:", string(b))

	dc := gg.NewContext(imgReq.CanvasWidth, imgReq.CanvasHeight)

	for idx, cmd := range imgReq.ImgCommands {
		func() {
			defer func() {
				e := recover()
				if e != nil {
					log.Printf("error processing command at index %d: %v\n", idx, e)
				}
			}()
			switch cmd.Cmd {
			case "SetRGB":
				dc.SetRGB(cmd.Args[0].(float64), cmd.Args[1].(float64), cmd.Args[2].(float64))
			case "SetRGBA":
				dc.SetRGBA(cmd.Args[0].(float64), cmd.Args[1].(float64), cmd.Args[2].(float64), cmd.Args[3].(float64))
			case "DrawRectangle":
				dc.DrawRectangle(cmd.Args[0].(float64), cmd.Args[1].(float64), cmd.Args[2].(float64), cmd.Args[3].(float64))
			case "DrawLine":
				dc.DrawLine(cmd.Args[0].(float64), cmd.Args[1].(float64), cmd.Args[2].(float64), cmd.Args[3].(float64))
				dc.Stroke()
			case "SetLineWidth":
				dc.SetLineWidth(cmd.Args[0].(float64))
			case "DrawCircle":
				dc.DrawCircle(cmd.Args[0].(float64), cmd.Args[1].(float64), cmd.Args[2].(float64))
			case "Fill":
				dc.Fill()
			case "Stroke":
				dc.Stroke()
			default:
				log.Println("invalid command", cmd.Cmd)
			}
		}()
	}
	w.Header().Set("Content-Type", "image/png")
	dc.EncodePNG(w)
}

func main() {
	flag.IntVar(&webPort, "port", 8088, "listening port for web server")
	flag.Parse()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", webPort),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.HandleFunc("/", imgHandler)

	fmt.Printf("Listening on http://localhost:%d\n\n", webPort)

	log.Fatal(s.ListenAndServe())
}
