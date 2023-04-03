docker build -t go-rest-api .
docker run go-rest-api
docker-compose up
% brew install go-task
task build
% brew install golangci-lint
task lint
task run

# git add -A
# git commit -m "Update with docker env"
# git push