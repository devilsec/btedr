// Get info about the running machine
package tasks

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/user"
	"strconv"
	"strings"

	"github.com/devilsec/btedr/proto/agentpb"
	"github.com/nogproject/nog/backend/pkg/getent"
)

func Info() *agentpb.Registration {
	return &agentpb.Registration{
		Id:       "",
		Os:       osType(),
		Ip:       ip(),
		Hostname: hostname(),
		User:     currentUser(),
		Groups:   groups(),
		Users:    users(),
	}
}

// Get the distro name from /etc/os-release, ie "Ubuntu 22.04.1 LTS"
func osType() string {
	readFile, err := os.Open("/etc/os-release")

	//Read OS Name and Version in the file
	if err == nil {
		scanner := bufio.NewScanner(readFile)
		scanner.Split(bufio.ScanLines)
		var Name string
		var Version string
		for scanner.Scan() {

			line := scanner.Text()
			vals := strings.Split(line, "=")
			if len(vals) < 2 {
				continue
			}

			if vals[0] == "NAME" {
				Name = vals[1]
			} else if vals[0] == "VERSION" {
				Version = vals[1]
			}

		}
		readFile.Close()
		osType := Name + Version
		println(osType)
		return osType
	} else {
		println("Error while getting os type " + err.Error())
	}

	return err.Error()
}

// Get the IP address of this agent that connects back to the server
func ip() string {
	//Based on the connection to google, ip is taken. net.Dial doesn't create a connection
	//but it gets the interface through which the connection was attempted.
	conn, err := net.Dial("udp", "8.8.8.8:80")

	if err != nil {
		println("Error while getting IP address" + err.Error())
		return err.Error()
	}

	IPAddress := conn.LocalAddr().(*net.UDPAddr).IP.String()
	println("IP address " + IPAddress)
	return IPAddress
}

// Get the hostname
func hostname() string {
	hostname, err := os.Hostname()
	if err == nil {
		println("Hostname:" + hostname)
		return hostname
	}
	println("Error obtained" + err.Error())
	return err.Error()
}

// Get the name of the user running this agent
func currentUser() *agentpb.User {
	currUser, err := user.Current()
	//println("current User: " + currUser.Username)
	if err != nil {
		log.Println("Error to get current User: " + err.Error())
		errU := &agentpb.User{
			Id:   0,
			Name: err.Error(),
		}
		return errU
	}

	id, _ := strconv.ParseInt(currUser.Uid, 10, 32)

	u := &agentpb.User{
		Id:   uint32(id),
		Name: currUser.Username,
	}

	return u
}

// Get the groups this user is a part of
func groups() []*agentpb.User {
	groupIds, _ := os.Getgroups()
	groups := []*agentpb.User{}

	for _, id := range groupIds {
		group, err := user.LookupGroupId(fmt.Sprint(id))

		if err == nil {
			u := &agentpb.User{
				Id:   uint32(id),
				Name: group.Name,
			}
			groups = append(groups, u)
		} else {
			println(err.Error())
		}
	}

	return groups
}

func UsersInpasswd() []string {
	readpwdFile, err := os.Open("/etc/passwd")
	var usersList []string
	if err == nil {
		scanner := bufio.NewScanner(readpwdFile)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {

			line := scanner.Text()

			if strings.HasSuffix(line, "sh") {
				vals := strings.Split(line, ":")
				if len(vals) >= 2 {
					usersList = append(usersList, vals[0])
				}
			}
		}
	} else {
		fmt.Printf(err.Error())
	}
	return usersList
}

func homeUsers() []string {
	homeDirSubs, err := os.ReadDir("/home")
	var usersList []string
	if err == nil {
		for _, f := range homeDirSubs {
			if f.IsDir() {
				usersList = append(usersList, f.Name())
			}
		}
	} else {
		fmt.Printf(err.Error())
	}
	return usersList
}

// Get all the users running on this machine
func users() []*agentpb.User {
	ctx := context.Background()

	users := []*agentpb.User{}
	usersInpasswd := UsersInpasswd()
	homeUsersList := homeUsers()
	checkHomeUsers := make(map[string]int)

	for _, val := range homeUsersList {
		checkHomeUsers[val] = 1
	}

	// get id's and append to normal users
	for _, username := range usersInpasswd {
		if username == "" {
			continue
		}
		sysUser, err2 := user.Lookup(username)

		if err2 != nil {
			println("Error obtained: " + err2.Error())
			continue
		}
		uid, _ := strconv.ParseInt(sysUser.Uid, 10, 32)

		u := &agentpb.User{
			Id:   uint32(uid),
			Name: sysUser.Username,
		}
		users = append(users, u)
	}
	isIdObtainedUsers := make(map[string]int)
	for _, val := range users {
		isIdObtainedUsers[val.Name] = 1
	}

	//get id's the remaining AD/Samba users
	allUsers, _ := getent.Passwds(ctx)
	for _, user := range allUsers {
		if _, ok := isIdObtainedUsers[user.User]; ok {
			continue
		}
		// If the user is in home directory and id is not obtained, get id using getent.
		if _, ok := checkHomeUsers[user.User]; ok {
			u := &agentpb.User{
				Id:   uint32(user.Uid),
				Name: user.User + "(AD/Samba)",
			}
			users = append(users, u)
		}
	}

	return users
}
