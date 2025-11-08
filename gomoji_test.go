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
	}{
		{
			name:         "transform emojis to shortcodes",
			input:        "Hello 😊 I love 🌈 and ✨!",
			targetFormat: FormatShortcode,
			expected:     "Hello :blush: I love :rainbow: and :sparkles:!",
		},
		{
			name:         "transform shortcodes to emojis",
			input:        "Hello :blush: I love :rainbow: and :sparkles:!",
			targetFormat: FormatEmoji,
			expected:     "Hello 😊 I love 🌈 and ✨!",
		},
		{
			name:         "transform mixed formats",
			input:        "🎤 Today we have :sparkles: and &#x1f308;",
			targetFormat: FormatEmoji,
			expected:     "🎤 Today we have ✨ and 🌈",
		},
		{
			name:         "transform to html",
			input:        "Check out this 🔥 content!",
			targetFormat: FormatHTML,
			expected:     "Check out this &#x1f525; content!",
		},
		{
			name:         "no emojis in text",
			input:        "This is just plain text",
			targetFormat: FormatEmoji,
			expected:     "This is just plain text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TransformText(tt.input, tt.targetFormat)
			if result != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, result)
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

// Test newly added emojis
func TestNewlyAddedEmojis(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// Transportation
		{"bicycle", "bicycle", "🚲"},
		{"scooter", "scooter", "🛵"},
		{"motorcycle", "motorcycle", "🏍️"},
		{"racing_car", "racing_car", "🏎️"},

		// Technology & Objects
		{"floppy_disk", "floppy_disk", "💾"},
		{"desktop_computer", "desktop_computer", "🖥️"},
		{"camera_flash", "camera_flash", "📸"},
		{"calendar", "calendar", "📅"},

		// Nature - Celestial
		{"sun", "sun", "☀️"},
		{"fire", "fire", "🔥"},
		{"moon", "moon", "🌙"},
		{"star2", "star2", "🌟"},

		// Nature - Water & Earth
		{"droplet", "droplet", "💧"},
		{"ocean", "ocean", "🌊"},
		{"earth_africa", "earth_africa", "🌍"},
		{"earth_americas", "earth_americas", "🌎"},
		{"earth_asia", "earth_asia", "🌏"},
		{"desert_island", "desert_island", "🏝️"},

		// Buildings & Places
		{"classical_building", "classical_building", "🏛️"},

		// Symbols
		{"bomb", "bomb", "💣"},

		// Flags by Continent
		{"flag_it", "flag_it", "🇮🇹"},
		{"flag_fr", "flag_fr", "🇫🇷"},
		{"flag_us", "flag_us", "🇺🇸"},
		{"flag_co", "flag_co", "🇨🇴"},
		{"flag_ar", "flag_ar", "🇦🇷"},
		{"flag_mx", "flag_mx", "🇲🇽"},
		{"flag_br", "flag_br", "🇧🇷"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Transform(tt.input, FormatEmoji)
			if err != nil {
				t.Errorf("Transform(%s) returned error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("Transform(%s) = %s, expected %s", tt.input, result, tt.expected)
			}

			// Test that all new emojis are supported
			if !IsSupported(tt.input) {
				t.Errorf("IsSupported(%s) = false, expected true", tt.input)
			}

			// Test GetEmojiInfo works for all new emojis
			info, err := GetEmojiInfo(tt.input)
			if err != nil {
				t.Errorf("GetEmojiInfo(%s) returned error: %v", tt.input, err)
				return
			}
			if info.Emoji != tt.expected {
				t.Errorf("GetEmojiInfo(%s).Emoji = %s, expected %s", tt.input, info.Emoji, tt.expected)
			}
		})
	}
}

// Test text transformation with multiple new emojis
func TestTransformTextWithNewEmojis(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		format   Format
		expected string
	}{
		{
			name:     "transportation emojis",
			input:    "I love riding my 🚲 and driving my 🏎️!",
			format:   FormatShortcode,
			expected: "I love riding my :bicycle: and driving my :racing_car:!",
		},
		{
			name:     "nature emojis",
			input:    "The ☀️ is shining and the 🌊 are beautiful!",
			format:   FormatShortcode,
			expected: "The :sun: is shining and the :ocean: are beautiful!",
		},
		{
			name:     "technology emojis",
			input:    "Working on my 🖥️ and 💻 with a 📸 nearby.",
			format:   FormatShortcode,
			expected: "Working on my :desktop_computer: and :computer: with a :camera_flash: nearby.",
		},
		{
			name:     "flags emojis",
			input:    "Visiting 🇮🇹 and 🇫🇷 this summer!",
			format:   FormatShortcode,
			expected: "Visiting :flag_it: and :flag_fr: this summer!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TransformText(tt.input, tt.format)
			if result != tt.expected {
				t.Errorf("TransformText(%s, %v) = %s, expected %s", tt.input, tt.format, result, tt.expected)
			}
		})
	}
}

