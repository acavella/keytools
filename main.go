package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

func init() {

	viper.SetConfigName("config")  // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./conf/") // optionally look for config in the working directory

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		fmt.Printf("fatal error config file: %s", err.Error())
	}

}

func main() {
	// Define command line flag input
	subject := flag.String("cn", "default", "string, defines request common name.")
	rsaPtr := flag.Bool("rsa", false, "boolean, informs use of rsa cipher.")
	eccPtr := flag.Bool("ecc", false, "boolean, informs use of ecc cipher.")

	err, key := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Printf("test %s", err)
	} else {
		fmt.Printf("test %s", key)
	}

	// Parse command line flags
	flag.Parse()
	keyUsage := viper.GetStringSlice("ku")
	exkeyUsage := viper.GetStringSlice("eku")

	fmt.Println("CN:", *subject)

	for i := 0; i < len(keyUsage); i++ {
		fmt.Printf("KU: %s\n", keyUsage[i])
	}

	for i := 0; i < len(keyUsage); i++ {
		fmt.Printf("EKU: %s\n", exkeyUsage[i])
	}

	if *rsaPtr {
		fmt.Println("Cipher: RSA")
	} else if *eccPtr {
		fmt.Println("Cipher: ECC")
	} else {
		fmt.Println("Cipher spec was not defined.")
	}

}
