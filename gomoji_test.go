package gomoji

import "testing"

func TestTransform(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		targetFormat Format
		expected     string
		shouldError  bool
	}{
		// Test transforming emoji name to different formats
		{
			name:         "smile to emoji",
			input:        "smile",
			targetFormat: FormatEmoji,
			expected:     "😄",
			shouldError:  false,
		},
		{
			name:         "smile to shortcode",
			input:        "smile",
			targetFormat: FormatShortcode,
			expected:     ":smile:",
			shouldError:  false,
		},
		{
			name:         "smile to html",
			input:        "smile",
			targetFormat: FormatHTML,
			expected:     "&#x1f604;",
			shouldError:  false,
		},
		{
			name:         "smile to unicode",
			input:        "smile",
			targetFormat: FormatUnicode,
			expected:     "\\U0001F604",
			shouldError:  false,
		},

		// Test transforming actual emoji to other formats
		{
			name:         "emoji to shortcode",
			input:        "😊",
			targetFormat: FormatShortcode,
			expected:     ":blush:",
			shouldError:  false,
		},
		{
			name:         "emoji to html",
			input:        "🌈",
			targetFormat: FormatHTML,
			expected:     "&#x1f308;",
			shouldError:  false,
		},
		{
			name:         "emoji to unicode",
			input:        "🌈",
			targetFormat: FormatUnicode,
			expected:     "\\U0001F308",
			shouldError:  false,
		},

		// Test transforming shortcode to other formats
		{
			name:         "shortcode to emoji",
			input:        ":sparkles:",
			targetFormat: FormatEmoji,
			expected:     "✨",
			shouldError:  false,
		},
		{
			name:         "shortcode to html",
			input:        ":rainbow:",
			targetFormat: FormatHTML,
			expected:     "&#x1f308;",
			shouldError:  false,
		},

		// Test transforming HTML to other formats
		{
			name:         "html to emoji",
			input:        "&#x1f60a;",
			targetFormat: FormatEmoji,
			expected:     "😊",
			shouldError:  false,
		},
		{
			name:         "html to shortcode",
			input:        "&#x2728;",
			targetFormat: FormatShortcode,
			expected:     ":sparkles:",
			shouldError:  false,
		},

		// Test transforming Unicode to other formats
		{
			name:         "unicode to emoji",
			input:        "\\U0001F60A",
			targetFormat: FormatEmoji,
			expected:     "😊",
			shouldError:  false,
		},
		{
			name:         "unicode to shortcode",
			input:        "\\U00002728",
			targetFormat: FormatShortcode,
			expected:     ":sparkles:",
			shouldError:  false,
		},

		// Test shortcode without colons
		{
			name:         "shortcode without colons",
			input:        "heart",
			targetFormat: FormatEmoji,
			expected:     "❤️",
			shouldError:  false,
		},

		// Test error cases
		{
			name:         "invalid emoji",
			input:        "invalid_emoji",
			targetFormat: FormatEmoji,
			expected:     "",
			shouldError:  true,
		},
		{
			name:         "invalid format",
			input:        "smile",
			targetFormat: Format("invalid_format"),
			expected:     "",
			shouldError:  true,
		},
		{
			name:         "empty input",
			input:        "",
			targetFormat: FormatEmoji,
			expected:     "",
			shouldError:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Transform(tt.input, tt.targetFormat)

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

func TestTransformText(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		targetFormat Format
		expected     string
		shouldError  bool
	}{
		{
			name:         "transform emojis to shortcodes",
			input:        "Hello 😊 I love 🌈 and ✨!",
			targetFormat: FormatShortcode,
			expected:     "Hello :blush: I love :rainbow: and :sparkles:!",
			shouldError:  false,
		},
		{
			name:         "transform shortcodes to emojis",
			input:        "Hello :blush: I love :rainbow: and :sparkles:!",
			targetFormat: FormatEmoji,
			expected:     "Hello 😊 I love 🌈 and ✨!",
			shouldError:  false,
		},
		{
			name:         "transform mixed formats",
			input:        "🎤 Today we have :sparkles: and &#x1f308;",
			targetFormat: FormatEmoji,
			expected:     "🎤 Today we have ✨ and 🌈",
			shouldError:  false,
		},
		{
			name:         "transform to html",
			input:        "Check out this 🔥 content!",
			targetFormat: FormatHTML,
			expected:     "Check out this &#x1f525; content!",
			shouldError:  false,
		},
		{
			name:         "no emojis in text",
			input:        "This is just plain text",
			targetFormat: FormatEmoji,
			expected:     "This is just plain text",
			shouldError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := TransformText(tt.input, tt.targetFormat)

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("expected %q, got %q", tt.expected, result)
				}
			}
		})
	}
}

func TestGetEmojiInfo(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		shouldError bool
	}{
		{
			name:        "get info for blush emoji",
			input:       "😊",
			shouldError: false,
		},
		{
			name:        "get info for smile name",
			input:       "smile",
			shouldError: false,
		},
		{
			name:        "get info for shortcode",
			input:       ":smile:",
			shouldError: false,
		},
		{
			name:        "get info for html",
			input:       "&#x1f604;",
			shouldError: false,
		},
		{
			name:        "get info for unicode",
			input:       "\\U0001F604",
			shouldError: false,
		},
		{
			name:        "get info for invalid emoji",
			input:       "invalid",
			shouldError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			info, err := GetEmojiInfo(tt.input)

			if tt.shouldError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				if info != nil {
					t.Errorf("expected nil info but got %v", info)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if info == nil {
					t.Errorf("expected emoji info but got nil")
				} else {
					// Verify that all fields are populated
					if info.Emoji == "" || info.Shortcode == "" || info.HTML == "" || info.Unicode == "" {
						t.Errorf("emoji info has empty fields: %+v", info)
					}
				}
			}
		})
	}
}

func TestGetSupportedEmojis(t *testing.T) {
	emojis := GetSupportedEmojis()

	if len(emojis) == 0 {
		t.Error("expected at least one supported emoji")
	}

	// Check that some expected emojis are in the list
	expectedEmojis := []string{"smile", "heart", "fire", "microphone", "rainbow"}
	for _, expected := range expectedEmojis {
		found := false
		for _, emoji := range emojis {
			if emoji == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected emoji %q not found in supported emojis", expected)
		}
	}
}

func TestIsSupported(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "supported emoji name",
			input:    "smile",
			expected: true,
		},
		{
			name:     "supported actual emoji",
			input:    "😊",
			expected: true,
		},
		{
			name:     "supported shortcode",
			input:    ":smile:",
			expected: true,
		},
		{
			name:     "supported html",
			input:    "&#x1f604;",
			expected: true,
		},
		{
			name:     "supported unicode",
			input:    "\\U0001F604",
			expected: true,
		},
		{
			name:     "unsupported emoji",
			input:    "invalid",
			expected: false,
		},
		{
			name:     "empty input",
			input:    "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSupported(tt.input)
			if result != tt.expected {
				t.Errorf("expected %t, got %t", tt.expected, result)
			}
		})
	}
}

// Benchmark tests for performance
func BenchmarkTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Transform("smile", FormatEmoji)
	}
}

func BenchmarkTransformText(b *testing.B) {
	text := "Hello 😊 I love 🌈 and ✨ content! 🔥"
	for i := 0; i < b.N; i++ {
		_, _ = TransformText(text, FormatShortcode)
	}
}

func BenchmarkGetEmojiInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GetEmojiInfo("smile")
	}
}

func BenchmarkIsSupported(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = IsSupported("smile")
	}
}
