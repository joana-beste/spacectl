package spaces

import (
	"fmt"
	"net/url"

	"github.com/mittwald/spacectl/client/lowlevel"

	"github.com/mittwald/spacectl/client/errors"
)

// GetStageProtection returns the current Stage Protection type
func (c *spacesClient) GetStageProtection(spaceID, stage string) (*StageProtection, error) {
	var protection StageProtection

	listPath := fmt.Sprintf("/spaces/%s/stages/%s/protection", url.PathEscape(spaceID), url.PathEscape(stage))
	err := c.client.Get(listPath, &protection)
	if err != nil {
		statusErr, ok := err.(lowlevel.ErrUnexpectedStatusCode)
		if ok && statusErr.StatusCode == 404 { // returns 404 if no protection exists
			return &protection, nil
		}
		return nil, errors.ErrNested{Inner: err, Msg: fmt.Sprintf("could not access protection for space: %s, stage: %s", spaceID, stage)}
	}

	return &protection, nil
}

// CreateStageProtection updates or creates the Stage Protection
func (c *spacesClient) CreateStageProtection(spaceID, stage string, inputProtection StageProtection) (*StageProtection, error) {
	var protection StageProtection

	createPath := fmt.Sprintf("/spaces/%s/stages/%s/protection", url.PathEscape(spaceID), url.PathEscape(stage))
	err := c.client.Put(createPath, inputProtection, &protection)
	if err != nil {
		return nil, errors.ErrNested{Inner: err, Msg: fmt.Sprintf("could not create/update protection for space: %s, stage: %s", spaceID, stage)}
	}

	return &protection, err
}

// DeleteStageProtection disables the Stage Protection for the given stage
func (c *spacesClient) DeleteStageProtection(spaceID, stage string) error {
	deletePath := fmt.Sprintf("/spaces/%s/stages/%s/protection", url.PathEscape(spaceID), url.PathEscape(stage))
	err := c.client.Delete(deletePath, nil)
	if err != nil {
		return errors.ErrNested{Inner: err, Msg: fmt.Sprintf("could not delete protection for space: %s, stage: %s", spaceID, stage)}
	}
	return nil
}
