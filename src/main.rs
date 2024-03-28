use std::{env::{self, args}, path::Path, process::{Child, Command}, thread, time::Duration};

use config::Config;
use notify::{event::DataChange, Result, Watcher};
use regex::Regex;
use walkdir::{DirEntry, WalkDir};

mod config;

fn is_hidden(entry: &DirEntry, excluded_dir: &Vec<String>) -> bool {
    entry.file_name()
         .to_str()
         .map(|s| excluded_dir.contains(&s.to_string()))
         .unwrap_or(false)
}


fn run_command(cmd: &str) -> Child {
    let argument:Vec<&str> = cmd.split(" ").collect::<Vec<&str>>();
    Command::new(&argument[0])
        .args(&argument[1..])
        .spawn()
        .unwrap()
}

fn main() {


    let arg = args().collect::<Vec<String>>();

    match arg[1].as_str() {
        "init" => config::init(),
        "run" => {

            let cfg = config::new();


            for cmd in cfg.build_commands {
                let mut proc = run_command(&cmd);
                let _ = proc.wait();
            }

            run_command(&cfg.final_build_command);

            let mut command = run_command(&cfg.start_process);


            let event_fn = move | res: Result<notify::Event>| {
                match res {
                    Ok(event) => {
                        match event.kind {
                            notify::EventKind::Modify(notify::event::ModifyKind::Data(_)) => {
                                
                                
                                let temp_cfg = config::new();

                                for reg in temp_cfg.exclude_regex.into_iter() {
                                   let re = Regex::new(&reg).unwrap();
                                   for cur in event.paths.clone().into_iter() {
                                        if re.is_match(&cur.to_str().unwrap()) {
                                            return ();
                                        }
                                   }
                                }


                                command.kill().expect("command couldn't be killed");
                                command.wait().unwrap();
                                println!("Build Cmds: {:?}", &temp_cfg.build_commands);

                                for cmd in temp_cfg.build_commands.into_iter() {
                                    let mut proc = run_command(&cmd);
                                    let _ = proc.wait();
                                }


                                let mut build = run_command(&cfg.final_build_command.clone());
                                let _ = build.wait();

                                command = run_command(&cfg.start_process.clone()); 

                                println!("event: {:?}", event)
                            },
                            _ => ()
                        }


                    },
                    Err(e) => println!("watch error: {:?}", e),
                }
            };

            let mut watcher = notify::recommended_watcher(event_fn)
                .unwrap();


            for entry in WalkDir::new(env::current_dir().unwrap()).into_iter().filter_entry(|e| !is_hidden(e, &cfg.exclude_directories)) {
                let path = entry.unwrap();
                println!("Path: {}", path.path().display());
                let _ = watcher.watch(path.path(), notify::RecursiveMode::NonRecursive);
            }

            loop {}
        }
        _ => println!("Argument not recognized")
    }



}
