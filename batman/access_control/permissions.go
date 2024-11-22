package accesscontrol

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/fahadPathan7/socialmedia_backend/batman"
	"github.com/fahadPathan7/socialmedia_backend/batman/auth"
	"github.com/fahadPathan7/socialmedia_backend/batman/validate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"gopkg.in/mgo.v2/bson"
)

// Action action def
type Action int

const (
	Create Action = iota
	Read
	Update
	Delete
)

func (d Action) String() string {
	return [...]string{"Create", "Read", "Update", "Delete"}[d]
}

var cs = os.Getenv("MONGO_URI")

const databaseName = "accesscontrol"
const recordPermissionsCollName = "recordPermissions"
const entityPermissionsCollName = "entityPermissions"

// recordPermissions stores permission for different records
type recordPermissions struct {
	ID                 string   `json:"id" bson:"id"`
	Entity             string   `json:"entity" bson:"entity"`
	IDsAllowedToRead   []string `json:"ids_allowed_to_read" bson:"ids_allowed_to_read"`
	IDsAllowedToUpdate []string `json:"ids_allowed_to_update" bson:"ids_allowed_to_update"`
	IDsAllowedToDelete []string `json:"ids_allowed_to_delete" bson:"ids_allowed_to_delete"`
}

// entityPermissions stores permission for different entities
type entityPermissions struct {
	Entity               string   `json:"entity" bson:"entity"`
	RolesAllowedToRead   []string `json:"roles_allowed_to_read" bson:"roles_allowed_to_read"`
	RolesAllowedToCreate []string `json:"roles_allowed_to_create" bson:"roles_allowed_to_create"`
	RolesAllowedToUpdate []string `json:"roles_allowed_to_update" bson:"roles_allowed_to_update"`
	RolesAllowedToDelete []string `json:"roles_allowed_to_delete" bson:"roles_allowed_to_delete"`
}

func fetchRecordPermissions(id, entityName string) (recordPermissions, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getDBURL()))
	if err != nil {
		return recordPermissions{}, err
	}
	defer client.Disconnect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	collection := client.Database(databaseName).Collection(recordPermissionsCollName)
	filter := bson.M{"id": id, "entity": entityName}
	result := collection.FindOne(ctx, filter)

	data := recordPermissions{}
	err = result.Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return recordPermissions{}, nil
		}
		return recordPermissions{}, err
	}

	return data, nil
}

func fetchEntityPermissions(entityName string) (entityPermissions, error) {
	entityPermissionMap := entityRolePermissionsMap()
	ep, ok := entityPermissionMap[entityName]
	if !ok {
		return entityPermissions{}, fmt.Errorf("no entity permission setting not found for %v", entityName)
	}
	return ep, nil
}

func storeRecordPermissions(rp recordPermissions) error {
	err := store(rp, recordPermissionsCollName)
	if err != nil {
		return err
	}
	return nil
}

func storEntityPermissions(ep entityPermissions) error {
	err := store(ep, entityPermissionsCollName)
	if err != nil {
		return err
	}
	return nil
}

func store(data interface{}, collectionName string) error {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(getDBURL()))
	if err != nil {
		return err
	}
	defer client.Disconnect(ctx)

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	collection := client.Database(databaseName).Collection(collectionName)
	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

// CheckPermissionFromCtx checks permission, it extracts the user from context header
func CheckPermissionFromCtx(ctx context.Context, action Action, entityName string, recordID string) (bool, error) {
	authSvc := auth.New()

	userID := "guest"
	roles := []string{Guest}
	u, err := authSvc.UserFromContext(ctx)
	if err == nil {
		userID = u.Id
		roles = append(roles, u.Roles...)
	}

	// Fetch record permissions
	recordPermissions, err := fetchRecordPermissions(recordID, entityName)
	if err != nil {
		fmt.Printf("could not fetch record permissions :%v\n", err)
		return false, err
	}

	// fetch entity level permission
	entityPermissions, err := fetchEntityPermissions(entityName)
	if err != nil {
		fmt.Printf("could not fetch entity permissions :%v\n", err)
		return false, nil
	}

	fmt.Println(roles)
	fmt.Println(entityPermissions)

	switch action {
	case Create:
		for _, role := range roles {
			if stringInSlice(role, entityPermissions.RolesAllowedToCreate) {
				return true, nil
			}
		}
	case Read:
		if stringInSlice(userID, recordPermissions.IDsAllowedToRead) {
			return true, nil
		}

		for _, role := range roles {
			if stringInSlice(role, entityPermissions.RolesAllowedToRead) {
				return true, nil
			}
		}
	case Update:
		if stringInSlice(userID, recordPermissions.IDsAllowedToUpdate) {
			return true, nil
		}

		for _, role := range roles {
			if stringInSlice(role, entityPermissions.RolesAllowedToUpdate) {
				return true, nil
			}
		}
	case Delete:
		if stringInSlice(userID, recordPermissions.IDsAllowedToDelete) {
			return true, nil
		}

		for _, role := range roles {
			if stringInSlice(role, entityPermissions.RolesAllowedToDelete) {
				return true, nil
			}
		}
	}

	return false, nil
}

// GRPCCheckPermissionFromCtx checks permission and returns with proper error
func GRPCCheckPermissionFromCtx(ctx context.Context, action Action, entityName string, recordID string) error {
	ok, err := CheckPermissionFromCtx(ctx, action, entityName, recordID)
	if !ok || err != nil {
		st := batman.ComposeMultipleErrorStr(
			codes.PermissionDenied,
			validate.PERMISSIONDENIED,
			[]string{
				"Permission denied",
			},
		)
		return st.Err()
	}

	return nil
}

// StoreRecordPermissions stores record level permission
func StoreRecordPermissions(id, entityName string, userIDs []string, actions ...Action) error {
	rp := recordPermissions{
		ID:     id,
		Entity: entityName,
	}

	for _, action := range actions {
		switch action {
		case Create:

		case Read:
			rp.IDsAllowedToRead = append(rp.IDsAllowedToRead, userIDs...)
		case Update:
			rp.IDsAllowedToUpdate = append(rp.IDsAllowedToUpdate, userIDs...)
		case Delete:
			rp.IDsAllowedToDelete = append(rp.IDsAllowedToDelete, userIDs...)
		}
	}

	// Store in database
	err := storeRecordPermissions(rp)
	if err != nil {
		return err
	}

	return nil
}

func getDBURL() string {
	return os.Getenv("MONGO_URI")
}
