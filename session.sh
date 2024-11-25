#!/bin/zsh

session="bulldogs"

tmux new-session -d -s $session

tmux rename-window -t 1 'Main'
tmux new-window -t $session:2 -n 'Cli'
tmux send-keys -t 'Main' 'zsh' C-m 'clear' C-m
tmux attach -t bulldogs
