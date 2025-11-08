# Gomoji 🎉

A powerful and easy-to-use Go package for converting between different emoji formats. Gomoji supports conversion between emoji unicode characters, shortcodes, HTML entities, and unicode escape sequences.

[![Go Reference](https://pkg.go.dev/badge/github.com/Santiago-Balcero/gomoji.svg)](https://pkg.go.dev/github.com/Santiago-Balcero/gomoji)
[![Go Report Card](https://goreportcard.com/badge/github.com/Santiago-Balcero/gomoji)](https://goreportcard.com/report/github.com/Santiago-Balcero/gomoji)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

## Features

- **Multiple Format Support**: Convert between emoji unicode, shortcodes, HTML entities, and unicode escape sequences
- **Individual & Bulk Processing**: Transform single emojis or entire text containing multiple emojis
- **Fast Performance**: Optimized reverse mappings for quick lookups
- **Comprehensive Database**: Support for 215+ commonly used emojis organized by category
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

Gomoji includes support for 215+ commonly used emojis organized across various categories:

### 😊 Faces & Emotions
`smile`, `joy`, `heart_eyes`, `wink`, `blush`, `thinking`, `cry`, `angry`, `scream`, etc.

### 🐶 Animals & Nature
`dog`, `cat`, `bear`, `lion_face`, `frog`, `sunflower`, `rose`, `sun`, `fire`, `moon`, `star2`, `droplet`, `ocean`, `earth_africa`, `earth_americas`, `earth_asia`, `desert_island`, `classical_building`, `seedling`, `evergreen_tree`, `rainbow`, etc.

### 👍 People & Body
`thumbs_up`, `clap`, `wave`, `pray`, `point_up`, `ok_hand`, `peace`, `fist`, `woman_teacher`, `raised_hand`, `vulcan`, etc.

### ❤️ Hearts & Symbols
`heart`, `yellow_heart`, `broken_heart`, `star`, `sparkles`, `zap`, `gem`, `bomb`, etc.

### 🍕 Food & Drink
`pizza`, `hamburger`, `apple`, `coffee`, `beer`, `wine_glass`, `cake`, etc.

### 🚗 Travel & Transportation
`car`, `airplane`, `rocket`, `ship`, `bicycle`, `scooter`, `motorcycle`, `racing_car`, `train`, etc.

### 💻 Objects & Technology
`calendar`, `computer`, `desktop_computer`, `floppy_disk`, `phone`, `camera`, `camera_flash`, `headphones`, `microphone`, `studio_microphone`, `chair`, `eyes`, `guitar`, etc.

### 🇺🇸 Flags (Organized by Continent)
**North America**: `flag_us`
**Europe**: `flag_gb`, `flag_fr`, `flag_it`, `flag_de`, `flag_es`
**Asia**: `flag_jp`, `flag_cn`
**South America**: `flag_co`, `flag_ar`, `flag_mx`, `flag_br`

For a complete list of supported emojis, use:
```go
emojis := gomoji.GetSupportedEmojis()
```

## Flexible Format Support

Gomoji automatically supports **multiple format variations** for emojis with variation selectors, making it extremely flexible for real-world usage:

### Supported Format Variations

For emojis that have variation selectors (like 🎙️, ⌨️, 🖥️), Gomoji recognizes these formats:

```go
// All these inputs work for the studio microphone emoji:
emoji1, _ := gomoji.Transform("&#x1f399;&#xfe0f;", gomoji.FormatEmoji) // Complete HTML: 🎙️
emoji2, _ := gomoji.Transform("&#x1f399;", gomoji.FormatEmoji)         // Base HTML: 🎙️  
emoji3, _ := gomoji.Transform("&#x1f399;️", gomoji.FormatEmoji)        // Hybrid HTML: 🎙️
emoji4, _ := gomoji.Transform("\\U0001F399\\uFE0F", gomoji.FormatEmoji) // Complete Unicode: 🎙️
emoji5, _ := gomoji.Transform("\\U0001F399", gomoji.FormatEmoji)        // Base Unicode: 🎙️

// All return: 🎙️
```

### Why This Matters

- **Web scraping**: Handle emojis copied from different websites with varying formats
- **User input**: Accept emojis pasted from various sources (browsers, documents, etc.)
- **API integrations**: Work seamlessly with different systems that use different emoji encodings
- **Backward compatibility**: All existing code continues to work while gaining new flexibility

### Real-World Example

```go
// These mixed inputs all work correctly:
inputs := []string{
    "&#x2328;",      // Keyboard (base HTML)
    "&#x1f5b1;️",    // Mouse (hybrid HTML) 
    "\\U0001F399",   // Studio mic (base Unicode)
    "🖥️",           // Desktop computer (actual emoji)
}

for _, input := range inputs {
    if gomoji.IsSupported(input) {
        emoji, _ := gomoji.Transform(input, gomoji.FormatEmoji)
        shortcode, _ := gomoji.Transform(input, gomoji.FormatShortcode)
        fmt.Printf("%s -> %s -> %s\n", input, emoji, shortcode)
    }
}
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

Gomoji is designed for high performance with optimized reverse mappings and efficient lookups:

```go
// Benchmark results (215 emojis, MacBook Pro M4 Pro)
BenchmarkTransform-14                   58742662    20.83 ns/op       0 B/op   0 allocs/op
BenchmarkTransformText-14                   3160   372098 ns/op    5683 B/op  72 allocs/op
BenchmarkGetEmojiInfo-14                30463682    37.76 ns/op      64 B/op   1 allocs/op
BenchmarkIsSupported-14                100000000    10.72 ns/op       0 B/op   0 allocs/op
BenchmarkTransformBaseHTML-14           20315640    62.12 ns/op       0 B/op   0 allocs/op
BenchmarkTransformHybridHTML-14         17585305    68.13 ns/op       0 B/op   0 allocs/op
BenchmarkTransformBaseUnicode-14        16274421    75.47 ns/op       0 B/op   0 allocs/op
```

### Understanding the Metrics

**What these numbers mean:**
- **Operations/sec**: Higher is better - shows how many times the function can run per second
- **ns/op**: Nanoseconds per operation - lower is better (1 second = 1 billion nanoseconds)
- **B/op**: Bytes allocated per operation - lower is better for memory efficiency
- **allocs/op**: Memory allocations per operation - zero allocations = no garbage collection overhead

**Performance Analysis:**
- **`Transform()`**: ~21ns per conversion - extremely fast single emoji transformations with zero memory allocations
- **`IsSupported()`**: ~11ns per check - lightning-fast emoji validation, perfect for hot paths
- **`GetEmojiInfo()`**: ~38ns per lookup - fast metadata retrieval with minimal memory usage (64 bytes)
- **`TransformText()`**: ~372μs for text with multiple emojis - efficient bulk processing with reasonable memory usage
- **Base HTML Format**: ~62ns - fast recognition of HTML base entities (e.g., `&#x1f399;`)
- **Hybrid HTML Format**: ~68ns - efficient handling of mixed formats (e.g., `&#x1f399;️`)
- **Base Unicode Format**: ~75ns - quick processing of Unicode base codes (e.g., `\\U0001F399`)

**Why it's fast:**
- Pre-built reverse mapping tables for O(1) lookups
- Zero memory allocations for basic operations
- Optimized string processing for bulk text transformation

Run benchmarks yourself:
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