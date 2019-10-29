package logic

import (
	"fmt"
	"testing"

	"github.com/sskmy1024/PartnerAssistant/infrastructures"
	topiccategory "github.com/sskmy1024/PartnerAssistant/models/category/topic_category"
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

func TestRequestAI(t *testing.T) {
	infrastructures.InitEnvWithPath("../../")

	replyData, err := RequestAI("ラーメン食べたい")
	if err != nil {
		t.Errorf("Cannot access watson: %+v", err)
	}
	fmt.Printf("%+v", replyData.Result.Context)

	assert.Equal(t, topiccategory.Gourmet, replyData.TopicCategory())
	assert.Equal(t, "ラーメン", replyData.Result.Context.Value)
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
