package main

func main() {
	goco :=goco.Server(redisURL)
	account := goco.Account(goco.DOAccount{api: "test"})
	instance:= account.CreateInstance(goco.DODroplet{type: "small"})
	ssh:=instance.waitForBoot().login()
	ssh.Install(goco.Go)
	ssh.Install(goco.Iptable(rules()))
	var a int
}
