namespace go example

struct Person {
    1: required string name,
    2: optional i32 age,
}

service format_data {
    Person sayHello(1:string name),
}