package dl

import (
	"os"
	"testing"
)

func TestGetFile(t *testing.T) {
	url := "https://gist.githubusercontent.com/phillipsj/07fed8ce06f932c19ab7613d8426d922/raw/13d3fc0ca54d136ad5744fd4448b65dbc87f32dc/random.txt"
	file := "getfile.txt"
	if err := GetFile(url, file); err != nil {
		t.Errorf("Error downloading the file: %v", err)
	}

	if _, err := os.Stat(file); os.IsNotExist(err) {
		t.Errorf("file not downloaded: %v", err)
	}

	os.Remove(file)
}
