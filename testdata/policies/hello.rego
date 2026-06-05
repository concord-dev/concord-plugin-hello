package concord.hello

default allow := false

allow {
  input.evidence.greeting.message == "hello, concord"
}
