module bitbucket.org/kravalsergey/ecat

go 1.17

replace bitbucket.org/kravalsergey/gocurl => ../gocurl

require (
	bitbucket.org/kravalsergey/gocurl v0.0.0-00010101000000-000000000000
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/kraser/errorshandler v0.0.0-20181012014344-40a6026a0d12
)

require (
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/kraser/logger v0.0.0-20181013171132-dd9ebf86848a // indirect
	github.com/sirupsen/logrus v1.8.1 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
)

replace github.com/kraser/errorshandler => ../errorshandler

replace github.com/kraser/logger => ../logger
