# golang-firebase-example

This is a sample respository to learn how we can integrate firebase realtime db operations

## before creating database
- Goto [firebase](https://console.firebase.google.com/)
- Create new project
- Project Overview -> Settings Icon -> Projct Settings -> Service Accounts -> Click on generate new private key
- copy the private key file into project folder and renamed as **serviceAccountKey.json**

## Steps to configure Firebase realtime db
- Realtime database -> create database -> choose location -> select mode (locked mode / test mode)
- It will redirect to another window and you can see the database url there.

    ```
    Choose a region for the database. Depending on your choice of region, the database namespace will be of the form <databaseName>.firebaseio.com or <databaseName>.<region>.firebasedatabase.app
    ```


## run
```go
gr main.go --db-url <firebase_realtime_db_url> --add
gr main.go --db-url <firebase_realtime_db_url> --update
gr main.go --db-url <firebase_realtime_db_url> --delete
```