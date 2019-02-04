package main

type mailbox struct {
	Selected *item
	All      []*item
}

type item struct {
	ID, Date, From, To, Subject, Body string
}
