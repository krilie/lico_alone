//+build wireinject

package ndb

func InitMission(name string) (Mission, error) {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}, nil
}
