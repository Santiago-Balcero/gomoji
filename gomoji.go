// Package gomoji provides utilities for converting between different emoji formats.
//
// This package supports conversion between emoji unicode characters, shortcodes,
// HTML entities, and unicode escape sequences. It can handle individual emoji
// transformations as well as bulk text processing.
//
// Example usage:
//
//	import "github.com/Santiago-Balcero/gomoji"
//
//	// Convert emoji name to actual emoji
//	emoji, err := gomoji.Transform("smile", gomoji.FormatEmoji)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(emoji) // Output: 😄
//
//	// Convert emoji to shortcode
//	shortcode, err := gomoji.Transform("😄", gomoji.FormatShortcode)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(shortcode) // Output: :smile:
//
//	// Transform all emojis in text
//	text := "I'm happy 😄 and winking 😉!"
//	result, err := gomoji.TransformText(text, gomoji.FormatShortcode)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(result) // Output: I'm happy :smile: and winking :wink:!
package gomoji

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Format represents the different emoji format types.
type Format string

const (
	// FormatEmoji represents the actual unicode emoji character (🎙️).
	FormatEmoji Format = "emoji"
	// FormatShortcode represents the shortcode format (:microphone:).
	FormatShortcode Format = "shortcode"
	// FormatHTML represents the HTML entity format (&#x1f399;&#xfe0f;).
	FormatHTML Format = "html"
	// FormatUnicode represents the unicode escape sequence format (\U0001F399\uFE0F).
	FormatUnicode Format = "unicode"
)

// Mapping represents all possible formats for a single emoji.
type Mapping struct {
	// Emoji is the actual unicode emoji character.
	Emoji string
	// Shortcode is the textual shortcode representation.
	Shortcode string
	// HTML is the HTML entity representation.
	HTML string
	// Unicode is the unicode escape sequence representation.
	Unicode string
}

// Transform converts between different emoji formats.
//
// The input can be:
//   - An emoji name (e.g., "smile")
//   - An actual emoji character (e.g., "😄")
//   - A shortcode (e.g., ":smile:" or "smile")
//   - An HTML entity (e.g., "&#x1f604;")
//   - A unicode escape sequence (e.g., "\\U0001F604")
//
// The targetFormat specifies the desired output format.
//
// Returns an error if the input emoji is not supported or the target format is invalid.
func Transform(input string, targetFormat Format) (string, error) {
	// Validate target format
	switch targetFormat {
	case FormatEmoji, FormatShortcode, FormatHTML, FormatUnicode:
		// Valid format
	default:
		return "", fmt.Errorf("invalid target format: %s. Valid formats: emoji, shortcode, html, unicode", targetFormat)
	}

	// First, try to identify what format the input is and find the emoji name
	emojiName := findEmojiName(input)
	if emojiName == "" {
		return "", fmt.Errorf("emoji not found or not supported: %s", input)
	}

	// Get the mapping for this emoji
	mapping, exists := emojiMappings[emojiName]
	if !exists {
		return "", fmt.Errorf("emoji mapping not found: %s", emojiName)
	}

	// Return the requested format
	switch targetFormat {
	case FormatEmoji:
		return mapping.Emoji, nil
	case FormatShortcode:
		return mapping.Shortcode, nil
	case FormatHTML:
		return mapping.HTML, nil
	case FormatUnicode:
		return mapping.Unicode, nil
	default:
		return "", fmt.Errorf("unexpected format: %s", targetFormat)
	}
}

// TransformText transforms all emojis found in a text to the target format.
//
// This function can handle mixed emoji formats within the same text and will
// convert all recognized emojis to the specified target format.
//
// Example:
//
//	text := "Hello 😄 :wink: &#x1f44d; world!"
//	result, err := TransformText(text, FormatShortcode)
//	// result: "Hello :smile: :wink: :thumbs_up: world!"
func TransformText(text string, targetFormat Format) string {
	result := text

	// Transform actual emojis
	for emoji, name := range emojiToName {
		if strings.Contains(result, emoji) {
			transformed, err := Transform(name, targetFormat)
			if err != nil {
				log.Printf("transformation for emoji %q with name %q failed: %v", emoji, name, err)
				continue // Skip if transformation fails
			}
			result = strings.ReplaceAll(result, emoji, transformed)
		}
	}

	// Transform shortcodes
	shortcodeRegex := regexp.MustCompile(`:[a-zA-Z_]+:`)
	result = shortcodeRegex.ReplaceAllStringFunc(result, func(match string) string {
		if name, exists := shortcodeToName[match]; exists {
			transformed, err := Transform(name, targetFormat)
			if err != nil {
				log.Printf("transformation for shortcode %q with name %q failed: %v", match, name, err)
				return match // Return original if transformation fails
			}
			return transformed
		}
		return match
	})

	// Transform HTML entities
	htmlRegex := regexp.MustCompile(`&#x[0-9a-fA-F]+;(?:&#x[0-9a-fA-F]+;)*`)
	result = htmlRegex.ReplaceAllStringFunc(result, func(match string) string {
		if name, exists := htmlToName[match]; exists {
			transformed, err := Transform(name, targetFormat)
			if err != nil {
				log.Printf("transformation for HTML %q with name %q failed: %v", match, name, err)
				return match
			}
			return transformed
		}
		return match
	})

	return result
}

// GetSupportedEmojis returns a list of all supported emoji names.
//
// This can be useful for validation or for displaying available emojis to users.
func GetSupportedEmojis() []string {
	var names []string
	for name := range emojiMappings {
		names = append(names, name)
	}
	return names
}

// GetEmojiInfo returns complete information about an emoji.
//
// The input can be in any supported format (name, emoji, shortcode, HTML, unicode).
// Returns a Mapping struct containing all format representations of the emoji.
func GetEmojiInfo(input string) (*Mapping, error) {
	name := findEmojiName(input)
	if name == "" {
		return nil, fmt.Errorf("emoji not found: %s", input)
	}

	mapping := emojiMappings[name]
	return &mapping, nil
}

// IsSupported checks if an emoji is supported by the library.
//
// The input can be in any format (name, emoji, shortcode, HTML, unicode).
func IsSupported(input string) bool {
	return findEmojiName(input) != ""
}

// findEmojiName attempts to identify the emoji name from various input formats.
func findEmojiName(input string) string {
	// Clean input
	input = strings.TrimSpace(input)

	// Check if it's a direct emoji name (like "smile")
	if _, exists := emojiMappings[input]; exists {
		return input
	}

	// Check if it's an actual emoji
	if name, exists := emojiToName[input]; exists {
		return name
	}

	// Check if it's a shortcode
	if name, exists := shortcodeToName[input]; exists {
		return name
	}

	// Check if it's HTML encoded
	if name, exists := htmlToName[input]; exists {
		return name
	}

	// Handle hybrid HTML format: &#x1f399;️ (HTML entity + actual variation selector emoji)
	// Convert trailing variation selector emoji (️) to HTML entity (&#xfe0f;)
	if strings.Contains(input, "&#x") && strings.HasSuffix(input, "️") {
		normalizedInput := strings.Replace(input, "️", "&#xfe0f;", 1)
		if name, exists := htmlToName[normalizedInput]; exists {
			return name
		}
	}

	// Check if it's Unicode escaped
	if name, exists := unicodeToName[input]; exists {
		return name
	}

	// Try to match shortcode without colons
	shortcodeWithColons := fmt.Sprintf(":%s:", input)
	if name, exists := shortcodeToName[shortcodeWithColons]; exists {
		return name
	}

	return ""
}
