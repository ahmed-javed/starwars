Prepopulated mysql container
================================

The mysql container image contain an init script that will execute everything in `/docker-entrypoint-initdb.d/`

see `Initializing a fresh instance` @ https://hub.docker.com/_/mysql/

Run this initialization in a multi-stage build and copy the initialized DB in the new image :D

Execute Commands!
======

Within this directory run

```
}> docker build --t mysql57 .
...

}> docker run -d --rm --name toptal-mysql --expose 3306 -p 3306:3306 mysql57
...

}> docker logs toptal-mysql
...

}> docker run -it --rm --link toptal-mysql mysql57 mysql -htoptal-mysql -uroot -proot toptal -e "select * from user;"

}> Or ignore the last command and connect to MySQL using workbench. Host=localhost, User=root, Password=root.
```
