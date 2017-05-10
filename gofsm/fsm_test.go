package gofsm

import (
	"fmt"
	"log"
	"testing"
)

type Turnstile struct {
	ID         uint64
	EventCount uint64
	CoinCount  uint64
	PassCount  uint64
	State      string
	States     []string
}

// TurnstileEventProcessor is used to handle turnstile actions.
type TurnstileEventProcessor struct{}

func (p *TurnstileEventProcessor) OnExit(fromState string, args []interface{}) {
	t := args[0].(*Turnstile)
	if t.State != fromState {
		panic(fmt.Errorf("转门 %v 的状态与期望的状态 %s 不一致，可能在状态机外被改变了", t, fromState))
	}

	log.Printf("转门 %d 从状态 %s 改变", t.ID, fromState)
}

func (p *TurnstileEventProcessor) Action(action string, fromState string, toState string, args []interface{}) {
	t := args[0].(*Turnstile)
	t.EventCount++

	switch action {
	case "pass": //用户通过的action
		t.PassCount++
	case "check", "repeat-check": //刷卡或者投币的action
		t.CoinCount++
	default: //其它action
	}
}

func (p *TurnstileEventProcessor) OnEnter(toState string, args []interface{}) {
	t := args[0].(*Turnstile)
	t.State = toState
	t.States = append(t.States, toState)

	log.Printf("转门 %d 的状态改变为 %s ", t.ID, toState)
}
func TestFSM(t *testing.T) {
	ts := &Turnstile{
		ID:     1,
		State:  "Locked",
		States: []string{"Locked"},
	}
	fsm := initFSM()

	//推门
	//没刷卡/投币不可进入
	err := fsm.Trigger(ts.State, "Push", ts)
	if err != nil {
		log.Printf("trigger err: %v", err)
	}

	//推门
	//没刷卡/投币不可进入
	err = fsm.Trigger(ts.State, "Push", ts)
	if err != nil {
		log.Printf("trigger err: %v", err)
	}

	//刷卡或者投币
	//不容易啊，终于解锁了
	err = fsm.Trigger(ts.State, "Coin", ts)
	if err != nil {
		log.Printf("trigger err: %v", err)
	}

	//刷卡或者投币
	//这时才解锁
	err = fsm.Trigger(ts.State, "Coin", ts)
	if err != nil {
		log.Printf("trigger err: %v", err)
	}

	//推门
	//这时才能进入，进入后闸门被锁
	err = fsm.Trigger(ts.State, "Push", ts)
	if err != nil {
		log.Printf("trigger err: %v", err)
	}

	//推门
	//无法进入，闸门已锁
	err = fsm.Trigger(ts.State, "Push", ts)
	if err != nil {
		log.Printf("trigger err: %v", err)
	}

	lastState := Turnstile{
		ID:         1,
		EventCount: 6,
		CoinCount:  2,
		PassCount:  1,
		State:      "Locked",
		States:     []string{"Locked", "Unlocked", "Locked"},
	}

	errExport := fsm.Export("fsm.png")
	if errExport != nil {
		log.Printf("%s", errExport)
		return
	}

	if !compareTurnstile(&lastState, ts) {
		log.Printf("Expected last state: %+v, but got %+v", lastState, ts)
	} else {
		log.Printf("最终的状态: %+v", ts)
	}
}

func compareTurnstile(t1 *Turnstile, t2 *Turnstile) bool {
	if t1.ID != t2.ID || t1.CoinCount != t2.CoinCount || t1.EventCount != t2.EventCount || t1.PassCount != t2.PassCount ||
		t1.State != t2.State {
		return false
	}

	return fmt.Sprint(t1.States) == fmt.Sprint(t2.States)
}

func initFSM() *StateMachine {
	delegate := &DefaultDelegate{P: &TurnstileEventProcessor{}}

	transitions := []Transition{
		{From: "Locked", Event: "Coin", To: "Unlocked", Action: "check"},
		{From: "Locked", Event: "Push", To: "Locked", Action: "invalid-push"},
		{From: "Unlocked", Event: "Push", To: "Locked", Action: "pass"},
		{From: "Unlocked", Event: "Coin", To: "Unlocked", Action: "repeat-check"},
	}

	return NewStateMachine(delegate, transitions...)
}
