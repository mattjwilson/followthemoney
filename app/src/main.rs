mod setup;
use std::io::{self, Read, Write};

fn main() {
    let config_file_name = String::from("config.mjw");
    println!("Follow the money, the app for looking at what our elected officials are actually doing with our money.");

    let mut config = setup::Config::new(&config_file_name);
}

fn print_output(message: String) {
    println!("{}", message);
    io::stdout()
        .flush()
        .unwrap(); // unsafe, but I'm ok with it here.
}
