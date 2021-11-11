env GOOS=linux go build main.go
zip getresults.zip main
mv getresults.zip ../dist/
rm main