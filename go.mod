module api.welooky

go 1.12

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.40.0
	golang.org/golang.org/x/tools => github.com/golang/text v0.3.2
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190513172903-22d7a77e9e5f
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190510132918-efd6b22b2522
	golang.org/x/image => github.com/golang/image v0.0.0-20190523035834-f03afa92d3ff
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190607214518-6fa95d984e88
	golang.org/x/net => github.com/golang/net v0.0.0-20190509222800-a4d6f7feada5
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190402181905-9f3314589c9a
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190509141414-a5b02f93d862
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190513233021-7d589f28aaf4
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.6.0
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190605220351-eb0b1bdb6ae6
	google.golang.org/grpc => github.com/grpc/grpc-go v1.21.1
	k8s.io/api => github.com/kubernetes/api v0.0.0-20190602205700-9b8cae951d65
	k8s.io/apimachinery => github.com/kubernetes/apimachinery v0.0.0-20190607205628-5fbcd19f360b
	k8s.io/client-go => github.com/kubernetes/client-go v11.0.0+incompatible
	k8s.io/gengo => github.com/kubernetes/gengo v0.0.0-20190327210449-e17681d19d3a
	k8s.io/klog => github.com/kubernetes/klog v0.3.2
	k8s.io/kube-openapi => github.com/kubernetes/kube-openapi v0.0.0-20190603182131-db7b694dc208
	sigs.k8s.io/structured-merge-diff => github.com/kubernetes-sigs/structured-merge-diff v0.0.0-20190525122527-15d366b2352e
	sigs.k8s.io/yaml => github.com/kubernetes-sigs/yaml v1.1.0
)

require (
	github.com/Joker/jade v1.0.0 // indirect
	github.com/Shopify/goreferrer v0.0.0-20181106222321-ec9c9a553398 // indirect
	github.com/ahmetb/go-linq v3.0.0+incompatible
	github.com/ajg/form v1.5.1 // indirect
	github.com/aymerick/raymond v2.0.2+incompatible // indirect
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2 v0.0.0-20190505152737-8914e1cf9164 // indirect
	github.com/gavv/monotime v0.0.0-20190418164738-30dba4353424 // indirect
	github.com/go-redis/redis v6.15.2+incompatible
	github.com/go-xorm/core v0.6.2
	github.com/go-xorm/xorm v0.7.3
	github.com/golang/protobuf v1.3.1
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/hashicorp/consul/api v1.1.0
	github.com/imkira/go-interpol v1.1.0 // indirect
	github.com/iris-contrib/blackfriday v2.0.0+incompatible // indirect
	github.com/iris-contrib/formBinder v0.0.0-20190104093907-fbd5963f41e1 // indirect
	github.com/iris-contrib/go.uuid v2.0.0+incompatible // indirect
	github.com/iris-contrib/httpexpect v0.0.0-20180314041918-ebe99fcebbce // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/juju/errors v0.0.0-20190207033735-e65537c515d7
	github.com/juju/testing v0.0.0-20190613124551-e81189438503 // indirect
	github.com/k0kubun/colorstring v0.0.0-20150214042306-9440f1994b88 // indirect
	github.com/kataras/golog v0.0.0-20190624001437-99c81de45f40 // indirect
	github.com/kataras/iris v11.1.1+incompatible
	github.com/kataras/pio v0.0.0-20190103105442-ea782b38602d // indirect
	github.com/klauspost/compress v1.7.1 // indirect
	github.com/klauspost/cpuid v1.2.1 // indirect
	github.com/mattn/go-oci8 v0.0.0-20190524004114-abcafa411cc8
	github.com/microcosm-cc/bluemonday v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/moul/http2curl v1.0.0 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/sergi/go-diff v1.0.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/goconvey v0.0.0-20190330032615-68dc04aab96a // indirect
	github.com/xeipuuv/gojsonpointer v0.0.0-20180127040702-4e3ac2762d5f // indirect
	github.com/xeipuuv/gojsonreference v0.0.0-20180127040603-bd5ef7bd5415 // indirect
	github.com/xeipuuv/gojsonschema v1.1.0 // indirect
	github.com/yalp/jsonpath v0.0.0-20180802001716-5cc68e5049a0 // indirect
	github.com/yudai/gojsondiff v1.0.0 // indirect
	github.com/yudai/golcs v0.0.0-20170316035057-ecda9a501e82 // indirect
	github.com/yudai/pp v2.0.1+incompatible // indirect
	google.golang.org/grpc v1.20.1
	gopkg.in/yaml.v2 v2.2.2
)
