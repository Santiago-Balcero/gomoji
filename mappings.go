package gomoji

import "strings"

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

		// For emojis with variation selectors, also support base formats
		if strings.Contains(mapping.HTML, "&#xfe0f;") {
			// Support HTML base entity: &#x1f399;&#xfe0f; -> &#x1f399;
			htmlBase := strings.Replace(mapping.HTML, "&#xfe0f;", "", 1)
			htmlToName[htmlBase] = name

			// Support hybrid format: &#x1f399; + ️ -> &#x1f399;️
			hybridHTML := htmlBase + "️"
			htmlToName[hybridHTML] = name
		}

		if strings.Contains(mapping.Unicode, "\\uFE0F") {
			// Support Unicode base: \\U0001F399\\uFE0F -> \\U0001F399
			unicodeBase := strings.Replace(mapping.Unicode, "\\uFE0F", "", 1)
			unicodeToName[unicodeBase] = name
		}
	}
}
