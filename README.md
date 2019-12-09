# Cicerone

Project started following this guide: https://astaxie.gitbooks.io/build-web-application-with-golang/en/

Cicerone is basically a booking app for guided tours!
Register and become a Cicerone to create a new event, it could be a tour of your city, or just a meetup like event.
Or if you just want to use the app and partecipate to something interesting register as a Globetrotter, you could upgrade to cicerone later on.

# How To Use?

Via script: bash install.sh

This will generate the binary and set up the database. If you want, you can copy the binary and the public folder into a folder of your choice.

Manually:

    go get github.com/alexanderi96/cicerone
    change dir to the respective folder and create the db file: cat schema.sql | sqlite3 cicerone.db
    run go build
    ./cicerone
    open localhost:8081

You can change the port in the config file

# License

The MIT License (MIT)

Copyright (c) 2015 Suraj Patil

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
