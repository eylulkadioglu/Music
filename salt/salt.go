package salt

var projectSalt string

func SetSalt(salt string) {
	projectSalt = salt
}

func GetSalt() string {
	return projectSalt
}
