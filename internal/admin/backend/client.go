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

type IPAUser struct {
	Username  string
	FirstName string
	LastName  string
	Email     string
}

type IPAClient struct {
	api *freeipa.Client
}

type IPAGroup struct {
	Name string
}

type IPAHost struct {
	FQDN string
}

type IPAHostGroup struct {
	Name string
}

type IPAHBACRule struct {
	Name    string
	Enabled bool
}

type IPASudoRule struct {
	Name    string
	Enabled bool
}

// * означает что мы работает с одним клиетом а не перезаписываем все
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
	// Отправляем в FreeIPA запрос поиска пользователей.
	result, err := client.api.UserFind(
		search,
		&freeipa.UserFindArgs{},
		nil,
	)

	// Если FreeIPA вернул ошибку, передаём её вызывающему коду.
	if err != nil {
		return nil, err
	}

	// Создаём пустой список для наших пользователей.
	users := []IPAUser{}

	// Wикл обработки найденных пользователей.
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

func (client *IPAClient) FindHosts(search string) ([]IPAHost, error) {
	result, err := client.api.HostFind(
		search,
		&freeipa.HostFindArgs{},
		nil,
	)

	if err != nil {
		return nil, err
	}

	hosts := []IPAHost{}

	for _, sourceHost := range result.Result {
		host := IPAHost{}
		host.FQDN = sourceHost.Fqdn

		hosts = append(hosts, host)
	}

	return hosts, nil
}

func (client *IPAClient) FindHostGroups(search string) ([]IPAHostGroup, error) {
	result, err := client.api.HostgroupFind(
		search,
		&freeipa.HostgroupFindArgs{},
		nil,
	)
	if err != nil {
		return nil, err
	}

	hostGroups := []IPAHostGroup{}
	for _, sourceHostGroup := range result.Result {
		hostGroup := IPAHostGroup{}
		hostGroup.Name = sourceHostGroup.Cn

		hostGroups = append(hostGroups, hostGroup)
	}

	return hostGroups, nil
}

func (client *IPAClient) FindHBACRules(search string) ([]IPAHBACRule, error) {
	result, err := client.api.HbacruleFind(
		search,
		&freeipa.HbacruleFindArgs{},
		nil,
	)
	if err != nil {
		return nil, err
	}

	rules := []IPAHBACRule{}
	for _, sourceRule := range result.Result {
		rule := IPAHBACRule{}
		rule.Name = sourceRule.Cn
		if sourceRule.Ipaenabledflag != nil {
			rule.Enabled = *sourceRule.Ipaenabledflag
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

func (client *IPAClient) FindSudoRules(search string) ([]IPASudoRule, error) {
	result, err := client.api.SudoruleFind(
		search,
		&freeipa.SudoruleFindArgs{},
		nil,
	)
	if err != nil {
		return nil, err
	}

	rules := []IPASudoRule{}
	for _, sourceRule := range result.Result {
		rule := IPASudoRule{}
		rule.Name = sourceRule.Cn
		if sourceRule.Ipaenabledflag != nil {
			rule.Enabled = *sourceRule.Ipaenabledflag
		}

		rules = append(rules, rule)
	}

	return rules, nil
}
