package cdsclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/ovh/cds/sdk"
)

func (c *client) ApplicationCreate(key string, app *sdk.Application) error {
	_, err := c.PostJSON("/project/"+key+"/applications", app, nil)
	return err
}

func (c *client) ApplicationDelete(key string, appName string) error {
	_, err := c.DeleteJSON("/project/"+key+"/application/"+appName, nil)
	return err
}

func (c *client) ApplicationGet(key string, appName string, mods ...RequestModifier) (*sdk.Application, error) {
	app := &sdk.Application{}
	if _, err := c.GetJSON("/project/"+key+"/application/"+appName, app, mods...); err != nil {
		return nil, err
	}
	return app, nil
}

func (c *client) ApplicationList(key string) ([]sdk.Application, error) {
	apps := []sdk.Application{}
	if _, err := c.GetJSON("/project/"+key+"/applications", &apps); err != nil {
		return nil, err
	}
	return apps, nil
}

func (c *client) ApplicationGroupsImport(projectKey, appName string, content io.Reader, format string, force bool) (sdk.Application, error) {
	var app sdk.Application
	uri := fmt.Sprintf("/project/%s/application/%s/group/import?format=%s", projectKey, appName, format)

	if force {
		uri += "&forceUpdate=true"
	}

	btes, _, _, errReq := c.Request("POST", uri, content)
	if errReq != nil {
		return app, errReq
	}

	if err := json.Unmarshal(btes, &app); err != nil {
		return app, errReq
	}

	return app, errReq
}

//ApplicationAttachToReposistoriesManager attachs the application to the repo identified by its fullname in the reposManager
func (c *client) ApplicationAttachToReposistoriesManager(projectKey, appName, reposManager, repoFullname string) error {
	uri := fmt.Sprintf("/project/%s/repositories_manager/%s/application/%s/attach?fullname=%s", projectKey, reposManager, appName, url.QueryEscape(repoFullname))
	_, _, _, err := c.Request("POST", uri, nil)
	return err
}
