package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/BurntSushi/toml"
	dockerapi "github.com/docker/docker/api"
	dockerclient "github.com/docker/engine-api/client"
	"github.com/docker/go-plugins-helpers/authz"
)

type registryacl struct {
	// we need the client to ask Docker for additional registry (currently only
	// supported in RedHat Docker's fork
	client *dockerclient.Client
	config config
}

type config struct {
	Registries map[string]policy
}

const (
	all = "*"

	actionPull   = "pull"
	actionPush   = "push"
	actionSearch = "search"
	actionLogin  = "login"
)

type policy struct {
	Allow []string
	Deny  []string
}

func newPlugin(dockerHost, configPath string) (*registryacl, error) {
	client, err := dockerclient.NewClient(dockerHost, dockerapi.DefaultVersion.String(), nil, nil)
	if err != nil {
		return nil, err
	}
	if configPath == "" {
		return nil, fmt.Errorf("Please specify a configuration file")
	}
	absConfig, err := filepath.Abs(configPath)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadFile(absConfig)
	if err != nil {
		return nil, err
	}
	var conf config
	if _, err := toml.Decode(string(b), &conf); err != nil {
		return nil, err
	}
	return &registryacl{client: client, config: conf}, nil
}

var (
	// POST
	pushRegExp = regexp.MustCompile(`/images/(.*)/push$`)
	pullRegExp = regexp.MustCompile(`/images/create`)
	authRegExp = regexp.MustCompile(`/auth$`)
	// GET
	searchRegExp = regexp.MustCompile(`/images/search`)
	// TODO(runcom): block build from?!
)

func (p *registryacl) AuthZReq(req authz.Request) authz.Response {
	var (
		image  string
		action string
	)
	switch req.RequestMethod {
	case "GET":
		if searchRegExp.MatchString(req.RequestURI) { // && term != "" -> then try/block search
			action = actionSearch

		}
	case "POST":
		if pushRegExp.MatchString(req.RequestURI) {
			res := pushRegExp.FindStringSubmatch(req.RequestURI)
			if len(res) < 1 {
				return authz.Response{Err: "unable to find image name"}
			}
			image = res[1]
			action = actionPush
			// decide what to allow/block

		}
		if pullRegExp.MatchString(req.RequestURI) { // && &fromImage != "" -> then we're pulling
			action = actionPull

		}
		if authRegExp.MatchString(req.RequestURI) {
			action = actionAuth

		}
	}

	//container, err := p.client.ContainerInspect(res[1])
	//if err != nil {
	//return authz.Response{Err: err.Error()}
	//}
	//image, _, err := p.client.ImageInspectWithRaw(container.Image, false)
	//if err != nil {
	//return authz.Response{Err: err.Error()}
	//}
	//if len(image.Config.Volumes) > 0 {
	//goto noallow
	//}
	//for _, m := range container.Mounts {
	//if m.Driver != "" {
	//goto noallow
	//}
	//}
	//if len(container.HostConfig.VolumesFrom) > 0 {
	//goto noallow
	//}
	//}
	return authz.Response{Allow: true}

	//noallow:
	//return authz.Response{Msg: "volumes are not allowed"}
}

func (p *registryacl) AuthZRes(req authz.Request) authz.Response {
	return authz.Response{Allow: true}
}
