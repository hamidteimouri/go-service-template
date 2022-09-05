package bootstrap

import "laramanpurego/cmd/initialize"

func Start()  {
	//initialize.DatabaseInitialization()
	initialize.DatabaseMigration()
}
