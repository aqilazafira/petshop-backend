package config

var allowedOrigins = []string{
	"http://localhost:3000/",
	"https://Fadhail.github.io",
	"https://petshop.xeroon.xyz",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
