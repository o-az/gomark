go fmt *.go && go mod tidy

rm -rf main.zip main

GOOS=linux GOARCH=amd64 go build -o main *.go

zip main.zip main

aws lambda update-function-code --function-name "$AWS_LAMBDA_FUNCTION_NAME" --zip-file fileb://main.zip
