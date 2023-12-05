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
				<head>
					<style>
						body {
							font-family: Arial, sans-serif;
							margin: 20px;
							padding: 0;
							background-color: #f2f2f2;
						}
						.container {
							max-width: 400px;
							margin: 0 auto;
							padding: 20px;
							background-color: #fff;
							border-radius: 5px;
							box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
						}
						.form-group {
							margin-bottom: 20px;
						}
						.form-group label {
							display: block;
							margin-bottom: 5px;
							font-weight: bold;
						}
						.form-group input {
							width: 100%;
							padding: 10px;
							border: 1px solid #ccc;
							border-radius: 3px;
						}
						#butty {
							padding: 10px 20px;
							background-color: #4CAF50;
							color: #fff;
							border: none;
							border-radius: 3px;
							cursor: pointer;
							font-size: 16px;
							transition: background-color 0.3s ease;
							width: 100%;
						}
						#butty:hover {
							background-color: #45a049;
						}
					</style>
				</head>
				<body>
					<div class="container">
						<form method="POST">
							<div class="form-group">
								<label for="pin">Enter PIN:</label>
								<input type="password" id="pin" name="pin" required>
							</div>
							<div class="form-group">
								<label for="time">Enter time in seconds:</label>
								<input type="number" id="time" name="time" required>
							</div>
							<button id="butty" type="submit">Shutdown</button>
						</form>
					</div>
				</body>
			</html>
		`)
	})

	log.Fatal(http.ListenAndServe(":2576", nil))
}