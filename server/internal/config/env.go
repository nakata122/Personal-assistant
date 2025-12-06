package config

import (
	"os"
	"bufio"
	"log"
)

func LoadEnv() {
	env, err := os.Open(".env");
	if err != nil {
		log.Println(err);
		return;
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