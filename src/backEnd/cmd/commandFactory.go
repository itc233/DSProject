package cmd

import (
    "reflect"
    "encoding/json"
    "fmt"
)

type CommandFactory struct {
    commandMap map[string] reflect.Type
}

func (factory *CommandFactory) Init() {
    factory.commandMap = make(map[string]reflect.Type)
    factory.commandMap["RegisterUser"] = reflect.TypeOf(RegisterUserArgs{})
    factory.commandMap["DeleteUser"] = reflect.TypeOf(DeleteUserArgs{})
    factory.commandMap["UserLogin"] = reflect.TypeOf(UserLoginArgs{})
    factory.commandMap["UserLogout"] = reflect.TypeOf(UserLogoutArgs{})
    factory.commandMap["Follow"] = reflect.TypeOf(FollowArgs{})
    factory.commandMap["UnFollow"] = reflect.TypeOf(UnFollowArgs{})
    factory.commandMap["Post"] = reflect.TypeOf(PostArgs{})
    factory.commandMap["GetMyContent"] = reflect.TypeOf(GetMyContentArgs{})
    factory.commandMap["GetFollower"] = reflect.TypeOf(GetFollowerArgs{})
}

type Command struct {
    Name string
    Args string
}

func (factory *CommandFactory) Encode(value reflect.Value) string {
    args, _ := json.Marshal(reflect.Indirect(value.Field(0)).Interface())
    command := Command{
        Name:value.Type().Name(),
        Args:string(args),
    }
    encoded, _ := json.Marshal(command)
    return string(encoded)
}

func (factory *CommandFactory) Decode(encoded string) (string, []reflect.Value) {
    fmt.Printf("Command encoded %s\n", encoded)
    command := Command{}
    err := json.Unmarshal([]byte(encoded), &command)
    fmt.Print(err)
    fmt.Printf("Command %v\n", command)

    cmdArgsType := factory.commandMap[command.Name]
    fmt.Printf("Command Name: %s, Type: %v\n", command.Name, cmdArgsType)
    cmdArgs := reflect.New(cmdArgsType)
    json.Unmarshal([]byte(command.Args), cmdArgs.Interface())
    parameters := make([]reflect.Value, 0)
    for i:=0; i<cmdArgs.Elem().NumField(); i++ {
        parameters = append(parameters, cmdArgs.Elem().Field(i))
    }
    return command.Name, parameters
}


