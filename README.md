# cicerone

Cicerone is basically a booking app for guided tours!
Register and become a Cicerone to create a new event, it could be a tour of your city, or just a meetup like event.
Or if you just want to use the app and partecipate to something interesting register as a Globetrotter, you could upgrade to cicerone later on.

# How To Use?

Via script: bash install.sh

This will generate the binary and set up the database. If you want, you can copy the binary and the public folder into a folder of your choice.

Manually:

    go get github.com/thewhitetulip/Tasks
    change dir to the respective folder and create the db file: cat schema.sql | sqlite3 tasks.db
    run go build
    ./Tasks
    open localhost:8081

You can change the port in the config file
