# village

`village` is the communtity website for Mountain Home Village. To run the site: `go run main.go` and go to `localhost:8100`.


## Deploy Notes

I have monit and systemd set up to run the site with default settings.

To deploy:
    - Make changes, save, push up changes to github.
    - `GOOS=linux go build`
    - `scp -r * $MYIP:~/mountainhomevillage`
    - `ssh $MYIP` and `sudo service mountainhomevillage restart`
