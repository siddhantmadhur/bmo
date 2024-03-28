use std::{fs, io::Write};
use std::fs::File;
use serde::{Deserialize, Serialize};



#[derive(Deserialize, Serialize)]
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


pub fn init() {
    let cfg: Config = Config{
        build_commands: vec![],
        final_build_command: String::from(""),
        start_process: String::from(""),
        exclude_regex: vec![],
        exclude_directories: vec![]
    };

    let toml_cfg = toml::to_string(&cfg).unwrap();

    let mut f = File::create(".bmo.toml").unwrap();
    let _ = f.write_all(toml_cfg.as_bytes());

}

impl Config {


}
