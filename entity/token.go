package entity

type Token struct {
	ClusterName string `json:"clusterName"`
	ClusterId   string `json:"clusterId"`
	Token       string `json:"token"`
}
