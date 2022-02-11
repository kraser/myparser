module bitbucket.org/kravalsergey/ecat

go 1.17

replace bitbucket.org/kravalsergey/gocurl => ../gocurl

require (
	bitbucket.org/kravalsergey/gocurl v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/kraser/errorshandler v0.0.0-20181012014344-40a6026a0d12
	github.com/kraser/logger v0.0.0-20181013171132-dd9ebf86848a
)

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/kraser/goprice v0.0.0-00010101000000-000000000000 // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20220128215802-99c3d69c2c27 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/kraser/errorshandler => ../errorshandler

replace github.com/kraser/logger => ../logger

replace github.com/kraser/goprice => ../goprice
