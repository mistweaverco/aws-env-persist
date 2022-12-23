package main

type model struct {
	cursor int
	choice string
}

type AWSEnv struct {
	AccessKey    string
	SecretKey    string
	SessionToken string
	Region       string
}
