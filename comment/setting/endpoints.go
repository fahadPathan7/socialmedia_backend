package setting

import accesscontrol "github.com/fahadPathan7/socialmedia_backend/batman/access_control"

const servicePath = "/comment.CommentService/"

// NotAuthGuardedEndpoints map contains all the endpoints that are not auth guarded.
// By default, all the endpoints are auth guarded
var NotAuthGuardedEndpoints = map[string]bool{
}


// AccessableRoles maps endpoints that are accessable by roles
var AccessableRoles = map[string][]string{
	servicePath + "CreateComment":                       {accesscontrol.Admin, accesscontrol.User},
	servicePath + "ReadAComment":                         {accesscontrol.Admin, accesscontrol.User},
	servicePath + "ReadAllCommentsOfAPost":                      {accesscontrol.Admin, accesscontrol.User},
	servicePath + "UpdateAComment":                       {accesscontrol.Admin, accesscontrol.User},
	servicePath + "DeleteAComment":                       {accesscontrol.Admin, accesscontrol.User},
}
