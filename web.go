package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			if r.FormValue("pin") != "1234" {
				fmt.Fprint(w, "Invalid PIN")
				return
			}
			command := fmt.Sprintf("shutdown /s /t %s", r.FormValue("time"))
			cmd := exec.Command("cmd", "/C", command)
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(w, "Error: %v", err)
				return
			}
			fmt.Fprint(w, "Shutdown command sent")
			return
		}
		fmt.Fprint(w, `
			<html>
				<body>
					<form method="POST">
						<label for="pin">Enter PIN:</label>
						<input type="password" id="pin" name="pin" required>
						<label for="time">Enter time in seconds:</label>
						<input type="number" id="time" name="time" required>
						<button type="submit">Shutdown</button>
					</form>
				</body>
			</html>
		`)
	})

	log.Fatal(http.ListenAndServe(":2576", nil))
}