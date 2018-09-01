## CSV READ

    Supports loading csv and xls files.
    # csv files use standard library "encoding/csv"
    # xls files use "github.com/tealeg/xlsx"

    GoLevelDB is used as database.
    # github.com/syndtr/goleveldb

    Testify is used for testing.
    # github.com/stretchr/testify

    Docker single stage production build is used.

    Default configuration can be changed in 
    # .env
    # src/conf/conf.yaml

### Running
    make
### Clean
    make clean

### Testing
    make test