package mongo

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var Client, err = mongo.Connect(
	context.Background(),
	"mongodb+srv://root:lL5f7Rmrk2AOIv6L@cluster0-6mvgl.gcp.mongodb.net/tasks?retryWrites=true",
	nil
)