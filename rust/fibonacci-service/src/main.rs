#![feature(proc_macro_hygiene, decl_macro)]

#[macro_use] extern crate rocket;
extern crate fibonacci;

#[get("/<position>")]
fn index(position: u32) -> String {
    return fibonacci::fibonacci(position).to_string();
}

fn main() {
    rocket::ignite().mount("/", routes![index]).launch();
}
