package setting

import accesscontrol "github.com/fahadPathan7/socialmedia_backend/batman/access_control"

const servicePath = "/user.UserService/"

// NotAuthGuardedEndpoints map contains all the endpoints that are not auth guarded.
// By default, all the endpoints are auth guarded
var NotAuthGuardedEndpoints = map[string]bool{
	servicePath + "Login":                   true,
	servicePath + "Register":                true,
}


// AccessableRoles maps endpoints that are accessable by roles
var AccessableRoles = map[string][]string{
	servicePath + "Register":                      {accesscontrol.Admin, accesscontrol.User, accesscontrol.Guest},
	servicePath + "Login":                         {accesscontrol.Admin, accesscontrol.User, accesscontrol.Guest},
	servicePath + "GetUserById":                   {accesscontrol.Admin, accesscontrol.User},
	servicePath + "GetUserByEmail":                {accesscontrol.Admin, accesscontrol.User},
}
