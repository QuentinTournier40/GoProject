mkdir output
mkdir csv
start mosquitto -v
start go clean ./...
start go build -o ./output ./...

cd output

start api.exe
start Pressure.exe
start Temp.exe
start Wind.exe
start subscriber_api.exe
start subscriber_csv.exe