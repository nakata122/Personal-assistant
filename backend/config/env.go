package config

import (
	"os"
	"bufio"
	"fmt"
)

func LoadEnv() {
	env, err := os.Open(".env");
	if err != nil {
		fmt.Println(err)
	}
	
	s := bufio.NewScanner(env);
	s.Split(bufio.ScanWords);

	for s.Scan() {
		key := s.Text();
		s.Scan();

		os.Setenv(key, s.Text());
	}

	defer env.Close();
}