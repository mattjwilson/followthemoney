use std::fs;
use std::path::Path;


pub struct Config {
    pub key: String,
    pub is_setup: bool,
}

impl Config {
    pub fn new(filename_and_path: &String) -> Config {
        let mut config = Config {
            key: String::from(""),
            is_setup: false,
        };

        if !Path::new(filename_and_path).exists() {
            return config;
        }

        let contents = fs::read_to_string(filename_and_path).unwrap();

        for line in contents.lines() {
            if line.contains("key") {
                config.key = String::from(line);
                config.set_is_setup();
            } else {
                println!("Unexpected config: {}", line);
            }
        }

        config
    }

    fn set_is_setup(&mut self) {
        if !self.is_setup {
            self.is_setup = true;
        }
    }
}