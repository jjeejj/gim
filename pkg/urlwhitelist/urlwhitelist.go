package urlwhitelist

var Business = map[string]int{
	"/pb.BusinessExt/SignIn":     0,
	"/pb.BusinessExt/UpdateUser": 0,
	"/pb.BusinessExt/GetUser":    0,
}

var Logic = map[string]int{
	"/pb.LogicExt/RegisterDevice": 0,
	"/pb.LogicExt/GetFriends":     0,
	"/pb.LogicExt/GetFriend":      0,
	"/pb.LogicExt/AddFriend":      0,
}
