package component

type AuthInfo struct {
	ClusterName      string
	IdentityEndpoint string
	Username         string
	Password         string
	DomainName       string
	ProjectName      string
}

var ClusterInfo = []AuthInfo{
	{"SJ", "http://1.1.1.1:5000/v3", "admin", "123456", "default", "admin"},
}
