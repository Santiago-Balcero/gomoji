package main

import (
	"fmt"
	"log"

	"github.com/Santiago-Balcero/gomoji"
)

func main() {
	fmt.Println("🎉 Gomoji Examples 🎉")
	fmt.Println("=====================")

	// Example 1: Basic emoji transformations
	fmt.Println("\n1. Basic Emoji Transformations:")
	basicTransformations()

	// Example 2: Text processing
	fmt.Println("\n2. Text Processing:")
	textProcessing()

	// Example 3: Getting emoji information
	fmt.Println("\n3. Emoji Information:")
	emojiInformation()

	// Example 4: Validation
	fmt.Println("\n4. Emoji Validation:")
	emojiValidation()

	// Example 5: Web development scenarios
	fmt.Println("\n5. Web Development Use Cases:")
	webDevelopment()

	// Example 6: Format conversions
	fmt.Println("\n6. Format Conversions:")
	formatConversions()
}

func basicTransformations() {
	// Convert emoji name to actual emoji
	emoji, err := gomoji.Transform("smile", gomoji.FormatEmoji)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   smile -> %s\n", emoji)

	// Convert emoji to shortcode
	shortcode, err := gomoji.Transform("😄", gomoji.FormatShortcode)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   😄 -> %s\n", shortcode)

	// Convert to HTML entity
	html, err := gomoji.Transform("heart", gomoji.FormatHTML)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   heart -> %s\n", html)

	// Convert to Unicode escape
	unicode, err := gomoji.Transform("🔥", gomoji.FormatUnicode)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   🔥 -> %s\n", unicode)
}

func textProcessing() {
	// Transform entire text with mixed emoji formats
	text := "Hello 😊 I'm :heart: coding! &#x1f525;"

	// Convert everything to shortcodes
	shortcodes, err := gomoji.TransformText(text, gomoji.FormatShortcode)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   Original: %s\n", text)
	fmt.Printf("   Shortcodes: %s\n", shortcodes)

	// Convert everything to emojis
	emojis, err := gomoji.TransformText(shortcodes, gomoji.FormatEmoji)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   Back to emojis: %s\n", emojis)

	// Real-world example: processing user messages
	userMessage := "Great work! 👍 The project is on 🔥! Keep it up! ⭐"
	processed, err := gomoji.TransformText(userMessage, gomoji.FormatShortcode)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   User message: %s\n", userMessage)
	fmt.Printf("   For storage: %s\n", processed)
}

func emojiInformation() {
	// Get complete information about an emoji
	emojis := []string{"rocket", "😊", ":coffee:", "&#x1f308;"}

	for _, input := range emojis {
		info, err := gomoji.GetEmojiInfo(input)
		if err != nil {
			fmt.Printf("   %s: Error - %v\n", input, err)
			continue
		}

		fmt.Printf("   Input: %s\n", input)
		fmt.Printf("     Emoji: %s\n", info.Emoji)
		fmt.Printf("     Shortcode: %s\n", info.Shortcode)
		fmt.Printf("     HTML: %s\n", info.HTML)
		fmt.Printf("     Unicode: %s\n", info.Unicode)
		fmt.Println()
	}
}

func emojiValidation() {
	// Check if emojis are supported
	inputs := []string{"smile", "invalid_emoji", "heart", "❤️", ":nonexistent:", "🚀"}

	fmt.Println("   Checking emoji support:")
	for _, input := range inputs {
		if gomoji.IsSupported(input) {
			emoji, _ := gomoji.Transform(input, gomoji.FormatEmoji)
			fmt.Printf("   ✓ %s -> %s (supported)\n", input, emoji)
		} else {
			fmt.Printf("   ✗ %s (not supported)\n", input)
		}
	}

	// Show supported emoji count
	supported := gomoji.GetSupportedEmojis()
	fmt.Printf("\n   Total supported emojis: %d\n", len(supported))

	// Show some examples
	fmt.Println("   Some supported emojis:")
	for i, name := range supported[:10] {
		emoji, _ := gomoji.Transform(name, gomoji.FormatEmoji)
		fmt.Printf("     %d. %s -> %s\n", i+1, name, emoji)
	}
}

func webDevelopment() {
	// Scenario 1: Converting emojis for safe HTML rendering
	content := "Welcome to our site! 😊 We hope you enjoy your stay! ⭐"

	htmlSafe, err := gomoji.TransformText(content, gomoji.FormatHTML)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   Original: %s\n", content)
	fmt.Printf("   HTML-safe: %s\n", htmlSafe)

	// Scenario 2: Converting emojis for Markdown
	markdown, err := gomoji.TransformText(content, gomoji.FormatShortcode)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	fmt.Printf("   Markdown: %s\n", markdown)

	// Scenario 3: Database storage and retrieval
	fmt.Println("\n   Database workflow:")
	userPost := "Just deployed my app! 🚀 So excited! 🎉"

	// Store as shortcodes
	forStorage, _ := gomoji.TransformText(userPost, gomoji.FormatShortcode)
	fmt.Printf("   Store in DB: %s\n", forStorage)

	// Retrieve and display
	forDisplay, _ := gomoji.TransformText(forStorage, gomoji.FormatEmoji)
	fmt.Printf("   Display to user: %s\n", forDisplay)
}

func formatConversions() {
	// Demonstrate all possible format conversions
	original := "smile"

	fmt.Printf("   Starting with: %s\n", original)

	// Convert to all formats
	emoji, _ := gomoji.Transform(original, gomoji.FormatEmoji)
	shortcode, _ := gomoji.Transform(original, gomoji.FormatShortcode)
	html, _ := gomoji.Transform(original, gomoji.FormatHTML)
	unicode, _ := gomoji.Transform(original, gomoji.FormatUnicode)

	fmt.Printf("   Emoji: %s\n", emoji)
	fmt.Printf("   Shortcode: %s\n", shortcode)
	fmt.Printf("   HTML: %s\n", html)
	fmt.Printf("   Unicode: %s\n", unicode)

	// Now convert from emoji back to other formats
	fmt.Printf("\n   Starting with emoji %s:\n", emoji)

	backToShortcode, _ := gomoji.Transform(emoji, gomoji.FormatShortcode)
	backToHTML, _ := gomoji.Transform(emoji, gomoji.FormatHTML)
	backToUnicode, _ := gomoji.Transform(emoji, gomoji.FormatUnicode)

	fmt.Printf("   -> Shortcode: %s\n", backToShortcode)
	fmt.Printf("   -> HTML: %s\n", backToHTML)
	fmt.Printf("   -> Unicode: %s\n", backToUnicode)

	// Demonstrate round-trip conversion
	fmt.Println("\n   Round-trip conversion test:")
	testEmojis := []string{"heart", "fire", "rocket", "coffee", "pizza"}

	for _, name := range testEmojis {
		// name -> emoji -> shortcode -> emoji
		e1, _ := gomoji.Transform(name, gomoji.FormatEmoji)
		s1, _ := gomoji.Transform(e1, gomoji.FormatShortcode)
		e2, _ := gomoji.Transform(s1, gomoji.FormatEmoji)

		if e1 == e2 {
			fmt.Printf("   ✓ %s: %s -> %s -> %s ✓\n", name, e1, s1, e2)
		} else {
			fmt.Printf("   ✗ %s: %s != %s\n", name, e1, e2)
		}
	}
}
