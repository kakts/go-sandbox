package main

import (
    "container/list"
    "fmt"
)

func main() {
    // create a new list
    l := list.New()
    l.PushBack(1) // l {1}

    // list mに要素を2つ追加
    m := list.New()
    m.PushBack(2) // m {2}
    m.PushBack(3) // m {2, 3}
    fmt.Printf("Length of m is %d\n", m.Len())

    // lの後ろにmをつなげる
    l.PushBackList(m)

    // Iterate through list and print its contents.
    // 1 2 3と表示される
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

    // l の先頭に要素追加
    l.PushFront(5) // l {5, 1, 2, 3}

    // l の先頭にmの要素をつなげる
    l.PushFrontList(m) // l {2, 3, 5, 1 , 2, 3}

    // Iterate through list and print its contents.
    // 2 3 5 1 2 3 と表示される
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

    // 最後の要素を削除する
    l.Remove(l.Back())
    // Iterate through list and print its contents.
    // 2 3 5 1 2 と表示される
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

}
