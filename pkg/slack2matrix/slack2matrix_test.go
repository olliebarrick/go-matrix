package slack2matrix

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMarkdownStringToHTML(t *testing.T) {
	codeBlock, err := MarkdownString("```hello```").ToHTML()
	assert.Nil(t, err)
	assert.Equal(t, `<code>hello</code>`, codeBlock)
}

func TestMarkdownStringToHTMLLinks(t *testing.T) {
	codeBlock, err := MarkdownString("Pushed to tag <https://gitlab/kubernetes/manifests/commits/flux-sync-flux|flux-sync-flux> of <https://gitlab/kubernetes/manifests|kubernetes/manifests> (<https://gitlab/kubernetes/manifests/compare/cb8aedae1951dcd340740a2fcc3c7c0336371054...029f886cd4f5e0220ddb13d749c068fae5c610bd|Compare changes>)").ToHTML()
	assert.Nil(t, err)
	assert.Equal(t, `Pushed to tag <a href="https://gitlab/kubernetes/manifests/commits/flux-sync-flux">flux-sync-flux</a> of <a href="https://gitlab/kubernetes/manifests">kubernetes/manifests</a> (<a href="https://gitlab/kubernetes/manifests/compare/cb8aedae1951dcd340740a2fcc3c7c0336371054...029f886cd4f5e0220ddb13d749c068fae5c610bd">Compare changes</a>)`, codeBlock)
}

func TestMarkdownStringReplaceLinks(t *testing.T) {
	blob := MarkdownString("Pushed to tag <https://gitlab/kubernetes/manifests/commits/flux-sync-flux|flux-sync-flux> of <https://gitlab/kubernetes/manifests|kubernetes/manifests> (<https://gitlab/kubernetes/manifests/compare/cb8aedae1951dcd340740a2fcc3c7c0336371054...029f886cd4f5e0220ddb13d749c068fae5c610bd|Compare changes>)").ReplaceLinks()
	assert.Equal(t, `Pushed to tag [flux-sync-flux](https://gitlab/kubernetes/manifests/commits/flux-sync-flux) of [kubernetes/manifests](https://gitlab/kubernetes/manifests) ([Compare changes](https://gitlab/kubernetes/manifests/compare/cb8aedae1951dcd340740a2fcc3c7c0336371054...029f886cd4f5e0220ddb13d749c068fae5c610bd))`, blob)
}

func TestSlackMessageToHTML(t *testing.T) {
	msg := SlackMessage{
		Text: "Justin Barrick pushed to tag <https://gitlab/kubernetes/manifests/commits/flux-sync-flux|flux-sync-flux> of <https://gitlab/kubernetes/manifests|kubernetes/manifests> (<https://gitlab/kubernetes/manifests/compare/cb8aedae1951dcd340740a2fcc3c7c0336371054...029f886cd4f5e0220ddb13d749c068fae5c610bd|Compare changes>)",
		Attachments: []SlackAttachment{
			SlackAttachment{
				Text: "<https://gitlab/kubernetes/manifests/commit/93a98d81006985e03b1bb2b5f72ccfdd2a40eb8a|93a98d81>: gitlab change\n - Justin Barrick",
				Color: "#345",
			},
		},
	}

	body, err := msg.ToHTML()
	assert.Nil(t, err)
	assert.Equal(t, `Justin Barrick pushed to tag <a href="https://gitlab/kubernetes/manifests/commits/flux-sync-flux">flux-sync-flux</a> of <a href="https://gitlab/kubernetes/manifests">kubernetes/manifests</a> (<a href="https://gitlab/kubernetes/manifests/compare/cb8aedae1951dcd340740a2fcc3c7c0336371054...029f886cd4f5e0220ddb13d749c068fae5c610bd">Compare changes</a>)<br><br><a href="https://gitlab/kubernetes/manifests/commit/93a98d81006985e03b1bb2b5f72ccfdd2a40eb8a">93a98d81</a>: gitlab change
 - Justin Barrick<br>`, body)
}