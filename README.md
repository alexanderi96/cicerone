# Cicerone

Project started following this guide:
https://astaxie.gitbooks.io/build-web-application-with-golang/en/

Cicerone is basically a booking app for guided tours
Register and become a Cicerone to create a new event, it could be a tour of your city, or just a meetup like event.
Or if you just want to use the app and partecipate to something interesting register as a Globetrotter, you could upgrade to cicerone later on.

# How To Use?

1. Clone the repo
`git clone github.com/alexanderi96/cicerone`
   
2. Create the db file:
`cat schema.sql | sqlite3 cicerone.db`

3. Build the source
`go build`
   
4. Remember to set the environment variable for the session cookie store:
   	Under windows: 
`$Env:CICERONE_SESSION_KEY="your super secret key"`

4. Run
`./cicerone`

Open localhost:8081 and try it

You can change the port in the config file
