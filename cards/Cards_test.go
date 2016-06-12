package cards

import (
	"fmt"
	"testing"
)

func TestCards(t *testing.T) {
	var card,ok = CreateCardInstance(10000001,12,233)
	if !ok {
		t.Fatal("创建卡片实例失败")
	}
	//card.Init(nil)
	fmt.Println(card)
	fmt.Println(card.Origin())
}