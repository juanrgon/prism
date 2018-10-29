# Prism

💎 -> ❤️ 💛 💚 💙 💜

_A golang library for colorizing terminal output_

## Installation
```sh
go get github.com/juanrgon/prism
```

## Usage
**Coloring Text**

Prism lets you turn any string into a colored a string with the functions `InRed`, `InGreen` , `InBlue`, `InYellow`, `InCyan`, `InMagenta`, `InBlack`, and `InWhite`.

```go
package main

import (
	"fmt"

	"github.com/juanrgon/prism"
)

func main() {
    r := prism.InRed("some red text")
    g := prism.InGreen("some green text")
    b := prism.InBlue("some blue text")
    fmt.Println(r, g, "and", b)
}
```

**Coloring Background**

Prism lets you change the background of any string with the functions `OnRed`, `OnGreen`, `OnBlue`, `OnYellow`, `OnCyan`, `OnMagenta`, `OnBlack`, and `OnWhite`.

```go
package main

import (
	"fmt"

	"github.com/juanrgon/prism"
)

func main() {
    y := prism.OnYellow("some text on yellow")
    c := prism.OnCyan("some text on cyan")
    m := prism.OnMagenta("some text on magenta")
    fmt.Println(y, c, "and", m)
}
```
**Bold Text**

Prism lets you turn any string into a bold string with the function `InBold`

```go
package main

import (
	"fmt"

	"github.com/juanrgon/prism"
)

func main() {
    b := prism.InBold("Very Important")
    fmt.Println(b)
}
```


**Underlined Text**

Prism lets you underline any string with the function `Underlined`

```go
package main

import (
	"fmt"

	"github.com/juanrgon/prism"
)

func main() {
    u := prism.Underlined("Emphasized text")
    fmt.Println(u)
}
```

**Combining Styles**

The prism functions return a `prism.DecoratedString` type. This type has the above `In*` (e.g. `InBlue`), `On*` (e.g. `OnGreen`), `InBold`, and `Underlined` functions registered. This allows you to apply multiple stylings to a string.

```go
package main

import (
	"fmt"

	"github.com/juanrgon/prism"
)

func main() {
    text := "A bold, black, underlined string on a yellow background"
    complex := prism.InBlack(text).OnYellow().InBold().Underlined()
    fmt.Println(complex)
}
```

For more examples see [github.com/juangrgon/prism-examples](https://github.com/juanrgon/prism-examples)
