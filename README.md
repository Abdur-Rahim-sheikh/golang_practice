## At this point
Instead of testing go-routine in main logic we test it seperate at `main.go` so we get full flexibility.

And this part can be simply tested via `go run main.go` from `backend` repo.

## Database
Now we have the need to use `database` and `pg-admin` to monitor the database.

So this is the best time to dockerize them.
You just need to run 
```bash
docker compose up --watch
```
So you can run and modify code without issue.
You also need not to restart the `main.go` file as I have set to auto restart the project of code changes.

*I have intentionally left out the client side, as from now we just use the backend and `postman` to test it out.*
## Project running

### Client
```bash
cd client
npm start
```
You mostly need to checkout of for `client/src/pages/api.js` where all the backend communications are written.

### Backend
```bash
cd backend
go run main.go
```


