package spacefile

import (
	"github.com/mittwald/spacectl/client/software"
	"io"
	"text/template"
)

const SpacefileTemplate = `version = "1"

// This file was auto-generated by "spacectl space init"
// Feel free to adjust it to your own needs.

space "{{ .SpaceDNSLabel }}" {
  name = "{{ .SpaceName }}"
  team = "{{ .TeamName }}"

  stage production {
    application {{ .Software.Identifier }} {

      // The "version" field supports semantic version ranges. Valid
      // examples might be "8.7.0", "~8.7.0", ">=7.0.0, <7.3".
      // We will always pick the latest available version that matches
      // this constraint and update them regularly, so do not specify
      // your version range too loosly for a production environment.
      version = "~{{ .Software.LatestVersion.Number }}"

      userData {
        initialAdminUser {
          username = "admin"
          password = "my-password"
        }
      }
    }

    // cron helloworld {
    //   schedule = "30 * * * *"
    //   command {
    //     command = "echo"
    //     arguments = ["Hello World"]
    //     workingDirectory = "/var/www"
    //   }
    // }
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
	Software *software.Software
}

func Generate(teamName string, spaceName string, spaceDNSLabel string, software *software.Software, out io.Writer) (error) {
	t := template.Must(template.New("spacefile").Parse(SpacefileTemplate))
	return t.Execute(out, templateData{
		teamName,
		spaceName,
		spaceDNSLabel,
		software,
	})
}