package replaceme

// import (
// 	"testing"
// 	"fmt"
// 	"os"
// )


// func Test_OnlyEnv(t *testing.T) {
// 	data := "User is : USER and his home directory is: HOME"
// 	want := fmt.Sprintf("User is : %v and his home directory is: %v", os.Getenv("USER"), os.Getenv("HOME"))
// 	got := ReplaceData([]byte(data), "")

// 	if want != string(got) {
// 		t.Fatalf("want %s, but got %s", want, got)
// 	}
// }