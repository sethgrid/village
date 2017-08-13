# village

`village` is the communtity website for Mountain Home Village. To run the site: `go run main.go` and go to `localhost:8100`.


## Deploy Notes

I have monit and systemd set up to run the site with default settings.

To deploy:
    - Make changes, save, push up changes to github.
    - `GOOS=linux go build`
    - `scp -r * $MYIP:~/mountainhomevillage`
    - `ssh $MYIP` and `sudo service mountainhomevillage restart`


## Integrations

I propose that we sign up for [DonerBox](https://donorbox.org/orgs/new). They have reasonable fees and are free for the first $1k in donations per month. That should do really well for us.

## Ideas

 - Show upcoming event details from the calendar in the side bar
 - Allow users to sign up or contact Dan for community watch
 - Have a page for more details on financials, upcoming projects, etc