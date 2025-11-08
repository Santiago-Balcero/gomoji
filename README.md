# Gomoji 🎉

A powerful and easy-to-use Go package for converting between different emoji formats. Gomoji supports conversion between emoji unicode characters, shortcodes, HTML entities, and unicode escape sequences.

[![Go Reference](https://pkg.go.dev/badge/github.com/Santiago-Balcero/gomoji.svg)](https://pkg.go.dev/github.com/Santiago-Balcero/gomoji)
[![Go Report Card](https://goreportcard.com/badge/github.com/Santiago-Balcero/gomoji)](https://goreportcard.com/report/github.com/Santiago-Balcero/gomoji)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- **Multiple Format Support**: Convert between emoji unicode, shortcodes, HTML entities, and unicode escape sequences
- **Individual & Bulk Processing**: Transform single emojis or entire text containing multiple emojis
- **Fast Performance**: Optimized reverse mappings for quick lookups
- **Comprehensive Database**: Support for hundreds of commonly used emojis
- **Type Safety**: Strongly typed API with custom format types
- **Zero Dependencies**: Pure Go implementation with no external dependencies
- **Full Test Coverage**: Extensively tested with benchmarks

## Installation

```bash
go get github.com/Santiago-Balcero/gomoji
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/Santiago-Balcero/gomoji"
)

func main() {
    // Convert emoji name to actual emoji
    emoji, err := gomoji.Transform("smile", gomoji.FormatEmoji)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(emoji) // Output: 😄

    // Convert emoji to shortcode
    shortcode, err := gomoji.Transform("😄", gomoji.FormatShortcode)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(shortcode) // Output: :smile:

    // Transform all emojis in text
    text := "I'm happy 😄 and love coding! 💻"
    result, err := gomoji.TransformText(text, gomoji.FormatShortcode)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result) // Output: I'm happy :smile: and love coding! :computer:
}
```

## API Reference

### Formats

Gomoji supports four emoji formats:

```go
const (
    FormatEmoji     Format = "emoji"     // 😄
    FormatShortcode Format = "shortcode" // :smile:
    FormatHTML      Format = "html"      // &#x1f604;
    FormatUnicode   Format = "unicode"   // \\U0001F604
)
```

### Core Functions

#### `Transform(input string, targetFormat Format) (string, error)`

Converts a single emoji between different formats.

**Input formats accepted:**
- Emoji name: `"smile"`
- Actual emoji: `"😄"`
- Shortcode: `":smile:"` or `"smile"`
- HTML entity: `"&#x1f604;"`
- Unicode escape: `"\\U0001F604"`

```go
// Convert from various formats to emoji
emoji, _ := gomoji.Transform("smile", gomoji.FormatEmoji)           // 😄
emoji, _ := gomoji.Transform(":heart:", gomoji.FormatEmoji)         // ❤️
emoji, _ := gomoji.Transform("&#x1f604;", gomoji.FormatEmoji)       // 😄

// Convert to different target formats
shortcode, _ := gomoji.Transform("😄", gomoji.FormatShortcode)      // :smile:
html, _ := gomoji.Transform("smile", gomoji.FormatHTML)             // &#x1f604;
unicode, _ := gomoji.Transform("😄", gomoji.FormatUnicode)          // \\U0001F604
```

#### `TransformText(text string, targetFormat Format) (string, error)`

Transforms all emojis found in a text to the specified format. Handles mixed emoji formats within the same text.

```go
// Convert mixed emoji formats in text
text := "Hello 😄 :heart: &#x1f44d; world!"
result, _ := gomoji.TransformText(text, gomoji.FormatShortcode)
// Output: "Hello :smile: :heart: :thumbs_up: world!"

// Convert all emojis to HTML entities
text = "Great work! 👍 🎉"
html, _ := gomoji.TransformText(text, gomoji.FormatHTML)
// Output: "Great work! &#x1f44d; &#x1f389;"
```

#### `GetEmojiInfo(input string) (*Mapping, error)`

Returns complete information about an emoji in all supported formats.

```go
info, err := gomoji.GetEmojiInfo("smile")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Emoji: %s\n", info.Emoji)         // 😄
fmt.Printf("Shortcode: %s\n", info.Shortcode) // :smile:
fmt.Printf("HTML: %s\n", info.HTML)           // &#x1f604;
fmt.Printf("Unicode: %s\n", info.Unicode)     // \\U0001F604
```

#### `GetSupportedEmojis() []string`

Returns a list of all supported emoji names.

```go
emojis := gomoji.GetSupportedEmojis()
fmt.Printf("Supported emojis: %d\n", len(emojis))
// Output: Supported emojis: 200+

// Check first few emojis
for i, emoji := range emojis[:5] {
    fmt.Printf("%d. %s\n", i+1, emoji)
}
```

#### `IsSupported(input string) bool`

Checks if an emoji is supported by the library.

```go
// Check various formats
fmt.Println(gomoji.IsSupported("smile"))        // true
fmt.Println(gomoji.IsSupported("😄"))           // true
fmt.Println(gomoji.IsSupported(":smile:"))      // true
fmt.Println(gomoji.IsSupported("&#x1f604;"))    // true
fmt.Println(gomoji.IsSupported("invalid"))      // false
```

## Supported Emojis

Gomoji includes support for 200+ commonly used emojis across various categories:

### 😊 Faces & Emotions
`smile`, `joy`, `heart_eyes`, `wink`, `blush`, `thinking`, `cry`, `angry`, `scream`, etc.

### 🐶 Animals & Nature
`dog`, `cat`, `bear`, `lion_face`, `frog`, `sunflower`, `rose`, `sun`, `rainbow`, etc.

### 👍 People & Body
`thumbs_up`, `clap`, `wave`, `pray`, `point_up`, `ok_hand`, `peace`, `fist`, etc.

### ❤️ Hearts & Symbols
`heart`, `yellow_heart`, `broken_heart`, `star`, `fire`, `sparkles`, `zap`, `gem`, etc.

### 🍕 Food & Drink
`pizza`, `hamburger`, `apple`, `coffee`, `beer`, `wine_glass`, `cake`, etc.

### 🚗 Travel & Transportation
`car`, `airplane`, `rocket`, `ship`, `bike`, `train`, etc.

### 💻 Objects & Technology
`computer`, `phone`, `camera`, `headphones`, `microphone`, `guitar`, etc.

### 🇺🇸 Flags
`flag_us`, `flag_gb`, `flag_fr`, `flag_de`, `flag_jp`, `flag_cn`, etc.

For a complete list of supported emojis, use:
```go
emojis := gomoji.GetSupportedEmojis()
```

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/Santiago-Balcero/gomoji"
)

func main() {
    // Simple emoji conversion
    emoji, _ := gomoji.Transform("heart", gomoji.FormatEmoji)
    fmt.Println("Heart emoji:", emoji) // ❤️
    
    // Get complete emoji information
    info, _ := gomoji.GetEmojiInfo("rocket")
    fmt.Printf("Rocket - Emoji: %s, HTML: %s\n", info.Emoji, info.HTML)
}
```

### Text Processing

```go
package main

import (
    "fmt"
    "github.com/Santiago-Balcero/gomoji"
)

func main() {
    // Process user messages
    message := "Great job! 👍 Keep it up! 🔥"
    
    // Convert to shortcodes for storage
    stored, _ := gomoji.TransformText(message, gomoji.FormatShortcode)
    fmt.Println("Stored:", stored) // "Great job! :thumbs_up: Keep it up! :fire:"
    
    // Convert back to emojis for display
    display, _ := gomoji.TransformText(stored, gomoji.FormatEmoji)
    fmt.Println("Display:", display) // "Great job! 👍 Keep it up! 🔥"
}
```

### Web Development

```go
package main

import (
    "fmt"
    "github.com/Santiago-Balcero/gomoji"
)

func main() {
    // Convert emojis for HTML display
    content := "Welcome! 😊 Enjoy your stay! ⭐"
    
    // Convert to HTML entities for safe HTML rendering
    htmlSafe, _ := gomoji.TransformText(content, gomoji.FormatHTML)
    fmt.Println("HTML:", htmlSafe)
    // Output: "Welcome! &#x1f60a; Enjoy your stay! &#x2b50;"
    
    // Convert to shortcodes for Markdown
    markdown, _ := gomoji.TransformText(content, gomoji.FormatShortcode)
    fmt.Println("Markdown:", markdown)
    // Output: "Welcome! :blush: Enjoy your stay! :star:"
}
```

### Validation and Safety

```go
package main

import (
    "fmt"
    "github.com/Santiago-Balcero/gomoji"
)

func main() {
    inputs := []string{"smile", "invalid_emoji", "heart", "❤️"}
    
    for _, input := range inputs {
        if gomoji.IsSupported(input) {
            emoji, _ := gomoji.Transform(input, gomoji.FormatEmoji)
            fmt.Printf("✓ %s -> %s\n", input, emoji)
        } else {
            fmt.Printf("✗ %s is not supported\n", input)
        }
    }
}
```

## Performance

Gomoji is designed for high performance with optimized reverse mappings:

```go
// Benchmark results (example)
BenchmarkTransform-8           5000000    250 ns/op
BenchmarkTransformText-8       1000000   1500 ns/op
BenchmarkIsSupported-8        10000000    120 ns/op
```

Run benchmarks:
```bash
go test -bench=.
```

## Error Handling

Gomoji returns descriptive errors for various scenarios:

```go
// Unsupported emoji
_, err := gomoji.Transform("nonexistent", gomoji.FormatEmoji)
fmt.Println(err) // "emoji not found or not supported: nonexistent"

// Invalid format
_, err = gomoji.Transform("smile", gomoji.Format("invalid"))
fmt.Println(err) // "invalid target format: invalid. Valid formats: emoji, shortcode, html, unicode"

// Empty input
_, err = gomoji.Transform("", gomoji.FormatEmoji)
fmt.Println(err) // "emoji not found or not supported: "
```

## Testing

Run the test suite:

```bash
# Run all tests
go test

# Run tests with coverage
go test -cover

# Run benchmarks
go test -bench=.

# Verbose output
go test -v
```

## Contributing

Contributions are welcome! Please feel free to:

1. Report bugs
2. Suggest new features
3. Submit pull requests
4. Add more emoji support
5. Improve documentation

### Adding New Emojis

To add new emojis to the database:

1. Add the emoji mapping to `data.go`
2. Ensure all formats are included (emoji, shortcode, HTML, unicode)
3. Add tests for the new emoji
4. Update documentation

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

### v1.0.0
- Initial release
- Support for 200+ emojis
- Four format types (emoji, shortcode, HTML, unicode)
- Comprehensive API with Transform, TransformText, GetEmojiInfo, etc.
- Full test coverage
- Performance optimizations

---

Made with ❤️ by [Santiago Balcero](https://github.com/Santiago-Balcero)

**Star ⭐ this repository if you find it useful!**