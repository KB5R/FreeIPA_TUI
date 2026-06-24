package backend

import (
	"crypto/tls"
	"net/http"

	"github.com/infra-monkey/go-freeipa/freeipa"
)

type IPAConfig struct {
	Host     string
	Username string
	Password string
	Insecure bool
}

type IPAClient struct {
	api *freeipa.Client
}

type IPAUser struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
}

type IPAGroup struct {
	Name string
}

func NewIPAClient(config IPAConfig) (*IPAClient, error) {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.Insecure,
		},
	}

	api, err := freeipa.Connect(
		config.Host,
		transport,
		config.Username,
		config.Password,
	)
	if err != nil {
		return nil, err
	}

	return &IPAClient{
		api: api,
	}, nil
}

func (client *IPAClient) FindUsers(search string) ([]IPAUser, error) {
	result, err := client.api.UserFind(
		search,
		&freeipa.UserFindArgs{},
		nil,
	)
	if err != nil {
		return nil, err
	}

	users := []IPAUser{}
	for _, sourceUser := range result.Result {
		user := IPAUser{}
		user.Username = sourceUser.UID
		user.LastName = sourceUser.Sn
		if sourceUser.Givenname != nil {
			user.FirstName = *sourceUser.Givenname
		}

		users = append(users, user)
	}

	return users, nil
}

func (client *IPAClient) FindGroups(search string) ([]IPAGroup, error) {
	result, err := client.api.GroupFind(
		search,
		&freeipa.GroupFindArgs{},
		nil,
	)
	if err != nil {
		return nil, err
	}

	groups := []IPAGroup{}
	for _, sourceGroup := range result.Result {
		group := IPAGroup{}
		group.Name = sourceGroup.Cn

		groups = append(groups, group)
	}

	return groups, nil
}
