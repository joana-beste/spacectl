package spacefile

import (
	"html/template"
	"io"
)

const SpacefileTemplate = `version = "1"

// This file was auto-generated by "spacectl space init"
// Feel free to adjust it to your own needs.

space "{{ .SpaceDNSLabel }}" {
  name = "{{ .SpaceName }}"
  team = "{{ .TeamName }}"

  stage production {
    application typo3 {

      // The "version" field supports semantic version ranges. Valid
      // examples might be "8.7.0", "~8.7.0", ">=7.0.0, <7.3".
      // We will always pick the latest available version that matches
      // this constraint and update them regularly, so do not specify
      // your version range too loosly for a production environment.
      version = "~8.7.0"

      userData {
        initialAdminUser {
          username = "admin"
          password = "my-password"
        }
      }
    }
  }

  // You can add additional stages to your Space definition
  // Use the "inherit" attribute to have individual stages inherit
  // their configuration from each other.

  // stage development {
  //   inherit = "production"
  // }
}
`

type templateData struct {
	TeamName string
	SpaceName string
	SpaceDNSLabel string
}

func Generate(teamName string, spaceName string, spaceDNSLabel string, out io.Writer) (error) {
	t := template.Must(template.New("spacefile").Parse(SpacefileTemplate))
	return t.Execute(out, templateData{
		teamName,
		spaceName,
		spaceDNSLabel,
	})
}