echo "changing directory to Cicerone"
echo "cd $GOPATH/src/gitlab.com/alexanderi96/cicerone"
echo "creating table"
cat schema.sql | sqlite3 cicerone.db

#echo "building the go binary"
#go build -o cicerone

#echo "starting the binary"
#./cicerone

echo -e "\ndatabase created, now build the project with: \n\ngo build -o cicerone \n\nand run it with ./cicerone"
