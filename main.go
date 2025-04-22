package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "darwin":
		err = exec.Command("open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	default: // linux
		err = exec.Command("xdg-open", url).Start()
	}
	if err != nil {
		log.Println("無法自動開啟瀏覽器，請手動開啟：", url)
	}
}

func main() {
	const port = 8080
	const folder = "./" // 靜態檔案目錄，根據你的 wedding 資料夾而定

	fs := http.FileServer(http.Dir(folder))
	http.Handle("/", fs)

	url := fmt.Sprintf("http://localhost:%d", port)
	fmt.Println("伺服器啟動中，請打開：", url)

	go openBrowser(url)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("伺服器啟動失敗:", err)
	}
}
