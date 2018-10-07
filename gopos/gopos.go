package main

import (
	"fmt"
	"sync"
	"time"
)

// -------------- context

type data string

type context struct {
	g       string
	resolve chan data
}

func (ctx *context) dispose(data data) {
	fmt.Println(time.Now(), ctx.g, "stores", data, "and is then disposed of")
}

// -------------- contexts

type contexts struct {
	m    sync.Mutex
	list map[string]context
}

func newContexts() contexts {
	return contexts{list: make(map[string]context)}
}

func provideData(ctx context) {
	<-time.After(time.Second * 1)
	s := fmt.Sprint(ctx.g, " is here")
	fmt.Println(time.Now(), "Data loaded", s)
	ctx.resolve <- data(s)
}

func (cs *contexts) Get(g string) context {
	cs.m.Lock()
	defer cs.m.Unlock()

	if ctx, ok := cs.list[g]; ok {
		fmt.Println(time.Now(), cs.list)
		return ctx
	}
	ctx := context{g: g, resolve: make(chan data)}
	cs.list[g] = ctx
	fmt.Println(time.Now(), cs.list)
	go provideData(ctx)
	return ctx
}

func (cs *contexts) Return(ctx context, data data, s string) {
	fmt.Println(time.Now(), "Returning", ctx.g, "of", s)

	cs.m.Lock()
	defer cs.m.Unlock()

	select {
	case ctx.resolve <- data:
		// nothing to do, next one with a ready channel receives the data
	default:
		delete(cs.list, ctx.g)
		fmt.Println(time.Now(), cs.list)
		ctx.dispose(data)
	}
}

// -------------- main

func p(cs *contexts, ctx context, s string) {
	dat := <-ctx.resolve
	fmt.Println(time.Now(), "resolved", s, "with", dat)
	cs.Return(ctx, data("("+dat+")"), s)
}

func main() {
	fmt.Println(time.Now(), "This is go POS - a channel based workitem dispatcher")
	cs := newContexts()

	for i := 0; i < 10; i++ {
		i:=i
		ctx1 := cs.Get("1")
		go p(&cs, ctx1, "ctx1")

		ctx2a := cs.Get("ab")
		go p(&cs, ctx2a, "ctx2a")

		name := fmt.Sprint(i)
		ctx2b := cs.Get(name)
		go p(&cs, ctx2b, "ctx2b")

		<-time.After(time.Millisecond * 600)
		ctx2c := cs.Get(name)
		go p(&cs, ctx2c, "ctx2c")

		<-time.After(time.Millisecond * 600)
		ctx2d := cs.Get(name)
		go p(&cs, ctx2d, "ctx2d")
	}

	<-time.After(time.Second * 2)
}