// Test hybrid HTML format handling
func TestHybridHTMLFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "studio microphone hybrid HTML",
			input:    "&#x1f399;️", // HTML entity + actual variation selector emoji
			expected: "🎙️",
		},
		{
			name:     "keyboard hybrid HTML",
			input:    "&#x2328;️", // HTML entity + actual variation selector emoji
			expected: "⌨️",
		},
		{
			name:     "mouse hybrid HTML",
			input:    "&#x1f5b1;️", // HTML entity + actual variation selector emoji
			expected: "🖱️",
		},
		{
			name:     "desktop computer hybrid HTML",
			input:    "&#x1f5a5;️", // HTML entity + actual variation selector emoji
			expected: "🖥️",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Transform(tt.input, FormatEmoji)
			if err != nil {
				t.Errorf("Transform(%s) returned error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("Transform(%s) = %s, expected %s", tt.input, result, tt.expected)
			}

			// Also test that it's properly supported
			if !IsSupported(tt.input) {
				t.Errorf("IsSupported(%s) = false, expected true", tt.input)
			}
		})
	}
}

// Test flexible format support for all emojis with variation selectors
func TestFlexibleFormatSupport(t *testing.T) {
	// Test all emojis that have variation selectors
	testCases := []struct {
		name            string
		completeHTML    string
		baseHTML        string
		hybridHTML      string
		completeUnicode string
		baseUnicode     string
		expectedEmoji   string
	}{
		{
			name:            "relaxed",
			completeHTML:    "&#x263a;&#xfe0f;",
			baseHTML:        "&#x263a;",
			hybridHTML:      "&#x263a;️",
			completeUnicode: "\\U0000263A\\uFE0F",
			baseUnicode:     "\\U0000263A",
			expectedEmoji:   "☺️",
		},
		{
			name:            "frowning",
			completeHTML:    "&#x2639;&#xfe0f;",
			baseHTML:        "&#x2639;",
			hybridHTML:      "&#x2639;️",
			completeUnicode: "\\U00002639\\uFE0F",
			baseUnicode:     "\\U00002639",
			expectedEmoji:   "☹️",
		},
		{
			name:            "point_up",
			completeHTML:    "&#x261d;&#xfe0f;",
			baseHTML:        "&#x261d;",
			hybridHTML:      "&#x261d;️",
			completeUnicode: "\\U0000261D\\uFE0F",
			baseUnicode:     "\\U0000261D",
			expectedEmoji:   "☝️",
		},
		{
			name:            "peace",
			completeHTML:    "&#x270c;&#xfe0f;",
			baseHTML:        "&#x270c;",
			hybridHTML:      "&#x270c;️",
			completeUnicode: "\\U0000270C\\uFE0F",
			baseUnicode:     "\\U0000270C",
			expectedEmoji:   "✌️",
		},
		{
			name:            "hand_splayed",
			completeHTML:    "&#x1f590;&#xfe0f;",
			baseHTML:        "&#x1f590;",
			hybridHTML:      "&#x1f590;️",
			completeUnicode: "\\U0001F590\\uFE0F",
			baseUnicode:     "\\U0001F590",
			expectedEmoji:   "🖐️",
		},
		{
			name:            "heart",
			completeHTML:    "&#x2764;&#xfe0f;",
			baseHTML:        "&#x2764;",
			hybridHTML:      "&#x2764;️",
			completeUnicode: "\\U00002764\\uFE0F",
			baseUnicode:     "\\U00002764",
			expectedEmoji:   "❤️",
		},
		{
			name:            "sun",
			completeHTML:    "&#x2600;&#xfe0f;",
			baseHTML:        "&#x2600;",
			hybridHTML:      "&#x2600;️",
			completeUnicode: "\\U00002600\\uFE0F",
			baseUnicode:     "\\U00002600",
			expectedEmoji:   "☀️",
		},
		{
			name:            "cloud",
			completeHTML:    "&#x2601;&#xfe0f;",
			baseHTML:        "&#x2601;",
			hybridHTML:      "&#x2601;️",
			completeUnicode: "\\U00002601\\uFE0F",
			baseUnicode:     "\\U00002601",
			expectedEmoji:   "☁️",
		},
		{
			name:            "airplane",
			completeHTML:    "&#x2708;&#xfe0f;",
			baseHTML:        "&#x2708;",
			hybridHTML:      "&#x2708;️",
			completeUnicode: "\\U00002708\\uFE0F",
			baseUnicode:     "\\U00002708",
			expectedEmoji:   "✈️",
		},
		{
			name:            "keyboard",
			completeHTML:    "&#x2328;&#xfe0f;",
			baseHTML:        "&#x2328;",
			hybridHTML:      "&#x2328;️",
			completeUnicode: "\\U00002328\\uFE0F",
			baseUnicode:     "\\U00002328",
			expectedEmoji:   "⌨️",
		},
		{
			name:            "mouse_three_button",
			completeHTML:    "&#x1f5b1;&#xfe0f;",
			baseHTML:        "&#x1f5b1;",
			hybridHTML:      "&#x1f5b1;️",
			completeUnicode: "\\U0001F5B1\\uFE0F",
			baseUnicode:     "\\U0001F5B1",
			expectedEmoji:   "🖱️",
		},
		{
			name:            "studio_microphone",
			completeHTML:    "&#x1f399;&#xfe0f;",
			baseHTML:        "&#x1f399;",
			hybridHTML:      "&#x1f399;️",
			completeUnicode: "\\U0001F399\\uFE0F",
			baseUnicode:     "\\U0001F399",
			expectedEmoji:   "🎙️",
		},
		{
			name:            "desktop_computer",
			completeHTML:    "&#x1f5a5;&#xfe0f;",
			baseHTML:        "&#x1f5a5;",
			hybridHTML:      "&#x1f5a5;️",
			completeUnicode: "\\U0001F5A5\\uFE0F",
			baseUnicode:     "\\U0001F5A5",
			expectedEmoji:   "🖥️",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test complete formats (should work)
			result, err := Transform(tc.completeHTML, FormatEmoji)
			if err != nil {
				t.Errorf("Transform(%s) failed: %v", tc.completeHTML, err)
			} else if result != tc.expectedEmoji {
				t.Errorf("Transform(%s) = %s, expected %s", tc.completeHTML, result, tc.expectedEmoji)
			}

			result, err = Transform(tc.completeUnicode, FormatEmoji)
			if err != nil {
				t.Errorf("Transform(%s) failed: %v", tc.completeUnicode, err)
			} else if result != tc.expectedEmoji {
				t.Errorf("Transform(%s) = %s, expected %s", tc.completeUnicode, result, tc.expectedEmoji)
			}

			// Test base formats (should now work)
			result, err = Transform(tc.baseHTML, FormatEmoji)
			if err != nil {
				t.Errorf("Transform(%s) failed: %v", tc.baseHTML, err)
			} else if result != tc.expectedEmoji {
				t.Errorf("Transform(%s) = %s, expected %s", tc.baseHTML, result, tc.expectedEmoji)
			}

			result, err = Transform(tc.baseUnicode, FormatEmoji)
			if err != nil {
				t.Errorf("Transform(%s) failed: %v", tc.baseUnicode, err)
			} else if result != tc.expectedEmoji {
				t.Errorf("Transform(%s) = %s, expected %s", tc.baseUnicode, result, tc.expectedEmoji)
			}

			// Test hybrid format (should now work)
			result, err = Transform(tc.hybridHTML, FormatEmoji)
			if err != nil {
				t.Errorf("Transform(%s) failed: %v", tc.hybridHTML, err)
			} else if result != tc.expectedEmoji {
				t.Errorf("Transform(%s) = %s, expected %s", tc.hybridHTML, result, tc.expectedEmoji)
			}

			// Test IsSupported for all formats
			if !IsSupported(tc.completeHTML) {
				t.Errorf("IsSupported(%s) = false, expected true", tc.completeHTML)
			}
			if !IsSupported(tc.baseHTML) {
				t.Errorf("IsSupported(%s) = false, expected true", tc.baseHTML)
			}
			if !IsSupported(tc.hybridHTML) {
				t.Errorf("IsSupported(%s) = false, expected true", tc.hybridHTML)
			}
			if !IsSupported(tc.completeUnicode) {
				t.Errorf("IsSupported(%s) = false, expected true", tc.completeUnicode)
			}
			if !IsSupported(tc.baseUnicode) {
				t.Errorf("IsSupported(%s) = false, expected true", tc.baseUnicode)
			}
		})
	}
}

// Test backward compatibility - ensure all existing functionality still works
func TestBackwardCompatibility(t *testing.T) {
	// Test that all formats still work for emojis without variation selectors
	simpleTests := []struct {
		name     string
		input    string
		expected string
	}{
		{"smile by name", "smile", "😄"},
		{"smile by emoji", "😄", "😄"},
		{"smile by shortcode", ":smile:", "😄"},
		{"smile by html", "&#x1f604;", "😄"},
		{"smile by unicode", "\\U0001F604", "😄"},
	}

	for _, tt := range simpleTests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Transform(tt.input, FormatEmoji)
			if err != nil {
				t.Errorf("Transform(%s) failed: %v", tt.input, err)
			}
			if result != tt.expected {
				t.Errorf("Transform(%s) = %s, expected %s", tt.input, result, tt.expected)
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
		_ = TransformText(text, FormatShortcode)
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

func BenchmarkTransformBaseHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Transform("&#x1f399;", FormatEmoji)
	}
}

func BenchmarkTransformHybridHTML(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Transform("&#x1f399;️", FormatEmoji)
	}
}

func BenchmarkTransformBaseUnicode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Transform("\\U0001F399", FormatEmoji)
	}
}
