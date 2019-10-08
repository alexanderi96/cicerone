echo "changing directory to Cicerone"
cd $GOPATH/src/github.com/alexanderi96/cicerone
echo "creating table"
cat schema.sql | sqlite3 tasks.db
echo "building the go binary"
go build -o cicerone

echo "starting the binary"
./cicerone
