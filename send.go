package main

import (
	"fmt"
	"github.com/googollee/go-gcm"
)

func main() {
	client := gcm.New("AIzaSyDXmMGUx4XNFhMx7rx-EgxOpRhUhOKDlpw")

	//dA89VAlE8CY:APA91bEA-Mk8guvD8Geuo6tQQ0Aw73wTCrAkWtnD_DhtbXLAevoTNDrB6Q3vAJ1xnOPNo0tV_8_V8EwrP4DNF7Et-dMsSA1Of47yhUNqGlWS4m6p86gjGvRtMwZUYku_g9WLARwegp32
	//APA91bFaEbRSP3o4_O4pi97rYPoVhJWAAadK9toVfX9V33BveHqol9N1jIstjURUyrlPxAgt_S-gk-i6ZUfoI4QDSISvPhQ7tPDTudRAf9BmBxdJ4v-uiikc7BBlQWNwqeXtVmg9ngO4
	load := gcm.NewMessage("dA89VAlE8CY:APA91bHQrsLKJqv_F-M3Zuv4a3Kl9ptNo-uFkmM1jMVViCDKrAIfECKTunvkAkrcwTXyUQ3SxOSSPXffDrykzVUDU75x5Ktyofz2yyc2-fBtxoRey3JA3FJoslLuQAFL_zP3CVAl-Ah5")
	load.AddRecipient("abc")
	load.SetPayload("title", "VertragsNavi")
	load.SetPayload("message", "Halllo")
	load.SetPayload("extra", "{\"url\":\"someurl.js\"}")

	load.CollapseKey = "demo"
	load.DelayWhileIdle = true
	load.TimeToLive = 10

	resp, err := client.Send(load)

	fmt.Printf("id: %+v\n", resp)
	fmt.Println("err:", err)
	fmt.Println("err index:", resp.ErrorIndexes())
	fmt.Println("reg index:", resp.RefreshIndexes())
}
