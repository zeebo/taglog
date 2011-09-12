Taglog implementes tagged based logging for go.

Example:

```go
package main

import (
	"github.com/zeebo/taglog"
	"log"
	"os"
)

var stdlog = log.New(os.Stderr, "", LstdFlags)

func main() {
	logger := taglog.New(stdlog)

	logger.Enable("debug")
	logger.Enable("feature")

	logger.Print("debug, other_tag", "Some statement") //prints
	logger.Print("other_tag", "Won't work!") //no tags
	logger.Print("debug, feature", "Debugging some feature") //prints once

	logger.Disable("debug")
	logger.Print("debug", "I hope you know this wont print!")
	logger.Print("feature", "But this will :)")
}
```

Coming soon is a way to enable/disable tags based on command line arguments in
a standard way.