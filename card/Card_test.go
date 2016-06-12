package card

import (
	"fmt"
	"testing"
)

func TestCardInfo(t *testing.T) {
	var x interface{} = new(MonsterCard)
	fmt.Println(x)
	
	var _,ok = x.(ICard)
	if !ok {
		t.Error("MonsterCard不符合ICard接口")
	}
	_,ok = x.(IMonsterCard)
	if !ok {
		t.Error("MonsterCard不符合IMonsterCard接口")
	}
	
	x = new(SpellCard)
	fmt.Println(x)
	_,ok = x.(ICard)
	if !ok {
		t.Error("SpellCard不符合ICard接口")
	}
	_,ok = x.(ISpellCard)
	if !ok {
		t.Error("SpellCard不符合ISpellCard接口")
	}
	
	x = new(TrapCard)
	fmt.Println(x)
	_,ok = x.(ICard)
	if !ok {
		t.Error("TrapCard不符合ICard接口")
	}
	_,ok = x.(ITrapCard)
	if !ok {
		t.Error("TrapCard不符合ITrapCard接口")
	}
	
	
}