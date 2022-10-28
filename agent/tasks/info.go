// Get info about the running machine
package tasks

import "github.com/devilsec/btedr/proto/agentpb"

func Info() *agentpb.Registration {
  return &agentpb.Registration{
      Id: "",
      Os: os(),
      Ip: ip(),
      Hostname: hostname(),
      User: user(),
      Groups: groups(),
      Users: users(),
    }
}

// Get the distro name from /etc/os-release, ie "Ubuntu 22.04.1 LTS"
func os() string {
	return "TODO"
}

// Get the IP address of this agent that connects back to the server
func ip() string {
	return "TODO"
}

// Get the hostname
func hostname() string {
	return "TODO"
}

// Get the name of the user running this agent
func user() *agentpb.User {
	u := &agentpb.User{
		Id:   "TODO",
		Name: "TODO",
	}
	return u
}

// Get the groups this user is a part of
func groups() []*agentpb.User {
	groups := []*agentpb.User{}
	u := &agentpb.User{
		Id:   "TODO",
		Name: "TODO",
	}
	groups = append(groups, u)
	return groups
}

// Get all the users running on this machine
func users() []*agentpb.User {
	users := []*agentpb.User{}
	u := &agentpb.User{
		Id:   "TODO",
		Name: "TODO",
	}
	users = append(users, u)
	return users
}
