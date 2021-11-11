env GOOS=linux go build main.go
zip postresults.zip main
mv postresults.zip ../dist/
rm main