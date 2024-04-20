> docker build -t go-rest-api .
> docker run go-rest-api
> docker-compose up
% brew install go-task
> task build
% brew install golangci-lint
> task lint
> task run

# git add -A
# git commit -m "Added docker any yml dependencies"
# git push

# migration:
> task run

# 2-d terminal window:
> docker ps
> docker exec -it 9ca2613ac2ea bash # container id
> psql -U postgres
> postgres=# \dt # all tables
> postgres=# \d+ comments;  # all comment's fields

> postgres=# \q
> exit

> postgres=# insert into comments (id) values('9a31bf83-28dc-4b8d-bf70-7d347a24ff2e');
> postgres=# select * from comments;

> 1-st terminal window:
> task run
> 2-d terminal window:
> curl http://localhost:8080/hello


> 1-st terminal window:
> task run
> 2-d terminal window:
> curl --location --request POST 'http://localhost:8080/api/v1/comment' --header 'Content-Type: application/json' --data-raw '{"slug": "hello", "body": "body", "author": "me"}'

> 1-st terminal window:
> task run
> 2-d terminal window:
> curl --location --request GET 'http://localhost:8080/api/v1/comment/74bdff3b-5ebb-415b-beb7-a26feb9b93c6'

> 1-st terminal window:
> task run
> 2-d terminal window:
> curl --location --request PUT 'http://localhost:8080/api/v1/comment/74bdff3b-5ebb-415b-beb7-a26feb9b93c6' --header 'Content-Type: application/json' --data-raw '{"slug": "/testing-put", "author": "Elliot Forbes", "body": "body"}'

> 1-st terminal window:
> task run
> 2-d terminal window:
> curl --location --request DELETE 'http://localhost:8080/api/v1/comment/74bdff3b-5ebb-415b-beb7-a26feb9b93c6'

> 1-st terminal window:
> task run
> 2-d terminal window:
> curl --location --request POST 'http://localhost:8080/api/v1/comment' -v --header 'Content-Type: application/json' --data-raw '{"slug": "hello", "body": "body", "author": "me"}'

> curl --location --request GET 'http://localhost:8080/api/v1/comment/d836ef56-6c58-4581-a15d-17b65a192080'

> go https://jwt.io, copy the left token after adding 'missionimpossible' signKey on the right
> 2-d terminal window: 
> curl --location --request POST 'http://localhost:8080/api/v1/comment' -v --header 'Content-Type: application/json' --header 'Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.2yPNTTY3Y5jUwYBPAJAfUc84Ybv2qPbZY_OHI7tzuug' --data-raw '{"slug": "hello", "body": "body", "author": "me"}'

> 1-st terminal window:
> task run
> 2-d terminal window:
> curl --location --request GET 'http://localhost:8080/api/v1/comments'

# testing:
> go test ./... -v
> go test -tags=integration ./... -v

> task integration-test

> task acceptance-test