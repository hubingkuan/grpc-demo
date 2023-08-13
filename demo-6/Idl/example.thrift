// 生成的目录名称和package 名称
namespace go example

//include "tweet.thrift"

struct Person {
    1: required string name,
    2: optional i32 age,
}

service format_data {
    void ping(),
    Person sayHello(1:string name),
}