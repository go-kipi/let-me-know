package mongo

import "go.uber.org/fx"

var Module = fx.Module("mongo",
	fx.Provide(
		fx.Annotate(
			NewMongo,
			fx.As(new(MongoI)),
		),
	),
)
