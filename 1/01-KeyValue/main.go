package main

import (
	"fmt"
	"sync"
)

// KeyValue struct
type KeyValue struct {
	store map[string]string
	mu    sync.RWMutex // 排他制御用
}

// NewKeyValue make store
// ポインタ型で返す
func NewKeyValue() *KeyValue {
	return &KeyValue{store: make(map[string]string)} // ポインタを生成して返す
}

// Set on KeyValue
// ポインタのインスタンスを受け取る
func (kv *KeyValue) Set(key, val string) {
	kv.mu.Lock()         // 書き込む前にロック
	defer kv.mu.Unlock() // 処理終了したらロック解除
	kv.store[key] = val
}

// Get from KeyValue
// ポインタのインスタンスを受け取る
func (kv *KeyValue) Get(key string) (string, bool) {
	kv.mu.RLock()         // 読む前にロック
	defer kv.mu.RUnlock() // 処理を抜けたらロック解除
	val, ok := kv.store[key]
	return val, ok
}

func main() {
	kv := NewKeyValue() // ポインタ型の変数
	kv.Set("key", "value")
	value, ok := kv.Get("key")
	if ok {
		fmt.Println(value)
	}
}
