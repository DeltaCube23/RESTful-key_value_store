package main

import (
	"fmt"
)

//MIHIR CHINTU SAME HASH VALUE

const (
	prime = 17003
)

func hash(str string) int {
	var res = 1
	for i := 0; i < len(str); i++ {
		res += (res*prime + int(str[i]-'A'))
		res = res % 101
	}
	//fmt.Println(res)
	return res
}

//Push is to create new key
func Push(key, value string) {
	id := hash(key)
	if store[id] == nil {
		store[id] = newLinkedList()
	}
	store[id].add(key, value)
}

//Get is to return key
func Get(key string) {
	id := hash(key)
	if store[id] == nil {
		fmt.Println("No such key value pair exists")
		return
	}
	newNode := store[id].head
	for {
		if newNode == nil {
			break
		} else if newNode.key == key {
			fmt.Println(newNode.data)
			return
		}
		newNode = newNode.next
	}
	fmt.Println("No such key value pair exists")
}

//Put is to update key
func Put(key, value string) {
	id := hash(key)
	if store[id] == nil {
		fmt.Println("No such key value pair exists")
		return
	}
	newNode := store[id].head
	for {
		if newNode == nil {
			break
		} else if newNode.key == key {
			newNode.data = value
			return
		}
		newNode = newNode.next
	}
	fmt.Println("No such key value pair exists")
}

//Delete the key
func Delete(key string) {
	id := hash(key)
	if store[id] == nil {
		fmt.Println("No such key value pair exists")
		return
	}
	store[id].remove(key)
}

/*func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">> ")
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := args[0]

		switch cmd {
		case "PUSH":
			Push(args[1], args[2])
		case "GET":
			Get(args[1])
		case "PUT":
			Put(args[1], args[2])
		case "DELETE":
			Delete(args[1])
		case "EXIT":
			return
		default:
			fmt.Println("Unrecognized Operation", cmd)
		}
	}
}*/
