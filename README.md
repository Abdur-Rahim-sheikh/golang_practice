## Go with Habib

```I am thankful to Habib bhai, for creating this awesome playlist on Golang with indepth explanation in simple and joyfull words.```


[Tutorial](https://github.com/Abdur-Rahim-sheikh/golang_practice)

### How this repo helps you?

About two third of the course is theoritical and very fun to go along. 

After that, Bhai starts implementation. And I felt the need to branch out and go along.

So if you are one of them. Then you sometime need to sync what's going on. At that time you can `switch` to that branch and match if you are on track.

Beside that, in database connection, Bhai used postgres and `pg-admin` for better explanation. I have dockerized the thing. No, one needs to learn docker, they just need to run `docker compose up` and backend, db and pg-admin will be working together without any issue. And pg-admin will be available at `localhost:8080` with mail: `test@abc.com` and password `test`


*I have intentionally left out the client side, as from now we just use the backend and `postman` to test it out.*


## Running the project in normal use case

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