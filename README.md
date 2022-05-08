# litestream-sample

## Get Started
```
$ cd docker
$ docker-compose up -d
```

Access to http://localhost:9001 and make `mybkt` bucket to store data.

```
$ docker exec -it litestream-sample-app bash
# sqlite3 docker/datafiles/db
sqlite> CREATE TABLE fruits (name TEXT, color TEXT);
sqlite> INSERT INTO fruits (name, color) VALUES ('apple', 'red');
sqlite> INSERT INTO fruits (name, color) VALUES ('banana', 'yellow');
```

Access to minio bucket to check if db data exits.