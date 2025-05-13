func interpret(command string) string {
    str := ""
    b := []byte(command)
    for i := 0; i < len(b); i++{
        if i != 0 && string(command[i-1]) == "(" && string(command[i]) == ")" {
            str += "o"
        }else if string(command[i]) == "(" || string(command[i]) == ")" {
            continue
        }else{
            str += string(command[i])
        }
    }

    return str
}