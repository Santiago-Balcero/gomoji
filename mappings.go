package gomoji

// Reverse mappings for quick lookups from any format to emoji name
var (
	emojiToName     map[string]string
	shortcodeToName map[string]string
	htmlToName      map[string]string
	unicodeToName   map[string]string
)

// init initializes reverse mappings for fast emoji lookups.
func init() {
	emojiToName = make(map[string]string)
	shortcodeToName = make(map[string]string)
	htmlToName = make(map[string]string)
	unicodeToName = make(map[string]string)

	for name, mapping := range emojiMappings {
		emojiToName[mapping.Emoji] = name
		shortcodeToName[mapping.Shortcode] = name
		htmlToName[mapping.HTML] = name
		unicodeToName[mapping.Unicode] = name
	}
}
