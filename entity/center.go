package entity

type CenterUrl struct {
	NodeName        string `json:"nodeName,omitempty"`
	Enable          string `json:"enable"`
	FastTransEnable string `json:"fastTransEnable,omitempty"`
	IsManagerNode   string `json:"isManagerNode,omitempty"`
	UdpPort         string `json:"udpPort,omitempty"`
	Version         string `json:"version"`
	Url             string `json:"url"`
}

type ClusterUserInfo struct {
	UserName string `json:"userName"`
	HomePath string `json:"homePath"`
}

type Center struct {
	Id              int              `json:"id"`
	Name            string           `json:"name"`
	Description     string           `json:"description,omitempty"`
	ClusterUserInfo *ClusterUserInfo `json:"clusterUserInfo"`
	IngressUrls     []*CenterUrl     `json:"ingressUrls,omitempty"`
	EFileUrls       []*CenterUrl     `json:"efileUrls,omitempty"`
	EShellUrls      []*CenterUrl     `json:"eshellUrls,omitempty"`
	HPCUrls         []*CenterUrl     `json:"hpcUrls,omitempty"`
	AIUrls          []*CenterUrl     `json:"aiUrls,omitempty"`
}
