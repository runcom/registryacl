package main

import (
	dockerapi "github.com/docker/docker/api"
	dockerclient "github.com/docker/engine-api/client"
	"github.com/docker/go-plugins-helpers/authz"
)

func newPlugin(dockerHost string) (*registryacl, error) {
	client, err := dockerclient.NewClient(dockerHost, dockerapi.DefaultVersion.String(), nil, nil)
	if err != nil {
		return nil, err
	}
	return &registryacl{client: client}, nil
}

var (
//startRegExp = regexp.MustCompile(`/containers/(.*)/start$`)
// TODO(runcom): pull/push/search and other action that deal with registries
)

type registryacl struct {
	client *dockerclient.Client
}

func (p *registryacl) AuthZReq(req authz.Request) authz.Response {
	//if req.RequestMethod == "POST" && startRegExp.MatchString(req.RequestURI) {
	//// this is deprecated in docker, remove once hostConfig is dropped to
	//// being available at start time
	//if req.RequestBody != nil {
	//type vfrom struct {
	//VolumesFrom []string
	//}
	//vf := &vfrom{}
	//if err := json.NewDecoder(bytes.NewReader(req.RequestBody)).Decode(vf); err != nil {
	//return authz.Response{Err: err.Error()}
	//}
	//if len(vf.VolumesFrom) > 0 {
	//goto noallow
	//}
	//}
	//res := startRegExp.FindStringSubmatch(req.RequestURI)
	//if len(res) < 1 {
	//return authz.Response{Err: "unable to find container name"}
	//}

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
