module github.com/GoogleCloudPlatform/testgrid

replace k8s.io/test-infra => github.com/mpherman2/test-infra v0.0.0-20201203221419-259733261dca

require (
	cloud.google.com/go/storage v1.12.0
	github.com/client9/misspell v0.3.4
	github.com/fvbommel/sortorder v1.0.1
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.2
	github.com/google/uuid v1.1.1
	github.com/hashicorp/go-multierror v1.1.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/stretchr/testify v1.5.1
	github.com/tektoncd/pipeline v0.13.1-0.20200625065359-44f22a067b75 // indirect
	google.golang.org/api v0.32.0
	google.golang.org/genproto v0.0.0-20200921151605-7abf4a1a14d5
	google.golang.org/grpc v1.32.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	sigs.k8s.io/yaml v1.2.0
)

go 1.14
