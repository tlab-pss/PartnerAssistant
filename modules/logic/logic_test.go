package logic

import (
	"testing"

	"github.com/sskmy1024/PartnerAssistant/infrastructures"
	"github.com/stretchr/testify/assert"
)

func TestMainLogic(t *testing.T) {
	infrastructures.InitEnvWithPath("../../")

	res, err := ExecuteLogic("お腹すいた")
	if err != nil {
		t.Errorf("Cannot execute Logic: %+v", err)
	}
	assert.Equal(t, "そうですか", res.Message)
}

func TestGetRequireService(t *testing.T) {
	infrastructures.InitEnvWithPath("../../")

	replyData, err := RequestAI("お腹すいた")
	if err != nil {
		t.Errorf("Cannot access watson: %+v", err)
	}

	requestArgs := ConvertRequireServiceType(replyData)

	assert.Equal(t, "Gourmet", requestArgs.TopicCategory.String())
	assert.Equal(t, true, requestArgs.RequireService)
}
