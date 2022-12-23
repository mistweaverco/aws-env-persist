package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"gopkg.in/ini.v1"
)

func printHelp() {
	fmt.Println("Usage: aws-env-persist [get-account|get-env|save|version|help]")
	fmt.Println("")
	fmt.Println("Source the output of get-env to set your AWS environment variables")
	fmt.Println("Example: source <(aws-env-persist get-env)")
}

func createDirectoryIfNotExists(dir string) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
}

func getAWSCallerIdentity() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalln(err)
	}
	stsClient := sts.NewFromConfig(cfg)
	resp, err := stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Account:", *resp.Account)
	fmt.Println("UserId:", *resp.UserId)
	fmt.Println("ARN:", *resp.Arn)
}

func getAWSEnv() (error, AWSEnv) {
	errorsSlice := []error{}
	access_key := os.Getenv("AWS_ACCESS_KEY_ID")
	if access_key == "" {
		errorsSlice = append(errorsSlice, errors.New("AWS_ACCESS_KEY_ID is not set"))
	}
	secret_key := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secret_key == "" {
		errorsSlice = append(errorsSlice, errors.New("AWS_SECRET_ACCESS_KEY is not set"))
	}
	session_token := os.Getenv("AWS_SESSION_TOKEN")
	if session_token == "" {
		errorsSlice = append(errorsSlice, errors.New("AWS_SESSION_TOKEN is not set"))
	}
	// region is not required
	region := os.Getenv("AWS_DEFAULT_REGION")
	if len(errorsSlice) > 0 {
		errorString := ""
		for _, err := range errorsSlice {
			errorString += err.Error() + "\n"
		}
		return errors.New(errorString), AWSEnv{}
	}
	return nil, AWSEnv{access_key, secret_key, session_token, region}
}

func getAWSConfigPathDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	awsConfigPathDir := home + "/.aws"
	return awsConfigPathDir
}

func getAWSConfigPath() string {
	awsConfigPathDir := getAWSConfigPathDir()
	awsConfigPath := awsConfigPathDir + "/env.ini"
	return awsConfigPath
}

func clearIniFile() {
	awsConfigPath := getAWSConfigPath()
	inifile, err := ini.Load(awsConfigPath)
	if err != nil {
		inifile = ini.Empty()
	}
	inifile.Section("default").DeleteKey("aws_access_key_id")
	inifile.Section("default").DeleteKey("aws_secret_access_key")
	inifile.Section("default").DeleteKey("aws_session_token")
	inifile.SaveTo(awsConfigPath)
}

func writeIniFile(awsEnv AWSEnv) {
	awsConfigPath := getAWSConfigPath()
	inifile, err := ini.Load(awsConfigPath)
	if err != nil {
		inifile = ini.Empty()
	}
	// because region is not required we need to check if it is set
	if awsEnv.Region != "" {
		inifile.Section("default").Key("region").SetValue(awsEnv.Region)
	}
	inifile.Section("default").Key("aws_access_key_id").SetValue(awsEnv.AccessKey)
	inifile.Section("default").Key("aws_secret_access_key").SetValue(awsEnv.SecretKey)
	inifile.Section("default").Key("aws_session_token").SetValue(awsEnv.SessionToken)
	inifile.SaveTo(awsConfigPath)
}

func getAWSEnvFromIniFile() AWSEnv {
	awsConfigPath := getAWSConfigPath()
	cfg, err := ini.Load(awsConfigPath)
	if err != nil {
		os.Exit(1)
	}
	awsEnv := AWSEnv{}
	awsEnv.Region = cfg.Section("default").Key("region").String()
	awsEnv.AccessKey = cfg.Section("default").Key("aws_access_key_id").String()
	awsEnv.SecretKey = cfg.Section("default").Key("aws_secret_access_key").String()
	awsEnv.SessionToken = cfg.Section("default").Key("aws_session_token").String()
	return awsEnv
}

func outputEnvironmentExports() {
	awsEnv := getAWSEnvFromIniFile()
	if awsEnv.Region != "" {
		fmt.Println("export AWS_DEFAULT_REGION=" + awsEnv.Region)
	}
	if awsEnv.AccessKey != "" {
		fmt.Println("export AWS_ACCESS_KEY_ID=" + awsEnv.AccessKey)
	}
	if awsEnv.SecretKey != "" {
		fmt.Println("export AWS_SECRET_ACCESS_KEY=" + awsEnv.SecretKey)
	}
	if awsEnv.SessionToken != "" {
		fmt.Println("export AWS_SESSION_TOKEN=" + awsEnv.SessionToken)
	}
}
