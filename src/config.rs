use std::fs;
use serde::Deserialize;


#[derive(Deserialize)]
pub struct Config {
    pub build_commands: Vec<String>,
    pub final_build_command: String,
    pub start_process: String,
    pub exclude_regex: Vec<String>,
    pub exclude_directories: Vec<String>
}

pub fn new() -> Config  {
    let cfg = fs::read_to_string(".bmo.toml").unwrap();

    let config: Config = toml::from_str(cfg.as_str()).unwrap();

    config
}



impl Config {



}
