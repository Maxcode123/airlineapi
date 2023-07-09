package main

func main() {
	InitConfig()
	createEngine().Run("localhost:" + Conf.APIPort)
}
