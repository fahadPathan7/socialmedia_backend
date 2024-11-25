package setting

import accesscontrol "github.com/fahadPathan7/socialmedia_backend/batman/access_control"

const servicePath = "/react.ReactService/"

// NotAuthGuardedEndpoints map contains all the endpoints that are not auth guarded.
// By default, all the endpoints are auth guarded
var NotAuthGuardedEndpoints = map[string]bool{
}


// AccessableRoles maps endpoints that are accessable by roles
var AccessableRoles = map[string][]string{
	servicePath + "CreateAReact":                       {accesscontrol.Admin, accesscontrol.User},
	servicePath + "ReadAReact":                         {accesscontrol.Admin, accesscontrol.User},
	servicePath + "ReadAllReactsOfAPost":                      {accesscontrol.Admin, accesscontrol.User},
	servicePath + "UpdateAReact":                       {accesscontrol.Admin, accesscontrol.User},
	servicePath + "DeleteAReact":                       {accesscontrol.Admin, accesscontrol.User},
	servicePath + "DeleteAllReactsOfAPost":                       {accesscontrol.Admin, accesscontrol.User},
}
