DEL c:\Users\Administrator\go\src\github.com\cloudfoundry\cli-acceptance-tests\cf.exe
bitsadmin.exe /transfer "DownloadStableCLI" https://s3.amazonaws.com/go-cli/builds/cf-windows-amd64.exe c:\Users\Administrator\go\src\github.com\cloudfoundry\cli-acceptance-tests\cf.exe

go get -u github.com/cloudfoundry/cli-acceptance-tests/...

SET CLIACCEPTANCEPATH=%GOPATH%\src\github.com\cloudfoundry\cli-acceptance-tests
SET PATH=%CLIACCEPTANCEPATH%;%PATH%;C:\Program Files\cURL\bin
SET CONFIG=%CD%\config.json
SET LOCAL_GOPATH=%CLIACCEPTANCEPATH%\Godeps\_workspace
MKDIR %LOCAL_GOPATH%\bin
SET GOPATH=%LOCAL_GOPATH%;%GOPATH%
SET PATH=%LOCAL_GOPATH%\bin;%PATH%

go install -v github.com/onsi/ginkgo/ginkgo
ginkgo.exe -r -slowSpecThreshold=120 ./gats
