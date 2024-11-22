package setting

import accesscontrol "github.com/fahadPathan7/socialmedia_backend/batman/access_control"

const servicePath = "/post.PostService/"

// NotAuthGuardedEndpoints map contains all the endpoints that are not auth guarded.
// By default, all the endpoints are auth guarded
var NotAuthGuardedEndpoints = map[string]bool{
}


// AccessableRoles maps endpoints that are accessable by roles
var AccessableRoles = map[string][]string{
	servicePath + "Create":                       {accesscontrol.Admin, accesscontrol.User},
	servicePath + "Read":                         {accesscontrol.Admin, accesscontrol.User},
	servicePath + "ReadAll":                      {accesscontrol.Admin, accesscontrol.User},
	servicePath + "Update":                       {accesscontrol.Admin, accesscontrol.User},
	servicePath + "Delete":                       {accesscontrol.Admin, accesscontrol.User},
}
